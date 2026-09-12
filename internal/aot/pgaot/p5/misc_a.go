package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AdvanceXLInsertBuffer(m *base.Module, l0 int64, l1 int32, l2 int32) {
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v36 int32
	_ = v36
	var v43 int64
	_ = v43
	var v47 int32
	_ = v47
	var v51 int64
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int64
	_ = v57
	var v60 int64
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int64
	_ = v78
	var v83 int64
	_ = v83
	var v88 int32
	_ = v88
	var v89 int64
	_ = v89
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int64
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int64
	_ = v111
	var v116 int32
	_ = v116
	var v117 int64
	_ = v117
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int64
	_ = v128
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 int64
	_ = v148
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int64
	_ = v161
	var v167 int32
	_ = v167
	var v168 int64
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v204 int32
	_ = v204
	var v205 int64
	_ = v205
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v217 int64
	_ = v217
	var v220 int32
	_ = v220
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	v13 = m.G0
	v15 = v13 - int32(32)
	m.G0 = v15
	v18 = *(*int32)(unsafe.Add(mBase, _consts[114]))
	v20 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v24 = F_LWLockAcquire(m, v20+int32(896), int32(0))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _consts[114]))
	v30 = *(*int64)(unsafe.Add(mBase, uint32(v29)+288))
	if base.B2i32(l2 == int32(0))&base.B2i32(base.Ui64(l0) < base.Ui64(v30)) != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v236 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	F_LWLockRelease(m, v236+int32(896))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L39
	}
L4:
	;
	v36 = v29
	v43 = v30
	goto L5
L5:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v36)+304))
	v51 = base.I64_rem_u_s(int64(base.Ui64(v43)>>(uint(int64(13))%64)), base.I64_extend_i32_s(v47+int32(1)))
	v52 = base.I32_wrap_i64(v51)
	v54 = v52 << (uint(int32(3)) % 32)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v36)+300))
	v56 = v54 + v55
	v57 = *(*int64)(unsafe.Add(mBase, uint32(v56)))
	*(*int64)(unsafe.Add(mBase, uint32(v56))) = v57
	v60 = *(*int64)(unsafe.Add(mBase, _consts[118]))
	if base.Ui64(v57) <= base.Ui64(v60) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L3
L7:
	;
	v167 = *(*int32)(unsafe.Add(mBase, _consts[114]))
	v168 = *(*int64)(unsafe.Add(mBase, uint32(v167)+288))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v167)+296))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v167)+300))
	*(*int64)(unsafe.Add(mBase, uint32(v170+v54))) = int64(0)
	v174 = int32(2)
	v177 = v169 + v52<<(uint(int32(13))%32)
	v183 = F__emscripten_memset_bulkmem(m, v177+v174, base.I32_extend8_s(int32(0)), int32(8190))
	mBase = m.M
	goto L30
L8:
	;
	if l2 != 0 {
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _consts[114]))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+440))
	*(*int32)(unsafe.Add(mBase, uint32(v63)+440)) = int32(1)
	if v64 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v68 = *(*int32)(unsafe.Add(mBase, _consts[114]))
	F_s_lock(m, v68+int32(440), int32(510097), int32(2025), int32(230953))
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
	v77 = *(*int32)(unsafe.Add(mBase, _consts[114]))
	v78 = *(*int64)(unsafe.Add(mBase, uint32(v77)+184))
	if base.Ui64(v78) < base.Ui64(v57) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L12
L14:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v77)+184)) = v57
	goto L16
L15:
	;
	goto L16
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v77)+440)) = int32(0)
	v83 = *(*int64)(unsafe.Add(mBase, uint32(v77)+280))
	*(*int64)(unsafe.Add(mBase, uint32(v77)+280)) = v83
	*(*int64)(unsafe.Add(mBase, _consts[119])) = v83
	v88 = *(*int32)(unsafe.Add(mBase, _consts[114]))
	v89 = *(*int64)(unsafe.Add(mBase, uint32(v88)+272))
	*(*int64)(unsafe.Add(mBase, uint32(v88)+272)) = v89
	*(*int64)(unsafe.Add(mBase, _consts[118])) = v89
	if base.Ui64(v57) <= base.Ui64(v89) {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	v95 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	F_LWLockRelease(m, v95+int32(896))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v100 = F_WaitXLogInsertionsToFinish(m, v57)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v103 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v107 = F_LWLockAcquire(m, v103+int32(1024), int32(0))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v109 = int32(4431840)
	v110 = *(*int32)(unsafe.Add(mBase, _consts[114]))
	v111 = *(*int64)(unsafe.Add(mBase, uint32(v110)+280))
	*(*int64)(unsafe.Add(mBase, uint32(v110)+280)) = v111
	*(*int64)(unsafe.Add(mBase, _consts[119])) = v111
	v116 = *(*int32)(unsafe.Add(mBase, _consts[114]))
	v117 = *(*int64)(unsafe.Add(mBase, uint32(v116)+272))
	*(*int64)(unsafe.Add(mBase, uint32(v116)+272)) = v117
	*(*int64)(unsafe.Add(mBase, _consts[118])) = v117
	if base.Ui64(v57) <= base.Ui64(v117) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v153 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v157 = F_LWLockAcquire(m, v153+int32(896), int32(0))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L28
	}
L22:
	;
	v123 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	F_LWLockRelease(m, v123+int32(1024))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v128 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+24)) = v128
	*(*int64)(unsafe.Add(mBase, uint32(v15)+8)) = v128
	*(*int64)(unsafe.Add(mBase, uint32(v15)+16)) = v57
	*(*int64)(unsafe.Add(mBase, uint32(v15))) = v57
	F_XLogWrite(m, v15, l1, int32(0))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L26
	}
L25:
	;
	goto L21
L26:
	;
	v138 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	F_LWLockRelease(m, v138+int32(1024))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v144 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[90])) = uint8(v144)
	v146 = int32(4434688)
	v148 = *(*int64)(unsafe.Add(mBase, _consts[120]))
	*(*int64)(unsafe.Add(mBase, _consts[120])) = v148 + int64(1)
	goto L21
L28:
	;
	v160 = *(*int32)(unsafe.Add(mBase, _consts[114]))
	v161 = *(*int64)(unsafe.Add(mBase, uint32(v160)+288))
	if base.Ui64(v161) <= base.Ui64(l0) {
		v36 = v160
		v43 = v161
		goto L5
	} else {
		goto L29
	}
L29:
	;
	goto L3
L30:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v177)+8)) = v168
	*(*int32)(unsafe.Add(mBase, uint32(v177)+4)) = l1
	v186 = int32(53528)
	*(*uint16)(unsafe.Add(mBase, uint32(v177))) = uint16(v186)
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v18)+164))
	if v188 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v191 = int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v177)+2)) = uint16(v191)
	v194 = int32(6)
	goto L33
L32:
	;
	v194 = v174
	goto L33
L33:
	;
	v196 = *(*int32)(unsafe.Add(mBase, _consts[117]))
	if v168&base.I64_extend_i32_s(v196-int32(1)) == int64(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v204 = *(*int32)(unsafe.Add(mBase, _consts[121]))
	v205 = *(*int64)(unsafe.Add(mBase, uint32(v204)))
	*(*int32)(unsafe.Add(mBase, uint32(v177)+36)) = int32(8192)
	*(*int32)(unsafe.Add(mBase, uint32(v177)+32)) = v196
	*(*int64)(unsafe.Add(mBase, uint32(v177)+24)) = v205
	*(*uint16)(unsafe.Add(mBase, uint32(v177)+2)) = uint16(v194)
	goto L36
L35:
	;
	goto L36
L36:
	;
	v212 = int32(4431840)
	v213 = *(*int32)(unsafe.Add(mBase, _consts[114]))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v213)+300))
	v217 = v168 - int64(-8192)
	*(*int64)(unsafe.Add(mBase, uint32(v214+v54))) = v217
	v220 = *(*int32)(unsafe.Add(mBase, _consts[114]))
	*(*int64)(unsafe.Add(mBase, uint32(v220)+288)) = v217
	if l2 != 0 {
		v36 = v220
		v43 = v217
		goto L5
	} else {
		goto L37
	}
L37:
	;
	if base.Ui64(v217) <= base.Ui64(l0) {
		v36 = v220
		v43 = v217
		goto L5
	} else {
		goto L38
	}
L38:
	;
	goto L6
L39:
	;
	m.G0 = v15 + int32(32)
	return
}
func F_AllocSetAllocLarge(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
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
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v53 int32
	_ = v53
	if base.Ui32(int32(1073741824)) <= base.Ui32(l1) {
		if l1 < int32(0) {
			F_MemoryContextSizeFailure(m, l1)
			mBase = m.M
			v53 = m.ExcPending
			if v53 != 0 {
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
				v53 = m.ExcPending
				if v53 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v19 = (l1+int32(7))&int32(-8) + int32(32)
				v20 = F_emscripten_builtin_malloc(m, v19)
				mBase = m.M
				if v20 == int32(0) {
					v23 = F_MemoryContextAllocationFailure(m, l0, l1, l2)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						return v23
					}
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v28 + v19
					v31 = v20 + v19
					*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v31
					*(*int32)(unsafe.Add(mBase, uint32(v20))) = l0
					*(*int64)(unsafe.Add(mBase, uint32(v20)+24)) = int64(-5645020766237429837)
					*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v31
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
					if v37 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v37
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v39
						if v39 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v39)+4)) = v20
							v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
							v43 = v42
						} else {
							v43 = v37
						}
						*(*int32)(unsafe.Add(mBase, uint32(v43)+8)) = v20
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v20)+4)) = int64(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v20
					}
					return v20 + int32(32)
				}
			}
		}
	} else {
		v19 = (l1+int32(7))&int32(-8) + int32(32)
		v20 = F_emscripten_builtin_malloc(m, v19)
		mBase = m.M
		if v20 == int32(0) {
			v23 = F_MemoryContextAllocationFailure(m, l0, l1, l2)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				return v23
			}
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v28 + v19
			v31 = v20 + v19
			*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v31
			*(*int32)(unsafe.Add(mBase, uint32(v20))) = l0
			*(*int64)(unsafe.Add(mBase, uint32(v20)+24)) = int64(-5645020766237429837)
			*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v31
			v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
			if v37 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v37
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v39
				if v39 != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(v39)+4)) = v20
					v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
					v43 = v42
				} else {
					v43 = v37
				}
				*(*int32)(unsafe.Add(mBase, uint32(v43)+8)) = v20
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v20)+4)) = int64(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v20
			}
			return v20 + int32(32)
		}
	}
}
func F_AllocSetContextCreateInternal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v46 int64
	_ = v46
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v137 int32
	_ = v137
	var v151 int32
	_ = v151
	var v167 int32
	_ = v167
	var v170 int64
	_ = v170
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	if l2 != 0 {
		v18 = int32(-1)
		if l2 != 0 {
			v73 = v18
			if l2 != 0 {
				v75 = l2
			} else {
				v75 = l3
			}
			if base.Ui32(v75) <= base.Ui32(int32(144)) {
				v78 = int32(144)
			} else {
				v78 = v75
			}
			v79 = F_emscripten_builtin_malloc(m, v78)
			mBase = m.M
			if v79 == int32(0) {
				v83 = *(*int32)(unsafe.Add(mBase, _consts[147]))
				if v83 != 0 {
					F_MemoryContextStats(m, v83)
					mBase = m.M
					v87 = m.ExcPending
					if v87 != 0 {
						return int32(0)
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v91 = m.ExcPending
						if v91 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(8389))
							mBase = m.M
							v94 = m.ExcPending
							if v94 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(13961), int32(0))
								mBase = m.M
								v98 = m.ExcPending
								if v98 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v13))) = l1
									F_errdetail(m, int32(678867), v13)
									mBase = m.M
									v102 = m.ExcPending
									if v102 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(504340), int32(453), int32(319332))
										mBase = m.M
										v107 = m.ExcPending
										if v107 != 0 {
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
					v91 = m.ExcPending
					if v91 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(8389))
						mBase = m.M
						v94 = m.ExcPending
						if v94 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(13961), int32(0))
							mBase = m.M
							v98 = m.ExcPending
							if v98 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v13))) = l1
								F_errdetail(m, int32(678867), v13)
								mBase = m.M
								v102 = m.ExcPending
								if v102 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(504340), int32(453), int32(319332))
									mBase = m.M
									v107 = m.ExcPending
									if v107 != 0 {
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
				*(*int32)(unsafe.Add(mBase, uint32(v79)+128)) = v79 + v78
				*(*int32)(unsafe.Add(mBase, uint32(v79)+124)) = v79 + int32(136)
				v114 = v79 + int32(112)
				*(*int32)(unsafe.Add(mBase, uint32(v114))) = v79
				*(*int64)(unsafe.Add(mBase, uint32(v79)+116)) = int64(0)
				*(*int32)(unsafe.Add(mBase, uint32(v79)+44)) = v114
				v120 = v79 + int32(48)
				v122 = v79 + int32(92)
				if base.Ui32(v120) < base.Ui32(v122) {
					v126 = v79 + int32(52)
					if base.Ui32(v126) < base.Ui32(v122) {
						v128 = v122
					} else {
						v128 = v126
					}
					v137 = F__emscripten_memset_bulkmem(m, v120, base.I32_extend8_s(int32(0)), (v128-v79-int32(49))&int32(-4)+int32(4))
					mBase = m.M
				} else {
				}
				*(*int32)(unsafe.Add(mBase, uint32(v79)+108)) = v73
				*(*int32)(unsafe.Add(mBase, uint32(v79)+100)) = l3
				*(*int32)(unsafe.Add(mBase, uint32(v79)+96)) = l4
				*(*int32)(unsafe.Add(mBase, uint32(v122))) = l3
				v151 = int32(8192)
				for {
					if base.Ui32(int32(base.Ui32(l4-int32(24))>>(uint(int32(2))%32))) < base.Ui32(v151+int32(8)) {
						v151 = int32(base.Ui32(v151) >> (uint(int32(1)) % 32))
						continue
					} else {
						break
					}
					break
				}
				*(*int32)(unsafe.Add(mBase, uint32(v79)+104)) = v151
				*(*int32)(unsafe.Add(mBase, uint32(v79)+16)) = l0
				v167 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v79)+4)) = uint8(v167)
				*(*int32)(unsafe.Add(mBase, uint32(v79))) = int32(474)
				v170 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v79)+36)) = v170
				*(*int32)(unsafe.Add(mBase, uint32(v79)+32)) = l1
				*(*int64)(unsafe.Add(mBase, uint32(v79)+20)) = v170
				*(*int32)(unsafe.Add(mBase, uint32(v79)+8)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v79)+12)) = int32(1785164)
				if l0 != 0 {
					v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v79)+28)) = v182
					if v182 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v182)+24)) = v79
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v79
					v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
					*(*uint8)(unsafe.Add(mBase, uint32(v79)+5)) = uint8(v186)
				} else {
					v188 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v79)+28)) = v188
					*(*uint8)(unsafe.Add(mBase, uint32(v79)+5)) = uint8(v188)
				}
				v198 = v79
				v199 = v78
				*(*int32)(unsafe.Add(mBase, uint32(v198)+8)) = v199
				m.G0 = v13 + int32(16)
				return v198
			}
		} else {
			if l3 != int32(1024) {
				v73 = v18
				if l2 != 0 {
					v75 = l2
				} else {
					v75 = l3
				}
				if base.Ui32(v75) <= base.Ui32(int32(144)) {
					v78 = int32(144)
				} else {
					v78 = v75
				}
				v79 = F_emscripten_builtin_malloc(m, v78)
				mBase = m.M
				if v79 == int32(0) {
					v83 = *(*int32)(unsafe.Add(mBase, _consts[147]))
					if v83 != 0 {
						F_MemoryContextStats(m, v83)
						mBase = m.M
						v87 = m.ExcPending
						if v87 != 0 {
							return int32(0)
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v91 = m.ExcPending
							if v91 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(8389))
								mBase = m.M
								v94 = m.ExcPending
								if v94 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(13961), int32(0))
									mBase = m.M
									v98 = m.ExcPending
									if v98 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v13))) = l1
										F_errdetail(m, int32(678867), v13)
										mBase = m.M
										v102 = m.ExcPending
										if v102 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(504340), int32(453), int32(319332))
											mBase = m.M
											v107 = m.ExcPending
											if v107 != 0 {
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
						v91 = m.ExcPending
						if v91 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(8389))
							mBase = m.M
							v94 = m.ExcPending
							if v94 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(13961), int32(0))
								mBase = m.M
								v98 = m.ExcPending
								if v98 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v13))) = l1
									F_errdetail(m, int32(678867), v13)
									mBase = m.M
									v102 = m.ExcPending
									if v102 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(504340), int32(453), int32(319332))
										mBase = m.M
										v107 = m.ExcPending
										if v107 != 0 {
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
					*(*int32)(unsafe.Add(mBase, uint32(v79)+128)) = v79 + v78
					*(*int32)(unsafe.Add(mBase, uint32(v79)+124)) = v79 + int32(136)
					v114 = v79 + int32(112)
					*(*int32)(unsafe.Add(mBase, uint32(v114))) = v79
					*(*int64)(unsafe.Add(mBase, uint32(v79)+116)) = int64(0)
					*(*int32)(unsafe.Add(mBase, uint32(v79)+44)) = v114
					v120 = v79 + int32(48)
					v122 = v79 + int32(92)
					if base.Ui32(v120) < base.Ui32(v122) {
						v126 = v79 + int32(52)
						if base.Ui32(v126) < base.Ui32(v122) {
							v128 = v122
						} else {
							v128 = v126
						}
						v137 = F__emscripten_memset_bulkmem(m, v120, base.I32_extend8_s(int32(0)), (v128-v79-int32(49))&int32(-4)+int32(4))
						mBase = m.M
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(v79)+108)) = v73
					*(*int32)(unsafe.Add(mBase, uint32(v79)+100)) = l3
					*(*int32)(unsafe.Add(mBase, uint32(v79)+96)) = l4
					*(*int32)(unsafe.Add(mBase, uint32(v122))) = l3
					v151 = int32(8192)
					for {
						if base.Ui32(int32(base.Ui32(l4-int32(24))>>(uint(int32(2))%32))) < base.Ui32(v151+int32(8)) {
							v151 = int32(base.Ui32(v151) >> (uint(int32(1)) % 32))
							continue
						} else {
							break
						}
						break
					}
					*(*int32)(unsafe.Add(mBase, uint32(v79)+104)) = v151
					*(*int32)(unsafe.Add(mBase, uint32(v79)+16)) = l0
					v167 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v79)+4)) = uint8(v167)
					*(*int32)(unsafe.Add(mBase, uint32(v79))) = int32(474)
					v170 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v79)+36)) = v170
					*(*int32)(unsafe.Add(mBase, uint32(v79)+32)) = l1
					*(*int64)(unsafe.Add(mBase, uint32(v79)+20)) = v170
					*(*int32)(unsafe.Add(mBase, uint32(v79)+8)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v79)+12)) = int32(1785164)
					if l0 != 0 {
						v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						*(*int32)(unsafe.Add(mBase, uint32(v79)+28)) = v182
						if v182 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v182)+24)) = v79
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v79
						v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
						*(*uint8)(unsafe.Add(mBase, uint32(v79)+5)) = uint8(v186)
					} else {
						v188 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v79)+28)) = v188
						*(*uint8)(unsafe.Add(mBase, uint32(v79)+5)) = uint8(v188)
					}
					v198 = v79
					v199 = v78
					*(*int32)(unsafe.Add(mBase, uint32(v198)+8)) = v199
					m.G0 = v13 + int32(16)
					return v198
				}
			} else {
				v23 = int32(1)
				v25 = v23 << (uint(int32(3)) % 32)
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_consts[1297])))
				if v28 == int32(0) {
					v73 = v23
					if l2 != 0 {
						v75 = l2
					} else {
						v75 = l3
					}
					if base.Ui32(v75) <= base.Ui32(int32(144)) {
						v78 = int32(144)
					} else {
						v78 = v75
					}
					v79 = F_emscripten_builtin_malloc(m, v78)
					mBase = m.M
					if v79 == int32(0) {
						v83 = *(*int32)(unsafe.Add(mBase, _consts[147]))
						if v83 != 0 {
							F_MemoryContextStats(m, v83)
							mBase = m.M
							v87 = m.ExcPending
							if v87 != 0 {
								return int32(0)
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v91 = m.ExcPending
								if v91 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(8389))
									mBase = m.M
									v94 = m.ExcPending
									if v94 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(13961), int32(0))
										mBase = m.M
										v98 = m.ExcPending
										if v98 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v13))) = l1
											F_errdetail(m, int32(678867), v13)
											mBase = m.M
											v102 = m.ExcPending
											if v102 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(504340), int32(453), int32(319332))
												mBase = m.M
												v107 = m.ExcPending
												if v107 != 0 {
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
							v91 = m.ExcPending
							if v91 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(8389))
								mBase = m.M
								v94 = m.ExcPending
								if v94 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(13961), int32(0))
									mBase = m.M
									v98 = m.ExcPending
									if v98 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v13))) = l1
										F_errdetail(m, int32(678867), v13)
										mBase = m.M
										v102 = m.ExcPending
										if v102 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(504340), int32(453), int32(319332))
											mBase = m.M
											v107 = m.ExcPending
											if v107 != 0 {
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
						*(*int32)(unsafe.Add(mBase, uint32(v79)+128)) = v79 + v78
						*(*int32)(unsafe.Add(mBase, uint32(v79)+124)) = v79 + int32(136)
						v114 = v79 + int32(112)
						*(*int32)(unsafe.Add(mBase, uint32(v114))) = v79
						*(*int64)(unsafe.Add(mBase, uint32(v79)+116)) = int64(0)
						*(*int32)(unsafe.Add(mBase, uint32(v79)+44)) = v114
						v120 = v79 + int32(48)
						v122 = v79 + int32(92)
						if base.Ui32(v120) < base.Ui32(v122) {
							v126 = v79 + int32(52)
							if base.Ui32(v126) < base.Ui32(v122) {
								v128 = v122
							} else {
								v128 = v126
							}
							v137 = F__emscripten_memset_bulkmem(m, v120, base.I32_extend8_s(int32(0)), (v128-v79-int32(49))&int32(-4)+int32(4))
							mBase = m.M
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(v79)+108)) = v73
						*(*int32)(unsafe.Add(mBase, uint32(v79)+100)) = l3
						*(*int32)(unsafe.Add(mBase, uint32(v79)+96)) = l4
						*(*int32)(unsafe.Add(mBase, uint32(v122))) = l3
						v151 = int32(8192)
						for {
							if base.Ui32(int32(base.Ui32(l4-int32(24))>>(uint(int32(2))%32))) < base.Ui32(v151+int32(8)) {
								v151 = int32(base.Ui32(v151) >> (uint(int32(1)) % 32))
								continue
							} else {
								break
							}
							break
						}
						*(*int32)(unsafe.Add(mBase, uint32(v79)+104)) = v151
						*(*int32)(unsafe.Add(mBase, uint32(v79)+16)) = l0
						v167 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v79)+4)) = uint8(v167)
						*(*int32)(unsafe.Add(mBase, uint32(v79))) = int32(474)
						v170 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v79)+36)) = v170
						*(*int32)(unsafe.Add(mBase, uint32(v79)+32)) = l1
						*(*int64)(unsafe.Add(mBase, uint32(v79)+20)) = v170
						*(*int32)(unsafe.Add(mBase, uint32(v79)+8)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v79)+12)) = int32(1785164)
						if l0 != 0 {
							v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v79)+28)) = v182
							if v182 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v182)+24)) = v79
							} else {
							}
							*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v79
							v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
							*(*uint8)(unsafe.Add(mBase, uint32(v79)+5)) = uint8(v186)
						} else {
							v188 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v79)+28)) = v188
							*(*uint8)(unsafe.Add(mBase, uint32(v79)+5)) = uint8(v188)
						}
						v198 = v79
						v199 = v78
						*(*int32)(unsafe.Add(mBase, uint32(v198)+8)) = v199
						m.G0 = v13 + int32(16)
						return v198
					}
				} else {
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v28)+28))
					*(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_consts[1297]))) = v33
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_consts[1298])))
					v36 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_consts[1298]))) = v35 - v36
					*(*int32)(unsafe.Add(mBase, uint32(v28)+96)) = l4
					*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = l0
					*(*uint8)(unsafe.Add(mBase, uint32(v28)+4)) = uint8(v36)
					*(*int32)(unsafe.Add(mBase, uint32(v28))) = int32(474)
					v46 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v28)+36)) = v46
					*(*int32)(unsafe.Add(mBase, uint32(v28)+32)) = l1
					*(*int64)(unsafe.Add(mBase, uint32(v28)+20)) = v46
					*(*int32)(unsafe.Add(mBase, uint32(v28)+8)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v28)+12)) = int32(1785164)
					if l0 != 0 {
						v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						*(*int32)(unsafe.Add(mBase, uint32(v28)+28)) = v58
						if v58 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v58)+24)) = v28
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v28
						v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
						*(*uint8)(unsafe.Add(mBase, uint32(v28)+5)) = uint8(v62)
					} else {
						v64 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v28)+28)) = v64
						*(*uint8)(unsafe.Add(mBase, uint32(v28)+5)) = uint8(v64)
					}
					v69 = *(*int32)(unsafe.Add(mBase, uint32(v28)+128))
					v198 = v28
					v199 = v69 - v28
					*(*int32)(unsafe.Add(mBase, uint32(v198)+8)) = v199
					m.G0 = v13 + int32(16)
					return v198
				}
			}
		}
	} else {
		if l3 != int32(8192) {
			v18 = int32(-1)
			if l2 != 0 {
				v73 = v18
				if l2 != 0 {
					v75 = l2
				} else {
					v75 = l3
				}
				if base.Ui32(v75) <= base.Ui32(int32(144)) {
					v78 = int32(144)
				} else {
					v78 = v75
				}
				v79 = F_emscripten_builtin_malloc(m, v78)
				mBase = m.M
				if v79 == int32(0) {
					v83 = *(*int32)(unsafe.Add(mBase, _consts[147]))
					if v83 != 0 {
						F_MemoryContextStats(m, v83)
						mBase = m.M
						v87 = m.ExcPending
						if v87 != 0 {
							return int32(0)
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v91 = m.ExcPending
							if v91 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(8389))
								mBase = m.M
								v94 = m.ExcPending
								if v94 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(13961), int32(0))
									mBase = m.M
									v98 = m.ExcPending
									if v98 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v13))) = l1
										F_errdetail(m, int32(678867), v13)
										mBase = m.M
										v102 = m.ExcPending
										if v102 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(504340), int32(453), int32(319332))
											mBase = m.M
											v107 = m.ExcPending
											if v107 != 0 {
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
						v91 = m.ExcPending
						if v91 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(8389))
							mBase = m.M
							v94 = m.ExcPending
							if v94 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(13961), int32(0))
								mBase = m.M
								v98 = m.ExcPending
								if v98 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v13))) = l1
									F_errdetail(m, int32(678867), v13)
									mBase = m.M
									v102 = m.ExcPending
									if v102 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(504340), int32(453), int32(319332))
										mBase = m.M
										v107 = m.ExcPending
										if v107 != 0 {
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
					*(*int32)(unsafe.Add(mBase, uint32(v79)+128)) = v79 + v78
					*(*int32)(unsafe.Add(mBase, uint32(v79)+124)) = v79 + int32(136)
					v114 = v79 + int32(112)
					*(*int32)(unsafe.Add(mBase, uint32(v114))) = v79
					*(*int64)(unsafe.Add(mBase, uint32(v79)+116)) = int64(0)
					*(*int32)(unsafe.Add(mBase, uint32(v79)+44)) = v114
					v120 = v79 + int32(48)
					v122 = v79 + int32(92)
					if base.Ui32(v120) < base.Ui32(v122) {
						v126 = v79 + int32(52)
						if base.Ui32(v126) < base.Ui32(v122) {
							v128 = v122
						} else {
							v128 = v126
						}
						v137 = F__emscripten_memset_bulkmem(m, v120, base.I32_extend8_s(int32(0)), (v128-v79-int32(49))&int32(-4)+int32(4))
						mBase = m.M
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(v79)+108)) = v73
					*(*int32)(unsafe.Add(mBase, uint32(v79)+100)) = l3
					*(*int32)(unsafe.Add(mBase, uint32(v79)+96)) = l4
					*(*int32)(unsafe.Add(mBase, uint32(v122))) = l3
					v151 = int32(8192)
					for {
						if base.Ui32(int32(base.Ui32(l4-int32(24))>>(uint(int32(2))%32))) < base.Ui32(v151+int32(8)) {
							v151 = int32(base.Ui32(v151) >> (uint(int32(1)) % 32))
							continue
						} else {
							break
						}
						break
					}
					*(*int32)(unsafe.Add(mBase, uint32(v79)+104)) = v151
					*(*int32)(unsafe.Add(mBase, uint32(v79)+16)) = l0
					v167 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v79)+4)) = uint8(v167)
					*(*int32)(unsafe.Add(mBase, uint32(v79))) = int32(474)
					v170 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v79)+36)) = v170
					*(*int32)(unsafe.Add(mBase, uint32(v79)+32)) = l1
					*(*int64)(unsafe.Add(mBase, uint32(v79)+20)) = v170
					*(*int32)(unsafe.Add(mBase, uint32(v79)+8)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v79)+12)) = int32(1785164)
					if l0 != 0 {
						v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						*(*int32)(unsafe.Add(mBase, uint32(v79)+28)) = v182
						if v182 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v182)+24)) = v79
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v79
						v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
						*(*uint8)(unsafe.Add(mBase, uint32(v79)+5)) = uint8(v186)
					} else {
						v188 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v79)+28)) = v188
						*(*uint8)(unsafe.Add(mBase, uint32(v79)+5)) = uint8(v188)
					}
					v198 = v79
					v199 = v78
					*(*int32)(unsafe.Add(mBase, uint32(v198)+8)) = v199
					m.G0 = v13 + int32(16)
					return v198
				}
			} else {
				if l3 != int32(1024) {
					v73 = v18
					if l2 != 0 {
						v75 = l2
					} else {
						v75 = l3
					}
					if base.Ui32(v75) <= base.Ui32(int32(144)) {
						v78 = int32(144)
					} else {
						v78 = v75
					}
					v79 = F_emscripten_builtin_malloc(m, v78)
					mBase = m.M
					if v79 == int32(0) {
						v83 = *(*int32)(unsafe.Add(mBase, _consts[147]))
						if v83 != 0 {
							F_MemoryContextStats(m, v83)
							mBase = m.M
							v87 = m.ExcPending
							if v87 != 0 {
								return int32(0)
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v91 = m.ExcPending
								if v91 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(8389))
									mBase = m.M
									v94 = m.ExcPending
									if v94 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(13961), int32(0))
										mBase = m.M
										v98 = m.ExcPending
										if v98 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v13))) = l1
											F_errdetail(m, int32(678867), v13)
											mBase = m.M
											v102 = m.ExcPending
											if v102 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(504340), int32(453), int32(319332))
												mBase = m.M
												v107 = m.ExcPending
												if v107 != 0 {
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
							v91 = m.ExcPending
							if v91 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(8389))
								mBase = m.M
								v94 = m.ExcPending
								if v94 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(13961), int32(0))
									mBase = m.M
									v98 = m.ExcPending
									if v98 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v13))) = l1
										F_errdetail(m, int32(678867), v13)
										mBase = m.M
										v102 = m.ExcPending
										if v102 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(504340), int32(453), int32(319332))
											mBase = m.M
											v107 = m.ExcPending
											if v107 != 0 {
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
						*(*int32)(unsafe.Add(mBase, uint32(v79)+128)) = v79 + v78
						*(*int32)(unsafe.Add(mBase, uint32(v79)+124)) = v79 + int32(136)
						v114 = v79 + int32(112)
						*(*int32)(unsafe.Add(mBase, uint32(v114))) = v79
						*(*int64)(unsafe.Add(mBase, uint32(v79)+116)) = int64(0)
						*(*int32)(unsafe.Add(mBase, uint32(v79)+44)) = v114
						v120 = v79 + int32(48)
						v122 = v79 + int32(92)
						if base.Ui32(v120) < base.Ui32(v122) {
							v126 = v79 + int32(52)
							if base.Ui32(v126) < base.Ui32(v122) {
								v128 = v122
							} else {
								v128 = v126
							}
							v137 = F__emscripten_memset_bulkmem(m, v120, base.I32_extend8_s(int32(0)), (v128-v79-int32(49))&int32(-4)+int32(4))
							mBase = m.M
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(v79)+108)) = v73
						*(*int32)(unsafe.Add(mBase, uint32(v79)+100)) = l3
						*(*int32)(unsafe.Add(mBase, uint32(v79)+96)) = l4
						*(*int32)(unsafe.Add(mBase, uint32(v122))) = l3
						v151 = int32(8192)
						for {
							if base.Ui32(int32(base.Ui32(l4-int32(24))>>(uint(int32(2))%32))) < base.Ui32(v151+int32(8)) {
								v151 = int32(base.Ui32(v151) >> (uint(int32(1)) % 32))
								continue
							} else {
								break
							}
							break
						}
						*(*int32)(unsafe.Add(mBase, uint32(v79)+104)) = v151
						*(*int32)(unsafe.Add(mBase, uint32(v79)+16)) = l0
						v167 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v79)+4)) = uint8(v167)
						*(*int32)(unsafe.Add(mBase, uint32(v79))) = int32(474)
						v170 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v79)+36)) = v170
						*(*int32)(unsafe.Add(mBase, uint32(v79)+32)) = l1
						*(*int64)(unsafe.Add(mBase, uint32(v79)+20)) = v170
						*(*int32)(unsafe.Add(mBase, uint32(v79)+8)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v79)+12)) = int32(1785164)
						if l0 != 0 {
							v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v79)+28)) = v182
							if v182 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v182)+24)) = v79
							} else {
							}
							*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v79
							v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
							*(*uint8)(unsafe.Add(mBase, uint32(v79)+5)) = uint8(v186)
						} else {
							v188 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v79)+28)) = v188
							*(*uint8)(unsafe.Add(mBase, uint32(v79)+5)) = uint8(v188)
						}
						v198 = v79
						v199 = v78
						*(*int32)(unsafe.Add(mBase, uint32(v198)+8)) = v199
						m.G0 = v13 + int32(16)
						return v198
					}
				} else {
					v23 = int32(1)
					v25 = v23 << (uint(int32(3)) % 32)
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_consts[1297])))
					if v28 == int32(0) {
						v73 = v23
						if l2 != 0 {
							v75 = l2
						} else {
							v75 = l3
						}
						if base.Ui32(v75) <= base.Ui32(int32(144)) {
							v78 = int32(144)
						} else {
							v78 = v75
						}
						v79 = F_emscripten_builtin_malloc(m, v78)
						mBase = m.M
						if v79 == int32(0) {
							v83 = *(*int32)(unsafe.Add(mBase, _consts[147]))
							if v83 != 0 {
								F_MemoryContextStats(m, v83)
								mBase = m.M
								v87 = m.ExcPending
								if v87 != 0 {
									return int32(0)
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v91 = m.ExcPending
									if v91 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(8389))
										mBase = m.M
										v94 = m.ExcPending
										if v94 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(13961), int32(0))
											mBase = m.M
											v98 = m.ExcPending
											if v98 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v13))) = l1
												F_errdetail(m, int32(678867), v13)
												mBase = m.M
												v102 = m.ExcPending
												if v102 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(504340), int32(453), int32(319332))
													mBase = m.M
													v107 = m.ExcPending
													if v107 != 0 {
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
								v91 = m.ExcPending
								if v91 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(8389))
									mBase = m.M
									v94 = m.ExcPending
									if v94 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(13961), int32(0))
										mBase = m.M
										v98 = m.ExcPending
										if v98 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v13))) = l1
											F_errdetail(m, int32(678867), v13)
											mBase = m.M
											v102 = m.ExcPending
											if v102 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(504340), int32(453), int32(319332))
												mBase = m.M
												v107 = m.ExcPending
												if v107 != 0 {
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
							*(*int32)(unsafe.Add(mBase, uint32(v79)+128)) = v79 + v78
							*(*int32)(unsafe.Add(mBase, uint32(v79)+124)) = v79 + int32(136)
							v114 = v79 + int32(112)
							*(*int32)(unsafe.Add(mBase, uint32(v114))) = v79
							*(*int64)(unsafe.Add(mBase, uint32(v79)+116)) = int64(0)
							*(*int32)(unsafe.Add(mBase, uint32(v79)+44)) = v114
							v120 = v79 + int32(48)
							v122 = v79 + int32(92)
							if base.Ui32(v120) < base.Ui32(v122) {
								v126 = v79 + int32(52)
								if base.Ui32(v126) < base.Ui32(v122) {
									v128 = v122
								} else {
									v128 = v126
								}
								v137 = F__emscripten_memset_bulkmem(m, v120, base.I32_extend8_s(int32(0)), (v128-v79-int32(49))&int32(-4)+int32(4))
								mBase = m.M
							} else {
							}
							*(*int32)(unsafe.Add(mBase, uint32(v79)+108)) = v73
							*(*int32)(unsafe.Add(mBase, uint32(v79)+100)) = l3
							*(*int32)(unsafe.Add(mBase, uint32(v79)+96)) = l4
							*(*int32)(unsafe.Add(mBase, uint32(v122))) = l3
							v151 = int32(8192)
							for {
								if base.Ui32(int32(base.Ui32(l4-int32(24))>>(uint(int32(2))%32))) < base.Ui32(v151+int32(8)) {
									v151 = int32(base.Ui32(v151) >> (uint(int32(1)) % 32))
									continue
								} else {
									break
								}
								break
							}
							*(*int32)(unsafe.Add(mBase, uint32(v79)+104)) = v151
							*(*int32)(unsafe.Add(mBase, uint32(v79)+16)) = l0
							v167 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v79)+4)) = uint8(v167)
							*(*int32)(unsafe.Add(mBase, uint32(v79))) = int32(474)
							v170 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v79)+36)) = v170
							*(*int32)(unsafe.Add(mBase, uint32(v79)+32)) = l1
							*(*int64)(unsafe.Add(mBase, uint32(v79)+20)) = v170
							*(*int32)(unsafe.Add(mBase, uint32(v79)+8)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v79)+12)) = int32(1785164)
							if l0 != 0 {
								v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								*(*int32)(unsafe.Add(mBase, uint32(v79)+28)) = v182
								if v182 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(v182)+24)) = v79
								} else {
								}
								*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v79
								v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
								*(*uint8)(unsafe.Add(mBase, uint32(v79)+5)) = uint8(v186)
							} else {
								v188 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v79)+28)) = v188
								*(*uint8)(unsafe.Add(mBase, uint32(v79)+5)) = uint8(v188)
							}
							v198 = v79
							v199 = v78
							*(*int32)(unsafe.Add(mBase, uint32(v198)+8)) = v199
							m.G0 = v13 + int32(16)
							return v198
						}
					} else {
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v28)+28))
						*(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_consts[1297]))) = v33
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_consts[1298])))
						v36 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_consts[1298]))) = v35 - v36
						*(*int32)(unsafe.Add(mBase, uint32(v28)+96)) = l4
						*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = l0
						*(*uint8)(unsafe.Add(mBase, uint32(v28)+4)) = uint8(v36)
						*(*int32)(unsafe.Add(mBase, uint32(v28))) = int32(474)
						v46 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v28)+36)) = v46
						*(*int32)(unsafe.Add(mBase, uint32(v28)+32)) = l1
						*(*int64)(unsafe.Add(mBase, uint32(v28)+20)) = v46
						*(*int32)(unsafe.Add(mBase, uint32(v28)+8)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v28)+12)) = int32(1785164)
						if l0 != 0 {
							v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v28)+28)) = v58
							if v58 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v58)+24)) = v28
							} else {
							}
							*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v28
							v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
							*(*uint8)(unsafe.Add(mBase, uint32(v28)+5)) = uint8(v62)
						} else {
							v64 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v28)+28)) = v64
							*(*uint8)(unsafe.Add(mBase, uint32(v28)+5)) = uint8(v64)
						}
						v69 = *(*int32)(unsafe.Add(mBase, uint32(v28)+128))
						v198 = v28
						v199 = v69 - v28
						*(*int32)(unsafe.Add(mBase, uint32(v198)+8)) = v199
						m.G0 = v13 + int32(16)
						return v198
					}
				}
			}
		} else {
			v23 = int32(0)
			v25 = v23 << (uint(int32(3)) % 32)
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_consts[1297])))
			if v28 == int32(0) {
				v73 = v23
				if l2 != 0 {
					v75 = l2
				} else {
					v75 = l3
				}
				if base.Ui32(v75) <= base.Ui32(int32(144)) {
					v78 = int32(144)
				} else {
					v78 = v75
				}
				v79 = F_emscripten_builtin_malloc(m, v78)
				mBase = m.M
				if v79 == int32(0) {
					v83 = *(*int32)(unsafe.Add(mBase, _consts[147]))
					if v83 != 0 {
						F_MemoryContextStats(m, v83)
						mBase = m.M
						v87 = m.ExcPending
						if v87 != 0 {
							return int32(0)
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v91 = m.ExcPending
							if v91 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(8389))
								mBase = m.M
								v94 = m.ExcPending
								if v94 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(13961), int32(0))
									mBase = m.M
									v98 = m.ExcPending
									if v98 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v13))) = l1
										F_errdetail(m, int32(678867), v13)
										mBase = m.M
										v102 = m.ExcPending
										if v102 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(504340), int32(453), int32(319332))
											mBase = m.M
											v107 = m.ExcPending
											if v107 != 0 {
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
						v91 = m.ExcPending
						if v91 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(8389))
							mBase = m.M
							v94 = m.ExcPending
							if v94 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(13961), int32(0))
								mBase = m.M
								v98 = m.ExcPending
								if v98 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v13))) = l1
									F_errdetail(m, int32(678867), v13)
									mBase = m.M
									v102 = m.ExcPending
									if v102 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(504340), int32(453), int32(319332))
										mBase = m.M
										v107 = m.ExcPending
										if v107 != 0 {
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
					*(*int32)(unsafe.Add(mBase, uint32(v79)+128)) = v79 + v78
					*(*int32)(unsafe.Add(mBase, uint32(v79)+124)) = v79 + int32(136)
					v114 = v79 + int32(112)
					*(*int32)(unsafe.Add(mBase, uint32(v114))) = v79
					*(*int64)(unsafe.Add(mBase, uint32(v79)+116)) = int64(0)
					*(*int32)(unsafe.Add(mBase, uint32(v79)+44)) = v114
					v120 = v79 + int32(48)
					v122 = v79 + int32(92)
					if base.Ui32(v120) < base.Ui32(v122) {
						v126 = v79 + int32(52)
						if base.Ui32(v126) < base.Ui32(v122) {
							v128 = v122
						} else {
							v128 = v126
						}
						v137 = F__emscripten_memset_bulkmem(m, v120, base.I32_extend8_s(int32(0)), (v128-v79-int32(49))&int32(-4)+int32(4))
						mBase = m.M
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(v79)+108)) = v73
					*(*int32)(unsafe.Add(mBase, uint32(v79)+100)) = l3
					*(*int32)(unsafe.Add(mBase, uint32(v79)+96)) = l4
					*(*int32)(unsafe.Add(mBase, uint32(v122))) = l3
					v151 = int32(8192)
					for {
						if base.Ui32(int32(base.Ui32(l4-int32(24))>>(uint(int32(2))%32))) < base.Ui32(v151+int32(8)) {
							v151 = int32(base.Ui32(v151) >> (uint(int32(1)) % 32))
							continue
						} else {
							break
						}
						break
					}
					*(*int32)(unsafe.Add(mBase, uint32(v79)+104)) = v151
					*(*int32)(unsafe.Add(mBase, uint32(v79)+16)) = l0
					v167 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v79)+4)) = uint8(v167)
					*(*int32)(unsafe.Add(mBase, uint32(v79))) = int32(474)
					v170 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v79)+36)) = v170
					*(*int32)(unsafe.Add(mBase, uint32(v79)+32)) = l1
					*(*int64)(unsafe.Add(mBase, uint32(v79)+20)) = v170
					*(*int32)(unsafe.Add(mBase, uint32(v79)+8)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v79)+12)) = int32(1785164)
					if l0 != 0 {
						v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						*(*int32)(unsafe.Add(mBase, uint32(v79)+28)) = v182
						if v182 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v182)+24)) = v79
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v79
						v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
						*(*uint8)(unsafe.Add(mBase, uint32(v79)+5)) = uint8(v186)
					} else {
						v188 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v79)+28)) = v188
						*(*uint8)(unsafe.Add(mBase, uint32(v79)+5)) = uint8(v188)
					}
					v198 = v79
					v199 = v78
					*(*int32)(unsafe.Add(mBase, uint32(v198)+8)) = v199
					m.G0 = v13 + int32(16)
					return v198
				}
			} else {
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v28)+28))
				*(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_consts[1297]))) = v33
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_consts[1298])))
				v36 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_consts[1298]))) = v35 - v36
				*(*int32)(unsafe.Add(mBase, uint32(v28)+96)) = l4
				*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = l0
				*(*uint8)(unsafe.Add(mBase, uint32(v28)+4)) = uint8(v36)
				*(*int32)(unsafe.Add(mBase, uint32(v28))) = int32(474)
				v46 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v28)+36)) = v46
				*(*int32)(unsafe.Add(mBase, uint32(v28)+32)) = l1
				*(*int64)(unsafe.Add(mBase, uint32(v28)+20)) = v46
				*(*int32)(unsafe.Add(mBase, uint32(v28)+8)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v28)+12)) = int32(1785164)
				if l0 != 0 {
					v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v28)+28)) = v58
					if v58 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v58)+24)) = v28
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v28
					v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
					*(*uint8)(unsafe.Add(mBase, uint32(v28)+5)) = uint8(v62)
				} else {
					v64 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v28)+28)) = v64
					*(*uint8)(unsafe.Add(mBase, uint32(v28)+5)) = uint8(v64)
				}
				v69 = *(*int32)(unsafe.Add(mBase, uint32(v28)+128))
				v198 = v28
				v199 = v69 - v28
				*(*int32)(unsafe.Add(mBase, uint32(v198)+8)) = v199
				m.G0 = v13 + int32(16)
				return v198
			}
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
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = l0 - int32(8)
	v15 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	if v15&int64(16) != int64(0) {
		v21 = l0 - int32(32)
		if v21 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v85 = m.ExcPending
			if v85 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = v14
				F_errmsg_internal(m, int32(244181), v11)
				mBase = m.M
				v89 = m.ExcPending
				if v89 != 0 {
					return
				} else {
					F_errfinish(m, int32(504340), int32(1080), int32(419579))
					mBase = m.M
					v94 = m.ExcPending
					if v94 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
			if v24 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v85 = m.ExcPending
				if v85 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v11))) = v14
					F_errmsg_internal(m, int32(244181), v11)
					mBase = m.M
					v89 = m.ExcPending
					if v89 != 0 {
						return
					} else {
						F_errfinish(m, int32(504340), int32(1080), int32(419579))
						mBase = m.M
						v94 = m.ExcPending
						if v94 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
				if v27 != int32(474) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v85 = m.ExcPending
					if v85 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v11))) = v14
						F_errmsg_internal(m, int32(244181), v11)
						mBase = m.M
						v89 = m.ExcPending
						if v89 != 0 {
							return
						} else {
							F_errfinish(m, int32(504340), int32(1080), int32(419579))
							mBase = m.M
							v94 = m.ExcPending
							if v94 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(l0-int32(20))))
					v34 = l0 - int32(16)
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
					if v32 != v35 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v85 = m.ExcPending
						if v85 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11))) = v14
							F_errmsg_internal(m, int32(244181), v11)
							mBase = m.M
							v89 = m.ExcPending
							if v89 != 0 {
								return
							} else {
								F_errfinish(m, int32(504340), int32(1080), int32(419579))
								mBase = m.M
								v94 = m.ExcPending
								if v94 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v39 = *(*int32)(unsafe.Add(mBase, uint32(l0-int32(24))))
						v41 = l0 - int32(28)
						v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
						if v42 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v42)+8)) = v39
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v24)+44)) = v39
						}
						if v39 != 0 {
							v45 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
							*(*int32)(unsafe.Add(mBase, uint32(v39)+4)) = v45
						} else {
						}
						v47 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
						v48 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
						*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v47 + (v21 - v48)
						F_emscripten_builtin_free(m, v21)
						mBase = m.M
						m.G0 = v11 + int32(16)
						return
					}
				}
			}
		}
	} else {
		v59 = *(*int32)(unsafe.Add(mBase, uint32(v14-base.I32_wrap_i64(int64(base.Ui64(v15)>>(uint(int64(34))%64)))&int32(1073741822))))
		v67 = v59 + base.I32_wrap_i64(int64(base.Ui64(v15)>>(uint(int64(5))%64)))<<(uint(int32(2))%32) + int32(48)
		v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v68
		*(*int32)(unsafe.Add(mBase, uint32(v67))) = v14
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
		v18 = *(*int32)(unsafe.Add(mBase, _consts[237]))
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
									F_errmsg(m, int32(715961), v8)
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return
									} else {
										F_errhint(m, int32(623275), int32(0))
										mBase = m.M
										v72 = m.ExcPending
										if v72 != 0 {
											return
										} else {
											F_errfinish(m, int32(507053), int32(558), int32(317634))
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
									v43 = *(*int32)(unsafe.Add(mBase, _consts[231]))
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
								F_errmsg(m, int32(715961), v8)
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return
								} else {
									F_errhint(m, int32(623275), int32(0))
									mBase = m.M
									v72 = m.ExcPending
									if v72 != 0 {
										return
									} else {
										F_errfinish(m, int32(507053), int32(558), int32(317634))
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
								v43 = *(*int32)(unsafe.Add(mBase, _consts[231]))
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
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
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v371 int32
	_ = v371
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v402 int32
	_ = v402
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
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
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
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v539 int32
	_ = v539
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v566 int32
	_ = v566
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
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
	var v702 int32
	_ = v702
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v714 int64
	_ = v714
	var v716 int32
	_ = v716
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v8 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v381 != int32(1024) {
		goto L67
	} else {
		goto L68
	}
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v381 = v11
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v371
	v381 = v371
	goto L1
L6:
	;
	v24 = v14
	v25 = v12
	v27 = l0 + int32(28)
	goto L9
L7:
	;
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14+v13))) = v8
	v365 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v371 = v365 + int32(4)
	goto L5
