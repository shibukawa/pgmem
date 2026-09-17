package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CacheInvalidateRelSync(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
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
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	v8 = F_PrepareInvalidationState(m)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_CacheInvalidateRelSync[0]))
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_CacheInvalidateRelSync[1]))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	if v14 < v15 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	return
L4:
	;
	v19 = v14
	goto L7
L5:
	;
	goto L6
L6:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_CacheInvalidateRelSync[2]))
	if v47 <= v15 {
		goto L14
	} else {
		goto L15
	}
L7:
	;
	v26 = v11 + v19<<(uint(int32(4))%32)
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	if v27 == int32(250) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L6
L9:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
	if base.B2i32(v30 == l0)|base.B2i32(v30 == int32(0)) != 0 {
		goto L3
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v37 = v19 + int32(1)
	if v37 != v15 {
		v19 = v37
		goto L7
	} else {
		goto L13
	}
L12:
	;
	goto L11
L13:
	;
	goto L8
L14:
	;
	if v11 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	v69 = v11
	goto L16
L16:
	;
	v73 = v69 + v15<<(uint(int32(4))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v73)+8)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v73)+4)) = v13
	v76 = int32(250)
	*(*uint8)(unsafe.Add(mBase, uint32(v73))) = uint8(v76)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v78 + int32(1)
	goto L3
L17:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CacheInvalidateRelSync[2])) = v63
	*(*int32)(unsafe.Add(mBase, _c_F_CacheInvalidateRelSync[0])) = v64
	v69 = v64
	goto L16
L18:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_CacheInvalidateRelSync[3]))
	v55 = F_MemoryContextAlloc(m, v53, int32(512))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v61 = F_repalloc(m, v11, v47<<(uint(int32(5))%32))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L22
	}
L21:
	;
	v63 = int32(32)
	v64 = v55
	goto L17
