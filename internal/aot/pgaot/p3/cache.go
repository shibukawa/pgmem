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
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
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
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v12 = F_SearchSysCache1(m, int32(57), base.I64_extend_i32_u(l0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		if v12 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
				F_errmsg_internal(m, int32(_a_F_CacheInvalidateRelcacheByRelid_0), v8)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_CacheInvalidateRelcacheByRelid_1), int32(1694), int32(_a_F_CacheInvalidateRelcacheByRelid_2))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+22)))
			v31 = v29 + v30
			v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+117)))
			v34 = *(*int32)(unsafe.Add(mBase, _c_F_CacheInvalidateRelcacheByRelid[0]))
			v35 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
			v36 = F_PrepareInvalidationState(m)
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return
			} else {
				if v32 != 0 {
					v39 = int32(0)
				} else {
					v39 = v34
				}
				F_RegisterRelcacheInvalidation(m, v36, v39, v35)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return
				} else {
					F_ReleaseCatCache(m, v12)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
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
				F_errfinish(m, int32(_a_F_CacheRegisterRelcacheCallback_1), int32(1859), int32(_a_F_CacheRegisterRelcacheCallback_2))
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
		v22 = v5 << (uint(int32(4)) % 32)
		*(*int64)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_CacheRegisterRelcacheCallback[1]))) = int64(0)
		*(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_CacheRegisterRelcacheCallback[2]))) = l0
		*(*int32)(unsafe.Add(mBase, _c_F_CacheRegisterRelcacheCallback[0])) = v5 + int32(1)
		return
	}
}
func F_CacheRegisterSyscacheCallback(m *base.Module, l0 int32, l1 int32, l2 int64) {
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
					F_errfinish(m, int32(_a_F_CacheRegisterSyscacheCallback_1), int32(1820), int32(_a_F_CacheRegisterSyscacheCallback_2))
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
					v36 = v31 << (uint(int32(4)) % 32)
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
			v53 = v15 << (uint(int32(4)) % 32)
			*(*int64)(unsafe.Add(mBase, uint32(v53)+uint32(_c_F_CacheRegisterSyscacheCallback[3]))) = l2
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
				F_errfinish(m, int32(_a_F_CacheRegisterSyscacheCallback_1), int32(1818), int32(_a_F_CacheRegisterSyscacheCallback_2))
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
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int64
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int64
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int64
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v267 int32
	_ = v267
	var v280 int32
	_ = v280
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
		v280 = int32(1)
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
	return v280
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
		v280 = int32(0)
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
		v280 = v57
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
	v210 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v65)+4)))
	v212 = v210 & int32(_a_F_cache_store_tuple_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v65)+4)) = uint16(v212)
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v65)+12))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v214)))
	*(*uint16)(unsafe.Add(mBase, uint32(v65)+6)) = uint16(v215)
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
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	v179 = F_ExecStoreMinimalTuple(m, v177, v63, int32(0))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
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
	v139 = v80
	goto L25
L25:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v149+v139<<(uint(int32(2))%32))))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v65)+20))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v153)+24))
	v157 = m.T0[v156].(func(*base.Module, int32, int32, int32) int64)(m, v153, v75, v154+v139)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L32
	}
L26:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v99+v89<<(uint(int32(2))%32))))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v65)+20))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v103)+24))
	v107 = m.T0[v106].(func(*base.Module, int32, int32, int32) int64)(m, v103, v75, v104+v89)
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
	*(*int64)(unsafe.Add(mBase, uint32(v109+v89<<(uint(int32(3))%32)))) = v107
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v116 = v89 | int32(1)
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v114+v116<<(uint(int32(2))%32))))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v65)+20))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v120)+24))
	v124 = m.T0[v123].(func(*base.Module, int32, int32, int32) int64)(m, v120, v75, v121+v116)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v65)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v126+v116<<(uint(int32(3))%32)))) = v124
	v131 = int32(2)
	v132 = v89 + v131
	v134 = v91 + v131
	if v134 != v64&int32(-2) {
		v89 = v132
		v91 = v134
		goto L26
	} else {
		goto L30
	}
L30:
	;
	goto L27
L31:
	;
	v139 = v132
	goto L25
L32:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v65)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v159+v139<<(uint(int32(3))%32)))) = v157
	goto L21
L33:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v181)))
	v183 = int32(*(*int16)(unsafe.Add(mBase, uint32(v63)+6)))
	if v183 < v182 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v63)+8))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v185)+16))
	m.T0[v186].(func(*base.Module, int32, int32))(m, v63, v182)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v190 = v64 << (uint(int32(3)) % 32)
	if v190 != 0 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	goto L36
L38:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v65)+16))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v63)+16))
	base.MemoryCopy(m, v191, v192, v190)
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
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v65)+20))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v63)+20))
	base.MemoryCopy(m, v196, v197, v64)
	goto L17
L42:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v218 = F_MemoizeHash_hash(m, v217)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v217)+20))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v217)+12))
	v222 = v218 & v221
	v225 = v220 + v222<<(uint(int32(4))%32)
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225)+12)))
	if v226 != 0 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = v267
	v280 = int32(1)
	goto L8
L45:
	;
	v228 = v225
	v230 = v221
	v232 = v220
	v233 = v222
	goto L48
L46:
	;
	goto L47
L47:
	;
	v267 = int32(0)
	goto L44
L48:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v228)+8))
	if v238 == v218 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	goto L47
L50:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v228)))
	v241 = F_MemoizeHash_equal(m, v217, v240)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L53
	}
L51:
	;
	v245 = v230
	v246 = v232
	goto L52
L52:
	;
	v249 = v245 & (v233 + int32(1))
	v252 = v246 + v249<<(uint(int32(4))%32)
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v252)+12)))
	if v253 != 0 {
		v228 = v252
		v230 = v245
		v232 = v246
		v233 = v249
		goto L48
	} else {
		goto L55
	}
L53:
	;
	if v241 != 0 {
		v267 = v228
		goto L44
	} else {
		goto L54
	}
L54:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v217)+20))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v217)+12))
	v245 = v244
	v246 = v243
	goto L52
L55:
	;
	goto L49
}