L9:
	;
	if base.Ui32(int32(1024)) <= base.Ui32(v24) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v371 = v361
	goto L5
L11:
	;
	v30 = int32(1024)
	v37 = int32(-1636607408)
	goto L16
L12:
	;
	v352 = v24
	goto L13
L13:
	;
	v355 = int32(1024) - v352
	if base.Ui32(v25) < base.Ui32(v355) {
		goto L59
	} else {
		goto L60
	}
L14:
	;
	v347 = F_Int64GetDatum(m, base.I64_extend_i32_u(v337)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v329^v337-base.I32_rotl(v337, int32(24))))
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
	v357 = v25
	goto L61
L60:
	;
	v357 = v355
	goto L61
L61:
	;
	if v357 != 0 {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v361 = v357 + v352
	v362 = v25 - v357
	if v362 != 0 {
		v24 = v361
		v25 = v362
		v27 = v357 + v27
		goto L9
	} else {
		goto L66
	}
L63:
	;
	v358 = F__emscripten_memcpy_bulkmem(m, v352+v13, v27, v357)
	mBase = m.M
	goto L65
L64:
	;
	goto L65
L65:
	;
	goto L62
L66:
	;
	goto L10
L67:
	;
	v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	*(*uint8)(unsafe.Add(mBase, uint32(v385+v381))) = uint8(v389)
	v391 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v391 + int32(1)
	return
L68:
	;
	goto L69
L69:
	;
	v395 = int32(1024)
	v402 = int32(-1636607408)
	goto L72
L70:
	;
	v712 = F_Int64GetDatum(m, base.I64_extend_i32_u(v702)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v694^v702-base.I32_rotl(v702, int32(24))))
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L57
	} else {
		goto L113
	}
L71:
	;
	if v385&int32(3) != 0 {
		goto L87
	} else {
		goto L88
	}
L72:
	;
	goto L71
L75:
	;
	v680 = int32(14)
	v682 = v676 ^ v677 - base.I32_rotl(v676, v680)
	v686 = v682 ^ v675 - base.I32_rotl(v682, int32(11))
	v690 = v686 ^ v676 - base.I32_rotl(v686, int32(25))
	v694 = v690 ^ v682 - base.I32_rotl(v690, int32(16))
	v698 = v694 ^ v686 - base.I32_rotl(v694, int32(4))
	v702 = v698 ^ v690 - base.I32_rotl(v698, v680)
	goto L70
L76:
	;
	v670 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v492))))
	v675 = v667 + v670
	v676 = v668
	v677 = v669
	goto L75
L77:
	;
	v663 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v492)+1)))
	v667 = v663<<(uint(int32(8))%32) + v660
	v668 = v661
	v669 = v662
	goto L76
L78:
	;
	v656 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v492)+2)))
	v660 = v656<<(uint(int32(16))%32) + v653
	v661 = v654
	v662 = v655
	goto L77
L79:
	;
	v649 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v492)+3)))
	v653 = v649<<(uint(int32(24))%32) + v485
	v654 = v647
	v655 = v648
	goto L78
L80:
	;
	v645 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v492)+4)))
	v647 = v643 + v645
	v648 = v644
	goto L79
L81:
	;
	v639 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v492)+5)))
	v643 = v639<<(uint(int32(8))%32) + v637
	v644 = v638
	goto L80
L82:
	;
	v633 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v492)+6)))
	v637 = v633<<(uint(int32(16))%32) + v631
	v638 = v632
	goto L81
L83:
	;
	v627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v492)+7)))
	v631 = v627<<(uint(int32(24))%32) + v486
	v632 = v626
	goto L82
L84:
	;
	v622 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v492)+8)))
	v626 = v622<<(uint(int32(8))%32) + v621
	goto L83
L85:
	;
	v617 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v492)+9)))
	v621 = v617<<(uint(int32(16))%32) + v616
	goto L84
L86:
	;
	v612 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v492)+10)))
	v616 = v612<<(uint(int32(24))%32) + v490
	goto L85
L87:
	;
	goto L90
L88:
	;
	goto L89
L89:
	;
	goto L96
L90:
	;
	v448 = v385
	v449 = v395
	v451 = v402
	v452 = v402
	v453 = v402
	goto L93
L92:
	;
	switch v494 - int32(1) {
	case 0:
		v667 = v485
		v668 = v486
		v669 = v490
		goto L76
	case 1:
		v660 = v485
		v661 = v486
		v662 = v490
		goto L77
	case 2:
		v653 = v485
		v654 = v486
		v655 = v490
		goto L78
	case 3:
		v647 = v486
		v648 = v490
		goto L79
	case 4:
		v643 = v486
		v644 = v490
		goto L80
	case 5:
		v637 = v486
		v638 = v490
		goto L81
	case 6:
		v631 = v486
		v632 = v490
		goto L82
	case 7:
		v626 = v490
		goto L83
	case 8:
		v621 = v490
		goto L84
	case 9:
		v616 = v490
		goto L85
	case 10:
		goto L86
	default:
		v675 = v485
		v676 = v486
		v677 = v490
		goto L75
	}
L93:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v448)+4))
	v456 = v455 + v452
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v448)))
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v448)+8))
	v460 = v459 + v453
	v462 = int32(4)
	v464 = v457 + v451 - v460 ^ base.I32_rotl(v460, v462)
	v468 = v456 - v464 ^ base.I32_rotl(v464, int32(6))
	v469 = v460 + v456
	v470 = v464 + v469
	v471 = v468 + v470
	v475 = v469 - v468 ^ base.I32_rotl(v468, int32(8))
	v479 = v470 - v475 ^ base.I32_rotl(v475, int32(16))
	v483 = v471 - v479 ^ base.I32_rotl(v479, int32(19))
	v484 = v475 + v471
	v485 = v479 + v484
	v486 = v483 + v485
	v490 = v484 - v483 ^ base.I32_rotl(v483, v462)
	v491 = int32(12)
	v492 = v448 + v491
	v494 = v449 - v491
	if base.Ui32(int32(11)) < base.Ui32(v494) {
		v448 = v492
		v449 = v494
		v451 = v485
		v452 = v486
		v453 = v490
		goto L93
	} else {
		goto L95
	}
L94:
	;
	goto L92
L95:
	;
	goto L94
L96:
	;
	v508 = v385
	v509 = v395
	v511 = v402
	v512 = v402
	v513 = v402
	goto L99
L98:
	;
	switch v554 - int32(1) {
	case 0:
		v609 = v545
		goto L102
	case 1:
		v604 = v545
		goto L103
	case 2:
		goto L104
	case 3:
		v597 = v546
		goto L105
	case 4:
		v594 = v546
		goto L106
	case 5:
		v589 = v546
		goto L107
	case 6:
		goto L108
	case 7:
		v580 = v550
		goto L109
	case 8:
		v575 = v550
		goto L110
	case 9:
		v570 = v550
		goto L111
	case 10:
		goto L112
	default:
		v675 = v545
		v676 = v546
		v677 = v550
		goto L75
	}
L99:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v508)+4))
	v516 = v515 + v512
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v508)))
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v508)+8))
	v520 = v519 + v513
	v522 = int32(4)
	v524 = v517 + v511 - v520 ^ base.I32_rotl(v520, v522)
	v528 = v516 - v524 ^ base.I32_rotl(v524, int32(6))
	v529 = v520 + v516
	v530 = v524 + v529
	v531 = v528 + v530
	v535 = v529 - v528 ^ base.I32_rotl(v528, int32(8))
	v539 = v530 - v535 ^ base.I32_rotl(v535, int32(16))
	v543 = v531 - v539 ^ base.I32_rotl(v539, int32(19))
	v544 = v535 + v531
	v545 = v539 + v544
	v546 = v543 + v545
	v550 = v544 - v543 ^ base.I32_rotl(v543, v522)
	v551 = int32(12)
	v552 = v508 + v551
	v554 = v509 - v551
	if base.Ui32(int32(11)) < base.Ui32(v554) {
		v508 = v552
		v509 = v554
		v511 = v545
		v512 = v546
		v513 = v550
		goto L99
	} else {
		goto L101
	}
L100:
	;
	goto L98
L101:
	;
	goto L100
L102:
	;
	v610 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v552))))
	v675 = v609 + v610
	v676 = v546
	v677 = v550
	goto L75
L103:
	;
	v605 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v552)+1)))
	v609 = v605<<(uint(int32(8))%32) + v604
	goto L102
L104:
	;
	v600 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v552)+2)))
	v604 = v600<<(uint(int32(16))%32) + v545
	goto L103
L105:
	;
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v552)))
	v675 = v598 + v545
	v676 = v597
	v677 = v550
	goto L75
L106:
	;
	v595 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v552)+4)))
	v597 = v594 + v595
	goto L105
L107:
	;
	v590 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v552)+5)))
	v594 = v590<<(uint(int32(8))%32) + v589
	goto L106
L108:
	;
	v585 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v552)+6)))
	v589 = v585<<(uint(int32(16))%32) + v546
	goto L107
L109:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v552)))
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v552)+4))
	v675 = v581 + v545
	v676 = v583 + v546
	v677 = v580
	goto L75
L110:
	;
	v576 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v552)+8)))
	v580 = v576<<(uint(int32(8))%32) + v575
	goto L109
L111:
	;
	v571 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v552)+9)))
	v575 = v571<<(uint(int32(16))%32) + v570
	goto L110
L112:
	;
	v566 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v552)+10)))
	v570 = v566<<(uint(int32(24))%32) + v550
	goto L111