L22:
	;
	v63 = v47 << (uint(int32(1)) % 32)
	v64 = v61
	goto L17
}
func F_CacheInvalidateRelcacheByRelid(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
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
	var v42 int32
	_ = v42
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = F_SearchSysCache1(m, int32(57), l0)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		if v11 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
				F_errmsg_internal(m, int32(_a_F_CacheInvalidateRelcacheByRelid_0), v8)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_CacheInvalidateRelcacheByRelid_1), int32(1697), int32(_a_F_CacheInvalidateRelcacheByRelid_2))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+22)))
			v30 = v28 + v29
			v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+117)))
			v33 = *(*int32)(unsafe.Add(mBase, _c_F_CacheInvalidateRelcacheByRelid[0]))
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
			v35 = F_PrepareInvalidationState(m)
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return
			} else {
				if v31 != 0 {
					v38 = int32(0)
				} else {
					v38 = v33
				}
				F_RegisterRelcacheInvalidation(m, v35, v38, v34)
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return
				} else {
					F_ReleaseCatCache(m, v11)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
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
func F_CacheRegisterRelcacheCallback(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_CacheRegisterRelcacheCallback[0]))
	if int32(10) <= v5 {
		F_errstart_cold(m, int32(22), int32(0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(_a_F_CacheRegisterRelcacheCallback_0), int32(0))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_CacheRegisterRelcacheCallback_1), int32(1862), int32(_a_F_CacheRegisterRelcacheCallback_2))
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
		v22 = v5 << (uint(int32(3)) % 32)
		*(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_CacheRegisterRelcacheCallback[1]))) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_CacheRegisterRelcacheCallback[2]))) = l0
		*(*int32)(unsafe.Add(mBase, _c_F_CacheRegisterRelcacheCallback[0])) = v5 + int32(1)
		return
	}
}
func F_CacheRegisterSyscacheCallback(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v53 int32
	_ = v53
	var v62 int32
	_ = v62
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	v1 = l0
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if base.Ui32(v1) < base.Ui32(int32(85)) {
		v15 = *(*int32)(unsafe.Add(mBase, _c_F_CacheRegisterSyscacheCallback[0]))
		if int32(64) <= v15 {
			F_errstart_cold(m, int32(22), int32(0))
			mBase = m.M
			v90 = m.ExcPending
			if v90 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(_a_F_CacheRegisterSyscacheCallback_0), int32(0))
				mBase = m.M
				v94 = m.ExcPending
				if v94 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_CacheRegisterSyscacheCallback_1), int32(1823), int32(_a_F_CacheRegisterSyscacheCallback_2))
					mBase = m.M
					v99 = m.ExcPending
					if v99 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v19 = v1 << (uint(int32(1)) % 32)
			v20 = int32(*(*int16)(unsafe.Add(mBase, uint32(v19)+uint32(_c_F_CacheRegisterSyscacheCallback[1]))))
			if v20 == int32(0) {
				v26 = v15 + int32(1)
				*(*uint16)(unsafe.Add(mBase, uint32(v19)+uint32(_c_F_CacheRegisterSyscacheCallback[1]))) = uint16(v26)
			} else {
				v31 = v20
				for {
					v36 = v31 * int32(12)
					v39 = int32(*(*int16)(unsafe.Add(mBase, uint32(v36)+uint32(_c_F_CacheRegisterSyscacheCallback[2]))))
					if int32(0) < v39 {
						v31 = v39
						continue
					} else {
						break
					}
					break
				}
				v43 = v15 + int32(1)
				*(*uint16)(unsafe.Add(mBase, uint32(v36)+uint32(_c_F_CacheRegisterSyscacheCallback[2]))) = uint16(v43)
			}
			v53 = v15 * int32(12)
			*(*int32)(unsafe.Add(mBase, uint32(v53)+uint32(_c_F_CacheRegisterSyscacheCallback[3]))) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v53)+uint32(_c_F_CacheRegisterSyscacheCallback[4]))) = l1
			v62 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(v53)+uint32(_c_F_CacheRegisterSyscacheCallback[5]))) = uint16(v62)
			*(*uint16)(unsafe.Add(mBase, uint32(v53)+uint32(_c_F_CacheRegisterSyscacheCallback[6]))) = uint16(v1)
			*(*int32)(unsafe.Add(mBase, _c_F_CacheRegisterSyscacheCallback[0])) = v15 + int32(1)
			m.G0 = v10 + int32(16)
			return
		}
	} else {
		F_errstart_cold(m, int32(22), int32(0))
		mBase = m.M
		v77 = m.ExcPending
		if v77 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v10))) = v1
			F_errmsg_internal(m, int32(_a_F_CacheRegisterSyscacheCallback_3), v10)
			mBase = m.M
			v81 = m.ExcPending
			if v81 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_CacheRegisterSyscacheCallback_1), int32(1821), int32(_a_F_CacheRegisterSyscacheCallback_2))
				mBase = m.M
				v86 = m.ExcPending
				if v86 != 0 {
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
func F_cache_store_tuple(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int64
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v48 int64
	_ = v48
	var v49 int64
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
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
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
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
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
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
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v259 int32
	_ = v259
	var v272 int32
	_ = v272
	v12 = int32(_a_F_cache_store_tuple_0)
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_cache_store_tuple[0]))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	*(*int32)(unsafe.Add(mBase, _c_F_cache_store_tuple[0])) = v16
	v19 = F_palloc(m, int32(8))
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
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+48))
	v26 = m.T0[v25].(func(*base.Module, int32, int32) int32)(m, l1, int32(0))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v28 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v26
	v31 = *(*int64)(unsafe.Add(mBase, uint32(l0)+160))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+160)) = v31 + base.I64_extend_i32_u(v32+int32(8))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v38 == v28 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = v19
	*(*int32)(unsafe.Add(mBase, _c_F_cache_store_tuple[0])) = v13
	v48 = *(*int64)(unsafe.Add(mBase, uint32(l0)+160))
	v49 = *(*int64)(unsafe.Add(mBase, uint32(l0)+168))
	if base.Ui64(v48) <= base.Ui64(v49) {
		v272 = int32(1)
		goto L8
	} else {
		goto L9
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v19
	goto L4
L6:
	;
	goto L7
L7:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+4)) = v19
	goto L4
L8:
	;
	return v272
L9:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v53 = F_cache_reduce_memory(m, l0, v52)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	if v53 == int32(0) {
		v272 = int32(0)
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v57 = int32(1)
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+12)))
	if v58 == v57 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if v61 == v52 {
		v272 = v57
		goto L8
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+8))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+12))
	m.T0[v67].(func(*base.Module, int32))(m, v65)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L16
	}
L15:
	;
	goto L14