L113:
	;
	v714 = *(*int64)(unsafe.Add(mBase, uint32(v712)))
	*(*int64)(unsafe.Add(mBase, uint32(v385))) = v714
	v716 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	*(*uint8)(unsafe.Add(mBase, uint32(v385)+8)) = uint8(v716)
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
	v2 = int32(*(*uint8)(unsafe.Add(mBase, _consts[301])))
	if v2 == int32(0) {
		v6 = int32(1)
		*(*uint8)(unsafe.Add(mBase, _consts[301])) = uint8(v6)
	} else {
	}
	return
}
func F_ApplyWorkerMain(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int64
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
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
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
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
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
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
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	v7 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[512])) = uint8(v7)
	F_SetupApplyOrSyncWorker(m, l0)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		v12 = int32(0)
		*(*uint8)(unsafe.Add(mBase, _consts[512])) = uint8(v12)
		v14 = m.G0
		v16 = v14 - int32(160)
		m.G0 = v16
		v19 = *(*int32)(unsafe.Add(mBase, _consts[527]))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+40))
		if v20 != 0 {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
			*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v21
			v29 = F_pg_snprintf(m, v16+int32(96), int32(64), int32(39486), v16+int32(32))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return
			} else {
				F_StartTransactionCommand(m)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					v36 = F_replorigin_by_name(m, v16+int32(96), int32(1))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						if v36 == int32(0) {
							v42 = F_replorigin_create(m, v16+int32(96))
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return
							} else {
								v44 = v42
								F_replorigin_session_setup(m, v44, int32(0))
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return
								} else {
									*(*uint16)(unsafe.Add(mBase, _consts[110])) = uint16(v44)
									v50 = F_replorigin_session_get_progress(m)
									mBase = m.M
									v51 = m.ExcPending
									if v51 != 0 {
										return
									} else {
										F_CommitTransactionCommand(m)
										mBase = m.M
										v53 = m.ExcPending
										if v53 != 0 {
											return
										} else {
											v55 = *(*int32)(unsafe.Add(mBase, _consts[527]))
											v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+30)))
											if v56 == int32(1) {
												v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+24)))
												v62 = v59 ^ int32(1)
											} else {
												v62 = int32(0)
											}
											v64 = *(*int32)(unsafe.Add(mBase, uint32(v55)+36))
											v65 = int32(1)
											v69 = *(*int32)(unsafe.Add(mBase, uint32(v55)+16))
											v73 = *(*int32)(unsafe.Add(mBase, _consts[300]))
											v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
											v75 = m.T0[v74].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v64, v65, v65, v62&v65, v69, v16+int32(48))
											mBase = m.M
											v76 = m.ExcPending
											if v76 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, _consts[508])) = v75
												if v75 == int32(0) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v270 = m.ExcPending
													if v270 != 0 {
														return
													} else {
														F_errcode(m, int32(100663808))
														mBase = m.M
														v273 = m.ExcPending
														if v273 != 0 {
															return
														} else {
															v275 = *(*int32)(unsafe.Add(mBase, _consts[527]))
															v276 = *(*int32)(unsafe.Add(mBase, uint32(v275)+16))
															*(*int32)(unsafe.Add(mBase, uint32(v16))) = v276
															v278 = *(*int32)(unsafe.Add(mBase, uint32(v16)+48))
															*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v278
															F_errmsg(m, int32(205616), v16)
															mBase = m.M
															v282 = m.ExcPending
															if v282 != 0 {
																return
															} else {
																F_errfinish(m, int32(506994), int32(4593), int32(224471))
																mBase = m.M
																v287 = m.ExcPending
																if v287 != 0 {
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
													v84 = *(*int32)(unsafe.Add(mBase, _consts[300]))
													v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+16))
													v86 = m.T0[v85].(func(*base.Module, int32, int32) int32)(m, v75, v16+int32(52))
													mBase = m.M
													v87 = m.ExcPending
													if v87 != 0 {
														return
													} else {
														v90 = *(*int32)(unsafe.Add(mBase, _consts[531]))
														v93 = F_MemoryContextStrdup(m, v90, v16+int32(96))
														mBase = m.M
														v94 = m.ExcPending
														if v94 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, _consts[555])) = v93
															*(*int64)(unsafe.Add(mBase, uint32(v16)+64)) = v50
															v97 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v16)+56)) = uint8(v97)
															*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v20
															v101 = *(*int32)(unsafe.Add(mBase, _consts[508]))
															v103 = *(*int32)(unsafe.Add(mBase, _consts[300]))
															v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+24))
															v105 = m.T0[v104].(func(*base.Module, int32) int32)(m, v101)
															mBase = m.M
															v106 = m.ExcPending
															if v106 != 0 {
																return
															} else {
																if v105 <= int32(159999) {
																	if int32(139999) < v105 {
																		v114 = int32(2)
																	} else {
																		v114 = int32(1)
																	}
																	if int32(149999) < v105 {
																		v117 = int32(3)
																	} else {
																		v117 = v114
																	}
																	*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = v117
																	v120 = *(*int32)(unsafe.Add(mBase, _consts[527]))
																	v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)+48))
																	*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = v121
																	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+26)))
																	*(*uint8)(unsafe.Add(mBase, uint32(v16)+80)) = uint8(v123)
																	if int32(140000) <= v105 {
																		v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+27)))
																		v142 = v120
																		v143 = v141
																		if v143&int32(255) != int32(102) {
																			v150 = int32(278626)
																		} else {
																			v150 = int32(0)
																		}
																		v152 = v142
																		v153 = v150
																		v154 = int32(0)
																	} else {
																		v152 = v120
																		v153 = int32(0)
																		v154 = int32(0)
																	}
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = int32(4)
																	v131 = *(*int32)(unsafe.Add(mBase, _consts[527]))
																	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+48))
																	*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = v132
																	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131)+26)))
																	*(*uint8)(unsafe.Add(mBase, uint32(v16)+80)) = uint8(v134)
																	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131)+27)))
																	if v136 != int32(112) {
																		v142 = v131
																		v143 = v136
																		if v143&int32(255) != int32(102) {
																			v150 = int32(278626)
																		} else {
																			v150 = int32(0)
																		}
																		v152 = v142
																		v153 = v150
																		v154 = int32(0)
																	} else {
																		v152 = v131
																		v153 = int32(314228)
																		v154 = int32(1)
																	}
																}
																*(*int32)(unsafe.Add(mBase, uint32(v16)+84)) = v153
																v157 = *(*int32)(unsafe.Add(mBase, _consts[507]))
																*(*uint8)(unsafe.Add(mBase, uint32(v157)+68)) = uint8(v154)
																v159 = int32(0)
																*(*uint8)(unsafe.Add(mBase, uint32(v16)+88)) = uint8(v159)
																v161 = *(*int32)(unsafe.Add(mBase, uint32(v152)+52))
																v162 = F_pstrdup(m, v161)
																mBase = m.M
																v163 = m.ExcPending
																if v163 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v16)+92)) = v162
																	v166 = *(*int32)(unsafe.Add(mBase, _consts[527]))
																	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+28)))
																	if v167 != int32(112) {
																		v205 = *(*int32)(unsafe.Add(mBase, _consts[508]))
																		v209 = *(*int32)(unsafe.Add(mBase, _consts[300]))
																		v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)+32))
																		v211 = m.T0[v210].(func(*base.Module, int32, int32) int32)(m, v205, v16+int32(56))
																		mBase = m.M
																		v212 = m.ExcPending
																		if v212 != 0 {
																			return
																		} else {
																			v215 = F_errstart(m, int32(14), int32(0))
																			mBase = m.M
																			v216 = m.ExcPending
																			if v216 != 0 {
																				return
																			} else {
																				if v215 != 0 {
																					v218 = *(*int32)(unsafe.Add(mBase, _consts[527]))
																					v219 = *(*int32)(unsafe.Add(mBase, uint32(v218)+16))
																					v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+28)))
																					switch v221 - int32(100) {
																					case 0:
																						v230 = int32(556578)
																					default:
																						if v221 == int32(101) {
																							v228 = int32(556587)
																						} else {
																							v228 = int32(558895)
																						}
																						v230 = v228
																					case 12:
																						v230 = int32(549382)
																					}
																					*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v230
																					*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v219
																					F_errmsg_internal(m, int32(184241), v16+int32(16))
																					mBase = m.M
																					v237 = m.ExcPending
																					if v237 != 0 {
																						return
																					} else {
																						F_errfinish(m, int32(506994), int32(4645), int32(224471))
																						mBase = m.M
																						v242 = m.ExcPending
																						if v242 != 0 {
																							return
																						} else {
																							F_start_apply(m, v50)
																							mBase = m.M
																							v247 = m.ExcPending
																							if v247 != 0 {
																								return
																							} else {
																								m.G0 = v16 + int32(160)
																								F_proc_exit(m, int32(0))
																								mBase = m.M
																								v290 = m.ExcPending
																								if v290 != 0 {
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
																					F_start_apply(m, v50)
																					mBase = m.M
																					v247 = m.ExcPending
																					if v247 != 0 {
																						return
																					} else {
																						m.G0 = v16 + int32(160)
																						F_proc_exit(m, int32(0))
																						mBase = m.M
																						v290 = m.ExcPending
																						if v290 != 0 {
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
																		v170 = F_AllTablesyncsReady(m)
																		mBase = m.M
																		v171 = m.ExcPending
																		if v171 != 0 {
																			return
																		} else {
																			if v170 == int32(0) {
																				v205 = *(*int32)(unsafe.Add(mBase, _consts[508]))
																				v209 = *(*int32)(unsafe.Add(mBase, _consts[300]))
																				v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)+32))
																				v211 = m.T0[v210].(func(*base.Module, int32, int32) int32)(m, v205, v16+int32(56))
																				mBase = m.M
																				v212 = m.ExcPending
																				if v212 != 0 {
																					return
																				} else {
																					v215 = F_errstart(m, int32(14), int32(0))
																					mBase = m.M
																					v216 = m.ExcPending
																					if v216 != 0 {
																						return
																					} else {
																						if v215 != 0 {
																							v218 = *(*int32)(unsafe.Add(mBase, _consts[527]))
																							v219 = *(*int32)(unsafe.Add(mBase, uint32(v218)+16))
																							v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+28)))
																							switch v221 - int32(100) {
																							case 0:
																								v230 = int32(556578)
																							default:
																								if v221 == int32(101) {
																									v228 = int32(556587)
																								} else {
																									v228 = int32(558895)
																								}
																								v230 = v228
																							case 12:
																								v230 = int32(549382)
																							}
																							*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v230
																							*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v219
																							F_errmsg_internal(m, int32(184241), v16+int32(16))
																							mBase = m.M
																							v237 = m.ExcPending
																							if v237 != 0 {
																								return
																							} else {
																								F_errfinish(m, int32(506994), int32(4645), int32(224471))
																								mBase = m.M
																								v242 = m.ExcPending
																								if v242 != 0 {
																									return
																								} else {
																									F_start_apply(m, v50)
																									mBase = m.M
																									v247 = m.ExcPending
																									if v247 != 0 {
																										return
																									} else {
																										m.G0 = v16 + int32(160)
																										F_proc_exit(m, int32(0))
																										mBase = m.M
																										v290 = m.ExcPending
																										if v290 != 0 {
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
																							F_start_apply(m, v50)
																							mBase = m.M
																							v247 = m.ExcPending
																							if v247 != 0 {
																								return
																							} else {
																								m.G0 = v16 + int32(160)
																								F_proc_exit(m, int32(0))
																								mBase = m.M
																								v290 = m.ExcPending
																								if v290 != 0 {
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
																				v174 = int32(1)
																				*(*uint8)(unsafe.Add(mBase, uint32(v16)+88)) = uint8(v174)
																				v177 = *(*int32)(unsafe.Add(mBase, _consts[508]))
																				v181 = *(*int32)(unsafe.Add(mBase, _consts[300]))
																				v182 = *(*int32)(unsafe.Add(mBase, uint32(v181)+32))
																				v183 = m.T0[v182].(func(*base.Module, int32, int32) int32)(m, v177, v16+int32(56))
																				mBase = m.M
																				v184 = m.ExcPending
																				if v184 != 0 {
																					return
																				} else {
																					F_StartTransactionCommand(m)
																					mBase = m.M
																					v186 = m.ExcPending
																					if v186 != 0 {
																						return
																					} else {
																						v187 = F_GetTransactionSnapshot(m)
																						mBase = m.M
																						v188 = m.ExcPending
																						if v188 != 0 {
																							return
																						} else {
																							F_PushActiveSnapshot(m, v187)
																							mBase = m.M
																							v190 = m.ExcPending
																							if v190 != 0 {
																								return
																							} else {
																								v192 = *(*int32)(unsafe.Add(mBase, _consts[527]))
																								v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)))
																								F_UpdateTwoPhaseState(m, v193)
																								mBase = m.M
																								v195 = m.ExcPending
																								if v195 != 0 {
																									return
																								} else {
																									v197 = *(*int32)(unsafe.Add(mBase, _consts[527]))
																									v198 = int32(101)
																									*(*uint8)(unsafe.Add(mBase, uint32(v197)+28)) = uint8(v198)
																									F_PopActiveSnapshot(m)
																									mBase = m.M
																									v201 = m.ExcPending
																									if v201 != 0 {
																										return
																									} else {
																										F_CommitTransactionCommand(m)
																										mBase = m.M
																										v203 = m.ExcPending
																										if v203 != 0 {
																											return
																										} else {
																											v215 = F_errstart(m, int32(14), int32(0))
																											mBase = m.M
																											v216 = m.ExcPending
																											if v216 != 0 {
																												return
																											} else {
																												if v215 != 0 {
																													v218 = *(*int32)(unsafe.Add(mBase, _consts[527]))
																													v219 = *(*int32)(unsafe.Add(mBase, uint32(v218)+16))
																													v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+28)))
																													switch v221 - int32(100) {
																													case 0:
																														v230 = int32(556578)
																													default:
																														if v221 == int32(101) {
																															v228 = int32(556587)
																														} else {
																															v228 = int32(558895)
																														}
																														v230 = v228
																													case 12:
																														v230 = int32(549382)
																													}
																													*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v230
																													*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v219
																													F_errmsg_internal(m, int32(184241), v16+int32(16))
																													mBase = m.M
																													v237 = m.ExcPending
																													if v237 != 0 {
																														return
																													} else {
																														F_errfinish(m, int32(506994), int32(4645), int32(224471))
																														mBase = m.M
																														v242 = m.ExcPending
																														if v242 != 0 {
																															return
																														} else {
																															F_start_apply(m, v50)
																															mBase = m.M
																															v247 = m.ExcPending
																															if v247 != 0 {
																																return
																															} else {
																																m.G0 = v16 + int32(160)
																																F_proc_exit(m, int32(0))
																																mBase = m.M
																																v290 = m.ExcPending
																																if v290 != 0 {
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
																													F_start_apply(m, v50)
																													mBase = m.M
																													v247 = m.ExcPending
																													if v247 != 0 {
																														return
																													} else {
																														m.G0 = v16 + int32(160)
																														F_proc_exit(m, int32(0))
																														mBase = m.M
																														v290 = m.ExcPending
																														if v290 != 0 {
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
							v44 = v36
							F_replorigin_session_setup(m, v44, int32(0))
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return
							} else {
								*(*uint16)(unsafe.Add(mBase, _consts[110])) = uint16(v44)
								v50 = F_replorigin_session_get_progress(m)
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return
								} else {
									F_CommitTransactionCommand(m)
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
										return
									} else {
										v55 = *(*int32)(unsafe.Add(mBase, _consts[527]))
										v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+30)))
										if v56 == int32(1) {
											v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+24)))
											v62 = v59 ^ int32(1)
										} else {
											v62 = int32(0)
										}
										v64 = *(*int32)(unsafe.Add(mBase, uint32(v55)+36))
										v65 = int32(1)
										v69 = *(*int32)(unsafe.Add(mBase, uint32(v55)+16))
										v73 = *(*int32)(unsafe.Add(mBase, _consts[300]))
										v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
										v75 = m.T0[v74].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v64, v65, v65, v62&v65, v69, v16+int32(48))
										mBase = m.M
										v76 = m.ExcPending
										if v76 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, _consts[508])) = v75
											if v75 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v270 = m.ExcPending
												if v270 != 0 {
													return
												} else {
													F_errcode(m, int32(100663808))
													mBase = m.M
													v273 = m.ExcPending
													if v273 != 0 {
														return
													} else {
														v275 = *(*int32)(unsafe.Add(mBase, _consts[527]))
														v276 = *(*int32)(unsafe.Add(mBase, uint32(v275)+16))
														*(*int32)(unsafe.Add(mBase, uint32(v16))) = v276
														v278 = *(*int32)(unsafe.Add(mBase, uint32(v16)+48))
														*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v278
														F_errmsg(m, int32(205616), v16)
														mBase = m.M
														v282 = m.ExcPending
														if v282 != 0 {
															return
														} else {
															F_errfinish(m, int32(506994), int32(4593), int32(224471))
															mBase = m.M
															v287 = m.ExcPending
															if v287 != 0 {
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
												v84 = *(*int32)(unsafe.Add(mBase, _consts[300]))
												v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+16))
												v86 = m.T0[v85].(func(*base.Module, int32, int32) int32)(m, v75, v16+int32(52))
												mBase = m.M
												v87 = m.ExcPending
												if v87 != 0 {
													return
												} else {
													v90 = *(*int32)(unsafe.Add(mBase, _consts[531]))
													v93 = F_MemoryContextStrdup(m, v90, v16+int32(96))
													mBase = m.M
													v94 = m.ExcPending
													if v94 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, _consts[555])) = v93
														*(*int64)(unsafe.Add(mBase, uint32(v16)+64)) = v50
														v97 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v16)+56)) = uint8(v97)
														*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v20
														v101 = *(*int32)(unsafe.Add(mBase, _consts[508]))
														v103 = *(*int32)(unsafe.Add(mBase, _consts[300]))
														v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+24))
														v105 = m.T0[v104].(func(*base.Module, int32) int32)(m, v101)
														mBase = m.M
														v106 = m.ExcPending
														if v106 != 0 {
															return
														} else {
															if v105 <= int32(159999) {
																if int32(139999) < v105 {
																	v114 = int32(2)
																} else {
																	v114 = int32(1)
																}
																if int32(149999) < v105 {
																	v117 = int32(3)
																} else {
																	v117 = v114
																}
																*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = v117
																v120 = *(*int32)(unsafe.Add(mBase, _consts[527]))
																v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)+48))
																*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = v121
																v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+26)))
																*(*uint8)(unsafe.Add(mBase, uint32(v16)+80)) = uint8(v123)
																if int32(140000) <= v105 {
																	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+27)))
																	v142 = v120
																	v143 = v141
																	if v143&int32(255) != int32(102) {
																		v150 = int32(278626)
																	} else {
																		v150 = int32(0)
																	}
																	v152 = v142
																	v153 = v150
																	v154 = int32(0)
																} else {
																	v152 = v120
																	v153 = int32(0)
																	v154 = int32(0)
																}
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = int32(4)
																v131 = *(*int32)(unsafe.Add(mBase, _consts[527]))
																v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+48))
																*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = v132
																v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131)+26)))
																*(*uint8)(unsafe.Add(mBase, uint32(v16)+80)) = uint8(v134)
																v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131)+27)))
																if v136 != int32(112) {
																	v142 = v131
																	v143 = v136
																	if v143&int32(255) != int32(102) {
																		v150 = int32(278626)
																	} else {
																		v150 = int32(0)
																	}
																	v152 = v142
																	v153 = v150
																	v154 = int32(0)
																} else {
																	v152 = v131
																	v153 = int32(314228)
																	v154 = int32(1)
																}
															}
															*(*int32)(unsafe.Add(mBase, uint32(v16)+84)) = v153
															v157 = *(*int32)(unsafe.Add(mBase, _consts[507]))
															*(*uint8)(unsafe.Add(mBase, uint32(v157)+68)) = uint8(v154)
															v159 = int32(0)
															*(*uint8)(unsafe.Add(mBase, uint32(v16)+88)) = uint8(v159)
															v161 = *(*int32)(unsafe.Add(mBase, uint32(v152)+52))
															v162 = F_pstrdup(m, v161)
															mBase = m.M
															v163 = m.ExcPending
															if v163 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v16)+92)) = v162
																v166 = *(*int32)(unsafe.Add(mBase, _consts[527]))
																v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+28)))
																if v167 != int32(112) {
																	v205 = *(*int32)(unsafe.Add(mBase, _consts[508]))
																	v209 = *(*int32)(unsafe.Add(mBase, _consts[300]))
																	v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)+32))
																	v211 = m.T0[v210].(func(*base.Module, int32, int32) int32)(m, v205, v16+int32(56))
																	mBase = m.M
																	v212 = m.ExcPending
																	if v212 != 0 {
																		return
																	} else {
																		v215 = F_errstart(m, int32(14), int32(0))
																		mBase = m.M
																		v216 = m.ExcPending
																		if v216 != 0 {
																			return
																		} else {
																			if v215 != 0 {
																				v218 = *(*int32)(unsafe.Add(mBase, _consts[527]))
																				v219 = *(*int32)(unsafe.Add(mBase, uint32(v218)+16))
																				v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+28)))
																				switch v221 - int32(100) {
																				case 0:
																					v230 = int32(556578)
																				default:
																					if v221 == int32(101) {
																						v228 = int32(556587)
																					} else {
																						v228 = int32(558895)
																					}
																					v230 = v228
																				case 12:
																					v230 = int32(549382)
																				}
																				*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v230
																				*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v219
																				F_errmsg_internal(m, int32(184241), v16+int32(16))
																				mBase = m.M
																				v237 = m.ExcPending
																				if v237 != 0 {
																					return
																				} else {
																					F_errfinish(m, int32(506994), int32(4645), int32(224471))
																					mBase = m.M
																					v242 = m.ExcPending
																					if v242 != 0 {
																						return
																					} else {
																						F_start_apply(m, v50)
																						mBase = m.M
																						v247 = m.ExcPending
																						if v247 != 0 {
																							return
																						} else {
																							m.G0 = v16 + int32(160)
																							F_proc_exit(m, int32(0))
																							mBase = m.M
																							v290 = m.ExcPending
																							if v290 != 0 {
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
																				F_start_apply(m, v50)
																				mBase = m.M
																				v247 = m.ExcPending
																				if v247 != 0 {
																					return
																				} else {
																					m.G0 = v16 + int32(160)
																					F_proc_exit(m, int32(0))
																					mBase = m.M
																					v290 = m.ExcPending
																					if v290 != 0 {
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
																	v170 = F_AllTablesyncsReady(m)
																	mBase = m.M
																	v171 = m.ExcPending
																	if v171 != 0 {
																		return
																	} else {
																		if v170 == int32(0) {
																			v205 = *(*int32)(unsafe.Add(mBase, _consts[508]))
																			v209 = *(*int32)(unsafe.Add(mBase, _consts[300]))
																			v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)+32))
																			v211 = m.T0[v210].(func(*base.Module, int32, int32) int32)(m, v205, v16+int32(56))
																			mBase = m.M
																			v212 = m.ExcPending
																			if v212 != 0 {
																				return
																			} else {
																				v215 = F_errstart(m, int32(14), int32(0))
																				mBase = m.M
																				v216 = m.ExcPending
																				if v216 != 0 {
																					return
																				} else {
																					if v215 != 0 {
																						v218 = *(*int32)(unsafe.Add(mBase, _consts[527]))
																						v219 = *(*int32)(unsafe.Add(mBase, uint32(v218)+16))
																						v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+28)))
																						switch v221 - int32(100) {
																						case 0:
																							v230 = int32(556578)
																						default:
																							if v221 == int32(101) {
																								v228 = int32(556587)
																							} else {
																								v228 = int32(558895)
																							}
																							v230 = v228
																						case 12:
																							v230 = int32(549382)
																						}
																						*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v230
																						*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v219
																						F_errmsg_internal(m, int32(184241), v16+int32(16))
																						mBase = m.M
																						v237 = m.ExcPending
																						if v237 != 0 {
																							return
																						} else {
																							F_errfinish(m, int32(506994), int32(4645), int32(224471))
																							mBase = m.M
																							v242 = m.ExcPending
																							if v242 != 0 {
																								return
																							} else {
																								F_start_apply(m, v50)
																								mBase = m.M
																								v247 = m.ExcPending
																								if v247 != 0 {
																									return
																								} else {
																									m.G0 = v16 + int32(160)
																									F_proc_exit(m, int32(0))
																									mBase = m.M
																									v290 = m.ExcPending
																									if v290 != 0 {
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
																						F_start_apply(m, v50)
																						mBase = m.M
																						v247 = m.ExcPending
																						if v247 != 0 {
																							return
																						} else {
																							m.G0 = v16 + int32(160)
																							F_proc_exit(m, int32(0))
																							mBase = m.M
																							v290 = m.ExcPending
																							if v290 != 0 {
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
																			v174 = int32(1)
																			*(*uint8)(unsafe.Add(mBase, uint32(v16)+88)) = uint8(v174)
																			v177 = *(*int32)(unsafe.Add(mBase, _consts[508]))
																			v181 = *(*int32)(unsafe.Add(mBase, _consts[300]))
																			v182 = *(*int32)(unsafe.Add(mBase, uint32(v181)+32))
																			v183 = m.T0[v182].(func(*base.Module, int32, int32) int32)(m, v177, v16+int32(56))
																			mBase = m.M
																			v184 = m.ExcPending
																			if v184 != 0 {
																				return
																			} else {
																				F_StartTransactionCommand(m)
																				mBase = m.M
																				v186 = m.ExcPending
																				if v186 != 0 {
																					return
																				} else {
																					v187 = F_GetTransactionSnapshot(m)
																					mBase = m.M
																					v188 = m.ExcPending
																					if v188 != 0 {
																						return
																					} else {
																						F_PushActiveSnapshot(m, v187)
																						mBase = m.M
																						v190 = m.ExcPending
																						if v190 != 0 {
																							return
																						} else {
																							v192 = *(*int32)(unsafe.Add(mBase, _consts[527]))
																							v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)))
																							F_UpdateTwoPhaseState(m, v193)
																							mBase = m.M
																							v195 = m.ExcPending
																							if v195 != 0 {
																								return
																							} else {
																								v197 = *(*int32)(unsafe.Add(mBase, _consts[527]))
																								v198 = int32(101)
																								*(*uint8)(unsafe.Add(mBase, uint32(v197)+28)) = uint8(v198)
																								F_PopActiveSnapshot(m)
																								mBase = m.M
																								v201 = m.ExcPending
																								if v201 != 0 {
																									return
																								} else {
																									F_CommitTransactionCommand(m)
																									mBase = m.M
																									v203 = m.ExcPending
																									if v203 != 0 {
																										return
																									} else {
																										v215 = F_errstart(m, int32(14), int32(0))
																										mBase = m.M
																										v216 = m.ExcPending
																										if v216 != 0 {
																											return
																										} else {
																											if v215 != 0 {
																												v218 = *(*int32)(unsafe.Add(mBase, _consts[527]))
																												v219 = *(*int32)(unsafe.Add(mBase, uint32(v218)+16))
																												v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+28)))
																												switch v221 - int32(100) {
																												case 0:
																													v230 = int32(556578)
																												default:
																													if v221 == int32(101) {
																														v228 = int32(556587)
																													} else {
																														v228 = int32(558895)
																													}
																													v230 = v228
																												case 12:
																													v230 = int32(549382)
																												}
																												*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v230
																												*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v219
																												F_errmsg_internal(m, int32(184241), v16+int32(16))
																												mBase = m.M
																												v237 = m.ExcPending
																												if v237 != 0 {
																													return
																												} else {
																													F_errfinish(m, int32(506994), int32(4645), int32(224471))
																													mBase = m.M
																													v242 = m.ExcPending
																													if v242 != 0 {
																														return
																													} else {
																														F_start_apply(m, v50)
																														mBase = m.M
																														v247 = m.ExcPending
																														if v247 != 0 {
																															return
																														} else {
																															m.G0 = v16 + int32(160)
																															F_proc_exit(m, int32(0))
																															mBase = m.M
																															v290 = m.ExcPending
																															if v290 != 0 {
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
																												F_start_apply(m, v50)
																												mBase = m.M
																												v247 = m.ExcPending
																												if v247 != 0 {
																													return
																												} else {
																													m.G0 = v16 + int32(160)
																													F_proc_exit(m, int32(0))
																													mBase = m.M
																													v290 = m.ExcPending
																													if v290 != 0 {
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
			v254 = m.ExcPending
			if v254 != 0 {
				return
			} else {
				F_errcode(m, int32(325))
				mBase = m.M
				v257 = m.ExcPending
				if v257 != 0 {
					return
				} else {
					F_errmsg(m, int32(108461), int32(0))
					mBase = m.M
					v261 = m.ExcPending
					if v261 != 0 {
						return
					} else {
						F_errfinish(m, int32(506994), int32(4567), int32(224471))
						mBase = m.M
						v266 = m.ExcPending
						if v266 != 0 {
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
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v14 = *(*int32)(unsafe.Add(mBase, _consts[281]))
	F_hash_seq_init(m, v9+int32(12), v14)
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
	v19 = F_hash_seq_search(m, v9+int32(12))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v19 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v24 = v19
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
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v24)+64))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
	if l0 != v28 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L6
L9:
	;
	v122 = F_hash_seq_search(m, v9+int32(12))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L49
	}
L10:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)+24))
	if v30 != l0 {
		goto L9
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v27)+80))
	if v87&int32(-2) == int32(2) {
		goto L37
	} else {
		goto L38
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+24)) = l1
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v27)+80))
	if v33 == int32(3) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	if v51 == int32(0) {
		goto L9
	} else {
		goto L21
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+80)) = int32(5)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
	if v38 == int32(0) {
		goto L14
	} else {
		goto L18
	}
L16:
	;
	v47 = v33
	goto L17
L17:
	;
	if v47 != int32(5) {
		goto L9
	} else {
		goto L20
	}
L18:
	;
	m.T0[v38].(func(*base.Module, int32))(m, v27)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = int32(0)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v27)+80))
	v47 = v45
	goto L17
L20:
	;
	goto L14
L21:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	if v56 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+12)) = int32(0)
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
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	if v59 == v51 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v51)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+4)) = v61
	goto L23
L26:
	;
	goto L27
L27:
	;
	v66 = v59
	goto L28
L28:
	;
	if v66 == int32(0) {
		goto L23
	} else {
		goto L30
	}
L29:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v51)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v66)+8)) = v71
	goto L23
L30:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v66)+8))
	if v51 != v69 {
		v66 = v69
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51))) = l2
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v51)+8)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v51
	goto L22
L33:
	;
	goto L34
L34:
	;
	v81 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v51)+8)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v51))) = v81
	goto L22
L35:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v27)+60))
	if v108 != 0 {
		goto L44
	} else {
		goto L45
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = int32(0)
	goto L35
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+80)) = int32(5)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
	if v94 == int32(0) {
		goto L35
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
	if v99 == int32(0) {
		goto L35
	} else {
		goto L42
	}
L40:
	;
	m.T0[v94].(func(*base.Module, int32))(m, v27)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	goto L36
L42:
	;
	m.T0[v99].(func(*base.Module, int32))(m, v27)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	goto L36
L44:
	;
	F_ReleaseCachedPlan(m, v108, int32(0))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+12)) = int32(0)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	F_MemoryContextDeleteChildren(m, v116)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L48
	}
L47:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v27)+56)) = int64(0)
	goto L46
L48:
	;
	goto L9
L49:
	;
	if v122 != 0 {
		v24 = v122
		goto L7
	} else {
		goto L50
	}
L50:
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
			F_errmsg_internal(m, int32(249588), int32(0))
			mBase = m.M
			v8 = m.ExcPending
			if v8 != 0 {
				return
			} else {
				F_errfinish(m, int32(509000), int32(795), int32(249207))
				mBase = m.M
				v13 = m.ExcPending
				if v13 != 0 {
					return
				} else {
					v15 = *(*int32)(unsafe.Add(mBase, _consts[399]))
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
			v15 = *(*int32)(unsafe.Add(mBase, _consts[399]))
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
func F_a_swap(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
	return v3
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
				v24 = F_AllocSetContextCreateInternal(m, l4, int32(100001), int32(0), int32(8192), int32(8388608))
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v142 int32
	_ = v142
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
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v224 int32
	_ = v224
	var v232 int32
	_ = v232
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v413 int32
	_ = v413
	v3 = int32(0)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v17 == int32(9999) {
		v21 = v16 - int32(1)
		if v21 < int32(0) {
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v21 == int32(0) {
				v85 = v3
				v86 = v21
				v88 = v3
			} else {
				v36 = v21
				v38 = v3
				v39 = v3
				for {
					v49 = v36 << (uint(int32(2)) % 32)
					v50 = v24 + v49
					v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
					v52 = v51 + v38
					if v52 < int32(10000) {
						v61 = v52
						v62 = int32(0)
					} else {
						v57 = base.I32_div_u_s(v52, int32(10000))
						v61 = v57*int32(-10000) + v52
						v62 = v57
					}
					*(*int32)(unsafe.Add(mBase, uint32(v50))) = v61
					v65 = v49 + (v24 - int32(4))
					v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
					v67 = v66 + v62
					if int32(10000) <= v67 {
						v71 = base.I32_div_u_s(v67, int32(10000))
						v75 = v71*int32(-10000) + v67
						v76 = v71
					} else {
						v75 = v67
						v76 = int32(0)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v65))) = v75
					v78 = int32(2)
					v79 = v36 - v78
					v81 = v39 + v78
					if v81 != v16&int32(-2) {
						v36 = v79
						v38 = v76
						v39 = v81
						continue
					} else {
						break
					}
					break
				}
				v85 = v75
				v86 = v79
				v88 = v76
			}
			if v16&int32(1) != 0 {
				v100 = v24 + v86<<(uint(int32(2))%32)
				v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
				v102 = v101 + v88
				v104 = base.I32_rem_u_s(v102, int32(10000))
				if int32(9999) < v102 {
					v107 = v104
				} else {
					v107 = v102
				}
				*(*int32)(unsafe.Add(mBase, uint32(v100))) = v107
				v109 = v107
			} else {
				v109 = v85
			}
			if int32(0) < v109 {
				v112 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v112)
			} else {
			}
			v114 = int32(1)
			v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			if v16 == v114 {
				v178 = v109
				v179 = int32(0)
				v180 = v21
			} else {
				v124 = int32(0)
				v129 = v124
				v130 = v21
				v131 = v124
				for {
					v142 = v130 << (uint(int32(2)) % 32)
					v143 = v116 + v142
					v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
					v145 = v144 + v129
					if v145 < int32(10000) {
						v154 = v145
						v155 = int32(0)
					} else {
						v150 = base.I32_div_u_s(v145, int32(10000))
						v154 = v150*int32(-10000) + v145
						v155 = v150
					}
					*(*int32)(unsafe.Add(mBase, uint32(v143))) = v154
					v158 = v142 + (v116 - int32(4))
					v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
					v160 = v159 + v155
					if int32(10000) <= v160 {
						v164 = base.I32_div_u_s(v160, int32(10000))
						v168 = v164*int32(-10000) + v160
						v169 = v164
					} else {
						v168 = v160
						v169 = int32(0)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v158))) = v168
					v171 = int32(2)
					v172 = v130 - v171
					v174 = v131 + v171
					if v174 != v16&int32(-2) {
						v129 = v169
						v130 = v172
						v131 = v174
						continue
					} else {
						break
					}
					break
				}
				v178 = v168
				v179 = v169
				v180 = v172
			}
			if v16&v114 != 0 {
				v193 = v116 + v180<<(uint(int32(2))%32)
				v194 = *(*int32)(unsafe.Add(mBase, uint32(v193)))
				v195 = v194 + v179
				v197 = base.I32_rem_u_s(v195, int32(10000))
				if int32(9999) < v195 {
					v200 = v197
				} else {
					v200 = v195
				}
				*(*int32)(unsafe.Add(mBase, uint32(v193))) = v200
				v202 = v200
			} else {
				v202 = v178
			}
			if v202 <= int32(0) {
			} else {
				v205 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v205)
			}
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
		v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v232 = v224
	} else {
		v232 = v16
	}
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v241 <= v240 {
		v244 = v240 + int32(1)
		v252 = v244
		v254 = v244 + (v232 - v241)
	} else {
		v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
		if v247 != 0 {
			v252 = v241
			v254 = v232
		} else {
			v248 = int32(1)
			v252 = v241 + v248
			v254 = v232 + v248
		}
	}
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v256 = int32(-1)
	v258 = v255 + (v240 ^ v256)
	v261 = v254 + (v252 ^ v256)
	if v261 < v258 {
		v265 = v258 - v261
	} else {
		v265 = int32(0)
	}
	v266 = v254 + v265
	if v252 != v241 {
		v270 = v266 << (uint(int32(2)) % 32)
		v271 = F_palloc0(m, v270)
		mBase = m.M
		v272 = m.ExcPending
		if v272 != 0 {
			return
		} else {
			v273 = F_palloc0(m, v270)
			mBase = m.M
			v274 = m.ExcPending
			if v274 != 0 {
				return
			} else {
				v275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v275 != 0 {
					v277 = int32(2)
					v278 = (v252 - v241) << (uint(v277) % 32)
					v281 = v232 << (uint(v277) % 32)
					if v281 != 0 {
						v282 = F__emscripten_memcpy_bulkmem(m, v271+v278, v275, v281)
						mBase = m.M
					} else {
					}
					v284 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					F_pfree(m, v284)
					mBase = m.M
					v286 = m.ExcPending
					if v286 != 0 {
						return
					} else {
						v288 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						if v281 != 0 {
							v289 = F__emscripten_memcpy_bulkmem(m, v278+v273, v288, v281)
							mBase = m.M
						} else {
						}
						v291 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						F_pfree(m, v291)
						mBase = m.M
						v293 = m.ExcPending
						if v293 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v273
							*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v271
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v252
							v299 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v299)
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = v266
							v302 = v252
							v308 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
							v309 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							if v309 < v308 {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v308
							} else {
							}
							v312 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							if v312 <= int32(0) {
							} else {
								v317 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
								if v317 != 0 {
									v318 = int32(24)
								} else {
									v318 = int32(20)
								}
								v320 = *(*int32)(unsafe.Add(mBase, uint32(l0+v318)))
								v321 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
								v322 = int32(1)
								v324 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
								v325 = v302 - v324
								v326 = int32(0)
								if v312 != v322 {
									v336 = v326
									v337 = v325
									v338 = int32(0)
									for {
										v349 = int32(2)
										v350 = v337 << (uint(v349) % 32)
										v351 = v320 + v350
										v352 = *(*int32)(unsafe.Add(mBase, uint32(v351)))
										v355 = v321 + v336<<(uint(int32(1))%32)
										v356 = int32(*(*int16)(unsafe.Add(mBase, uint32(v355))))
										*(*int32)(unsafe.Add(mBase, uint32(v351))) = v352 + v356
										v359 = v350 + (v320 + int32(4))
										v360 = *(*int32)(unsafe.Add(mBase, uint32(v359)))
										v361 = int32(*(*int16)(unsafe.Add(mBase, uint32(v355)+2)))
										*(*int32)(unsafe.Add(mBase, uint32(v359))) = v360 + v361
										v365 = v336 + v349
										v367 = v337 + v349
										v369 = v338 + v349
										if v369 != v312&int32(2147483646) {
											v336 = v365
											v337 = v367
											v338 = v369
											continue
										} else {
											break
										}
										break
									}
									v373 = v365
									v374 = v367
								} else {
									v373 = v326
									v374 = v325
								}
								if v312&v322 == int32(0) {
								} else {
									v390 = v320 + v374<<(uint(int32(2))%32)
									v391 = *(*int32)(unsafe.Add(mBase, uint32(v390)))
									v395 = int32(*(*int16)(unsafe.Add(mBase, uint32(v321+v373<<(uint(int32(1))%32)))))
									*(*int32)(unsafe.Add(mBase, uint32(v390))) = v391 + v395
								}
							}
							v413 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v413 + int32(1)
							return
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v273
					*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v271
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v252
					v299 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v299)
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = v266
					v302 = v252
					v308 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					v309 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					if v309 < v308 {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v308
					} else {
					}
					v312 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					if v312 <= int32(0) {
					} else {
						v317 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						if v317 != 0 {
							v318 = int32(24)
						} else {
							v318 = int32(20)
						}
						v320 = *(*int32)(unsafe.Add(mBase, uint32(l0+v318)))
						v321 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
						v322 = int32(1)
						v324 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						v325 = v302 - v324
						v326 = int32(0)
						if v312 != v322 {
							v336 = v326
							v337 = v325
							v338 = int32(0)
							for {
								v349 = int32(2)
								v350 = v337 << (uint(v349) % 32)
								v351 = v320 + v350
								v352 = *(*int32)(unsafe.Add(mBase, uint32(v351)))
								v355 = v321 + v336<<(uint(int32(1))%32)
								v356 = int32(*(*int16)(unsafe.Add(mBase, uint32(v355))))
								*(*int32)(unsafe.Add(mBase, uint32(v351))) = v352 + v356
								v359 = v350 + (v320 + int32(4))
								v360 = *(*int32)(unsafe.Add(mBase, uint32(v359)))
								v361 = int32(*(*int16)(unsafe.Add(mBase, uint32(v355)+2)))
								*(*int32)(unsafe.Add(mBase, uint32(v359))) = v360 + v361
								v365 = v336 + v349
								v367 = v337 + v349
								v369 = v338 + v349
								if v369 != v312&int32(2147483646) {
									v336 = v365
									v337 = v367
									v338 = v369
									continue
								} else {
									break
								}
								break
							}
							v373 = v365
							v374 = v367
						} else {
							v373 = v326
							v374 = v325
						}
						if v312&v322 == int32(0) {
						} else {
							v390 = v320 + v374<<(uint(int32(2))%32)
							v391 = *(*int32)(unsafe.Add(mBase, uint32(v390)))
							v395 = int32(*(*int16)(unsafe.Add(mBase, uint32(v321+v373<<(uint(int32(1))%32)))))
							*(*int32)(unsafe.Add(mBase, uint32(v390))) = v391 + v395
						}
					}
					v413 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v413 + int32(1)
					return
				}
			}
		}
	} else {
		if v266 != v232 {
			v270 = v266 << (uint(int32(2)) % 32)
			v271 = F_palloc0(m, v270)
			mBase = m.M
			v272 = m.ExcPending
			if v272 != 0 {
				return
			} else {
				v273 = F_palloc0(m, v270)
				mBase = m.M
				v274 = m.ExcPending
				if v274 != 0 {
					return
				} else {
					v275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v275 != 0 {
						v277 = int32(2)
						v278 = (v252 - v241) << (uint(v277) % 32)
						v281 = v232 << (uint(v277) % 32)
						if v281 != 0 {
							v282 = F__emscripten_memcpy_bulkmem(m, v271+v278, v275, v281)
							mBase = m.M
						} else {
						}
						v284 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						F_pfree(m, v284)
						mBase = m.M
						v286 = m.ExcPending
						if v286 != 0 {
							return
						} else {
							v288 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							if v281 != 0 {
								v289 = F__emscripten_memcpy_bulkmem(m, v278+v273, v288, v281)
								mBase = m.M
							} else {
							}
							v291 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							F_pfree(m, v291)
							mBase = m.M
							v293 = m.ExcPending
							if v293 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v273
								*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v271
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v252
								v299 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v299)
								*(*int32)(unsafe.Add(mBase, uint32(l0))) = v266
								v302 = v252
								v308 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
								v309 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								if v309 < v308 {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v308
								} else {
								}
								v312 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								if v312 <= int32(0) {
								} else {
									v317 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
									if v317 != 0 {
										v318 = int32(24)
									} else {
										v318 = int32(20)
									}
									v320 = *(*int32)(unsafe.Add(mBase, uint32(l0+v318)))
									v321 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
									v322 = int32(1)
									v324 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
									v325 = v302 - v324
									v326 = int32(0)
									if v312 != v322 {
										v336 = v326
										v337 = v325
										v338 = int32(0)
										for {
											v349 = int32(2)
											v350 = v337 << (uint(v349) % 32)
											v351 = v320 + v350
											v352 = *(*int32)(unsafe.Add(mBase, uint32(v351)))
											v355 = v321 + v336<<(uint(int32(1))%32)
											v356 = int32(*(*int16)(unsafe.Add(mBase, uint32(v355))))
											*(*int32)(unsafe.Add(mBase, uint32(v351))) = v352 + v356
											v359 = v350 + (v320 + int32(4))
											v360 = *(*int32)(unsafe.Add(mBase, uint32(v359)))
											v361 = int32(*(*int16)(unsafe.Add(mBase, uint32(v355)+2)))
											*(*int32)(unsafe.Add(mBase, uint32(v359))) = v360 + v361
											v365 = v336 + v349
											v367 = v337 + v349
											v369 = v338 + v349
											if v369 != v312&int32(2147483646) {
												v336 = v365
												v337 = v367
												v338 = v369
												continue
											} else {
												break
											}
											break
										}
										v373 = v365
										v374 = v367
									} else {
										v373 = v326
										v374 = v325
									}
									if v312&v322 == int32(0) {
									} else {
										v390 = v320 + v374<<(uint(int32(2))%32)
										v391 = *(*int32)(unsafe.Add(mBase, uint32(v390)))
										v395 = int32(*(*int16)(unsafe.Add(mBase, uint32(v321+v373<<(uint(int32(1))%32)))))
										*(*int32)(unsafe.Add(mBase, uint32(v390))) = v391 + v395
									}
								}
								v413 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v413 + int32(1)
								return
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v273
						*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v271
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v252
						v299 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v299)
						*(*int32)(unsafe.Add(mBase, uint32(l0))) = v266
						v302 = v252
						v308 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
						v309 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						if v309 < v308 {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v308
						} else {
						}
						v312 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						if v312 <= int32(0) {
						} else {
							v317 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
							if v317 != 0 {
								v318 = int32(24)
							} else {
								v318 = int32(20)
							}
							v320 = *(*int32)(unsafe.Add(mBase, uint32(l0+v318)))
							v321 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
							v322 = int32(1)
							v324 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
							v325 = v302 - v324
							v326 = int32(0)
							if v312 != v322 {
								v336 = v326
								v337 = v325
								v338 = int32(0)
								for {
									v349 = int32(2)
									v350 = v337 << (uint(v349) % 32)
									v351 = v320 + v350
									v352 = *(*int32)(unsafe.Add(mBase, uint32(v351)))
									v355 = v321 + v336<<(uint(int32(1))%32)
									v356 = int32(*(*int16)(unsafe.Add(mBase, uint32(v355))))
									*(*int32)(unsafe.Add(mBase, uint32(v351))) = v352 + v356
									v359 = v350 + (v320 + int32(4))
									v360 = *(*int32)(unsafe.Add(mBase, uint32(v359)))
									v361 = int32(*(*int16)(unsafe.Add(mBase, uint32(v355)+2)))
									*(*int32)(unsafe.Add(mBase, uint32(v359))) = v360 + v361
									v365 = v336 + v349
									v367 = v337 + v349
									v369 = v338 + v349
									if v369 != v312&int32(2147483646) {
										v336 = v365
										v337 = v367
										v338 = v369
										continue
									} else {
										break
									}
									break
								}
								v373 = v365
								v374 = v367
							} else {
								v373 = v326
								v374 = v325
							}
							if v312&v322 == int32(0) {
							} else {
								v390 = v320 + v374<<(uint(int32(2))%32)
								v391 = *(*int32)(unsafe.Add(mBase, uint32(v390)))
								v395 = int32(*(*int16)(unsafe.Add(mBase, uint32(v321+v373<<(uint(int32(1))%32)))))
								*(*int32)(unsafe.Add(mBase, uint32(v390))) = v391 + v395
							}
						}
						v413 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v413 + int32(1)
						return
					}
				}
			}
		} else {
			v302 = v241
			v308 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v309 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if v309 < v308 {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v308
			} else {
			}
			v312 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			if v312 <= int32(0) {
			} else {
				v317 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				if v317 != 0 {
					v318 = int32(24)
				} else {
					v318 = int32(20)
				}
				v320 = *(*int32)(unsafe.Add(mBase, uint32(l0+v318)))
				v321 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
				v322 = int32(1)
				v324 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				v325 = v302 - v324
				v326 = int32(0)
				if v312 != v322 {
					v336 = v326
					v337 = v325
					v338 = int32(0)
					for {
						v349 = int32(2)
						v350 = v337 << (uint(v349) % 32)
						v351 = v320 + v350
						v352 = *(*int32)(unsafe.Add(mBase, uint32(v351)))
						v355 = v321 + v336<<(uint(int32(1))%32)
						v356 = int32(*(*int16)(unsafe.Add(mBase, uint32(v355))))
						*(*int32)(unsafe.Add(mBase, uint32(v351))) = v352 + v356
						v359 = v350 + (v320 + int32(4))
						v360 = *(*int32)(unsafe.Add(mBase, uint32(v359)))
						v361 = int32(*(*int16)(unsafe.Add(mBase, uint32(v355)+2)))
						*(*int32)(unsafe.Add(mBase, uint32(v359))) = v360 + v361
						v365 = v336 + v349
						v367 = v337 + v349
						v369 = v338 + v349
						if v369 != v312&int32(2147483646) {
							v336 = v365
							v337 = v367
							v338 = v369
							continue
						} else {
							break
						}
						break
					}
					v373 = v365
					v374 = v367
				} else {
					v373 = v326
					v374 = v325
				}
				if v312&v322 == int32(0) {
				} else {
					v390 = v320 + v374<<(uint(int32(2))%32)
					v391 = *(*int32)(unsafe.Add(mBase, uint32(v390)))
					v395 = int32(*(*int16)(unsafe.Add(mBase, uint32(v321+v373<<(uint(int32(1))%32)))))
					*(*int32)(unsafe.Add(mBase, uint32(v390))) = v391 + v395
				}
			}
			v413 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v413 + int32(1)
			return
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
					F_errmsg_internal(m, int32(495085), v7+int32(32))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return
					} else {
						F_errfinish(m, int32(509618), int32(2790), int32(216318))
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
				v61 = int32(191723)
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
							F_errfinish(m, int32(509618), int32(2795), int32(216318))
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
				v61 = int32(188660)
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
							F_errfinish(m, int32(509618), int32(2795), int32(216318))
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
				v61 = int32(187919)
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
							F_errfinish(m, int32(509618), int32(2795), int32(216318))
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
				v61 = int32(188218)
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
							F_errfinish(m, int32(509618), int32(2795), int32(216318))
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
				v61 = int32(191839)
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
							F_errfinish(m, int32(509618), int32(2795), int32(216318))
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
				v61 = int32(188784)
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
							F_errfinish(m, int32(509618), int32(2795), int32(216318))
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
				v61 = int32(186307)
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
							F_errfinish(m, int32(509618), int32(2795), int32(216318))
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
				v61 = int32(188354)
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
							F_errfinish(m, int32(509618), int32(2795), int32(216318))
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
				v61 = int32(186083)
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
							F_errfinish(m, int32(509618), int32(2795), int32(216318))
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
				v61 = int32(185889)
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
							F_errfinish(m, int32(509618), int32(2795), int32(216318))
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
				v61 = int32(199191)
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
							F_errfinish(m, int32(509618), int32(2795), int32(216318))
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
				v61 = int32(187683)
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
							F_errfinish(m, int32(509618), int32(2795), int32(216318))
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
				v61 = int32(182013)
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
							F_errfinish(m, int32(509618), int32(2795), int32(216318))
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
				v61 = int32(199506)
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
							F_errfinish(m, int32(509618), int32(2795), int32(216318))
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
				v61 = int32(183351)
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
							F_errfinish(m, int32(509618), int32(2795), int32(216318))
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
				v61 = int32(182265)
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
							F_errfinish(m, int32(509618), int32(2795), int32(216318))
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
				v61 = int32(183956)
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
							F_errfinish(m, int32(509618), int32(2795), int32(216318))
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
				v61 = int32(185189)
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
							F_errfinish(m, int32(509618), int32(2795), int32(216318))
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
				v61 = int32(181821)
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
							F_errfinish(m, int32(509618), int32(2795), int32(216318))
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
				v61 = int32(186026)
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
							F_errfinish(m, int32(509618), int32(2795), int32(216318))
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
				v61 = int32(181924)
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
							F_errfinish(m, int32(509618), int32(2795), int32(216318))
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
				v61 = int32(191920)
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
							F_errfinish(m, int32(509618), int32(2795), int32(216318))
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
				v61 = int32(188052)
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
							F_errfinish(m, int32(509618), int32(2795), int32(216318))
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
				v61 = int32(197885)
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
							F_errfinish(m, int32(509618), int32(2795), int32(216318))
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
				v61 = int32(201489)
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
							F_errfinish(m, int32(509618), int32(2795), int32(216318))
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
				v61 = int32(199846)
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
							F_errfinish(m, int32(509618), int32(2795), int32(216318))
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
				v61 = int32(187521)
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
							F_errfinish(m, int32(509618), int32(2795), int32(216318))
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
				v61 = int32(183270)
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
							F_errfinish(m, int32(509618), int32(2795), int32(216318))
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
				v61 = int32(199160)
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
							F_errfinish(m, int32(509618), int32(2795), int32(216318))
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
				v61 = int32(199943)
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
							F_errfinish(m, int32(509618), int32(2795), int32(216318))
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
				v61 = int32(187822)
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
							F_errfinish(m, int32(509618), int32(2795), int32(216318))
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
				v61 = int32(181700)
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
							F_errfinish(m, int32(509618), int32(2795), int32(216318))
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
				v61 = int32(194184)
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
							F_errfinish(m, int32(509618), int32(2795), int32(216318))
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
				v61 = int32(182210)
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
							F_errfinish(m, int32(509618), int32(2795), int32(216318))
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
				v61 = int32(558629)
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
							F_errfinish(m, int32(509618), int32(2795), int32(216318))
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
					F_errmsg_internal(m, int32(495085), v7-int32(-64))
					mBase = m.M
					v120 = m.ExcPending
					if v120 != 0 {
						return
					} else {
						F_errfinish(m, int32(509618), int32(2928), int32(216318))
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
				v127 = int32(191758)
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
							F_errfinish(m, int32(509618), int32(2933), int32(216318))
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
				v127 = int32(187984)
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
							F_errfinish(m, int32(509618), int32(2933), int32(216318))
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
				v127 = int32(187954)
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
							F_errfinish(m, int32(509618), int32(2933), int32(216318))
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
				v127 = int32(188254)
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
							F_errfinish(m, int32(509618), int32(2933), int32(216318))
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
				v127 = int32(191873)
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
							F_errfinish(m, int32(509618), int32(2933), int32(216318))
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
				v127 = int32(188816)
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
							F_errfinish(m, int32(509618), int32(2933), int32(216318))
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
				v127 = int32(186346)
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
							F_errfinish(m, int32(509618), int32(2933), int32(216318))
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
				v127 = int32(188389)
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
							F_errfinish(m, int32(509618), int32(2933), int32(216318))
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
				v127 = int32(186129)
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
							F_errfinish(m, int32(509618), int32(2933), int32(216318))
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
				v127 = int32(185929)
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
							F_errfinish(m, int32(509618), int32(2933), int32(216318))
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
				v127 = int32(199230)
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
							F_errfinish(m, int32(509618), int32(2933), int32(216318))
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
				v127 = int32(187772)
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
							F_errfinish(m, int32(509618), int32(2933), int32(216318))
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
				v127 = int32(182044)
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
							F_errfinish(m, int32(509618), int32(2933), int32(216318))
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
				v127 = int32(199540)
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
							F_errfinish(m, int32(509618), int32(2933), int32(216318))
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
				v127 = int32(183389)
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
							F_errfinish(m, int32(509618), int32(2933), int32(216318))
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
				v127 = int32(182308)
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
							F_errfinish(m, int32(509618), int32(2933), int32(216318))
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
				v127 = int32(183996)
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
							F_errfinish(m, int32(509618), int32(2933), int32(216318))
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
				v127 = int32(185410)
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
							F_errfinish(m, int32(509618), int32(2933), int32(216318))
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
				v127 = int32(181862)
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
							F_errfinish(m, int32(509618), int32(2933), int32(216318))
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
				v127 = int32(191955)
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
							F_errfinish(m, int32(509618), int32(2933), int32(216318))
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
				v127 = int32(188168)
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
							F_errfinish(m, int32(509618), int32(2933), int32(216318))
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
				v127 = int32(197918)
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
							F_errfinish(m, int32(509618), int32(2933), int32(216318))
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
				v127 = int32(201850)
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
							F_errfinish(m, int32(509618), int32(2933), int32(216318))
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
				v127 = int32(199880)
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
							F_errfinish(m, int32(509618), int32(2933), int32(216318))
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
				v127 = int32(187559)
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
							F_errfinish(m, int32(509618), int32(2933), int32(216318))
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
				v127 = int32(183313)
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
							F_errfinish(m, int32(509618), int32(2933), int32(216318))
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
				v127 = int32(199264)
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
							F_errfinish(m, int32(509618), int32(2933), int32(216318))
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
				v127 = int32(199979)
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
							F_errfinish(m, int32(509618), int32(2933), int32(216318))
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
				v127 = int32(187873)
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
							F_errfinish(m, int32(509618), int32(2933), int32(216318))
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
				v127 = int32(181748)
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
							F_errfinish(m, int32(509618), int32(2933), int32(216318))
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
				v127 = int32(196473)
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
							F_errfinish(m, int32(509618), int32(2933), int32(216318))
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
				v127 = int32(182240)
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
							F_errfinish(m, int32(509618), int32(2933), int32(216318))
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
				v127 = int32(558629)
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
							F_errfinish(m, int32(509618), int32(2933), int32(216318))
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
				F_errmsg_internal(m, int32(491723), v7)
				mBase = m.M
				v152 = m.ExcPending
				if v152 != 0 {
					return
				} else {
					F_errfinish(m, int32(509618), int32(2937), int32(216318))
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
	var v34 int32
	_ = v34
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
	v34 = int32(0)
	goto L10
L8:
	;
	goto L9
L9:
	;
	return int32(0)
L10:
	;
	v42 = v27 + v9 + v34<<(uint(int32(4))%32)
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
	v56 = v34 + int32(1)
	if v56 != v16 {
		v34 = v56
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
			F_pg_qsort(m, v17+l0, v6, int32(16), int32(1252))
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
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int64
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v132 int32
	_ = v132
	var v133 int64
	_ = v133
	var v137 int64
	_ = v137
	var v138 int32
	_ = v138
	var v139 int64
	_ = v139
	var v146 int32
	_ = v146
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
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
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v278 int32
	_ = v278
	var v288 int32
	_ = v288
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
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
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v336 int32
	_ = v336
	var v337 int64
	_ = v337
	var v340 int64
	_ = v340
	var v342 int64
	_ = v342
	var v346 int64
	_ = v346
	var v347 int64
	_ = v347
	var v349 int64
	_ = v349
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v399 int32
	_ = v399
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v414 int32
	_ = v414
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v434 int32
	_ = v434
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v491 int32
	_ = v491
	var v504 int32
	_ = v504
	var v514 int64
	_ = v514
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v524 int64
	_ = v524
	var v525 int32
	_ = v525
	var v528 int64
	_ = v528
	var v532 int64
	_ = v532
	var v542 int32
	_ = v542
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v554 int32
	_ = v554
	var v561 int32
	_ = v561
	var v571 int32
	_ = v571
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v585 int64
	_ = v585
	var v592 int32
	_ = v592
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v611 int32
	_ = v611
	var v626 int32
	_ = v626
	var v630 int32
	_ = v630
	var v635 int32
	_ = v635
	var v639 int32
	_ = v639
	var v642 int32
	_ = v642
	var v646 int32
	_ = v646
	var v650 int32
	_ = v650
	var v655 int32
	_ = v655
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
	v639 = m.ExcPending
	if v639 != 0 {
		goto L1
	} else {
		goto L149
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L1
	} else {
		goto L146
	}
L5:
	;
	v336 = v320 + v325<<(uint(int32(4))%32)
	v337 = *(*int64)(unsafe.Add(mBase, uint32(v336)+8))
	switch l2 - int32(1) {
	case 0:
		goto L75
	case 1:
		goto L74
	case 2:
		goto L73
	default:
		v349 = v337
		goto L71
	}
L6:
	;
	v288 = v161 + int32(1)
	if v288 < int32(0) {
		goto L4
	} else {
		goto L65
	}
L7:
	;
	if v263 == v161 {
		v278 = v263
		goto L6
	} else {
		goto L64
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L60
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L57
	}
L10:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v162 != 0 {
		goto L40
	} else {
		goto L41
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
	*(*int32)(unsafe.Add(mBase, uint32(v39)+12)) = int32(1033)
	*(*int64)(unsafe.Add(mBase, uint32(v39)+4)) = int64(1)
	v45 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = v38 << (uint(v45) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+20)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+16)) = v32
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v53 = int32(base.Ui32(v51) >> (uint(v45) % 32))
	if v53 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v65 = v39
	goto L21
L18:
	;
	v54 = F__emscripten_memcpy_bulkmem(m, v39, l0, v53)
	mBase = m.M
	goto L20
L19:
	;
	goto L20
L20:
	;
	goto L17
L21:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v65)+16))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v65)+8))
	if v71 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v133 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l1)+12)))
	v137 = F_aclmask(m, v65, v132, l3, v133<<(uint(int64(32))%64), int32(0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L37
	}
L23:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	v81 = (v74<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L25
L24:
	;
	v81 = v71
	goto L25
L25:
	;
	if int32(0) < v70 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v92 = int32(0)
	goto L29
L27:
	;
	goto L28
L28:
	;
	goto L22
L29:
	;
	v103 = v81 + v65 + v92<<(uint(int32(4))%32)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	if v104 != v85 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	goto L28
L31:
	;
	v116 = v92 + int32(1)
	if v116 != v70 {
		v92 = v116
		goto L29
	} else {
		goto L36
	}
L32:
	;
	v106 = *(*int64)(unsafe.Add(mBase, uint32(v103)+8))
	if base.Ui64(v106) < base.Ui64(int64(4294967296)) {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v111 = F_aclupdate(m, v65, v103, int32(2), l3, int32(1))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	F_pfree(m, v65)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v65 = v111
	goto L21
L36:
	;
	goto L30
L37:
	;
	v139 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	if base.Ui64(int64(4294967296)) <= base.Ui64(v139&(v137^int64(-1))) {
		goto L8
	} else {
		goto L38
	}
L38:
	;
	F_pfree(m, v65)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	goto L10
L40:
	;
	v170 = v162
	goto L42
L41:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v170 = (v163<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L42
L42:
	;
	v171 = v170 + l0
	if v161 <= int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v174 = int32(0)
	v258 = v174
	v263 = v174
	v264 = v174
	goto L7
L44:
	;
	goto L45
L45:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v184 = int32(0)
	goto L46
L46:
	;
	v195 = v171 + v184<<(uint(int32(4))%32)
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v195)))
	if v177 != v196 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v278 = v161
	goto L6
L48:
	;
	v225 = v184 + int32(1)
	if v225 != v161 {
		v184 = v225
		goto L46
	} else {
		goto L56
	}
L49:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v195)+4))
	if v198 != v199 {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v204 = v161<<(uint(int32(4))%32) + int32(24)
	v205 = F_palloc0(m, v204)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v205)+12)) = int32(1033)
	*(*int64)(unsafe.Add(mBase, uint32(v205)+4)) = int64(1)
	v211 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v205))) = v204 << (uint(v211) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v205)+20)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v205)+16)) = v161
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v219 = int32(base.Ui32(v217) >> (uint(v211) % 32))
	if v219 != 0 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v258 = v221 + int32(24)
	v263 = v184
	v264 = v205
	goto L7
L53:
	;
	v220 = F__emscripten_memcpy_bulkmem(m, v205, l0, v219)
	mBase = m.M
	v221 = v220
	goto L55
L54:
	;
	v221 = v205
	goto L55
L55:
	;
	goto L52
L56:
	;
	goto L47
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v32
	F_errmsg_internal(m, int32(493877), v17+int32(16))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(509494), int32(433), int32(314845))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L60:
	;
	F_errcode(m, int32(16910080))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	F_errmsg(m, int32(211918), int32(0))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(509494), int32(1281), int32(11231))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
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
	v320 = v258
	v325 = v263
	v326 = v264
	v328 = v161
	goto L5
L65:
	;
	v294 = v288<<(uint(int32(4))%32) + int32(24)
	v295 = F_palloc0(m, v294)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v295)+12)) = int32(1033)
	*(*int64)(unsafe.Add(mBase, uint32(v295)+4)) = int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v295))) = v294 << (uint(int32(2)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v295)+20)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v295)+16)) = v288
	v308 = v295 + int32(24)
	v310 = v161 << (uint(int32(4)) % 32)
	if v310 != 0 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	v313 = v312 + v310
	v314 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v313))) = v314
	v316 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v313)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v313)+4)) = v316
	v320 = v308
	v325 = v278
	v326 = v295
	v328 = v288
	goto L5
L68:
	;
	v311 = F__emscripten_memcpy_bulkmem(m, v308, v171, v310)
	mBase = m.M
	v312 = v311
	goto L70
L69:
	;
	v312 = v308
	goto L70
L70:
	;
	goto L67
L71:
	;
	if v349 == int64(0) {
		goto L76
	} else {
		goto L77
	}
L72:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v336)+8)) = v347
	v349 = v347
	goto L71
L73:
	;
	v346 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	v347 = v346
	goto L72
L74:
	;
	v342 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	v347 = v337 & (v342 ^ int64(-1))
	goto L72
L75:
	;
	v340 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	v347 = v340 | v337
	goto L72
L76:
	;
	v353 = v336 + int32(16)
	v358 = (v328 + (v325 ^ int32(-1))) << (uint(int32(4)) % 32)
	if v336 == v353 {
		goto L80
	} else {
		goto L81
	}
L77:
	;
	goto L78
L78:
	;
	v514 = v337 & (v349 ^ int64(-1))
	if base.Ui64(v514) < base.Ui64(int64(4294967296)) {
		v611 = v326
		goto L125
	} else {
		goto L126
	}
L79:
	;
	v504 = v328 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v326))) = v504<<(uint(int32(6))%32) + int32(96)
	*(*int32)(unsafe.Add(mBase, uint32(v326)+16)) = v504
	goto L78
L80:
	;
	goto L79
L81:
	;
	v362 = v336 + v358
	if base.Ui32(v353-v362) <= base.Ui32(int32(0)-v358<<(uint(int32(1))%32)) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v369 = F___memcpy(m, v336, v353, v358)
	mBase = m.M
	goto L79
L83:
	;
	goto L84
L84:
	;
	v372 = (v336 ^ v353) & int32(3)
	if base.Ui32(v336) < base.Ui32(v353) {
		goto L87
	} else {
		goto L88
	}
L85:
	;
	if v474 == int32(0) {
		goto L80
	} else {
		goto L121
	}
L86:
	;
	if base.Ui32(v452) <= base.Ui32(int32(3)) {
		v473 = v451
		v474 = v452
		v475 = v453
		goto L85
	} else {
		goto L117
	}
L87:
	;
	if v372 != 0 {
		goto L90
	} else {
		goto L91
	}
L88:
	;
	goto L89
L89:
	;
	if v372 != 0 {
		v434 = v358
		goto L100
	} else {
		goto L101
	}
L90:
	;
	v473 = v353
	v474 = v358
	v475 = v336
	goto L85
L91:
	;
	goto L92
L92:
	;
	if v336&int32(3) == int32(0) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v451 = v353
	v452 = v358
	v453 = v336
	goto L86
L94:
	;
	goto L95
L95:
	;
	v379 = v353
	v380 = v358
	v381 = v336
	goto L96
L96:
	;
	if v380 == int32(0) {
		goto L80
	} else {
		goto L98
	}
L97:
	;
	v451 = v388
	v452 = v390
	v453 = v392
	goto L86
L98:
	;
	v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v379))))
	*(*uint8)(unsafe.Add(mBase, uint32(v381))) = uint8(v385)
	v387 = int32(1)
	v388 = v379 + v387
	v390 = v380 - v387
	v392 = v381 + v387
	if v392&int32(3) != 0 {
		v379 = v388
		v380 = v390
		v381 = v392
		goto L96
	} else {
		goto L99
	}
L99:
	;
	goto L97
L100:
	;
	if v434 == int32(0) {
		goto L80
	} else {
		goto L113
	}
L101:
	;
	if v362&int32(3) != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v399 = v358
	goto L105
L103:
	;
	v414 = v358
	goto L104
L104:
	;
	if base.Ui32(v414) <= base.Ui32(int32(3)) {
		v434 = v414
		goto L100
	} else {
		goto L109
	}
L105:
	;
	if v399 == int32(0) {
		goto L80
	} else {
		goto L107
	}
L106:
	;
	v414 = v405
	goto L104
L107:
	;
	v405 = v399 - int32(1)
	v406 = v336 + v405
	v408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v353+v405))))
	*(*uint8)(unsafe.Add(mBase, uint32(v406))) = uint8(v408)
	if v406&int32(3) != 0 {
		v399 = v405
		goto L105
	} else {
		goto L108
	}
L108:
	;
	goto L106
L109:
	;
	v421 = v414
	goto L110
L110:
	;
	v425 = v421 - int32(4)
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v353+v425)))
	*(*int32)(unsafe.Add(mBase, uint32(v336+v425))) = v428
	if base.Ui32(int32(3)) < base.Ui32(v425) {
		v421 = v425
		goto L110
	} else {
		goto L112
	}
L111:
	;
	v434 = v425
	goto L100
L112:
	;
	goto L111
L113:
	;
	v441 = v434
	goto L114
L114:
	;
	v445 = v441 - int32(1)
	v448 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v353+v445))))
	*(*uint8)(unsafe.Add(mBase, uint32(v336+v445))) = uint8(v448)
	if v445 != 0 {
		v441 = v445
		goto L114
	} else {
		goto L116
	}
L115:
	;
	goto L80
L116:
	;
	goto L115
L117:
	;
	v458 = v451
	v459 = v452
	v460 = v453
	goto L118
L118:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v458)))
	*(*int32)(unsafe.Add(mBase, uint32(v460))) = v462
	v464 = int32(4)
	v465 = v458 + v464
	v467 = v460 + v464
	v469 = v459 - v464
	if base.Ui32(int32(3)) < base.Ui32(v469) {
		v458 = v465
		v459 = v469
		v460 = v467
		goto L118
	} else {
		goto L120
	}
L119:
	;
	v473 = v465
	v474 = v469
	v475 = v467
	goto L85
L120:
	;
	goto L119
L121:
	;
	v480 = v473
	v481 = v474
	v482 = v475
	goto L122
L122:
	;
	v484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v480))))
	*(*uint8)(unsafe.Add(mBase, uint32(v482))) = uint8(v484)
	v486 = int32(1)
	v491 = v481 - v486
	if v491 != 0 {
		v480 = v480 + v486
		v481 = v491
		v482 = v482 + v486
		goto L122
	} else {
		goto L124
	}
L123:
	;
	goto L80
L124:
	;
	goto L123
L125:
	;
	m.G0 = v17 + int32(48)
	return v611
L126:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_check_acl(m, v326)
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	if l3 == v517 {
		v611 = v326
		goto L125
	} else {
		goto L128
	}
L128:
	;
	v524 = F_aclmask(m, v326, v517, l3, v514&int64(-4294967296), int32(0))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	v528 = v514 & (v524 ^ int64(-1))
	if base.Ui64(v528) < base.Ui64(int64(4294967296)) {
		v611 = v326
		goto L125
	} else {
		goto L130
	}
L130:
	;
	v532 = int64(base.Ui64(v528) >> (uint(int64(32)) % 64))
	v542 = v326
	goto L131
L131:
	;
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v542)+16))
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v542)+8))
	if v551 == int32(0) {
		goto L133
	} else {
		goto L134
	}
L132:
	;
	v611 = v542
	goto L125
L133:
	;
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v542)+4))
	v561 = (v554<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L135
L134:
	;
	v561 = v551
	goto L135
L135:
	;
	if v550 <= int32(0) {
		v611 = v542
		goto L125
	} else {
		goto L136
	}
L136:
	;
	v571 = int32(0)
	goto L137
L137:
	;
	v582 = v561 + v542 + v571<<(uint(int32(4))%32)
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v582)+4))
	if v583 != v517 {
		goto L139
	} else {
		goto L140
	}
L138:
	;
	goto L132
L139:
	;
	v603 = v571 + int32(1)
	if v603 != v550 {
		v571 = v603
		goto L137
	} else {
		goto L145
	}
L140:
	;
	v585 = *(*int64)(unsafe.Add(mBase, uint32(v582)+8))
	if v585&v532 == int64(0) {
		goto L139
	} else {
		goto L141
	}
L141:
	;
	if l4 == int32(0) {
		goto L3
	} else {
		goto L142
	}
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = v517
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v582)))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+40)) = v532 | v528&int64(-4294967296)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v592
	v598 = F_aclupdate(m, v542, v17+int32(32), int32(2), l3, l4)
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	F_pfree(m, v542)
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	v542 = v598
	goto L131
L145:
	;
	goto L138
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v288
	F_errmsg_internal(m, int32(493877), v17)
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	F_errfinish(m, int32(509494), int32(433), int32(314845))
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L1
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
	F_errcode(m, int32(16909442))
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	F_errmsg(m, int32(74471), int32(0))
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	F_errhint(m, int32(626259), int32(0))
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	F_errfinish(m, int32(509494), int32(1343), int32(407164))
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L1
	} else {
		goto L153
	}