L16:
	;
	if v52 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v202 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v65)+4)))
	v204 = v202 & int32(_a_F_cache_store_tuple_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v65)+4)) = uint16(v204)
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v65)+12))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
	*(*uint16)(unsafe.Add(mBase, uint32(v65)+6)) = uint16(v207)
	goto L42
L18:
	;
	v72 = int32(_a_F_cache_store_tuple_0)
	v73 = *(*int32)(unsafe.Add(mBase, _c_F_cache_store_tuple[0]))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_cache_store_tuple[0])) = v76
	if v64 <= int32(0) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	v173 = F_ExecStoreMinimalTuple(m, v171, v63, int32(0))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L33
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_cache_store_tuple[0])) = v73
	goto L17
L22:
	;
	v80 = int32(0)
	if v64 != int32(1) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v89 = v80
	v91 = int32(0)
	goto L26
L24:
	;
	v135 = v80
	goto L25
L25:
	;
	v146 = v135 << (uint(int32(2)) % 32)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v146+v147)))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v65)+20))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v149)+20))
	v153 = m.T0[v152].(func(*base.Module, int32, int32, int32) int32)(m, v149, v75, v150+v135)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L32
	}
L26:
	;
	v100 = v89 << (uint(int32(2)) % 32)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v100+v101)))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v65)+20))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v103)+20))
	v107 = m.T0[v106].(func(*base.Module, int32, int32, int32) int32)(m, v103, v75, v104+v89)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L28
	}
L27:
	;
	if v64&int32(1) == int32(0) {
		goto L21
	} else {
		goto L31
	}
L28:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v65)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v109+v100))) = v107
	v113 = v89 | int32(1)
	v115 = v113 << (uint(int32(2)) % 32)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v115+v116)))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v65)+20))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v118)+20))
	v122 = m.T0[v121].(func(*base.Module, int32, int32, int32) int32)(m, v118, v75, v119+v113)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v65)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v124+v115))) = v122
	v127 = int32(2)
	v128 = v89 + v127
	v130 = v91 + v127
	if v130 != v64&int32(-2) {
		v89 = v128
		v91 = v130
		goto L26
	} else {
		goto L30
	}
L30:
	;
	goto L27
L31:
	;
	v135 = v128
	goto L25
L32:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v65)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v155+v146))) = v153
	goto L21
L33:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v175)))
	v177 = int32(*(*int16)(unsafe.Add(mBase, uint32(v63)+6)))
	if v177 < v176 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	F_slot_getsomeattrs_int(m, v63, v176)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v182 = v64 << (uint(int32(2)) % 32)
	if v182 != 0 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	goto L36
L38:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v65)+16))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v63)+16))
	base.MemoryCopy(m, v183, v184, v182)
	goto L40
L39:
	;
	goto L40
L40:
	;
	if v64 == int32(0) {
		goto L17
	} else {
		goto L41
	}
L41:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v65)+20))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v63)+20))
	base.MemoryCopy(m, v188, v189, v64)
	goto L17
L42:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v210 = F_MemoizeHash_hash(m, v209)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v209)+20))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v209)+12))
	v214 = v210 & v213
	v217 = v212 + v214<<(uint(int32(4))%32)
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217)+12)))
	if v218 != 0 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = v259
	v272 = int32(1)
	goto L8
L45:
	;
	v220 = v217
	v222 = v212
	v224 = v214
	v227 = v213
	goto L48
L46:
	;
	goto L47
L47:
	;
	v259 = int32(0)
	goto L44
L48:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v220)+8))
	if v230 == v210 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	goto L47
L50:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v220)))
	v233 = F_MemoizeHash_equal(m, v209, v232)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L53
	}
L51:
	;
	v237 = v222
	v238 = v227
	goto L52
L52:
	;
	v241 = v238 & (v224 + int32(1))
	v244 = v237 + v241<<(uint(int32(4))%32)
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244)+12)))
	if v245 != 0 {
		v220 = v244
		v222 = v237
		v224 = v241
		v227 = v238
		goto L48
	} else {
		goto L55
	}
L53:
	;
	if v233 != 0 {
		v259 = v220
		goto L44
	} else {
		goto L54
	}
L54:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v209)+12))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v209)+20))
	v237 = v236
	v238 = v235
	goto L52
L55:
	;
	goto L49
}