L153:
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
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = l1
	v16 = F_query_or_expression_tree_mutator_impl(m, l0, int32(1052), v7+int32(4))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		m.G0 = v7 + int32(16)
		return v16
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
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
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
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
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
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v114 float64
	_ = v114
	var v122 int32
	_ = v122
	var v129 float64
	_ = v129
	var v135 float64
	_ = v135
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
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
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
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
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v268 int32
	_ = v268
	var v277 float64
	_ = v277
	var v278 float64
	_ = v278
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
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
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
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
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v544 int32
	_ = v544
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v565 int32
	_ = v565
	var v587 int32
	_ = v587
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v599 int32
	_ = v599
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v617 int32
	_ = v617
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v633 int32
	_ = v633
	var v638 int32
	_ = v638
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v690 int32
	_ = v690
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v741 int32
	_ = v741
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v758 int32
	_ = v758
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v788 int32
	_ = v788
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v811 int32
	_ = v811
	var v813 int32
	_ = v813
	var v815 int32
	_ = v815
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v828 int32
	_ = v828
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v846 int32
	_ = v846
	var v850 int32
	_ = v850
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v863 float64
	_ = v863
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v893 int32
	_ = v893
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v931 int32
	_ = v931
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v971 int32
	_ = v971
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v995 int32
	_ = v995
	var v1000 int32
	_ = v1000
	var v1022 int32
	_ = v1022
	var v1026 int32
	_ = v1026
	var v1029 int32
	_ = v1029
	var v1033 int32
	_ = v1033
	var v1037 int32
	_ = v1037
	var v1041 int32
	_ = v1041
	var v1044 int32
	_ = v1044
	var v1046 int32
	_ = v1046
	var v1049 int32
	_ = v1049
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1053 float64
	_ = v1053
	var v1055 int32
	_ = v1055
	var v1080 float64
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1090 int32
	_ = v1090
	var v1093 int32
	_ = v1093
	var v1095 int32
	_ = v1095
	var v1098 int32
	_ = v1098
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1118 int32
	_ = v1118
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1140 int32
	_ = v1140
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1154 int32
	_ = v1154
	var v1156 int32
	_ = v1156
	var v1160 int32
	_ = v1160
	var v1161 int32
	_ = v1161
	var v1184 int32
	_ = v1184
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1214 int32
	_ = v1214
	var v1215 int32
	_ = v1215
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1221 int32
	_ = v1221
	var v1226 int32
	_ = v1226
	var v1250 int32
	_ = v1250
	var v1254 int32
	_ = v1254
	var v1262 int32
	_ = v1262
	var v1289 int32
	_ = v1289
	var v1311 int32
	_ = v1311
	var v1314 int32
	_ = v1314
	var v1316 int32
	_ = v1316
	var v1318 int32
	_ = v1318
	var v1319 int32
	_ = v1319
	var v1321 int32
	_ = v1321
	var v1349 int32
	_ = v1349
	var v1353 int32
	_ = v1353
	var v1356 int32
	_ = v1356
	var v1357 int32
	_ = v1357
	var v1358 int32
	_ = v1358
	var v1359 int32
	_ = v1359
	var v1362 int32
	_ = v1362
	var v1363 int32
	_ = v1363
	var v1366 int32
	_ = v1366
	var v1367 int32
	_ = v1367
	var v1369 int32
	_ = v1369
	var v1370 int32
	_ = v1370
	var v1379 int32
	_ = v1379
	var v1380 int32
	_ = v1380
	var v1382 int32
	_ = v1382
	var v1385 int32
	_ = v1385
	var v1386 int32
	_ = v1386
	var v1391 int32
	_ = v1391
	var v1398 int32
	_ = v1398
	var v1400 int32
	_ = v1400
	var v1402 int32
	_ = v1402
	var v1403 int32
	_ = v1403
	var v1405 int32
	_ = v1405
	var v1407 int32
	_ = v1407
	var v1411 int32
	_ = v1411
	var v1421 int32
	_ = v1421
	var v1427 int32
	_ = v1427
	var v1428 int32
	_ = v1428
	var v1432 int32
	_ = v1432
	var v1433 int32
	_ = v1433
	var v1437 int32
	_ = v1437
	var v1438 int32
	_ = v1438
	var v1441 int32
	_ = v1441
	var v1465 int32
	_ = v1465
	var v1469 int32
	_ = v1469
	var v1470 int32
	_ = v1470
	var v1474 int32
	_ = v1474
	var v1483 int32
	_ = v1483
	var v1487 int32
	_ = v1487
	var v1491 int32
	_ = v1491
	var v1493 int32
	_ = v1493
	var v1497 int32
	_ = v1497
	var v1498 int32
	_ = v1498
	var v1503 int32
	_ = v1503
	var v1506 int32
	_ = v1506
	var v1509 int32
	_ = v1509
	var v1514 int32
	_ = v1514
	var v1515 int32
	_ = v1515
	var v1519 int32
	_ = v1519
	var v1527 int32
	_ = v1527
	var v1528 int32
	_ = v1528
	var v1539 int32
	_ = v1539
	var v1543 int32
	_ = v1543
	var v1547 int32
	_ = v1547
	var v1549 int32
	_ = v1549
	var v1553 int32
	_ = v1553
	var v1554 int32
	_ = v1554
	var v1559 int32
	_ = v1559
	var v1562 int32
	_ = v1562
	var v1565 int32
	_ = v1565
	var v1570 int32
	_ = v1570
	var v1571 int32
	_ = v1571
	var v1575 int32
	_ = v1575
	var v1583 int32
	_ = v1583
	var v1592 int32
	_ = v1592
	var v1596 int32
	_ = v1596
	var v1600 int32
	_ = v1600
	var v1602 int32
	_ = v1602
	var v1606 int32
	_ = v1606
	var v1607 int32
	_ = v1607
	var v1612 int32
	_ = v1612
	var v1615 int32
	_ = v1615
	var v1618 int32
	_ = v1618
	var v1623 int32
	_ = v1623
	var v1624 int32
	_ = v1624
	var v1628 int32
	_ = v1628
	var v1636 int32
	_ = v1636
	var v1639 int32
	_ = v1639
	var v1640 int32
	_ = v1640
	var v1649 int32
	_ = v1649
	var v1653 int32
	_ = v1653
	var v1657 int32
	_ = v1657
	var v1659 int32
	_ = v1659
	var v1663 int32
	_ = v1663
	var v1664 int32
	_ = v1664
	var v1669 int32
	_ = v1669
	var v1672 int32
	_ = v1672
	var v1675 int32
	_ = v1675
	var v1680 int32
	_ = v1680
	var v1681 int32
	_ = v1681
	var v1685 int32
	_ = v1685
	var v1693 int32
	_ = v1693
	var v1697 int32
	_ = v1697
	var v1701 int32
	_ = v1701
	var v1704 int32
	_ = v1704
	var v1707 int32
	_ = v1707
	var v1708 int32
	_ = v1708
	var v1709 int32
	_ = v1709
	var v1713 int32
	_ = v1713
	var v1714 int32
	_ = v1714
	var v1715 int32
	_ = v1715
	var v1716 int32
	_ = v1716
	var v1717 int32
	_ = v1717
	var v1729 int32
	_ = v1729
	var v1730 int32
	_ = v1730
	var v1732 int32
	_ = v1732
	var v1733 int32
	_ = v1733
	var v1736 int32
	_ = v1736
	var v1747 int32
	_ = v1747
	var v1751 int32
	_ = v1751
	var v1752 int32
	_ = v1752
	var v1753 int32
	_ = v1753
	var v1763 int32
	_ = v1763
	var v1772 int32
	_ = v1772
	var v1775 int32
	_ = v1775
	var v1778 int32
	_ = v1778
	var v1782 int32
	_ = v1782
	var v1786 int32
	_ = v1786
	var v1789 int32
	_ = v1789
	var v1797 int32
	_ = v1797
	var v1804 int32
	_ = v1804
	var v1808 int32
	_ = v1808
	var v1810 int32
	_ = v1810
	var v1814 int32
	_ = v1814
	var v1815 int32
	_ = v1815
	var v1821 int32
	_ = v1821
	var v1824 int32
	_ = v1824
	var v1829 int32
	_ = v1829
	var v1830 int32
	_ = v1830
	var v1845 int32
	_ = v1845
	var v1846 int32
	_ = v1846
	var v1848 int32
	_ = v1848
	var v1849 int32
	_ = v1849
	var v1850 int32
	_ = v1850
	var v1857 int32
	_ = v1857
	var v1864 int32
	_ = v1864
	var v1865 int32
	_ = v1865
	var v1873 int32
	_ = v1873
	var v1879 int32
	_ = v1879
	var v1880 int32
	_ = v1880
	var v1891 int32
	_ = v1891
	var v1900 int32
	_ = v1900
	var v1903 int32
	_ = v1903
	var v1906 int32
	_ = v1906
	var v1910 int32
	_ = v1910
	var v1914 int32
	_ = v1914
	var v1917 int32
	_ = v1917
	var v1925 int32
	_ = v1925
	var v1932 int32
	_ = v1932
	var v1936 int32
	_ = v1936
	var v1938 int32
	_ = v1938
	var v1942 int32
	_ = v1942
	var v1943 int32
	_ = v1943
	var v1949 int32
	_ = v1949
	var v1952 int32
	_ = v1952
	var v1957 int32
	_ = v1957
	var v1958 int32
	_ = v1958
	var v1973 int32
	_ = v1973
	var v1974 int32
	_ = v1974
	var v1976 int32
	_ = v1976
	var v1977 int32
	_ = v1977
	var v1978 int32
	_ = v1978
	var v1985 int32
	_ = v1985
	var v1992 int32
	_ = v1992
	var v1993 int32
	_ = v1993
	var v2001 int32
	_ = v2001
	var v2007 int32
	_ = v2007
	var v2010 int32
	_ = v2010
	var v2011 int32
	_ = v2011
	var v2012 int32
	_ = v2012
	var v2014 float64
	_ = v2014
	var v2017 int32
	_ = v2017
	var v2020 float64
	_ = v2020
	var v2022 float64
	_ = v2022
	var v2023 int32
	_ = v2023
	var v2030 int32
	_ = v2030
	var v2038 int32
	_ = v2038
	var v2041 int32
	_ = v2041
	var v2043 int32
	_ = v2043
	var v2047 int32
	_ = v2047
	var v2048 int32
	_ = v2048
	var v2051 int32
	_ = v2051
	var v2057 int32
	_ = v2057
	var v2064 int32
	_ = v2064
	var v2068 int32
	_ = v2068
	var v2070 int32
	_ = v2070
	var v2074 int32
	_ = v2074
	var v2075 int32
	_ = v2075
	var v2081 int32
	_ = v2081
	var v2084 int32
	_ = v2084
	var v2089 int32
	_ = v2089
	var v2090 int32
	_ = v2090
	var v2103 int32
	_ = v2103
	var v2104 int32
	_ = v2104
	var v2106 int32
	_ = v2106
	var v2108 int32
	_ = v2108
	var v2109 int32
	_ = v2109
	var v2115 int32
	_ = v2115
	var v2121 int32
	_ = v2121
	var v2122 int32
	_ = v2122
	var v2129 int32
	_ = v2129
	var v2134 int32
	_ = v2134
	var v2135 int32
	_ = v2135
	var v2137 int32
	_ = v2137
	var v2140 int32
	_ = v2140
	var v2143 int32
	_ = v2143
	var v2146 int32
	_ = v2146
	var v2149 int32
	_ = v2149
	var v2152 int32
	_ = v2152
	var v2153 int32
	_ = v2153
	var v2154 int32
	_ = v2154
	var v2155 int32
	_ = v2155
	var v2157 int32
	_ = v2157
	var v2160 int32
	_ = v2160
	var v2163 int32
	_ = v2163
	var v2166 int32
	_ = v2166
	var v2169 int32
	_ = v2169
	var v2172 int32
	_ = v2172
	var v2173 int32
	_ = v2173
	var v2174 int32
	_ = v2174
	var v2175 int32
	_ = v2175
	var v2177 int32
	_ = v2177
	var v2178 int32
	_ = v2178
	var v2179 int32
	_ = v2179
	var v2180 int32
	_ = v2180
	var v2183 int32
	_ = v2183
	var v2186 int32
	_ = v2186
	var v2189 int32
	_ = v2189
	var v2192 int32
	_ = v2192
	var v2195 int32
	_ = v2195
	var v2198 int32
	_ = v2198
	var v2199 int32
	_ = v2199
	var v2200 int32
	_ = v2200
	var v2201 int32
	_ = v2201
	var v2202 int32
	_ = v2202
	var v2205 int32
	_ = v2205
	var v2208 int32
	_ = v2208
	var v2209 int32
	_ = v2209
	var v2210 int32
	_ = v2210
	var v2211 int32
	_ = v2211
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
	var v2221 int32
	_ = v2221
	var v2224 int32
	_ = v2224
	var v2225 int32
	_ = v2225
	var v2226 int32
	_ = v2226
	var v2227 int32
	_ = v2227
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
	var v2236 int32
	_ = v2236
	var v2239 int32
	_ = v2239
	var v2242 int32
	_ = v2242
	var v2243 int32
	_ = v2243
	var v2244 int32
	_ = v2244
	var v2245 int32
	_ = v2245
	var v2246 int32
	_ = v2246
	var v2247 int32
	_ = v2247
	var v2248 int32
	_ = v2248
	var v2251 int32
	_ = v2251
	var v2252 int32
	_ = v2252
	var v2254 int32
	_ = v2254
	var v2255 int32
	_ = v2255
	var v2258 int32
	_ = v2258
	var v2259 int32
	_ = v2259
	var v2260 int32
	_ = v2260
	var v2262 int32
	_ = v2262
	var v2263 int32
	_ = v2263
	var v2265 int32
	_ = v2265
	var v2276 int32
	_ = v2276
	var v2277 int32
	_ = v2277
	var v2279 int32
	_ = v2279
	var v2283 int32
	_ = v2283
	var v2303 int32
	_ = v2303
	var v2304 int32
	_ = v2304
	var v2306 int32
	_ = v2306
	var v2310 int32
	_ = v2310
	var v2321 int32
	_ = v2321
	var v2326 int32
	_ = v2326
	var v2327 int32
	_ = v2327
	var v2329 int32
	_ = v2329
	var v2332 int32
	_ = v2332
	var v2337 int32
	_ = v2337
	var v2338 int32
	_ = v2338
	var v2340 int32
	_ = v2340
	var v2343 int32
	_ = v2343
	var v2348 int32
	_ = v2348
	var v2349 int32
	_ = v2349
	var v2357 int32
	_ = v2357
	var v2358 int32
	_ = v2358
	var v2360 int32
	_ = v2360
	var v2364 int32
	_ = v2364
	var v2375 int32
	_ = v2375
	var v2376 int32
	_ = v2376
	var v2378 int32
	_ = v2378
	var v2381 int32
	_ = v2381
	var v2382 int32
	_ = v2382
	var v2384 int32
	_ = v2384
	var v2387 int32
	_ = v2387
	var v2388 int32
	_ = v2388
	var v2414 int32
	_ = v2414
	var v2416 int32
	_ = v2416
	var v2443 int32
	_ = v2443
	var v2444 int32
	_ = v2444
	var v2473 int32
	_ = v2473
	var v2474 int32
	_ = v2474
	var v2484 int32
	_ = v2484
	var v2502 int32
	_ = v2502
	var v2506 int32
	_ = v2506
	var v2507 int32
	_ = v2507
	var v2510 int32
	_ = v2510
	var v2511 int32
	_ = v2511
	var v2522 int32
	_ = v2522
	var v2528 int32
	_ = v2528
	var v2539 int32
	_ = v2539
	var v2543 int32
	_ = v2543
	var v2544 int32
	_ = v2544
	var v2547 int32
	_ = v2547
	var v2558 int32
	_ = v2558
	var v2567 int32
	_ = v2567
	var v2570 int32
	_ = v2570
	var v2573 int32
	_ = v2573
	var v2577 int32
	_ = v2577
	var v2581 int32
	_ = v2581
	var v2584 int32
	_ = v2584
	var v2610 int32
	_ = v2610
	var v2640 int32
	_ = v2640
	var v2641 int32
	_ = v2641
	var v2643 int32
	_ = v2643
	var v2644 int32
	_ = v2644
	var v2645 int32
	_ = v2645
	var v2652 int32
	_ = v2652
	var v2659 int32
	_ = v2659
	var v2660 int32
	_ = v2660
	var v2668 int32
	_ = v2668
	var v2674 int32
	_ = v2674
	var v2675 int32
	_ = v2675
	var v2677 int32
	_ = v2677
	var v2678 int32
	_ = v2678
	var v2685 int32
	_ = v2685
	var v2691 int32
	_ = v2691
	var v2692 int32
	_ = v2692
	var v2694 int32
	_ = v2694
	var v2697 int32
	_ = v2697
	var v2698 int32
	_ = v2698
	var v2703 int32
	_ = v2703
	var v2711 int32
	_ = v2711
	var v2713 int32
	_ = v2713
	var v2715 int32
	_ = v2715
	var v2716 int32
	_ = v2716
	var v2719 int32
	_ = v2719
	var v2724 int32
	_ = v2724
	var v2730 int32
	_ = v2730
	var v2733 int32
	_ = v2733
	var v2734 int32
	_ = v2734
	var v2740 int32
	_ = v2740
	var v2752 int32
	_ = v2752
	var v2762 int32
	_ = v2762
	var v2766 int32
	_ = v2766
	var v2767 int32
	_ = v2767
	var v2768 int32
	_ = v2768
	var v2770 int32
	_ = v2770
	var v2771 int32
	_ = v2771
	var v2780 int32
	_ = v2780
	var v2781 int32
	_ = v2781
	var v2783 int32
	_ = v2783
	var v2786 int32
	_ = v2786
	var v2787 int32
	_ = v2787
	var v2792 int32
	_ = v2792
	var v2799 int32
	_ = v2799
	var v2801 int32
	_ = v2801
	var v2803 int32
	_ = v2803
	var v2806 int32
	_ = v2806
	var v2808 int32
	_ = v2808
	var v2810 int32
	_ = v2810
	var v2815 int32
	_ = v2815
	var v2824 int32
	_ = v2824
	var v2834 int32
	_ = v2834
	var v2835 int32
	_ = v2835
	var v2840 int32
	_ = v2840
	var v2856 int32
	_ = v2856
	var v2857 float64
	_ = v2857
	var v2858 float64
	_ = v2858
	var v2862 float64
	_ = v2862
	var v2863 float64
	_ = v2863
	var v2871 int32
	_ = v2871
	var v2877 int32
	_ = v2877
	var v2880 int32
	_ = v2880
	var v2881 int32
	_ = v2881
	var v2883 int32
	_ = v2883
	var v2884 int32
	_ = v2884
	var v2891 int32
	_ = v2891
	var v2897 int32
	_ = v2897
	var v2898 int32
	_ = v2898
	var v2900 int32
	_ = v2900
	var v2903 int32
	_ = v2903
	var v2904 int32
	_ = v2904
	var v2909 int32
	_ = v2909
	var v2917 int32
	_ = v2917
	var v2919 int32
	_ = v2919
	var v2921 int32
	_ = v2921
	var v2922 int32
	_ = v2922
	var v2925 int32
	_ = v2925
	var v2930 int32
	_ = v2930
	var v2935 int32
	_ = v2935
	var v2936 int32
	_ = v2936
	var v2937 int32
	_ = v2937
	var v2947 int32
	_ = v2947
	var v2948 int32
	_ = v2948
	var v2953 int32
	_ = v2953
	var v2969 int32
	_ = v2969
	var v2970 float64
	_ = v2970
	var v2971 float64
	_ = v2971
	var v2975 float64
	_ = v2975
	var v2976 float64
	_ = v2976
	var v2984 int32
	_ = v2984
	var v2990 int32
	_ = v2990
	var v2993 int32
	_ = v2993
	var v2994 int32
	_ = v2994
	var v2997 int32
	_ = v2997
	var v2998 int32
	_ = v2998
	var v3005 int32
	_ = v3005
	var v3027 int32
	_ = v3027
	var v3030 int32
	_ = v3030
	var v3033 int32
	_ = v3033
	var v3034 int32
	_ = v3034
	var v3035 int32
	_ = v3035
	var v3036 int32
	_ = v3036
	var v3037 int32
	_ = v3037
	var v3038 int32
	_ = v3038
	var v3039 int32
	_ = v3039
	var v3040 int32
	_ = v3040
	var v3041 int32
	_ = v3041
	var v3042 int32
	_ = v3042
	var v3044 int32
	_ = v3044
	var v3045 int32
	_ = v3045
	var v3055 int32
	_ = v3055
	var v3072 int32
	_ = v3072
	var v3077 int32
	_ = v3077
	var v3078 int32
	_ = v3078
	var v3080 int32
	_ = v3080
	var v3107 int32
	_ = v3107
	var v3108 int32
	_ = v3108
	var v3137 int32
	_ = v3137
	var v3140 int32
	_ = v3140
	var v3141 int32
	_ = v3141
	var v3142 int32
	_ = v3142
	var v3145 int32
	_ = v3145
	var v3152 int32
	_ = v3152
	var v3174 int32
	_ = v3174
	var v3178 int32
	_ = v3178
	var v3179 int32
	_ = v3179
	var v3186 int32
	_ = v3186
	var v3187 int32
	_ = v3187
	var v3188 int32
	_ = v3188
	var v3190 int32
	_ = v3190
	var v3192 int32
	_ = v3192
	var v3193 int32
	_ = v3193
	var v3195 int32
	_ = v3195
	var v3197 int32
	_ = v3197
	var v3198 int32
	_ = v3198
	v4 = int32(0)
	v26 = m.G0
	v28 = v26 - int32(32)
	m.G0 = v28
	*(*int32)(unsafe.Add(mBase, uint32(v28)+24)) = v4
	*(*int32)(unsafe.Add(mBase, uint32(v28)+20)) = v4
	v34 = int32(1)
	v36 = int32(*(*uint8)(unsafe.Add(mBase, _consts[383])))
	if v36 == v34 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	v40 = v39
	goto L3
L2:
	;
	v40 = v4
	goto L3
L3:
	;
	if l2 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v863 = float64(-1)
	if v856&int32(1) == int32(0) {
		v1080 = v863
		goto L212
	} else {
		goto L213
	}
L5:
	;
	v828 = int32(0)
	v834 = F_create_append_path(m, v803, v804, v813, v828, v828, v828, v828, v828, float64(-1))
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L30
	} else {
		goto L210
	}
L6:
	;
	v788 = int32(0)
	v794 = F_create_append_path(m, l0, l1, v779, v788, v788, v788, v788, v788, float64(-1))
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L30
	} else {
		goto L207
	}
L7:
	;
	v770 = v34
	v771 = v4
	v773 = v4
	v777 = v40
	v778 = v4
	v779 = v4
	v780 = v4
	v781 = int32(1)
	goto L6
L8:
	;
	goto L9
L9:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v44 <= int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if v744&int32(1) != 0 {
		v770 = v738
		v771 = v739
		v773 = v741
		v777 = v745
		v778 = v746
		v779 = v747
		v780 = v748
		v781 = v749
		goto L6
	} else {
		goto L205
	}
L11:
	;
	v47 = int32(1)
	v738 = v34
	v739 = v4
	v741 = v4
	v744 = v47
	v745 = v40
	v746 = v4
	v747 = v4
	v748 = v4
	v749 = v47
	goto L10
L12:
	;
	goto L13
L13:
	;
	v49 = int32(1)
	v58 = v34
	v59 = v4
	v60 = v4
	v61 = v4
	v64 = v49
	v65 = v40
	v66 = v4
	v67 = v4
	v68 = v4
	v69 = v49
	goto L14
L14:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v76+v60<<(uint(int32(2))%32))))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+32))
	if v81 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v738 = v191
	v739 = v217
	v741 = v192
	v744 = v104
	v745 = v358
	v746 = v717
	v747 = v105
	v748 = v719
	v749 = v218
	goto L10
L16:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)))
	if v106 != int32(1) {
		goto L35
	} else {
		goto L36
	}
L17:
	;
	v104 = int32(0)
	v105 = v67
	goto L16
L18:
	;
	goto L19
L19:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v80)+48))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+16))
	if v86 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v104 = int32(0)
	v105 = v67
	goto L16
L21:
	;
	goto L22
L22:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	switch v88 - int32(290) {
	case 0:
		goto L25
	case 1:
		goto L24
	default:
		goto L23
	}
L23:
	;
	v101 = F_lappend(m, v67, v85)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L30
	} else {
		goto L33
	}
L24:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v85)+72))
	v99 = F_list_concat(m, v67, v98)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L30
	} else {
		goto L32
	}
L25:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+20)))
	if v91 == int32(1) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v85)+76))
	if v94 != 0 {
		goto L23
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v85)+72))
	v96 = F_list_concat(m, v67, v95)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	goto L28
L30:
	;
	return
L31:
	;
	v104 = v64
	v105 = v96
	goto L16
L32:
	;
	v104 = v64
	v105 = v99
	goto L16
L33:
	;
	v104 = v64
	v105 = v101
	goto L16
L34:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v80)+40))
	if v194 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L35:
	;
	v191 = int32(0)
	v192 = v61
	goto L34
L36:
	;
	goto L37
L37:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v80)+44))
	if v110 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v191 = int32(0)
	v192 = v61
	goto L34
L39:
	;
	goto L40
L40:
	;
	v114 = *(*float64)(unsafe.Add(mBase, uint32(l0)+296))
	if base.F64_gt(v114, float64(0)) != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v80)+48))
	if base.F64_le(v114, float64(0)) != 0 {
		v169 = v122
		goto L45
	} else {
		goto L46
	}
L42:
	;
	v174 = v110
	goto L43
L43:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
	switch v175 - int32(290) {
	case 0:
		goto L63
	case 1:
		goto L62
	default:
		goto L61
	}
L44:
	;
	v174 = v169
	goto L43
L45:
	;
	goto L44
L46:
	;
	if base.F64_ge(v114, float64(1)) == int32(0) {
		v135 = v114
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v80)+32))
	if v137 == int32(0) {
		v169 = v122
		goto L45
	} else {
		goto L50
	}
L48:
	;
	v129 = *(*float64)(unsafe.Add(mBase, uint32(v122)+32))
	if base.F64_gt(v129, float64(0)) == int32(0) {
		v135 = v114
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v135 = base.F64_div(v114, v129)
	goto L47
L50:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v137)+4))
	if v140 <= int32(0) {
		v169 = v122
		goto L45
	} else {
		goto L51
	}
L51:
	;
	v145 = v122
	v147 = int32(0)
	goto L52
L52:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v137)+12))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v150+v147<<(uint(int32(2))%32))))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v154)+16))
	if v155 != 0 {
		v162 = v145
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v169 = v162
	goto L45
L54:
	;
	v164 = v147 + int32(1)
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v137)+4))
	if v164 < v165 {
		v145 = v162
		v147 = v164
		goto L52
	} else {
		goto L60
	}
L55:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v80)+48))
	if v154 == v156 {
		v162 = v145
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v158 = F_compare_fractional_path_costs(m, v145, v154, v135)
	mBase = m.M
	if v158 <= int32(0) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v161 = v145
	goto L59
L58:
	;
	v161 = v154
	goto L59
L59:
	;
	v162 = v161
	goto L54
L60:
	;
	goto L53
L61:
	;
	v188 = F_lappend(m, v61, v174)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L30
	} else {
		goto L70
	}
L62:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v174)+72))
	v186 = F_list_concat(m, v61, v185)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L30
	} else {
		goto L69
	}
L63:
	;
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+20)))
	if v178 == int32(1) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v174)+76))
	if v181 != 0 {
		goto L61
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v174)+72))
	v183 = F_list_concat(m, v61, v182)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L30
	} else {
		goto L68
	}
L67:
	;
	goto L66
L68:
	;
	v191 = v58
	v192 = v183
	goto L34
L69:
	;
	v191 = v58
	v192 = v186
	goto L34
L70:
	;
	v191 = v58
	v192 = v188
	goto L34
L71:
	;
	if v65&int32(1) == int32(0) {
		goto L86
	} else {
		goto L87
	}
L72:
	;
	v197 = int32(0)
	v216 = v197
	v217 = v59
	v218 = v197
	goto L71
L73:
	;
	goto L74
L74:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v194)+12))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v199)))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v200)))
	switch v201 - int32(290) {
	case 0:
		goto L77
	case 1:
		goto L76
	default:
		goto L75
	}
L75:
	;
	v214 = F_lappend(m, v59, v200)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L30
	} else {
		goto L84
	}
L76:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v200)+72))
	v212 = F_list_concat(m, v59, v211)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L30
	} else {
		goto L83
	}
L77:
	;
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200)+20)))
	if v204 == int32(1) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v200)+76))
	if v207 != 0 {
		goto L75
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v200)+72))
	v209 = F_list_concat(m, v59, v208)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L30
	} else {
		goto L82
	}
L81:
	;
	goto L80
L82:
	;
	v216 = v200
	v217 = v209
	v218 = v69
	goto L71
L83:
	;
	v216 = v200
	v217 = v212
	v218 = v69
	goto L71
L84:
	;
	v216 = v200
	v217 = v214
	v218 = v69
	goto L71
L85:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v80)+32))
	if v359 == int32(0) {
		v717 = v66
		v719 = v68
		goto L141
	} else {
		goto L142
	}
L86:
	;
	v358 = int32(0)
	goto L85
L87:
	;
	goto L88
L88:
	;
	v224 = int32(0)
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v80)+32))
	if v225 != 0 {
		goto L91
	} else {
		goto L92
	}
L89:
	;
	if v216|v268 == int32(0) {
		v358 = v224
		goto L85
	} else {
		goto L106
	}
L90:
	;
	goto L89
L91:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v225)+4))
	if v230 <= int32(0) {
		v268 = v224
		goto L90
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v268 = int32(0)
	goto L90
L94:
	;
	v233 = int32(0)
	if v233 < v230 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v236 = v230
	goto L97
L96:
	;
	v236 = v233
	goto L97
L97:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v225)+12))
	v239 = int32(0)
	goto L98
L98:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v237+v239<<(uint(int32(2))%32))))
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v247)+21)))
	if v248 == int32(1) {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	goto L93
L100:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v247)+16))
	if v251 == int32(0) {
		v268 = v247
		goto L90
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	v259 = v239 + int32(1)
	if v259 != v236 {
		v239 = v259
		goto L98
	} else {
		goto L105
	}
L103:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v251)+4))
	if v254 == int32(0) {
		v268 = v247
		goto L90
	} else {
		goto L104
	}
L104:
	;
	goto L102
L105:
	;
	goto L99
L106:
	;
	if v268 != 0 {
		goto L109
	} else {
		goto L110
	}
L107:
	;
	v358 = int32(1)
	goto L85
L108:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v268)))
	switch v328 - int32(290) {
	case 0:
		goto L133
	case 1:
		goto L132
	default:
		goto L131
	}
L109:
	;
	if v216 == int32(0) {
		goto L108
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	v283 = v28 + int32(24)
	v285 = v28 + int32(20)
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
	switch v286 - int32(290) {
	case 0:
		goto L117
	case 1:
		goto L116
	default:
		goto L115
	}
L112:
	;
	v277 = *(*float64)(unsafe.Add(mBase, uint32(v216)+56))
	v278 = *(*float64)(unsafe.Add(mBase, uint32(v268)+56))
	if base.F64_lt(v277, v278) == int32(0) {
		goto L108
	} else {
		goto L113
	}
L113:
	;
	goto L111
L114:
	;
	goto L107
L115:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v283)))
	v323 = F_lappend(m, v322, v216)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L30
	} else {
		goto L130
	}
L116:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v283)))
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v216)+72))
	v318 = F_list_concat(m, v316, v317)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L30
	} else {
		goto L129
	}
L117:
	;
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216)+20)))
	if v289 == int32(1) {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	if v285 == int32(0) {
		goto L115
	} else {
		goto L124
	}
L119:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v216)+76))
	if v292 != 0 {
		goto L118
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v283)))
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v216)+72))
	v296 = F_list_concat(m, v294, v295)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L30
	} else {
		goto L123
	}
L122:
	;
	goto L121
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v283))) = v296
	goto L114
L124:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v283)))
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v216)+72))
	v303 = F_list_copy_tail(m, v302, v292)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L30
	} else {
		goto L125
	}
L125:
	;
	v305 = F_list_concat(m, v301, v303)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L30
	} else {
		goto L126
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v283))) = v305
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v216)+72))
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v216)+76))
	v310 = F_list_copy_head(m, v308, v309)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L30
	} else {
		goto L127
	}
L127:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v285)))
	v313 = F_list_concat(m, v312, v310)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L30
	} else {
		goto L128
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v285))) = v313
	goto L114
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v283))) = v318
	goto L114
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v283))) = v323
	goto L114
L131:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v28)+20))
	v346 = F_lappend(m, v345, v268)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L30
	} else {
		goto L140
	}
L132:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v28)+20))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v268)+72))
	v342 = F_list_concat(m, v340, v341)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L30
	} else {
		goto L139
	}
L133:
	;
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268)+20)))
	if v331 == int32(1) {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v268)+76))
	if v334 != 0 {
		goto L131
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v28)+20))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v268)+72))
	v337 = F_list_concat(m, v335, v336)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L30
	} else {
		goto L138
	}
L137:
	;
	goto L136
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+20)) = v337
	goto L107
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+20)) = v342
	goto L107
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+20)) = v346
	goto L107
L141:
	;
	v728 = v60 + int32(1)
	v729 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v728 < v729 {
		v58 = v191
		v59 = v217
		v60 = v728
		v61 = v192
		v64 = v104
		v65 = v358
		v66 = v717
		v67 = v105
		v68 = v719
		v69 = v218
		goto L14
	} else {
		goto L204
	}
L142:
	;
	v362 = int32(0)
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v359)+4))
	if v363 <= v362 {
		v717 = v66
		v719 = v68
		goto L141
	} else {
		goto L143
	}
L143:
	;
	v377 = v362
	v381 = v66
	v383 = v68
	goto L144
L144:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v359)+12))
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v391+v377<<(uint(int32(2))%32))))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v395)+64))
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v395)+16))
	if v398 != 0 {
		goto L146
	} else {
		goto L147
	}
L145:
	;
	v717 = v544
	v719 = v690
	goto L141
L146:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v398)+4))
	v400 = v399
	goto L148
L147:
	;
	v400 = int32(0)
	goto L148
L148:
	;
	if v396 == int32(0) {
		v544 = v381
		goto L149
	} else {
		goto L150
	}
L149:
	;
	if v400 == int32(0) {
		v690 = v383
		goto L181
	} else {
		goto L182
	}
L150:
	;
	if v381 == int32(0) {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v527 = F_lappend(m, v381, v396)
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L30
	} else {
		goto L180
	}
L152:
	;
	v405 = int32(0)
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v381)+4))
	if v406 <= v405 {
		goto L151
	} else {
		goto L153
	}
L153:
	;
	v412 = v405
	goto L154
L154:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v381)+12))
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v434+v412<<(uint(int32(2))%32))))
	if v438 == v396 {
		goto L157
	} else {
		goto L158
	}
L155:
	;
	goto L151
L156:
	;
	if v495 == int32(0) {
		v544 = v381
		goto L149
	} else {
		goto L178
	}
L157:
	;
	v495 = int32(0)
	goto L156
L158:
	;
	goto L159
L159:
	;
	v446 = int32(0)
	goto L162
L160:
	;
	if v485 != 0 {
		goto L175
	} else {
		goto L176
	}
L161:
	;
	v479 = int32(0)
	v485 = base.B2i32(v459 == v479)
	v487 = base.B2i32(v468 != v479) << (uint(int32(1)) % 32)
	goto L160
L162:
	;
	v449 = int32(0)
	if v438 == v449 {
		v459 = v449
		goto L164
	} else {
		goto L165
	}
L163:
	;
	v495 = int32(3)
	goto L156
L164:
	;
	if v396 != 0 {
		goto L168
	} else {
		goto L169
	}
L165:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v438)+4))
	if v453 <= v446 {
		v459 = int32(0)
		goto L164
	} else {
		goto L166
	}
L166:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v438)+12))
	v459 = v455 + v446<<(uint(int32(2))%32)
	goto L164
L167:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v396)+12))
	v468 = v465 + v446<<(uint(int32(2))%32)
	if v459 == int32(0) {
		goto L161
	} else {
		goto L172
	}
L168:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v396)+4))
	if v446 < v460 {
		goto L167
	} else {
		goto L171
	}
L169:
	;
	goto L170
L170:
	;
	v462 = int32(0)
	v485 = base.B2i32(v459 == v462)
	v487 = v462
	goto L160
L171:
	;
	goto L170
L172:
	;
	if v468 == int32(0) {
		goto L161
	} else {
		goto L173
	}
L173:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v459)))
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v468)))
	if v475 == v476 {
		v446 = v446 + int32(1)
		goto L162
	} else {
		goto L174
	}
L174:
	;
	goto L163
L175:
	;
	v489 = v487
	goto L177
L176:
	;
	v489 = int32(1)
	goto L177
L177:
	;
	v495 = v489
	goto L156
L178:
	;
	v499 = v412 + int32(1)
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v381)+4))
	if v499 < v500 {
		v412 = v499
		goto L154
	} else {
		goto L179
	}
L179:
	;
	goto L155
L180:
	;
	v544 = v527
	goto L149
L181:
	;
	v699 = v377 + int32(1)
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v359)+4))
	if v699 < v700 {
		v377 = v699
		v381 = v544
		v383 = v690
		goto L144
	} else {
		goto L203
	}
L182:
	;
	if v383 == int32(0) {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v671 = F_lappend(m, v383, v400)
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		goto L30
	} else {
		goto L202
	}
L184:
	;
	v558 = int32(0)
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v383)+4))
	if v559 <= v558 {
		goto L183
	} else {
		goto L185
	}
L185:
	;
	v565 = v558
	goto L186
L186:
	;
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v383)+12))
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v587+v565<<(uint(int32(2))%32))))
	v592 = int32(0)
	v599 = base.B2i32(v591|v400 == v592)
	if v591 == v592 {
		v638 = v599
		goto L189
	} else {
		goto L190
	}
L187:
	;
	goto L183
L188:
	;
	if v638 != 0 {
		v690 = v383
		goto L181
	} else {
		goto L200
	}
L189:
	;
	goto L188
L190:
	;
	if v400 == int32(0) {
		v638 = v599
		goto L189
	} else {
		goto L191
	}
L191:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v591)+4))
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v400)+4))
	if v605 != v606 {
		v638 = int32(0)
		goto L189
	} else {
		goto L192
	}
L192:
	;
	v608 = int32(1)
	if v605 <= v608 {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v611 = v608
	goto L195
L194:
	;
	v611 = v605
	goto L195
L195:
	;
	v612 = int32(8)
	v617 = int32(0)
	goto L196
L196:
	;
	v625 = v617 << (uint(int32(2)) % 32)
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v591+v612+v625)))
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v625+(v400+v612))))
	v630 = base.B2i32(v627 == v629)
	if v629 != v627 {
		v638 = v630
		goto L189
	} else {
		goto L198
	}
L197:
	;
	v638 = v630
	goto L189
L198:
	;
	v633 = v617 + int32(1)
	if v633 != v611 {
		v617 = v633
		goto L196
	} else {
		goto L199
	}
L199:
	;
	goto L197
L200:
	;
	v643 = v565 + int32(1)
	v644 = *(*int32)(unsafe.Add(mBase, uint32(v383)+4))
	if v643 < v644 {
		v565 = v643
		goto L186
	} else {
		goto L201
	}
L201:
	;
	goto L187
L202:
	;
	v690 = v671
	goto L181
L203:
	;
	goto L145
L204:
	;
	goto L15
L205:
	;
	v758 = int32(0)
	if v738&int32(1) == v758 {
		v838 = l0
		v839 = l1
		v840 = l2
		v846 = v739
		v850 = v28
		v852 = v745
		v853 = v746
		v854 = v758
		v855 = v748
		v856 = v749
		v857 = v4
		v858 = v4
		v859 = v4
		goto L4
	} else {
		goto L206
	}
L206:
	;
	v803 = l0
	v804 = l1
	v805 = l2
	v811 = v739
	v813 = v741
	v815 = v28
	v817 = v745
	v818 = v746
	v819 = v758
	v820 = v748
	v821 = v749
	v822 = v4
	v823 = v4
	v824 = v4
	goto L5
L207:
	;
	F_add_path(m, l1, v794)
	mBase = m.M
	v797 = m.ExcPending
	if v797 != 0 {
		goto L30
	} else {
		goto L208
	}
L208:
	;
	v798 = int32(1)
	if v770&v798 == int32(0) {
		v838 = l0
		v839 = l1
		v840 = l2
		v846 = v771
		v850 = v28
		v852 = v777
		v853 = v778
		v854 = v798
		v855 = v780
		v856 = v781
		v857 = v4
		v858 = v4
		v859 = v4
		goto L4
	} else {
		goto L209
	}
L209:
	;
	v803 = l0
	v804 = l1
	v805 = l2
	v811 = v771
	v813 = v773
	v815 = v28
	v817 = v777
	v818 = v778
	v819 = v798
	v820 = v780
	v821 = v781
	v822 = v4
	v823 = v4
	v824 = v4
	goto L5
L210:
	;
	F_add_path(m, v804, v834)
	mBase = m.M
	v837 = m.ExcPending
	if v837 != 0 {
		goto L30
	} else {
		goto L211
	}
L211:
	;
	v838 = v803
	v839 = v804
	v840 = v805
	v846 = v811
	v850 = v815
	v852 = v817
	v853 = v818
	v854 = v819
	v855 = v820
	v856 = v821
	v857 = v822
	v858 = v823
	v859 = v824
	goto L4
L212:
	;
	v1081 = *(*int32)(unsafe.Add(mBase, uint32(v850)+20))
	if base.B2i32(v1081 != int32(0))&v852 != 0 {
		goto L257
	} else {
		goto L258
	}
L213:
	;
	if v846 == int32(0) {
		v1080 = v863
		goto L212
	} else {
		goto L214
	}
L214:
	;
	v870 = int32(0)
	v871 = *(*int32)(unsafe.Add(mBase, uint32(v846)+4))
	if v871 <= v870 {
		v1000 = v870
		goto L215
	} else {
		goto L216
	}
L215:
	;
	v1022 = int32(0)
	v1026 = int32(*(*uint8)(unsafe.Add(mBase, _consts[383])))
	if v1026 != 0 {
		goto L242
	} else {
		goto L243
	}
L216:
	;
	v875 = v871 & int32(3)
	v876 = *(*int32)(unsafe.Add(mBase, uint32(v846)+12))
	v877 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v871) {
		goto L217
	} else {
		goto L218
	}
L217:
	;
	v887 = v870
	v888 = v877
	v893 = int32(0)
	goto L220
L218:
	;
	v936 = v870
	v937 = v877
	goto L219
L219:
	;
	if v875 == int32(0) {
		v1000 = v936
		goto L215
	} else {
		goto L235
	}
L220:
	;
	v911 = v876 + v888<<(uint(int32(2))%32)
	v912 = *(*int32)(unsafe.Add(mBase, uint32(v911)))
	v913 = *(*int32)(unsafe.Add(mBase, uint32(v912)+24))
	if v913 < v887 {
		goto L222
	} else {
		goto L223
	}
L221:
	;
	v936 = v927
	v937 = v929
	goto L219
L222:
	;
	v915 = v887
	goto L224
L223:
	;
	v915 = v913
	goto L224
L224:
	;
	v916 = *(*int32)(unsafe.Add(mBase, uint32(v911)+4))
	v917 = *(*int32)(unsafe.Add(mBase, uint32(v916)+24))
	if v917 < v915 {
		goto L225
	} else {
		goto L226
	}
L225:
	;
	v919 = v915
	goto L227
L226:
	;
	v919 = v917
	goto L227
L227:
	;
	v920 = *(*int32)(unsafe.Add(mBase, uint32(v911)+8))
	v921 = *(*int32)(unsafe.Add(mBase, uint32(v920)+24))
	if v921 < v919 {
		goto L228
	} else {
		goto L229
	}
L228:
	;
	v923 = v919
	goto L230
L229:
	;
	v923 = v921
	goto L230
L230:
	;
	v924 = *(*int32)(unsafe.Add(mBase, uint32(v911)+12))
	v925 = *(*int32)(unsafe.Add(mBase, uint32(v924)+24))
	if v925 < v923 {
		goto L231
	} else {
		goto L232
	}
L231:
	;
	v927 = v923
	goto L233
L232:
	;
	v927 = v925
	goto L233
L233:
	;
	v928 = int32(4)
	v929 = v888 + v928
	v931 = v893 + v928
	if v931 != v871&int32(2147483644) {
		v887 = v927
		v888 = v929
		v893 = v931
		goto L220
	} else {
		goto L234
	}
L234:
	;
	goto L221
L235:
	;
	v963 = v936
	v964 = v937
	v971 = v877
	goto L236
L236:
	;
	v988 = *(*int32)(unsafe.Add(mBase, uint32(v876+v964<<(uint(int32(2))%32))))
	v989 = *(*int32)(unsafe.Add(mBase, uint32(v988)+24))
	if v989 < v963 {
		goto L238
	} else {
		goto L239
	}
L237:
	;
	v1000 = v991
	goto L215
L238:
	;
	v991 = v963
	goto L240
L239:
	;
	v991 = v989
	goto L240
L240:
	;
	v992 = int32(1)
	v995 = v971 + v992
	if v995 != v875 {
		v963 = v991
		v964 = v964 + v992
		v971 = v995
		goto L236
	} else {
		goto L241
	}
L241:
	;
	goto L237
L242:
	;
	if v840 == int32(0) {
		goto L247
	} else {
		goto L248
	}
L243:
	;
	v1049 = v1000
	goto L244
L244:
	;
	v1051 = F_create_append_path(m, v838, v839, v1022, v846, v1022, v1022, v1049, v1026, float64(-1))
	mBase = m.M
	v1052 = m.ExcPending
	if v1052 != 0 {
		goto L30
	} else {
		goto L255
	}
L245:
	;
	v1044 = *(*int32)(unsafe.Add(mBase, _consts[384]))
	if v1041 < v1044 {
		goto L252
	} else {
		goto L253
	}
L246:
	;
	v1041 = int32(32) - base.I32_clz(v1037)
	goto L245
L247:
	;
	v1029 = int32(0)
	if v1000 <= v1029 {
		v1037 = v1029
		goto L246
	} else {
		goto L250
	}
L248:
	;
	goto L249
L249:
	;
	v1033 = *(*int32)(unsafe.Add(mBase, uint32(v840)+4))
	if int32(32)-base.I32_clz(v1033) < v1000 {
		v1041 = v1000
		goto L245
	} else {
		goto L251
	}
L250:
	;
	v1041 = v1000
	goto L245
L251:
	;
	v1037 = v1033
	goto L246
L252:
	;
	v1046 = v1041
	goto L254
L253:
	;
	v1046 = v1044
	goto L254
L254:
	;
	v1049 = v1046
	goto L244
L255:
	;
	v1053 = *(*float64)(unsafe.Add(mBase, uint32(v1051)+32))
	F_add_partial_path(m, v839, v1051)
	mBase = m.M
	v1055 = m.ExcPending
	if v1055 != 0 {
		goto L30
	} else {
		goto L256
	}
L256:
	;
	v1080 = v1053
	goto L212
L257:
	;
	v1085 = *(*int32)(unsafe.Add(mBase, uint32(v850)+24))
	if v1085 != 0 {
		goto L266
	} else {
		goto L267
	}
L258:
	;
	goto L259
L259:
	;
	if v854 == int32(0) {
		goto L307
	} else {
		goto L308
	}
L260:
	;
	v1311 = int32(0)
	v1314 = *(*int32)(unsafe.Add(mBase, _consts[384]))
	if v1289 < v1314 {
		goto L302
	} else {
		goto L303
	}
L261:
	;
	v1289 = int32(32) - base.I32_clz(v1262)
	goto L260
L262:
	;
	if v840 == int32(0) {
		goto L297
	} else {
		goto L298
	}
L263:
	;
	if v1095 == int32(0) {
		v1226 = v1160
		goto L262
	} else {
		goto L290
	}
L264:
	;
	v1156 = *(*int32)(unsafe.Add(mBase, uint32(v840)+4))
	v1262 = v1156
	goto L261
L265:
	;
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(v1085)+12))
	v1104 = int32(0)
	v1110 = v1104
	v1111 = v1104
	v1118 = v1104
	goto L275
L266:
	;
	v1086 = int32(0)
	v1087 = *(*int32)(unsafe.Add(mBase, uint32(v1085)+4))
	if v1087 <= v1086 {
		v1226 = v1086
		goto L262
	} else {
		goto L269
	}
L267:
	;
	goto L268
L268:
	;
	if v840 != 0 {
		goto L264
	} else {
		goto L274
	}
L269:
	;
	v1090 = int32(0)
	if v1090 < v1087 {
		goto L270
	} else {
		goto L271
	}
L270:
	;
	v1093 = v1087
	goto L272
L271:
	;
	v1093 = v1090
	goto L272
L272:
	;
	v1095 = v1093 & int32(3)
	if int32(4) <= v1087 {
		goto L265
	} else {
		goto L273
	}
L273:
	;
	v1098 = int32(0)
	v1160 = v1098
	v1161 = v1098
	goto L263
L274:
	;
	v1262 = int32(0)
	goto L261
L275:
	;
	v1134 = v1103 + v1111<<(uint(int32(2))%32)
	v1135 = *(*int32)(unsafe.Add(mBase, uint32(v1134)))
	v1136 = *(*int32)(unsafe.Add(mBase, uint32(v1135)+24))
	if v1136 < v1110 {
		goto L277
	} else {
		goto L278
	}
L276:
	;
	v1160 = v1150
	v1161 = v1152
	goto L263
L277:
	;
	v1138 = v1110
	goto L279
L278:
	;
	v1138 = v1136
	goto L279
L279:
	;
	v1139 = *(*int32)(unsafe.Add(mBase, uint32(v1134)+4))
	v1140 = *(*int32)(unsafe.Add(mBase, uint32(v1139)+24))
	if v1140 < v1138 {
		goto L280
	} else {
		goto L281
	}
L280:
	;
	v1142 = v1138
	goto L282
L281:
	;
	v1142 = v1140
	goto L282
L282:
	;
	v1143 = *(*int32)(unsafe.Add(mBase, uint32(v1134)+8))
	v1144 = *(*int32)(unsafe.Add(mBase, uint32(v1143)+24))
	if v1144 < v1142 {
		goto L283
	} else {
		goto L284
	}
L283:
	;
	v1146 = v1142
	goto L285
L284:
	;
	v1146 = v1144
	goto L285
L285:
	;
	v1147 = *(*int32)(unsafe.Add(mBase, uint32(v1134)+12))
	v1148 = *(*int32)(unsafe.Add(mBase, uint32(v1147)+24))
	if v1148 < v1146 {
		goto L286
	} else {
		goto L287
	}
L286:
	;
	v1150 = v1146
	goto L288
L287:
	;
	v1150 = v1148
	goto L288
L288:
	;
	v1151 = int32(4)
	v1152 = v1111 + v1151
	v1154 = v1118 + v1151
	if v1154 != v1093&int32(2147483644) {
		v1110 = v1150
		v1111 = v1152
		v1118 = v1154
		goto L275
	} else {
		goto L289
	}
L289:
	;
	goto L276
L290:
	;
	v1184 = *(*int32)(unsafe.Add(mBase, uint32(v1085)+12))
	v1189 = v1160
	v1190 = v1161
	v1191 = int32(0)
	goto L291
L291:
	;
	v1214 = *(*int32)(unsafe.Add(mBase, uint32(v1184+v1190<<(uint(int32(2))%32))))
	v1215 = *(*int32)(unsafe.Add(mBase, uint32(v1214)+24))
	if v1215 < v1189 {
		goto L293
	} else {
		goto L294
	}
L292:
	;
	v1226 = v1217
	goto L262
L293:
	;
	v1217 = v1189
	goto L295
L294:
	;
	v1217 = v1215
	goto L295
L295:
	;
	v1218 = int32(1)
	v1221 = v1191 + v1218
	if v1221 != v1095 {
		v1189 = v1217
		v1190 = v1190 + v1218
		v1191 = v1221
		goto L291
	} else {
		goto L296
	}
L296:
	;
	goto L292
L297:
	;
	v1250 = int32(0)
	if v1226 <= v1250 {
		v1262 = v1250
		goto L261
	} else {
		goto L300
	}
L298:
	;
	goto L299
L299:
	;
	v1254 = *(*int32)(unsafe.Add(mBase, uint32(v840)+4))
	if int32(32)-base.I32_clz(v1254) < v1226 {
		v1289 = v1226
		goto L260
	} else {
		goto L301
	}
L300:
	;
	v1289 = v1226
	goto L260
L301:
	;
	v1262 = v1254
	goto L261
L302:
	;
	v1316 = v1289
	goto L304
L303:
	;
	v1316 = v1314
	goto L304
L304:
	;
	v1318 = F_create_append_path(m, v838, v839, v1081, v1085, v1311, v1311, v1316, int32(1), v1080)
	mBase = m.M
	v1319 = m.ExcPending
	if v1319 != 0 {
		goto L30
	} else {
		goto L305
	}
L305:
	;
	F_add_partial_path(m, v839, v1318)
	mBase = m.M
	v1321 = m.ExcPending
	if v1321 != 0 {
		goto L30
	} else {
		goto L306
	}
L306:
	;
	goto L259
L307:
	;
	if v855 == int32(0) {
		goto L664
	} else {
		goto L665
	}
L308:
	;
	v1349 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v850)+31)) = uint8(v1349)
	*(*uint8)(unsafe.Add(mBase, uint32(v850)+30)) = uint8(v1349)
	v1353 = *(*int32)(unsafe.Add(mBase, uint32(v839)+232))
	if v1353 == int32(0) {
		v1437 = v857
		v1438 = v858
		goto L309
	} else {
		goto L310
	}
L309:
	;
	if v853 == int32(0) {
		goto L307
	} else {
		goto L337
	}
L310:
	;
	v1356 = *(*int32)(unsafe.Add(mBase, uint32(v839)+4))
	switch v1356 {
	case 0, 2:
		goto L311
	default:
		v1437 = v857
		v1438 = v858
		goto L309
	}
L311:
	;
	v1357 = *(*int32)(unsafe.Add(mBase, uint32(v839)+256))
	v1358 = *(*int32)(unsafe.Add(mBase, uint32(v839)+240))
	v1359 = *(*int32)(unsafe.Add(mBase, uint32(v1358)))
	switch v1359 - int32(108) {
	case 0:
		goto L314
	default:
		goto L313
	case 6:
		goto L315
	}
L312:
	;
	if v1421 == int32(0) {
		v1437 = v857
		v1438 = v858
		goto L309
	} else {
		goto L334
	}
L313:
	;
	v1421 = int32(0)
	goto L312
L314:
	;
	v1369 = *(*int32)(unsafe.Add(mBase, uint32(v1358)+16))
	v1370 = int32(0)
	if v1357 == v1370 {
		v1411 = v1370
		goto L320
	} else {
		goto L321
	}
L315:
	;
	v1362 = int32(1)
	v1363 = *(*int32)(unsafe.Add(mBase, uint32(v1358)+32))
	if v1363 == int32(-1) {
		v1421 = v1362
		goto L312
	} else {
		goto L316
	}
L316:
	;
	v1366 = F_bms_is_member(m, v1363, v1357)
	mBase = m.M
	v1367 = m.ExcPending
	if v1367 != 0 {
		goto L30
	} else {
		goto L317
	}
L317:
	;
	if v1366 != 0 {
		goto L313
	} else {
		goto L318
	}
L318:
	;
	v1421 = v1362
	goto L312
L319:
	;
	if v1411 == int32(0) {
		v1421 = int32(1)
		goto L312
	} else {
		goto L333
	}
L320:
	;
	goto L319
L321:
	;
	if v1369 == int32(0) {
		v1411 = v1370
		goto L320
	} else {
		goto L322
	}
L322:
	;
	v1379 = *(*int32)(unsafe.Add(mBase, uint32(v1357)+4))
	v1380 = *(*int32)(unsafe.Add(mBase, uint32(v1369)+4))
	if v1379 < v1380 {
		goto L323
	} else {
		goto L324
	}
L323:
	;
	v1382 = v1379
	goto L325
L324:
	;
	v1382 = v1380
	goto L325
L325:
	;
	if v1382 <= int32(1) {
		goto L326
	} else {
		goto L327
	}
L326:
	;
	v1385 = int32(1)
	goto L328
L327:
	;
	v1385 = v1382
	goto L328
L328:
	;
	v1386 = int32(8)
	v1391 = int32(0)
	goto L329
L329:
	;
	v1398 = v1391 << (uint(int32(2)) % 32)
	v1400 = *(*int32)(unsafe.Add(mBase, uint32(v1369+v1386+v1398)))
	v1402 = *(*int32)(unsafe.Add(mBase, uint32(v1398+(v1357+v1386))))
	v1403 = v1400 & v1402
	v1405 = base.B2i32(v1403 != int32(0))
	if v1403 != 0 {
		v1411 = v1405
		goto L320
	} else {
		goto L331
	}
L330:
	;
	v1411 = v1405
	goto L320
L331:
	;
	v1407 = v1391 + int32(1)
	if v1407 != v1385 {
		v1391 = v1407
		goto L329
	} else {
		goto L332
	}
L332:
	;
	goto L330
L333:
	;
	goto L313
L334:
	;
	v1427 = F_build_partition_pathkeys(m, v838, v839, int32(1), v850+int32(31))
	mBase = m.M
	v1428 = m.ExcPending
	if v1428 != 0 {
		goto L30
	} else {
		goto L335
	}
L335:
	;
	v1432 = F_build_partition_pathkeys(m, v838, v839, int32(-1), v850+int32(30))
	mBase = m.M
	v1433 = m.ExcPending
	if v1433 != 0 {
		goto L30
	} else {
		goto L336
	}
L336:
	;
	v1437 = v1427
	v1438 = v1432
	goto L309
L337:
	;
	v1441 = *(*int32)(unsafe.Add(mBase, uint32(v853)+4))
	if v1441 <= int32(0) {
		goto L307
	} else {
		goto L338
	}
L338:
	;
	v1465 = v859
	goto L339
L339:
	;
	v1469 = int32(1)
	v1470 = *(*int32)(unsafe.Add(mBase, uint32(v853)+12))
	v1474 = *(*int32)(unsafe.Add(mBase, uint32(v1470+v1465<<(uint(int32(2))%32))))
	if v1474 == v1437 {
		goto L351
	} else {
		goto L352
	}
L340:
	;
	goto L307
L341:
	;
	v2443 = v1465 + int32(1)
	v2444 = *(*int32)(unsafe.Add(mBase, uint32(v853)+4))
	if v2443 < v2444 {
		v1465 = v2443
		goto L339
	} else {
		goto L663
	}
L342:
	;
	F_add_path(m, v839, v2414)
	mBase = m.M
	v2416 = m.ExcPending
	if v2416 != 0 {
		goto L30
	} else {
		goto L662
	}
L343:
	;
	v2375 = F_create_merge_append_path(m, v838, v839, v2358, v1474)
	mBase = m.M
	v2376 = m.ExcPending
	if v2376 != 0 {
		goto L30
	} else {
		goto L653
	}
L344:
	;
	v2321 = int32(0)
	v2326 = F_create_append_path(m, v838, v839, v2304, v2321, v1474, v2321, v2321, v2321, float64(-1))
	mBase = m.M
	v2327 = m.ExcPending
	if v2327 != 0 {
		goto L30
	} else {
		goto L644
	}
L345:
	;
	if v1714 == int32(0) {
		v2357 = v2276
		v2358 = v2277
		v2360 = v2279
		v2364 = v2283
		goto L343
	} else {
		goto L643
	}
L346:
	;
	v2265 = int32(0)
	if v1701 != 0 {
		v2303 = v2265
		v2304 = v2265
		v2306 = v2265
		v2310 = v2265
		goto L344
	} else {
		goto L642
	}
L347:
	;
	v1717 = int32(0)
	if v1713 == v1715 {
		v2276 = v1717
		v2277 = v1717
		v2279 = v1717
		v2283 = v1717
		goto L345
	} else {
		goto L434
	}
L348:
	;
	v1707 = int32(-1)
	v1708 = *(*int32)(unsafe.Add(mBase, uint32(v840)+4))
	v1709 = int32(1)
	v1713 = v1708 - v1709
	v1714 = v1709
	v1715 = v1707
	v1716 = v1707
	goto L347
L349:
	;
	if v840 == int32(0) {
		goto L346
	} else {
		goto L433
	}
L350:
	;
	if v1527 != 0 {
		v1701 = v1469
		goto L349
	} else {
		goto L368
	}
L351:
	;
	v1527 = int32(1)
	goto L350
L352:
	;
	goto L353
L353:
	;
	v1483 = int32(0)
	goto L355
L354:
	;
	v1527 = v1519
	goto L350
L355:
	;
	v1487 = int32(0)
	if v1474 == v1487 {
		v1497 = v1487
		goto L357
	} else {
		goto L358
	}
L356:
	;
	v1519 = int32(0)
	goto L354
L357:
	;
	if v1437 != 0 {
		goto L361
	} else {
		goto L362
	}
L358:
	;
	v1491 = *(*int32)(unsafe.Add(mBase, uint32(v1474)+4))
	if v1491 <= v1483 {
		v1497 = int32(0)
		goto L357
	} else {
		goto L359
	}
L359:
	;
	v1493 = *(*int32)(unsafe.Add(mBase, uint32(v1474)+12))
	v1497 = v1493 + v1483<<(uint(int32(2))%32)
	goto L357
L360:
	;
	v1503 = base.B2i32(v1497 == int32(0))
	if v1497 == int32(0) {
		v1519 = v1503
		goto L354
	} else {
		goto L365
	}
L361:
	;
	v1498 = *(*int32)(unsafe.Add(mBase, uint32(v1437)+4))
	if v1483 < v1498 {
		goto L360
	} else {
		goto L364
	}
L362:
	;
	goto L363
L363:
	;
	v1527 = base.B2i32(v1497 == int32(0))
	goto L350
L364:
	;
	goto L363
L365:
	;
	v1506 = *(*int32)(unsafe.Add(mBase, uint32(v1437)+12))
	v1509 = v1506 + v1483<<(uint(int32(2))%32)
	if v1509 == int32(0) {
		v1519 = v1503
		goto L354
	} else {
		goto L366
	}
L366:
	;
	v1514 = *(*int32)(unsafe.Add(mBase, uint32(v1497)))
	v1515 = *(*int32)(unsafe.Add(mBase, uint32(v1509)))
	if v1514 == v1515 {
		v1483 = v1483 + int32(1)
		goto L355
	} else {
		goto L367
	}
L367:
	;
	goto L356
L368:
	;
	v1528 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v850)+31)))
	if v1528 == int32(0) {
		goto L369
	} else {
		goto L370
	}
L369:
	;
	if v1437 == v1474 {
		goto L373
	} else {
		goto L374
	}
L370:
	;
	goto L371
L371:
	;
	if v1474 == v1438 {
		goto L392
	} else {
		goto L393
	}
L372:
	;
	if v1583 != 0 {
		v1701 = v1469
		goto L349
	} else {
		goto L390
	}
L373:
	;
	v1583 = int32(1)
	goto L372
L374:
	;
	goto L375
L375:
	;
	v1539 = int32(0)
	goto L377
L376:
	;
	v1583 = v1575
	goto L372
L377:
	;
	v1543 = int32(0)
	if v1437 == v1543 {
		v1553 = v1543
		goto L379
	} else {
		goto L380
	}
L378:
	;
	v1575 = int32(0)
	goto L376
L379:
	;
	if v1474 != 0 {
		goto L383
	} else {
		goto L384
	}
L380:
	;
	v1547 = *(*int32)(unsafe.Add(mBase, uint32(v1437)+4))
	if v1547 <= v1539 {
		v1553 = int32(0)
		goto L379
	} else {
		goto L381
	}
L381:
	;
	v1549 = *(*int32)(unsafe.Add(mBase, uint32(v1437)+12))
	v1553 = v1549 + v1539<<(uint(int32(2))%32)
	goto L379
L382:
	;
	v1559 = base.B2i32(v1553 == int32(0))
	if v1553 == int32(0) {
		v1575 = v1559
		goto L376
	} else {
		goto L387
	}
L383:
	;
	v1554 = *(*int32)(unsafe.Add(mBase, uint32(v1474)+4))
	if v1539 < v1554 {
		goto L382
	} else {
		goto L386
	}
L384:
	;
	goto L385
L385:
	;
	v1583 = base.B2i32(v1553 == int32(0))
	goto L372
L386:
	;
	goto L385
L387:
	;
	v1562 = *(*int32)(unsafe.Add(mBase, uint32(v1474)+12))
	v1565 = v1562 + v1539<<(uint(int32(2))%32)
	if v1565 == int32(0) {
		v1575 = v1559
		goto L376
	} else {
		goto L388
	}
L388:
	;
	v1570 = *(*int32)(unsafe.Add(mBase, uint32(v1553)))
	v1571 = *(*int32)(unsafe.Add(mBase, uint32(v1565)))
	if v1570 == v1571 {
		v1539 = v1539 + int32(1)
		goto L377
	} else {
		goto L389
	}
L389:
	;
	goto L378
L390:
	;
	goto L371
L391:
	;
	if v1636 == int32(0) {
		goto L409
	} else {
		goto L410
	}
L392:
	;
	v1636 = int32(1)
	goto L391
L393:
	;
	goto L394
L394:
	;
	v1592 = int32(0)
	goto L396
L395:
	;
	v1636 = v1628
	goto L391
L396:
	;
	v1596 = int32(0)
	if v1474 == v1596 {
		v1606 = v1596
		goto L398
	} else {
		goto L399
	}
L397:
	;
	v1628 = int32(0)
	goto L395
L398:
	;
	if v1438 != 0 {
		goto L402
	} else {
		goto L403
	}
L399:
	;
	v1600 = *(*int32)(unsafe.Add(mBase, uint32(v1474)+4))
	if v1600 <= v1592 {
		v1606 = int32(0)
		goto L398
	} else {
		goto L400
	}
L400:
	;
	v1602 = *(*int32)(unsafe.Add(mBase, uint32(v1474)+12))
	v1606 = v1602 + v1592<<(uint(int32(2))%32)
	goto L398
L401:
	;
	v1612 = base.B2i32(v1606 == int32(0))
	if v1606 == int32(0) {
		v1628 = v1612
		goto L395
	} else {
		goto L406
	}
L402:
	;
	v1607 = *(*int32)(unsafe.Add(mBase, uint32(v1438)+4))
	if v1592 < v1607 {
		goto L401
	} else {
		goto L405
	}
L403:
	;
	goto L404
L404:
	;
	v1636 = base.B2i32(v1606 == int32(0))
	goto L391
L405:
	;
	goto L404
L406:
	;
	v1615 = *(*int32)(unsafe.Add(mBase, uint32(v1438)+12))
	v1618 = v1615 + v1592<<(uint(int32(2))%32)
	if v1618 == int32(0) {
		v1628 = v1612
		goto L395
	} else {
		goto L407
	}
L407:
	;
	v1623 = *(*int32)(unsafe.Add(mBase, uint32(v1606)))
	v1624 = *(*int32)(unsafe.Add(mBase, uint32(v1618)))
	if v1623 == v1624 {
		v1592 = v1592 + int32(1)
		goto L396
	} else {
		goto L408
	}
L408:
	;
	goto L397
L409:
	;
	v1639 = int32(0)
	v1640 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v850)+30)))
	if v1640 != 0 {
		v1701 = v1639
		goto L349
	} else {
		goto L412
	}
L410:
	;
	goto L411
L411:
	;
	if v840 != 0 {
		goto L348
	} else {
		goto L432
	}
L412:
	;
	if v1438 == v1474 {
		goto L414
	} else {
		goto L415
	}
L413:
	;
	if v1693 == int32(0) {
		v1701 = v1639
		goto L349
	} else {
		goto L431
	}
L414:
	;
	v1693 = int32(1)
	goto L413
L415:
	;
	goto L416
L416:
	;
	v1649 = int32(0)
	goto L418
L417:
	;
	v1693 = v1685
	goto L413
L418:
	;
	v1653 = int32(0)
	if v1438 == v1653 {
		v1663 = v1653
		goto L420
	} else {
		goto L421
	}
L419:
	;
	v1685 = int32(0)
	goto L417
L420:
	;
	if v1474 != 0 {
		goto L424
	} else {
		goto L425
	}
L421:
	;
	v1657 = *(*int32)(unsafe.Add(mBase, uint32(v1438)+4))
	if v1657 <= v1649 {
		v1663 = int32(0)
		goto L420
	} else {
		goto L422
	}
L422:
	;
	v1659 = *(*int32)(unsafe.Add(mBase, uint32(v1438)+12))
	v1663 = v1659 + v1649<<(uint(int32(2))%32)
	goto L420
L423:
	;
	v1669 = base.B2i32(v1663 == int32(0))
	if v1663 == int32(0) {
		v1685 = v1669
		goto L417
	} else {
		goto L428
	}
L424:
	;
	v1664 = *(*int32)(unsafe.Add(mBase, uint32(v1474)+4))
	if v1649 < v1664 {
		goto L423
	} else {
		goto L427
	}
L425:
	;
	goto L426
L426:
	;
	v1693 = base.B2i32(v1663 == int32(0))
	goto L413
L427:
	;
	goto L426
L428:
	;
	v1672 = *(*int32)(unsafe.Add(mBase, uint32(v1474)+12))
	v1675 = v1672 + v1649<<(uint(int32(2))%32)
	if v1675 == int32(0) {
		v1685 = v1669
		goto L417
	} else {
		goto L429
	}
L429:
	;
	v1680 = *(*int32)(unsafe.Add(mBase, uint32(v1663)))
	v1681 = *(*int32)(unsafe.Add(mBase, uint32(v1675)))
	if v1680 == v1681 {
		v1649 = v1649 + int32(1)
		goto L418
	} else {
		goto L430
	}
L430:
	;
	goto L419
L431:
	;
	goto L411
L432:
	;
	v1697 = int32(0)
	v2303 = v1697
	v2304 = v1697
	v2306 = v1697
	v2310 = v1697
	goto L344
L433:
	;
	v1704 = *(*int32)(unsafe.Add(mBase, uint32(v840)+4))
	v1713 = int32(0)
	v1714 = v1701
	v1715 = v1704
	v1716 = int32(1)
	goto L347
L434:
	;
	v1729 = v1717
	v1730 = v1717
	v1732 = v1717
	v1733 = v1713
	v1736 = v1717
	goto L435
L435:
	;
	v1747 = *(*int32)(unsafe.Add(mBase, uint32(v840)+12))
	v1751 = *(*int32)(unsafe.Add(mBase, uint32(v1747+v1733<<(uint(int32(2))%32))))
	v1752 = *(*int32)(unsafe.Add(mBase, uint32(v1751)+32))
	v1753 = int32(0)
	if v1752 != 0 {
		goto L440
	} else {
		goto L441
	}
L436:
	;
	v2276 = v2258
	v2277 = v2259
	v2279 = v2260
	v2283 = v2262
	goto L345
L437:
	;
	v1879 = int32(0)
	v1880 = *(*int32)(unsafe.Add(mBase, uint32(v1751)+32))
	if v1880 != 0 {
		goto L483
	} else {
		goto L484
	}
L438:
	;
	goto L437
L439:
	;
	v1772 = v1753
	v1775 = v1753
	goto L444
L440:
	;
	v1763 = *(*int32)(unsafe.Add(mBase, uint32(v1752)+4))
	if int32(0) < v1763 {
		goto L439
	} else {
		goto L443
	}
L441:
	;
	goto L442
L442:
	;
	v1873 = v1753
	goto L438
L443:
	;
	goto L442
L444:
	;
	v1778 = *(*int32)(unsafe.Add(mBase, uint32(v1752)+12))
	v1782 = *(*int32)(unsafe.Add(mBase, uint32(v1778+v1775<<(uint(int32(2))%32))))
	goto L448
L445:
	;
	v1873 = v1857
	goto L438
L446:
	;
	v1864 = v1775 + int32(1)
	v1865 = *(*int32)(unsafe.Add(mBase, uint32(v1752)+4))
	if v1864 < v1865 {
		v1772 = v1857
		v1775 = v1864
		goto L444
	} else {
		goto L479
	}
L448:
	;
	goto L449
L449:
	;
	if v1772 != 0 {
		goto L451
	} else {
		goto L452
	}
L451:
	;
	v1786 = F_compare_path_costs(m, v1772, v1782, v1753)
	mBase = m.M
	if v1786 <= int32(0) {
		v1857 = v1772
		goto L446
	} else {
		goto L454
	}
L452:
	;
	goto L453
L453:
	;
	v1789 = *(*int32)(unsafe.Add(mBase, uint32(v1782)+64))
	if v1474 == v1789 {
		goto L455
	} else {
		goto L456
	}
L454:
	;
	goto L453
L455:
	;
	v1845 = *(*int32)(unsafe.Add(mBase, uint32(v1782)+16))
	if v1845 != 0 {
		goto L473
	} else {
		goto L474
	}
L456:
	;
	v1797 = int32(0)
	goto L457
L457:
	;
	v1804 = int32(0)
	if v1474 == v1804 {
		v1814 = v1804
		goto L459
	} else {
		goto L460
	}
L458:
	;
	if v1814 != 0 {
		v1857 = v1772
		goto L446
	} else {
		goto L472
	}
L459:
	;
	if v1789 != 0 {
		goto L463
	} else {
		goto L464
	}
L460:
	;
	v1808 = *(*int32)(unsafe.Add(mBase, uint32(v1474)+4))
	if v1808 <= v1797 {
		v1814 = int32(0)
		goto L459
	} else {
		goto L461
	}
L461:
	;
	v1810 = *(*int32)(unsafe.Add(mBase, uint32(v1474)+12))
	v1814 = v1810 + v1797<<(uint(int32(2))%32)
	goto L459
L462:
	;
	if v1814 == int32(0) {
		goto L468
	} else {
		goto L469
	}
L463:
	;
	v1815 = *(*int32)(unsafe.Add(mBase, uint32(v1789)+4))
	if v1797 < v1815 {
		goto L462
	} else {
		goto L466
	}
L464:
	;
	goto L465
L465:
	;
	if v1814 == int32(0) {
		goto L455
	} else {
		goto L467
	}
L466:
	;
	goto L465
L467:
	;
	v1857 = v1772
	goto L446
L468:
	;
	goto L458
L469:
	;
	v1821 = *(*int32)(unsafe.Add(mBase, uint32(v1789)+12))
	v1824 = v1821 + v1797<<(uint(int32(2))%32)
	if v1824 == int32(0) {
		goto L468
	} else {
		goto L470
	}
L470:
	;
	v1829 = *(*int32)(unsafe.Add(mBase, uint32(v1814)))
	v1830 = *(*int32)(unsafe.Add(mBase, uint32(v1824)))
	if v1829 == v1830 {
		v1797 = v1797 + int32(1)
		goto L457
	} else {
		goto L471
	}
L471:
	;
	v1857 = v1772
	goto L446
L472:
	;
	goto L455
L473:
	;
	v1846 = *(*int32)(unsafe.Add(mBase, uint32(v1845)+4))
	v1848 = v1846
	goto L475
L474:
	;
	v1848 = int32(0)
	goto L475
L475:
	;
	v1849 = F_bms_is_subset(m, v1848, v1753)
	mBase = m.M
	if v1849 != 0 {
		goto L476
	} else {
		goto L477
	}
L476:
	;
	v1850 = v1782
	goto L478
L477:
	;
	v1850 = v1772
	goto L478
L478:
	;
	v1857 = v1850
	goto L446
L479:
	;
	goto L445
L480:
	;
	if v2001 != 0 {
		goto L523
	} else {
		goto L524
	}
L481:
	;
	goto L480
L482:
	;
	v1900 = v1879
	v1903 = v1879
	goto L487
L483:
	;
	v1891 = *(*int32)(unsafe.Add(mBase, uint32(v1880)+4))
	if int32(0) < v1891 {
		goto L482
	} else {
		goto L486
	}
L484:
	;
	goto L485
L485:
	;
	v2001 = v1879
	goto L481
L486:
	;
	goto L485
L487:
	;
	v1906 = *(*int32)(unsafe.Add(mBase, uint32(v1880)+12))
	v1910 = *(*int32)(unsafe.Add(mBase, uint32(v1906+v1903<<(uint(int32(2))%32))))
	goto L491
L488:
	;
	v2001 = v1985
	goto L481
L489:
	;
	v1992 = v1903 + int32(1)
	v1993 = *(*int32)(unsafe.Add(mBase, uint32(v1880)+4))
	if v1992 < v1993 {
		v1900 = v1985
		v1903 = v1992
		goto L487
	} else {
		goto L522
	}
L491:
	;
	goto L492
L492:
	;
	if v1900 != 0 {
		goto L494
	} else {
		goto L495
	}
L494:
	;
	v1914 = F_compare_path_costs(m, v1900, v1910, int32(1))
	mBase = m.M
	if v1914 <= int32(0) {
		v1985 = v1900
		goto L489
	} else {
		goto L497
	}
L495:
	;
	goto L496
L496:
	;
	v1917 = *(*int32)(unsafe.Add(mBase, uint32(v1910)+64))
	if v1474 == v1917 {
		goto L498
	} else {
		goto L499
	}
L497:
	;
	goto L496
L498:
	;
	v1973 = *(*int32)(unsafe.Add(mBase, uint32(v1910)+16))
	if v1973 != 0 {
		goto L516
	} else {
		goto L517
	}
L499:
	;
	v1925 = int32(0)
	goto L500
L500:
	;
	v1932 = int32(0)
	if v1474 == v1932 {
		v1942 = v1932
		goto L502
	} else {
		goto L503
	}
L501:
	;
	if v1942 != 0 {
		v1985 = v1900
		goto L489
	} else {
		goto L515
	}
L502:
	;
	if v1917 != 0 {
		goto L506
	} else {
		goto L507
	}
L503:
	;
	v1936 = *(*int32)(unsafe.Add(mBase, uint32(v1474)+4))
	if v1936 <= v1925 {
		v1942 = int32(0)
		goto L502
	} else {
		goto L504
	}
L504:
	;
	v1938 = *(*int32)(unsafe.Add(mBase, uint32(v1474)+12))
	v1942 = v1938 + v1925<<(uint(int32(2))%32)
	goto L502
L505:
	;
	if v1942 == int32(0) {
		goto L511
	} else {
		goto L512
	}
L506:
	;
	v1943 = *(*int32)(unsafe.Add(mBase, uint32(v1917)+4))
	if v1925 < v1943 {
		goto L505
	} else {
		goto L509
	}
L507:
	;
	goto L508
L508:
	;
	if v1942 == int32(0) {
		goto L498
	} else {
		goto L510
	}
L509:
	;
	goto L508
L510:
	;
	v1985 = v1900
	goto L489
L511:
	;
	goto L501
L512:
	;
	v1949 = *(*int32)(unsafe.Add(mBase, uint32(v1917)+12))
	v1952 = v1949 + v1925<<(uint(int32(2))%32)
	if v1952 == int32(0) {
		goto L511
	} else {
		goto L513
	}
L513:
	;
	v1957 = *(*int32)(unsafe.Add(mBase, uint32(v1942)))
	v1958 = *(*int32)(unsafe.Add(mBase, uint32(v1952)))
	if v1957 == v1958 {
		v1925 = v1925 + int32(1)
		goto L500
	} else {
		goto L514
	}
L514:
	;
	v1985 = v1900
	goto L489
L515:
	;
	goto L498
L516:
	;
	v1974 = *(*int32)(unsafe.Add(mBase, uint32(v1973)+4))
	v1976 = v1974
	goto L518
L517:
	;
	v1976 = int32(0)
	goto L518
L518:
	;
	v1977 = F_bms_is_subset(m, v1976, v1879)
	mBase = m.M
	if v1977 != 0 {
		goto L519
	} else {
		goto L520
	}
L519:
	;
	v1978 = v1910
	goto L521
L520:
	;
	v1978 = v1900
	goto L521
L521:
	;
	v1985 = v1978
	goto L489
L522:
	;
	goto L488
L523:
	;
	v2007 = v1873
	goto L525
L524:
	;
	v2007 = v1879
	goto L525
L525:
	;
	if v2007 == int32(0) {
		goto L526
	} else {
		goto L527
	}
L526:
	;
	v2010 = *(*int32)(unsafe.Add(mBase, uint32(v1751)+48))
	v2011 = v2010
	v2012 = v2010
	goto L528
L527:
	;
	v2011 = v2001
	v2012 = v1873
	goto L528
L528:
	;
	v2014 = *(*float64)(unsafe.Add(mBase, uint32(v838)+296))
	if base.F64_gt(v2014, float64(0)) != 0 {
		goto L529
	} else {
		goto L530
	}
L529:
	;
	v2017 = *(*int32)(unsafe.Add(mBase, uint32(v1751)+32))
	if base.F64_ge(v2014, float64(1)) != 0 {
		goto L532
	} else {
		goto L533
	}
L530:
	;
	v2135 = int32(0)
	goto L531
L531:
	;
	v2137 = *(*int32)(unsafe.Add(mBase, uint32(v2012)))
	if v1714 != 0 {
		goto L579
	} else {
		goto L580
	}
L532:
	;
	v2020 = *(*float64)(unsafe.Add(mBase, uint32(v2011)+32))
	v2022 = base.F64_div(v2014, v2020)
	goto L534
L533:
	;
	v2022 = v2014
	goto L534
L534:
	;
	v2023 = int32(0)
	if v2017 != 0 {
		goto L538
	} else {
		goto L539
	}
L535:
	;
	if v2129 != 0 {
		goto L574
	} else {
		goto L575
	}
L536:
	;
	goto L535
L537:
	;
	v2038 = v2023
	v2041 = v2023
	goto L542
L538:
	;
	v2030 = *(*int32)(unsafe.Add(mBase, uint32(v2017)+4))
	if int32(0) < v2030 {
		goto L537
	} else {
		goto L541
	}
L539:
	;
	goto L540
L540:
	;
	v2129 = v2023
	goto L536
L541:
	;
	goto L540
L542:
	;
	v2043 = *(*int32)(unsafe.Add(mBase, uint32(v2017)+12))
	v2047 = *(*int32)(unsafe.Add(mBase, uint32(v2043+v2041<<(uint(int32(2))%32))))
	if v2038 != 0 {
		goto L545
	} else {
		goto L546
	}
L543:
	;
	v2129 = v2115
	goto L536
L544:
	;
	v2121 = v2041 + int32(1)
	v2122 = *(*int32)(unsafe.Add(mBase, uint32(v2017)+4))
	if v2121 < v2122 {
		v2038 = v2115
		v2041 = v2121
		goto L542
	} else {
		goto L573
	}
L545:
	;
	v2048 = F_compare_fractional_path_costs(m, v2038, v2047, v2022)
	mBase = m.M
	if v2048 <= int32(0) {
		v2115 = v2038
		goto L544
	} else {
		goto L548
	}
L546:
	;
	goto L547
L547:
	;
	v2051 = *(*int32)(unsafe.Add(mBase, uint32(v2047)+64))
	if v1474 == v2051 {
		goto L549
	} else {
		goto L550
	}
L548:
	;
	goto L547
L549:
	;
	v2103 = *(*int32)(unsafe.Add(mBase, uint32(v2047)+16))
	if v2103 != 0 {
		goto L567
	} else {
		goto L568
	}
L550:
	;
	v2057 = int32(0)
	goto L551
L551:
	;
	v2064 = int32(0)
	if v1474 == v2064 {
		v2074 = v2064
		goto L553
	} else {
		goto L554
	}
L552:
	;
	if v2074 != 0 {
		v2115 = v2038
		goto L544
	} else {
		goto L566
	}
L553:
	;
	if v2051 != 0 {
		goto L557
	} else {
		goto L558
	}
L554:
	;
	v2068 = *(*int32)(unsafe.Add(mBase, uint32(v1474)+4))
	if v2068 <= v2057 {
		v2074 = int32(0)
		goto L553
	} else {
		goto L555
	}
L555:
	;
	v2070 = *(*int32)(unsafe.Add(mBase, uint32(v1474)+12))
	v2074 = v2070 + v2057<<(uint(int32(2))%32)
	goto L553
L556:
	;
	if v2074 == int32(0) {
		goto L562
	} else {
		goto L563
	}
L557:
	;
	v2075 = *(*int32)(unsafe.Add(mBase, uint32(v2051)+4))
	if v2057 < v2075 {
		goto L556
	} else {
		goto L560
	}
L558:
	;
	goto L559
L559:
	;
	if v2074 == int32(0) {
		goto L549
	} else {
		goto L561
	}
L560:
	;
	goto L559
L561:
	;
	v2115 = v2038
	goto L544
L562:
	;
	goto L552
L563:
	;
	v2081 = *(*int32)(unsafe.Add(mBase, uint32(v2051)+12))
	v2084 = v2081 + v2057<<(uint(int32(2))%32)
	if v2084 == int32(0) {
		goto L562
	} else {
		goto L564
	}
L564:
	;
	v2089 = *(*int32)(unsafe.Add(mBase, uint32(v2074)))
	v2090 = *(*int32)(unsafe.Add(mBase, uint32(v2084)))
	if v2089 == v2090 {
		v2057 = v2057 + int32(1)
		goto L551
	} else {
		goto L565
	}
L565:
	;
	v2115 = v2038
	goto L544
L566:
	;
	goto L549
L567:
	;
	v2104 = *(*int32)(unsafe.Add(mBase, uint32(v2103)+4))
	v2106 = v2104
	goto L569
L568:
	;
	v2106 = int32(0)
	goto L569
L569:
	;
	v2108 = F_bms_is_subset(m, v2106, int32(0))
	mBase = m.M
	if v2108 != 0 {
		goto L570
	} else {
		goto L571
	}
L570:
	;
	v2109 = v2047
	goto L572
L571:
	;
	v2109 = v2038
	goto L572
L572:
	;
	v2115 = v2109
	goto L544
L573:
	;
	goto L543
L574:
	;
	v2134 = v2129
	goto L576
L575:
	;
	v2134 = v2011
	goto L576
L576:
	;
	v2135 = v2134
	goto L531
L577:
	;
	v2262 = base.B2i32(v2011 != v2012) | v1736
	v2263 = v1733 + v1716
	if v1715 != v2263 {
		v1729 = v2258
		v1730 = v2259
		v1732 = v2260
		v1733 = v2263
		v1736 = v2262
		goto L435
	} else {
		goto L641
	}
L578:
	;
	v2254 = F_lappend(m, v1732, v2135)
	mBase = m.M
	v2255 = m.ExcPending
	if v2255 != 0 {
		goto L30
	} else {
		goto L640
	}
L579:
	;
	switch v2137 - int32(290) {
	case 0:
		goto L585
	case 1:
		goto L584
	default:
		v2155 = v2012
		goto L582
	}
L580:
	;
	goto L581
L581:
	;
	switch v2137 - int32(290) {
	case 0:
		goto L612
	case 1:
		goto L611
	default:
		goto L610
	}
L582:
	;
	v2157 = *(*int32)(unsafe.Add(mBase, uint32(v2011)))
	switch v2157 - int32(290) {
	case 0:
		goto L593
	case 1:
		goto L592
	default:
		v2175 = v2011
		goto L590
	}
L583:
	;
	v2153 = *(*int32)(unsafe.Add(mBase, uint32(v2152)+12))
	v2154 = *(*int32)(unsafe.Add(mBase, uint32(v2153)))
	v2155 = v2154
	goto L582
L584:
	;
	v2146 = *(*int32)(unsafe.Add(mBase, uint32(v2012)+72))
	if v2146 == int32(0) {
		v2155 = v2012
		goto L582
	} else {
		goto L588
	}
L585:
	;
	v2140 = *(*int32)(unsafe.Add(mBase, uint32(v2012)+72))
	if v2140 == int32(0) {
		v2155 = v2012
		goto L582
	} else {
		goto L586
	}
L586:
	;
	v2143 = *(*int32)(unsafe.Add(mBase, uint32(v2140)+4))
	if v2143 == int32(1) {
		v2152 = v2140
		goto L583
	} else {
		goto L587
	}
L587:
	;
	v2155 = v2012
	goto L582
L588:
	;
	v2149 = *(*int32)(unsafe.Add(mBase, uint32(v2146)+4))
	if v2149 != int32(1) {
		v2155 = v2012
		goto L582
	} else {
		goto L589
	}
L589:
	;
	v2152 = v2146
	goto L583
L590:
	;
	v2177 = F_lappend(m, v1730, v2155)
	mBase = m.M
	v2178 = m.ExcPending
	if v2178 != 0 {
		goto L30
	} else {
		goto L598
	}
L591:
	;
	v2173 = *(*int32)(unsafe.Add(mBase, uint32(v2172)+12))
	v2174 = *(*int32)(unsafe.Add(mBase, uint32(v2173)))
	v2175 = v2174
	goto L590
L592:
	;
	v2166 = *(*int32)(unsafe.Add(mBase, uint32(v2011)+72))
	if v2166 == int32(0) {
		v2175 = v2011
		goto L590
	} else {
		goto L596
	}
L593:
	;
	v2160 = *(*int32)(unsafe.Add(mBase, uint32(v2011)+72))
	if v2160 == int32(0) {
		v2175 = v2011
		goto L590
	} else {
		goto L594
	}
L594:
	;
	v2163 = *(*int32)(unsafe.Add(mBase, uint32(v2160)+4))
	if v2163 == int32(1) {
		v2172 = v2160
		goto L591
	} else {
		goto L595
	}
L595:
	;
	v2175 = v2011
	goto L590
L596:
	;
	v2169 = *(*int32)(unsafe.Add(mBase, uint32(v2166)+4))
	if v2169 != int32(1) {
		v2175 = v2011
		goto L590
	} else {
		goto L597
	}
L597:
	;
	v2172 = v2166
	goto L591
L598:
	;
	v2179 = F_lappend(m, v1729, v2175)
	mBase = m.M
	v2180 = m.ExcPending
	if v2180 != 0 {
		goto L30
	} else {
		goto L599
	}
L599:
	;
	if v2135 == int32(0) {
		v2258 = v2179
		v2259 = v2177
		v2260 = v1732
		goto L577
	} else {
		goto L600
	}
L600:
	;
	v2183 = *(*int32)(unsafe.Add(mBase, uint32(v2135)))
	switch v2183 - int32(290) {
	case 0:
		goto L603
	case 1:
		goto L602
	default:
		v2251 = v2179
		v2252 = v2177
		goto L578
	}
L601:
	;
	v2199 = *(*int32)(unsafe.Add(mBase, uint32(v2198)+12))
	v2200 = *(*int32)(unsafe.Add(mBase, uint32(v2199)))
	v2201 = F_lappend(m, v1732, v2200)
	mBase = m.M
	v2202 = m.ExcPending
	if v2202 != 0 {
		goto L30
	} else {
		goto L608
	}
L602:
	;
	v2192 = *(*int32)(unsafe.Add(mBase, uint32(v2135)+72))
	if v2192 == int32(0) {
		v2251 = v2179
		v2252 = v2177
		goto L578
	} else {
		goto L606
	}
L603:
	;
	v2186 = *(*int32)(unsafe.Add(mBase, uint32(v2135)+72))
	if v2186 == int32(0) {
		v2251 = v2179
		v2252 = v2177
		goto L578
	} else {
		goto L604
	}
L604:
	;
	v2189 = *(*int32)(unsafe.Add(mBase, uint32(v2186)+4))
	if v2189 == int32(1) {
		v2198 = v2186
		goto L601
	} else {
		goto L605
	}
L605:
	;
	v2251 = v2179
	v2252 = v2177
	goto L578
L606:
	;
	v2195 = *(*int32)(unsafe.Add(mBase, uint32(v2192)+4))
	if v2195 != int32(1) {
		v2251 = v2179
		v2252 = v2177
		goto L578
	} else {
		goto L607
	}
L607:
	;
	v2198 = v2192
	goto L601
L608:
	;
	v2258 = v2179
	v2259 = v2177
	v2260 = v2201
	goto L577
L609:
	;
	v2218 = *(*int32)(unsafe.Add(mBase, uint32(v2011)))
	switch v2218 - int32(290) {
	case 0:
		goto L623
	case 1:
		goto L622
	default:
		goto L621
	}
L610:
	;
	v2215 = F_lappend(m, v1730, v2012)
	mBase = m.M
	v2216 = m.ExcPending
	if v2216 != 0 {
		goto L30
	} else {
		goto L619
	}
L611:
	;
	v2212 = *(*int32)(unsafe.Add(mBase, uint32(v2012)+72))
	v2213 = F_list_concat(m, v1730, v2212)
	mBase = m.M
	v2214 = m.ExcPending
	if v2214 != 0 {
		goto L30
	} else {
		goto L618
	}
L612:
	;
	v2205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2012)+20)))
	if v2205 == int32(1) {
		goto L613
	} else {
		goto L614
	}
L613:
	;
	v2208 = *(*int32)(unsafe.Add(mBase, uint32(v2012)+76))
	if v2208 != 0 {
		goto L610
	} else {
		goto L616
	}
L614:
	;
	goto L615
L615:
	;
	v2209 = *(*int32)(unsafe.Add(mBase, uint32(v2012)+72))
	v2210 = F_list_concat(m, v1730, v2209)
	mBase = m.M
	v2211 = m.ExcPending
	if v2211 != 0 {
		goto L30
	} else {
		goto L617
	}
L616:
	;
	goto L615
L617:
	;
	v2217 = v2210
	goto L609
L618:
	;
	v2217 = v2213
	goto L609
L619:
	;
	v2217 = v2215
	goto L609
L620:
	;
	if v2135 == int32(0) {
		v2258 = v2233
		v2259 = v2217
		v2260 = v1732
		goto L577
	} else {
		goto L631
	}
L621:
	;
	v2231 = F_lappend(m, v1729, v2011)
	mBase = m.M
	v2232 = m.ExcPending
	if v2232 != 0 {
		goto L30
	} else {
		goto L630
	}
L622:
	;
	v2228 = *(*int32)(unsafe.Add(mBase, uint32(v2011)+72))
	v2229 = F_list_concat(m, v1729, v2228)
	mBase = m.M
	v2230 = m.ExcPending
	if v2230 != 0 {
		goto L30
	} else {
		goto L629
	}
L623:
	;
	v2221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2011)+20)))
	if v2221 == int32(1) {
		goto L624
	} else {
		goto L625
	}
L624:
	;
	v2224 = *(*int32)(unsafe.Add(mBase, uint32(v2011)+76))
	if v2224 != 0 {
		goto L621
	} else {
		goto L627
	}
L625:
	;
	goto L626
L626:
	;
	v2225 = *(*int32)(unsafe.Add(mBase, uint32(v2011)+72))
	v2226 = F_list_concat(m, v1729, v2225)
	mBase = m.M
	v2227 = m.ExcPending
	if v2227 != 0 {
		goto L30
	} else {
		goto L628
	}
L627:
	;
	goto L626
L628:
	;
	v2233 = v2226
	goto L620
L629:
	;
	v2233 = v2229
	goto L620
L630:
	;
	v2233 = v2231
	goto L620
L631:
	;
	v2236 = *(*int32)(unsafe.Add(mBase, uint32(v2135)))
	switch v2236 - int32(290) {
	case 0:
		goto L633
	case 1:
		goto L632
	default:
		v2251 = v2233
		v2252 = v2217
		goto L578
	}
L632:
	;
	v2246 = *(*int32)(unsafe.Add(mBase, uint32(v2135)+72))
	v2247 = F_list_concat(m, v1732, v2246)
	mBase = m.M
	v2248 = m.ExcPending
	if v2248 != 0 {
		goto L30
	} else {
		goto L639
	}
L633:
	;
	v2239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2135)+20)))
	if v2239 == int32(1) {
		goto L634
	} else {
		goto L635
	}
L634:
	;
	v2242 = *(*int32)(unsafe.Add(mBase, uint32(v2135)+76))
	if v2242 != 0 {
		v2251 = v2233
		v2252 = v2217
		goto L578
	} else {
		goto L637
	}
L635:
	;
	goto L636
L636:
	;
	v2243 = *(*int32)(unsafe.Add(mBase, uint32(v2135)+72))
	v2244 = F_list_concat(m, v1732, v2243)
	mBase = m.M
	v2245 = m.ExcPending
	if v2245 != 0 {
		goto L30
	} else {
		goto L638
	}
L637:
	;
	goto L636
L638:
	;
	v2258 = v2233
	v2259 = v2217
	v2260 = v2244
	goto L577
L639:
	;
	v2258 = v2233
	v2259 = v2217
	v2260 = v2247
	goto L577
L640:
	;
	v2258 = v2251
	v2259 = v2252
	v2260 = v2254
	goto L577
L641:
	;
	goto L436
L642:
	;
	v2357 = v2265
	v2358 = v2265
	v2360 = v2265
	v2364 = v2265
	goto L343
L643:
	;
	v2303 = v2276
	v2304 = v2277
	v2306 = v2279
	v2310 = v2283
	goto L344
L644:
	;
	F_add_path(m, v839, v2326)
	mBase = m.M
	v2329 = m.ExcPending
	if v2329 != 0 {
		goto L30
	} else {
		goto L645
	}
L645:
	;
	if v2310&int32(1) != 0 {
		goto L646
	} else {
		goto L647
	}
L646:
	;
	v2332 = int32(0)
	v2337 = F_create_append_path(m, v838, v839, v2303, v2332, v1474, v2332, v2332, v2332, float64(-1))
	mBase = m.M
	v2338 = m.ExcPending
	if v2338 != 0 {
		goto L30
	} else {
		goto L649
	}
L647:
	;
	goto L648
L648:
	;
	if v2306 == int32(0) {
		goto L341
	} else {
		goto L651
	}
L649:
	;
	F_add_path(m, v839, v2337)
	mBase = m.M
	v2340 = m.ExcPending
	if v2340 != 0 {
		goto L30
	} else {
		goto L650
	}
L650:
	;
	goto L648
L651:
	;
	v2343 = int32(0)
	v2348 = F_create_append_path(m, v838, v839, v2306, v2343, v1474, v2343, v2343, v2343, float64(-1))
	mBase = m.M
	v2349 = m.ExcPending
	if v2349 != 0 {
		goto L30
	} else {
		goto L652
	}
L652:
	;
	v2414 = v2348
	goto L342
L653:
	;
	F_add_path(m, v839, v2375)
	mBase = m.M
	v2378 = m.ExcPending
	if v2378 != 0 {
		goto L30
	} else {
		goto L654
	}
L654:
	;
	if v2364&int32(1) != 0 {
		goto L655
	} else {
		goto L656
	}
L655:
	;
	v2381 = F_create_merge_append_path(m, v838, v839, v2357, v1474)
	mBase = m.M
	v2382 = m.ExcPending
	if v2382 != 0 {
		goto L30
	} else {
		goto L658
	}
L656:
	;
	goto L657
L657:
	;
	if v2360 == int32(0) {
		goto L341
	} else {
		goto L660
	}
L658:
	;
	F_add_path(m, v839, v2381)
	mBase = m.M
	v2384 = m.ExcPending
	if v2384 != 0 {
		goto L30
	} else {
		goto L659
	}
L659:
	;
	goto L657
L660:
	;
	v2387 = F_create_merge_append_path(m, v838, v839, v2360, v1474)
	mBase = m.M
	v2388 = m.ExcPending
	if v2388 != 0 {
		goto L30
	} else {
		goto L661
	}
L661:
	;
	v2414 = v2387
	goto L342
L662:
	;
	goto L341
L663:
	;
	goto L340
L664:
	;
	if v840 == int32(0) {
		goto L855
	} else {
		goto L856
	}
L665:
	;
	v2473 = int32(0)
	v2474 = *(*int32)(unsafe.Add(mBase, uint32(v855)+4))
	if v2474 <= v2473 {
		goto L664
	} else {
		goto L666
	}
L666:
	;
	v2484 = v2473
	goto L667
L667:
	;
	v2502 = *(*int32)(unsafe.Add(mBase, uint32(v855)+12))
	v2506 = *(*int32)(unsafe.Add(mBase, uint32(v2502+v2484<<(uint(int32(2))%32))))
	v2507 = int32(0)
	if v840 == v2507 {
		v3055 = v2507
		goto L670
	} else {
		goto L671
	}
L668:
	;
	goto L664
L669:
	;
	v3107 = v2484 + int32(1)
	v3108 = *(*int32)(unsafe.Add(mBase, uint32(v855)+4))
	if v3107 < v3108 {
		v2484 = v3107
		goto L667
	} else {
		goto L854
	}
L670:
	;
	v3072 = int32(0)
	v3077 = F_create_append_path(m, v838, v839, v3055, v3072, v3072, v2506, v3072, v3072, float64(-1))
	mBase = m.M
	v3078 = m.ExcPending
	if v3078 != 0 {
		goto L30
	} else {
		goto L852
	}
L671:
	;
	v2510 = int32(0)
	v2511 = *(*int32)(unsafe.Add(mBase, uint32(v840)+4))
	if v2511 <= v2510 {
		v3055 = v2507
		goto L670
	} else {
		goto L672
	}
L672:
	;
	v2522 = v2507
	v2528 = v2510
	goto L673
L673:
	;
	v2539 = *(*int32)(unsafe.Add(mBase, uint32(v840)+12))
	v2543 = *(*int32)(unsafe.Add(mBase, uint32(v2539+v2528<<(uint(int32(2))%32))))
	v2544 = *(*int32)(unsafe.Add(mBase, uint32(v2543)+32))
	if v2544 == int32(0) {
		goto L669
	} else {
		goto L675
	}
L674:
	;
	v3055 = v3042
	goto L670
L675:
	;
	v2547 = int32(0)
	if v2544 != 0 {
		goto L679
	} else {
		goto L680
	}
L676:
	;
	v2674 = *(*int32)(unsafe.Add(mBase, uint32(v2668)+16))
	if v2674 != 0 {
		goto L719
	} else {
		goto L720
	}
L677:
	;
	goto L676
L678:
	;
	v2567 = v2547
	v2570 = v2547
	goto L683
L679:
	;
	v2558 = *(*int32)(unsafe.Add(mBase, uint32(v2544)+4))
	if int32(0) < v2558 {
		goto L678
	} else {
		goto L682
	}
L680:
	;
	goto L681
L681:
	;
	v2668 = v2547
	goto L677
L682:
	;
	goto L681
L683:
	;
	v2573 = *(*int32)(unsafe.Add(mBase, uint32(v2544)+12))
	v2577 = *(*int32)(unsafe.Add(mBase, uint32(v2573+v2570<<(uint(int32(2))%32))))
	goto L687
L684:
	;
	v2668 = v2652
	goto L677
L685:
	;
	v2659 = v2570 + int32(1)
	v2660 = *(*int32)(unsafe.Add(mBase, uint32(v2544)+4))
	if v2659 < v2660 {
		v2567 = v2652
		v2570 = v2659
		goto L683
	} else {
		goto L718
	}
L687:
	;
	goto L688
L688:
	;
	if v2567 != 0 {
		goto L690
	} else {
		goto L691
	}
L690:
	;
	v2581 = F_compare_path_costs(m, v2567, v2577, int32(1))
	mBase = m.M
	if v2581 <= int32(0) {
		v2652 = v2567
		goto L685
	} else {
		goto L693
	}
L691:
	;
	goto L692
L692:
	;
	v2584 = *(*int32)(unsafe.Add(mBase, uint32(v2577)+64))
	if v2547 == v2584 {
		goto L694
	} else {
		goto L695
	}
L693:
	;
	goto L692
L694:
	;
	v2640 = *(*int32)(unsafe.Add(mBase, uint32(v2577)+16))
	if v2640 != 0 {
		goto L712
	} else {
		goto L713
	}
L695:
	;
	goto L696
L696:
	;
	goto L698
L697:
	;
	goto L711
L698:
	;
	if v2584 != 0 {
		goto L702
	} else {
		goto L703
	}
L701:
	;
	goto L707
L702:
	;
	v2610 = *(*int32)(unsafe.Add(mBase, uint32(v2584)+4))
	if int32(0) < v2610 {
		goto L701
	} else {
		goto L705
	}
L703:
	;
	goto L704
L704:
	;
	goto L694
L705:
	;
	goto L704
L707:
	;
	goto L697
L711:
	;
	goto L694
L712:
	;
	v2641 = *(*int32)(unsafe.Add(mBase, uint32(v2640)+4))
	v2643 = v2641
	goto L714
L713:
	;
	v2643 = int32(0)
	goto L714
L714:
	;
	v2644 = F_bms_is_subset(m, v2643, v2506)
	mBase = m.M
	if v2644 != 0 {
		goto L715
	} else {
		goto L716
	}
L715:
	;
	v2645 = v2577
	goto L717
L716:
	;
	v2645 = v2567
	goto L717
L717:
	;
	v2652 = v2645
	goto L685
L718:
	;
	goto L684
L719:
	;
	v2675 = *(*int32)(unsafe.Add(mBase, uint32(v2674)+4))
	v2677 = v2675
	goto L721
L720:
	;
	v2677 = int32(0)
	goto L721
L721:
	;
	v2678 = int32(0)
	v2685 = base.B2i32(v2677|v2506 == v2678)
	if v2677 == v2678 {
		v2724 = v2685
		goto L723
	} else {
		goto L724
	}
L722:
	;
	if v2724 == int32(0) {
		goto L734
	} else {
		goto L735
	}
L723:
	;
	goto L722
L724:
	;
	if v2506 == int32(0) {
		v2724 = v2685
		goto L723
	} else {
		goto L725
	}
L725:
	;
	v2691 = *(*int32)(unsafe.Add(mBase, uint32(v2677)+4))
	v2692 = *(*int32)(unsafe.Add(mBase, uint32(v2506)+4))
	if v2691 != v2692 {
		v2724 = int32(0)
		goto L723
	} else {
		goto L726
	}
L726:
	;
	v2694 = int32(1)
	if v2691 <= v2694 {
		goto L727
	} else {
		goto L728
	}
L727:
	;
	v2697 = v2694
	goto L729
L728:
	;
	v2697 = v2691
	goto L729
L729:
	;
	v2698 = int32(8)
	v2703 = int32(0)
	goto L730
L730:
	;
	v2711 = v2703 << (uint(int32(2)) % 32)
	v2713 = *(*int32)(unsafe.Add(mBase, uint32(v2677+v2698+v2711)))
	v2715 = *(*int32)(unsafe.Add(mBase, uint32(v2711+(v2506+v2698))))
	v2716 = base.B2i32(v2713 == v2715)
	if v2715 != v2713 {
		v2724 = v2716
		goto L723
	} else {
		goto L732
	}
L731:
	;
	v2724 = v2716
	goto L723
L732:
	;
	v2719 = v2703 + int32(1)
	if v2719 != v2697 {
		v2703 = v2719
		goto L730
	} else {
		goto L733
	}
L733:
	;
	goto L731
L734:
	;
	v2730 = *(*int32)(unsafe.Add(mBase, uint32(v2543)+32))
	if v2730 == int32(0) {
		goto L669
	} else {
		goto L737
	}
L735:
	;
	v3005 = v2668
	goto L736
L736:
	;
	v3027 = *(*int32)(unsafe.Add(mBase, uint32(v3005)))
	switch v3027 - int32(290) {
	case 0:
		goto L843
	case 1:
		goto L842
	default:
		goto L841
	}
L737:
	;
	v2733 = int32(0)
	v2734 = *(*int32)(unsafe.Add(mBase, uint32(v2730)+4))
	if v2734 <= v2733 {
		goto L669
	} else {
		goto L738
	}
L738:
	;
	v2740 = v2733
	v2752 = v2547
	goto L739
L739:
	;
	v2762 = *(*int32)(unsafe.Add(mBase, uint32(v2730)+12))
	v2766 = *(*int32)(unsafe.Add(mBase, uint32(v2762+v2752<<(uint(int32(2))%32))))
	v2767 = *(*int32)(unsafe.Add(mBase, uint32(v2766)+16))
	if v2767 != 0 {
		goto L742
	} else {
		goto L743
	}
L740:
	;
	if v2994 == int32(0) {
		goto L669
	} else {
		goto L839
	}
L741:
	;
	v2997 = v2752 + int32(1)
	v2998 = *(*int32)(unsafe.Add(mBase, uint32(v2730)+4))
	if v2997 < v2998 {
		v2740 = v2994
		v2752 = v2997
		goto L739
	} else {
		goto L838
	}
L742:
	;
	v2768 = *(*int32)(unsafe.Add(mBase, uint32(v2767)+4))
	v2770 = v2768
	goto L744
L743:
	;
	v2770 = int32(0)
	goto L744
L744:
	;
	v2771 = int32(0)
	if v2770 == v2771 {
		goto L746
	} else {
		goto L747
	}
L745:
	;
	if v2824 == int32(0) {
		goto L759
	} else {
		goto L760
	}
L746:
	;
	v2824 = int32(1)
	goto L745
L747:
	;
	goto L748
L748:
	;
	if v2506 == int32(0) {
		v2815 = v2771
		goto L749
	} else {
		goto L750
	}
L749:
	;
	v2824 = v2815
	goto L745
L750:
	;
	v2780 = *(*int32)(unsafe.Add(mBase, uint32(v2770)+4))
	v2781 = *(*int32)(unsafe.Add(mBase, uint32(v2506)+4))
	if v2781 < v2780 {
		v2815 = v2771
		goto L749
	} else {
		goto L751
	}
L751:
	;
	v2783 = int32(1)
	if v2780 <= v2783 {
		goto L752
	} else {
		goto L753
	}
L752:
	;
	v2786 = v2783
	goto L754
L753:
	;
	v2786 = v2780
	goto L754
L754:
	;
	v2787 = int32(8)
	v2792 = int32(0)
	goto L755
L755:
	;
	v2799 = v2792 << (uint(int32(2)) % 32)
	v2801 = *(*int32)(unsafe.Add(mBase, uint32(v2770+v2787+v2799)))
	v2803 = *(*int32)(unsafe.Add(mBase, uint32(v2799+(v2506+v2787))))
	v2806 = v2801 & (v2803 ^ int32(-1))
	v2808 = base.B2i32(v2806 == int32(0))
	if v2806 != 0 {
		v2815 = v2808
		goto L749
	} else {
		goto L757
	}
L756:
	;
	v2815 = v2808
	goto L749
L757:
	;
	v2810 = v2792 + int32(1)
	if v2810 != v2786 {
		v2792 = v2810
		goto L755
	} else {
		goto L758
	}
L758:
	;
	goto L756
L759:
	;
	v2994 = v2740
	goto L741
L760:
	;
	goto L761
L761:
	;
	if v2740 == int32(0) {
		goto L762
	} else {
		goto L763
	}
L762:
	;
	v2880 = *(*int32)(unsafe.Add(mBase, uint32(v2766)+16))
	if v2880 != 0 {
		goto L789
	} else {
		goto L790
	}
L763:
	;
	v2834 = *(*int32)(unsafe.Add(mBase, uint32(v2740)+40))
	v2835 = *(*int32)(unsafe.Add(mBase, uint32(v2766)+40))
	if v2834 != v2835 {
		goto L765
	} else {
		goto L766
	}
L764:
	;
	if int32(0) < v2877 {
		goto L762
	} else {
		goto L788
	}
L765:
	;
	if v2834 < v2835 {
		goto L768
	} else {
		goto L769
	}
L766:
	;
	goto L767
L767:
	;
	goto L774
L768:
	;
	v2840 = int32(-1)
	goto L770
L769:
	;
	v2840 = int32(1)
	goto L770
L770:
	;
	v2877 = v2840
	goto L764
L771:
	;
	v2877 = v2871
	goto L764
L772:
	;
	v2871 = int32(0)
	goto L771
L774:
	;
	goto L775
L775:
	;
	v2856 = int32(-1)
	v2857 = *(*float64)(unsafe.Add(mBase, uint32(v2740)+56))
	v2858 = *(*float64)(unsafe.Add(mBase, uint32(v2766)+56))
	if base.F64_lt(v2857, v2858) != 0 {
		v2871 = v2856
		goto L771
	} else {
		goto L782
	}
L782:
	;
	if base.F64_gt(v2857, v2858) != 0 {
		goto L783
	} else {
		goto L784
	}
L783:
	;
	v2877 = int32(1)
	goto L764
L784:
	;
	goto L785
L785:
	;
	v2862 = *(*float64)(unsafe.Add(mBase, uint32(v2740)+48))
	v2863 = *(*float64)(unsafe.Add(mBase, uint32(v2766)+48))
	if base.F64_lt(v2862, v2863) != 0 {
		v2871 = v2856
		goto L771
	} else {
		goto L786
	}
L786:
	;
	if base.F64_gt(v2862, v2863) != 0 {
		v2871 = int32(1)
		goto L771
	} else {
		goto L787
	}
L787:
	;
	goto L772
L788:
	;
	v2994 = v2740
	goto L741
L789:
	;
	v2881 = *(*int32)(unsafe.Add(mBase, uint32(v2880)+4))
	v2883 = v2881
	goto L791
L790:
	;
	v2883 = int32(0)
	goto L791
L791:
	;
	v2884 = int32(0)
	v2891 = base.B2i32(v2883|v2506 == v2884)
	if v2883 == v2884 {
		v2930 = v2891
		goto L793
	} else {
		goto L794
	}
L792:
	;
	if v2930 != 0 {
		v2994 = v2766
		goto L741
	} else {
		goto L804
	}
L793:
	;
	goto L792
L794:
	;
	if v2506 == int32(0) {
		v2930 = v2891
		goto L793
	} else {
		goto L795
	}
L795:
	;
	v2897 = *(*int32)(unsafe.Add(mBase, uint32(v2883)+4))
	v2898 = *(*int32)(unsafe.Add(mBase, uint32(v2506)+4))
	if v2897 != v2898 {
		v2930 = int32(0)
		goto L793
	} else {
		goto L796
	}
L796:
	;
	v2900 = int32(1)
	if v2897 <= v2900 {
		goto L797
	} else {
		goto L798
	}
L797:
	;
	v2903 = v2900
	goto L799
L798:
	;
	v2903 = v2897
	goto L799
L799:
	;
	v2904 = int32(8)
	v2909 = int32(0)
	goto L800
L800:
	;
	v2917 = v2909 << (uint(int32(2)) % 32)
	v2919 = *(*int32)(unsafe.Add(mBase, uint32(v2883+v2904+v2917)))
	v2921 = *(*int32)(unsafe.Add(mBase, uint32(v2917+(v2506+v2904))))
	v2922 = base.B2i32(v2919 == v2921)
	if v2921 != v2919 {
		v2930 = v2922
		goto L793
	} else {
		goto L802
	}
L801:
	;
	v2930 = v2922
	goto L793
L802:
	;
	v2925 = v2909 + int32(1)
	if v2925 != v2903 {
		v2909 = v2925
		goto L800
	} else {
		goto L803
	}
L803:
	;
	goto L801
L804:
	;
	v2935 = F_reparameterize_path(m, v838, v2766, v2506, float64(1))
	mBase = m.M
	v2936 = m.ExcPending
	if v2936 != 0 {
		goto L30
	} else {
		goto L805
	}
L805:
	;
	if v2935 != 0 {
		goto L806
	} else {
		goto L807
	}
L806:
	;
	v2937 = v2935
	goto L808
L807:
	;
	v2937 = v2740
	goto L808
L808:
	;
	if v2740 == int32(0) {
		v2994 = v2937
		goto L741
	} else {
		goto L809
	}
L809:
	;
	if v2935 == int32(0) {
		v2994 = v2937
		goto L741
	} else {
		goto L810
	}
L810:
	;
	v2947 = *(*int32)(unsafe.Add(mBase, uint32(v2740)+40))
	v2948 = *(*int32)(unsafe.Add(mBase, uint32(v2935)+40))
	if v2947 != v2948 {
		goto L812
	} else {
		goto L813
	}
L811:
	;
	if v2990 <= int32(0) {
		goto L835
	} else {
		goto L836
	}
L812:
	;
	if v2947 < v2948 {
		goto L815
	} else {
		goto L816
	}
L813:
	;
	goto L814
L814:
	;
	goto L821
L815:
	;
	v2953 = int32(-1)
	goto L817
L816:
	;
	v2953 = int32(1)
	goto L817
L817:
	;
	v2990 = v2953
	goto L811
L818:
	;
	v2990 = v2984
	goto L811
L819:
	;
	v2984 = int32(0)
	goto L818
L821:
	;
	goto L822
L822:
	;
	v2969 = int32(-1)
	v2970 = *(*float64)(unsafe.Add(mBase, uint32(v2740)+56))
	v2971 = *(*float64)(unsafe.Add(mBase, uint32(v2935)+56))
	if base.F64_lt(v2970, v2971) != 0 {
		v2984 = v2969
		goto L818
	} else {
		goto L829
	}
L829:
	;
	if base.F64_gt(v2970, v2971) != 0 {
		goto L830
	} else {
		goto L831
	}
L830:
	;
	v2990 = int32(1)
	goto L811
L831:
	;
	goto L832
L832:
	;
	v2975 = *(*float64)(unsafe.Add(mBase, uint32(v2740)+48))
	v2976 = *(*float64)(unsafe.Add(mBase, uint32(v2935)+48))
	if base.F64_lt(v2975, v2976) != 0 {
		v2984 = v2969
		goto L818
	} else {
		goto L833
	}
L833:
	;
	if base.F64_gt(v2975, v2976) != 0 {
		v2984 = int32(1)
		goto L818
	} else {
		goto L834
	}
L834:
	;
	goto L819
L835:
	;
	v2993 = v2740
	goto L837
L836:
	;
	v2993 = v2935
	goto L837
L837:
	;
	v2994 = v2993
	goto L741
L838:
	;
	goto L740
L839:
	;
	v3005 = v2994
	goto L736
L840:
	;
	v3044 = v2528 + int32(1)
	v3045 = *(*int32)(unsafe.Add(mBase, uint32(v840)+4))
	if v3044 < v3045 {
		v2522 = v3042
		v2528 = v3044
		goto L673
	} else {
		goto L851
	}
L841:
	;
	v3040 = F_lappend(m, v2522, v3005)
	mBase = m.M
	v3041 = m.ExcPending
	if v3041 != 0 {
		goto L30
	} else {
		goto L850
	}
L842:
	;
	v3037 = *(*int32)(unsafe.Add(mBase, uint32(v3005)+72))
	v3038 = F_list_concat(m, v2522, v3037)
	mBase = m.M
	v3039 = m.ExcPending
	if v3039 != 0 {
		goto L30
	} else {
		goto L849
	}
L843:
	;
	v3030 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3005)+20)))
	if v3030 == int32(1) {
		goto L844
	} else {
		goto L845
	}
L844:
	;
	v3033 = *(*int32)(unsafe.Add(mBase, uint32(v3005)+76))
	if v3033 != 0 {
		goto L841
	} else {
		goto L847
	}
L845:
	;
	goto L846
L846:
	;
	v3034 = *(*int32)(unsafe.Add(mBase, uint32(v3005)+72))
	v3035 = F_list_concat(m, v2522, v3034)
	mBase = m.M
	v3036 = m.ExcPending
	if v3036 != 0 {
		goto L30
	} else {
		goto L848
	}
L847:
	;
	goto L846
L848:
	;
	v3042 = v3035
	goto L840
L849:
	;
	v3042 = v3038
	goto L840
L850:
	;
	v3042 = v3040
	goto L840
L851:
	;
	goto L674
L852:
	;
	F_add_path(m, v839, v3077)
	mBase = m.M
	v3080 = m.ExcPending
	if v3080 != 0 {
		goto L30
	} else {
		goto L853
	}
L853:
	;
	goto L669
L854:
	;
	goto L668
L855:
	;
	m.G0 = v850 + int32(32)
	return
L856:
	;
	v3137 = *(*int32)(unsafe.Add(mBase, uint32(v840)+4))
	if v3137 != int32(1) {
		goto L855
	} else {
		goto L857
	}
L857:
	;
	v3140 = *(*int32)(unsafe.Add(mBase, uint32(v840)+12))
	v3141 = *(*int32)(unsafe.Add(mBase, uint32(v3140)))
	v3142 = *(*int32)(unsafe.Add(mBase, uint32(v3141)+40))
	if v3142 == int32(0) {
		goto L855
	} else {
		goto L858
	}
L858:
	;
	v3145 = *(*int32)(unsafe.Add(mBase, uint32(v3142)+4))
	if v3145 < int32(2) {
		goto L855
	} else {
		goto L859
	}
L859:
	;
	v3152 = int32(1)
	goto L860
L860:
	;
	v3174 = *(*int32)(unsafe.Add(mBase, uint32(v3142)+12))
	v3178 = *(*int32)(unsafe.Add(mBase, uint32(v3174+v3152<<(uint(int32(2))%32))))
	v3179 = *(*int32)(unsafe.Add(mBase, uint32(v3178)+64))
	if v3179 != 0 {
		goto L862
	} else {
		goto L863
	}
L861:
	;
	goto L855
L862:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v850)+12)) = v3178
	*(*int32)(unsafe.Add(mBase, uint32(v850)+16)) = v3178
	v3186 = F_list_make1_impl(m, int32(1), v850+int32(12))
	mBase = m.M
	v3187 = m.ExcPending
	if v3187 != 0 {
		goto L30
	} else {
		goto L865
	}
L863:
	;
	goto L864
L864:
	;
	v3197 = v3152 + int32(1)
	v3198 = *(*int32)(unsafe.Add(mBase, uint32(v3142)+4))
	if v3197 < v3198 {
		v3152 = v3197
		goto L860
	} else {
		goto L868
	}
L865:
	;
	v3188 = int32(0)
	v3190 = *(*int32)(unsafe.Add(mBase, uint32(v3178)+24))
	v3192 = F_create_append_path(m, v838, v839, int32(0), v3186, v3188, v3188, v3190, int32(1), v1080)
	mBase = m.M
	v3193 = m.ExcPending
	if v3193 != 0 {
		goto L30
	} else {
		goto L866
	}
L866:
	;
	F_add_partial_path(m, v839, v3192)
	mBase = m.M
	v3195 = m.ExcPending
	if v3195 != 0 {
		goto L30
	} else {
		goto L867
	}
L867:
	;
	goto L864
L868:
	;
	goto L861
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
	v3 = *(*int32)(unsafe.Add(mBase, _consts[7]))
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
				F_errmsg(m, int32(471921), int32(0))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(504839), int32(700), int32(434852))
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
		*(*int32)(unsafe.Add(mBase, _consts[7])) = v26
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
	var v21 int32
	_ = v21
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
	var v49 int32
	_ = v49
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
		goto L12
	} else {
		goto L36
	}
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v12 <= int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if v49 == int32(0) {
		goto L1
	} else {
		goto L16
	}
L4:
	;
	v49 = v6
	goto L3
L5:
	;
	goto L6
L6:
	;
	v20 = v6
	v21 = v6
	goto L7
L7:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v24+v20<<(uint(int32(2))%32))))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	if v29 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v49 = v38
	goto L3
L9:
	;
	v30 = F_copyObjectImpl(m, v29)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v38 = v21
	goto L11
L11:
	;
	v40 = v20 + int32(1)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v40 < v41 {
		v20 = v40
		v21 = v38
		goto L7
	} else {
		goto L15
	}
L12:
	;
	return
L13:
	;
	v32 = F_lappend(m, v21, v30)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+24)))
	v36 = v34 | v35
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v36)
	v38 = v32
	goto L11
L15:
	;
	goto L8
L16:
	;
	if l2 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if v102 == int32(1) {
		goto L30
	} else {
		goto L31
	}
L18:
	;
	v56 = int32(0)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v57 <= v56 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v65 = v56
	goto L20
L20:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v69+v65<<(uint(int32(2))%32))))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+16))
	if v74 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L17
L22:
	;
	v75 = F_copyObjectImpl(m, v74)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L12
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v90 = v65 + int32(1)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v90 < v91 {
		v65 = v90
		goto L20
	} else {
		goto L28
	}
L25:
	;
	F_ChangeVarNodes(m, v75, int32(1), l0)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L12
	} else {
		goto L26
	}
L26:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v81 = F_list_append_unique(m, v80, v75)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L12
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v81
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+24)))
	v86 = v84 | v85
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v86)
	goto L24
L28:
	;
	goto L21
L29:
	;
	F_ChangeVarNodes(m, v111, int32(1), l0)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L12
	} else {
		goto L34
	}
L30:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	v111 = v106
	goto L29
L31:
	;
	goto L32
L32:
	;
	v109 = F_makeBoolExpr(m, int32(1), v49, int32(-1))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L12
	} else {
		goto L33
	}
L33:
	;
	v111 = v109
	goto L29
L34:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v116 = F_list_append_unique(m, v115, v111)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L12
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v116
	return
L36:
	;
	v138 = F_lappend(m, v128, v136)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L12
	} else {
		goto L37
	}
L37:
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
				F_errmsg(m, int32(24882), int32(0))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(506097), int32(1908), int32(84083))
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
						F_errmsg(m, int32(575542), v6)
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(507143), int32(53), int32(282886))
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
				if base.Ui32(int32(10485761)) <= base.Ui32(v17) {
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
							*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = int32(10485760)
							*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = l1
							F_errmsg(m, int32(489346), v6+int32(16))
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(507143), int32(58), int32(282886))
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
					F_errmsg(m, int32(226996), int32(0))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(507143), int32(48), int32(282886))
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
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(409441)
			F_errmsg(m, int32(196564), v4)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(505570), int32(207), int32(285371))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
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
	var v63 int32
	_ = v63
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
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
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
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v85 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v83+v34))) = uint8(v85)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v89 = v87 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v89
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v93 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v91+v89))) = uint8(v93)
	goto L3
L12:
	;
	goto L11
L13:
	;
	v62 = int32(1)
	v63 = v27 + v62
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
	v27 = v63
	goto L10
L27:
	;
	v27 = v63
	goto L10
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
	v14 = *(*int32)(unsafe.Add(mBase, _consts[532]))
	if v14 == int32(0) {
		m.G0 = v11 + int32(192)
		return
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, _consts[535]))
		if v18 == int32(0) {
			v22 = *(*int32)(unsafe.Add(mBase, _consts[536]))
			if v22 == int32(0) {
				F_set_errcontext_domain(m, int32(0))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, _consts[555]))
					v31 = *(*int32)(unsafe.Add(mBase, _consts[532]))
					v32 = F_logicalrep_message_type(m, v31)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v32
						*(*int32)(unsafe.Add(mBase, uint32(v11))) = v29
						F_errcontext_msg(m, int32(728489), v11)
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
				v40 = *(*int64)(unsafe.Add(mBase, _consts[534]))
				F_set_errcontext_domain(m, int32(0))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return
				} else {
					v45 = *(*int32)(unsafe.Add(mBase, _consts[555]))
					v47 = *(*int32)(unsafe.Add(mBase, _consts[532]))
					v48 = F_logicalrep_message_type(m, v47)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return
					} else {
						v51 = *(*int32)(unsafe.Add(mBase, _consts[536]))
						if v40 == int64(0) {
							*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v51
							*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v48
							*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v45
							F_errcontext_msg(m, int32(46534), v11+int32(16))
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return
							} else {
								m.G0 = v11 + int32(192)
								return
							}
						} else {
							v63 = *(*int64)(unsafe.Add(mBase, _consts[534]))
							*(*uint32)(unsafe.Add(mBase, uint32(v11)+48)) = uint32(v63)
							*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v45
							*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v48
							*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v51
							v69 = int64(base.Ui64(v63) >> (uint(int64(32)) % 64))
							*(*uint32)(unsafe.Add(mBase, uint32(v11)+44)) = uint32(v69)
							F_errcontext_msg(m, int32(526322), v11+int32(32))
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
			v77 = *(*int64)(unsafe.Add(mBase, _consts[534]))
			v79 = *(*int32)(unsafe.Add(mBase, _consts[533]))
			F_set_errcontext_domain(m, int32(0))
			mBase = m.M
			v82 = m.ExcPending
			if v82 != 0 {
				return
			} else {
				v84 = *(*int32)(unsafe.Add(mBase, _consts[555]))
				v86 = *(*int32)(unsafe.Add(mBase, _consts[532]))
				v87 = F_logicalrep_message_type(m, v86)
				mBase = m.M
				v88 = m.ExcPending
				if v88 != 0 {
					return
				} else {
					v90 = *(*int32)(unsafe.Add(mBase, _consts[535]))
					v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+8))
					v92 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
					if v79 < int32(0) {
						v96 = *(*int32)(unsafe.Add(mBase, _consts[536]))
						if v77 == int64(0) {
							*(*int32)(unsafe.Add(mBase, uint32(v11)+80)) = v96
							*(*int32)(unsafe.Add(mBase, uint32(v11)+76)) = v91
							*(*int32)(unsafe.Add(mBase, uint32(v11)+72)) = v92
							*(*int32)(unsafe.Add(mBase, uint32(v11)+68)) = v87
							*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = v84
							F_errcontext_msg(m, int32(46254), v11-int32(-64))
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
							v111 = *(*int64)(unsafe.Add(mBase, _consts[534]))
							*(*uint32)(unsafe.Add(mBase, uint32(v11)+120)) = uint32(v111)
							v114 = int64(base.Ui64(v111) >> (uint(int64(32)) % 64))
							*(*uint32)(unsafe.Add(mBase, uint32(v11)+116)) = uint32(v114)
							*(*int32)(unsafe.Add(mBase, uint32(v11)+96)) = v84
							*(*int32)(unsafe.Add(mBase, uint32(v11)+100)) = v87
							*(*int32)(unsafe.Add(mBase, uint32(v11)+104)) = v92
							*(*int32)(unsafe.Add(mBase, uint32(v11)+108)) = v91
							F_errcontext_msg(m, int32(526004), v11+int32(96))
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
						v127 = *(*int32)(unsafe.Add(mBase, _consts[533]))
						v131 = *(*int32)(unsafe.Add(mBase, uint32(v125+v127<<(uint(int32(2))%32))))
						v133 = *(*int32)(unsafe.Add(mBase, _consts[536]))
						if v77 == int64(0) {
							*(*int32)(unsafe.Add(mBase, uint32(v11)+148)) = v133
							*(*int32)(unsafe.Add(mBase, uint32(v11)+144)) = v131
							*(*int32)(unsafe.Add(mBase, uint32(v11)+140)) = v91
							*(*int32)(unsafe.Add(mBase, uint32(v11)+136)) = v92
							*(*int32)(unsafe.Add(mBase, uint32(v11)+132)) = v87
							*(*int32)(unsafe.Add(mBase, uint32(v11)+128)) = v84
							F_errcontext_msg(m, int32(46388), v11+int32(128))
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
							v150 = *(*int64)(unsafe.Add(mBase, _consts[534]))
							*(*uint32)(unsafe.Add(mBase, uint32(v11)+188)) = uint32(v150)
							v153 = int64(base.Ui64(v150) >> (uint(int64(32)) % 64))
							*(*uint32)(unsafe.Add(mBase, uint32(v11)+184)) = uint32(v153)
							*(*int32)(unsafe.Add(mBase, uint32(v11)+160)) = v84
							*(*int32)(unsafe.Add(mBase, uint32(v11)+164)) = v87
							*(*int32)(unsafe.Add(mBase, uint32(v11)+168)) = v92
							*(*int32)(unsafe.Add(mBase, uint32(v11)+172)) = v91
							F_errcontext_msg(m, int32(526157), v11+int32(160))
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
	var v28 int32
	_ = v28
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
	v28 = v3
	goto L7
L7:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v36 = v28 << (uint(int32(2)) % 32)
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
	v157 = v28 + int32(1)
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v157 < v158 {
		v28 = v157
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
	F_errmsg_internal(m, int32(156725), int32(0))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L32
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(503336), int32(824), int32(74879))
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
	F_errmsg_internal(m, int32(74758), int32(0))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L32
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(503336), int32(821), int32(74879))
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
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+68))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v4)+72))
	if v5 < v6 {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v4)+76))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v8+v5<<(uint(int32(2))%32))))
		*(*int32)(unsafe.Add(mBase, uint32(v4)+56)) = v12
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v4)+80))
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14+v5))))
		*(*uint8)(unsafe.Add(mBase, uint32(v4)+60)) = uint8(v16)
		*(*int32)(unsafe.Add(mBase, uint32(v4)+68)) = v5 + int32(1)
		v21 = v4
	} else {
		v21 = int32(0)
	}
	return v21
}
func F_arrayconst_startup_fn(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v16 = F_palloc(m, int32(84))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v16
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
		v23 = F_pg_detoast_datum(m, v22)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return
		} else {
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
			F_get_typlenbyvalalign(m, v25, v13+int32(14), v13+int32(13), v13+int32(12))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return
			} else {
				v35 = int32(*(*int16)(unsafe.Add(mBase, uint32(v13)+14)))
				v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+13)))
				v37 = int32(*(*int8)(unsafe.Add(mBase, uint32(v13)+12)))
				F_deconstruct_array(m, v23, v35, v36, v37, v16+int32(76), v16+int32(80), v16+int32(72))
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v16))) = int32(17)
					v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v48
					v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v51 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v51
					*(*uint8)(unsafe.Add(mBase, uint32(v16)+16)) = uint8(v51)
					*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = int32(16)
					*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v50
					v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v58
					v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					v61 = F_list_copy(m, v60)
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = int32(7)
						*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v61
						v66 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = int32(-1)
						*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v66
						v70 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v70
						v72 = int32(*(*int16)(unsafe.Add(mBase, uint32(v13)+14)))
						*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v72
						v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+13)))
						*(*uint8)(unsafe.Add(mBase, uint32(v16)+61)) = uint8(v74)
						v76 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v76)+4)) = v16 + int32(36)
						*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = int32(0)
						m.G0 = v13 + int32(16)
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
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, _consts[249]))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
		if v14 == int32(1) {
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+1)))
			if base.Ui32((v17-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v44 = int32(1)
				if v14&v44 != 0 {
					v48 = v44
				} else {
					v48 = int32(4)
				}
				v49 = v7 + v48
				if v13 != int32(6) {
					if base.Ui32(v13) <= base.Ui32(int32(41)) {
						v100 = *(*int32)(unsafe.Add(mBase, uint32(v13*int32(28))+uint32(_consts[1056])))
						v101 = v100
					} else {
						v101 = int32(1)
					}
					v102 = int32(*(*int8)(unsafe.Add(mBase, uint32(v49))))
					v104 = v102 & int32(255)
					if v101 < int32(2) {
						v126 = v104
						return v126
					} else {
						if int32(0) <= v102 {
							v126 = v104
							return v126
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v112 = m.ExcPending
							if v112 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(261))
								mBase = m.M
								v115 = m.ExcPending
								if v115 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(409257), int32(0))
									mBase = m.M
									v119 = m.ExcPending
									if v119 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(504517), int32(1001), int32(326738))
										mBase = m.M
										v124 = m.ExcPending
										if v124 != 0 {
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
					v52 = int32(*(*int8)(unsafe.Add(mBase, uint32(v49))))
					if int32(0) <= v52 {
						if base.Ui32(v13) <= base.Ui32(int32(41)) {
							v100 = *(*int32)(unsafe.Add(mBase, uint32(v13*int32(28))+uint32(_consts[1056])))
							v101 = v100
						} else {
							v101 = int32(1)
						}
						v102 = int32(*(*int8)(unsafe.Add(mBase, uint32(v49))))
						v104 = v102 & int32(255)
						if v101 < int32(2) {
							v126 = v104
							return v126
						} else {
							if int32(0) <= v102 {
								v126 = v104
								return v126
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v112 = m.ExcPending
								if v112 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(261))
									mBase = m.M
									v115 = m.ExcPending
									if v115 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(409257), int32(0))
										mBase = m.M
										v119 = m.ExcPending
										if v119 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(504517), int32(1001), int32(326738))
											mBase = m.M
											v124 = m.ExcPending
											if v124 != 0 {
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
						v59 = base.B2i32(base.Ui32(int32(-33)) < base.Ui32(v52))
						if base.Ui32(int32(-33)) < base.Ui32(v52) {
							v60 = int32(2)
						} else {
							v60 = int32(1)
						}
						v62 = base.B2i32(base.Ui32(int32(-17)) < base.Ui32(v52))
						if base.Ui32(int32(-17)) < base.Ui32(v52) {
							v63 = int32(3)
						} else {
							v63 = v60
						}
						if base.Ui32(int32(-33)) < base.Ui32(v52) {
							v69 = int32(15)
						} else {
							v69 = int32(31)
						}
						if base.Ui32(int32(-17)) < base.Ui32(v52) {
							v70 = int32(7)
						} else {
							v70 = v69
						}
						v74 = int32(1)
						v75 = v52 & int32(255) & v70
						v76 = int32(0)
						for {
							v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74+v49))))
							v85 = v80&int32(63) | v75<<(uint(int32(6))%32)
							v86 = int32(1)
							v89 = v76 + v86
							if v89 != v63 {
								v74 = v74 + v86
								v75 = v85
								v76 = v89
								continue
							} else {
								break
							}
							break
						}
						v126 = v85
						return v126
					}
				}
			} else {
				v40 = base.B2i32(v17 == int32(18)) << (uint(int32(4)) % 32)
				if v40 != 0 {
					v44 = int32(1)
					if v14&v44 != 0 {
						v48 = v44
					} else {
						v48 = int32(4)
					}
					v49 = v7 + v48
					if v13 != int32(6) {
						if base.Ui32(v13) <= base.Ui32(int32(41)) {
							v100 = *(*int32)(unsafe.Add(mBase, uint32(v13*int32(28))+uint32(_consts[1056])))
							v101 = v100
						} else {
							v101 = int32(1)
						}
						v102 = int32(*(*int8)(unsafe.Add(mBase, uint32(v49))))
						v104 = v102 & int32(255)
						if v101 < int32(2) {
							v126 = v104
							return v126
						} else {
							if int32(0) <= v102 {
								v126 = v104
								return v126
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v112 = m.ExcPending
								if v112 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(261))
									mBase = m.M
									v115 = m.ExcPending
									if v115 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(409257), int32(0))
										mBase = m.M
										v119 = m.ExcPending
										if v119 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(504517), int32(1001), int32(326738))
											mBase = m.M
											v124 = m.ExcPending
											if v124 != 0 {
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
						v52 = int32(*(*int8)(unsafe.Add(mBase, uint32(v49))))
						if int32(0) <= v52 {
							if base.Ui32(v13) <= base.Ui32(int32(41)) {
								v100 = *(*int32)(unsafe.Add(mBase, uint32(v13*int32(28))+uint32(_consts[1056])))
								v101 = v100
							} else {
								v101 = int32(1)
							}
							v102 = int32(*(*int8)(unsafe.Add(mBase, uint32(v49))))
							v104 = v102 & int32(255)
							if v101 < int32(2) {
								v126 = v104
								return v126
							} else {
								if int32(0) <= v102 {
									v126 = v104
									return v126
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v112 = m.ExcPending
									if v112 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(261))
										mBase = m.M
										v115 = m.ExcPending
										if v115 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(409257), int32(0))
											mBase = m.M
											v119 = m.ExcPending
											if v119 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(504517), int32(1001), int32(326738))
												mBase = m.M
												v124 = m.ExcPending
												if v124 != 0 {
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
							v59 = base.B2i32(base.Ui32(int32(-33)) < base.Ui32(v52))
							if base.Ui32(int32(-33)) < base.Ui32(v52) {
								v60 = int32(2)
							} else {
								v60 = int32(1)
							}
							v62 = base.B2i32(base.Ui32(int32(-17)) < base.Ui32(v52))
							if base.Ui32(int32(-17)) < base.Ui32(v52) {
								v63 = int32(3)
							} else {
								v63 = v60
							}
							if base.Ui32(int32(-33)) < base.Ui32(v52) {
								v69 = int32(15)
							} else {
								v69 = int32(31)
							}
							if base.Ui32(int32(-17)) < base.Ui32(v52) {
								v70 = int32(7)
							} else {
								v70 = v69
							}
							v74 = int32(1)
							v75 = v52 & int32(255) & v70
							v76 = int32(0)
							for {
								v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74+v49))))
								v85 = v80&int32(63) | v75<<(uint(int32(6))%32)
								v86 = int32(1)
								v89 = v76 + v86
								if v89 != v63 {
									v74 = v74 + v86
									v75 = v85
									v76 = v89
									continue
								} else {
									break
								}
								break
							}
							v126 = v85
							return v126
						}
					}
				} else {
					return int32(0)
				}
			}
		} else {
			v28 = int32(1)
			if v14&v28 != 0 {
				v40 = int32(base.Ui32(v14)>>(uint(v28)%32)) - v28
			} else {
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
				v40 = int32(base.Ui32(v34)>>(uint(int32(2))%32)) - int32(4)
			}
			if v40 != 0 {
				v44 = int32(1)
				if v14&v44 != 0 {
					v48 = v44
				} else {
					v48 = int32(4)
				}
				v49 = v7 + v48
				if v13 != int32(6) {
					if base.Ui32(v13) <= base.Ui32(int32(41)) {
						v100 = *(*int32)(unsafe.Add(mBase, uint32(v13*int32(28))+uint32(_consts[1056])))
						v101 = v100
					} else {
						v101 = int32(1)
					}
					v102 = int32(*(*int8)(unsafe.Add(mBase, uint32(v49))))
					v104 = v102 & int32(255)
					if v101 < int32(2) {
						v126 = v104
						return v126
					} else {
						if int32(0) <= v102 {
							v126 = v104
							return v126
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v112 = m.ExcPending
							if v112 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(261))
								mBase = m.M
								v115 = m.ExcPending
								if v115 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(409257), int32(0))
									mBase = m.M
									v119 = m.ExcPending
									if v119 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(504517), int32(1001), int32(326738))
										mBase = m.M
										v124 = m.ExcPending
										if v124 != 0 {
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
					v52 = int32(*(*int8)(unsafe.Add(mBase, uint32(v49))))
					if int32(0) <= v52 {
						if base.Ui32(v13) <= base.Ui32(int32(41)) {
							v100 = *(*int32)(unsafe.Add(mBase, uint32(v13*int32(28))+uint32(_consts[1056])))
							v101 = v100
						} else {
							v101 = int32(1)
						}
						v102 = int32(*(*int8)(unsafe.Add(mBase, uint32(v49))))
						v104 = v102 & int32(255)
						if v101 < int32(2) {
							v126 = v104
							return v126
						} else {
							if int32(0) <= v102 {
								v126 = v104
								return v126
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v112 = m.ExcPending
								if v112 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(261))
									mBase = m.M
									v115 = m.ExcPending
									if v115 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(409257), int32(0))
										mBase = m.M
										v119 = m.ExcPending
										if v119 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(504517), int32(1001), int32(326738))
											mBase = m.M
											v124 = m.ExcPending
											if v124 != 0 {
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
						v59 = base.B2i32(base.Ui32(int32(-33)) < base.Ui32(v52))
						if base.Ui32(int32(-33)) < base.Ui32(v52) {
							v60 = int32(2)
						} else {
							v60 = int32(1)
						}
						v62 = base.B2i32(base.Ui32(int32(-17)) < base.Ui32(v52))
						if base.Ui32(int32(-17)) < base.Ui32(v52) {
							v63 = int32(3)
						} else {
							v63 = v60
						}
						if base.Ui32(int32(-33)) < base.Ui32(v52) {
							v69 = int32(15)
						} else {
							v69 = int32(31)
						}
						if base.Ui32(int32(-17)) < base.Ui32(v52) {
							v70 = int32(7)
						} else {
							v70 = v69
						}
						v74 = int32(1)
						v75 = v52 & int32(255) & v70
						v76 = int32(0)
						for {
							v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74+v49))))
							v85 = v80&int32(63) | v75<<(uint(int32(6))%32)
							v86 = int32(1)
							v89 = v76 + v86
							if v89 != v63 {
								v74 = v74 + v86
								v75 = v85
								v76 = v89
								continue
							} else {
								break
							}
							break
						}
						v126 = v85
						return v126
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
	v5 = *(*int32)(unsafe.Add(mBase, _consts[334]))
	if l0 < v5 {
		v7 = l0
	} else {
		v7 = v5
	}
	*(*int32)(unsafe.Add(mBase, _consts[335])) = v7
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
