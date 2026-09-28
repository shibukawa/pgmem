package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_IdleSessionTimeoutHandler(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	v2 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_IdleSessionTimeoutHandler[0])) = v2
	*(*int32)(unsafe.Add(mBase, _c_F_IdleSessionTimeoutHandler[1])) = v2
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_IdleSessionTimeoutHandler[2]))
	F_SetLatch(m, v8)
	mBase = m.M
	return
}
func F_InitSparseVector(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
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
	var v19 int32
	_ = v19
	v7 = F_mul_size(m, int32(4), l1)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = F_add_size(m, int32(16), v7)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v14 = F_mul_size(m, int32(4), l1)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				v16 = F_add_size(m, v11, v14)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return int32(0)
				} else {
					v18 = F_palloc0(m, v16)
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = l0
						*(*int32)(unsafe.Add(mBase, uint32(v18))) = v16 << (uint(int32(2)) % 32)
						return v18
					}
				}
			}
		}
	}
}
func F_InitializeFastPathLocks(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	v5 = int32(1)
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeFastPathLocks[0]))
	if v8&(v8-v5) != 0 {
		v15 = v5 << (uint(int32(32)-base.I32_clz(v8)) % 32)
	} else {
		v15 = v8
	}
	if base.Ui32(v15) <= base.Ui32(int32(31)) {
		v18 = int32(31)
	} else {
		v18 = v15
	}
	if base.Ui32(int32(_a_F_InitializeFastPathLocks_0)) <= base.Ui32(v15) {
		v23 = int32(1024)
	} else {
		v23 = int32(base.Ui32(v18) >> (uint(int32(4)) % 32))
	}
	*(*int32)(unsafe.Add(mBase, _c_F_InitializeFastPathLocks[1])) = v23
	return
}
func F_InitializeLatchWaitSet(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	v5 = F_CreateWaitEventSet(m, int32(0), int32(2))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_InitializeLatchWaitSet[0])) = v5
		v11 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeLatchWaitSet[1]))
		F_AddWaitEventToSet(m, v5, int32(1), int32(-1), v11)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			v15 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitializeLatchWaitSet[2])))
			if v15 == int32(1) {
				v19 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeLatchWaitSet[0]))
				F_AddWaitEventToSet(m, v19, int32(32), int32(-1), int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					return
				}
			} else {
				return
			}
		}
	}
}
func F_IsAbortedTransactionBlockState(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_IsAbortedTransactionBlockState[0]))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+24))
	return base.B2i32((v3-int32(7))&int32(-9) == int32(0))
}
func F__intbig_out(m *base.Module, l0 int32) int64 {
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14235(m, l0, int32(_a_F__intbig_out_0), int32(44), int32(_a_F__intbig_out_1), int32(_a_F__intbig_out_2), int32(_a_F__intbig_out_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		return v7
	}
}
func F_i4toi2(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v33 int64
	_ = v33
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if base.Ui32(base.I32_wrap_i64(v3)-int32(_a_F_i4toi2_0)) <= base.Ui32(int32(-65537)) {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v11 = F_errsave_start(m, v10)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int64(0)
		} else {
			if v11 == int32(0) {
				v33 = int64(0)
				return v33
			} else {
				F_errcode(m, int32(50331778))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, int32(_a_F_i4toi2_1), int32(0))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int64(0)
					} else {
						F_errsave_finish(m, v10, int32(_a_F_i4toi2_2), int32(384), int32(_a_F_i4toi2_3))
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return int64(0)
						} else {
							return int64(0)
						}
					}
				}
			}
		}
	} else {
		v33 = base.I64_extend16_s(v3)
		return v33
	}
}
func F_iclikesel(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14370(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_icu_language_tag(m *base.Module) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	F_errstart_cold(m, int32(21), int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			F_errmsg(m, int32(_a_F_icu_language_tag_0), int32(0))
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_icu_language_tag_1), int32(1893), int32(_a_F_icu_language_tag_2))
				v18 = m.ExcPending
				if v18 != 0 {
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
func F_inc_lex_level(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v7 = v5 + int32(1)
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v8 == int32(0) {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v7
		return
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		if v7 < v13 {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
			v17 = v15 + int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v17
			v51 = v17
			v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+12))
			*(*int32)(unsafe.Add(mBase, uint32(v54+v51<<(uint(int32(2))%32)))) = int32(0)
			return
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
			v21 = v13 - int32(-64)
			v24 = F_repalloc(m, v19, v21*int32(10))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v24
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
				v32 = F_repalloc(m, v29, v21<<(uint(int32(2))%32))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					*(*int32)(unsafe.Add(mBase, uint32(v34)+12)) = v32
					v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
					v38 = F_repalloc(m, v37, v21)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return
					} else {
						v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
						*(*int32)(unsafe.Add(mBase, uint32(v40)+16)) = v38
						v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
						*(*int32)(unsafe.Add(mBase, uint32(v42))) = v21
						v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
						v45 = int32(1)
						v46 = v44 + v45
						*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v46
						v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
						if v48 != v45 {
						} else {
							v51 = v46
							v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
							v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v54+v51<<(uint(int32(2))%32)))) = int32(0)
						}
						return
					}
				}
			}
		}
	}
}
func F_inetnot(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
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
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		v14 = F_palloc0(m, int32(22))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int64(0)
		} else {
			v16 = int32(4)
			v18 = int32(1)
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
			if v20&v18 != 0 {
				v23 = v18
			} else {
				v23 = v16
			}
			v24 = v9 + v23
			v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
			if v25 == int32(2) {
				v28 = v16
			} else {
				v28 = int32(16)
			}
			v29 = int32(1)
			v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
			if v31&v29 != 0 {
				v34 = v29
			} else {
				v34 = int32(4)
			}
			v36 = int32(2)
			v40 = v28
			for {
				v47 = int32(1)
				v48 = v40 - v47
				v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48+(v24+v36)))))
				v53 = v51 ^ int32(-1)
				*(*uint8)(unsafe.Add(mBase, uint32(v14+v34+v36+v48))) = uint8(v53)
				if base.Ui32(v47) < base.Ui32(v40) {
					v40 = v48
					continue
				} else {
					break
				}
				break
			}
			v57 = int32(1)
			v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
			if v59&v57 != 0 {
				v62 = v57
			} else {
				v62 = int32(4)
			}
			v63 = v14 + v62
			v64 = int32(1)
			v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
			if v66&v64 != 0 {
				v69 = v64
			} else {
				v69 = int32(4)
			}
			v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9+v69)+1)))
			*(*uint8)(unsafe.Add(mBase, uint32(v63)+1)) = uint8(v71)
			v73 = int32(1)
			v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
			if v75&v73 != 0 {
				v78 = v73
			} else {
				v78 = int32(4)
			}
			v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9+v78))))
			*(*uint8)(unsafe.Add(mBase, uint32(v63))) = uint8(v80)
			if v80 == int32(2) {
				v86 = int32(40)
			} else {
				v86 = int32(88)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v14))) = v86
			return base.I64_extend_i32_u(v14)
		}
	}
}
func F_init_dummy_sjinfo(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	v4 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = v4
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(322)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+20)) = v4
	*(*int64)(unsafe.Add(mBase, uint32(l0)+28)) = v4
	*(*int64)(unsafe.Add(mBase, uint32(l0)+36)) = v4
	*(*int32)(unsafe.Add(mBase, uint32(l0)+43)) = int32(0)
	return
}
func F_init_params(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v50 int32
	_ = v50
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
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v364 int32
	_ = v364
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
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
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v400 int32
	_ = v400
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
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v432 int64
	_ = v432
	var v434 int32
	_ = v434
	var v435 int64
	_ = v435
	var v439 int64
	_ = v439
	var v441 int32
	_ = v441
	var v442 int64
	_ = v442
	var v446 int64
	_ = v446
	var v449 int64
	_ = v449
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v470 int32
	_ = v470
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v482 int32
	_ = v482
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int64
	_ = v504
	var v505 int32
	_ = v505
	var v509 int32
	_ = v509
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v606 int64
	_ = v606
	var v607 int32
	_ = v607
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v661 int64
	_ = v661
	var v664 int32
	_ = v664
	var v689 int32
	_ = v689
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v715 int32
	_ = v715
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v728 int64
	_ = v728
	var v733 int64
	_ = v733
	var v738 int32
	_ = v738
	var v741 int64
	_ = v741
	var v742 int32
	_ = v742
	var v752 int64
	_ = v752
	var v755 int64
	_ = v755
	var v757 int64
	_ = v757
	var v760 int64
	_ = v760
	var v761 int64
	_ = v761
	var v764 int32
	_ = v764
	var v767 int64
	_ = v767
	var v768 int32
	_ = v768
	var v779 int64
	_ = v779
	var v781 int64
	_ = v781
	var v782 int32
	_ = v782
	var v784 int64
	_ = v784
	var v787 int64
	_ = v787
	var v788 int64
	_ = v788
	var v793 int64
	_ = v793
	var v794 int64
	_ = v794
	var v797 int64
	_ = v797
	var v798 int64
	_ = v798
	var v800 int32
	_ = v800
	var v801 int64
	_ = v801
	var v802 int32
	_ = v802
	var v803 int64
	_ = v803
	var v805 int32
	_ = v805
	var v807 int32
	_ = v807
	var v812 int32
	_ = v812
	var v814 int64
	_ = v814
	var v815 int64
	_ = v815
	var v817 int64
	_ = v817
	var v819 int64
	_ = v819
	var v820 int32
	_ = v820
	var v824 int32
	_ = v824
	var v837 int32
	_ = v837
	var v840 int32
	_ = v840
	var v841 int64
	_ = v841
	var v842 int64
	_ = v842
	var v849 int32
	_ = v849
	var v854 int32
	_ = v854
	var v858 int32
	_ = v858
	var v861 int32
	_ = v861
	var v862 int64
	_ = v862
	var v863 int64
	_ = v863
	var v870 int32
	_ = v870
	var v875 int32
	_ = v875
	var v879 int32
	_ = v879
	var v882 int32
	_ = v882
	var v883 int64
	_ = v883
	var v884 int64
	_ = v884
	var v889 int32
	_ = v889
	var v894 int32
	_ = v894
	var v898 int32
	_ = v898
	var v901 int32
	_ = v901
	var v902 int64
	_ = v902
	var v903 int64
	_ = v903
	var v910 int32
	_ = v910
	var v915 int32
	_ = v915
	var v919 int32
	_ = v919
	var v922 int32
	_ = v922
	var v923 int64
	_ = v923
	var v924 int64
	_ = v924
	var v931 int32
	_ = v931
	var v936 int32
	_ = v936
	var v940 int32
	_ = v940
	var v943 int32
	_ = v943
	var v944 int64
	_ = v944
	var v950 int32
	_ = v950
	var v955 int32
	_ = v955
	var v959 int32
	_ = v959
	var v962 int32
	_ = v962
	var v963 int64
	_ = v963
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v973 int32
	_ = v973
	var v978 int32
	_ = v978
	var v982 int32
	_ = v982
	var v985 int32
	_ = v985
	var v986 int64
	_ = v986
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v996 int32
	_ = v996
	var v1001 int32
	_ = v1001
	var v1003 int32
	_ = v1003
	v13 = int32(0)
	v25 = m.G0
	v27 = v25 - int32(144)
	m.G0 = v27
	*(*uint8)(unsafe.Add(mBase, uint32(l8))) = uint8(v13)
	*(*int32)(unsafe.Add(mBase, uint32(l9))) = v13
	if l1 != 0 {
		goto L8
	} else {
		goto L9
	}
L1:
	;
	F_errorConflictingDefElem(m, v64, l0)
	mBase = m.M
	v1003 = m.ExcPending
	if v1003 != 0 {
		goto L116
	} else {
		goto L298
	}
L2:
	;
	v725 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	switch v725 - int32(21) {
	case 0:
		goto L193
	default:
		goto L192
	case 2:
		goto L194
	}
L3:
	;
	v699 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v699)
	v715 = v689
	v720 = v694
	v722 = v696
	v723 = v697
	v724 = v698
	goto L2
L4:
	;
	if v648 == int32(0) {
		goto L184
	} else {
		goto L185
	}
L5:
	;
	if l3 != 0 {
		v648 = v622
		v649 = v623
		v654 = v628
		v656 = v630
		v657 = v631
		v658 = v632
		goto L4
	} else {
		goto L181
	}
L6:
	;
	if v597 == int32(0) {
		v622 = v590
		v623 = v591
		v628 = v596
		v630 = v598
		v631 = v599
		v632 = v600
		goto L5
	} else {
		goto L178
	}
L7:
	;
	v575 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l4)+48)) = uint8(v575)
	v590 = v564
	v591 = v565
	v596 = v570
	v597 = v571
	v598 = v572
	v599 = v573
	v600 = v574
	goto L6
L8:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v33 {
		goto L16
	} else {
		goto L17
	}
L9:
	;
	goto L10
L10:
	;
	if l3 == int32(0) {
		v622 = v13
		v623 = v13
		v628 = v13
		v630 = v13
		v631 = v13
		v632 = v13
		goto L5
	} else {
		goto L177
	}
L11:
	;
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v401)+12))
	v539 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v538)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(l4)+48)) = uint8(v539)
	v541 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v541)
	v590 = v536
	v591 = v537
	v596 = v403
	v597 = v404
	v598 = v405
	v599 = v406
	v600 = v407
	goto L6
L12:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l4)+16)) = int64(1)
	v532 = int32(0)
	if v401 != 0 {
		v536 = v532
		v537 = v532
		goto L11
	} else {
		goto L176
	}
L13:
	;
	v504 = F_defGetInt64(m, v400)
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L116
	} else {
		goto L166
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L116
	} else {
		goto L159
	}
L15:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L116
	} else {
		goto L155
	}
L16:
	;
	v50 = v13
	v52 = v13
	v53 = v13
	v54 = v13
	v55 = v13
	v56 = v13
	v57 = v13
	v58 = v13
	v59 = v13
	goto L19
L17:
	;
	v400 = v13
	v401 = v13
	v402 = v13
	v403 = v13
	v404 = v13
	v405 = v13
	v406 = v13
	v407 = v13
	goto L18
L18:
	;
	if l3 != 0 {
		goto L134
	} else {
		goto L135
	}
L19:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v60+v50<<(uint(int32(2))%32))))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	if v66 != int32(97) {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	v400 = v372
	v401 = v373
	v402 = v374
	v403 = v375
	v404 = v376
	v405 = v377
	v406 = v378
	v407 = v379
	goto L18
L21:
	;
	v381 = v50 + int32(1)
	v382 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v381 < v382 {
		v50 = v381
		v52 = v372
		v53 = v373
		v54 = v374
		v55 = v375
		v56 = v376
		v57 = v377
		v58 = v378
		v59 = v379
		goto L19
	} else {
		goto L129
	}
L22:
	;
	v370 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l8))) = uint8(v370)
	v372 = v52
	v373 = v53
	v374 = v64
	v375 = v55
	v376 = v56
	v377 = v57
	v378 = v58
	v379 = v59
	goto L21
L23:
	;
	v75 = int32(_a_F_init_params_0)
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	v81 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_init_params[0])))
	if base.B2i32(v78 == int32(0))|base.B2i32(v78 != v81) != 0 {
		v99 = v78
		v100 = v81
		goto L29
	} else {
		goto L30
	}
L24:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+1)))
	if v69 != int32(115) {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+2)))
	if v72 != 0 {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	if v54 == int32(0) {
		goto L22
	} else {
		goto L27
	}
L27:
	;
	goto L1
L28:
	;
	if v99-v100 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L29:
	;
	goto L28
L30:
	;
	v84 = v65
	v85 = v75
	goto L31
L31:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+1)))
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+1)))
	if v89 == int32(0) {
		v99 = v89
		v100 = v88
		goto L29
	} else {
		goto L33
	}
L32:
	;
	v99 = v89
	v100 = v88
	goto L29
L33:
	;
	v92 = int32(1)
	if v89 == v88 {
		v84 = v84 + v92
		v85 = v85 + v92
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	if v52 != 0 {
		goto L1
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v106 = int32(_a_F_init_params_1)
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	v112 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_init_params[1])))
	if base.B2i32(v109 == int32(0))|base.B2i32(v109 != v112) != 0 {
		v130 = v109
		v131 = v112
		goto L40
	} else {
		goto L41
	}
L38:
	;
	v104 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l8))) = uint8(v104)
	v372 = v64
	v373 = v53
	v374 = v54
	v375 = v55
	v376 = v56
	v377 = v57
	v378 = v58
	v379 = v59
	goto L21
L39:
	;
	if v130-v131 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L40:
	;
	goto L39
L41:
	;
	v115 = v65
	v116 = v106
	goto L42
L42:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+1)))
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+1)))
	if v120 == int32(0) {
		v130 = v120
		v131 = v119
		goto L40
	} else {
		goto L44
	}
L43:
	;
	v130 = v120
	v131 = v119
	goto L40
L44:
	;
	v123 = int32(1)
	if v120 == v119 {
		v115 = v115 + v123
		v116 = v116 + v123
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	if v58 != 0 {
		goto L1
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v137 = int32(_a_F_init_params_2)
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	v143 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_init_params[2])))
	if base.B2i32(v140 == int32(0))|base.B2i32(v140 != v143) != 0 {
		v161 = v140
		v162 = v143
		goto L51
	} else {
		goto L52
	}
L49:
	;
	v135 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l8))) = uint8(v135)
	v372 = v52
	v373 = v53
	v374 = v54
	v375 = v55
	v376 = v56
	v377 = v57
	v378 = v64
	v379 = v59
	goto L21
L50:
	;
	if v161-v162 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L51:
	;
	goto L50
L52:
	;
	v146 = v65
	v147 = v137
	goto L53
L53:
	;
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147)+1)))
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+1)))
	if v151 == int32(0) {
		v161 = v151
		v162 = v150
		goto L51
	} else {
		goto L55
	}
L54:
	;
	v161 = v151
	v162 = v150
	goto L51
L55:
	;
	v154 = int32(1)
	if v151 == v150 {
		v146 = v146 + v154
		v147 = v147 + v154
		goto L53
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	if v55 != 0 {
		goto L1
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v168 = int32(_a_F_init_params_3)
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	v174 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_init_params[3])))
	if base.B2i32(v171 == int32(0))|base.B2i32(v171 != v174) != 0 {
		v192 = v171
		v193 = v174
		goto L62
	} else {
		goto L63
	}
L60:
	;
	v166 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l8))) = uint8(v166)
	v372 = v52
	v373 = v53
	v374 = v54
	v375 = v64
	v376 = v56
	v377 = v57
	v378 = v58
	v379 = v59
	goto L21
L61:
	;
	if v192-v193 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L62:
	;
	goto L61
L63:
	;
	v177 = v65
	v178 = v168
	goto L64
L64:
	;
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+1)))
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+1)))
	if v182 == int32(0) {
		v192 = v182
		v193 = v181
		goto L62
	} else {
		goto L66
	}
L65:
	;
	v192 = v182
	v193 = v181
	goto L62
L66:
	;
	v185 = int32(1)
	if v182 == v181 {
		v177 = v177 + v185
		v178 = v178 + v185
		goto L64
	} else {
		goto L67
	}
L67:
	;
	goto L65
L68:
	;
	if v56 != 0 {
		goto L1
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v199 = int32(_a_F_init_params_4)
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	v205 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_init_params[4])))
	if base.B2i32(v202 == int32(0))|base.B2i32(v202 != v205) != 0 {
		v223 = v202
		v224 = v205
		goto L73
	} else {
		goto L74
	}
L71:
	;
	v197 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l8))) = uint8(v197)
	v372 = v52
	v373 = v53
	v374 = v54
	v375 = v55
	v376 = v64
	v377 = v57
	v378 = v58
	v379 = v59
	goto L21
L72:
	;
	if v223-v224 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L73:
	;
	goto L72
L74:
	;
	v208 = v65
	v209 = v199
	goto L75
L75:
	;
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+1)))
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+1)))
	if v213 == int32(0) {
		v223 = v213
		v224 = v212
		goto L73
	} else {
		goto L77
	}
L76:
	;
	v223 = v213
	v224 = v212
	goto L73
L77:
	;
	v216 = int32(1)
	if v213 == v212 {
		v208 = v208 + v216
		v209 = v209 + v216
		goto L75
	} else {
		goto L78
	}
L78:
	;
	goto L76
L79:
	;
	if v57 != 0 {
		goto L1
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v230 = int32(_a_F_init_params_5)
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	v236 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_init_params[5])))
	if base.B2i32(v233 == int32(0))|base.B2i32(v233 != v236) != 0 {
		v254 = v233
		v255 = v236
		goto L84
	} else {
		goto L85
	}
L82:
	;
	v228 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l8))) = uint8(v228)
	v372 = v52
	v373 = v53
	v374 = v54
	v375 = v55
	v376 = v56
	v377 = v64
	v378 = v58
	v379 = v59
	goto L21
L83:
	;
	if v254-v255 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L84:
	;
	goto L83
L85:
	;
	v239 = v65
	v240 = v230
	goto L86
L86:
	;
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240)+1)))
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239)+1)))
	if v244 == int32(0) {
		v254 = v244
		v255 = v243
		goto L84
	} else {
		goto L88
	}
L87:
	;
	v254 = v244
	v255 = v243
	goto L84
L88:
	;
	v247 = int32(1)
	if v244 == v243 {
		v239 = v239 + v247
		v240 = v240 + v247
		goto L86
	} else {
		goto L89
	}
L89:
	;
	goto L87
L90:
	;
	if v59 != 0 {
		goto L1
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v261 = int32(_a_F_init_params_6)
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	v267 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_init_params[6])))
	if base.B2i32(v264 == int32(0))|base.B2i32(v264 != v267) != 0 {
		v285 = v264
		v286 = v267
		goto L95
	} else {
		goto L96
	}
L93:
	;
	v259 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l8))) = uint8(v259)
	v372 = v52
	v373 = v53
	v374 = v54
	v375 = v55
	v376 = v56
	v377 = v57
	v378 = v58
	v379 = v64
	goto L21
L94:
	;
	if v285-v286 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L95:
	;
	goto L94
L96:
	;
	v270 = v65
	v271 = v261
	goto L97
L97:
	;
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271)+1)))
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270)+1)))
	if v275 == int32(0) {
		v285 = v275
		v286 = v274
		goto L95
	} else {
		goto L99
	}
L98:
	;
	v285 = v275
	v286 = v274
	goto L95
L99:
	;
	v278 = int32(1)
	if v275 == v274 {
		v270 = v270 + v278
		v271 = v271 + v278
		goto L97
	} else {
		goto L100
	}
L100:
	;
	goto L98
L101:
	;
	if v53 != 0 {
		goto L1
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	v292 = int32(_a_F_init_params_7)
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	v298 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_init_params[7])))
	if base.B2i32(v295 == int32(0))|base.B2i32(v295 != v298) != 0 {
		v316 = v295
		v317 = v298
		goto L106
	} else {
		goto L107
	}
L104:
	;
	v290 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l8))) = uint8(v290)
	v372 = v52
	v373 = v64
	v374 = v54
	v375 = v55
	v376 = v56
	v377 = v57
	v378 = v58
	v379 = v59
	goto L21
L105:
	;
	if v316-v317 == int32(0) {
		goto L112
	} else {
		goto L113
	}
L106:
	;
	goto L105
L107:
	;
	v301 = v65
	v302 = v292
	goto L108
L108:
	;
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v302)+1)))
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301)+1)))
	if v306 == int32(0) {
		v316 = v306
		v317 = v305
		goto L106
	} else {
		goto L110
	}
L109:
	;
	v316 = v306
	v317 = v305
	goto L106
L110:
	;
	v309 = int32(1)
	if v306 == v305 {
		v301 = v301 + v309
		v302 = v302 + v309
		goto L108
	} else {
		goto L111
	}
L111:
	;
	goto L109
L112:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(l9)))
	if v321 != 0 {
		goto L1
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	v325 = int32(_a_F_init_params_8)
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	v331 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_init_params[8])))
	if base.B2i32(v328 == int32(0))|base.B2i32(v328 != v331) != 0 {
		v349 = v328
		v350 = v331
		goto L119
	} else {
		goto L120
	}
L115:
	;
	v322 = F_defGetQualifiedName(m, v64)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	return
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l9))) = v322
	v372 = v52
	v373 = v53
	v374 = v54
	v375 = v55
	v376 = v56
	v377 = v57
	v378 = v58
	v379 = v59
	goto L21
L118:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L116
	} else {
		goto L125
	}
L119:
	;
	goto L118
L120:
	;
	v334 = v65
	v335 = v325
	goto L121
L121:
	;
	v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335)+1)))
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v334)+1)))
	if v339 == int32(0) {
		v349 = v339
		v350 = v338
		goto L119
	} else {
		goto L123
	}
L122:
	;
	v349 = v339
	v350 = v338
	goto L119
L123:
	;
	v342 = int32(1)
	if v339 == v338 {
		v334 = v334 + v342
		v335 = v335 + v342
		goto L121
	} else {
		goto L124
	}
L124:
	;
	goto L122
L125:
	;
	if v349-v350 == int32(0) {
		goto L15
	} else {
		goto L126
	}
L126:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+128)) = v358
	F_errmsg_internal(m, int32(_a_F_init_params_9), v27+int32(128))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L116
	} else {
		goto L127
	}
L127:
	;
	F_errfinish(m, int32(_a_F_init_params_10), int32(1367), int32(_a_F_init_params_11))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L116
	} else {
		goto L128
	}
L128:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L129:
	;
	goto L20
L130:
	;
	if v400 != 0 {
		v502 = v464
		v503 = v465
		goto L13
	} else {
		goto L153
	}
L131:
	;
	v461 = int32(0)
	v464 = v461
	v465 = v461
	goto L130
L132:
	;
	v459 = int32(0)
	v502 = v459
	v503 = v459
	goto L13
L133:
	;
	v414 = F_defGetTypeName(m, v402)
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L116
	} else {
		goto L140
	}
L134:
	;
	v408 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v408)
	if v402 != 0 {
		goto L133
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	if v402 == int32(0) {
		goto L131
	} else {
		goto L139
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = int32(20)
	if v400 != 0 {
		goto L132
	} else {
		goto L138
	}
L138:
	;
	goto L12
L139:
	;
	goto L133
L140:
	;
	v416 = F_typenameTypeId(m, l0, v414)
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L116
	} else {
		goto L141
	}
L141:
	;
	if base.B2i32(base.Ui32(int32(2)) <= base.Ui32(v416-int32(20)))&base.B2i32(v416 != int32(23)) != 0 {
		goto L14
	} else {
		goto L142
	}
L142:
	;
	if l3 == int32(0) {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v427 = int32(0)
	v429 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	switch v429 - int32(20) {
	case 0:
		goto L147
	case 1:
		goto L149
	default:
		v452 = v427
		v453 = v427
		goto L146
	case 3:
		goto L148
	}
L144:
	;
	goto L145
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = v416
	if v400 == int32(0) {
		goto L12
	} else {
		goto L152
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = v416
	v464 = v452
	v465 = v453
	goto L130
L147:
	;
	v446 = *(*int64)(unsafe.Add(mBase, uint32(l4)+32))
	v449 = *(*int64)(unsafe.Add(mBase, uint32(l4)+24))
	v452 = base.B2i32(v449 == int64(9223372036854775807))
	v453 = base.B2i32(v446 == int64(-9223372036854775807-1))
	goto L146
L148:
	;
	v439 = *(*int64)(unsafe.Add(mBase, uint32(l4)+24))
	v441 = base.B2i32(v439 == int64(2147483647))
	v442 = *(*int64)(unsafe.Add(mBase, uint32(l4)+32))
	if v442 != int64(-2147483648) {
		v452 = v441
		v453 = v427
		goto L146
	} else {
		goto L151
	}
L149:
	;
	v432 = *(*int64)(unsafe.Add(mBase, uint32(l4)+24))
	v434 = base.B2i32(v432 == int64(32767))
	v435 = *(*int64)(unsafe.Add(mBase, uint32(l4)+32))
	if v435 != int64(-32768) {
		v452 = v434
		v453 = v427
		goto L146
	} else {
		goto L150
	}
L150:
	;
	v452 = v434
	v453 = int32(1)
	goto L146
L151:
	;
	v452 = v441
	v453 = int32(1)
	goto L146
L152:
	;
	goto L132
L153:
	;
	if v401 == int32(0) {
		v590 = v464
		v591 = v465
		v596 = v403
		v597 = v404
		v598 = v405
		v599 = v406
		v600 = v407
		goto L6
	} else {
		goto L154
	}
L154:
	;
	v536 = v464
	v537 = v465
	goto L11
L155:
	;
	F_errmsg(m, int32(_a_F_init_params_12), int32(0))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L116
	} else {
		goto L156
	}
L156:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v64)+20))
	F_parser_errposition(m, l0, v475)
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L116
	} else {
		goto L157
	}
L157:
	;
	F_errfinish(m, int32(_a_F_init_params_10), int32(1363), int32(_a_F_init_params_11))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L116
	} else {
		goto L158
	}
L158:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L159:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L116
	} else {
		goto L160
	}
L160:
	;
	if l2 != 0 {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v492 = int32(_a_F_init_params_13)
	goto L163
L162:
	;
	v492 = int32(_a_F_init_params_14)
	goto L163
L163:
	;
	F_errmsg(m, v492, int32(0))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L116
	} else {
		goto L164
	}
L164:
	;
	F_errfinish(m, int32(_a_F_init_params_10), int32(1389), int32(_a_F_init_params_11))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L116
	} else {
		goto L165
	}
L165:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L166:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l4)+16)) = v504
	if v504 != int64(0) {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v509 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v509)
	if v401 != 0 {
		v536 = v502
		v537 = v503
		goto L11
	} else {
		goto L170
	}
L168:
	;
	goto L169
L169:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L116
	} else {
		goto L172
	}
L170:
	;
	if l3 == int32(0) {
		v590 = v502
		v591 = v503
		v596 = v403
		v597 = v404
		v598 = v405
		v599 = v406
		v600 = v407
		goto L6
	} else {
		goto L171
	}
L171:
	;
	v564 = v502
	v565 = v503
	v570 = v403
	v571 = v404
	v572 = v405
	v573 = v406
	v574 = v407
	goto L7
L172:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L116
	} else {
		goto L173
	}
L173:
	;
	F_errmsg(m, int32(_a_F_init_params_15), int32(0))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L116
	} else {
		goto L174
	}
L174:
	;
	F_errfinish(m, int32(_a_F_init_params_10), int32(1423), int32(_a_F_init_params_11))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L116
	} else {
		goto L175
	}
L175:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L176:
	;
	v564 = v532
	v565 = v532
	v570 = v403
	v571 = v404
	v572 = v405
	v573 = v406
	v574 = v407
	goto L7
L177:
	;
	v545 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v545)
	*(*int64)(unsafe.Add(mBase, uint32(l4)+16)) = int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = int32(20)
	v564 = v13
	v565 = v13
	v570 = v13
	v571 = v13
	v572 = v13
	v573 = v13
	v574 = v13
	goto L7
L178:
	;
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v597)+12))
	if v603 == int32(0) {
		v648 = v590
		v649 = v591
		v654 = v596
		v656 = v598
		v657 = v599
		v658 = v600
		goto L4
	} else {
		goto L179
	}
L179:
	;
	v606 = F_defGetInt64(m, v597)
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L116
	} else {
		goto L180
	}
L180:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l4)+24)) = v606
	v689 = v591
	v694 = v596
	v696 = v598
	v697 = v599
	v698 = v600
	goto L3
L181:
	;
	if v622 == int32(0) {
		v715 = v623
		v720 = v628
		v722 = v630
		v723 = v631
		v724 = v632
		goto L2
	} else {
		goto L182
	}
L182:
	;
	v648 = v622
	v649 = v623
	v654 = v628
	v656 = v630
	v657 = v631
	v658 = v632
	goto L4
L183:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l4)+24)) = int64(-1)
	v689 = v649
	v694 = v654
	v696 = v656
	v697 = v657
	v698 = v658
	goto L3
L184:
	;
	v661 = *(*int64)(unsafe.Add(mBase, uint32(l4)+16))
	if v661 <= int64(0) {
		goto L183
	} else {
		goto L187
	}
L185:
	;
	goto L186
L186:
	;
	v664 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	switch v664 - int32(21) {
	case 0:
		goto L190
	default:
		goto L188
	case 2:
		goto L189
	}
L187:
	;
	goto L186
L188:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l4)+24)) = int64(9223372036854775807)
	v689 = v649
	v694 = v654
	v696 = v656
	v697 = v657
	v698 = v658
	goto L3
L189:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l4)+24)) = int64(2147483647)
	v689 = v649
	v694 = v654
	v696 = v656
	v697 = v657
	v698 = v658
	goto L3
L190:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l4)+24)) = int64(32767)
	v689 = v649
	v694 = v654
	v696 = v656
	v697 = v657
	v698 = v658
	goto L3
L191:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v982 = m.ExcPending
	if v982 != 0 {
		goto L116
	} else {
		goto L293
	}
L192:
	;
	if v722 != 0 {
		goto L200
	} else {
		goto L201
	}
L193:
	;
	v733 = *(*int64)(unsafe.Add(mBase, uint32(l4)+24))
	if base.Ui64(v733-int64(32768)) < base.Ui64(int64(-65536)) {
		goto L191
	} else {
		goto L196
	}
L194:
	;
	v728 = *(*int64)(unsafe.Add(mBase, uint32(l4)+24))
	if base.Ui64(int64(-4294967297)) < base.Ui64(v728-int64(2147483648)) {
		goto L192
	} else {
		goto L195
	}
L195:
	;
	goto L191
L196:
	;
	goto L192
L197:
	;
	v767 = *(*int64)(unsafe.Add(mBase, uint32(l4)+32))
	v768 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	switch v768 - int32(21) {
	case 0:
		goto L220
	default:
		goto L219
	case 2:
		goto L221
	}
L198:
	;
	v764 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v764)
	goto L197
L199:
	;
	if v725 == int32(23) {
		goto L206
	} else {
		goto L207
	}
L200:
	;
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v722)+12))
	if v738 == int32(0) {
		goto L199
	} else {
		goto L203
	}
L201:
	;
	goto L202
L202:
	;
	if l3|v715 != int32(1) {
		goto L197
	} else {
		goto L205
	}
L203:
	;
	v741 = F_defGetInt64(m, v722)
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L116
	} else {
		goto L204
	}
L204:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l4)+32)) = v741
	goto L198
L205:
	;
	goto L199
L206:
	;
	v752 = int64(-2147483648)
	goto L208
L207:
	;
	v752 = int64(-9223372036854775807 - 1)
	goto L208
L208:
	;
	if v725 == int32(21) {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	v755 = int64(-32768)
	goto L211
L210:
	;
	v755 = v752
	goto L211
L211:
	;
	v757 = *(*int64)(unsafe.Add(mBase, uint32(l4)+16))
	if int64(0) <= v757 {
		goto L212
	} else {
		goto L213
	}
L212:
	;
	v760 = int64(1)
	goto L214
L213:
	;
	v760 = v755
	goto L214
L214:
	;
	if v715 != 0 {
		goto L215
	} else {
		goto L216
	}
L215:
	;
	v761 = v755
	goto L217
L216:
	;
	v761 = v760
	goto L217
L217:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l4)+32)) = v761
	goto L198
L218:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		goto L116
	} else {
		goto L288
	}
L219:
	;
	v779 = *(*int64)(unsafe.Add(mBase, uint32(l4)+24))
	if v767 < v779 {
		goto L229
	} else {
		goto L230
	}
L220:
	;
	if base.Ui64(v767-int64(32768)) < base.Ui64(int64(-65536)) {
		goto L218
	} else {
		goto L223
	}
L221:
	;
	if base.Ui64(int64(-4294967297)) < base.Ui64(v767-int64(2147483648)) {
		goto L219
	} else {
		goto L222
	}
L222:
	;
	goto L218
L223:
	;
	goto L219
L224:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v940 = m.ExcPending
	if v940 != 0 {
		goto L116
	} else {
		goto L284
	}
L225:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v919 = m.ExcPending
	if v919 != 0 {
		goto L116
	} else {
		goto L280
	}
L226:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v898 = m.ExcPending
	if v898 != 0 {
		goto L116
	} else {
		goto L276
	}
L227:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
		goto L116
	} else {
		goto L272
	}
L228:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		goto L116
	} else {
		goto L268
	}
L229:
	;
	if v723 != 0 {
		goto L234
	} else {
		goto L235
	}
L230:
	;
	goto L231
L231:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v837 = m.ExcPending
	if v837 != 0 {
		goto L116
	} else {
		goto L264
	}
L232:
	;
	v798 = *(*int64)(unsafe.Add(mBase, uint32(l4)+24))
	if v798 < v797 {
		goto L227
	} else {
		goto L245
	}
L233:
	;
	if v794 < v793 {
		goto L228
	} else {
		goto L244
	}
L234:
	;
	v781 = F_defGetInt64(m, v723)
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L116
	} else {
		goto L237
	}
L235:
	;
	goto L236
L236:
	;
	if l3 == int32(0) {
		goto L238
	} else {
		goto L239
	}
L237:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l4)+8)) = v781
	v784 = *(*int64)(unsafe.Add(mBase, uint32(l4)+32))
	v793 = v784
	v794 = v781
	goto L233
L238:
	;
	v787 = *(*int64)(unsafe.Add(mBase, uint32(l4)+8))
	v793 = v767
	v794 = v787
	goto L233
L239:
	;
	goto L240
L240:
	;
	v788 = *(*int64)(unsafe.Add(mBase, uint32(l4)+16))
	if int64(0) < v788 {
		goto L241
	} else {
		goto L242
	}
L241:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l4)+8)) = v767
	v797 = v767
	goto L232
L242:
	;
	goto L243
L243:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l4)+8)) = v779
	v793 = v767
	v794 = v779
	goto L233
L244:
	;
	v797 = v794
	goto L232
L245:
	;
	if v720 != 0 {
		goto L247
	} else {
		goto L248
	}
L246:
	;
	v814 = *(*int64)(unsafe.Add(mBase, uint32(l5)))
	v815 = *(*int64)(unsafe.Add(mBase, uint32(l4)+32))
	if v814 < v815 {
		goto L226
	} else {
		goto L255
	}
L247:
	;
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v720)+12))
	if v800 != 0 {
		goto L250
	} else {
		goto L251
	}
L248:
	;
	goto L249
L249:
	;
	if l3 == int32(0) {
		goto L246
	} else {
		goto L254
	}
L250:
	;
	v801 = F_defGetInt64(m, v720)
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L116
	} else {
		goto L253
	}
L251:
	;
	v803 = v797
	goto L252
L252:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l5))) = v803
	v805 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v805)
	v807 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v807)
	goto L246
L253:
	;
	v803 = v801
	goto L252
L254:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l5))) = v797
	v812 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v812)
	goto L246
L255:
	;
	v817 = *(*int64)(unsafe.Add(mBase, uint32(l4)+24))
	if v817 < v814 {
		goto L225
	} else {
		goto L256
	}
L256:
	;
	if v724 != 0 {
		goto L258
	} else {
		goto L259
	}
L257:
	;
	m.G0 = v27 + int32(144)
	return
L258:
	;
	v819 = F_defGetInt64(m, v724)
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L116
	} else {
		goto L261
	}
L259:
	;
	goto L260
L260:
	;
	if l3 == int32(0) {
		goto L257
	} else {
		goto L263
	}
L261:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l4)+40)) = v819
	if v819 <= int64(0) {
		goto L224
	} else {
		goto L262
	}
L262:
	;
	v824 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v824)
	goto L257
L263:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l4)+40)) = int64(1)
	goto L257
L264:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v840 = m.ExcPending
	if v840 != 0 {
		goto L116
	} else {
		goto L265
	}
L265:
	;
	v841 = *(*int64)(unsafe.Add(mBase, uint32(l4)+32))
	v842 = *(*int64)(unsafe.Add(mBase, uint32(l4)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+88)) = v842
	*(*int64)(unsafe.Add(mBase, uint32(v27)+80)) = v841
	F_errmsg(m, int32(_a_F_init_params_16), v27+int32(80))
	mBase = m.M
	v849 = m.ExcPending
	if v849 != 0 {
		goto L116
	} else {
		goto L266
	}
L266:
	;
	F_errfinish(m, int32(_a_F_init_params_10), int32(1513), int32(_a_F_init_params_11))
	mBase = m.M
	v854 = m.ExcPending
	if v854 != 0 {
		goto L116
	} else {
		goto L267
	}
L267:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L268:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v861 = m.ExcPending
	if v861 != 0 {
		goto L116
	} else {
		goto L269
	}
L269:
	;
	v862 = *(*int64)(unsafe.Add(mBase, uint32(l4)+8))
	v863 = *(*int64)(unsafe.Add(mBase, uint32(l4)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+72)) = v863
	*(*int64)(unsafe.Add(mBase, uint32(v27)+64)) = v862
	F_errmsg(m, int32(_a_F_init_params_17), v27-int32(-64))
	mBase = m.M
	v870 = m.ExcPending
	if v870 != 0 {
		goto L116
	} else {
		goto L270
	}
L270:
	;
	F_errfinish(m, int32(_a_F_init_params_10), int32(1534), int32(_a_F_init_params_11))
	mBase = m.M
	v875 = m.ExcPending
	if v875 != 0 {
		goto L116
	} else {
		goto L271
	}
L271:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L272:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v882 = m.ExcPending
	if v882 != 0 {
		goto L116
	} else {
		goto L273
	}
L273:
	;
	v883 = *(*int64)(unsafe.Add(mBase, uint32(l4)+8))
	v884 = *(*int64)(unsafe.Add(mBase, uint32(l4)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+8)) = v884
	*(*int64)(unsafe.Add(mBase, uint32(v27))) = v883
	F_errmsg(m, int32(_a_F_init_params_18), v27)
	mBase = m.M
	v889 = m.ExcPending
	if v889 != 0 {
		goto L116
	} else {
		goto L274
	}
L274:
	;
	F_errfinish(m, int32(_a_F_init_params_10), int32(1540), int32(_a_F_init_params_11))
	mBase = m.M
	v894 = m.ExcPending
	if v894 != 0 {
		goto L116
	} else {
		goto L275
	}
L275:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L276:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v901 = m.ExcPending
	if v901 != 0 {
		goto L116
	} else {
		goto L277
	}
L277:
	;
	v902 = *(*int64)(unsafe.Add(mBase, uint32(l5)))
	v903 = *(*int64)(unsafe.Add(mBase, uint32(l4)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+24)) = v903
	*(*int64)(unsafe.Add(mBase, uint32(v27)+16)) = v902
	F_errmsg(m, int32(_a_F_init_params_19), v27+int32(16))
	mBase = m.M
	v910 = m.ExcPending
	if v910 != 0 {
		goto L116
	} else {
		goto L278
	}
L278:
	;
	F_errfinish(m, int32(_a_F_init_params_10), int32(1564), int32(_a_F_init_params_11))
	mBase = m.M
	v915 = m.ExcPending
	if v915 != 0 {
		goto L116
	} else {
		goto L279
	}
L279:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L280:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v922 = m.ExcPending
	if v922 != 0 {
		goto L116
	} else {
		goto L281
	}
L281:
	;
	v923 = *(*int64)(unsafe.Add(mBase, uint32(l5)))
	v924 = *(*int64)(unsafe.Add(mBase, uint32(l4)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+40)) = v924
	*(*int64)(unsafe.Add(mBase, uint32(v27)+32)) = v923
	F_errmsg(m, int32(_a_F_init_params_20), v27+int32(32))
	mBase = m.M
	v931 = m.ExcPending
	if v931 != 0 {
		goto L116
	} else {
		goto L282
	}
L282:
	;
	F_errfinish(m, int32(_a_F_init_params_10), int32(1570), int32(_a_F_init_params_11))
	mBase = m.M
	v936 = m.ExcPending
	if v936 != 0 {
		goto L116
	} else {
		goto L283
	}
L283:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L284:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v943 = m.ExcPending
	if v943 != 0 {
		goto L116
	} else {
		goto L285
	}
L285:
	;
	v944 = *(*int64)(unsafe.Add(mBase, uint32(l4)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+48)) = v944
	F_errmsg(m, int32(_a_F_init_params_21), v27+int32(48))
	mBase = m.M
	v950 = m.ExcPending
	if v950 != 0 {
		goto L116
	} else {
		goto L286
	}
L286:
	;
	F_errfinish(m, int32(_a_F_init_params_10), int32(1580), int32(_a_F_init_params_11))
	mBase = m.M
	v955 = m.ExcPending
	if v955 != 0 {
		goto L116
	} else {
		goto L287
	}
L287:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L288:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v962 = m.ExcPending
	if v962 != 0 {
		goto L116
	} else {
		goto L289
	}
L289:
	;
	v963 = *(*int64)(unsafe.Add(mBase, uint32(l4)+32))
	v964 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	v965 = F_format_type_be(m, v964)
	mBase = m.M
	v966 = m.ExcPending
	if v966 != 0 {
		goto L116
	} else {
		goto L290
	}
L290:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+104)) = v965
	*(*int64)(unsafe.Add(mBase, uint32(v27)+96)) = v963
	F_errmsg(m, int32(_a_F_init_params_22), v27+int32(96))
	mBase = m.M
	v973 = m.ExcPending
	if v973 != 0 {
		goto L116
	} else {
		goto L291
	}
L291:
	;
	F_errfinish(m, int32(_a_F_init_params_10), int32(1505), int32(_a_F_init_params_11))
	mBase = m.M
	v978 = m.ExcPending
	if v978 != 0 {
		goto L116
	} else {
		goto L292
	}
L292:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L293:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v985 = m.ExcPending
	if v985 != 0 {
		goto L116
	} else {
		goto L294
	}
L294:
	;
	v986 = *(*int64)(unsafe.Add(mBase, uint32(l4)+24))
	v987 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	v988 = F_format_type_be(m, v987)
	mBase = m.M
	v989 = m.ExcPending
	if v989 != 0 {
		goto L116
	} else {
		goto L295
	}
L295:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+120)) = v988
	*(*int64)(unsafe.Add(mBase, uint32(v27)+112)) = v986
	F_errmsg(m, int32(_a_F_init_params_23), v27+int32(112))
	mBase = m.M
	v996 = m.ExcPending
	if v996 != 0 {
		goto L116
	} else {
		goto L296
	}
L296:
	;
	F_errfinish(m, int32(_a_F_init_params_10), int32(1473), int32(_a_F_init_params_11))
	mBase = m.M
	v1001 = m.ExcPending
	if v1001 != 0 {
		goto L116
	} else {
		goto L297
	}
L297:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L298:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_init_ps_display(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	if l0 == int32(0) {
		v5 = *(*int32)(unsafe.Add(mBase, _c_F_init_ps_display[0]))
		if base.Ui32(v5) <= base.Ui32(int32(17)) {
		} else {
		}
	} else {
	}
	return
}
func F_init_span(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
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
	var v32 int32
	_ = v32
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
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
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
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	v6 = l5
	if l1 != 0 {
		v13 = int32(0)
		v16 = base.AtomicRmwOr32(m, v13, int32(_a_F_init_span_0), v13)
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+1468))
		if v17 != v19 {
			v24 = F_LWLockAcquire(m, v18+int32(1476), int32(0))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				F_check_for_freed_segments_locked(m, l0)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					F_LWLockRelease(m, v28+int32(1476))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return
					} else {
						v36 = int32(base.Ui32(l1) >> (uint(int32(27)) % 32))
						v39 = l0 + v36*int32(20)
						v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
						if v40 != 0 {
							v44 = v40
							v51 = v44 + l1&int32(134217727)
							v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6<<(uint(int32(1))%32))+uint32(_c_F_init_span[0]))))
							v55 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
							if v55 != 0 {
								v56 = int32(0)
								v59 = base.AtomicRmwOr32(m, v56, int32(_a_F_init_span_0), v56)
								v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
								v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+1468))
								if v60 != v62 {
									v67 = F_LWLockAcquire(m, v61+int32(1476), int32(0))
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return
									} else {
										F_check_for_freed_segments_locked(m, l0)
										mBase = m.M
										v70 = m.ExcPending
										if v70 != 0 {
											return
										} else {
											v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											F_LWLockRelease(m, v71+int32(1476))
											mBase = m.M
											v75 = m.ExcPending
											if v75 != 0 {
												return
											} else {
												v79 = int32(base.Ui32(v55) >> (uint(int32(27)) % 32))
												v82 = l0 + v79*int32(20)
												v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
												if v83 != 0 {
													v87 = v83
													*(*int32)(unsafe.Add(mBase, uint32(v87+v55&int32(134217727))+4)) = l1
													v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
													*(*int32)(unsafe.Add(mBase, uint32(v51))) = l2 - v94
													v97 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
													v98 = int32(0)
													*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = v98
													*(*int32)(unsafe.Add(mBase, uint32(v51)+8)) = v97
													*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = l1
													*(*uint16)(unsafe.Add(mBase, uint32(v51)+22)) = uint16(v98)
													*(*uint16)(unsafe.Add(mBase, uint32(v51)+20)) = uint16(v6)
													*(*int32)(unsafe.Add(mBase, uint32(v51)+16)) = l4
													*(*int32)(unsafe.Add(mBase, uint32(v51)+12)) = l3
													switch v6 {
													case 0:
														v108 = int32(1)
														*(*uint16)(unsafe.Add(mBase, uint32(v51)+22)) = uint16(v108)
														v111 = base.I32_div_u_s(int32(_a_F_init_span_1), v54)
														v116 = v111 - v108
														*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)) = uint16(v116)
														v118 = v116
													case 1:
														v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)))
														v118 = v107
													default:
														v115 = base.I32_div_u_s(int32(_a_F_init_span_2), v54)
														v116 = v115
														*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)) = uint16(v116)
														v118 = v116
													}
													v119 = int32(1)
													*(*uint16)(unsafe.Add(mBase, uint32(v51)+30)) = uint16(v119)
													*(*uint16)(unsafe.Add(mBase, uint32(v51)+28)) = uint16(v118)
													v122 = int32(_a_F_init_span_3)
													*(*uint16)(unsafe.Add(mBase, uint32(v51)+26)) = uint16(v122)
													return
												} else {
													v84 = F_get_segment_by_index(m, l0, v79)
													mBase = m.M
													v85 = m.ExcPending
													if v85 != 0 {
														return
													} else {
														v86 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
														v87 = v86
														*(*int32)(unsafe.Add(mBase, uint32(v87+v55&int32(134217727))+4)) = l1
														v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
														*(*int32)(unsafe.Add(mBase, uint32(v51))) = l2 - v94
														v97 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
														v98 = int32(0)
														*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = v98
														*(*int32)(unsafe.Add(mBase, uint32(v51)+8)) = v97
														*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = l1
														*(*uint16)(unsafe.Add(mBase, uint32(v51)+22)) = uint16(v98)
														*(*uint16)(unsafe.Add(mBase, uint32(v51)+20)) = uint16(v6)
														*(*int32)(unsafe.Add(mBase, uint32(v51)+16)) = l4
														*(*int32)(unsafe.Add(mBase, uint32(v51)+12)) = l3
														switch v6 {
														case 0:
															v108 = int32(1)
															*(*uint16)(unsafe.Add(mBase, uint32(v51)+22)) = uint16(v108)
															v111 = base.I32_div_u_s(int32(_a_F_init_span_1), v54)
															v116 = v111 - v108
															*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)) = uint16(v116)
															v118 = v116
														case 1:
															v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)))
															v118 = v107
														default:
															v115 = base.I32_div_u_s(int32(_a_F_init_span_2), v54)
															v116 = v115
															*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)) = uint16(v116)
															v118 = v116
														}
														v119 = int32(1)
														*(*uint16)(unsafe.Add(mBase, uint32(v51)+30)) = uint16(v119)
														*(*uint16)(unsafe.Add(mBase, uint32(v51)+28)) = uint16(v118)
														v122 = int32(_a_F_init_span_3)
														*(*uint16)(unsafe.Add(mBase, uint32(v51)+26)) = uint16(v122)
														return
													}
												}
											}
										}
									}
								} else {
									v79 = int32(base.Ui32(v55) >> (uint(int32(27)) % 32))
									v82 = l0 + v79*int32(20)
									v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
									if v83 != 0 {
										v87 = v83
										*(*int32)(unsafe.Add(mBase, uint32(v87+v55&int32(134217727))+4)) = l1
										v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										*(*int32)(unsafe.Add(mBase, uint32(v51))) = l2 - v94
										v97 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
										v98 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = v98
										*(*int32)(unsafe.Add(mBase, uint32(v51)+8)) = v97
										*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = l1
										*(*uint16)(unsafe.Add(mBase, uint32(v51)+22)) = uint16(v98)
										*(*uint16)(unsafe.Add(mBase, uint32(v51)+20)) = uint16(v6)
										*(*int32)(unsafe.Add(mBase, uint32(v51)+16)) = l4
										*(*int32)(unsafe.Add(mBase, uint32(v51)+12)) = l3
										switch v6 {
										case 0:
											v108 = int32(1)
											*(*uint16)(unsafe.Add(mBase, uint32(v51)+22)) = uint16(v108)
											v111 = base.I32_div_u_s(int32(_a_F_init_span_1), v54)
											v116 = v111 - v108
											*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)) = uint16(v116)
											v118 = v116
										case 1:
											v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)))
											v118 = v107
										default:
											v115 = base.I32_div_u_s(int32(_a_F_init_span_2), v54)
											v116 = v115
											*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)) = uint16(v116)
											v118 = v116
										}
										v119 = int32(1)
										*(*uint16)(unsafe.Add(mBase, uint32(v51)+30)) = uint16(v119)
										*(*uint16)(unsafe.Add(mBase, uint32(v51)+28)) = uint16(v118)
										v122 = int32(_a_F_init_span_3)
										*(*uint16)(unsafe.Add(mBase, uint32(v51)+26)) = uint16(v122)
										return
									} else {
										v84 = F_get_segment_by_index(m, l0, v79)
										mBase = m.M
										v85 = m.ExcPending
										if v85 != 0 {
											return
										} else {
											v86 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
											v87 = v86
											*(*int32)(unsafe.Add(mBase, uint32(v87+v55&int32(134217727))+4)) = l1
											v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											*(*int32)(unsafe.Add(mBase, uint32(v51))) = l2 - v94
											v97 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
											v98 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = v98
											*(*int32)(unsafe.Add(mBase, uint32(v51)+8)) = v97
											*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = l1
											*(*uint16)(unsafe.Add(mBase, uint32(v51)+22)) = uint16(v98)
											*(*uint16)(unsafe.Add(mBase, uint32(v51)+20)) = uint16(v6)
											*(*int32)(unsafe.Add(mBase, uint32(v51)+16)) = l4
											*(*int32)(unsafe.Add(mBase, uint32(v51)+12)) = l3
											switch v6 {
											case 0:
												v108 = int32(1)
												*(*uint16)(unsafe.Add(mBase, uint32(v51)+22)) = uint16(v108)
												v111 = base.I32_div_u_s(int32(_a_F_init_span_1), v54)
												v116 = v111 - v108
												*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)) = uint16(v116)
												v118 = v116
											case 1:
												v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)))
												v118 = v107
											default:
												v115 = base.I32_div_u_s(int32(_a_F_init_span_2), v54)
												v116 = v115
												*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)) = uint16(v116)
												v118 = v116
											}
											v119 = int32(1)
											*(*uint16)(unsafe.Add(mBase, uint32(v51)+30)) = uint16(v119)
											*(*uint16)(unsafe.Add(mBase, uint32(v51)+28)) = uint16(v118)
											v122 = int32(_a_F_init_span_3)
											*(*uint16)(unsafe.Add(mBase, uint32(v51)+26)) = uint16(v122)
											return
										}
									}
								}
							} else {
								v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v51))) = l2 - v94
								v97 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
								v98 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = v98
								*(*int32)(unsafe.Add(mBase, uint32(v51)+8)) = v97
								*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = l1
								*(*uint16)(unsafe.Add(mBase, uint32(v51)+22)) = uint16(v98)
								*(*uint16)(unsafe.Add(mBase, uint32(v51)+20)) = uint16(v6)
								*(*int32)(unsafe.Add(mBase, uint32(v51)+16)) = l4
								*(*int32)(unsafe.Add(mBase, uint32(v51)+12)) = l3
								switch v6 {
								case 0:
									v108 = int32(1)
									*(*uint16)(unsafe.Add(mBase, uint32(v51)+22)) = uint16(v108)
									v111 = base.I32_div_u_s(int32(_a_F_init_span_1), v54)
									v116 = v111 - v108
									*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)) = uint16(v116)
									v118 = v116
								case 1:
									v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)))
									v118 = v107
								default:
									v115 = base.I32_div_u_s(int32(_a_F_init_span_2), v54)
									v116 = v115
									*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)) = uint16(v116)
									v118 = v116
								}
								v119 = int32(1)
								*(*uint16)(unsafe.Add(mBase, uint32(v51)+30)) = uint16(v119)
								*(*uint16)(unsafe.Add(mBase, uint32(v51)+28)) = uint16(v118)
								v122 = int32(_a_F_init_span_3)
								*(*uint16)(unsafe.Add(mBase, uint32(v51)+26)) = uint16(v122)
								return
							}
						} else {
							v41 = F_get_segment_by_index(m, l0, v36)
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return
							} else {
								v43 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
								v44 = v43
								v51 = v44 + l1&int32(134217727)
								v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6<<(uint(int32(1))%32))+uint32(_c_F_init_span[0]))))
								v55 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
								if v55 != 0 {
									v56 = int32(0)
									v59 = base.AtomicRmwOr32(m, v56, int32(_a_F_init_span_0), v56)
									v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
									v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+1468))
									if v60 != v62 {
										v67 = F_LWLockAcquire(m, v61+int32(1476), int32(0))
										mBase = m.M
										v68 = m.ExcPending
										if v68 != 0 {
											return
										} else {
											F_check_for_freed_segments_locked(m, l0)
											mBase = m.M
											v70 = m.ExcPending
											if v70 != 0 {
												return
											} else {
												v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												F_LWLockRelease(m, v71+int32(1476))
												mBase = m.M
												v75 = m.ExcPending
												if v75 != 0 {
													return
												} else {
													v79 = int32(base.Ui32(v55) >> (uint(int32(27)) % 32))
													v82 = l0 + v79*int32(20)
													v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
													if v83 != 0 {
														v87 = v83
														*(*int32)(unsafe.Add(mBase, uint32(v87+v55&int32(134217727))+4)) = l1
														v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
														*(*int32)(unsafe.Add(mBase, uint32(v51))) = l2 - v94
														v97 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
														v98 = int32(0)
														*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = v98
														*(*int32)(unsafe.Add(mBase, uint32(v51)+8)) = v97
														*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = l1
														*(*uint16)(unsafe.Add(mBase, uint32(v51)+22)) = uint16(v98)
														*(*uint16)(unsafe.Add(mBase, uint32(v51)+20)) = uint16(v6)
														*(*int32)(unsafe.Add(mBase, uint32(v51)+16)) = l4
														*(*int32)(unsafe.Add(mBase, uint32(v51)+12)) = l3
														switch v6 {
														case 0:
															v108 = int32(1)
															*(*uint16)(unsafe.Add(mBase, uint32(v51)+22)) = uint16(v108)
															v111 = base.I32_div_u_s(int32(_a_F_init_span_1), v54)
															v116 = v111 - v108
															*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)) = uint16(v116)
															v118 = v116
														case 1:
															v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)))
															v118 = v107
														default:
															v115 = base.I32_div_u_s(int32(_a_F_init_span_2), v54)
															v116 = v115
															*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)) = uint16(v116)
															v118 = v116
														}
														v119 = int32(1)
														*(*uint16)(unsafe.Add(mBase, uint32(v51)+30)) = uint16(v119)
														*(*uint16)(unsafe.Add(mBase, uint32(v51)+28)) = uint16(v118)
														v122 = int32(_a_F_init_span_3)
														*(*uint16)(unsafe.Add(mBase, uint32(v51)+26)) = uint16(v122)
														return
													} else {
														v84 = F_get_segment_by_index(m, l0, v79)
														mBase = m.M
														v85 = m.ExcPending
														if v85 != 0 {
															return
														} else {
															v86 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
															v87 = v86
															*(*int32)(unsafe.Add(mBase, uint32(v87+v55&int32(134217727))+4)) = l1
															v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
															*(*int32)(unsafe.Add(mBase, uint32(v51))) = l2 - v94
															v97 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
															v98 = int32(0)
															*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = v98
															*(*int32)(unsafe.Add(mBase, uint32(v51)+8)) = v97
															*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = l1
															*(*uint16)(unsafe.Add(mBase, uint32(v51)+22)) = uint16(v98)
															*(*uint16)(unsafe.Add(mBase, uint32(v51)+20)) = uint16(v6)
															*(*int32)(unsafe.Add(mBase, uint32(v51)+16)) = l4
															*(*int32)(unsafe.Add(mBase, uint32(v51)+12)) = l3
															switch v6 {
															case 0:
																v108 = int32(1)
																*(*uint16)(unsafe.Add(mBase, uint32(v51)+22)) = uint16(v108)
																v111 = base.I32_div_u_s(int32(_a_F_init_span_1), v54)
																v116 = v111 - v108
																*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)) = uint16(v116)
																v118 = v116
															case 1:
																v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)))
																v118 = v107
															default:
																v115 = base.I32_div_u_s(int32(_a_F_init_span_2), v54)
																v116 = v115
																*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)) = uint16(v116)
																v118 = v116
															}
															v119 = int32(1)
															*(*uint16)(unsafe.Add(mBase, uint32(v51)+30)) = uint16(v119)
															*(*uint16)(unsafe.Add(mBase, uint32(v51)+28)) = uint16(v118)
															v122 = int32(_a_F_init_span_3)
															*(*uint16)(unsafe.Add(mBase, uint32(v51)+26)) = uint16(v122)
															return
														}
													}
												}
											}
										}
									} else {
										v79 = int32(base.Ui32(v55) >> (uint(int32(27)) % 32))
										v82 = l0 + v79*int32(20)
										v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
										if v83 != 0 {
											v87 = v83
											*(*int32)(unsafe.Add(mBase, uint32(v87+v55&int32(134217727))+4)) = l1
											v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											*(*int32)(unsafe.Add(mBase, uint32(v51))) = l2 - v94
											v97 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
											v98 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = v98
											*(*int32)(unsafe.Add(mBase, uint32(v51)+8)) = v97
											*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = l1
											*(*uint16)(unsafe.Add(mBase, uint32(v51)+22)) = uint16(v98)
											*(*uint16)(unsafe.Add(mBase, uint32(v51)+20)) = uint16(v6)
											*(*int32)(unsafe.Add(mBase, uint32(v51)+16)) = l4
											*(*int32)(unsafe.Add(mBase, uint32(v51)+12)) = l3
											switch v6 {
											case 0:
												v108 = int32(1)
												*(*uint16)(unsafe.Add(mBase, uint32(v51)+22)) = uint16(v108)
												v111 = base.I32_div_u_s(int32(_a_F_init_span_1), v54)
												v116 = v111 - v108
												*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)) = uint16(v116)
												v118 = v116
											case 1:
												v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)))
												v118 = v107
											default:
												v115 = base.I32_div_u_s(int32(_a_F_init_span_2), v54)
												v116 = v115
												*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)) = uint16(v116)
												v118 = v116
											}
											v119 = int32(1)
											*(*uint16)(unsafe.Add(mBase, uint32(v51)+30)) = uint16(v119)
											*(*uint16)(unsafe.Add(mBase, uint32(v51)+28)) = uint16(v118)
											v122 = int32(_a_F_init_span_3)
											*(*uint16)(unsafe.Add(mBase, uint32(v51)+26)) = uint16(v122)
											return
										} else {
											v84 = F_get_segment_by_index(m, l0, v79)
											mBase = m.M
											v85 = m.ExcPending
											if v85 != 0 {
												return
											} else {
												v86 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
												v87 = v86
												*(*int32)(unsafe.Add(mBase, uint32(v87+v55&int32(134217727))+4)) = l1
												v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												*(*int32)(unsafe.Add(mBase, uint32(v51))) = l2 - v94
												v97 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
												v98 = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = v98
												*(*int32)(unsafe.Add(mBase, uint32(v51)+8)) = v97
												*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = l1
												*(*uint16)(unsafe.Add(mBase, uint32(v51)+22)) = uint16(v98)
												*(*uint16)(unsafe.Add(mBase, uint32(v51)+20)) = uint16(v6)
												*(*int32)(unsafe.Add(mBase, uint32(v51)+16)) = l4
												*(*int32)(unsafe.Add(mBase, uint32(v51)+12)) = l3
												switch v6 {
												case 0:
													v108 = int32(1)
													*(*uint16)(unsafe.Add(mBase, uint32(v51)+22)) = uint16(v108)
													v111 = base.I32_div_u_s(int32(_a_F_init_span_1), v54)
													v116 = v111 - v108
													*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)) = uint16(v116)
													v118 = v116
												case 1:
													v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)))
													v118 = v107
												default:
													v115 = base.I32_div_u_s(int32(_a_F_init_span_2), v54)
													v116 = v115
													*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)) = uint16(v116)
													v118 = v116
												}
												v119 = int32(1)
												*(*uint16)(unsafe.Add(mBase, uint32(v51)+30)) = uint16(v119)
												*(*uint16)(unsafe.Add(mBase, uint32(v51)+28)) = uint16(v118)
												v122 = int32(_a_F_init_span_3)
												*(*uint16)(unsafe.Add(mBase, uint32(v51)+26)) = uint16(v122)
												return
											}
										}
									}
								} else {
									v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									*(*int32)(unsafe.Add(mBase, uint32(v51))) = l2 - v94
									v97 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
									v98 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = v98
									*(*int32)(unsafe.Add(mBase, uint32(v51)+8)) = v97
									*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = l1
									*(*uint16)(unsafe.Add(mBase, uint32(v51)+22)) = uint16(v98)
									*(*uint16)(unsafe.Add(mBase, uint32(v51)+20)) = uint16(v6)
									*(*int32)(unsafe.Add(mBase, uint32(v51)+16)) = l4
									*(*int32)(unsafe.Add(mBase, uint32(v51)+12)) = l3
									switch v6 {
									case 0:
										v108 = int32(1)
										*(*uint16)(unsafe.Add(mBase, uint32(v51)+22)) = uint16(v108)
										v111 = base.I32_div_u_s(int32(_a_F_init_span_1), v54)
										v116 = v111 - v108
										*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)) = uint16(v116)
										v118 = v116
									case 1:
										v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)))
										v118 = v107
									default:
										v115 = base.I32_div_u_s(int32(_a_F_init_span_2), v54)
										v116 = v115
										*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)) = uint16(v116)
										v118 = v116
									}
									v119 = int32(1)
									*(*uint16)(unsafe.Add(mBase, uint32(v51)+30)) = uint16(v119)
									*(*uint16)(unsafe.Add(mBase, uint32(v51)+28)) = uint16(v118)
									v122 = int32(_a_F_init_span_3)
									*(*uint16)(unsafe.Add(mBase, uint32(v51)+26)) = uint16(v122)
									return
								}
							}
						}
					}
				}
			}
		} else {
			v36 = int32(base.Ui32(l1) >> (uint(int32(27)) % 32))
			v39 = l0 + v36*int32(20)
			v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
			if v40 != 0 {
				v44 = v40
				v51 = v44 + l1&int32(134217727)
				v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6<<(uint(int32(1))%32))+uint32(_c_F_init_span[0]))))
				v55 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
				if v55 != 0 {
					v56 = int32(0)
					v59 = base.AtomicRmwOr32(m, v56, int32(_a_F_init_span_0), v56)
					v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
					v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+1468))
					if v60 != v62 {
						v67 = F_LWLockAcquire(m, v61+int32(1476), int32(0))
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return
						} else {
							F_check_for_freed_segments_locked(m, l0)
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return
							} else {
								v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								F_LWLockRelease(m, v71+int32(1476))
								mBase = m.M
								v75 = m.ExcPending
								if v75 != 0 {
									return
								} else {
									v79 = int32(base.Ui32(v55) >> (uint(int32(27)) % 32))
									v82 = l0 + v79*int32(20)
									v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
									if v83 != 0 {
										v87 = v83
										*(*int32)(unsafe.Add(mBase, uint32(v87+v55&int32(134217727))+4)) = l1
										v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										*(*int32)(unsafe.Add(mBase, uint32(v51))) = l2 - v94
										v97 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
										v98 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = v98
										*(*int32)(unsafe.Add(mBase, uint32(v51)+8)) = v97
										*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = l1
										*(*uint16)(unsafe.Add(mBase, uint32(v51)+22)) = uint16(v98)
										*(*uint16)(unsafe.Add(mBase, uint32(v51)+20)) = uint16(v6)
										*(*int32)(unsafe.Add(mBase, uint32(v51)+16)) = l4
										*(*int32)(unsafe.Add(mBase, uint32(v51)+12)) = l3
										switch v6 {
										case 0:
											v108 = int32(1)
											*(*uint16)(unsafe.Add(mBase, uint32(v51)+22)) = uint16(v108)
											v111 = base.I32_div_u_s(int32(_a_F_init_span_1), v54)
											v116 = v111 - v108
											*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)) = uint16(v116)
											v118 = v116
										case 1:
											v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)))
											v118 = v107
										default:
											v115 = base.I32_div_u_s(int32(_a_F_init_span_2), v54)
											v116 = v115
											*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)) = uint16(v116)
											v118 = v116
										}
										v119 = int32(1)
										*(*uint16)(unsafe.Add(mBase, uint32(v51)+30)) = uint16(v119)
										*(*uint16)(unsafe.Add(mBase, uint32(v51)+28)) = uint16(v118)
										v122 = int32(_a_F_init_span_3)
										*(*uint16)(unsafe.Add(mBase, uint32(v51)+26)) = uint16(v122)
										return
									} else {
										v84 = F_get_segment_by_index(m, l0, v79)
										mBase = m.M
										v85 = m.ExcPending
										if v85 != 0 {
											return
										} else {
											v86 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
											v87 = v86
											*(*int32)(unsafe.Add(mBase, uint32(v87+v55&int32(134217727))+4)) = l1
											v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											*(*int32)(unsafe.Add(mBase, uint32(v51))) = l2 - v94
											v97 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
											v98 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = v98
											*(*int32)(unsafe.Add(mBase, uint32(v51)+8)) = v97
											*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = l1
											*(*uint16)(unsafe.Add(mBase, uint32(v51)+22)) = uint16(v98)
											*(*uint16)(unsafe.Add(mBase, uint32(v51)+20)) = uint16(v6)
											*(*int32)(unsafe.Add(mBase, uint32(v51)+16)) = l4
											*(*int32)(unsafe.Add(mBase, uint32(v51)+12)) = l3
											switch v6 {
											case 0:
												v108 = int32(1)
												*(*uint16)(unsafe.Add(mBase, uint32(v51)+22)) = uint16(v108)
												v111 = base.I32_div_u_s(int32(_a_F_init_span_1), v54)
												v116 = v111 - v108
												*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)) = uint16(v116)
												v118 = v116
											case 1:
												v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)))
												v118 = v107
											default:
												v115 = base.I32_div_u_s(int32(_a_F_init_span_2), v54)
												v116 = v115
												*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)) = uint16(v116)
												v118 = v116
											}
											v119 = int32(1)
											*(*uint16)(unsafe.Add(mBase, uint32(v51)+30)) = uint16(v119)
											*(*uint16)(unsafe.Add(mBase, uint32(v51)+28)) = uint16(v118)
											v122 = int32(_a_F_init_span_3)
											*(*uint16)(unsafe.Add(mBase, uint32(v51)+26)) = uint16(v122)
											return
										}
									}
								}
							}
						}
					} else {
						v79 = int32(base.Ui32(v55) >> (uint(int32(27)) % 32))
						v82 = l0 + v79*int32(20)
						v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
						if v83 != 0 {
							v87 = v83
							*(*int32)(unsafe.Add(mBase, uint32(v87+v55&int32(134217727))+4)) = l1
							v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v51))) = l2 - v94
							v97 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
							v98 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = v98
							*(*int32)(unsafe.Add(mBase, uint32(v51)+8)) = v97
							*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = l1
							*(*uint16)(unsafe.Add(mBase, uint32(v51)+22)) = uint16(v98)
							*(*uint16)(unsafe.Add(mBase, uint32(v51)+20)) = uint16(v6)
							*(*int32)(unsafe.Add(mBase, uint32(v51)+16)) = l4
							*(*int32)(unsafe.Add(mBase, uint32(v51)+12)) = l3
							switch v6 {
							case 0:
								v108 = int32(1)
								*(*uint16)(unsafe.Add(mBase, uint32(v51)+22)) = uint16(v108)
								v111 = base.I32_div_u_s(int32(_a_F_init_span_1), v54)
								v116 = v111 - v108
								*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)) = uint16(v116)
								v118 = v116
							case 1:
								v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)))
								v118 = v107
							default:
								v115 = base.I32_div_u_s(int32(_a_F_init_span_2), v54)
								v116 = v115
								*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)) = uint16(v116)
								v118 = v116
							}
							v119 = int32(1)
							*(*uint16)(unsafe.Add(mBase, uint32(v51)+30)) = uint16(v119)
							*(*uint16)(unsafe.Add(mBase, uint32(v51)+28)) = uint16(v118)
							v122 = int32(_a_F_init_span_3)
							*(*uint16)(unsafe.Add(mBase, uint32(v51)+26)) = uint16(v122)
							return
						} else {
							v84 = F_get_segment_by_index(m, l0, v79)
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
								return
							} else {
								v86 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
								v87 = v86
								*(*int32)(unsafe.Add(mBase, uint32(v87+v55&int32(134217727))+4)) = l1
								v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v51))) = l2 - v94
								v97 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
								v98 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = v98
								*(*int32)(unsafe.Add(mBase, uint32(v51)+8)) = v97
								*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = l1
								*(*uint16)(unsafe.Add(mBase, uint32(v51)+22)) = uint16(v98)
								*(*uint16)(unsafe.Add(mBase, uint32(v51)+20)) = uint16(v6)
								*(*int32)(unsafe.Add(mBase, uint32(v51)+16)) = l4
								*(*int32)(unsafe.Add(mBase, uint32(v51)+12)) = l3
								switch v6 {
								case 0:
									v108 = int32(1)
									*(*uint16)(unsafe.Add(mBase, uint32(v51)+22)) = uint16(v108)
									v111 = base.I32_div_u_s(int32(_a_F_init_span_1), v54)
									v116 = v111 - v108
									*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)) = uint16(v116)
									v118 = v116
								case 1:
									v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)))
									v118 = v107
								default:
									v115 = base.I32_div_u_s(int32(_a_F_init_span_2), v54)
									v116 = v115
									*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)) = uint16(v116)
									v118 = v116
								}
								v119 = int32(1)
								*(*uint16)(unsafe.Add(mBase, uint32(v51)+30)) = uint16(v119)
								*(*uint16)(unsafe.Add(mBase, uint32(v51)+28)) = uint16(v118)
								v122 = int32(_a_F_init_span_3)
								*(*uint16)(unsafe.Add(mBase, uint32(v51)+26)) = uint16(v122)
								return
							}
						}
					}
				} else {
					v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v51))) = l2 - v94
					v97 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
					v98 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = v98
					*(*int32)(unsafe.Add(mBase, uint32(v51)+8)) = v97
					*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = l1
					*(*uint16)(unsafe.Add(mBase, uint32(v51)+22)) = uint16(v98)
					*(*uint16)(unsafe.Add(mBase, uint32(v51)+20)) = uint16(v6)
					*(*int32)(unsafe.Add(mBase, uint32(v51)+16)) = l4
					*(*int32)(unsafe.Add(mBase, uint32(v51)+12)) = l3
					switch v6 {
					case 0:
						v108 = int32(1)
						*(*uint16)(unsafe.Add(mBase, uint32(v51)+22)) = uint16(v108)
						v111 = base.I32_div_u_s(int32(_a_F_init_span_1), v54)
						v116 = v111 - v108
						*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)) = uint16(v116)
						v118 = v116
					case 1:
						v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)))
						v118 = v107
					default:
						v115 = base.I32_div_u_s(int32(_a_F_init_span_2), v54)
						v116 = v115
						*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)) = uint16(v116)
						v118 = v116
					}
					v119 = int32(1)
					*(*uint16)(unsafe.Add(mBase, uint32(v51)+30)) = uint16(v119)
					*(*uint16)(unsafe.Add(mBase, uint32(v51)+28)) = uint16(v118)
					v122 = int32(_a_F_init_span_3)
					*(*uint16)(unsafe.Add(mBase, uint32(v51)+26)) = uint16(v122)
					return
				}
			} else {
				v41 = F_get_segment_by_index(m, l0, v36)
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return
				} else {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
					v44 = v43
					v51 = v44 + l1&int32(134217727)
					v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6<<(uint(int32(1))%32))+uint32(_c_F_init_span[0]))))
					v55 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
					if v55 != 0 {
						v56 = int32(0)
						v59 = base.AtomicRmwOr32(m, v56, int32(_a_F_init_span_0), v56)
						v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
						v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+1468))
						if v60 != v62 {
							v67 = F_LWLockAcquire(m, v61+int32(1476), int32(0))
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return
							} else {
								F_check_for_freed_segments_locked(m, l0)
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return
								} else {
									v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									F_LWLockRelease(m, v71+int32(1476))
									mBase = m.M
									v75 = m.ExcPending
									if v75 != 0 {
										return
									} else {
										v79 = int32(base.Ui32(v55) >> (uint(int32(27)) % 32))
										v82 = l0 + v79*int32(20)
										v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
										if v83 != 0 {
											v87 = v83
											*(*int32)(unsafe.Add(mBase, uint32(v87+v55&int32(134217727))+4)) = l1
											v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											*(*int32)(unsafe.Add(mBase, uint32(v51))) = l2 - v94
											v97 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
											v98 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = v98
											*(*int32)(unsafe.Add(mBase, uint32(v51)+8)) = v97
											*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = l1
											*(*uint16)(unsafe.Add(mBase, uint32(v51)+22)) = uint16(v98)
											*(*uint16)(unsafe.Add(mBase, uint32(v51)+20)) = uint16(v6)
											*(*int32)(unsafe.Add(mBase, uint32(v51)+16)) = l4
											*(*int32)(unsafe.Add(mBase, uint32(v51)+12)) = l3
											switch v6 {
											case 0:
												v108 = int32(1)
												*(*uint16)(unsafe.Add(mBase, uint32(v51)+22)) = uint16(v108)
												v111 = base.I32_div_u_s(int32(_a_F_init_span_1), v54)
												v116 = v111 - v108
												*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)) = uint16(v116)
												v118 = v116
											case 1:
												v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)))
												v118 = v107
											default:
												v115 = base.I32_div_u_s(int32(_a_F_init_span_2), v54)
												v116 = v115
												*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)) = uint16(v116)
												v118 = v116
											}
											v119 = int32(1)
											*(*uint16)(unsafe.Add(mBase, uint32(v51)+30)) = uint16(v119)
											*(*uint16)(unsafe.Add(mBase, uint32(v51)+28)) = uint16(v118)
											v122 = int32(_a_F_init_span_3)
											*(*uint16)(unsafe.Add(mBase, uint32(v51)+26)) = uint16(v122)
											return
										} else {
											v84 = F_get_segment_by_index(m, l0, v79)
											mBase = m.M
											v85 = m.ExcPending
											if v85 != 0 {
												return
											} else {
												v86 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
												v87 = v86
												*(*int32)(unsafe.Add(mBase, uint32(v87+v55&int32(134217727))+4)) = l1
												v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												*(*int32)(unsafe.Add(mBase, uint32(v51))) = l2 - v94
												v97 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
												v98 = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = v98
												*(*int32)(unsafe.Add(mBase, uint32(v51)+8)) = v97
												*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = l1
												*(*uint16)(unsafe.Add(mBase, uint32(v51)+22)) = uint16(v98)
												*(*uint16)(unsafe.Add(mBase, uint32(v51)+20)) = uint16(v6)
												*(*int32)(unsafe.Add(mBase, uint32(v51)+16)) = l4
												*(*int32)(unsafe.Add(mBase, uint32(v51)+12)) = l3
												switch v6 {
												case 0:
													v108 = int32(1)
													*(*uint16)(unsafe.Add(mBase, uint32(v51)+22)) = uint16(v108)
													v111 = base.I32_div_u_s(int32(_a_F_init_span_1), v54)
													v116 = v111 - v108
													*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)) = uint16(v116)
													v118 = v116
												case 1:
													v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)))
													v118 = v107
												default:
													v115 = base.I32_div_u_s(int32(_a_F_init_span_2), v54)
													v116 = v115
													*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)) = uint16(v116)
													v118 = v116
												}
												v119 = int32(1)
												*(*uint16)(unsafe.Add(mBase, uint32(v51)+30)) = uint16(v119)
												*(*uint16)(unsafe.Add(mBase, uint32(v51)+28)) = uint16(v118)
												v122 = int32(_a_F_init_span_3)
												*(*uint16)(unsafe.Add(mBase, uint32(v51)+26)) = uint16(v122)
												return
											}
										}
									}
								}
							}
						} else {
							v79 = int32(base.Ui32(v55) >> (uint(int32(27)) % 32))
							v82 = l0 + v79*int32(20)
							v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
							if v83 != 0 {
								v87 = v83
								*(*int32)(unsafe.Add(mBase, uint32(v87+v55&int32(134217727))+4)) = l1
								v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v51))) = l2 - v94
								v97 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
								v98 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = v98
								*(*int32)(unsafe.Add(mBase, uint32(v51)+8)) = v97
								*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = l1
								*(*uint16)(unsafe.Add(mBase, uint32(v51)+22)) = uint16(v98)
								*(*uint16)(unsafe.Add(mBase, uint32(v51)+20)) = uint16(v6)
								*(*int32)(unsafe.Add(mBase, uint32(v51)+16)) = l4
								*(*int32)(unsafe.Add(mBase, uint32(v51)+12)) = l3
								switch v6 {
								case 0:
									v108 = int32(1)
									*(*uint16)(unsafe.Add(mBase, uint32(v51)+22)) = uint16(v108)
									v111 = base.I32_div_u_s(int32(_a_F_init_span_1), v54)
									v116 = v111 - v108
									*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)) = uint16(v116)
									v118 = v116
								case 1:
									v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)))
									v118 = v107
								default:
									v115 = base.I32_div_u_s(int32(_a_F_init_span_2), v54)
									v116 = v115
									*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)) = uint16(v116)
									v118 = v116
								}
								v119 = int32(1)
								*(*uint16)(unsafe.Add(mBase, uint32(v51)+30)) = uint16(v119)
								*(*uint16)(unsafe.Add(mBase, uint32(v51)+28)) = uint16(v118)
								v122 = int32(_a_F_init_span_3)
								*(*uint16)(unsafe.Add(mBase, uint32(v51)+26)) = uint16(v122)
								return
							} else {
								v84 = F_get_segment_by_index(m, l0, v79)
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
									return
								} else {
									v86 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
									v87 = v86
									*(*int32)(unsafe.Add(mBase, uint32(v87+v55&int32(134217727))+4)) = l1
									v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									*(*int32)(unsafe.Add(mBase, uint32(v51))) = l2 - v94
									v97 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
									v98 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = v98
									*(*int32)(unsafe.Add(mBase, uint32(v51)+8)) = v97
									*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = l1
									*(*uint16)(unsafe.Add(mBase, uint32(v51)+22)) = uint16(v98)
									*(*uint16)(unsafe.Add(mBase, uint32(v51)+20)) = uint16(v6)
									*(*int32)(unsafe.Add(mBase, uint32(v51)+16)) = l4
									*(*int32)(unsafe.Add(mBase, uint32(v51)+12)) = l3
									switch v6 {
									case 0:
										v108 = int32(1)
										*(*uint16)(unsafe.Add(mBase, uint32(v51)+22)) = uint16(v108)
										v111 = base.I32_div_u_s(int32(_a_F_init_span_1), v54)
										v116 = v111 - v108
										*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)) = uint16(v116)
										v118 = v116
									case 1:
										v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)))
										v118 = v107
									default:
										v115 = base.I32_div_u_s(int32(_a_F_init_span_2), v54)
										v116 = v115
										*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)) = uint16(v116)
										v118 = v116
									}
									v119 = int32(1)
									*(*uint16)(unsafe.Add(mBase, uint32(v51)+30)) = uint16(v119)
									*(*uint16)(unsafe.Add(mBase, uint32(v51)+28)) = uint16(v118)
									v122 = int32(_a_F_init_span_3)
									*(*uint16)(unsafe.Add(mBase, uint32(v51)+26)) = uint16(v122)
									return
								}
							}
						}
					} else {
						v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v51))) = l2 - v94
						v97 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
						v98 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = v98
						*(*int32)(unsafe.Add(mBase, uint32(v51)+8)) = v97
						*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = l1
						*(*uint16)(unsafe.Add(mBase, uint32(v51)+22)) = uint16(v98)
						*(*uint16)(unsafe.Add(mBase, uint32(v51)+20)) = uint16(v6)
						*(*int32)(unsafe.Add(mBase, uint32(v51)+16)) = l4
						*(*int32)(unsafe.Add(mBase, uint32(v51)+12)) = l3
						switch v6 {
						case 0:
							v108 = int32(1)
							*(*uint16)(unsafe.Add(mBase, uint32(v51)+22)) = uint16(v108)
							v111 = base.I32_div_u_s(int32(_a_F_init_span_1), v54)
							v116 = v111 - v108
							*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)) = uint16(v116)
							v118 = v116
						case 1:
							v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)))
							v118 = v107
						default:
							v115 = base.I32_div_u_s(int32(_a_F_init_span_2), v54)
							v116 = v115
							*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)) = uint16(v116)
							v118 = v116
						}
						v119 = int32(1)
						*(*uint16)(unsafe.Add(mBase, uint32(v51)+30)) = uint16(v119)
						*(*uint16)(unsafe.Add(mBase, uint32(v51)+28)) = uint16(v118)
						v122 = int32(_a_F_init_span_3)
						*(*uint16)(unsafe.Add(mBase, uint32(v51)+26)) = uint16(v122)
						return
					}
				}
			}
		}
	} else {
		v51 = int32(0)
		v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6<<(uint(int32(1))%32))+uint32(_c_F_init_span[0]))))
		v55 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
		if v55 != 0 {
			v56 = int32(0)
			v59 = base.AtomicRmwOr32(m, v56, int32(_a_F_init_span_0), v56)
			v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
			v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+1468))
			if v60 != v62 {
				v67 = F_LWLockAcquire(m, v61+int32(1476), int32(0))
				mBase = m.M
				v68 = m.ExcPending
				if v68 != 0 {
					return
				} else {
					F_check_for_freed_segments_locked(m, l0)
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return
					} else {
						v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						F_LWLockRelease(m, v71+int32(1476))
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
							return
						} else {
							v79 = int32(base.Ui32(v55) >> (uint(int32(27)) % 32))
							v82 = l0 + v79*int32(20)
							v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
							if v83 != 0 {
								v87 = v83
								*(*int32)(unsafe.Add(mBase, uint32(v87+v55&int32(134217727))+4)) = l1
								v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v51))) = l2 - v94
								v97 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
								v98 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = v98
								*(*int32)(unsafe.Add(mBase, uint32(v51)+8)) = v97
								*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = l1
								*(*uint16)(unsafe.Add(mBase, uint32(v51)+22)) = uint16(v98)
								*(*uint16)(unsafe.Add(mBase, uint32(v51)+20)) = uint16(v6)
								*(*int32)(unsafe.Add(mBase, uint32(v51)+16)) = l4
								*(*int32)(unsafe.Add(mBase, uint32(v51)+12)) = l3
								switch v6 {
								case 0:
									v108 = int32(1)
									*(*uint16)(unsafe.Add(mBase, uint32(v51)+22)) = uint16(v108)
									v111 = base.I32_div_u_s(int32(_a_F_init_span_1), v54)
									v116 = v111 - v108
									*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)) = uint16(v116)
									v118 = v116
								case 1:
									v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)))
									v118 = v107
								default:
									v115 = base.I32_div_u_s(int32(_a_F_init_span_2), v54)
									v116 = v115
									*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)) = uint16(v116)
									v118 = v116
								}
								v119 = int32(1)
								*(*uint16)(unsafe.Add(mBase, uint32(v51)+30)) = uint16(v119)
								*(*uint16)(unsafe.Add(mBase, uint32(v51)+28)) = uint16(v118)
								v122 = int32(_a_F_init_span_3)
								*(*uint16)(unsafe.Add(mBase, uint32(v51)+26)) = uint16(v122)
								return
							} else {
								v84 = F_get_segment_by_index(m, l0, v79)
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
									return
								} else {
									v86 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
									v87 = v86
									*(*int32)(unsafe.Add(mBase, uint32(v87+v55&int32(134217727))+4)) = l1
									v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									*(*int32)(unsafe.Add(mBase, uint32(v51))) = l2 - v94
									v97 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
									v98 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = v98
									*(*int32)(unsafe.Add(mBase, uint32(v51)+8)) = v97
									*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = l1
									*(*uint16)(unsafe.Add(mBase, uint32(v51)+22)) = uint16(v98)
									*(*uint16)(unsafe.Add(mBase, uint32(v51)+20)) = uint16(v6)
									*(*int32)(unsafe.Add(mBase, uint32(v51)+16)) = l4
									*(*int32)(unsafe.Add(mBase, uint32(v51)+12)) = l3
									switch v6 {
									case 0:
										v108 = int32(1)
										*(*uint16)(unsafe.Add(mBase, uint32(v51)+22)) = uint16(v108)
										v111 = base.I32_div_u_s(int32(_a_F_init_span_1), v54)
										v116 = v111 - v108
										*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)) = uint16(v116)
										v118 = v116
									case 1:
										v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)))
										v118 = v107
									default:
										v115 = base.I32_div_u_s(int32(_a_F_init_span_2), v54)
										v116 = v115
										*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)) = uint16(v116)
										v118 = v116
									}
									v119 = int32(1)
									*(*uint16)(unsafe.Add(mBase, uint32(v51)+30)) = uint16(v119)
									*(*uint16)(unsafe.Add(mBase, uint32(v51)+28)) = uint16(v118)
									v122 = int32(_a_F_init_span_3)
									*(*uint16)(unsafe.Add(mBase, uint32(v51)+26)) = uint16(v122)
									return
								}
							}
						}
					}
				}
			} else {
				v79 = int32(base.Ui32(v55) >> (uint(int32(27)) % 32))
				v82 = l0 + v79*int32(20)
				v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
				if v83 != 0 {
					v87 = v83
					*(*int32)(unsafe.Add(mBase, uint32(v87+v55&int32(134217727))+4)) = l1
					v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v51))) = l2 - v94
					v97 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
					v98 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = v98
					*(*int32)(unsafe.Add(mBase, uint32(v51)+8)) = v97
					*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = l1
					*(*uint16)(unsafe.Add(mBase, uint32(v51)+22)) = uint16(v98)
					*(*uint16)(unsafe.Add(mBase, uint32(v51)+20)) = uint16(v6)
					*(*int32)(unsafe.Add(mBase, uint32(v51)+16)) = l4
					*(*int32)(unsafe.Add(mBase, uint32(v51)+12)) = l3
					switch v6 {
					case 0:
						v108 = int32(1)
						*(*uint16)(unsafe.Add(mBase, uint32(v51)+22)) = uint16(v108)
						v111 = base.I32_div_u_s(int32(_a_F_init_span_1), v54)
						v116 = v111 - v108
						*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)) = uint16(v116)
						v118 = v116
					case 1:
						v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)))
						v118 = v107
					default:
						v115 = base.I32_div_u_s(int32(_a_F_init_span_2), v54)
						v116 = v115
						*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)) = uint16(v116)
						v118 = v116
					}
					v119 = int32(1)
					*(*uint16)(unsafe.Add(mBase, uint32(v51)+30)) = uint16(v119)
					*(*uint16)(unsafe.Add(mBase, uint32(v51)+28)) = uint16(v118)
					v122 = int32(_a_F_init_span_3)
					*(*uint16)(unsafe.Add(mBase, uint32(v51)+26)) = uint16(v122)
					return
				} else {
					v84 = F_get_segment_by_index(m, l0, v79)
					mBase = m.M
					v85 = m.ExcPending
					if v85 != 0 {
						return
					} else {
						v86 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
						v87 = v86
						*(*int32)(unsafe.Add(mBase, uint32(v87+v55&int32(134217727))+4)) = l1
						v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v51))) = l2 - v94
						v97 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
						v98 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = v98
						*(*int32)(unsafe.Add(mBase, uint32(v51)+8)) = v97
						*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = l1
						*(*uint16)(unsafe.Add(mBase, uint32(v51)+22)) = uint16(v98)
						*(*uint16)(unsafe.Add(mBase, uint32(v51)+20)) = uint16(v6)
						*(*int32)(unsafe.Add(mBase, uint32(v51)+16)) = l4
						*(*int32)(unsafe.Add(mBase, uint32(v51)+12)) = l3
						switch v6 {
						case 0:
							v108 = int32(1)
							*(*uint16)(unsafe.Add(mBase, uint32(v51)+22)) = uint16(v108)
							v111 = base.I32_div_u_s(int32(_a_F_init_span_1), v54)
							v116 = v111 - v108
							*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)) = uint16(v116)
							v118 = v116
						case 1:
							v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)))
							v118 = v107
						default:
							v115 = base.I32_div_u_s(int32(_a_F_init_span_2), v54)
							v116 = v115
							*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)) = uint16(v116)
							v118 = v116
						}
						v119 = int32(1)
						*(*uint16)(unsafe.Add(mBase, uint32(v51)+30)) = uint16(v119)
						*(*uint16)(unsafe.Add(mBase, uint32(v51)+28)) = uint16(v118)
						v122 = int32(_a_F_init_span_3)
						*(*uint16)(unsafe.Add(mBase, uint32(v51)+26)) = uint16(v122)
						return
					}
				}
			}
		} else {
			v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			*(*int32)(unsafe.Add(mBase, uint32(v51))) = l2 - v94
			v97 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
			v98 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = v98
			*(*int32)(unsafe.Add(mBase, uint32(v51)+8)) = v97
			*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = l1
			*(*uint16)(unsafe.Add(mBase, uint32(v51)+22)) = uint16(v98)
			*(*uint16)(unsafe.Add(mBase, uint32(v51)+20)) = uint16(v6)
			*(*int32)(unsafe.Add(mBase, uint32(v51)+16)) = l4
			*(*int32)(unsafe.Add(mBase, uint32(v51)+12)) = l3
			switch v6 {
			case 0:
				v108 = int32(1)
				*(*uint16)(unsafe.Add(mBase, uint32(v51)+22)) = uint16(v108)
				v111 = base.I32_div_u_s(int32(_a_F_init_span_1), v54)
				v116 = v111 - v108
				*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)) = uint16(v116)
				v118 = v116
			case 1:
				v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)))
				v118 = v107
			default:
				v115 = base.I32_div_u_s(int32(_a_F_init_span_2), v54)
				v116 = v115
				*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)) = uint16(v116)
				v118 = v116
			}
			v119 = int32(1)
			*(*uint16)(unsafe.Add(mBase, uint32(v51)+30)) = uint16(v119)
			*(*uint16)(unsafe.Add(mBase, uint32(v51)+28)) = uint16(v118)
			v122 = int32(_a_F_init_span_3)
			*(*uint16)(unsafe.Add(mBase, uint32(v51)+26)) = uint16(v122)
			return
		}
	}
}
func F_initcap_wbnext(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
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
	var v75 int32
	_ = v75
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v198 int32
	_ = v198
	var v205 int32
	_ = v205
	var v216 int32
	_ = v216
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(v7) < base.Ui32(v8) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v11 = v7
	goto L4
L2:
	;
	v243 = v8
	goto L3
L3:
	;
	return v243
L4:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v18 = int32(*(*int8)(unsafe.Add(mBase, uint32(v16+v11))))
	if int32(0) <= v18 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v243 = v239
	goto L3
L6:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(v45) < base.Ui32(v42+v43) {
		goto L19
	} else {
		goto L20
	}
L7:
	;
	v42 = int32(1)
	goto L6
L8:
	;
	goto L9
L9:
	;
	v23 = v18 & int32(255)
	if v23&int32(224) == int32(192) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v42 = int32(2)
	goto L6
L11:
	;
	goto L12
L12:
	;
	if v23&int32(240) == int32(224) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v42 = int32(3)
	goto L6
L14:
	;
	goto L15
L15:
	;
	if v23&int32(248) == int32(240) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v40 = int32(4)
	goto L18
L17:
	;
	v40 = int32(1)
	goto L18
L18:
	;
	v42 = v40
	goto L6
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v45
	v48 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)) = uint8(v48)
	return v43
L20:
	;
	goto L21
L21:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v52 = v51 + v43
	v53 = int32(*(*int8)(unsafe.Add(mBase, uint32(v52))))
	v55 = v53 & int32(255)
	if int32(0) <= v53 {
		v112 = v55
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	if base.Ui32(int32(128)) <= base.Ui32(v112) {
		goto L37
	} else {
		goto L38
	}
L23:
	;
	if v55&int32(224) == int32(192) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105+v52))))
	v112 = v107&int32(63) | v104
	goto L22
L25:
	;
	v104 = v55 << (uint(int32(6)) % 32) & int32(1984)
	v105 = int32(1)
	goto L24
L26:
	;
	goto L27
L27:
	;
	if v55&int32(240) == int32(224) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+1)))
	v104 = v55<<(uint(int32(12))%32)&int32(_a_F_initcap_wbnext_0) | v75&int32(63)<<(uint(int32(6))%32)
	v105 = int32(2)
	goto L24
L29:
	;
	goto L30
L30:
	;
	if v55&int32(248) != int32(240) {
		v112 = int32(-1)
		goto L22
	} else {
		goto L31
	}
L31:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+1)))
	v92 = int32(63)
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+2)))
	v104 = v55<<(uint(int32(18))%32)&int32(_a_F_initcap_wbnext_1) | v91&v92<<(uint(int32(12))%32) | v97&v92<<(uint(int32(6))%32)
	v105 = int32(3)
	goto L24
L32:
	;
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)))
	if v224 == int32(1) {
		goto L63
	} else {
		goto L64
	}
L33:
	;
	v223 = v216
	goto L32
L34:
	;
	v216 = base.B2i32(v205&int32(255) == int32(9))
	goto L33
L35:
	;
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+uint32(_c_F_initcap_wbnext[0]))))
	v205 = v198
	goto L34
L36:
	;
	v223 = base.B2i32(base.Ui32(v112-int32(48)) < base.Ui32(int32(10)))
	goto L32
L37:
	;
	v123 = int32(1201)
	v124 = int32(0)
	goto L40
L38:
	;
	goto L39
L39:
	;
	v178 = int32(1)
	v180 = v112 << (uint(v178) % 32)
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+uint32(_c_F_initcap_wbnext[1]))))
	if v181&v178 != 0 {
		v216 = v178
		goto L33
	} else {
		goto L60
	}
L40:
	;
	v129 = base.I32_div_s(v123+v124, int32(2))
	v131 = v129 << (uint(int32(3)) % 32)
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v131)+uint32(_c_F_initcap_wbnext[2])))
	if base.Ui32(v134) < base.Ui32(v112) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	if v113 != 0 {
		goto L36
	} else {
		goto L50
	}
L42:
	;
	if v146 <= v145 {
		v123 = v145
		v124 = v146
		goto L40
	} else {
		goto L49
	}
L43:
	;
	v145 = v123
	v146 = v129 + int32(1)
	goto L42
L44:
	;
	goto L45
L45:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v131)+uint32(_c_F_initcap_wbnext[3])))
	if base.Ui32(v140) <= base.Ui32(v112) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v223 = int32(1)
	goto L32
L47:
	;
	goto L48
L48:
	;
	v145 = v129 - int32(1)
	v146 = v124
	goto L42
L49:
	;
	goto L41
L50:
	;
	v152 = int32(3408)
	v153 = int32(0)
	goto L52
L51:
	;
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160)+uint32(_c_F_initcap_wbnext[4]))))
	v205 = v177
	goto L34
L52:
	;
	v158 = base.I32_div_s(v152+v153, int32(2))
	v160 = v158 * int32(12)
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v160)+uint32(_c_F_initcap_wbnext[5])))
	if base.Ui32(v163) < base.Ui32(v112) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v205 = int32(0)
	goto L34
L54:
	;
	if v174 <= v173 {
		v152 = v173
		v153 = v174
		goto L52
	} else {
		goto L59
	}
L55:
	;
	v173 = v152
	v174 = v158 + int32(1)
	goto L54
L56:
	;
	goto L57
L57:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v160)+uint32(_c_F_initcap_wbnext[6])))
	if base.Ui32(v169) <= base.Ui32(v112) {
		goto L51
	} else {
		goto L58
	}
L58:
	;
	v173 = v158 - int32(1)
	v174 = v153
	goto L54
L59:
	;
	goto L53
L60:
	;
	if v113 == int32(0) {
		goto L35
	} else {
		goto L61
	}
L61:
	;
	goto L36
L62:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v237 = v236 + v42
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v237
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(v237) < base.Ui32(v239) {
		v11 = v237
		goto L4
	} else {
		goto L67
	}
L63:
	;
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+14)))
	if v227 == v223 {
		goto L62
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v229 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)) = uint8(v229)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+14)) = uint8(v223)
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v232 + v42
	return v43
L66:
	;
	goto L65
L67:
	;
	goto L5
}
func F_initial_cost_mergejoin(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v29 float64
	_ = v29
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 float64
	_ = v50
	var v51 float64
	_ = v51
	var v52 float64
	_ = v52
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
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
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
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int64
	_ = v203
	var v208 int32
	_ = v208
	var v209 int64
	_ = v209
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
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
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
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
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
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
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
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
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v421 int64
	_ = v421
	var v422 int32
	_ = v422
	var v423 float64
	_ = v423
	var v424 int32
	_ = v424
	var v431 int64
	_ = v431
	var v432 int32
	_ = v432
	var v433 float64
	_ = v433
	var v434 int32
	_ = v434
	var v437 float64
	_ = v437
	var v439 float64
	_ = v439
	var v440 float64
	_ = v440
	var v452 int64
	_ = v452
	var v453 int32
	_ = v453
	var v454 float64
	_ = v454
	var v455 int32
	_ = v455
	var v462 int64
	_ = v462
	var v463 int32
	_ = v463
	var v464 float64
	_ = v464
	var v465 int32
	_ = v465
	var v468 float64
	_ = v468
	var v470 float64
	_ = v470
	var v471 float64
	_ = v471
	var v482 int32
	_ = v482
	var v485 float64
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 float32
	_ = v489
	var v491 float64
	_ = v491
	var v493 float64
	_ = v493
	var v498 float64
	_ = v498
	var v503 float64
	_ = v503
	var v506 float64
	_ = v506
	var v507 float32
	_ = v507
	var v509 float64
	_ = v509
	var v511 float64
	_ = v511
	var v516 float64
	_ = v516
	var v521 float64
	_ = v521
	var v526 int32
	_ = v526
	var v529 float64
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 float32
	_ = v533
	var v535 float64
	_ = v535
	var v537 float64
	_ = v537
	var v542 float64
	_ = v542
	var v547 float64
	_ = v547
	var v550 float64
	_ = v550
	var v551 float32
	_ = v551
	var v553 float64
	_ = v553
	var v555 float64
	_ = v555
	var v560 float64
	_ = v560
	var v565 float64
	_ = v565
	var v570 float64
	_ = v570
	var v571 float64
	_ = v571
	var v577 float64
	_ = v577
	var v578 float64
	_ = v578
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v641 float64
	_ = v641
	var v643 float64
	_ = v643
	var v645 float64
	_ = v645
	var v647 float64
	_ = v647
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v667 int32
	_ = v667
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v726 int32
	_ = v726
	var v733 int32
	_ = v733
	var v735 int32
	_ = v735
	var v737 int32
	_ = v737
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v744 int32
	_ = v744
	var v751 int32
	_ = v751
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v761 float64
	_ = v761
	var v764 int32
	_ = v764
	var v766 float64
	_ = v766
	var v768 int32
	_ = v768
	var v773 int32
	_ = v773
	var v775 float64
	_ = v775
	var v778 int32
	_ = v778
	var v780 float64
	_ = v780
	var v783 int32
	_ = v783
	var v784 float64
	_ = v784
	var v786 float64
	_ = v786
	var v815 float64
	_ = v815
	var v816 float64
	_ = v816
	var v818 float64
	_ = v818
	var v821 float64
	_ = v821
	var v835 float64
	_ = v835
	var v840 float64
	_ = v840
	var v842 float64
	_ = v842
	var v843 float64
	_ = v843
	var v852 float64
	_ = v852
	var v856 float64
	_ = v856
	var v858 float64
	_ = v858
	var v859 float64
	_ = v859
	var v868 float64
	_ = v868
	var v872 float64
	_ = v872
	var v873 float64
	_ = v873
	var v874 float64
	_ = v874
	var v875 int32
	_ = v875
	var v879 int32
	_ = v879
	var v884 float64
	_ = v884
	var v885 float64
	_ = v885
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v889 int32
	_ = v889
	var v892 int32
	_ = v892
	var v893 float64
	_ = v893
	var v894 float64
	_ = v894
	var v895 int32
	_ = v895
	var v896 float64
	_ = v896
	var v899 float64
	_ = v899
	var v901 float64
	_ = v901
	var v905 float64
	_ = v905
	var v906 float64
	_ = v906
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v916 float64
	_ = v916
	var v918 int64
	_ = v918
	var v920 int64
	_ = v920
	var v921 float64
	_ = v921
	var v923 float64
	_ = v923
	var v929 float64
	_ = v929
	var v931 float64
	_ = v931
	var v934 int32
	_ = v934
	var v936 int64
	_ = v936
	var v937 int32
	_ = v937
	var v940 int32
	_ = v940
	var v943 int32
	_ = v943
	var v944 float64
	_ = v944
	var v946 float64
	_ = v946
	var v947 float64
	_ = v947
	var v951 float64
	_ = v951
	var v954 float64
	_ = v954
	var v958 float64
	_ = v958
	var v965 float64
	_ = v965
	var v966 float64
	_ = v966
	var v968 float64
	_ = v968
	var v972 float64
	_ = v972
	var v976 float64
	_ = v976
	var v977 float64
	_ = v977
	var v979 float64
	_ = v979
	var v983 int32
	_ = v983
	var v987 float64
	_ = v987
	var v988 float64
	_ = v988
	var v993 int32
	_ = v993
	var v994 float64
	_ = v994
	var v999 int32
	_ = v999
	var v1000 float64
	_ = v1000
	var v1001 float64
	_ = v1001
	var v1002 float64
	_ = v1002
	var v1008 int32
	_ = v1008
	var v1013 float64
	_ = v1013
	var v1016 float64
	_ = v1016
	var v1017 float64
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1019 float64
	_ = v1019
	var v1022 float64
	_ = v1022
	var v1024 float64
	_ = v1024
	var v1028 float64
	_ = v1028
	var v1029 float64
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1039 float64
	_ = v1039
	var v1041 int64
	_ = v1041
	var v1043 int64
	_ = v1043
	var v1044 float64
	_ = v1044
	var v1046 float64
	_ = v1046
	var v1052 float64
	_ = v1052
	var v1054 float64
	_ = v1054
	var v1057 int32
	_ = v1057
	var v1059 int64
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1063 int32
	_ = v1063
	var v1066 int32
	_ = v1066
	var v1067 float64
	_ = v1067
	var v1069 float64
	_ = v1069
	var v1070 float64
	_ = v1070
	var v1074 float64
	_ = v1074
	var v1077 float64
	_ = v1077
	var v1081 float64
	_ = v1081
	var v1088 float64
	_ = v1088
	var v1089 float64
	_ = v1089
	var v1091 float64
	_ = v1091
	var v1095 float64
	_ = v1095
	var v1099 float64
	_ = v1099
	var v1100 float64
	_ = v1100
	var v1103 int32
	_ = v1103
	var v1107 float64
	_ = v1107
	var v1110 float64
	_ = v1110
	var v1114 float64
	_ = v1114
	var v1115 float64
	_ = v1115
	var v1116 float64
	_ = v1116
	var v1120 int32
	_ = v1120
	var v1121 float64
	_ = v1121
	var v1127 float64
	_ = v1127
	var v1137 float64
	_ = v1137
	var v1143 float64
	_ = v1143
	var v1157 int32
	_ = v1157
	var v1161 int32
	_ = v1161
	var v1166 int32
	_ = v1166
	v10 = int32(0)
	v29 = float64(0)
	v46 = m.G0
	v48 = v46 - int32(96)
	m.G0 = v48
	v50 = *(*float64)(unsafe.Add(mBase, uint32(l5)+32))
	v51 = *(*float64)(unsafe.Add(mBase, uint32(l4)+32))
	v52 = float64(1)
	if base.B2i32(l3 == v10)|base.B2i32(l2 == int32(2)) != 0 {
		v815 = v52
		v816 = v29
		v818 = v52
		v821 = v29
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1157 = m.ExcPending
	if v1157 != 0 {
		goto L32
	} else {
		goto L246
	}
L2:
	;
	if base.F64_le(v50, float64(0)) != 0 {
		goto L180
	} else {
		goto L181
	}
L3:
	;
	if l6 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v60 = l6
	goto L6
L5:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l4)+64))
	v60 = v59
	goto L6
L6:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
	if l7 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v65 = l7
	goto L9
L8:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l5)+64))
	v65 = v64
	goto L9
L9:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+12))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+8))
	if v63 != v68 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+8))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+8))
	if v71 != v73 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v62)+12))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
	if v75 != v76 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+16)))
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+16)))
	if v78 != v79 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+116))
	if v83 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v82)+44))
	v703 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v703)+8))
	v705 = int32(0)
	if v702 == v705 {
		goto L148
	} else {
		goto L149
	}
L15:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	v197 = m.G0
	v199 = v197 - int32(112)
	m.G0 = v199
	v202 = v48 + int32(80)
	v203 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v202))) = v203
	*(*int64)(unsafe.Add(mBase, uint32(v48))) = v203
	v208 = v48 + int32(72)
	v209 = int64(4607182418800017408)
	*(*int64)(unsafe.Add(mBase, uint32(v208))) = v209
	v212 = v48 + int32(88)
	*(*int64)(unsafe.Add(mBase, uint32(v212))) = v209
	if v196 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L16:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	if v86 <= int32(0) {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v83)+12))
	v94 = int32(0)
	goto L18
L18:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v89+v94<<(uint(int32(2))%32))))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)))
	if v140 != v63 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L15
L20:
	;
	v149 = v94 + int32(1)
	if v86 != v149 {
		v94 = v149
		goto L18
	} else {
		goto L25
	}
L21:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
	if v142 != v71 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v139)+8))
	if v144 != v75 {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139)+12)))
	if v146 == v78 {
		v667 = v139
		goto L14
	} else {
		goto L24
	}
L24:
	;
	goto L20
L25:
	;
	goto L19
L26:
	;
	m.G0 = v199 + int32(112)
	v624 = int32(_a_F_initial_cost_mergejoin_0)
	v625 = *(*int32)(unsafe.Add(mBase, _c_F_initial_cost_mergejoin[0]))
	v627 = *(*int32)(unsafe.Add(mBase, uint32(l0)+300))
	*(*int32)(unsafe.Add(mBase, _c_F_initial_cost_mergejoin[0])) = v627
	v630 = F_palloc(m, int32(48))
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L32
	} else {
		goto L145
	}
L27:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v196)))
	if v217 != int32(17) {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v196)+28))
	if v220 == int32(0) {
		goto L26
	} else {
		goto L29
	}
L29:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v220)+4))
	if v223 < int32(2) {
		goto L26
	} else {
		goto L30
	}
L30:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v220)+12))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)+4))
	if v227 == int32(0) {
		goto L26
	} else {
		goto L31
	}
L31:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v196)+24))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v196)+4))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v226)))
	F_examine_variable(m, l0, v232, int32(0), v199+int32(80))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	return
L33:
	;
	F_examine_variable(m, l0, v227, int32(0), v199+int32(48))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v243 = F_get_opfamily_method(m, v63)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L32
	} else {
		goto L35
	}
L35:
	;
	F_get_op_opfamily_properties(m, v231, v63, int32(0), v199+int32(44), v199+int32(40), v199+int32(36))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L32
	} else {
		goto L36
	}
L36:
	;
	switch v75 - int32(1) {
	case 0:
		goto L40
	default:
		goto L37
	case 4:
		goto L39
	}
L37:
	;
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v199)+88))
	if v598 != 0 {
		goto L139
	} else {
		goto L140
	}
L38:
	;
	v359 = int32(0)
	if base.B2i32(v354 == v359)|base.B2i32(v350 == v359)|(base.B2i32(v352 == v359)|base.B2i32(v349 == v359))|(base.B2i32(v351 == v359)|base.B2i32(v357 == v359)|(base.B2i32(v353 == v359)|base.B2i32(v355 == v359))) != 0 {
		goto L37
	} else {
		goto L71
	}
L39:
	;
	v295 = int32(1)
	v298 = F_IndexAmTranslateCompareType(m, v295, v243, v63, v295)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L32
	} else {
		goto L54
	}
L40:
	;
	v256 = int32(1)
	v258 = F_IndexAmTranslateCompareType(m, v256, v243, v63, v256)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L32
	} else {
		goto L41
	}
L41:
	;
	v262 = F_IndexAmTranslateCompareType(m, int32(2), v243, v63, int32(1))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L32
	} else {
		goto L42
	}
L42:
	;
	v264 = base.I32_extend16_s(v262)
	v265 = base.I32_extend16_s(v258)
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v199)+40))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v199)+36))
	if v266 == v267 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v269 = F_get_opfamily_member(m, v63, v266, v266, v265)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L32
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v275 = F_get_opfamily_member(m, v63, v266, v267, v265)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L32
	} else {
		goto L48
	}
L46:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v199)+40))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v199)+36))
	v273 = F_get_opfamily_member(m, v63, v271, v272, v264)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L32
	} else {
		goto L47
	}
L47:
	;
	v349 = v269
	v350 = v269
	v351 = v269
	v352 = v269
	v353 = v269
	v354 = v269
	v355 = v273
	v356 = v10
	v357 = v273
	goto L38
L48:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v199)+40))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v199)+36))
	v279 = F_get_opfamily_member(m, v63, v277, v278, v264)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L32
	} else {
		goto L49
	}
L49:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v199)+40))
	v282 = F_get_opfamily_member(m, v63, v281, v281, v265)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L32
	} else {
		goto L50
	}
L50:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v199)+36))
	v285 = F_get_opfamily_member(m, v63, v284, v284, v265)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L32
	} else {
		goto L51
	}
L51:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v199)+36))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v199)+40))
	v289 = F_get_opfamily_member(m, v63, v287, v288, v265)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L32
	} else {
		goto L52
	}
L52:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v199)+36))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v199)+40))
	v293 = F_get_opfamily_member(m, v63, v291, v292, v264)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L32
	} else {
		goto L53
	}
L53:
	;
	v349 = v285
	v350 = v285
	v351 = v275
	v352 = v282
	v353 = v289
	v354 = v282
	v355 = v293
	v356 = v10
	v357 = v279
	goto L38
L54:
	;
	v302 = F_IndexAmTranslateCompareType(m, int32(5), v243, v63, int32(1))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L32
	} else {
		goto L55
	}
L55:
	;
	v304 = base.I32_extend16_s(v302)
	v307 = F_IndexAmTranslateCompareType(m, int32(4), v243, v63, int32(1))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L32
	} else {
		goto L56
	}
L56:
	;
	v309 = base.I32_extend16_s(v307)
	v310 = base.I32_extend16_s(v298)
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v199)+40))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v199)+36))
	if v311 == v312 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v314 = F_get_opfamily_member(m, v63, v311, v311, v304)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L32
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v323 = F_get_opfamily_member(m, v63, v311, v312, v304)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L32
	} else {
		goto L63
	}
L60:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v199)+40))
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v199)+36))
	v318 = F_get_opfamily_member(m, v63, v316, v317, v309)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L32
	} else {
		goto L61
	}
L61:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v199)+40))
	v321 = F_get_opfamily_member(m, v63, v320, v320, v310)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L32
	} else {
		goto L62
	}
L62:
	;
	v349 = v321
	v350 = v314
	v351 = v314
	v352 = v321
	v353 = v314
	v354 = v314
	v355 = v318
	v356 = v295
	v357 = v318
	goto L38
L63:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v199)+40))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v199)+36))
	v327 = F_get_opfamily_member(m, v63, v325, v326, v309)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L32
	} else {
		goto L64
	}
L64:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v199)+40))
	v330 = F_get_opfamily_member(m, v63, v329, v329, v304)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L32
	} else {
		goto L65
	}
L65:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v199)+36))
	v333 = F_get_opfamily_member(m, v63, v332, v332, v304)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L32
	} else {
		goto L66
	}
L66:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v199)+40))
	v336 = F_get_opfamily_member(m, v63, v335, v335, v310)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L32
	} else {
		goto L67
	}
L67:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v199)+36))
	v339 = F_get_opfamily_member(m, v63, v338, v338, v310)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L32
	} else {
		goto L68
	}
L68:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v199)+36))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v199)+40))
	v343 = F_get_opfamily_member(m, v63, v341, v342, v304)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L32
	} else {
		goto L69
	}
L69:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v199)+36))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v199)+40))
	v347 = F_get_opfamily_member(m, v63, v345, v346, v309)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L32
	} else {
		goto L70
	}
L70:
	;
	v349 = v339
	v350 = v333
	v351 = v323
	v352 = v336
	v353 = v343
	v354 = v330
	v355 = v347
	v356 = v295
	v357 = v327
	goto L38
L71:
	;
	if v356 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v421 = *(*int64)(unsafe.Add(mBase, uint32(v199)))
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v199)+36))
	v423 = F_scalarineqsel(m, l0, v357, v356, int32(1), v230, v199+int32(80), v421, v422)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L32
	} else {
		goto L84
	}
L73:
	;
	v390 = F_get_variable_range(m, v199+int32(80), v352, v230, v199+int32(24), v199+int32(16))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L32
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v406 = F_get_variable_range(m, v199+int32(80), v352, v230, v199+int32(16), v199+int32(24))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L32
	} else {
		goto L80
	}
L76:
	;
	if v390 == int32(0) {
		goto L37
	} else {
		goto L77
	}
L77:
	;
	v398 = F_get_variable_range(m, v199+int32(48), v349, v230, v199+int32(8), v199)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L32
	} else {
		goto L78
	}
L78:
	;
	if v398 != 0 {
		goto L72
	} else {
		goto L79
	}
L79:
	;
	goto L37
L80:
	;
	if v406 == int32(0) {
		goto L37
	} else {
		goto L81
	}
L81:
	;
	v414 = F_get_variable_range(m, v199+int32(48), v349, v230, v199, v199+int32(8))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L32
	} else {
		goto L82
	}
L82:
	;
	if v414 == int32(0) {
		goto L37
	} else {
		goto L83
	}
L83:
	;
	goto L72
L84:
	;
	if base.F64_ne(v423, float64(0.3333333333333333)) != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v212))) = v423
	goto L87
L86:
	;
	goto L87
L87:
	;
	v431 = *(*int64)(unsafe.Add(mBase, uint32(v199)+16))
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v199)+40))
	v433 = F_scalarineqsel(m, l0, v355, v356, int32(1), v230, v199+int32(48), v431, v432)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L32
	} else {
		goto L89
	}
L88:
	;
	v440 = *(*float64)(unsafe.Add(mBase, uint32(v212)))
	if base.F64_gt(v440, v439) == int32(0) {
		goto L94
	} else {
		goto L95
	}
L89:
	;
	if base.F64_eq(v433, float64(0.3333333333333333)) != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v437 = *(*float64)(unsafe.Add(mBase, uint32(v208)))
	v439 = v437
	goto L88
L91:
	;
	goto L92
L92:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v208))) = v433
	v439 = v433
	goto L88
L93:
	;
	v452 = *(*int64)(unsafe.Add(mBase, uint32(v199)+8))
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v199)+36))
	v454 = F_scalarineqsel(m, l0, v351, v356, int32(0), v230, v199+int32(80), v452, v453)
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L32
	} else {
		goto L98
	}
L94:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v208))) = int64(4607182418800017408)
	if base.F64_gt(v439, v440) != 0 {
		goto L93
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v212))) = int64(4607182418800017408)
	goto L93
L97:
	;
	goto L96
L98:
	;
	if base.F64_ne(v454, float64(0.3333333333333333)) != 0 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v48))) = v454
	goto L101
L100:
	;
	goto L101
L101:
	;
	v462 = *(*int64)(unsafe.Add(mBase, uint32(v199)+24))
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v199)+40))
	v464 = F_scalarineqsel(m, l0, v353, v356, int32(0), v230, v199+int32(48), v462, v463)
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L32
	} else {
		goto L103
	}
L102:
	;
	v471 = *(*float64)(unsafe.Add(mBase, uint32(v48)))
	if base.F64_lt(v471, v470) == int32(0) {
		goto L108
	} else {
		goto L109
	}
L103:
	;
	if base.F64_eq(v464, float64(0.3333333333333333)) != 0 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v468 = *(*float64)(unsafe.Add(mBase, uint32(v202)))
	v470 = v468
	goto L102
L105:
	;
	goto L106
L106:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v202))) = v464
	v470 = v464
	goto L102
L107:
	;
	if v78 == int32(0) {
		goto L112
	} else {
		goto L113
	}
L108:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v202))) = int64(0)
	if base.F64_lt(v470, v471) != 0 {
		goto L107
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v48))) = int64(0)
	goto L107
L111:
	;
	goto L110
L112:
	;
	v570 = *(*float64)(unsafe.Add(mBase, uint32(v48)))
	v571 = *(*float64)(unsafe.Add(mBase, uint32(v212)))
	if base.F64_ge(v570, v571) != 0 {
		goto L135
	} else {
		goto L136
	}
L113:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v199)+88))
	if v482 == int32(0) {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v199)+56))
	if v526 == int32(0) {
		goto L112
	} else {
		goto L125
	}
L115:
	;
	v485 = *(*float64)(unsafe.Add(mBase, uint32(v48)))
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v482)+16))
	v487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v486)+22)))
	v488 = v486 + v487
	v489 = *(*float32)(unsafe.Add(mBase, uint32(v488)+8))
	v491 = base.F64_add(v485, base.F64_promote_f32(v489))
	*(*float64)(unsafe.Add(mBase, uint32(v48))) = v491
	v493 = float64(0)
	if base.F64_lt(v491, v493) == int32(0) {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	v506 = *(*float64)(unsafe.Add(mBase, uint32(v212)))
	v507 = *(*float32)(unsafe.Add(mBase, uint32(v488)+8))
	v509 = base.F64_add(v506, base.F64_promote_f32(v507))
	*(*float64)(unsafe.Add(mBase, uint32(v212))) = v509
	v511 = float64(0)
	if base.F64_lt(v509, v511) == int32(0) {
		goto L121
	} else {
		goto L122
	}
L117:
	;
	v498 = float64(1)
	if base.F64_gt(v491, v498) == int32(0) {
		goto L116
	} else {
		goto L120
	}
L118:
	;
	v503 = v493
	goto L119
L119:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v48))) = v503
	goto L116
L120:
	;
	v503 = v498
	goto L119
L121:
	;
	v516 = float64(1)
	if base.F64_gt(v509, v516) == int32(0) {
		goto L114
	} else {
		goto L124
	}
L122:
	;
	v521 = v511
	goto L123
L123:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v212))) = v521
	goto L114
L124:
	;
	v521 = v516
	goto L123
L125:
	;
	v529 = *(*float64)(unsafe.Add(mBase, uint32(v202)))
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v526)+16))
	v531 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v530)+22)))
	v532 = v530 + v531
	v533 = *(*float32)(unsafe.Add(mBase, uint32(v532)+8))
	v535 = base.F64_add(v529, base.F64_promote_f32(v533))
	*(*float64)(unsafe.Add(mBase, uint32(v202))) = v535
	v537 = float64(0)
	if base.F64_lt(v535, v537) == int32(0) {
		goto L127
	} else {
		goto L128
	}
L126:
	;
	v550 = *(*float64)(unsafe.Add(mBase, uint32(v208)))
	v551 = *(*float32)(unsafe.Add(mBase, uint32(v532)+8))
	v553 = base.F64_add(v550, base.F64_promote_f32(v551))
	*(*float64)(unsafe.Add(mBase, uint32(v208))) = v553
	v555 = float64(0)
	if base.F64_lt(v553, v555) == int32(0) {
		goto L131
	} else {
		goto L132
	}
L127:
	;
	v542 = float64(1)
	if base.F64_gt(v535, v542) == int32(0) {
		goto L126
	} else {
		goto L130
	}
L128:
	;
	v547 = v537
	goto L129
L129:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v202))) = v547
	goto L126
L130:
	;
	v547 = v542
	goto L129
L131:
	;
	v560 = float64(1)
	if base.F64_gt(v553, v560) == int32(0) {
		goto L112
	} else {
		goto L134
	}
L132:
	;
	v565 = v555
	goto L133
L133:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v208))) = v565
	goto L112
L134:
	;
	v565 = v560
	goto L133
L135:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v48))) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v212))) = int64(4607182418800017408)
	goto L137
L136:
	;
	goto L137
L137:
	;
	v577 = *(*float64)(unsafe.Add(mBase, uint32(v202)))
	v578 = *(*float64)(unsafe.Add(mBase, uint32(v208)))
	if base.F64_ge(v577, v578) == int32(0) {
		goto L37
	} else {
		goto L138
	}
L138:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v202))) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v208))) = int64(4607182418800017408)
	goto L37
L139:
	;
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v199)+92))
	m.T0[v599].(func(*base.Module, int32))(m, v598)
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L32
	} else {
		goto L142
	}
L140:
	;
	goto L141
L141:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v199)+56))
	if v602 == int32(0) {
		goto L26
	} else {
		goto L143
	}
L142:
	;
	goto L141
L143:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v199)+60))
	m.T0[v605].(func(*base.Module, int32))(m, v602)
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L32
	} else {
		goto L144
	}
L144:
	;
	goto L26
L145:
	;
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v630))) = v632
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v634)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v630)+4)) = v635
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v62)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v630)+8)) = v637
	v639 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+16)))
	*(*uint8)(unsafe.Add(mBase, uint32(v630)+12)) = uint8(v639)
	v641 = *(*float64)(unsafe.Add(mBase, uint32(v48)))
	*(*float64)(unsafe.Add(mBase, uint32(v630)+16)) = v641
	v643 = *(*float64)(unsafe.Add(mBase, uint32(v48)+88))
	*(*float64)(unsafe.Add(mBase, uint32(v630)+24)) = v643
	v645 = *(*float64)(unsafe.Add(mBase, uint32(v48)+80))
	*(*float64)(unsafe.Add(mBase, uint32(v630)+32)) = v645
	v647 = *(*float64)(unsafe.Add(mBase, uint32(v48)+72))
	*(*float64)(unsafe.Add(mBase, uint32(v630)+40)) = v647
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v82)+116))
	v650 = F_lappend(m, v649, v630)
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L32
	} else {
		goto L146
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82)+116)) = v650
	*(*int32)(unsafe.Add(mBase, _c_F_initial_cost_mergejoin[0])) = v625
	v667 = v630
	goto L14
L147:
	;
	if v758 != 0 {
		goto L161
	} else {
		goto L162
	}
L148:
	;
	v758 = int32(1)
	goto L147
L149:
	;
	goto L150
L150:
	;
	if v704 == int32(0) {
		v751 = v705
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v758 = v751
	goto L147
L152:
	;
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v702)+4))
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v704)+4))
	if v715 < v714 {
		v751 = v705
		goto L151
	} else {
		goto L153
	}
L153:
	;
	v717 = int32(1)
	if v714 <= v717 {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v720 = v717
	goto L156
L155:
	;
	v720 = v714
	goto L156
L156:
	;
	v721 = int32(8)
	v726 = int32(0)
	goto L157
L157:
	;
	v733 = v726 << (uint(int32(2)) % 32)
	v735 = *(*int32)(unsafe.Add(mBase, uint32(v702+v721+v733)))
	v737 = *(*int32)(unsafe.Add(mBase, uint32(v704+v721+v733)))
	v740 = v735 & (v737 ^ int32(-1))
	v742 = base.B2i32(v740 == int32(0))
	if v740 != 0 {
		v751 = v742
		goto L151
	} else {
		goto L159
	}
L158:
	;
	v751 = v742
	goto L151
L159:
	;
	v744 = v726 + int32(1)
	if v744 != v720 {
		v726 = v744
		goto L157
	} else {
		goto L160
	}
L160:
	;
	goto L158
L161:
	;
	v759 = int32(32)
	goto L163
L162:
	;
	v759 = int32(16)
	goto L163
L163:
	;
	v761 = *(*float64)(unsafe.Add(mBase, uint32(v667+v759)))
	if v758 != 0 {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v764 = int32(40)
	goto L166
L165:
	;
	v764 = int32(24)
	goto L166
L166:
	;
	v766 = *(*float64)(unsafe.Add(mBase, uint32(v667+v764)))
	v768 = l2 & int32(-5)
	if v768 == int32(1) {
		v815 = v766
		v816 = v761
		v818 = v52
		v821 = v29
		goto L2
	} else {
		goto L167
	}
L167:
	;
	if v758 != 0 {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v773 = int32(24)
	goto L170
L169:
	;
	v773 = int32(40)
	goto L170
L170:
	;
	v775 = *(*float64)(unsafe.Add(mBase, uint32(v667+v773)))
	if v758 != 0 {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v778 = int32(16)
	goto L173
L172:
	;
	v778 = int32(32)
	goto L173
L173:
	;
	v780 = *(*float64)(unsafe.Add(mBase, uint32(v667+v778)))
	v783 = base.B2i32(v768 == int32(3))
	if v768 == int32(3) {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v784 = float64(1)
	goto L176
L175:
	;
	v784 = v766
	goto L176
L176:
	;
	if v768 == int32(3) {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	v786 = float64(0)
	goto L179
L178:
	;
	v786 = v761
	goto L179
L179:
	;
	v815 = v784
	v816 = v786
	v818 = v775
	v821 = v780
	goto L2
L180:
	;
	v835 = float64(1)
	goto L182
L181:
	;
	v835 = v50
	goto L182
L182:
	;
	if base.F64_le(v51, float64(0)) != 0 {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v840 = float64(1)
	goto L185
L184:
	;
	v840 = v51
	goto L185
L185:
	;
	v842 = float64(1e+100)
	v843 = base.F64_mul(v840, v818)
	if base.F64_gt(v843, v842) != 0 {
		v856 = v842
		goto L186
	} else {
		goto L187
	}
L186:
	;
	v858 = base.F64_nearest(base.F64_mul(v840, v821))
	v859 = base.F64_mul(v835, v815)
	if base.F64_gt(v859, float64(1e+100))|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v859)&int64(9223372036854775807))) != 0 {
		v872 = float64(1e+100)
		goto L190
	} else {
		goto L191
	}
L187:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v843)&int64(9223372036854775807)) {
		v856 = float64(1e+100)
		goto L186
	} else {
		goto L188
	}
L188:
	;
	v852 = float64(1)
	if base.F64_le(v843, v852) != 0 {
		v856 = v852
		goto L186
	} else {
		goto L189
	}
L189:
	;
	v856 = base.F64_nearest(v843)
	goto L186
L190:
	;
	v873 = base.F64_nearest(base.F64_mul(v835, v816))
	v874 = base.F64_div(v858, v840)
	if l6 != 0 {
		goto L194
	} else {
		goto L195
	}
L191:
	;
	v868 = float64(1)
	if base.F64_le(v859, v868) != 0 {
		v872 = v868
		goto L190
	} else {
		goto L192
	}
L192:
	;
	v872 = base.F64_nearest(v859)
	goto L190
L193:
	;
	v1017 = base.F64_div(v873, v835)
	v1018 = *(*int32)(unsafe.Add(mBase, uint32(l5)+40))
	if l7 != 0 {
		goto L223
	} else {
		goto L224
	}
L194:
	;
	v875 = *(*int32)(unsafe.Add(mBase, uint32(l4)+40))
	if l8 <= int32(0) {
		goto L198
	} else {
		goto L199
	}
L195:
	;
	goto L196
L196:
	;
	v999 = *(*int32)(unsafe.Add(mBase, uint32(l4)+40))
	v1000 = *(*float64)(unsafe.Add(mBase, uint32(l4)+56))
	v1001 = *(*float64)(unsafe.Add(mBase, uint32(l4)+48))
	v1002 = base.F64_sub(v1000, v1001)
	v1008 = v999
	v1013 = v1002
	v1016 = base.F64_add(base.F64_mul(v1002, v874), base.F64_add(v1001, float64(0)))
	goto L193
L197:
	;
	v994 = base.F64_sub(v987, v988)
	v1008 = v993
	v1013 = v994
	v1016 = base.F64_add(base.F64_mul(v994, v874), base.F64_add(v988, float64(0)))
	goto L193
L198:
	;
	v896 = float64(2)
	if base.F64_lt(v840, v896) != 0 {
		goto L202
	} else {
		goto L203
	}
L199:
	;
	v879 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_initial_cost_mergejoin[1])))
	if v879&int32(1) == int32(0) {
		goto L198
	} else {
		goto L200
	}
L200:
	;
	v884 = *(*float64)(unsafe.Add(mBase, uint32(l4)+48))
	v885 = *(*float64)(unsafe.Add(mBase, uint32(l4)+56))
	v886 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v886)+32))
	v889 = *(*int32)(unsafe.Add(mBase, _c_F_initial_cost_mergejoin[2]))
	F_cost_incremental_sort(m, v48, l0, l6, l8, v875, v884, v885, v840, v887, v889, float64(-1))
	mBase = m.M
	v892 = m.ExcPending
	if v892 != 0 {
		goto L32
	} else {
		goto L201
	}
L201:
	;
	v893 = *(*float64)(unsafe.Add(mBase, uint32(v48)+56))
	v894 = *(*float64)(unsafe.Add(mBase, uint32(v48)+48))
	v895 = *(*int32)(unsafe.Add(mBase, uint32(v48)+40))
	v987 = v893
	v988 = v894
	v993 = v895
	goto L197
L202:
	;
	v899 = v896
	goto L204
L203:
	;
	v899 = v840
	goto L204
L204:
	;
	v901 = *(*float64)(unsafe.Add(mBase, _c_F_initial_cost_mergejoin[3]))
	v905 = base.F64_mul(v899, base.F64_add(base.F64_add(v901, v901), float64(0)))
	v906 = *(*float64)(unsafe.Add(mBase, uint32(l4)+56))
	v907 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v908 = *(*int32)(unsafe.Add(mBase, uint32(v907)+32))
	v916 = base.F64_mul(v840, base.F64_convert_i32_u((v908+int32(7))&int32(-8)+int32(24)))
	v918 = int64(*(*int32)(unsafe.Add(mBase, _c_F_initial_cost_mergejoin[2])))
	v920 = v918 << (uint(int64(10)) % 64)
	v921 = base.F64_convert_i64_s(v920)
	if base.F64_gt(v916, v921) != 0 {
		goto L206
	} else {
		goto L207
	}
L205:
	;
	v979 = base.F64_add(v906, v977)
	v983 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_initial_cost_mergejoin[4])))
	v987 = base.F64_add(v979, base.F64_mul(v899, v976))
	v988 = v979
	v993 = v875 + (v983 ^ int32(1))
	goto L197
L206:
	;
	v923 = F_log(m, v899)
	mBase = m.M
	v929 = base.F64_ceil(base.F64_mul(v916, float64(0.0001220703125)))
	v931 = base.F64_div(v916, v921)
	v934 = int32(6)
	v936 = base.I64_div_s(v920, int64(278528))
	v937 = base.I32_wrap_i64(v936)
	if v937 <= v934 {
		goto L210
	} else {
		goto L211
	}
L207:
	;
	goto L208
L208:
	;
	v966 = base.F64_add(v899, v899)
	if base.F64_lt(v966, v899) != 0 {
		goto L219
	} else {
		goto L220
	}
L209:
	;
	v944 = base.F64_convert_i32_s(v943)
	if base.F64_gt(v931, v944) != 0 {
		goto L216
	} else {
		goto L217
	}
L210:
	;
	v940 = v934
	goto L212
L211:
	;
	v940 = v937
	goto L212
L212:
	;
	if int32(500) <= v940 {
		goto L213
	} else {
		goto L214
	}
L213:
	;
	v943 = int32(500)
	goto L215
L214:
	;
	v943 = v940
	goto L215
L215:
	;
	goto L209
L216:
	;
	v946 = F_log(m, v931)
	mBase = m.M
	v947 = F_log(m, v944)
	mBase = m.M
	v951 = base.F64_ceil(base.F64_div(v946, v947))
	goto L218
L217:
	;
	v951 = float64(1)
	goto L218
L218:
	;
	v954 = *(*float64)(unsafe.Add(mBase, _c_F_initial_cost_mergejoin[5]))
	v958 = *(*float64)(unsafe.Add(mBase, _c_F_initial_cost_mergejoin[6]))
	v965 = *(*float64)(unsafe.Add(mBase, _c_F_initial_cost_mergejoin[3]))
	v976 = v965
	v977 = base.F64_add(base.F64_mul(base.F64_mul(base.F64_add(v929, v929), v951), base.F64_add(base.F64_mul(v954, float64(0.75)), base.F64_mul(v958, float64(0.25)))), base.F64_mul(base.F64_div(v923, float64(0.693147180559945)), v905))
	goto L205
L219:
	;
	v968 = F_log(m, v966)
	mBase = m.M
	v976 = v901
	v977 = base.F64_mul(base.F64_div(v968, float64(0.693147180559945)), v905)
	goto L205
L220:
	;
	goto L221
L221:
	;
	v972 = F_log(m, v899)
	mBase = m.M
	v976 = v901
	v977 = base.F64_mul(base.F64_div(v972, float64(0.693147180559945)), v905)
	goto L205
L222:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l1)+72)) = v873
	*(*float64)(unsafe.Add(mBase, uint32(l1)+64)) = v858
	*(*float64)(unsafe.Add(mBase, uint32(l1)+56)) = v872
	*(*float64)(unsafe.Add(mBase, uint32(l1)+48)) = v856
	*(*float64)(unsafe.Add(mBase, uint32(l1)+8)) = v1127
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1120 + v1008
	v1137 = base.F64_mul(base.F64_sub(base.F64_div(v872, v835), v1017), v1121)
	*(*float64)(unsafe.Add(mBase, uint32(l1)+32)) = v1137
	v1143 = base.F64_add(base.F64_mul(v1013, base.F64_sub(base.F64_div(v856, v840), v874)), float64(0))
	*(*float64)(unsafe.Add(mBase, uint32(l1)+24)) = v1143
	*(*float64)(unsafe.Add(mBase, uint32(l1)+16)) = base.F64_add(v1137, base.F64_add(v1143, v1127))
	m.G0 = v48 + int32(96)
	return
L223:
	;
	v1019 = float64(2)
	if base.F64_lt(v835, v1019) != 0 {
		goto L226
	} else {
		goto L227
	}
L224:
	;
	goto L225
L225:
	;
	v1114 = *(*float64)(unsafe.Add(mBase, uint32(l5)+56))
	v1115 = *(*float64)(unsafe.Add(mBase, uint32(l5)+48))
	v1116 = base.F64_sub(v1114, v1115)
	v1120 = v1018
	v1121 = v1116
	v1127 = base.F64_add(base.F64_mul(v1116, v1017), base.F64_add(v1016, v1115))
	goto L222
L226:
	;
	v1022 = v1019
	goto L228
L227:
	;
	v1022 = v835
	goto L228
L228:
	;
	v1024 = *(*float64)(unsafe.Add(mBase, _c_F_initial_cost_mergejoin[3]))
	v1028 = base.F64_mul(v1022, base.F64_add(base.F64_add(v1024, v1024), float64(0)))
	v1029 = *(*float64)(unsafe.Add(mBase, uint32(l5)+56))
	v1030 = *(*int32)(unsafe.Add(mBase, uint32(l5)+12))
	v1031 = *(*int32)(unsafe.Add(mBase, uint32(v1030)+32))
	v1039 = base.F64_mul(v835, base.F64_convert_i32_u((v1031+int32(7))&int32(-8)+int32(24)))
	v1041 = int64(*(*int32)(unsafe.Add(mBase, _c_F_initial_cost_mergejoin[2])))
	v1043 = v1041 << (uint(int64(10)) % 64)
	v1044 = base.F64_convert_i64_s(v1043)
	if base.F64_gt(v1039, v1044) != 0 {
		goto L230
	} else {
		goto L231
	}
L229:
	;
	v1103 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_initial_cost_mergejoin[4])))
	v1107 = base.F64_add(v1029, v1100)
	v1110 = base.F64_sub(base.F64_add(v1107, base.F64_mul(v1022, v1099)), v1107)
	v1120 = v1018 + (v1103 ^ int32(1))
	v1121 = v1110
	v1127 = base.F64_add(base.F64_mul(v1110, v1017), base.F64_add(v1016, v1107))
	goto L222
L230:
	;
	v1046 = F_log(m, v1022)
	mBase = m.M
	v1052 = base.F64_ceil(base.F64_mul(v1039, float64(0.0001220703125)))
	v1054 = base.F64_div(v1039, v1044)
	v1057 = int32(6)
	v1059 = base.I64_div_s(v1043, int64(278528))
	v1060 = base.I32_wrap_i64(v1059)
	if v1060 <= v1057 {
		goto L234
	} else {
		goto L235
	}
L231:
	;
	goto L232
L232:
	;
	v1089 = base.F64_add(v1022, v1022)
	if base.F64_lt(v1089, v1022) != 0 {
		goto L243
	} else {
		goto L244
	}
L233:
	;
	v1067 = base.F64_convert_i32_s(v1066)
	if base.F64_gt(v1054, v1067) != 0 {
		goto L240
	} else {
		goto L241
	}
L234:
	;
	v1063 = v1057
	goto L236
L235:
	;
	v1063 = v1060
	goto L236
L236:
	;
	if int32(500) <= v1063 {
		goto L237
	} else {
		goto L238
	}
L237:
	;
	v1066 = int32(500)
	goto L239
L238:
	;
	v1066 = v1063
	goto L239
L239:
	;
	goto L233
L240:
	;
	v1069 = F_log(m, v1054)
	mBase = m.M
	v1070 = F_log(m, v1067)
	mBase = m.M
	v1074 = base.F64_ceil(base.F64_div(v1069, v1070))
	goto L242
L241:
	;
	v1074 = float64(1)
	goto L242
L242:
	;
	v1077 = *(*float64)(unsafe.Add(mBase, _c_F_initial_cost_mergejoin[5]))
	v1081 = *(*float64)(unsafe.Add(mBase, _c_F_initial_cost_mergejoin[6]))
	v1088 = *(*float64)(unsafe.Add(mBase, _c_F_initial_cost_mergejoin[3]))
	v1099 = v1088
	v1100 = base.F64_add(base.F64_mul(base.F64_mul(base.F64_add(v1052, v1052), v1074), base.F64_add(base.F64_mul(v1077, float64(0.75)), base.F64_mul(v1081, float64(0.25)))), base.F64_mul(base.F64_div(v1046, float64(0.693147180559945)), v1028))
	goto L229
L243:
	;
	v1091 = F_log(m, v1089)
	mBase = m.M
	v1099 = v1024
	v1100 = base.F64_mul(base.F64_div(v1091, float64(0.693147180559945)), v1028)
	goto L229
L244:
	;
	goto L245
L245:
	;
	v1095 = F_log(m, v1022)
	mBase = m.M
	v1099 = v1024
	v1100 = base.F64_mul(base.F64_div(v1095, float64(0.693147180559945)), v1028)
	goto L229
L246:
	;
	F_errmsg_internal(m, int32(_a_F_initial_cost_mergejoin_1), int32(0))
	mBase = m.M
	v1161 = m.ExcPending
	if v1161 != 0 {
		goto L32
	} else {
		goto L247
	}
L247:
	;
	F_errfinish(m, int32(_a_F_initial_cost_mergejoin_2), int32(3721), int32(_a_F_initial_cost_mergejoin_3))
	mBase = m.M
	v1166 = m.ExcPending
	if v1166 != 0 {
		goto L32
	} else {
		goto L248
	}
L248:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_initial_cost_nestloop(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v8 float64
	_ = v8
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v31 float64
	_ = v31
	var v32 int32
	_ = v32
	var v35 float64
	_ = v35
	var v36 float64
	_ = v36
	var v38 int32
	_ = v38
	var v41 float64
	_ = v41
	var v42 float64
	_ = v42
	var v45 float64
	_ = v45
	var v46 float64
	_ = v46
	var v47 float64
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v57 float64
	_ = v57
	var v59 int32
	_ = v59
	var v67 float64
	_ = v67
	var v74 float64
	_ = v74
	var v75 float64
	_ = v75
	var v76 float64
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v86 float64
	_ = v86
	var v88 int32
	_ = v88
	var v96 float64
	_ = v96
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 float64
	_ = v105
	var v106 float64
	_ = v106
	var v107 float64
	_ = v107
	var v108 float64
	_ = v108
	var v111 float64
	_ = v111
	var v113 int32
	_ = v113
	var v117 float64
	_ = v117
	var v118 float64
	_ = v118
	var v121 float64
	_ = v121
	var v135 float64
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v151 float64
	_ = v151
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 float64
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v188 float64
	_ = v188
	var v197 int32
	_ = v197
	var v207 float64
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 float64
	_ = v212
	var v218 float64
	_ = v218
	var v219 int32
	_ = v219
	var v220 float64
	_ = v220
	var v222 float64
	_ = v222
	var v225 float64
	_ = v225
	var v226 float64
	_ = v226
	var v229 float64
	_ = v229
	var v233 float64
	_ = v233
	var v236 float64
	_ = v236
	var v241 float64
	_ = v241
	var v243 float64
	_ = v243
	var v248 float64
	_ = v248
	var v256 float64
	_ = v256
	var v257 float64
	_ = v257
	var v266 float64
	_ = v266
	var v267 float64
	_ = v267
	var v282 float64
	_ = v282
	var v284 float64
	_ = v284
	var v285 float64
	_ = v285
	var v288 float64
	_ = v288
	var v292 float64
	_ = v292
	var v293 float64
	_ = v293
	var v294 float64
	_ = v294
	var v295 float64
	_ = v295
	var v296 float64
	_ = v296
	var v301 int32
	_ = v301
	var v306 float64
	_ = v306
	var v313 float64
	_ = v313
	var v317 float64
	_ = v317
	v8 = float64(0)
	v24 = m.G0
	v26 = v24 - int32(16)
	m.G0 = v26
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l4)+40))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l5)+40))
	v30 = *(*int64)(unsafe.Add(mBase, uint32(l6)+40))
	v31 = *(*float64)(unsafe.Add(mBase, uint32(l4)+32))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	switch v32 - int32(352) {
	case 0:
		goto L7
	default:
		goto L2
	case 3, 5:
		goto L5
	case 11:
		goto L6
	case 12, 14:
		goto L4
	case 13:
		goto L3
	}
L1:
	;
	v282 = base.F64_add(v31, float64(-1))
	v284 = *(*float64)(unsafe.Add(mBase, uint32(l4)+56))
	v285 = *(*float64)(unsafe.Add(mBase, uint32(l4)+48))
	v288 = base.F64_add(base.F64_sub(v284, v285), float64(0))
	if base.F64_gt(v31, float64(1)) != 0 {
		goto L41
	} else {
		goto L42
	}
L2:
	;
	v256 = *(*float64)(unsafe.Add(mBase, uint32(l5)+56))
	v257 = *(*float64)(unsafe.Add(mBase, uint32(l5)+48))
	v266 = v256
	v267 = v257
	goto L1
L3:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l5)+72))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+12))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+32))
	v105 = *(*float64)(unsafe.Add(mBase, uint32(v102)+32))
	v106 = *(*float64)(unsafe.Add(mBase, uint32(l5)+96))
	v107 = *(*float64)(unsafe.Add(mBase, uint32(v102)+56))
	v108 = *(*float64)(unsafe.Add(mBase, uint32(v102)+48))
	v111 = *(*float64)(unsafe.Add(mBase, _c_F_initial_cost_nestloop[0]))
	v113 = *(*int32)(unsafe.Add(mBase, _c_F_initial_cost_nestloop[1]))
	v117 = base.F64_mul(base.F64_mul(v111, base.F64_convert_i32_s(v113)), float64(1024))
	v118 = float64(4.294967295e+09)
	if base.F64_lt(v117, v118) != 0 {
		goto L12
	} else {
		goto L13
	}
L4:
	;
	v74 = *(*float64)(unsafe.Add(mBase, _c_F_initial_cost_nestloop[2]))
	v75 = *(*float64)(unsafe.Add(mBase, uint32(l5)+32))
	v76 = base.F64_mul(v74, v75)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l5)+12))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+32))
	v86 = base.F64_mul(v75, base.F64_convert_i32_u((v78+int32(7))&int32(-8)+int32(24)))
	v88 = *(*int32)(unsafe.Add(mBase, _c_F_initial_cost_nestloop[1]))
	if base.F64_gt(v86, base.F64_convert_i32_u(v88<<(uint(int32(10))%32))) == int32(0) {
		v266 = v76
		v267 = v8
		goto L1
	} else {
		goto L10
	}
L5:
	;
	v45 = *(*float64)(unsafe.Add(mBase, _c_F_initial_cost_nestloop[3]))
	v46 = *(*float64)(unsafe.Add(mBase, uint32(l5)+32))
	v47 = base.F64_mul(v45, v46)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l5)+12))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+32))
	v57 = base.F64_mul(v46, base.F64_convert_i32_u((v49+int32(7))&int32(-8)+int32(24)))
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_initial_cost_nestloop[1]))
	if base.F64_gt(v57, base.F64_convert_i32_u(v59<<(uint(int32(10))%32))) == int32(0) {
		v266 = v47
		v267 = v8
		goto L1
	} else {
		goto L9
	}
L6:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l5)+100))
	if v38 != int32(1) {
		goto L2
	} else {
		goto L8
	}
L7:
	;
	v35 = *(*float64)(unsafe.Add(mBase, uint32(l5)+56))
	v36 = *(*float64)(unsafe.Add(mBase, uint32(l5)+48))
	v266 = base.F64_sub(v35, v36)
	v267 = v8
	goto L1
L8:
	;
	v41 = *(*float64)(unsafe.Add(mBase, uint32(l5)+56))
	v42 = *(*float64)(unsafe.Add(mBase, uint32(l5)+48))
	v266 = base.F64_sub(v41, v42)
	v267 = v8
	goto L1
L9:
	;
	v67 = *(*float64)(unsafe.Add(mBase, _c_F_initial_cost_nestloop[4]))
	v266 = base.F64_add(base.F64_mul(v67, base.F64_ceil(base.F64_mul(v57, float64(0.0001220703125)))), v47)
	v267 = v8
	goto L1
L10:
	;
	v96 = *(*float64)(unsafe.Add(mBase, _c_F_initial_cost_nestloop[4]))
	v266 = base.F64_add(base.F64_mul(v96, base.F64_ceil(base.F64_mul(v86, float64(0.0001220703125)))), v76)
	v267 = v8
	goto L1
L11:
	;
	v135 = base.F64_add(base.F64_add(base.F64_mul(v105, float64(8)), float64(28)), base.F64_mul(v105, base.F64_convert_i32_u((v104+int32(7))&int32(-8)+int32(24))))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l5)+80))
	if v136 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L12:
	;
	v121 = v117
	goto L14
L13:
	;
	v121 = v118
	goto L14
L14:
	;
	goto L11
L15:
	;
	v207 = F_estimate_num_groups(m, l0, v197, v106, int32(0), v26+int32(12))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L22
	} else {
		goto L25
	}
L16:
	;
	v188 = v135
	v197 = int32(0)
	goto L15
L17:
	;
	goto L18
L18:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v136)+4))
	if v140 <= int32(0) {
		v188 = v135
		v197 = v136
		goto L15
	} else {
		goto L19
	}
L19:
	;
	v151 = v135
	v161 = int32(0)
	goto L20
L20:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v136)+12))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v167+v161<<(uint(int32(2))%32))))
	v172 = F_get_expr_width(m, l0, v171)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l5)+80))
	v188 = v175
	v197 = v180
	goto L15
L22:
	;
	return
L23:
	;
	v175 = base.F64_add(v151, base.F64_convert_i32_s(v172))
	v177 = v161 + int32(1)
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v136)+4))
	if v177 < v178 {
		v151 = v175
		v161 = v177
		goto L20
	} else {
		goto L24
	}
L24:
	;
	goto L21
L25:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	if v209&int32(1) != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v212 = v106
	goto L28
L27:
	;
	v212 = v207
	goto L28
L28:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l5)+104)) = v212
	v218 = base.F64_floor(base.F64_div(base.F64_convert_i32_u(base.I32_trunc_sat_f64_u(v121)), v188))
	v219 = base.F64_lt(v218, v212)
	if v219 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v220 = v212
	goto L31
L30:
	;
	v220 = v218
	goto L31
L31:
	;
	v222 = base.F64_mul(base.F64_div(base.F64_sub(v106, v212), v106), base.F64_div(v218, v220))
	*(*float64)(unsafe.Add(mBase, uint32(l5)+112)) = v222
	if base.F64_gt(v218, v212) != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v225 = v212
	goto L34
L33:
	;
	v225 = v218
	goto L34
L34:
	;
	v226 = float64(4.294967295e+09)
	if base.F64_lt(v225, v226) != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v229 = v225
	goto L37
L36:
	;
	v229 = v226
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5)+88)) = base.I32_trunc_sat_f64_u(v229)
	v233 = *(*float64)(unsafe.Add(mBase, _c_F_initial_cost_nestloop[2]))
	v236 = *(*float64)(unsafe.Add(mBase, _c_F_initial_cost_nestloop[3]))
	if v219 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v241 = v218
	goto L40
L39:
	;
	v241 = v212
	goto L40
L40:
	;
	v243 = base.F64_sub(float64(1), base.F64_div(v241, v212))
	v248 = base.F64_sub(float64(1), v222)
	v266 = base.F64_add(base.F64_add(base.F64_mul(v233, v105), v236), base.F64_add(base.F64_mul(base.F64_mul(base.F64_div(v233, float64(10)), v243), v105), base.F64_add(base.F64_mul(v236, v243), base.F64_add(base.F64_mul(v107, v248), v233))))
	v267 = base.F64_add(v236, base.F64_mul(v108, v248))
	goto L1
L41:
	;
	v292 = base.F64_add(base.F64_mul(v282, v267), v288)
	goto L43
L42:
	;
	v292 = v288
	goto L43
L43:
	;
	v293 = base.F64_sub(v266, v267)
	v294 = *(*float64)(unsafe.Add(mBase, uint32(l5)+56))
	v295 = *(*float64)(unsafe.Add(mBase, uint32(l5)+48))
	v296 = base.F64_sub(v294, v295)
	if l2&int32(-2) != int32(4) {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l1)+24)) = v313
	v317 = base.F64_add(base.F64_add(v285, v295), float64(0))
	*(*float64)(unsafe.Add(mBase, uint32(l1)+8)) = v317
	*(*float64)(unsafe.Add(mBase, uint32(l1)+16)) = base.F64_add(v317, v313)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v28 + v29 + base.B2i32(l3&v30 != l3)
	m.G0 = v26 + int32(16)
	return
L45:
	;
	v306 = base.F64_add(v296, v292)
	if base.F64_gt(v31, float64(1)) == int32(0) {
		v313 = v306
		goto L44
	} else {
		goto L50
	}
L46:
	;
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+8)))
	if v301 != int32(1) {
		goto L45
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l1)+40)) = v293
	*(*float64)(unsafe.Add(mBase, uint32(l1)+32)) = v296
	v313 = v292
	goto L44
L49:
	;
	goto L48
L50:
	;
	v313 = base.F64_add(base.F64_mul(v282, v293), v306)
	goto L44
}
func F_initialize_acl(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int64
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_initialize_acl[0]))
	if v2 != 0 {
		v6 = int64(*(*uint32)(unsafe.Add(mBase, _c_F_initialize_acl[1])))
		v8 = F_GetSysCacheHashValue(m, int32(21), v6, int64(0))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_initialize_acl[2])) = v8
			F_CacheRegisterSyscacheCallback(m, int32(9), int32(1367), int64(0))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				F_CacheRegisterSyscacheCallback(m, int32(11), int32(1367), int64(0))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					F_CacheRegisterSyscacheCallback(m, int32(21), int32(1367), int64(0))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	} else {
		return
	}
}
func F_inject_projection_plan(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v22 int32
	_ = v22
	var v24 float64
	_ = v24
	var v26 float64
	_ = v26
	var v28 float64
	_ = v28
	var v30 int32
	_ = v30
	v3 = l2
	v6 = F_palloc0(m, int32(88))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v6)+80)) = v10
		*(*int64)(unsafe.Add(mBase, uint32(v6)+72)) = int64(0)
		*(*int32)(unsafe.Add(mBase, uint32(v6)+56)) = v10
		*(*int32)(unsafe.Add(mBase, uint32(v6)+52)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v6)+48)) = v10
		*(*int32)(unsafe.Add(mBase, uint32(v6)+44)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(335)
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v22
		v24 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
		*(*float64)(unsafe.Add(mBase, uint32(v6)+8)) = v24
		v26 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
		*(*float64)(unsafe.Add(mBase, uint32(v6)+16)) = v26
		v28 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
		*(*float64)(unsafe.Add(mBase, uint32(v6)+24)) = v28
		v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
		*(*uint8)(unsafe.Add(mBase, uint32(v6)+37)) = uint8(v3)
		*(*uint8)(unsafe.Add(mBase, uint32(v6)+36)) = uint8(v10)
		*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = v30
		return v6
	}
}
func F_inline_cte_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
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
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int64
	_ = v75
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	v3 = int32(0)
	if l0 == v3 {
		v92 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v92
L2:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v7 != int32(101) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v89 = F_expression_tree_walker_impl(m, l0, int32(892), l1)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L8
	} else {
		goto L25
	}
L4:
	;
	if v7 != int32(67) {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v28 != int32(6) {
		v92 = v3
		goto L1
	} else {
		goto L10
	}
L7:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v12 + int32(1)
	v18 = F_query_tree_walker_impl(m, l0, int32(892), l1, int32(32))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return int32(0)
L9:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v22 - int32(1)
	return int32(0)
L10:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	if base.B2i32(v35 == int32(0))|base.B2i32(v35 != v38) != 0 {
		v56 = v35
		v57 = v38
		goto L12
	} else {
		goto L13
	}
L11:
	;
	if v56-v57 != 0 {
		v92 = v3
		goto L1
	} else {
		goto L18
	}
L12:
	;
	goto L11
L13:
	;
	v41 = v31
	v42 = v32
	goto L14
L14:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+1)))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1)))
	if v46 == int32(0) {
		v56 = v46
		v57 = v45
		goto L12
	} else {
		goto L16
	}
L15:
	;
	v56 = v46
	v57 = v45
	goto L12
L16:
	;
	v49 = int32(1)
	if v46 == v45 {
		v41 = v41 + v49
		v42 = v42 + v49
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v59 != v60 {
		v92 = v3
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v63 = l0 + int32(84)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v65 = F_copyObjectImpl(m, v64)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L8
	} else {
		goto L20
	}
L20:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v67 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	F_IncrementVarSublevelsUp(m, v65, v67, int32(1))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L8
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v73 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v73
	v75 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+96)) = v75
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)) = uint8(v73)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v63)+8)) = uint8(v73)
	*(*int64)(unsafe.Add(mBase, uint32(v63))) = v75
	return v73
L24:
	;
	goto L23
L25:
	;
	v92 = v89
	goto L1
}
func F_insertStatEntry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v184 int32
	_ = v184
	var v202 int32
	_ = v202
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v242 int32
	_ = v242
	var __phi242 int32
	_ = __phi242
	var v244 int32
	_ = v244
	var __phi244 int32
	_ = __phi244
	var v249 int32
	_ = v249
	var __phi249 int32
	_ = __phi249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
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
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
	v5 = int32(0)
	v17 = int32(1)
	v19 = l2 + int32(8)
	v22 = v19 + l3<<(uint(int32(2))%32)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v25 = v23 & v17
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v27 == v5 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	return
L2:
	;
	if v26 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L3:
	;
	if v184 == int32(0) {
		goto L1
	} else {
		goto L21
	}
L4:
	;
	if v25 == int32(0) {
		v202 = v17
		goto L2
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	if v25 == int32(0) {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v36 = int32(1)
	v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19+v32<<(uint(int32(2))%32)+(int32(base.Ui32(v23)>>(uint(v36)%32))&int32(2047)+int32(base.Ui32(v23)>>(uint(int32(12))%32))+v36)&int32(_a_F_insertStatEntry_0)))))
	v184 = v48
	goto L3
L8:
	;
	v51 = int32(1)
	v61 = (int32(base.Ui32(v23)>>(uint(v51)%32))&int32(2047) + int32(base.Ui32(v23)>>(uint(int32(12))%32)) + v51) & int32(_a_F_insertStatEntry_0)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v64 = v62 << (uint(int32(2)) % 32)
	v67 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61+(v19+v64)))))
	if v67 == int32(0) {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v71 = v67 & int32(3)
	v75 = l2 + v64 + v61 + int32(10)
	if base.Ui32(v67) < base.Ui32(int32(4)) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v152 = v136
	v156 = v140
	v157 = v5
	goto L18
L11:
	;
	v136 = v75
	v140 = int32(0)
	goto L10
L12:
	;
	goto L13
L13:
	;
	v85 = v75
	v89 = int32(0)
	v92 = v5
	goto L14
L14:
	;
	v98 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v85))))
	v99 = int32(14)
	v102 = int32(1)
	v105 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v85)+2)))
	v112 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v85)+4)))
	v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v85)+6)))
	v125 = int32(base.Ui32(v27)>>(uint(int32(base.Ui32(v98)>>(uint(v99)%32)))%32))&v102 + v89 + int32(base.Ui32(v27)>>(uint(int32(base.Ui32(v105)>>(uint(v99)%32)))%32))&v102 + int32(base.Ui32(v27)>>(uint(int32(base.Ui32(v112)>>(uint(v99)%32)))%32))&v102 + int32(base.Ui32(v27)>>(uint(int32(base.Ui32(v119)>>(uint(v99)%32)))%32))&v102
	v127 = v85 + int32(8)
	v129 = v92 + int32(4)
	if v129 != v67&int32(_a_F_insertStatEntry_1) {
		v85 = v127
		v89 = v125
		v92 = v129
		goto L14
	} else {
		goto L16
	}
L15:
	;
	if v71 == int32(0) {
		v184 = v125
		goto L3
	} else {
		goto L17
	}
L16:
	;
	goto L15
L17:
	;
	v136 = v127
	v140 = v125
	goto L10
L18:
	;
	v165 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v152))))
	v169 = int32(1)
	v171 = int32(base.Ui32(v27)>>(uint(int32(base.Ui32(v165)>>(uint(int32(14))%32)))%32))&v169 + v156
	v175 = v157 + v169
	if v175 != v71 {
		v152 = v152 + int32(2)
		v156 = v171
		v157 = v175
		goto L18
	} else {
		goto L20
	}
L19:
	;
	v184 = v171
	goto L3
L20:
	;
	goto L19
L21:
	;
	v202 = v184
	goto L2
L22:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.Ui32(v367) < base.Ui32(v358) {
		goto L71
	} else {
		goto L72
	}
L23:
	;
	v213 = int32(0)
	v214 = int32(1)
	v353 = v213
	v354 = v213
	v356 = v214
	v358 = v214
	v366 = v213
	goto L22
L24:
	;
	goto L25
L25:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v224 = v19 + v218<<(uint(int32(2))%32) + int32(base.Ui32(v23)>>(uint(int32(12))%32))
	v230 = int32(base.Ui32(v23)>>(uint(int32(1))%32)) & int32(2047)
	if v230 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v231 = int32(-1)
	goto L28
L27:
	;
	v231 = int32(0)
	goto L28
L28:
	;
	__phi242 = int32(1)
	__phi244 = int32(0)
	__phi249 = v26
	v242 = __phi242
	v244 = __phi244
	v249 = __phi249
	goto L29
L29:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v249)+16))
	if v250 != 0 {
		goto L34
	} else {
		goto L35
	}
L30:
	;
	v353 = v341
	v354 = v342
	v356 = v344
	v358 = v345
	v366 = int32(base.Ui32(v343) >> (uint(int32(31)) % 32))
	goto L22
L31:
	;
	goto L30
L32:
	;
	v330 = int32(1)
	v332 = v242 + v330
	v333 = int32(0)
	if v329 < v333 {
		goto L67
	} else {
		goto L68
	}
L33:
	;
	v257 = v249 + int32(20)
	if base.Ui32(v250) < base.Ui32(v230) {
		goto L39
	} else {
		goto L40
	}
L34:
	;
	if v230 != 0 {
		goto L33
	} else {
		goto L37
	}
L35:
	;
	v253 = v231
	goto L36
L36:
	;
	if v253 != 0 {
		v329 = v253
		goto L32
	} else {
		goto L38
	}
L37:
	;
	v253 = base.B2i32(int32(0) < v250)
	goto L36
L38:
	;
	v254 = int32(0)
	v341 = v244
	v342 = v249
	v343 = v254
	v344 = v254
	v345 = v242
	goto L31
L39:
	;
	v259 = v250
	goto L41
L40:
	;
	v259 = v230
	goto L41
L41:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v259) {
		goto L45
	} else {
		goto L46
	}
L42:
	;
	if v321 != 0 {
		v329 = v321
		goto L32
	} else {
		goto L60
	}
L43:
	;
	v321 = int32(0)
	goto L42
L44:
	;
	v295 = v290
	v296 = v291
	v297 = v292
	goto L54
L45:
	;
	if (v257|v224)&int32(3) != 0 {
		v290 = v257
		v291 = v224
		v292 = v259
		goto L44
	} else {
		goto L48
	}
L46:
	;
	v283 = v257
	v284 = v224
	v285 = v259
	goto L47
L47:
	;
	if v285 == int32(0) {
		goto L43
	} else {
		goto L53
	}
L48:
	;
	v267 = v257
	v268 = v224
	v269 = v259
	goto L49
L49:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v267)))
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v268)))
	if v272 != v273 {
		v290 = v267
		v291 = v268
		v292 = v269
		goto L44
	} else {
		goto L51
	}
L50:
	;
	v283 = v278
	v284 = v276
	v285 = v280
	goto L47
L51:
	;
	v275 = int32(4)
	v276 = v268 + v275
	v278 = v267 + v275
	v280 = v269 - v275
	if base.Ui32(int32(3)) < base.Ui32(v280) {
		v267 = v278
		v268 = v276
		v269 = v280
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	v290 = v283
	v291 = v284
	v292 = v285
	goto L44
L54:
	;
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295))))
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v296))))
	if v300 == v301 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v321 = v300 - v301
	goto L42
L56:
	;
	v303 = int32(1)
	v308 = v297 - v303
	if v308 != 0 {
		v295 = v295 + v303
		v296 = v296 + v303
		v297 = v308
		goto L54
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	goto L55
L59:
	;
	goto L43
L60:
	;
	if v250 == v230 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v323 = int32(0)
	v341 = v244
	v342 = v249
	v343 = v323
	v344 = v323
	v345 = v242
	goto L31
L62:
	;
	goto L63
L63:
	;
	if v250 < v230 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v328 = int32(-1)
	goto L66
L65:
	;
	v328 = int32(1)
	goto L66
L66:
	;
	v329 = v328
	goto L32
L67:
	;
	v338 = int32(8)
	goto L69
L68:
	;
	v338 = int32(12)
	goto L69
L69:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v249+v338)))
	if v340 != 0 {
		__phi242 = v332
		__phi244 = v249
		__phi249 = v340
		v242 = __phi242
		v244 = __phi244
		v249 = __phi249
		goto L29
	} else {
		goto L70
	}
L70:
	;
	v341 = v249
	v342 = v333
	v343 = v329
	v344 = v330
	v345 = v332
	goto L31
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v358
	goto L73
L72:
	;
	goto L73
L73:
	;
	if v356 != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v377 = F_MemoryContextAlloc(m, l0, int32(base.Ui32(v370)>>(uint(int32(1))%32))&int32(2047)+int32(20))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	goto L76
L76:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v354)))
	*(*int32)(unsafe.Add(mBase, uint32(v354))) = v406 + int32(1)
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v354)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v354)+4)) = v410 + v202
	goto L1
L77:
	;
	return
L78:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v377)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v377)+4)) = v202
	v382 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v377))) = v382
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v388 = int32(base.Ui32(v384)>>(uint(v382)%32)) & int32(2047)
	*(*int32)(unsafe.Add(mBase, uint32(v377)+16)) = v388
	if v388 != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	base.MemoryCopy(m, v377+int32(20), v19+v392<<(uint(int32(2))%32)+int32(base.Ui32(v396)>>(uint(int32(12))%32)), v388)
	goto L81
L80:
	;
	goto L81
L81:
	;
	if v353 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v377
	return
L83:
	;
	goto L84
L84:
	;
	if v366 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v353)+8)) = v377
	return
L86:
	;
	goto L87
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v353)+12)) = v377
	return
}
func F_insert_s(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	v12 = l1 - l2 + l3
	if v12 == int32(0) {
		if l3 != 0 {
			v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			base.MemoryCopy(m, v64+l1, l4, l3)
		} else {
		}
		v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v67 < l1 {
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v67 + v12
			v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			if v71 < l1 {
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v71 + v12
			}
		}
		return int32(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v15-int32(4))))
		v19 = v18 + v12
		v21 = v15 - int32(8)
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
		if v22 < v19 {
			v26 = F_repalloc(m, v21, v19+int32(29))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				if v26 == int32(0) {
					return int32(-1)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v26))) = v19 + int32(20)
					v38 = v26 + int32(8)
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = v38
					v40 = v38
					v41 = v18 - l2
					if v41 != 0 {
						v42 = l2 + v40
						base.MemoryCopy(m, v42+v12, v42, v41)
					} else {
					}
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v46-int32(4)))) = v19
					v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v50 + v12
					v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					if l2 <= v53 {
						v57 = v53 + v12
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v57
					} else {
						if v53 <= l1 {
						} else {
							v57 = l1
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v57
						}
					}
					if l3 != 0 {
						v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						base.MemoryCopy(m, v64+l1, l4, l3)
					} else {
					}
					v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v67 < l1 {
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v67 + v12
						v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						if v71 < l1 {
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v71 + v12
						}
					}
					return int32(0)
				}
			}
		} else {
			v40 = v15
			v41 = v18 - l2
			if v41 != 0 {
				v42 = l2 + v40
				base.MemoryCopy(m, v42+v12, v42, v41)
			} else {
			}
			v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			*(*int32)(unsafe.Add(mBase, uint32(v46-int32(4)))) = v19
			v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v50 + v12
			v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if l2 <= v53 {
				v57 = v53 + v12
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v57
			} else {
				if v53 <= l1 {
				} else {
					v57 = l1
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v57
				}
			}
			if l3 != 0 {
				v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				base.MemoryCopy(m, v64+l1, l4, l3)
			} else {
			}
			v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v67 < l1 {
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v67 + v12
				v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				if v71 < l1 {
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v71 + v12
				}
			}
			return int32(0)
		}
	}
}
func F_instantiate_empty_record_variable(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
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
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v8 == int32(2249) {
		F_errstart_cold(m, int32(21), int32(_a_F_instantiate_empty_record_variable_0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			F_errcode(m, int32(325))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = v18
				F_errmsg(m, int32(_a_F_instantiate_empty_record_variable_1), v6)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					v25 = F_errdetail(m, int32(_a_F_instantiate_empty_record_variable_2), int32(0))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_instantiate_empty_record_variable_3), int32(_a_F_instantiate_empty_record_variable_4), int32(_a_F_instantiate_empty_record_variable_5))
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
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
		F_revalidate_rectypeid(m, l1)
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return
		} else {
			v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
			v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
			v37 = F_make_expanded_record_from_typeid(m, v34, int32(-1), v36)
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v37
				m.G0 = v6 + int32(16)
				return
			}
		}
	}
}
func F_int24mi(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v6 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+24)))
	v7 = v6 - v3
	if base.B2i32(int32(0) < v3) != base.B2i32(v7 < v6) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_int24mi_0), int32(0))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_int24mi_1), int32(1040), int32(_a_F_int24mi_2))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
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
		return base.I64_extend_i32_s(v7)
	}
}
func F_int24mul(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	var v9 int32
	_ = v9
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	v3 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+24)))
	v4 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+40)))
	v5 = v3 * v4
	v9 = base.I32_wrap_i64(v5)
	if base.I32_wrap_i64(int64(base.Ui64(v5)>>(uint(int64(32))%64))) != v9>>(uint(int32(31))%32) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_int24mul_0), int32(0))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_int24mul_1), int32(1054), int32(_a_F_int24mul_2))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
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
		return base.I64_extend_i32_s(v9)
	}
}
func F_int24pl(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v6 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+24)))
	v7 = v6 + v3
	if base.B2i32(v3 < int32(0)) != base.B2i32(v7 < v6) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_int24pl_0), int32(0))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_int24pl_1), int32(1026), int32(_a_F_int24pl_2))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
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
		return base.I64_extend_i32_s(v7)
	}
}
func F_int28ge(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v3 int64
	_ = v3
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v3 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+24)))
	return base.I64_extend_i32_u(base.B2i32(v2 <= v3))
}
func F_int28mul(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int64
	_ = v8
	var v9 int64
	_ = v9
	var v11 int64
	_ = v11
	var v18 int64
	_ = v18
	var v19 int64
	_ = v19
	var v21 int64
	_ = v21
	var v24 int64
	_ = v24
	var v25 int64
	_ = v25
	var v27 int64
	_ = v27
	var v28 int64
	_ = v28
	var v32 int64
	_ = v32
	var v39 int64
	_ = v39
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
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
	v8 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+24)))
	v9 = int64(63)
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v18 = int64(32)
	v19 = int64(base.Ui64(v11) >> (uint(v18) % 64))
	v21 = int64(base.Ui64(v8) >> (uint(v18) % 64))
	v24 = int64(4294967295)
	v25 = v11 & v24
	v27 = v8 & v24
	v28 = v25 * v27
	v32 = int64(base.Ui64(v28)>>(uint(v18)%64)) + v25*v21
	v39 = v27*v19 + v32&v24
	*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = v8*(v11>>(uint(v9)%64)) + v8>>(uint(v9)%64)*v11 + v19*v21 + int64(base.Ui64(v32)>>(uint(v18)%64)) + int64(base.Ui64(v39)>>(uint(v18)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v6))) = v28&v24 | v39<<(uint(v18)%64)
	v50 = *(*int64)(unsafe.Add(mBase, uint32(v6)+8))
	v51 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
	if v50 != v51>>(uint(int64(63))%64) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v60 = m.ExcPending
		if v60 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v63 = m.ExcPending
			if v63 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_int28mul_0), int32(0))
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_int28mul_1), int32(1189), int32(_a_F_int28mul_2))
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
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
		m.G0 = v6 + int32(16)
		return v51
	}
}
func F_int28pl(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v7 int64
	_ = v7
	var v8 int64
	_ = v8
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v7 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+24)))
	v8 = v4 + v7
	if base.B2i32(v4 < int64(0)) != base.B2i32(v8 < v7) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_int28pl_0), int32(0))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_int28pl_1), int32(1161), int32(_a_F_int28pl_2))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
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
		return v8
	}
}
func F_int2mi(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	v2 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+24)))
	v3 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+40)))
	v4 = v2 - v3
	if base.I32_extend16_s(v4) != v4 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_int2mi_0), int32(0))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_int2mi_1), int32(958), int32(_a_F_int2mi_2))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
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
		return base.I64_extend16_s(base.I64_extend_i32_u(v4))
	}
}
func F_int2mod(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int64
	_ = v13
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	v3 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+40)))
	if v3 != int32(_a_F_int2mod_0) {
		if v3 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(33816706))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, int32(_a_F_int2mod_1), int32(0))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_int2mod_2), int32(1196), int32(_a_F_int2mod_3))
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
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
			v8 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+24)))
			v10 = base.I32_rem_s(v8, base.I32_extend16_s(v3))
			v13 = base.I64_extend_i32_s(v10)
			return v13
		}
	} else {
		v13 = int64(0)
		return v13
	}
}
func F_int2mul(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	v2 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+24)))
	v3 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+40)))
	v4 = v2 * v3
	if base.I32_extend16_s(v4) != v4 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_int2mul_0), int32(0))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_int2mul_1), int32(972), int32(_a_F_int2mul_2))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
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
		return base.I64_extend16_s(base.I64_extend_i32_u(v4))
	}
}
func F_int2send(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+24)))
	F_pq_begintypsend(m, v6)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		F_enlargeStringInfo(m, v6, int32(2))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int64(0)
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
			v19 = int32(8)
			v23 = v8<<(uint(v19)%32) | int32(base.Ui32(v8)>>(uint(v19)%32))
			*(*uint16)(unsafe.Add(mBase, uint32(v16+v17))) = uint16(v23)
			v25 = int32(2)
			v26 = v16 + v25
			*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v26
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
			*(*int32)(unsafe.Add(mBase, uint32(v29))) = v26 << (uint(v25) % 32)
			m.G0 = v6 + int32(16)
			return base.I64_extend_i32_u(v29)
		}
	}
}
func F_int2um(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int64
	_ = v27
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if v3&int64(65535) == int64(32768) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_int2um_0), int32(0))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_int2um_1), int32(922), int32(_a_F_int2um_2))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
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
		v27 = int64(48)
		return (int64(0) - v3<<(uint(v27)%64)) >> (uint(v27) % 64)
	}
}
func F_int42mul(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	var v9 int32
	_ = v9
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	v3 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+24)))
	v4 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+40)))
	v5 = v3 * v4
	v9 = base.I32_wrap_i64(v5)
	if base.I32_wrap_i64(int64(base.Ui64(v5)>>(uint(int64(32))%64))) != v9>>(uint(int32(31))%32) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_int42mul_0), int32(0))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_int42mul_1), int32(1115), int32(_a_F_int42mul_2))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
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
		return base.I64_extend_i32_s(v9)
	}
}
func F_int48mul(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int64
	_ = v8
	var v9 int64
	_ = v9
	var v11 int64
	_ = v11
	var v18 int64
	_ = v18
	var v19 int64
	_ = v19
	var v21 int64
	_ = v21
	var v24 int64
	_ = v24
	var v25 int64
	_ = v25
	var v27 int64
	_ = v27
	var v28 int64
	_ = v28
	var v32 int64
	_ = v32
	var v39 int64
	_ = v39
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
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
	v8 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+24)))
	v9 = int64(63)
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v18 = int64(32)
	v19 = int64(base.Ui64(v11) >> (uint(v18) % 64))
	v21 = int64(base.Ui64(v8) >> (uint(v18) % 64))
	v24 = int64(4294967295)
	v25 = v11 & v24
	v27 = v8 & v24
	v28 = v25 * v27
	v32 = int64(base.Ui64(v28)>>(uint(v18)%64)) + v25*v21
	v39 = v27*v19 + v32&v24
	*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = v8*(v11>>(uint(v9)%64)) + v8>>(uint(v9)%64)*v11 + v19*v21 + int64(base.Ui64(v32)>>(uint(v18)%64)) + int64(base.Ui64(v39)>>(uint(v18)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v6))) = v28&v24 | v39<<(uint(v18)%64)
	v50 = *(*int64)(unsafe.Add(mBase, uint32(v6)+8))
	v51 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
	if v50 != v51>>(uint(int64(63))%64) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v60 = m.ExcPending
		if v60 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v63 = m.ExcPending
			if v63 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_int48mul_0), int32(0))
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_int48mul_1), int32(1047), int32(_a_F_int48mul_2))
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
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
		m.G0 = v6 + int32(16)
		return v51
	}
}
func F_int4gcd(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v28 int32
	_ = v28
	var __phi28 int32
	_ = __phi28
	var v29 int32
	_ = v29
	var __phi29 int32
	_ = __phi29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v7 = int32(31)
	v8 = v5 >> (uint(v7) % 32)
	v12 = v6 >> (uint(v7) % 32)
	v15 = base.B2i32(v12-(v12^v6) < v8-(v8^v5))
	if v12-(v12^v6) < v8-(v8^v5) {
		v16 = v5
	} else {
		v16 = v6
	}
	if v12-(v12^v6) < v8-(v8^v5) {
		v17 = v6
	} else {
		v17 = v5
	}
	if v17 == int32(-2147483648) {
		if v16&int32(2147483647) == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(50331778))
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, int32(_a_F_int4gcd_0), int32(0))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_int4gcd_1), int32(1293), int32(_a_F_int4gcd_2))
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
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
			if v16 != int32(-1) {
				__phi28 = v16
				__phi29 = v17
				v28 = __phi28
				v29 = __phi29
				for {
					v32 = base.I32_rem_s(v29, v28)
					if v32 != 0 {
						__phi28 = v32
						__phi29 = v28
						v28 = __phi28
						v29 = __phi29
						continue
					} else {
						break
					}
					break
				}
				v34 = v28
				v38 = v34 >> (uint(int32(31)) % 32)
				return base.I64_extend_i32_s(v34 ^ v38 - v38)
			} else {
				return int64(1)
			}
		}
	} else {
		if v16 != 0 {
			__phi28 = v16
			__phi29 = v17
			v28 = __phi28
			v29 = __phi29
			for {
				v32 = base.I32_rem_s(v29, v28)
				if v32 != 0 {
					__phi28 = v32
					__phi29 = v28
					v28 = __phi28
					v29 = __phi29
					continue
				} else {
					break
				}
				break
			}
			v34 = v28
		} else {
			v34 = v17
		}
		v38 = v34 >> (uint(int32(31)) % 32)
		return base.I64_extend_i32_s(v34 ^ v38 - v38)
	}
}
func F_int4inc(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = v3 + int32(1)
	if v5 < v3 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_int4inc_0), int32(0))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_int4inc_1), int32(909), int32(_a_F_int4inc_2))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
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
		return base.I64_extend_i32_s(v5)
	}
}
func F_int4out(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_palloc(m, int32(12))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		if int32(0) <= v2 {
			v17 = v2
			v18 = int32(0)
		} else {
			v12 = int32(45)
			*(*uint8)(unsafe.Add(mBase, uint32(v4))) = uint8(v12)
			v17 = int32(0) - v2
			v18 = int32(1)
		}
		v20 = F_pg_ultoa_n(m, v17, v4+v18)
		mBase = m.M
		v23 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v4+(v20+v18)))) = uint8(v23)
		return base.I64_extend_i32_u(v4)
	}
}
func F_int4pl(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = v6 + v3
	if base.B2i32(v3 < int32(0)) != base.B2i32(v7 < v6) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_int4pl_0), int32(0))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_int4pl_1), int32(829), int32(_a_F_int4pl_2))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
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
		return base.I64_extend_i32_s(v7)
	}
}
func F_int4recv(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_pq_getmsgint(m, v2, int32(4))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_s(v4)
	}
}
func F_int82ge(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v3 int64
	_ = v3
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+40)))
	return base.I64_extend_i32_u(base.B2i32(v3 <= v2))
}
func F_int82lt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v3 int64
	_ = v3
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+40)))
	return base.I64_extend_i32_u(base.B2i32(v2 < v3))
}
func F_int82mul(m *base.Module, l0 int32) int64 {
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v6 = Fn14244(m, l0, int32(_a_F_int82mul_0), int32(1108), int32(_a_F_int82mul_1), int32(_a_F_int82mul_2))
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return v6
	}
}
func F_int84eq(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v3 int64
	_ = v3
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+40)))
	return base.I64_extend_i32_u(base.B2i32(v2 == v3))
}
func F_int8dec(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v6 int64
	_ = v6
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = v4 - int64(1)
	if v4 <= v6 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_int8dec_0), int32(0))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_int8dec_1), int32(750), int32(_a_F_int8dec_2))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
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
		return v6
	}
}
func F_int8inc(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v6 int64
	_ = v6
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = v4 + int64(1)
	if v6 < v4 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_int8inc_0), int32(0))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_int8inc_1), int32(736), int32(_a_F_int8inc_2))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
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
		return v6
	}
}
func F_int8lcm(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int64
	_ = v10
	var v13 int64
	_ = v13
	var v16 int64
	_ = v16
	var v17 int64
	_ = v17
	var v21 int64
	_ = v21
	var v24 int32
	_ = v24
	var v25 int64
	_ = v25
	var v26 int64
	_ = v26
	var v37 int64
	_ = v37
	var __phi37 int64
	_ = __phi37
	var v38 int64
	_ = v38
	var __phi38 int64
	_ = __phi38
	var v41 int64
	_ = v41
	var v45 int64
	_ = v45
	var v53 int64
	_ = v53
	var v54 int64
	_ = v54
	var v55 int64
	_ = v55
	var v63 int64
	_ = v63
	var v64 int64
	_ = v64
	var v66 int64
	_ = v66
	var v69 int64
	_ = v69
	var v70 int64
	_ = v70
	var v72 int64
	_ = v72
	var v73 int64
	_ = v73
	var v77 int64
	_ = v77
	var v84 int64
	_ = v84
	var v95 int64
	_ = v95
	var v96 int64
	_ = v96
	var v103 int64
	_ = v103
	var v107 int64
	_ = v107
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	v2 = int64(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if v10 == v2 {
		v107 = v2
		m.G0 = v8 + int32(16)
		return v107
	} else {
		v13 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		if v13 == int64(0) {
			v107 = v2
			m.G0 = v8 + int32(16)
			return v107
		} else {
			v16 = int64(63)
			v17 = v10 >> (uint(v16) % 64)
			v21 = v13 >> (uint(v16) % 64)
			v24 = base.B2i32(base.Ui64(v21-(v21^v13)) < base.Ui64(v17-(v17^v10)))
			if base.Ui64(v21-(v21^v13)) < base.Ui64(v17-(v17^v10)) {
				v25 = v10
			} else {
				v25 = v13
			}
			if base.Ui64(v21-(v21^v13)) < base.Ui64(v17-(v17^v10)) {
				v26 = v13
			} else {
				v26 = v10
			}
			if v26 != int64(-9223372036854775807-1) {
				__phi37 = v25
				__phi38 = v26
				v37 = __phi37
				v38 = __phi38
				for {
					v41 = base.I64_rem_s(v38, v37)
					if v41 != int64(0) {
						__phi37 = v41
						__phi38 = v37
						v37 = __phi37
						v38 = __phi38
						continue
					} else {
						break
					}
					break
				}
				v45 = v37 >> (uint(int64(63)) % 64)
				v53 = v37 ^ v45 - v45
				v54 = base.I64_div_s(v10, v53)
				v55 = int64(63)
				v63 = int64(32)
				v64 = int64(base.Ui64(v13) >> (uint(v63) % 64))
				v66 = int64(base.Ui64(v54) >> (uint(v63) % 64))
				v69 = int64(4294967295)
				v70 = v13 & v69
				v72 = v54 & v69
				v73 = v70 * v72
				v77 = int64(base.Ui64(v73)>>(uint(v63)%64)) + v70*v66
				v84 = v72*v64 + v77&v69
				*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v54*(v13>>(uint(v55)%64)) + v54>>(uint(v55)%64)*v13 + v64*v66 + int64(base.Ui64(v77)>>(uint(v63)%64)) + int64(base.Ui64(v84)>>(uint(v63)%64))
				*(*int64)(unsafe.Add(mBase, uint32(v8))) = v73&v69 | v84<<(uint(v63)%64)
				v95 = *(*int64)(unsafe.Add(mBase, uint32(v8)+8))
				v96 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
				if v95 != v96>>(uint(int64(63))%64) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v136 = m.ExcPending
					if v136 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(50331778))
						mBase = m.M
						v139 = m.ExcPending
						if v139 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_int8lcm_0), int32(0))
							mBase = m.M
							v143 = m.ExcPending
							if v143 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_int8lcm_1), int32(713), int32(_a_F_int8lcm_2))
								mBase = m.M
								v148 = m.ExcPending
								if v148 != 0 {
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
					if v96 == int64(-9223372036854775807-1) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v152 = m.ExcPending
						if v152 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(50331778))
							mBase = m.M
							v155 = m.ExcPending
							if v155 != 0 {
								return int64(0)
							} else {
								F_errmsg(m, int32(_a_F_int8lcm_0), int32(0))
								mBase = m.M
								v159 = m.ExcPending
								if v159 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_int8lcm_1), int32(719), int32(_a_F_int8lcm_2))
									mBase = m.M
									v164 = m.ExcPending
									if v164 != 0 {
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
						v103 = v96 >> (uint(int64(63)) % 64)
						v107 = v96 ^ v103 - v103
						m.G0 = v8 + int32(16)
						return v107
					}
				}
			} else {
				if v25&int64(9223372036854775807) == int64(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v120 = m.ExcPending
					if v120 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(50331778))
						mBase = m.M
						v123 = m.ExcPending
						if v123 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_int8lcm_0), int32(0))
							mBase = m.M
							v127 = m.ExcPending
							if v127 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_int8lcm_1), int32(645), int32(_a_F_int8lcm_3))
								mBase = m.M
								v132 = m.ExcPending
								if v132 != 0 {
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
					if v25 != int64(-1) {
						__phi37 = v25
						__phi38 = v26
						v37 = __phi37
						v38 = __phi38
						for {
							v41 = base.I64_rem_s(v38, v37)
							if v41 != int64(0) {
								__phi37 = v41
								__phi38 = v37
								v37 = __phi37
								v38 = __phi38
								continue
							} else {
								break
							}
							break
						}
						v45 = v37 >> (uint(int64(63)) % 64)
						v53 = v37 ^ v45 - v45
					} else {
						v53 = int64(1)
					}
					v54 = base.I64_div_s(v10, v53)
					v55 = int64(63)
					v63 = int64(32)
					v64 = int64(base.Ui64(v13) >> (uint(v63) % 64))
					v66 = int64(base.Ui64(v54) >> (uint(v63) % 64))
					v69 = int64(4294967295)
					v70 = v13 & v69
					v72 = v54 & v69
					v73 = v70 * v72
					v77 = int64(base.Ui64(v73)>>(uint(v63)%64)) + v70*v66
					v84 = v72*v64 + v77&v69
					*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v54*(v13>>(uint(v55)%64)) + v54>>(uint(v55)%64)*v13 + v64*v66 + int64(base.Ui64(v77)>>(uint(v63)%64)) + int64(base.Ui64(v84)>>(uint(v63)%64))
					*(*int64)(unsafe.Add(mBase, uint32(v8))) = v73&v69 | v84<<(uint(v63)%64)
					v95 = *(*int64)(unsafe.Add(mBase, uint32(v8)+8))
					v96 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
					if v95 != v96>>(uint(int64(63))%64) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v136 = m.ExcPending
						if v136 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(50331778))
							mBase = m.M
							v139 = m.ExcPending
							if v139 != 0 {
								return int64(0)
							} else {
								F_errmsg(m, int32(_a_F_int8lcm_0), int32(0))
								mBase = m.M
								v143 = m.ExcPending
								if v143 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_int8lcm_1), int32(713), int32(_a_F_int8lcm_2))
									mBase = m.M
									v148 = m.ExcPending
									if v148 != 0 {
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
						if v96 == int64(-9223372036854775807-1) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v152 = m.ExcPending
							if v152 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(50331778))
								mBase = m.M
								v155 = m.ExcPending
								if v155 != 0 {
									return int64(0)
								} else {
									F_errmsg(m, int32(_a_F_int8lcm_0), int32(0))
									mBase = m.M
									v159 = m.ExcPending
									if v159 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_int8lcm_1), int32(719), int32(_a_F_int8lcm_2))
										mBase = m.M
										v164 = m.ExcPending
										if v164 != 0 {
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
							v103 = v96 >> (uint(int64(63)) % 64)
							v107 = v96 ^ v103 - v103
							m.G0 = v8 + int32(16)
							return v107
						}
					}
				}
			}
		}
	}
}
func F_int8mi(m *base.Module, l0 int32) int64 {
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v6 = Fn14243(m, l0, int32(_a_F_int8mi_0), int32(492), int32(_a_F_int8mi_1), int32(_a_F_int8mi_2))
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return v6
	}
}
func F_int8mod(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	var v6 int64
	_ = v6
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int64
	_ = v31
	var v32 int64
	_ = v32
	var v33 int64
	_ = v33
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v5 = int64(1)
	v6 = v4 + v5
	if base.Ui64(v6) <= base.Ui64(v5) {
		if base.I32_wrap_i64(v6)-int32(1) != 0 {
			v33 = int64(0)
			return v33
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(33816706))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, int32(_a_F_int8mod_0), int32(0))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_int8mod_1), int32(581), int32(_a_F_int8mod_2))
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
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
	} else {
		v31 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
		v32 = base.I64_rem_s(v31, v4)
		v33 = v32
		return v33
	}
}
func F_int8pl(m *base.Module, l0 int32) int64 {
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v6 = Fn14247(m, l0, int32(_a_F_int8pl_0), int32(478), int32(_a_F_int8pl_1), int32(_a_F_int8pl_2))
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return v6
	}
}
func F_intarray_match_first(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v6 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L5
	} else {
		goto L21
	}
L2:
	;
	v7 = F_array_contains_nulls(m, l0)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = F_ArrayGetNItemsSafe(m, v11, l0+int32(16))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L5
	} else {
		goto L8
	}
L5:
	;
	return int32(0)
L6:
	;
	if v7 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	goto L4
L8:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v16 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v26 = (v19<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L11
L10:
	;
	v26 = v16
	goto L11
L11:
	;
	if int32(0) < v14 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v31 = int32(0)
	goto L15
L13:
	;
	goto L14
L14:
	;
	return int32(0)
L15:
	;
	v37 = v31 + int32(1)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0+v26+v31<<(uint(int32(2))%32))))
	if l1 == v41 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L14
L17:
	;
	return v37
L18:
	;
	goto L19
L19:
	;
	if v37 != v14 {
		v31 = v37
		goto L15
	} else {
		goto L20
	}
L20:
	;
	goto L16
L21:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L5
	} else {
		goto L22
	}
L22:
	;
	F_errmsg(m, int32(_a_F_intarray_match_first_0), int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(_a_F_intarray_match_first_1), int32(344), int32(_a_F_intarray_match_first_2))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L5
	} else {
		goto L24
	}
L24:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_intarray_push_elem(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = F_pg_detoast_datum(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v10 = F_intarray_add_elem(m, v5, v9)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int64(0)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			if v12 != v5 {
				F_pfree(m, v5)
				mBase = m.M
				v15 = m.ExcPending
				if v15 != 0 {
					return int64(0)
				} else {
					return base.I64_extend_i32_u(v10)
				}
			} else {
				return base.I64_extend_i32_u(v10)
			}
		}
	}
}
func F_interpret_function_parameter_list(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
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
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v416 int32
	_ = v416
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v480 int32
	_ = v480
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v582 int32
	_ = v582
	var v586 int32
	_ = v586
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v603 int32
	_ = v603
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v641 int32
	_ = v641
	var v645 int32
	_ = v645
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v663 int32
	_ = v663
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v677 int32
	_ = v677
	var v682 int32
	_ = v682
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v701 int32
	_ = v701
	var v705 int32
	_ = v705
	var v708 int32
	_ = v708
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v720 int32
	_ = v720
	var v724 int32
	_ = v724
	var v727 int32
	_ = v727
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v739 int32
	_ = v739
	var v762 int32
	_ = v762
	var v764 int32
	_ = v764
	var v768 int32
	_ = v768
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v784 int32
	_ = v784
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v803 int32
	_ = v803
	var v812 int32
	_ = v812
	var v849 int32
	_ = v849
	var v850 int64
	_ = v850
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v859 int32
	_ = v859
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v939 int32
	_ = v939
	v14 = int32(0)
	v38 = m.G0
	v40 = v38 - int32(80)
	m.G0 = v40
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v43 = v42
	goto L3
L2:
	;
	v43 = v14
	goto L3
L3:
	;
	v44 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l11))) = v44
	*(*int32)(unsafe.Add(mBase, uint32(l12))) = v44
	v50 = F_palloc(m, v43<<(uint(int32(2))%32))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	v53 = v43 << (uint(int32(3)) % 32)
	v54 = F_palloc(m, v53)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v56 = F_palloc(m, v53)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v58 = F_palloc0(m, v53)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v60 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l10))) = v60
	if l1 == v60 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v781 = F_buildoidvector(m, v50, v768)
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L4
	} else {
		goto L206
	}
L10:
	;
	v762 = v14
	v764 = v14
	v768 = v14
	v780 = int32(0)
	goto L9
L11:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v64 <= int32(0) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v90 = v14
	v92 = v14
	v93 = v14
	v95 = v14
	v96 = v14
	v102 = v14
	goto L23
L13:
	;
	v762 = v480
	v764 = v278
	v768 = v234
	v780 = base.B2i32(int32(0) < v295)
	goto L9
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L4
	} else {
		goto L201
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L4
	} else {
		goto L196
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L4
	} else {
		goto L191
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L4
	} else {
		goto L186
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L4
	} else {
		goto L181
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L4
	} else {
		goto L176
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L4
	} else {
		goto L171
	}
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L4
	} else {
		goto L165
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L4
	} else {
		goto L159
	}
L23:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v108+v95<<(uint(int32(2))%32))))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+12))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v112)+8))
	v116 = F_LookupTypeName(m, l0, v114, int32(0))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L4
	} else {
		goto L25
	}
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L4
	} else {
		goto L153
	}
L25:
	;
	if v116 == int32(0) {
		goto L21
	} else {
		goto L26
	}
L26:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v116)+16))
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+22)))
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120+v121)+82)))
	if v123 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L24
L28:
	;
	v153 = F_typeTypeId(m, v116)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L4
	} else {
		goto L39
	}
L29:
	;
	if base.B2i32(l2 != int32(14)) == int32(0) {
		goto L27
	} else {
		goto L30
	}
L30:
	;
	if l3 == int32(1) {
		goto L22
	} else {
		goto L31
	}
L31:
	;
	v130 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L4
	} else {
		goto L32
	}
L32:
	;
	if v130 == int32(0) {
		goto L28
	} else {
		goto L33
	}
L33:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	v137 = F_TypeNameToString(m, v114)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+64)) = v137
	F_errmsg(m, int32(_a_F_interpret_function_parameter_list_0), v40-int32(-64))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v114)+28))
	F_parser_errposition(m, l0, v145)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(_a_F_interpret_function_parameter_list_1), int32(274), int32(_a_F_interpret_function_parameter_list_2))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	goto L28
L39:
	;
	F_ReleaseCatCache(m, v116)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	v159 = *(*int32)(unsafe.Add(mBase, _c_F_interpret_function_parameter_list[0]))
	v161 = F_object_aclcheck(m, int32(1247), v153, v159, int64(256))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	if v161 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	F_aclcheck_error_type(m, v161, v153)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L4
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+12)))
	if v165 == int32(1) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	goto L44
L46:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L4
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v217 = v113 - int32(111)
	switch v217 {
	case 0, 5:
		v233 = int32(0)
		v234 = v96
		goto L65
	default:
		goto L66
	}
L49:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L4
	} else {
		goto L50
	}
L50:
	;
	if l3 != int32(29) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	F_errmsg(m, int32(_a_F_interpret_function_parameter_list_3), int32(0))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L4
	} else {
		goto L62
	}
L52:
	;
	if l3 != int32(1) {
		goto L51
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	F_errmsg(m, int32(_a_F_interpret_function_parameter_list_4), int32(0))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L4
	} else {
		goto L59
	}
L55:
	;
	F_errmsg(m, int32(_a_F_interpret_function_parameter_list_5), int32(0))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L4
	} else {
		goto L56
	}
L56:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v112)+20))
	F_parser_errposition(m, l0, v183)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L4
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(_a_F_interpret_function_parameter_list_1), int32(299), int32(_a_F_interpret_function_parameter_list_2))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L4
	} else {
		goto L58
	}
L58:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L59:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v112)+20))
	F_parser_errposition(m, l0, v195)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L4
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_interpret_function_parameter_list_1), int32(304), int32(_a_F_interpret_function_parameter_list_2))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L4
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
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v112)+20))
	F_parser_errposition(m, l0, v207)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L4
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_interpret_function_parameter_list_1), int32(309), int32(_a_F_interpret_function_parameter_list_2))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L4
	} else {
		goto L64
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L65:
	;
	if v113 == int32(100) {
		goto L70
	} else {
		goto L71
	}
L66:
	;
	if int32(0) < v93 {
		goto L20
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50+v96<<(uint(int32(2))%32)))) = v153
	v224 = int32(1)
	v226 = v96 + v224
	if l5 == int32(0) {
		v233 = v224
		v234 = v226
		goto L65
	} else {
		goto L68
	}
L68:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v230 = F_lappend_oid(m, v229, v153)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L4
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v230
	v233 = v224
	v234 = v226
	goto L65
L70:
	;
	v238 = int32(105)
	goto L72
L71:
	;
	v238 = v113
	goto L72
L72:
	;
	v240 = v238 - int32(105)
	v241 = int32(0)
	if base.B2i32(v240 == v241)|base.B2i32(v240 == int32(13)) == v241 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	if l3 == int32(29) {
		goto L78
	} else {
		goto L79
	}
L74:
	;
	v278 = v92
	goto L75
L75:
	;
	if v238 != int32(118) {
		v295 = v93
		goto L88
	} else {
		goto L89
	}
L76:
	;
	v278 = v92 + int32(1)
	goto L75
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l12))) = v272
	goto L76
L78:
	;
	if v93 <= int32(0) {
		v272 = int32(2249)
		goto L77
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	if v92 != 0 {
		goto L76
	} else {
		goto L87
	}
L81:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L4
	} else {
		goto L82
	}
L82:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L4
	} else {
		goto L83
	}
L83:
	;
	F_errmsg(m, int32(_a_F_interpret_function_parameter_list_6), int32(0))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L4
	} else {
		goto L84
	}
L84:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v112)+20))
	F_parser_errposition(m, l0, v264)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L4
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(_a_F_interpret_function_parameter_list_1), int32(341), int32(_a_F_interpret_function_parameter_list_2))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L4
	} else {
		goto L86
	}
L86:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L87:
	;
	v272 = v153
	goto L77
L88:
	;
	v297 = v95 << (uint(int32(3)) % 32)
	*(*int64)(unsafe.Add(mBase, uint32(v54+v297))) = base.I64_extend_i32_u(v153)
	*(*int64)(unsafe.Add(mBase, uint32(v56+v297))) = base.I64_extend8_s(base.I64_extend_i32_u(v238))
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
	if v305 == int32(0) {
		v480 = v90
		goto L93
	} else {
		goto L94
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l11))) = v153
	v283 = v93 + int32(1)
	if base.B2i32(v153 == int32(_a_F_interpret_function_parameter_list_7))|base.B2i32(base.Ui32(v153-int32(2276)) < base.Ui32(int32(2))) != 0 {
		v295 = v283
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v291 = F_get_element_type(m, v153)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L4
	} else {
		goto L91
	}
L91:
	;
	if v291 == int32(0) {
		goto L19
	} else {
		goto L92
	}
L92:
	;
	v295 = v283
	goto L88
L93:
	;
	if l9 != 0 {
		goto L129
	} else {
		goto L130
	}
L94:
	;
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v305))))
	if v308 == int32(0) {
		v480 = v90
		goto L93
	} else {
		goto L95
	}
L95:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v311 <= int32(0) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v456 = F_cstring_to_text(m, v305)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L4
	} else {
		goto L128
	}
L97:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v318 = int32(0)
	goto L98
L98:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v314+v318<<(uint(int32(2))%32))))
	if v356 == v112 {
		goto L96
	} else {
		goto L100
	}
L99:
	;
	goto L96
L100:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v356)+12))
	if v359 == int32(100) {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v362 = int32(105)
	goto L103
L102:
	;
	v362 = v359
	goto L103
L103:
	;
	if v240 != int32(13) {
		goto L107
	} else {
		goto L108
	}
L104:
	;
	v416 = v318 + int32(1)
	if v416 != v311 {
		v318 = v416
		goto L98
	} else {
		goto L127
	}
L105:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v356)+4))
	if v379 == int32(0) {
		goto L104
	} else {
		goto L117
	}
L106:
	;
	switch v217 {
	case 0, 5:
		goto L104
	default:
		goto L105
	}
L107:
	;
	v366 = v240
	goto L109
L108:
	;
	v366 = int32(0)
	goto L109
L109:
	;
	if v366 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	switch v362 - int32(105) {
	case 0, 13:
		goto L106
	default:
		goto L105
	case 6, 11:
		goto L104
	}
L111:
	;
	goto L112
L112:
	;
	v372 = v362 - int32(105)
	if v372 != int32(13) {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v376 = v372
	goto L115
L114:
	;
	v376 = int32(0)
	goto L115
L115:
	;
	if v376 != 0 {
		goto L105
	} else {
		goto L116
	}
L116:
	;
	goto L106
L117:
	;
	v382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v379))))
	if v382 == int32(0) {
		goto L104
	} else {
		goto L118
	}
L118:
	;
	v387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v379))))
	v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v305))))
	if base.B2i32(v387 == int32(0))|base.B2i32(v387 != v390) != 0 {
		v408 = v387
		v409 = v390
		goto L120
	} else {
		goto L121
	}
L119:
	;
	if v408-v409 == int32(0) {
		goto L18
	} else {
		goto L126
	}
L120:
	;
	goto L119
L121:
	;
	v393 = v379
	v394 = v305
	goto L122
L122:
	;
	v397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v394)+1)))
	v398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v393)+1)))
	if v398 == int32(0) {
		v408 = v398
		v409 = v397
		goto L120
	} else {
		goto L124
	}
L123:
	;
	v408 = v398
	v409 = v397
	goto L120
L124:
	;
	v401 = int32(1)
	if v398 == v397 {
		v393 = v393 + v401
		v394 = v394 + v401
		goto L122
	} else {
		goto L125
	}
L125:
	;
	goto L123
L126:
	;
	goto L104
L127:
	;
	goto L99
L128:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v58+v297))) = base.I64_extend_i32_u(v456)
	v480 = int32(1)
	goto L93
L129:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(l9)))
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
	if v499 != 0 {
		goto L132
	} else {
		goto L133
	}
L130:
	;
	goto L131
L131:
	;
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v112)+16))
	if v510 != 0 {
		goto L139
	} else {
		goto L140
	}
L132:
	;
	v503 = v499
	goto L134
L133:
	;
	v501 = F_pstrdup(m, int32(_a_F_interpret_function_parameter_list_8))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L4
	} else {
		goto L135
	}
L134:
	;
	v504 = F_makeString(m, v503)
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L4
	} else {
		goto L136
	}
L135:
	;
	v503 = v501
	goto L134
L136:
	;
	v506 = F_lappend(m, v498, v504)
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L4
	} else {
		goto L137
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l9))) = v506
	goto L131
L138:
	;
	v534 = v95 + int32(1)
	v535 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v535 <= v534 {
		goto L13
	} else {
		goto L152
	}
L139:
	;
	if v233 == int32(0) {
		goto L17
	} else {
		goto L142
	}
L140:
	;
	goto L141
L141:
	;
	if v233&v102 != 0 {
		goto L15
	} else {
		goto L150
	}
L142:
	;
	v514 = F_transformExpr(m, l0, v510, int32(31))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L4
	} else {
		goto L143
	}
L143:
	;
	v517 = F_coerce_to_specific_type(m, l0, v514, v153, int32(_a_F_interpret_function_parameter_list_9))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L4
	} else {
		goto L144
	}
L144:
	;
	F_assign_expr_collations(m, l0, v517)
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L4
	} else {
		goto L145
	}
L145:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v521 != 0 {
		goto L16
	} else {
		goto L146
	}
L146:
	;
	v522 = F_contain_var_clause(m, v517)
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L4
	} else {
		goto L147
	}
L147:
	;
	if v522 != 0 {
		goto L16
	} else {
		goto L148
	}
L148:
	;
	v524 = *(*int32)(unsafe.Add(mBase, uint32(l10)))
	v525 = F_lappend(m, v524, v517)
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L4
	} else {
		goto L149
	}
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l10))) = v525
	v532 = int32(1)
	goto L138
L150:
	;
	if v102&base.B2i32(l3 == int32(29)) != 0 {
		goto L14
	} else {
		goto L151
	}
L151:
	;
	v532 = v102
	goto L138
L152:
	;
	v90 = v480
	v92 = v278
	v93 = v295
	v95 = v534
	v96 = v234
	v102 = v532
	goto L23
L153:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L4
	} else {
		goto L154
	}
L154:
	;
	v544 = F_TypeNameToString(m, v114)
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L4
	} else {
		goto L155
	}
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v544
	F_errmsg(m, int32(_a_F_interpret_function_parameter_list_10), v40+int32(32))
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L4
	} else {
		goto L156
	}
L156:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v114)+28))
	F_parser_errposition(m, l0, v552)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L4
	} else {
		goto L157
	}
L157:
	;
	F_errfinish(m, int32(_a_F_interpret_function_parameter_list_1), int32(261), int32(_a_F_interpret_function_parameter_list_2))
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L4
	} else {
		goto L158
	}
L158:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L159:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L4
	} else {
		goto L160
	}
L160:
	;
	v567 = F_TypeNameToString(m, v114)
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L4
	} else {
		goto L161
	}
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+48)) = v567
	F_errmsg(m, int32(_a_F_interpret_function_parameter_list_11), v40+int32(48))
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L4
	} else {
		goto L162
	}
L162:
	;
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v114)+28))
	F_parser_errposition(m, l0, v575)
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L4
	} else {
		goto L163
	}
L163:
	;
	F_errfinish(m, int32(_a_F_interpret_function_parameter_list_1), int32(268), int32(_a_F_interpret_function_parameter_list_2))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L4
	} else {
		goto L164
	}
L164:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L165:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L4
	} else {
		goto L166
	}
L166:
	;
	v590 = F_TypeNameToString(m, v114)
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L4
	} else {
		goto L167
	}
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40))) = v590
	F_errmsg(m, int32(_a_F_interpret_function_parameter_list_12), v40)
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L4
	} else {
		goto L168
	}
L168:
	;
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v114)+28))
	F_parser_errposition(m, l0, v596)
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L4
	} else {
		goto L169
	}
L169:
	;
	F_errfinish(m, int32(_a_F_interpret_function_parameter_list_1), int32(285), int32(_a_F_interpret_function_parameter_list_2))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L4
	} else {
		goto L170
	}
L170:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L171:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L4
	} else {
		goto L172
	}
L172:
	;
	F_errmsg(m, int32(_a_F_interpret_function_parameter_list_13), int32(0))
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L4
	} else {
		goto L173
	}
L173:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v112)+20))
	F_parser_errposition(m, l0, v615)
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L4
	} else {
		goto L174
	}
L174:
	;
	F_errfinish(m, int32(_a_F_interpret_function_parameter_list_1), int32(320), int32(_a_F_interpret_function_parameter_list_2))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L4
	} else {
		goto L175
	}
L175:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L176:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L4
	} else {
		goto L177
	}
L177:
	;
	F_errmsg(m, int32(_a_F_interpret_function_parameter_list_14), int32(0))
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L4
	} else {
		goto L178
	}
L178:
	;
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v112)+20))
	F_parser_errposition(m, l0, v634)
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L4
	} else {
		goto L179
	}
L179:
	;
	F_errfinish(m, int32(_a_F_interpret_function_parameter_list_1), int32(367), int32(_a_F_interpret_function_parameter_list_2))
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L4
	} else {
		goto L180
	}
L180:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L181:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L4
	} else {
		goto L182
	}
L182:
	;
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+16)) = v649
	F_errmsg(m, int32(_a_F_interpret_function_parameter_list_15), v40+int32(16))
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L4
	} else {
		goto L183
	}
L183:
	;
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v112)+20))
	F_parser_errposition(m, l0, v656)
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L4
	} else {
		goto L184
	}
L184:
	;
	F_errfinish(m, int32(_a_F_interpret_function_parameter_list_1), int32(414), int32(_a_F_interpret_function_parameter_list_2))
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L4
	} else {
		goto L185
	}
L185:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L186:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L4
	} else {
		goto L187
	}
L187:
	;
	F_errmsg(m, int32(_a_F_interpret_function_parameter_list_16), int32(0))
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L4
	} else {
		goto L188
	}
L188:
	;
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v112)+20))
	F_parser_errposition(m, l0, v675)
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L4
	} else {
		goto L189
	}
L189:
	;
	F_errfinish(m, int32(_a_F_interpret_function_parameter_list_1), int32(432), int32(_a_F_interpret_function_parameter_list_2))
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L4
	} else {
		goto L190
	}
L190:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L191:
	;
	F_errcode(m, int32(_a_F_interpret_function_parameter_list_17))
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L4
	} else {
		goto L192
	}
L192:
	;
	F_errmsg(m, int32(_a_F_interpret_function_parameter_list_18), int32(0))
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L4
	} else {
		goto L193
	}
L193:
	;
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v112)+20))
	F_parser_errposition(m, l0, v694)
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L4
	} else {
		goto L194
	}
L194:
	;
	F_errfinish(m, int32(_a_F_interpret_function_parameter_list_1), int32(448), int32(_a_F_interpret_function_parameter_list_2))
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L4
	} else {
		goto L195
	}
L195:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L196:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L4
	} else {
		goto L197
	}
L197:
	;
	F_errmsg(m, int32(_a_F_interpret_function_parameter_list_19), int32(0))
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L4
	} else {
		goto L198
	}
L198:
	;
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v112)+20))
	F_parser_errposition(m, l0, v713)
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L4
	} else {
		goto L199
	}
L199:
	;
	F_errfinish(m, int32(_a_F_interpret_function_parameter_list_1), int32(473), int32(_a_F_interpret_function_parameter_list_2))
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L4
	} else {
		goto L200
	}
L200:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L201:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L4
	} else {
		goto L202
	}
L202:
	;
	F_errmsg(m, int32(_a_F_interpret_function_parameter_list_20), int32(0))
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L4
	} else {
		goto L203
	}
L203:
	;
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v112)+20))
	F_parser_errposition(m, l0, v732)
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L4
	} else {
		goto L204
	}
L204:
	;
	F_errfinish(m, int32(_a_F_interpret_function_parameter_list_1), int32(484), int32(_a_F_interpret_function_parameter_list_2))
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L4
	} else {
		goto L205
	}
L205:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L206:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v781
	v784 = int32(0)
	if base.B2i32(v780 == v784)&base.B2i32(v764 <= v784) == v784 {
		goto L208
	} else {
		goto L209
	}
L207:
	;
	if v762 != 0 {
		goto L214
	} else {
		goto L215
	}
L208:
	;
	v792 = F_construct_array_builtin(m, v54, v43, int32(26))
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L4
	} else {
		goto L211
	}
L209:
	;
	goto L210
L210:
	;
	v803 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v803
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v803
	goto L207
L211:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v792
	v796 = F_construct_array_builtin(m, v56, v43, int32(18))
	mBase = m.M
	v797 = m.ExcPending
	if v797 != 0 {
		goto L4
	} else {
		goto L212
	}
L212:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v796
	if v764 < int32(2) {
		goto L207
	} else {
		goto L213
	}
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l12))) = int32(2249)
	goto L207
L214:
	;
	if int32(0) < v43 {
		goto L217
	} else {
		goto L218
	}
L215:
	;
	v939 = int32(0)
	goto L216
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l8))) = v939
	m.G0 = v40 + int32(80)
	return
L217:
	;
	v812 = int32(0)
	goto L220
L218:
	;
	goto L219
L219:
	;
	v899 = F_construct_array_builtin(m, v58, v43, int32(25))
	mBase = m.M
	v900 = m.ExcPending
	if v900 != 0 {
		goto L4
	} else {
		goto L227
	}
L220:
	;
	v849 = v58 + v812<<(uint(int32(3))%32)
	v850 = *(*int64)(unsafe.Add(mBase, uint32(v849)))
	if v850 == int64(0) {
		goto L222
	} else {
		goto L223
	}
L221:
	;
	goto L219
L222:
	;
	v854 = F_cstring_to_text(m, int32(_a_F_interpret_function_parameter_list_8))
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
		goto L4
	} else {
		goto L225
	}
L223:
	;
	goto L224
L224:
	;
	v859 = v812 + int32(1)
	if v859 != v43 {
		v812 = v859
		goto L220
	} else {
		goto L226
	}
L225:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v849))) = base.I64_extend_i32_u(v854)
	goto L224
L226:
	;
	goto L221
L227:
	;
	v939 = v899
	goto L216
}
func F_inv_write(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int64
	_ = v25
	var v31 int32
	_ = v31
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int64
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v93 int64
	_ = v93
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v110 int64
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v179 int32
	_ = v179
	var v180 int64
	_ = v180
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v211 int32
	_ = v211
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v231 int64
	_ = v231
	var v235 int32
	_ = v235
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v292 int64
	_ = v292
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int64
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v331 int32
	_ = v331
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v348 int32
	_ = v348
	var v359 int64
	_ = v359
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v405 int32
	_ = v405
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v433 int32
	_ = v433
	var v438 int32
	_ = v438
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v455 int32
	_ = v455
	v4 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(2256)
	m.G0 = v23
	v25 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	base.MemoryFill(m, v23+int32(92), v4, int32(2052))
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v31&int32(2) != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if int32(0) < l2 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L18
	} else {
		goto L110
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L18
	} else {
		goto L106
	}
L5:
	;
	if base.Ui64(int64(4398046509057)) <= base.Ui64(v25+base.I64_extend_i32_u(l2)) {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	v405 = v4
	goto L7
L7:
	;
	m.G0 = v23 + int32(2256)
	return v405
L8:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_inv_write[0]))
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_inv_write[1]))
	if v44 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v45 = v41
	goto L11
L10:
	;
	v45 = int32(0)
	goto L11
L11:
	;
	if v45 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v48 = int32(_a_F_inv_write_0)
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_inv_write[2]))
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_inv_write[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_inv_write[2])) = v52
	if v41 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v79 = v41
	goto L14
L14:
	;
	v82 = int64(base.Ui64(v25) >> (uint(int64(11)) % 64))
	v85 = v23 + int32(96)
	v86 = F_CatalogOpenIndexes(m, v79)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L18
	} else {
		goto L24
	}
L15:
	;
	v64 = v41
	v65 = v44
	goto L17
L16:
	;
	v57 = F_table_open(m, int32(2613), int32(3))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if v65 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	return int32(0)
L19:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_inv_write[0])) = v57
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_inv_write[1]))
	v64 = v57
	v65 = v63
	goto L17
L20:
	;
	v71 = F_index_open(m, int32(2683), int32(3))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L18
	} else {
		goto L23
	}
L21:
	;
	v76 = v64
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_inv_write[2])) = v49
	v79 = v76
	goto L14
L23:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_inv_write[1])) = v71
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_inv_write[0]))
	v76 = v75
	goto L22
L24:
	;
	v89 = v23 + int32(2144)
	v93 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0))))
	F_ScanKeyInit(m, v89, int32(1), int32(3), int32(184), v93)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L18
	} else {
		goto L25
	}
L25:
	;
	F_ScanKeyInit(m, v23+int32(2200), int32(2), int32(4), int32(150), base.I64_extend32_s(v82))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L18
	} else {
		goto L26
	}
L26:
	;
	v105 = v23 + int32(100)
	v107 = v23 + int32(92)
	v110 = base.I64_extend_i32_u(v107)
	v112 = *(*int32)(unsafe.Add(mBase, _c_F_inv_write[0]))
	v114 = *(*int32)(unsafe.Add(mBase, _c_F_inv_write[1]))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v117 = F_systable_beginscan_ordered(m, v112, v114, v115, int32(2), v89)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L18
	} else {
		goto L27
	}
L27:
	;
	v127 = v4
	v129 = int32(0)
	v131 = int32(1)
	v132 = v4
	v133 = base.I32_wrap_i64(v82)
	goto L28
L28:
	;
	if v131 != 0 {
		goto L36
	} else {
		goto L37
	}
L29:
	;
	F_systable_endscan_ordered(m, v117)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L18
	} else {
		goto L103
	}
L30:
	;
	F_pfree(m, v378)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L18
	} else {
		goto L101
	}
L31:
	;
	v308 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v309 = base.I32_wrap_i64(v308)
	v311 = v309 & int32(2047)
	if v311 == int32(0) {
		goto L84
	} else {
		goto L85
	}
L32:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L18
	} else {
		goto L80
	}
L33:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L18
	} else {
		goto L77
	}
L34:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v154)+4))
	if v156 != v133 {
		v306 = v154
		v307 = v155
		goto L31
	} else {
		goto L43
	}
L35:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v142)+16))
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147)+20)))
	if v148&int32(1) != 0 {
		goto L33
	} else {
		goto L42
	}
L36:
	;
	v142 = F_systable_getnext_ordered(m, v117, int32(1))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L18
	} else {
		goto L39
	}
L37:
	;
	v145 = v132
	goto L38
L38:
	;
	if v127 != 0 {
		v154 = v127
		v155 = v145
		goto L34
	} else {
		goto L41
	}
L39:
	;
	if v142 != 0 {
		goto L35
	} else {
		goto L40
	}
L40:
	;
	v145 = int32(0)
	goto L38
L41:
	;
	v306 = int32(0)
	v307 = v145
	goto L31
L42:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147)+22)))
	v154 = v147 + v151
	v155 = v142
	goto L34
L43:
	;
	v159 = v154 + int32(8)
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+8)))
	v162 = v160 & int32(3)
	if v162 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v163 = F_detoast_attr(m, v159)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L18
	} else {
		goto L47
	}
L45:
	;
	v165 = v159
	goto L46
L46:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	v168 = int32(base.Ui32(v166) >> (uint(int32(2)) % 32))
	v170 = v168 - int32(4)
	if base.Ui32(v168-int32(2053)) <= base.Ui32(int32(-2050)) {
		goto L32
	} else {
		goto L48
	}
L47:
	;
	v165 = v163
	goto L46
L48:
	;
	if v170 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	base.MemoryCopy(m, v85, v165+int32(4), v170)
	goto L51
L50:
	;
	goto L51
L51:
	;
	if v162 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	F_pfree(m, v165)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L18
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v180 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v183 = base.I32_wrap_i64(v180) & int32(2047)
	if v183 <= v170 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	goto L54
L56:
	;
	v221 = int32(2048) - v183
	v222 = l2 - v129
	if v221 < v222 {
		goto L66
	} else {
		goto L67
	}
L57:
	;
	v187 = v23 + int32(92) + v168
	v188 = int32(3)
	v190 = v183 - v170
	if v187&v188|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v190))|v190&v188 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	if base.Ui32(v183+v85) <= base.Ui32(v187) {
		goto L56
	} else {
		goto L61
	}
L59:
	;
	v211 = v190
	goto L60
L60:
	;
	if v211 == int32(0) {
		goto L56
	} else {
		goto L65
	}
L61:
	;
	v201 = v85 + v168
	v202 = v183 + v85
	if base.Ui32(v202) < base.Ui32(v201) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v204 = v201
	goto L64
L63:
	;
	v204 = v202
	goto L64
L64:
	;
	v211 = (v204+(v107^int32(-1))-v168)&int32(-4) + int32(4)
	goto L60
L65:
	;
	base.MemoryFill(m, v187, int32(0), v211)
	goto L56
L66:
	;
	v224 = v221
	goto L68
L67:
	;
	v224 = v222
	goto L68
L68:
	;
	if v224 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	base.MemoryCopy(m, v183+v85, l1+v129, v224)
	goto L71
L70:
	;
	goto L71
L71:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v180 + base.I64_extend_i32_s(v224)
	v231 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+64)) = v231
	*(*int64)(unsafe.Add(mBase, uint32(v23)+72)) = v231
	v235 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+60)) = uint16(v235)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+62)) = uint8(v235)
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+56)) = uint16(v235)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+80)) = v110
	v243 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+58)) = uint8(v243)
	v246 = v224 + v183
	if v246 < v170 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v248 = v170
	goto L74
L73:
	;
	v248 = v246
	goto L74
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+92)) = v248<<(uint(int32(2))%32) + int32(16)
	v255 = *(*int32)(unsafe.Add(mBase, _c_F_inv_write[0]))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v255)+52))
	v263 = F_heap_modify_tuple(m, v155, v256, v23-int32(-64), v23+int32(60), v23+int32(56))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L18
	} else {
		goto L75
	}
L75:
	;
	v266 = *(*int32)(unsafe.Add(mBase, _c_F_inv_write[0]))
	F_CatalogTupleUpdateWithInfo(m, v266, v263+int32(4), v263, v86)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L18
	} else {
		goto L76
	}
L76:
	;
	v377 = v224
	v378 = v263
	v379 = v235
	v381 = v243
	v382 = int32(0)
	goto L30
L77:
	;
	F_errmsg_internal(m, int32(_a_F_inv_write_1), int32(0))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L18
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(_a_F_inv_write_2), int32(622), int32(_a_F_inv_write_3))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L18
	} else {
		goto L79
	}
L79:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L80:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L18
	} else {
		goto L81
	}
L81:
	;
	v292 = *(*int64)(unsafe.Add(mBase, uint32(v154)))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+40)) = v170
	*(*int64)(unsafe.Add(mBase, uint32(v23)+32)) = v292
	F_errmsg(m, int32(_a_F_inv_write_4), v23+int32(32))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L18
	} else {
		goto L82
	}
L82:
	;
	F_errfinish(m, int32(_a_F_inv_write_2), int32(153), int32(_a_F_inv_write_5))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L18
	} else {
		goto L83
	}
L83:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L84:
	;
	v338 = int32(2048) - v311
	v339 = l2 - v129
	if v338 < v339 {
		goto L93
	} else {
		goto L94
	}
L85:
	;
	if v309&int32(3)|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v311)) == int32(0) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v321 = v311 + v85
	if base.Ui32(v105) < base.Ui32(v321) {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	v331 = v311
	goto L88
L88:
	;
	if v331 == int32(0) {
		goto L84
	} else {
		goto L92
	}
L89:
	;
	v323 = v321
	goto L91
L90:
	;
	v323 = v105
	goto L91
L91:
	;
	v331 = (v323-v23-int32(97))&int32(-4) + int32(4)
	goto L88
L92:
	;
	base.MemoryFill(m, v85, int32(0), v331)
	goto L84
L93:
	;
	v341 = v338
	goto L95
L94:
	;
	v341 = v339
	goto L95
L95:
	;
	if v341 != 0 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	base.MemoryCopy(m, v311+v85, l1+v129, v341)
	goto L98
L97:
	;
	goto L98
L98:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v308 + base.I64_extend_i32_s(v341)
	v348 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+60)) = uint16(v348)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+92)) = (v341+v311)<<(uint(int32(2))%32) + int32(16)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+62)) = uint8(v348)
	v359 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0))))
	*(*int64)(unsafe.Add(mBase, uint32(v23)+80)) = v110
	*(*int64)(unsafe.Add(mBase, uint32(v23)+72)) = base.I64_extend_i32_s(v133)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+64)) = v359
	v365 = *(*int32)(unsafe.Add(mBase, _c_F_inv_write[0]))
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v365)+52))
	v371 = F_heap_form_tuple(m, v366, v23-int32(-64), v23+int32(60))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L18
	} else {
		goto L99
	}
L99:
	;
	v374 = *(*int32)(unsafe.Add(mBase, _c_F_inv_write[0]))
	F_CatalogTupleInsertWithInfo(m, v374, v371, v86)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L18
	} else {
		goto L100
	}
L100:
	;
	v377 = v341
	v378 = v371
	v379 = v306
	v381 = v348
	v382 = v307
	goto L30
L101:
	;
	v389 = v377 + v129
	if v389 < l2 {
		v127 = v379
		v129 = v389
		v131 = v381
		v132 = v382
		v133 = v133 + int32(1)
		goto L28
	} else {
		goto L102
	}
L102:
	;
	goto L29
L103:
	;
	F_CatalogCloseIndexes(m, v86)
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L18
	} else {
		goto L104
	}
L104:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L18
	} else {
		goto L105
	}
L105:
	;
	v405 = v389
	goto L7
L106:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L18
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = l2
	F_errmsg(m, int32(_a_F_inv_write_6), v23+int32(16))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L18
	} else {
		goto L108
	}
L108:
	;
	F_errfinish(m, int32(_a_F_inv_write_2), int32(588), int32(_a_F_inv_write_3))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L18
	} else {
		goto L109
	}
L109:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L110:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L18
	} else {
		goto L111
	}
L111:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v446
	F_errmsg(m, int32(_a_F_inv_write_7), v23)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L18
	} else {
		goto L112
	}
L112:
	;
	F_errfinish(m, int32(_a_F_inv_write_2), int32(578), int32(_a_F_inv_write_3))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L18
	} else {
		goto L113
	}
L113:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_is_admin_of_role(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14314(m, l0, l1, int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_is_code_in_table(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	v2 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = l0
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_is_code_in_table[0]))
	if base.Ui32(l0) < base.Ui32(v10) {
		v27 = v2
		m.G0 = v6 + int32(16)
		return v27
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, _c_F_is_code_in_table[1]))
		if base.Ui32(v13) < base.Ui32(l0) {
			v27 = v2
			m.G0 = v6 + int32(16)
			return v27
		} else {
			v21 = F_bsearch(m, v6+int32(12), int32(_a_F_is_code_in_table_0), int32(34), int32(8), int32(_a_F_is_code_in_table_1))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v27 = base.B2i32(v21 != int32(0))
				m.G0 = v6 + int32(16)
				return v27
			}
		}
	}
}
func F_is_dummy_rel(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v4 == v2 {
		v25 = v2
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
		v8 = v7
		for {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
			if base.Ui32(int32(2)) <= base.Ui32(v12-int32(303)) {
				break
			} else {
				v8 = v11 + int32(72)
				continue
			}
			break
		}
		if v12 != int32(293) {
			v25 = v2
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v11)+72))
			if v19 != 0 {
				v25 = v2
			} else {
				v25 = int32(1)
			}
		}
	}
	return v25
}
func F_is_leap(m *base.Module, l0 int32) int32 {
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	if int32(2147481747) < l0 {
		v6 = l0 - int32(2000)
	} else {
		v6 = l0
	}
	if v6&int32(3) != 0 {
		return int32(0)
	} else {
		v12 = v6 + int32(1900)
		v14 = base.I32_rem_s(v12, int32(100))
		if v14 != 0 {
			return int32(1)
		} else {
			v18 = base.I32_rem_s(v12, int32(400))
			return base.B2i32(v18 == int32(0))
		}
	}
}
func F_ismn_in(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14257(m, l0, int32(4))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_iso8859_1_to_utf8(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int64
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	v8 = *(*int64)(unsafe.Add(mBase, uint32(l0)+104))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	F_check_encoding_conversion_args(m, v11, v12, v13, int32(8), int32(6))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	if v13 <= int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v68 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v65))) = uint8(v68)
	return base.I64_extend_i32_s(v63 - v10)
L4:
	;
	v63 = v10
	v65 = v9
	goto L3
L5:
	;
	goto L6
L6:
	;
	v23 = v9
	v24 = v10
	v25 = v13
	goto L7
L7:
	;
	v29 = int32(*(*int8)(unsafe.Add(mBase, uint32(v24))))
	if v29 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v63 = v56
	v65 = v53
	goto L3
L9:
	;
	if v8 != int64(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	if int32(0) <= v29 {
		goto L16
	} else {
		goto L17
	}
L12:
	;
	v63 = v24
	v65 = v23
	goto L3
L13:
	;
	goto L14
L14:
	;
	F_report_invalid_encoding(m, int32(8), v24, v25)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L16:
	;
	v52 = v29
	v53 = v23 + int32(1)
	goto L18
L17:
	;
	v42 = v29 & int32(191)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)) = uint8(v42)
	v52 = int32(base.Ui32(v29&int32(192))>>(uint(int32(6))%32)) | int32(-64)
	v53 = v23 + int32(2)
	goto L18
L18:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v23))) = uint8(v52)
	v55 = int32(1)
	v56 = v24 + v55
	if v55 < v25 {
		v23 = v53
		v24 = v56
		v25 = v25 - v55
		goto L7
	} else {
		goto L19
	}
L19:
	;
	goto L8
}
func F_issn_cast_from_ean13(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14315(m, l0, int32(5))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_iswalpha(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	if base.Ui32(l0) <= base.Ui32(int32(_a_F_iswalpha_0)) {
		v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(l0)>>(uint(int32(8))%32)))+uint32(_c_F_iswalpha[0]))))
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(l0)>>(uint(int32(3))%32))&int32(31)|v10<<(uint(int32(5))%32))+uint32(_c_F_iswalpha[0]))))
		return int32(base.Ui32(v14)>>(uint(l0&int32(7))%32)) & int32(1)
	} else {
		return base.B2i32(base.Ui32(l0) < base.Ui32(int32(_a_F_iswalpha_1)))
	}
}
func F_italian_UTF_8_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v197 int32
	_ = v197
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v223 int32
	_ = v223
	var v235 int32
	_ = v235
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v328 int32
	_ = v328
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v354 int32
	_ = v354
	var v366 int32
	_ = v366
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v467 int32
	_ = v467
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v493 int32
	_ = v493
	var v505 int32
	_ = v505
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v571 int32
	_ = v571
	var v575 int32
	_ = v575
	var v592 int32
	_ = v592
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v613 int32
	_ = v613
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v622 int32
	_ = v622
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v638 int32
	_ = v638
	var v651 int32
	_ = v651
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v671 int32
	_ = v671
	var v677 int32
	_ = v677
	var v689 int32
	_ = v689
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v726 int32
	_ = v726
	var v728 int32
	_ = v728
	var v732 int32
	_ = v732
	var v735 int32
	_ = v735
	var v737 int32
	_ = v737
	var v741 int32
	_ = v741
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v757 int32
	_ = v757
	var v770 int32
	_ = v770
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v790 int32
	_ = v790
	var v796 int32
	_ = v796
	var v807 int32
	_ = v807
	var v814 int32
	_ = v814
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v836 int32
	_ = v836
	var v843 int32
	_ = v843
	var v845 int32
	_ = v845
	var v849 int32
	_ = v849
	var v852 int32
	_ = v852
	var v854 int32
	_ = v854
	var v858 int32
	_ = v858
	var v868 int32
	_ = v868
	var v870 int32
	_ = v870
	var v874 int32
	_ = v874
	var v887 int32
	_ = v887
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v907 int32
	_ = v907
	var v913 int32
	_ = v913
	var v920 int32
	_ = v920
	var v931 int32
	_ = v931
	var v934 int32
	_ = v934
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v966 int32
	_ = v966
	var v968 int32
	_ = v968
	var v972 int32
	_ = v972
	var v975 int32
	_ = v975
	var v977 int32
	_ = v977
	var v981 int32
	_ = v981
	var v991 int32
	_ = v991
	var v993 int32
	_ = v993
	var v997 int32
	_ = v997
	var v1010 int32
	_ = v1010
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1030 int32
	_ = v1030
	var v1036 int32
	_ = v1036
	var v1048 int32
	_ = v1048
	var v1055 int32
	_ = v1055
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1077 int32
	_ = v1077
	var v1084 int32
	_ = v1084
	var v1086 int32
	_ = v1086
	var v1090 int32
	_ = v1090
	var v1093 int32
	_ = v1093
	var v1095 int32
	_ = v1095
	var v1099 int32
	_ = v1099
	var v1109 int32
	_ = v1109
	var v1111 int32
	_ = v1111
	var v1115 int32
	_ = v1115
	var v1128 int32
	_ = v1128
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1148 int32
	_ = v1148
	var v1154 int32
	_ = v1154
	var v1162 int32
	_ = v1162
	var v1173 int32
	_ = v1173
	var v1176 int32
	_ = v1176
	var v1181 int32
	_ = v1181
	var v1183 int32
	_ = v1183
	var v1185 int32
	_ = v1185
	var v1189 int32
	_ = v1189
	var v1191 int32
	_ = v1191
	var v1195 int32
	_ = v1195
	var v1196 int32
	_ = v1196
	var v1210 int32
	_ = v1210
	var v1211 int32
	_ = v1211
	var v1226 int32
	_ = v1226
	var v1228 int32
	_ = v1228
	var v1232 int32
	_ = v1232
	var v1235 int32
	_ = v1235
	var v1237 int32
	_ = v1237
	var v1241 int32
	_ = v1241
	var v1251 int32
	_ = v1251
	var v1253 int32
	_ = v1253
	var v1257 int32
	_ = v1257
	var v1270 int32
	_ = v1270
	var v1285 int32
	_ = v1285
	var v1286 int32
	_ = v1286
	var v1290 int32
	_ = v1290
	var v1296 int32
	_ = v1296
	var v1307 int32
	_ = v1307
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1328 int32
	_ = v1328
	var v1329 int32
	_ = v1329
	var v1344 int32
	_ = v1344
	var v1346 int32
	_ = v1346
	var v1350 int32
	_ = v1350
	var v1353 int32
	_ = v1353
	var v1355 int32
	_ = v1355
	var v1359 int32
	_ = v1359
	var v1369 int32
	_ = v1369
	var v1371 int32
	_ = v1371
	var v1375 int32
	_ = v1375
	var v1388 int32
	_ = v1388
	var v1403 int32
	_ = v1403
	var v1404 int32
	_ = v1404
	var v1408 int32
	_ = v1408
	var v1414 int32
	_ = v1414
	var v1425 int32
	_ = v1425
	var v1432 int32
	_ = v1432
	var v1444 int32
	_ = v1444
	var v1445 int32
	_ = v1445
	var v1446 int32
	_ = v1446
	var v1454 int32
	_ = v1454
	var v1461 int32
	_ = v1461
	var v1463 int32
	_ = v1463
	var v1467 int32
	_ = v1467
	var v1470 int32
	_ = v1470
	var v1472 int32
	_ = v1472
	var v1476 int32
	_ = v1476
	var v1486 int32
	_ = v1486
	var v1488 int32
	_ = v1488
	var v1492 int32
	_ = v1492
	var v1505 int32
	_ = v1505
	var v1520 int32
	_ = v1520
	var v1521 int32
	_ = v1521
	var v1525 int32
	_ = v1525
	var v1531 int32
	_ = v1531
	var v1538 int32
	_ = v1538
	var v1549 int32
	_ = v1549
	var v1552 int32
	_ = v1552
	var v1568 int32
	_ = v1568
	var v1569 int32
	_ = v1569
	var v1584 int32
	_ = v1584
	var v1586 int32
	_ = v1586
	var v1590 int32
	_ = v1590
	var v1593 int32
	_ = v1593
	var v1595 int32
	_ = v1595
	var v1599 int32
	_ = v1599
	var v1609 int32
	_ = v1609
	var v1611 int32
	_ = v1611
	var v1615 int32
	_ = v1615
	var v1628 int32
	_ = v1628
	var v1643 int32
	_ = v1643
	var v1644 int32
	_ = v1644
	var v1648 int32
	_ = v1648
	var v1654 int32
	_ = v1654
	var v1666 int32
	_ = v1666
	var v1673 int32
	_ = v1673
	var v1674 int32
	_ = v1674
	var v1675 int32
	_ = v1675
	var v1676 int32
	_ = v1676
	var v1683 int32
	_ = v1683
	var v1685 int32
	_ = v1685
	var v1690 int32
	_ = v1690
	var v1692 int32
	_ = v1692
	var v1699 int32
	_ = v1699
	var v1702 int32
	_ = v1702
	var v1706 int32
	_ = v1706
	var v1713 int32
	_ = v1713
	var v1714 int32
	_ = v1714
	var v1728 int32
	_ = v1728
	var v1731 int32
	_ = v1731
	var v1749 int32
	_ = v1749
	var v1750 int32
	_ = v1750
	var v1758 int32
	_ = v1758
	var v1765 int32
	_ = v1765
	var v1767 int32
	_ = v1767
	var v1771 int32
	_ = v1771
	var v1774 int32
	_ = v1774
	var v1776 int32
	_ = v1776
	var v1780 int32
	_ = v1780
	var v1790 int32
	_ = v1790
	var v1792 int32
	_ = v1792
	var v1796 int32
	_ = v1796
	var v1809 int32
	_ = v1809
	var v1824 int32
	_ = v1824
	var v1825 int32
	_ = v1825
	var v1829 int32
	_ = v1829
	var v1835 int32
	_ = v1835
	var v1842 int32
	_ = v1842
	var v1853 int32
	_ = v1853
	var v1856 int32
	_ = v1856
	var v1857 int32
	_ = v1857
	var v1871 int32
	_ = v1871
	var v1872 int32
	_ = v1872
	var v1880 int32
	_ = v1880
	var v1887 int32
	_ = v1887
	var v1889 int32
	_ = v1889
	var v1893 int32
	_ = v1893
	var v1896 int32
	_ = v1896
	var v1898 int32
	_ = v1898
	var v1902 int32
	_ = v1902
	var v1912 int32
	_ = v1912
	var v1914 int32
	_ = v1914
	var v1918 int32
	_ = v1918
	var v1931 int32
	_ = v1931
	var v1946 int32
	_ = v1946
	var v1947 int32
	_ = v1947
	var v1951 int32
	_ = v1951
	var v1957 int32
	_ = v1957
	var v1965 int32
	_ = v1965
	var v1976 int32
	_ = v1976
	var v1979 int32
	_ = v1979
	var v1980 int32
	_ = v1980
	var v1995 int32
	_ = v1995
	var v1996 int32
	_ = v1996
	var v2004 int32
	_ = v2004
	var v2011 int32
	_ = v2011
	var v2013 int32
	_ = v2013
	var v2017 int32
	_ = v2017
	var v2020 int32
	_ = v2020
	var v2022 int32
	_ = v2022
	var v2026 int32
	_ = v2026
	var v2036 int32
	_ = v2036
	var v2038 int32
	_ = v2038
	var v2042 int32
	_ = v2042
	var v2055 int32
	_ = v2055
	var v2070 int32
	_ = v2070
	var v2071 int32
	_ = v2071
	var v2075 int32
	_ = v2075
	var v2081 int32
	_ = v2081
	var v2088 int32
	_ = v2088
	var v2099 int32
	_ = v2099
	var v2102 int32
	_ = v2102
	var v2103 int32
	_ = v2103
	var v2117 int32
	_ = v2117
	var v2118 int32
	_ = v2118
	var v2126 int32
	_ = v2126
	var v2133 int32
	_ = v2133
	var v2135 int32
	_ = v2135
	var v2139 int32
	_ = v2139
	var v2142 int32
	_ = v2142
	var v2144 int32
	_ = v2144
	var v2148 int32
	_ = v2148
	var v2158 int32
	_ = v2158
	var v2160 int32
	_ = v2160
	var v2164 int32
	_ = v2164
	var v2177 int32
	_ = v2177
	var v2192 int32
	_ = v2192
	var v2193 int32
	_ = v2193
	var v2197 int32
	_ = v2197
	var v2203 int32
	_ = v2203
	var v2211 int32
	_ = v2211
	var v2222 int32
	_ = v2222
	var v2225 int32
	_ = v2225
	var v2230 int32
	_ = v2230
	var v2234 int32
	_ = v2234
	var v2236 int32
	_ = v2236
	var v2238 int32
	_ = v2238
	var v2253 int32
	_ = v2253
	var v2254 int32
	_ = v2254
	var v2257 int32
	_ = v2257
	var v2260 int32
	_ = v2260
	var v2261 int32
	_ = v2261
	var v2263 int32
	_ = v2263
	var v2265 int32
	_ = v2265
	var v2271 int32
	_ = v2271
	var v2272 int32
	_ = v2272
	var v2275 int32
	_ = v2275
	var v2276 int32
	_ = v2276
	var v2280 int32
	_ = v2280
	var v2285 int32
	_ = v2285
	var v2286 int32
	_ = v2286
	var v2290 int32
	_ = v2290
	var v2296 int32
	_ = v2296
	var v2297 int32
	_ = v2297
	var v2300 int32
	_ = v2300
	var v2304 int32
	_ = v2304
	var v2306 int32
	_ = v2306
	var v2309 int32
	_ = v2309
	var v2311 int32
	_ = v2311
	var v2314 int32
	_ = v2314
	var v2316 int32
	_ = v2316
	var v2318 int32
	_ = v2318
	var v2321 int32
	_ = v2321
	var v2324 int32
	_ = v2324
	var v2327 int32
	_ = v2327
	var v2331 int32
	_ = v2331
	var v2334 int32
	_ = v2334
	var v2336 int32
	_ = v2336
	var v2338 int32
	_ = v2338
	var v2341 int32
	_ = v2341
	var v2345 int32
	_ = v2345
	var v2346 int32
	_ = v2346
	var v2349 int32
	_ = v2349
	var v2353 int32
	_ = v2353
	var v2354 int32
	_ = v2354
	var v2357 int32
	_ = v2357
	var v2361 int32
	_ = v2361
	var v2362 int32
	_ = v2362
	var v2365 int32
	_ = v2365
	var v2367 int32
	_ = v2367
	var v2370 int32
	_ = v2370
	var v2372 int32
	_ = v2372
	var v2375 int32
	_ = v2375
	var v2378 int32
	_ = v2378
	var v2379 int32
	_ = v2379
	var v2381 int32
	_ = v2381
	var v2383 int32
	_ = v2383
	var v2398 int32
	_ = v2398
	var v2399 int32
	_ = v2399
	var v2402 int32
	_ = v2402
	var v2404 int32
	_ = v2404
	var v2406 int32
	_ = v2406
	var v2411 int32
	_ = v2411
	var v2413 int32
	_ = v2413
	var v2415 int32
	_ = v2415
	var v2418 int32
	_ = v2418
	var v2421 int32
	_ = v2421
	var v2424 int32
	_ = v2424
	var v2428 int32
	_ = v2428
	var v2431 int32
	_ = v2431
	var v2433 int32
	_ = v2433
	var v2435 int32
	_ = v2435
	var v2438 int32
	_ = v2438
	var v2440 int32
	_ = v2440
	var v2443 int32
	_ = v2443
	var v2446 int32
	_ = v2446
	var v2447 int32
	_ = v2447
	var v2449 int32
	_ = v2449
	var v2451 int32
	_ = v2451
	var v2466 int32
	_ = v2466
	var v2467 int32
	_ = v2467
	var v2470 int32
	_ = v2470
	var v2472 int32
	_ = v2472
	var v2474 int32
	_ = v2474
	var v2477 int32
	_ = v2477
	var v2479 int32
	_ = v2479
	var v2482 int32
	_ = v2482
	var v2484 int32
	_ = v2484
	var v2486 int32
	_ = v2486
	var v2489 int32
	_ = v2489
	var v2492 int32
	_ = v2492
	var v2495 int32
	_ = v2495
	var v2499 int32
	_ = v2499
	var v2502 int32
	_ = v2502
	var v2504 int32
	_ = v2504
	var v2506 int32
	_ = v2506
	var v2509 int32
	_ = v2509
	var v2511 int32
	_ = v2511
	var v2513 int32
	_ = v2513
	var v2516 int32
	_ = v2516
	var v2519 int32
	_ = v2519
	var v2522 int32
	_ = v2522
	var v2526 int32
	_ = v2526
	var v2529 int32
	_ = v2529
	var v2531 int32
	_ = v2531
	var v2533 int32
	_ = v2533
	var v2537 int32
	_ = v2537
	var v2539 int32
	_ = v2539
	var v2542 int32
	_ = v2542
	var v2547 int32
	_ = v2547
	var v2548 int32
	_ = v2548
	var v2549 int32
	_ = v2549
	var v2551 int32
	_ = v2551
	var v2559 int32
	_ = v2559
	var v2575 int32
	_ = v2575
	var v2576 int32
	_ = v2576
	var v2592 int32
	_ = v2592
	var v2593 int32
	_ = v2593
	var v2595 int32
	_ = v2595
	var v2597 int32
	_ = v2597
	var v2604 int32
	_ = v2604
	var v2606 int32
	_ = v2606
	var v2608 int32
	_ = v2608
	var v2610 int32
	_ = v2610
	var v2623 int32
	_ = v2623
	var v2625 int32
	_ = v2625
	var v2627 int32
	_ = v2627
	var v2645 int32
	_ = v2645
	var v2647 int32
	_ = v2647
	var v2655 int32
	_ = v2655
	var v2659 int32
	_ = v2659
	var v2661 int32
	_ = v2661
	var v2667 int32
	_ = v2667
	var v2683 int32
	_ = v2683
	var v2690 int32
	_ = v2690
	var v2691 int32
	_ = v2691
	var v2693 int32
	_ = v2693
	var v2695 int32
	_ = v2695
	var v2698 int32
	_ = v2698
	var v2700 int32
	_ = v2700
	var v2702 int32
	_ = v2702
	var v2706 int32
	_ = v2706
	var v2710 int32
	_ = v2710
	var v2713 int32
	_ = v2713
	var v2715 int32
	_ = v2715
	var v2718 int32
	_ = v2718
	var v2721 int32
	_ = v2721
	var v2722 int32
	_ = v2722
	var v2727 int32
	_ = v2727
	var v2729 int32
	_ = v2729
	var v2732 int32
	_ = v2732
	var v2734 int32
	_ = v2734
	var v2738 int32
	_ = v2738
	var v2742 int32
	_ = v2742
	var v2758 int32
	_ = v2758
	var v2759 int32
	_ = v2759
	var v2775 int32
	_ = v2775
	var v2776 int32
	_ = v2776
	var v2778 int32
	_ = v2778
	var v2780 int32
	_ = v2780
	var v2787 int32
	_ = v2787
	var v2789 int32
	_ = v2789
	var v2791 int32
	_ = v2791
	var v2793 int32
	_ = v2793
	var v2806 int32
	_ = v2806
	var v2808 int32
	_ = v2808
	var v2810 int32
	_ = v2810
	var v2828 int32
	_ = v2828
	var v2830 int32
	_ = v2830
	var v2838 int32
	_ = v2838
	var v2842 int32
	_ = v2842
	var v2844 int32
	_ = v2844
	var v2850 int32
	_ = v2850
	var v2866 int32
	_ = v2866
	var v2873 int32
	_ = v2873
	var v2874 int32
	_ = v2874
	var v2875 int32
	_ = v2875
	var v2877 int32
	_ = v2877
	var v2881 int32
	_ = v2881
	var v2888 int32
	_ = v2888
	var v2890 int32
	_ = v2890
	var v2892 int32
	_ = v2892
	var v2894 int32
	_ = v2894
	var v2896 int32
	_ = v2896
	var v2897 int32
	_ = v2897
	var v2907 int32
	_ = v2907
	var v2908 int32
	_ = v2908
	var v2909 int32
	_ = v2909
	var v2913 int32
	_ = v2913
	var v2916 int32
	_ = v2916
	var v2917 int32
	_ = v2917
	var v2922 int32
	_ = v2922
	var v2923 int32
	_ = v2923
	var v2928 int32
	_ = v2928
	var v2929 int32
	_ = v2929
	var v2930 int32
	_ = v2930
	var v2937 int32
	_ = v2937
	var v2939 int32
	_ = v2939
	var v2944 int32
	_ = v2944
	var v2946 int32
	_ = v2946
	var v2953 int32
	_ = v2953
	var v2956 int32
	_ = v2956
	var v2960 int32
	_ = v2960
	var v2967 int32
	_ = v2967
	var v2968 int32
	_ = v2968
	var v2982 int32
	_ = v2982
	var v2988 int32
	_ = v2988
	var v2996 int32
	_ = v2996
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v8 = v6
	goto L2
L1:
	;
	return v2996
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v8
	v16 = F_find_among(m, l0, int32(_a_F_italian_UTF_8_stem_0), int32(7), int32(0))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v122 = v6
	goto L51
L4:
	;
	goto L3
L5:
	;
	return int32(0)
L6:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v20
	switch v16 - int32(1) {
	case 0:
		goto L14
	case 1:
		goto L13
	case 2:
		goto L12
	case 3:
		goto L11
	case 4:
		goto L10
	case 5:
		goto L9
	case 6:
		goto L8
	default:
		goto L7
	}
L7:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v8 = v118
	goto L2
L8:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L29
L9:
	;
	v56 = F_slice_from_s(m, l0, int32(2), int32(_a_F_italian_UTF_8_stem_1))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L5
	} else {
		goto L25
	}
L10:
	;
	v50 = F_slice_from_s(m, l0, int32(2), int32(_a_F_italian_UTF_8_stem_2))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L5
	} else {
		goto L23
	}
L11:
	;
	v44 = F_slice_from_s(m, l0, int32(2), int32(_a_F_italian_UTF_8_stem_3))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L5
	} else {
		goto L21
	}
L12:
	;
	v38 = F_slice_from_s(m, l0, int32(2), int32(_a_F_italian_UTF_8_stem_4))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L5
	} else {
		goto L19
	}
L13:
	;
	v32 = F_slice_from_s(m, l0, int32(2), int32(_a_F_italian_UTF_8_stem_5))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L5
	} else {
		goto L17
	}
L14:
	;
	v26 = F_slice_from_s(m, l0, int32(2), int32(_a_F_italian_UTF_8_stem_6))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L5
	} else {
		goto L15
	}
L15:
	;
	if int32(0) <= v26 {
		goto L7
	} else {
		goto L16
	}
L16:
	;
	v2996 = v26
	goto L1
L17:
	;
	if int32(0) <= v32 {
		goto L7
	} else {
		goto L18
	}
L18:
	;
	v2996 = v32
	goto L1
L19:
	;
	if int32(0) <= v38 {
		goto L7
	} else {
		goto L20
	}
L20:
	;
	v2996 = v38
	goto L1
L21:
	;
	if int32(0) <= v44 {
		goto L7
	} else {
		goto L22
	}
L22:
	;
	v2996 = v44
	goto L1
L23:
	;
	if int32(0) <= v50 {
		goto L7
	} else {
		goto L24
	}
L24:
	;
	v2996 = v50
	goto L1
L25:
	;
	if int32(0) <= v56 {
		goto L7
	} else {
		goto L26
	}
L26:
	;
	v2996 = v56
	goto L1
L27:
	;
	if v113 < int32(0) {
		goto L47
	} else {
		goto L48
	}
L29:
	;
	goto L30
L30:
	;
	goto L31
L31:
	;
	v68 = v20
	v70 = int32(1)
	goto L34
L33:
	;
	v113 = v98
	goto L27
L34:
	;
	if v61 <= v68 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	goto L33
L36:
	;
	v113 = int32(-1)
	goto L27
L37:
	;
	goto L38
L38:
	;
	v75 = v68 + int32(1)
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60+v68))))
	if base.Ui32(v77) < base.Ui32(int32(192)) {
		v98 = v75
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v99 = int32(1)
	if v99 < v70 {
		v68 = v98
		v70 = v70 - v99
		goto L34
	} else {
		goto L46
	}
L40:
	;
	if v61 <= v75 {
		v98 = v75
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v84 = v75
	goto L42
L42:
	;
	v87 = int32(*(*int8)(unsafe.Add(mBase, uint32(v60+v84))))
	if int32(-65) < v87 {
		v98 = v84
		goto L39
	} else {
		goto L44
	}
L43:
	;
	v98 = v61
	goto L39
L44:
	;
	v91 = v84 + int32(1)
	if v91 != v61 {
		v84 = v91
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	goto L35
L47:
	;
	goto L4
L48:
	;
	goto L49
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v113
	goto L7
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2729
	v2732 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2729 <= v2732 {
		goto L657
	} else {
		goto L658
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v122
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L56
L52:
	;
	v2727 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2727
	v2729 = v2727
	goto L50
L53:
	;
	goto L52
L54:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v242 != 0 {
		v517 = v243
		goto L79
	} else {
		goto L80
	}
L55:
	;
	v242 = v235
	goto L54
L56:
	;
	if v137 <= v122 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v235 = int32(0)
	goto L55
L58:
	;
	v242 = int32(-1)
	goto L54
L59:
	;
	goto L60
L60:
	;
	v153 = int32(1)
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122+v138))))
	if base.Ui32(v155) < base.Ui32(int32(192)) {
		v212 = v155
		v213 = v153
		goto L61
	} else {
		goto L62
	}
L61:
	;
	if int32(249) < v212 {
		v235 = v213
		goto L55
	} else {
		goto L74
	}
L62:
	;
	v159 = v122 + int32(1)
	if v159 == v137 {
		v212 = v155
		v213 = v153
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159+v138))))
	v164 = v162 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v155) {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168+v138))))
	v180 = v178 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v155) {
		goto L70
	} else {
		goto L71
	}
L65:
	;
	v168 = v122 + int32(2)
	if v168 != v137 {
		goto L64
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v212 = v155<<(uint(int32(6))%32)&int32(1984) | v164
	v213 = int32(2)
	goto L61
L68:
	;
	goto L67
L69:
	;
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138+v184))))
	v212 = v197&int32(63) | (v155<<(uint(int32(18))%32)&int32(_a_F_italian_UTF_8_stem_7) | v164<<(uint(int32(12))%32) | v180<<(uint(int32(6))%32))
	v213 = int32(4)
	goto L61
L70:
	;
	v184 = v122 + int32(3)
	if v184 != v137 {
		goto L69
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v212 = v155<<(uint(int32(12))%32)&int32(_a_F_italian_UTF_8_stem_8) | v164<<(uint(int32(6))%32) | v180
	v213 = int32(3)
	goto L61
L73:
	;
	goto L72
L74:
	;
	v217 = v212 - int32(97)
	if v217 < int32(0) {
		v235 = v213
		goto L55
	} else {
		goto L75
	}
L75:
	;
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v217)>>(uint(int32(3))%32)))+uint32(_c_F_italian_UTF_8_stem[0]))))
	if int32(base.Ui32(v223)>>(uint(v217&int32(7))%32))&int32(1) == int32(0) {
		v235 = v213
		goto L55
	} else {
		goto L76
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v213 + v122
	goto L77
L77:
	;
	goto L57
L78:
	;
	v2721 = F_slice_from_s(m, l0, int32(1), int32(_a_F_italian_UTF_8_stem_9))
	mBase = m.M
	v2722 = m.ExcPending
	if v2722 != 0 {
		goto L5
	} else {
		goto L655
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v122
	v519 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L144
L80:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v244
	if v244 == v243 {
		v383 = v243
		goto L81
	} else {
		goto L82
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v244
	if v244 == v383 {
		goto L113
	} else {
		goto L114
	}
L82:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v247+v244))))
	if v249 != int32(117) {
		v383 = v243
		goto L81
	} else {
		goto L83
	}
L83:
	;
	v253 = v244 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v253
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v253
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L86
L84:
	;
	if v373 == int32(0) {
		goto L108
	} else {
		goto L109
	}
L85:
	;
	v373 = v366
	goto L84
L86:
	;
	if v268 <= v253 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v366 = int32(0)
	goto L85
L88:
	;
	v373 = int32(-1)
	goto L84
L89:
	;
	goto L90
L90:
	;
	v284 = int32(1)
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253+v269))))
	if base.Ui32(v286) < base.Ui32(int32(192)) {
		v343 = v286
		v344 = v284
		goto L91
	} else {
		goto L92
	}
L91:
	;
	if int32(249) < v343 {
		v366 = v344
		goto L85
	} else {
		goto L104
	}
L92:
	;
	v290 = v244 + int32(2)
	if v290 == v268 {
		v343 = v286
		v344 = v284
		goto L91
	} else {
		goto L93
	}
L93:
	;
	v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v290+v269))))
	v295 = v293 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v286) {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v299+v269))))
	v311 = v309 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v286) {
		goto L100
	} else {
		goto L101
	}
L95:
	;
	v299 = v244 + int32(3)
	if v299 != v268 {
		goto L94
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	v343 = v286<<(uint(int32(6))%32)&int32(1984) | v295
	v344 = int32(2)
	goto L91
L98:
	;
	goto L97
L99:
	;
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269+v315))))
	v343 = v328&int32(63) | (v286<<(uint(int32(18))%32)&int32(_a_F_italian_UTF_8_stem_7) | v295<<(uint(int32(12))%32) | v311<<(uint(int32(6))%32))
	v344 = int32(4)
	goto L91
L100:
	;
	v315 = v244 + int32(4)
	if v315 != v268 {
		goto L99
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	v343 = v286<<(uint(int32(12))%32)&int32(_a_F_italian_UTF_8_stem_8) | v295<<(uint(int32(6))%32) | v311
	v344 = int32(3)
	goto L91
L103:
	;
	goto L102
L104:
	;
	v348 = v343 - int32(97)
	if v348 < int32(0) {
		v366 = v344
		goto L85
	} else {
		goto L105
	}
L105:
	;
	v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v348)>>(uint(int32(3))%32)))+uint32(_c_F_italian_UTF_8_stem[0]))))
	if int32(base.Ui32(v354)>>(uint(v348&int32(7))%32))&int32(1) == int32(0) {
		v366 = v344
		goto L85
	} else {
		goto L106
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v344 + v253
	goto L107
L107:
	;
	goto L87
L108:
	;
	v378 = F_slice_from_s(m, l0, int32(1), int32(_a_F_italian_UTF_8_stem_10))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L5
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v383 = v382
	goto L81
L111:
	;
	if v378 < int32(0) {
		v2996 = v378
		goto L1
	} else {
		goto L112
	}
L112:
	;
	goto L51
L113:
	;
	v517 = v244
	goto L79
L114:
	;
	goto L115
L115:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v386+v244))))
	if v388 != int32(105) {
		v517 = v383
		goto L79
	} else {
		goto L116
	}
L116:
	;
	v392 = v244 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v392
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v392
	v407 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v408 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L119
L117:
	;
	if v512 == int32(0) {
		goto L78
	} else {
		goto L141
	}
L118:
	;
	v512 = v505
	goto L117
L119:
	;
	if v407 <= v392 {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	v505 = int32(0)
	goto L118
L121:
	;
	v512 = int32(-1)
	goto L117
L122:
	;
	goto L123
L123:
	;
	v423 = int32(1)
	v425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v392+v408))))
	if base.Ui32(v425) < base.Ui32(int32(192)) {
		v482 = v425
		v483 = v423
		goto L124
	} else {
		goto L125
	}
L124:
	;
	if int32(249) < v482 {
		v505 = v483
		goto L118
	} else {
		goto L137
	}
L125:
	;
	v429 = v244 + int32(2)
	if v429 == v407 {
		v482 = v425
		v483 = v423
		goto L124
	} else {
		goto L126
	}
L126:
	;
	v432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v429+v408))))
	v434 = v432 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v425) {
		goto L128
	} else {
		goto L129
	}
L127:
	;
	v448 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438+v408))))
	v450 = v448 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v425) {
		goto L133
	} else {
		goto L134
	}
L128:
	;
	v438 = v244 + int32(3)
	if v438 != v407 {
		goto L127
	} else {
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	v482 = v425<<(uint(int32(6))%32)&int32(1984) | v434
	v483 = int32(2)
	goto L124
L131:
	;
	goto L130
L132:
	;
	v467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v408+v454))))
	v482 = v467&int32(63) | (v425<<(uint(int32(18))%32)&int32(_a_F_italian_UTF_8_stem_7) | v434<<(uint(int32(12))%32) | v450<<(uint(int32(6))%32))
	v483 = int32(4)
	goto L124
L133:
	;
	v454 = v244 + int32(4)
	if v454 != v407 {
		goto L132
	} else {
		goto L136
	}
L134:
	;
	goto L135
L135:
	;
	v482 = v425<<(uint(int32(12))%32)&int32(_a_F_italian_UTF_8_stem_8) | v434<<(uint(int32(6))%32) | v450
	v483 = int32(3)
	goto L124
L136:
	;
	goto L135
L137:
	;
	v487 = v482 - int32(97)
	if v487 < int32(0) {
		v505 = v483
		goto L118
	} else {
		goto L138
	}
L138:
	;
	v493 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v487)>>(uint(int32(3))%32)))+uint32(_c_F_italian_UTF_8_stem[0]))))
	if int32(base.Ui32(v493)>>(uint(v487&int32(7))%32))&int32(1) == int32(0) {
		v505 = v483
		goto L118
	} else {
		goto L139
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v483 + v392
	goto L140
L140:
	;
	goto L120
L141:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v517 = v515
	goto L79
L142:
	;
	if int32(0) <= v571 {
		v122 = v571
		goto L51
	} else {
		goto L162
	}
L144:
	;
	goto L145
L145:
	;
	goto L146
L146:
	;
	v526 = v122
	v528 = int32(1)
	goto L149
L148:
	;
	v571 = v556
	goto L142
L149:
	;
	if v517 <= v526 {
		goto L151
	} else {
		goto L152
	}
L150:
	;
	goto L148
L151:
	;
	v571 = int32(-1)
	goto L142
L152:
	;
	goto L153
L153:
	;
	v533 = v526 + int32(1)
	v535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v519+v526))))
	if base.Ui32(v535) < base.Ui32(int32(192)) {
		v556 = v533
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v557 = int32(1)
	if v557 < v528 {
		v526 = v556
		v528 = v528 - v557
		goto L149
	} else {
		goto L161
	}
L155:
	;
	if v517 <= v533 {
		v556 = v533
		goto L154
	} else {
		goto L156
	}
L156:
	;
	v542 = v533
	goto L157
L157:
	;
	v545 = int32(*(*int8)(unsafe.Add(mBase, uint32(v519+v542))))
	if int32(-65) < v545 {
		v556 = v542
		goto L154
	} else {
		goto L159
	}
L158:
	;
	v556 = v517
	goto L154
L159:
	;
	v549 = v542 + int32(1)
	if v549 != v517 {
		v542 = v549
		goto L157
	} else {
		goto L160
	}
L160:
	;
	goto L158
L161:
	;
	goto L150
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v6
	v575 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v575
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v575
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v575
	v592 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L168
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v6
	v1749 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1750 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1758 = v6
	goto L429
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v1731
	goto L163
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v6
	v1181 = int32(5)
	v1183 = int32(0)
	v1185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v1185-v6 < v1181 {
		v1195 = v1183
		goto L295
	} else {
		goto L296
	}
L166:
	;
	if v696 != 0 {
		goto L165
	} else {
		goto L190
	}
L167:
	;
	v696 = v689
	goto L166
L168:
	;
	if v575 <= v6 {
		goto L170
	} else {
		goto L171
	}
L169:
	;
	v689 = int32(0)
	goto L167
L170:
	;
	v696 = int32(-1)
	goto L166
L171:
	;
	goto L172
L172:
	;
	v607 = int32(1)
	v609 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6+v592))))
	if base.Ui32(v609) < base.Ui32(int32(192)) {
		v666 = v609
		v667 = v607
		goto L173
	} else {
		goto L174
	}
L173:
	;
	if int32(249) < v666 {
		v689 = v667
		goto L167
	} else {
		goto L186
	}
L174:
	;
	v613 = v6 + int32(1)
	if v613 == v575 {
		v666 = v609
		v667 = v607
		goto L173
	} else {
		goto L175
	}
L175:
	;
	v616 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v613+v592))))
	v618 = v616 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v609) {
		goto L177
	} else {
		goto L178
	}
L176:
	;
	v632 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v622+v592))))
	v634 = v632 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v609) {
		goto L182
	} else {
		goto L183
	}
L177:
	;
	v622 = v6 + int32(2)
	if v622 != v575 {
		goto L176
	} else {
		goto L180
	}
L178:
	;
	goto L179
L179:
	;
	v666 = v609<<(uint(int32(6))%32)&int32(1984) | v618
	v667 = int32(2)
	goto L173
L180:
	;
	goto L179
L181:
	;
	v651 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v592+v638))))
	v666 = v651&int32(63) | (v609<<(uint(int32(18))%32)&int32(_a_F_italian_UTF_8_stem_7) | v618<<(uint(int32(12))%32) | v634<<(uint(int32(6))%32))
	v667 = int32(4)
	goto L173
L182:
	;
	v638 = v6 + int32(3)
	if v638 != v575 {
		goto L181
	} else {
		goto L185
	}
L183:
	;
	goto L184
L184:
	;
	v666 = v609<<(uint(int32(12))%32)&int32(_a_F_italian_UTF_8_stem_8) | v618<<(uint(int32(6))%32) | v634
	v667 = int32(3)
	goto L173
L185:
	;
	goto L184
L186:
	;
	v671 = v666 - int32(97)
	if v671 < int32(0) {
		v689 = v667
		goto L167
	} else {
		goto L187
	}
L187:
	;
	v677 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v671)>>(uint(int32(3))%32)))+uint32(_c_F_italian_UTF_8_stem[0]))))
	if int32(base.Ui32(v677)>>(uint(v671&int32(7))%32))&int32(1) == int32(0) {
		v689 = v667
		goto L167
	} else {
		goto L188
	}
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v667 + v6
	goto L189
L189:
	;
	goto L169
L190:
	;
	v697 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v710 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v711 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L194
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v697
	v950 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v951 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L246
L192:
	;
	if v814 != 0 {
		goto L191
	} else {
		goto L217
	}
L193:
	;
	v814 = v807
	goto L192
L194:
	;
	if v710 <= v697 {
		goto L196
	} else {
		goto L197
	}
L195:
	;
	v807 = int32(0)
	goto L193
L196:
	;
	v814 = int32(-1)
	goto L192
L197:
	;
	goto L198
L198:
	;
	v726 = int32(1)
	v728 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v697+v711))))
	if base.Ui32(v728) < base.Ui32(int32(192)) {
		v785 = v728
		v786 = v726
		goto L199
	} else {
		goto L200
	}
L199:
	;
	if int32(249) < v785 {
		goto L212
	} else {
		goto L213
	}
L200:
	;
	v732 = v697 + int32(1)
	if v732 == v710 {
		v785 = v728
		v786 = v726
		goto L199
	} else {
		goto L201
	}
L201:
	;
	v735 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v732+v711))))
	v737 = v735 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v728) {
		goto L203
	} else {
		goto L204
	}
L202:
	;
	v751 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v741+v711))))
	v753 = v751 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v728) {
		goto L208
	} else {
		goto L209
	}
L203:
	;
	v741 = v697 + int32(2)
	if v741 != v710 {
		goto L202
	} else {
		goto L206
	}
L204:
	;
	goto L205
L205:
	;
	v785 = v728<<(uint(int32(6))%32)&int32(1984) | v737
	v786 = int32(2)
	goto L199
L206:
	;
	goto L205
L207:
	;
	v770 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v711+v757))))
	v785 = v770&int32(63) | (v728<<(uint(int32(18))%32)&int32(_a_F_italian_UTF_8_stem_7) | v737<<(uint(int32(12))%32) | v753<<(uint(int32(6))%32))
	v786 = int32(4)
	goto L199
L208:
	;
	v757 = v697 + int32(3)
	if v757 != v710 {
		goto L207
	} else {
		goto L211
	}
L209:
	;
	goto L210
L210:
	;
	v785 = v728<<(uint(int32(12))%32)&int32(_a_F_italian_UTF_8_stem_8) | v737<<(uint(int32(6))%32) | v753
	v786 = int32(3)
	goto L199
L211:
	;
	goto L210
L212:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v786 + v697
	goto L216
L213:
	;
	v790 = v785 - int32(97)
	if v790 < int32(0) {
		goto L212
	} else {
		goto L214
	}
L214:
	;
	v796 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v790)>>(uint(int32(3))%32)))+uint32(_c_F_italian_UTF_8_stem[0]))))
	if int32(base.Ui32(v796)>>(uint(v790&int32(7))%32))&int32(1) != 0 {
		v807 = v786
		goto L193
	} else {
		goto L215
	}
L215:
	;
	goto L212
L216:
	;
	goto L195
L217:
	;
	v826 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v827 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v828 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v836 = v826
	goto L220
L218:
	;
	if v931 < int32(0) {
		goto L191
	} else {
		goto L243
	}
L219:
	;
	v931 = v903
	goto L218
L220:
	;
	if v827 <= v836 {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	v931 = int32(-1)
	goto L218
L223:
	;
	goto L224
L224:
	;
	v843 = int32(1)
	v845 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v836+v828))))
	if base.Ui32(v845) < base.Ui32(int32(192)) {
		v902 = v845
		v903 = v843
		goto L225
	} else {
		goto L226
	}
L225:
	;
	if int32(249) < v902 {
		goto L238
	} else {
		goto L239
	}
L226:
	;
	v849 = v836 + int32(1)
	if v849 == v827 {
		v902 = v845
		v903 = v843
		goto L225
	} else {
		goto L227
	}
L227:
	;
	v852 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v849+v828))))
	v854 = v852 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v845) {
		goto L229
	} else {
		goto L230
	}
L228:
	;
	v868 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v858+v828))))
	v870 = v868 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v845) {
		goto L234
	} else {
		goto L235
	}
L229:
	;
	v858 = v836 + int32(2)
	if v858 != v827 {
		goto L228
	} else {
		goto L232
	}
L230:
	;
	goto L231
L231:
	;
	v902 = v845<<(uint(int32(6))%32)&int32(1984) | v854
	v903 = int32(2)
	goto L225
L232:
	;
	goto L231
L233:
	;
	v887 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v828+v874))))
	v902 = v887&int32(63) | (v845<<(uint(int32(18))%32)&int32(_a_F_italian_UTF_8_stem_7) | v854<<(uint(int32(12))%32) | v870<<(uint(int32(6))%32))
	v903 = int32(4)
	goto L225
L234:
	;
	v874 = v836 + int32(3)
	if v874 != v827 {
		goto L233
	} else {
		goto L237
	}
L235:
	;
	goto L236
L236:
	;
	v902 = v845<<(uint(int32(12))%32)&int32(_a_F_italian_UTF_8_stem_8) | v854<<(uint(int32(6))%32) | v870
	v903 = int32(3)
	goto L225
L237:
	;
	goto L236
L238:
	;
	v920 = v903 + v836
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v920
	v836 = v920
	goto L220
L239:
	;
	v907 = v902 - int32(97)
	if v907 < int32(0) {
		goto L238
	} else {
		goto L240
	}
L240:
	;
	v913 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v907)>>(uint(int32(3))%32)))+uint32(_c_F_italian_UTF_8_stem[0]))))
	if int32(base.Ui32(v913)>>(uint(v907&int32(7))%32))&int32(1) != 0 {
		goto L219
	} else {
		goto L241
	}
L241:
	;
	goto L238
L243:
	;
	v934 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1731 = v934 + v931
	goto L164
L244:
	;
	if v1055 != 0 {
		goto L165
	} else {
		goto L268
	}
L245:
	;
	v1055 = v1048
	goto L244
L246:
	;
	if v950 <= v697 {
		goto L248
	} else {
		goto L249
	}
L247:
	;
	v1048 = int32(0)
	goto L245
L248:
	;
	v1055 = int32(-1)
	goto L244
L249:
	;
	goto L250
L250:
	;
	v966 = int32(1)
	v968 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v697+v951))))
	if base.Ui32(v968) < base.Ui32(int32(192)) {
		v1025 = v968
		v1026 = v966
		goto L251
	} else {
		goto L252
	}
L251:
	;
	if int32(249) < v1025 {
		v1048 = v1026
		goto L245
	} else {
		goto L264
	}
L252:
	;
	v972 = v697 + int32(1)
	if v972 == v950 {
		v1025 = v968
		v1026 = v966
		goto L251
	} else {
		goto L253
	}
L253:
	;
	v975 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v972+v951))))
	v977 = v975 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v968) {
		goto L255
	} else {
		goto L256
	}
L254:
	;
	v991 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v981+v951))))
	v993 = v991 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v968) {
		goto L260
	} else {
		goto L261
	}
L255:
	;
	v981 = v697 + int32(2)
	if v981 != v950 {
		goto L254
	} else {
		goto L258
	}
L256:
	;
	goto L257
L257:
	;
	v1025 = v968<<(uint(int32(6))%32)&int32(1984) | v977
	v1026 = int32(2)
	goto L251
L258:
	;
	goto L257
L259:
	;
	v1010 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v951+v997))))
	v1025 = v1010&int32(63) | (v968<<(uint(int32(18))%32)&int32(_a_F_italian_UTF_8_stem_7) | v977<<(uint(int32(12))%32) | v993<<(uint(int32(6))%32))
	v1026 = int32(4)
	goto L251
L260:
	;
	v997 = v697 + int32(3)
	if v997 != v950 {
		goto L259
	} else {
		goto L263
	}
L261:
	;
	goto L262
L262:
	;
	v1025 = v968<<(uint(int32(12))%32)&int32(_a_F_italian_UTF_8_stem_8) | v977<<(uint(int32(6))%32) | v993
	v1026 = int32(3)
	goto L251
L263:
	;
	goto L262
L264:
	;
	v1030 = v1025 - int32(97)
	if v1030 < int32(0) {
		v1048 = v1026
		goto L245
	} else {
		goto L265
	}
L265:
	;
	v1036 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1030)>>(uint(int32(3))%32)))+uint32(_c_F_italian_UTF_8_stem[0]))))
	if int32(base.Ui32(v1036)>>(uint(v1030&int32(7))%32))&int32(1) == int32(0) {
		v1048 = v1026
		goto L245
	} else {
		goto L266
	}
L266:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1026 + v697
	goto L267
L267:
	;
	goto L247
L268:
	;
	v1067 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1068 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1069 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1077 = v1067
	goto L271
L269:
	;
	if v1173 < int32(0) {
		goto L165
	} else {
		goto L293
	}
L270:
	;
	v1173 = v1144
	goto L269
L271:
	;
	if v1068 <= v1077 {
		goto L273
	} else {
		goto L274
	}
L273:
	;
	v1173 = int32(-1)
	goto L269
L274:
	;
	goto L275
L275:
	;
	v1084 = int32(1)
	v1086 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1077+v1069))))
	if base.Ui32(v1086) < base.Ui32(int32(192)) {
		v1143 = v1086
		v1144 = v1084
		goto L276
	} else {
		goto L277
	}
L276:
	;
	if int32(249) < v1143 {
		goto L270
	} else {
		goto L289
	}
L277:
	;
	v1090 = v1077 + int32(1)
	if v1090 == v1068 {
		v1143 = v1086
		v1144 = v1084
		goto L276
	} else {
		goto L278
	}
L278:
	;
	v1093 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1090+v1069))))
	v1095 = v1093 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1086) {
		goto L280
	} else {
		goto L281
	}
L279:
	;
	v1109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1099+v1069))))
	v1111 = v1109 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1086) {
		goto L285
	} else {
		goto L286
	}
L280:
	;
	v1099 = v1077 + int32(2)
	if v1099 != v1068 {
		goto L279
	} else {
		goto L283
	}
L281:
	;
	goto L282
L282:
	;
	v1143 = v1086<<(uint(int32(6))%32)&int32(1984) | v1095
	v1144 = int32(2)
	goto L276
L283:
	;
	goto L282
L284:
	;
	v1128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1069+v1115))))
	v1143 = v1128&int32(63) | (v1086<<(uint(int32(18))%32)&int32(_a_F_italian_UTF_8_stem_7) | v1095<<(uint(int32(12))%32) | v1111<<(uint(int32(6))%32))
	v1144 = int32(4)
	goto L276
L285:
	;
	v1115 = v1077 + int32(3)
	if v1115 != v1068 {
		goto L284
	} else {
		goto L288
	}
L286:
	;
	goto L287
L287:
	;
	v1143 = v1086<<(uint(int32(12))%32)&int32(_a_F_italian_UTF_8_stem_8) | v1095<<(uint(int32(6))%32) | v1111
	v1144 = int32(3)
	goto L276
L288:
	;
	goto L287
L289:
	;
	v1148 = v1143 - int32(97)
	if v1148 < int32(0) {
		goto L270
	} else {
		goto L290
	}
L290:
	;
	v1154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1148)>>(uint(int32(3))%32)))+uint32(_c_F_italian_UTF_8_stem[0]))))
	if int32(base.Ui32(v1154)>>(uint(v1148&int32(7))%32))&int32(1) == int32(0) {
		goto L270
	} else {
		goto L291
	}
L291:
	;
	v1162 = v1144 + v1077
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1162
	v1077 = v1162
	goto L271
L293:
	;
	v1176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1731 = v1176 + v1173
	goto L164
L294:
	;
	if v1195 != 0 {
		goto L298
	} else {
		goto L299
	}
L295:
	;
	goto L294
L296:
	;
	v1189 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1191 = F_memcmp(m, v1189+v6, int32(_a_F_italian_UTF_8_stem_11), v1181)
	mBase = m.M
	if v1191 != 0 {
		v1195 = v1183
		goto L295
	} else {
		goto L297
	}
L297:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1181 + v6
	v1195 = int32(1)
	goto L295
L298:
	;
	v1196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1731 = v1196
	goto L164
L299:
	;
	goto L300
L300:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v6
	v1210 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1211 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L303
L301:
	;
	if v1314 != 0 {
		goto L163
	} else {
		goto L326
	}
L302:
	;
	v1314 = v1307
	goto L301
L303:
	;
	if v1210 <= v6 {
		goto L305
	} else {
		goto L306
	}
L304:
	;
	v1307 = int32(0)
	goto L302
L305:
	;
	v1314 = int32(-1)
	goto L301
L306:
	;
	goto L307
L307:
	;
	v1226 = int32(1)
	v1228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6+v1211))))
	if base.Ui32(v1228) < base.Ui32(int32(192)) {
		v1285 = v1228
		v1286 = v1226
		goto L308
	} else {
		goto L309
	}
L308:
	;
	if int32(249) < v1285 {
		goto L321
	} else {
		goto L322
	}
L309:
	;
	v1232 = v6 + int32(1)
	if v1232 == v1210 {
		v1285 = v1228
		v1286 = v1226
		goto L308
	} else {
		goto L310
	}
L310:
	;
	v1235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1232+v1211))))
	v1237 = v1235 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1228) {
		goto L312
	} else {
		goto L313
	}
L311:
	;
	v1251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1241+v1211))))
	v1253 = v1251 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1228) {
		goto L317
	} else {
		goto L318
	}
L312:
	;
	v1241 = v6 + int32(2)
	if v1241 != v1210 {
		goto L311
	} else {
		goto L315
	}
L313:
	;
	goto L314
L314:
	;
	v1285 = v1228<<(uint(int32(6))%32)&int32(1984) | v1237
	v1286 = int32(2)
	goto L308
L315:
	;
	goto L314
L316:
	;
	v1270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1211+v1257))))
	v1285 = v1270&int32(63) | (v1228<<(uint(int32(18))%32)&int32(_a_F_italian_UTF_8_stem_7) | v1237<<(uint(int32(12))%32) | v1253<<(uint(int32(6))%32))
	v1286 = int32(4)
	goto L308
L317:
	;
	v1257 = v6 + int32(3)
	if v1257 != v1210 {
		goto L316
	} else {
		goto L320
	}
L318:
	;
	goto L319
L319:
	;
	v1285 = v1228<<(uint(int32(12))%32)&int32(_a_F_italian_UTF_8_stem_8) | v1237<<(uint(int32(6))%32) | v1253
	v1286 = int32(3)
	goto L308
L320:
	;
	goto L319
L321:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1286 + v6
	goto L325
L322:
	;
	v1290 = v1285 - int32(97)
	if v1290 < int32(0) {
		goto L321
	} else {
		goto L323
	}
L323:
	;
	v1296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1290)>>(uint(int32(3))%32)))+uint32(_c_F_italian_UTF_8_stem[0]))))
	if int32(base.Ui32(v1296)>>(uint(v1290&int32(7))%32))&int32(1) != 0 {
		v1307 = v1286
		goto L302
	} else {
		goto L324
	}
L324:
	;
	goto L321
L325:
	;
	goto L304
L326:
	;
	v1315 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1328 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1329 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L330
L327:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1315
	v1568 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1569 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L382
L328:
	;
	if v1432 != 0 {
		goto L327
	} else {
		goto L353
	}
L329:
	;
	v1432 = v1425
	goto L328
L330:
	;
	if v1328 <= v1315 {
		goto L332
	} else {
		goto L333
	}
L331:
	;
	v1425 = int32(0)
	goto L329
L332:
	;
	v1432 = int32(-1)
	goto L328
L333:
	;
	goto L334
L334:
	;
	v1344 = int32(1)
	v1346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1315+v1329))))
	if base.Ui32(v1346) < base.Ui32(int32(192)) {
		v1403 = v1346
		v1404 = v1344
		goto L335
	} else {
		goto L336
	}
L335:
	;
	if int32(249) < v1403 {
		goto L348
	} else {
		goto L349
	}
L336:
	;
	v1350 = v1315 + int32(1)
	if v1350 == v1328 {
		v1403 = v1346
		v1404 = v1344
		goto L335
	} else {
		goto L337
	}
L337:
	;
	v1353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1350+v1329))))
	v1355 = v1353 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1346) {
		goto L339
	} else {
		goto L340
	}
L338:
	;
	v1369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1359+v1329))))
	v1371 = v1369 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1346) {
		goto L344
	} else {
		goto L345
	}
L339:
	;
	v1359 = v1315 + int32(2)
	if v1359 != v1328 {
		goto L338
	} else {
		goto L342
	}
L340:
	;
	goto L341
L341:
	;
	v1403 = v1346<<(uint(int32(6))%32)&int32(1984) | v1355
	v1404 = int32(2)
	goto L335
L342:
	;
	goto L341
L343:
	;
	v1388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1329+v1375))))
	v1403 = v1388&int32(63) | (v1346<<(uint(int32(18))%32)&int32(_a_F_italian_UTF_8_stem_7) | v1355<<(uint(int32(12))%32) | v1371<<(uint(int32(6))%32))
	v1404 = int32(4)
	goto L335
L344:
	;
	v1375 = v1315 + int32(3)
	if v1375 != v1328 {
		goto L343
	} else {
		goto L347
	}
L345:
	;
	goto L346
L346:
	;
	v1403 = v1346<<(uint(int32(12))%32)&int32(_a_F_italian_UTF_8_stem_8) | v1355<<(uint(int32(6))%32) | v1371
	v1404 = int32(3)
	goto L335
L347:
	;
	goto L346
L348:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1404 + v1315
	goto L352
L349:
	;
	v1408 = v1403 - int32(97)
	if v1408 < int32(0) {
		goto L348
	} else {
		goto L350
	}
L350:
	;
	v1414 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1408)>>(uint(int32(3))%32)))+uint32(_c_F_italian_UTF_8_stem[0]))))
	if int32(base.Ui32(v1414)>>(uint(v1408&int32(7))%32))&int32(1) != 0 {
		v1425 = v1404
		goto L329
	} else {
		goto L351
	}
L351:
	;
	goto L348
L352:
	;
	goto L331
L353:
	;
	v1444 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1445 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1446 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1454 = v1444
	goto L356
L354:
	;
	if v1549 < int32(0) {
		goto L327
	} else {
		goto L379
	}
L355:
	;
	v1549 = v1521
	goto L354
L356:
	;
	if v1445 <= v1454 {
		goto L358
	} else {
		goto L359
	}
L358:
	;
	v1549 = int32(-1)
	goto L354
L359:
	;
	goto L360
L360:
	;
	v1461 = int32(1)
	v1463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1454+v1446))))
	if base.Ui32(v1463) < base.Ui32(int32(192)) {
		v1520 = v1463
		v1521 = v1461
		goto L361
	} else {
		goto L362
	}
L361:
	;
	if int32(249) < v1520 {
		goto L374
	} else {
		goto L375
	}
L362:
	;
	v1467 = v1454 + int32(1)
	if v1467 == v1445 {
		v1520 = v1463
		v1521 = v1461
		goto L361
	} else {
		goto L363
	}
L363:
	;
	v1470 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1467+v1446))))
	v1472 = v1470 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1463) {
		goto L365
	} else {
		goto L366
	}
L364:
	;
	v1486 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1476+v1446))))
	v1488 = v1486 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1463) {
		goto L370
	} else {
		goto L371
	}
L365:
	;
	v1476 = v1454 + int32(2)
	if v1476 != v1445 {
		goto L364
	} else {
		goto L368
	}
L366:
	;
	goto L367
L367:
	;
	v1520 = v1463<<(uint(int32(6))%32)&int32(1984) | v1472
	v1521 = int32(2)
	goto L361
L368:
	;
	goto L367
L369:
	;
	v1505 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1446+v1492))))
	v1520 = v1505&int32(63) | (v1463<<(uint(int32(18))%32)&int32(_a_F_italian_UTF_8_stem_7) | v1472<<(uint(int32(12))%32) | v1488<<(uint(int32(6))%32))
	v1521 = int32(4)
	goto L361
L370:
	;
	v1492 = v1454 + int32(3)
	if v1492 != v1445 {
		goto L369
	} else {
		goto L373
	}
L371:
	;
	goto L372
L372:
	;
	v1520 = v1463<<(uint(int32(12))%32)&int32(_a_F_italian_UTF_8_stem_8) | v1472<<(uint(int32(6))%32) | v1488
	v1521 = int32(3)
	goto L361
L373:
	;
	goto L372
L374:
	;
	v1538 = v1521 + v1454
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1538
	v1454 = v1538
	goto L356
L375:
	;
	v1525 = v1520 - int32(97)
	if v1525 < int32(0) {
		goto L374
	} else {
		goto L376
	}
L376:
	;
	v1531 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1525)>>(uint(int32(3))%32)))+uint32(_c_F_italian_UTF_8_stem[0]))))
	if int32(base.Ui32(v1531)>>(uint(v1525&int32(7))%32))&int32(1) != 0 {
		goto L355
	} else {
		goto L377
	}
L377:
	;
	goto L374
L379:
	;
	v1552 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1731 = v1552 + v1549
	goto L164
L380:
	;
	if v1673 != 0 {
		goto L163
	} else {
		goto L404
	}
L381:
	;
	v1673 = v1666
	goto L380
L382:
	;
	if v1568 <= v1315 {
		goto L384
	} else {
		goto L385
	}
L383:
	;
	v1666 = int32(0)
	goto L381
L384:
	;
	v1673 = int32(-1)
	goto L380
L385:
	;
	goto L386
L386:
	;
	v1584 = int32(1)
	v1586 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1315+v1569))))
	if base.Ui32(v1586) < base.Ui32(int32(192)) {
		v1643 = v1586
		v1644 = v1584
		goto L387
	} else {
		goto L388
	}
L387:
	;
	if int32(249) < v1643 {
		v1666 = v1644
		goto L381
	} else {
		goto L400
	}
L388:
	;
	v1590 = v1315 + int32(1)
	if v1590 == v1568 {
		v1643 = v1586
		v1644 = v1584
		goto L387
	} else {
		goto L389
	}
L389:
	;
	v1593 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1590+v1569))))
	v1595 = v1593 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1586) {
		goto L391
	} else {
		goto L392
	}
L390:
	;
	v1609 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1599+v1569))))
	v1611 = v1609 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1586) {
		goto L396
	} else {
		goto L397
	}
L391:
	;
	v1599 = v1315 + int32(2)
	if v1599 != v1568 {
		goto L390
	} else {
		goto L394
	}
L392:
	;
	goto L393
L393:
	;
	v1643 = v1586<<(uint(int32(6))%32)&int32(1984) | v1595
	v1644 = int32(2)
	goto L387
L394:
	;
	goto L393
L395:
	;
	v1628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1569+v1615))))
	v1643 = v1628&int32(63) | (v1586<<(uint(int32(18))%32)&int32(_a_F_italian_UTF_8_stem_7) | v1595<<(uint(int32(12))%32) | v1611<<(uint(int32(6))%32))
	v1644 = int32(4)
	goto L387
L396:
	;
	v1615 = v1315 + int32(3)
	if v1615 != v1568 {
		goto L395
	} else {
		goto L399
	}
L397:
	;
	goto L398
L398:
	;
	v1643 = v1586<<(uint(int32(12))%32)&int32(_a_F_italian_UTF_8_stem_8) | v1595<<(uint(int32(6))%32) | v1611
	v1644 = int32(3)
	goto L387
L399:
	;
	goto L398
L400:
	;
	v1648 = v1643 - int32(97)
	if v1648 < int32(0) {
		v1666 = v1644
		goto L381
	} else {
		goto L401
	}
L401:
	;
	v1654 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1648)>>(uint(int32(3))%32)))+uint32(_c_F_italian_UTF_8_stem[0]))))
	if int32(base.Ui32(v1654)>>(uint(v1648&int32(7))%32))&int32(1) == int32(0) {
		v1666 = v1644
		goto L381
	} else {
		goto L402
	}
L402:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1644 + v1315
	goto L403
L403:
	;
	goto L383
L404:
	;
	v1674 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1675 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1676 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L407
L405:
	;
	if v1728 < int32(0) {
		goto L163
	} else {
		goto L425
	}
L407:
	;
	goto L408
L408:
	;
	goto L409
L409:
	;
	v1683 = v1675
	v1685 = int32(1)
	goto L412
L411:
	;
	v1728 = v1713
	goto L405
L412:
	;
	if v1676 <= v1683 {
		goto L414
	} else {
		goto L415
	}
L413:
	;
	goto L411
L414:
	;
	v1728 = int32(-1)
	goto L405
L415:
	;
	goto L416
L416:
	;
	v1690 = v1683 + int32(1)
	v1692 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1674+v1683))))
	if base.Ui32(v1692) < base.Ui32(int32(192)) {
		v1713 = v1690
		goto L417
	} else {
		goto L418
	}
L417:
	;
	v1714 = int32(1)
	if v1714 < v1685 {
		v1683 = v1713
		v1685 = v1685 - v1714
		goto L412
	} else {
		goto L424
	}
L418:
	;
	if v1676 <= v1690 {
		v1713 = v1690
		goto L417
	} else {
		goto L419
	}
L419:
	;
	v1699 = v1690
	goto L420
L420:
	;
	v1702 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1674+v1699))))
	if int32(-65) < v1702 {
		v1713 = v1699
		goto L417
	} else {
		goto L422
	}
L421:
	;
	v1713 = v1676
	goto L417
L422:
	;
	v1706 = v1699 + int32(1)
	if v1706 != v1676 {
		v1699 = v1706
		goto L420
	} else {
		goto L423
	}
L423:
	;
	goto L421
L424:
	;
	goto L413
L425:
	;
	v1731 = v1728
	goto L164
L426:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v6
	v2230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2230
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2230
	v2234 = v2230 - int32(1)
	if v2234 <= v6 {
		goto L529
	} else {
		goto L530
	}
L427:
	;
	if v1853 < int32(0) {
		goto L426
	} else {
		goto L452
	}
L428:
	;
	v1853 = v1825
	goto L427
L429:
	;
	if v1749 <= v1758 {
		goto L431
	} else {
		goto L432
	}
L431:
	;
	v1853 = int32(-1)
	goto L427
L432:
	;
	goto L433
L433:
	;
	v1765 = int32(1)
	v1767 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1758+v1750))))
	if base.Ui32(v1767) < base.Ui32(int32(192)) {
		v1824 = v1767
		v1825 = v1765
		goto L434
	} else {
		goto L435
	}
L434:
	;
	if int32(249) < v1824 {
		goto L447
	} else {
		goto L448
	}
L435:
	;
	v1771 = v1758 + int32(1)
	if v1771 == v1749 {
		v1824 = v1767
		v1825 = v1765
		goto L434
	} else {
		goto L436
	}
L436:
	;
	v1774 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1771+v1750))))
	v1776 = v1774 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1767) {
		goto L438
	} else {
		goto L439
	}
L437:
	;
	v1790 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1780+v1750))))
	v1792 = v1790 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1767) {
		goto L443
	} else {
		goto L444
	}
L438:
	;
	v1780 = v1758 + int32(2)
	if v1780 != v1749 {
		goto L437
	} else {
		goto L441
	}
L439:
	;
	goto L440
L440:
	;
	v1824 = v1767<<(uint(int32(6))%32)&int32(1984) | v1776
	v1825 = int32(2)
	goto L434
L441:
	;
	goto L440
L442:
	;
	v1809 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1750+v1796))))
	v1824 = v1809&int32(63) | (v1767<<(uint(int32(18))%32)&int32(_a_F_italian_UTF_8_stem_7) | v1776<<(uint(int32(12))%32) | v1792<<(uint(int32(6))%32))
	v1825 = int32(4)
	goto L434
L443:
	;
	v1796 = v1758 + int32(3)
	if v1796 != v1749 {
		goto L442
	} else {
		goto L446
	}
L444:
	;
	goto L445
L445:
	;
	v1824 = v1767<<(uint(int32(12))%32)&int32(_a_F_italian_UTF_8_stem_8) | v1776<<(uint(int32(6))%32) | v1792
	v1825 = int32(3)
	goto L434
L446:
	;
	goto L445
L447:
	;
	v1842 = v1825 + v1758
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1842
	v1758 = v1842
	goto L429
L448:
	;
	v1829 = v1824 - int32(97)
	if v1829 < int32(0) {
		goto L447
	} else {
		goto L449
	}
L449:
	;
	v1835 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1829)>>(uint(int32(3))%32)))+uint32(_c_F_italian_UTF_8_stem[0]))))
	if int32(base.Ui32(v1835)>>(uint(v1829&int32(7))%32))&int32(1) != 0 {
		goto L428
	} else {
		goto L450
	}
L450:
	;
	goto L447
L452:
	;
	v1856 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1857 = v1856 + v1853
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1857
	v1871 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1872 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1880 = v1857
	goto L455
L453:
	;
	if v1976 < int32(0) {
		goto L426
	} else {
		goto L477
	}
L454:
	;
	v1976 = v1947
	goto L453
L455:
	;
	if v1871 <= v1880 {
		goto L457
	} else {
		goto L458
	}
L457:
	;
	v1976 = int32(-1)
	goto L453
L458:
	;
	goto L459
L459:
	;
	v1887 = int32(1)
	v1889 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1880+v1872))))
	if base.Ui32(v1889) < base.Ui32(int32(192)) {
		v1946 = v1889
		v1947 = v1887
		goto L460
	} else {
		goto L461
	}
L460:
	;
	if int32(249) < v1946 {
		goto L454
	} else {
		goto L473
	}
L461:
	;
	v1893 = v1880 + int32(1)
	if v1893 == v1871 {
		v1946 = v1889
		v1947 = v1887
		goto L460
	} else {
		goto L462
	}
L462:
	;
	v1896 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1893+v1872))))
	v1898 = v1896 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1889) {
		goto L464
	} else {
		goto L465
	}
L463:
	;
	v1912 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1902+v1872))))
	v1914 = v1912 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1889) {
		goto L469
	} else {
		goto L470
	}
L464:
	;
	v1902 = v1880 + int32(2)
	if v1902 != v1871 {
		goto L463
	} else {
		goto L467
	}
L465:
	;
	goto L466
L466:
	;
	v1946 = v1889<<(uint(int32(6))%32)&int32(1984) | v1898
	v1947 = int32(2)
	goto L460
L467:
	;
	goto L466
L468:
	;
	v1931 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1872+v1918))))
	v1946 = v1931&int32(63) | (v1889<<(uint(int32(18))%32)&int32(_a_F_italian_UTF_8_stem_7) | v1898<<(uint(int32(12))%32) | v1914<<(uint(int32(6))%32))
	v1947 = int32(4)
	goto L460
L469:
	;
	v1918 = v1880 + int32(3)
	if v1918 != v1871 {
		goto L468
	} else {
		goto L472
	}
L470:
	;
	goto L471
L471:
	;
	v1946 = v1889<<(uint(int32(12))%32)&int32(_a_F_italian_UTF_8_stem_8) | v1898<<(uint(int32(6))%32) | v1914
	v1947 = int32(3)
	goto L460
L472:
	;
	goto L471
L473:
	;
	v1951 = v1946 - int32(97)
	if v1951 < int32(0) {
		goto L454
	} else {
		goto L474
	}
L474:
	;
	v1957 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1951)>>(uint(int32(3))%32)))+uint32(_c_F_italian_UTF_8_stem[0]))))
	if int32(base.Ui32(v1957)>>(uint(v1951&int32(7))%32))&int32(1) == int32(0) {
		goto L454
	} else {
		goto L475
	}
L475:
	;
	v1965 = v1947 + v1880
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1965
	v1880 = v1965
	goto L455
L477:
	;
	v1979 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1980 = v1979 + v1976
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v1980
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1980
	v1995 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1996 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2004 = v1980
	goto L480
L478:
	;
	if v2099 < int32(0) {
		goto L426
	} else {
		goto L503
	}
L479:
	;
	v2099 = v2071
	goto L478
L480:
	;
	if v1995 <= v2004 {
		goto L482
	} else {
		goto L483
	}
L482:
	;
	v2099 = int32(-1)
	goto L478
L483:
	;
	goto L484
L484:
	;
	v2011 = int32(1)
	v2013 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2004+v1996))))
	if base.Ui32(v2013) < base.Ui32(int32(192)) {
		v2070 = v2013
		v2071 = v2011
		goto L485
	} else {
		goto L486
	}
L485:
	;
	if int32(249) < v2070 {
		goto L498
	} else {
		goto L499
	}
L486:
	;
	v2017 = v2004 + int32(1)
	if v2017 == v1995 {
		v2070 = v2013
		v2071 = v2011
		goto L485
	} else {
		goto L487
	}
L487:
	;
	v2020 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2017+v1996))))
	v2022 = v2020 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v2013) {
		goto L489
	} else {
		goto L490
	}
L488:
	;
	v2036 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2026+v1996))))
	v2038 = v2036 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v2013) {
		goto L494
	} else {
		goto L495
	}
L489:
	;
	v2026 = v2004 + int32(2)
	if v2026 != v1995 {
		goto L488
	} else {
		goto L492
	}
L490:
	;
	goto L491
L491:
	;
	v2070 = v2013<<(uint(int32(6))%32)&int32(1984) | v2022
	v2071 = int32(2)
	goto L485
L492:
	;
	goto L491
L493:
	;
	v2055 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1996+v2042))))
	v2070 = v2055&int32(63) | (v2013<<(uint(int32(18))%32)&int32(_a_F_italian_UTF_8_stem_7) | v2022<<(uint(int32(12))%32) | v2038<<(uint(int32(6))%32))
	v2071 = int32(4)
	goto L485
L494:
	;
	v2042 = v2004 + int32(3)
	if v2042 != v1995 {
		goto L493
	} else {
		goto L497
	}
L495:
	;
	goto L496
L496:
	;
	v2070 = v2013<<(uint(int32(12))%32)&int32(_a_F_italian_UTF_8_stem_8) | v2022<<(uint(int32(6))%32) | v2038
	v2071 = int32(3)
	goto L485
L497:
	;
	goto L496
L498:
	;
	v2088 = v2071 + v2004
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2088
	v2004 = v2088
	goto L480
L499:
	;
	v2075 = v2070 - int32(97)
	if v2075 < int32(0) {
		goto L498
	} else {
		goto L500
	}
L500:
	;
	v2081 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v2075)>>(uint(int32(3))%32)))+uint32(_c_F_italian_UTF_8_stem[0]))))
	if int32(base.Ui32(v2081)>>(uint(v2075&int32(7))%32))&int32(1) != 0 {
		goto L479
	} else {
		goto L501
	}
L501:
	;
	goto L498
L503:
	;
	v2102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2103 = v2102 + v2099
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2103
	v2117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2118 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2126 = v2103
	goto L506
L504:
	;
	if v2222 < int32(0) {
		goto L426
	} else {
		goto L528
	}
L505:
	;
	v2222 = v2193
	goto L504
L506:
	;
	if v2117 <= v2126 {
		goto L508
	} else {
		goto L509
	}
L508:
	;
	v2222 = int32(-1)
	goto L504
L509:
	;
	goto L510
L510:
	;
	v2133 = int32(1)
	v2135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2126+v2118))))
	if base.Ui32(v2135) < base.Ui32(int32(192)) {
		v2192 = v2135
		v2193 = v2133
		goto L511
	} else {
		goto L512
	}
L511:
	;
	if int32(249) < v2192 {
		goto L505
	} else {
		goto L524
	}
L512:
	;
	v2139 = v2126 + int32(1)
	if v2139 == v2117 {
		v2192 = v2135
		v2193 = v2133
		goto L511
	} else {
		goto L513
	}
L513:
	;
	v2142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2139+v2118))))
	v2144 = v2142 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v2135) {
		goto L515
	} else {
		goto L516
	}
L514:
	;
	v2158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2148+v2118))))
	v2160 = v2158 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v2135) {
		goto L520
	} else {
		goto L521
	}
L515:
	;
	v2148 = v2126 + int32(2)
	if v2148 != v2117 {
		goto L514
	} else {
		goto L518
	}
L516:
	;
	goto L517
L517:
	;
	v2192 = v2135<<(uint(int32(6))%32)&int32(1984) | v2144
	v2193 = int32(2)
	goto L511
L518:
	;
	goto L517
L519:
	;
	v2177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2118+v2164))))
	v2192 = v2177&int32(63) | (v2135<<(uint(int32(18))%32)&int32(_a_F_italian_UTF_8_stem_7) | v2144<<(uint(int32(12))%32) | v2160<<(uint(int32(6))%32))
	v2193 = int32(4)
	goto L511
L520:
	;
	v2164 = v2126 + int32(3)
	if v2164 != v2117 {
		goto L519
	} else {
		goto L523
	}
L521:
	;
	goto L522
L522:
	;
	v2192 = v2135<<(uint(int32(12))%32)&int32(_a_F_italian_UTF_8_stem_8) | v2144<<(uint(int32(6))%32) | v2160
	v2193 = int32(3)
	goto L511
L523:
	;
	goto L522
L524:
	;
	v2197 = v2192 - int32(97)
	if v2197 < int32(0) {
		goto L505
	} else {
		goto L525
	}
L525:
	;
	v2203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v2197)>>(uint(int32(3))%32)))+uint32(_c_F_italian_UTF_8_stem[0]))))
	if int32(base.Ui32(v2203)>>(uint(v2197&int32(7))%32))&int32(1) == int32(0) {
		goto L505
	} else {
		goto L526
	}
L526:
	;
	v2211 = v2193 + v2126
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2211
	v2126 = v2211
	goto L506
L528:
	;
	v2225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v2225 + v2222
	goto L426
L529:
	;
	v2290 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2290
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2290
	v2296 = F_find_among_b(m, l0, int32(_a_F_italian_UTF_8_stem_12), int32(51), int32(0))
	mBase = m.M
	v2297 = m.ExcPending
	if v2297 != 0 {
		goto L5
	} else {
		goto L546
	}
L530:
	;
	v2236 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2236+v2234))))
	if base.B2i32(v2238&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v2238)%32)&int32(_a_F_italian_UTF_8_stem_13) == int32(0)) != 0 {
		goto L529
	} else {
		goto L531
	}
L531:
	;
	v2253 = F_find_among_b(m, l0, int32(_a_F_italian_UTF_8_stem_14), int32(37), int32(0))
	mBase = m.M
	v2254 = m.ExcPending
	if v2254 != 0 {
		goto L5
	} else {
		goto L532
	}
L532:
	;
	if v2253 == int32(0) {
		goto L529
	} else {
		goto L533
	}
L533:
	;
	v2257 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2257
	v2260 = v2257 - int32(1)
	v2261 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2260 <= v2261 {
		goto L529
	} else {
		goto L534
	}
L534:
	;
	v2263 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2263+v2260))))
	switch v2265 - int32(111) {
	case 0, 3:
		goto L535
	default:
		goto L529
	}
L535:
	;
	v2271 = F_find_among_b(m, l0, int32(_a_F_italian_UTF_8_stem_15), int32(5), int32(0))
	mBase = m.M
	v2272 = m.ExcPending
	if v2272 != 0 {
		goto L5
	} else {
		goto L536
	}
L536:
	;
	if v2271 == int32(0) {
		goto L529
	} else {
		goto L537
	}
L537:
	;
	v2275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v2276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v2276 < v2275 {
		goto L529
	} else {
		goto L538
	}
L538:
	;
	switch v2271 - int32(1) {
	case 0:
		goto L540
	case 1:
		goto L539
	default:
		goto L529
	}
L539:
	;
	v2285 = F_slice_from_s(m, l0, int32(1), int32(_a_F_italian_UTF_8_stem_16))
	mBase = m.M
	v2286 = m.ExcPending
	if v2286 != 0 {
		goto L5
	} else {
		goto L542
	}
L540:
	;
	v2280 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v2280 {
		goto L529
	} else {
		goto L541
	}
L541:
	;
	v2996 = v2280
	goto L1
L542:
	;
	if v2285 < int32(0) {
		v2996 = v2285
		goto L1
	} else {
		goto L543
	}
L543:
	;
	goto L529
L544:
	;
	v2559 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2559
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2559
	v2575 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v2576 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L627
L545:
	;
	v2537 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2537
	v2539 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v2537 < v2539 {
		goto L544
	} else {
		goto L619
	}
L546:
	;
	if v2296 == int32(0) {
		goto L545
	} else {
		goto L547
	}
L547:
	;
	v2300 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2300
	switch v2296 - int32(1) {
	case 0:
		goto L556
	case 1:
		goto L555
	case 2:
		goto L554
	case 3:
		goto L553
	case 4:
		goto L552
	case 5:
		goto L551
	case 6:
		goto L550
	case 7:
		goto L549
	case 8:
		goto L548
	default:
		goto L544
	}
L548:
	;
	v2477 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v2300 < v2477 {
		goto L545
	} else {
		goto L603
	}
L549:
	;
	v2438 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v2300 < v2438 {
		goto L545
	} else {
		goto L595
	}
L550:
	;
	v2370 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v2300 < v2370 {
		goto L545
	} else {
		goto L579
	}
L551:
	;
	v2365 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v2300 < v2365 {
		goto L545
	} else {
		goto L577
	}
L552:
	;
	v2357 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v2300 < v2357 {
		goto L545
	} else {
		goto L574
	}
L553:
	;
	v2349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v2300 < v2349 {
		goto L545
	} else {
		goto L571
	}
L554:
	;
	v2341 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v2300 < v2341 {
		goto L545
	} else {
		goto L568
	}
L555:
	;
	v2309 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v2300 < v2309 {
		goto L545
	} else {
		goto L559
	}
L556:
	;
	v2304 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v2300 < v2304 {
		goto L545
	} else {
		goto L557
	}
L557:
	;
	v2306 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v2306 {
		goto L544
	} else {
		goto L558
	}
L558:
	;
	v2996 = v2306
	goto L1
L559:
	;
	v2311 = F_slice_del(m, l0)
	mBase = m.M
	if v2311 < int32(0) {
		v2996 = v2311
		goto L1
	} else {
		goto L560
	}
L560:
	;
	v2314 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2314
	v2316 = int32(2)
	v2318 = int32(0)
	v2321 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2314-v2321 < v2316 {
		v2331 = v2318
		goto L562
	} else {
		goto L563
	}
L561:
	;
	if v2331 == int32(0) {
		goto L544
	} else {
		goto L565
	}
L562:
	;
	goto L561
L563:
	;
	v2324 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2327 = F_memcmp(m, v2324+v2314-v2316, int32(_a_F_italian_UTF_8_stem_17), v2316)
	mBase = m.M
	if v2327 != 0 {
		v2331 = v2318
		goto L562
	} else {
		goto L564
	}
L564:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2314 - v2316
	v2331 = int32(1)
	goto L562
L565:
	;
	v2334 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2334
	v2336 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v2334 < v2336 {
		goto L544
	} else {
		goto L566
	}
L566:
	;
	v2338 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v2338 {
		goto L544
	} else {
		goto L567
	}
L567:
	;
	v2996 = v2338
	goto L1
L568:
	;
	v2345 = F_slice_from_s(m, l0, int32(3), int32(_a_F_italian_UTF_8_stem_18))
	mBase = m.M
	v2346 = m.ExcPending
	if v2346 != 0 {
		goto L5
	} else {
		goto L569
	}
L569:
	;
	if int32(0) <= v2345 {
		goto L544
	} else {
		goto L570
	}
L570:
	;
	v2996 = v2345
	goto L1
L571:
	;
	v2353 = F_slice_from_s(m, l0, int32(1), int32(_a_F_italian_UTF_8_stem_19))
	mBase = m.M
	v2354 = m.ExcPending
	if v2354 != 0 {
		goto L5
	} else {
		goto L572
	}
L572:
	;
	if int32(0) <= v2353 {
		goto L544
	} else {
		goto L573
	}
L573:
	;
	v2996 = v2353
	goto L1
L574:
	;
	v2361 = F_slice_from_s(m, l0, int32(4), int32(_a_F_italian_UTF_8_stem_20))
	mBase = m.M
	v2362 = m.ExcPending
	if v2362 != 0 {
		goto L5
	} else {
		goto L575
	}
L575:
	;
	if int32(0) <= v2361 {
		goto L544
	} else {
		goto L576
	}
L576:
	;
	v2996 = v2361
	goto L1
L577:
	;
	v2367 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v2367 {
		goto L544
	} else {
		goto L578
	}
L578:
	;
	v2996 = v2367
	goto L1
L579:
	;
	v2372 = F_slice_del(m, l0)
	mBase = m.M
	if v2372 < int32(0) {
		v2996 = v2372
		goto L1
	} else {
		goto L580
	}
L580:
	;
	v2375 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2375
	v2378 = v2375 - int32(1)
	v2379 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2378 <= v2379 {
		goto L544
	} else {
		goto L581
	}
L581:
	;
	v2381 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2381+v2378))))
	if base.B2i32(v2383&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v2383)%32)&int32(_a_F_italian_UTF_8_stem_21) == int32(0)) != 0 {
		goto L544
	} else {
		goto L582
	}
L582:
	;
	v2398 = F_find_among_b(m, l0, int32(_a_F_italian_UTF_8_stem_22), int32(4), int32(0))
	mBase = m.M
	v2399 = m.ExcPending
	if v2399 != 0 {
		goto L5
	} else {
		goto L583
	}
L583:
	;
	if v2398 == int32(0) {
		goto L544
	} else {
		goto L584
	}
L584:
	;
	v2402 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2402
	v2404 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v2402 < v2404 {
		goto L544
	} else {
		goto L585
	}
L585:
	;
	v2406 = F_slice_del(m, l0)
	mBase = m.M
	if v2406 < int32(0) {
		v2996 = v2406
		goto L1
	} else {
		goto L586
	}
L586:
	;
	if v2398 != int32(1) {
		goto L544
	} else {
		goto L587
	}
L587:
	;
	v2411 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2411
	v2413 = int32(2)
	v2415 = int32(0)
	v2418 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2411-v2418 < v2413 {
		v2428 = v2415
		goto L589
	} else {
		goto L590
	}
L588:
	;
	if v2428 == int32(0) {
		goto L544
	} else {
		goto L592
	}
L589:
	;
	goto L588
L590:
	;
	v2421 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2424 = F_memcmp(m, v2421+v2411-v2413, int32(_a_F_italian_UTF_8_stem_23), v2413)
	mBase = m.M
	if v2424 != 0 {
		v2428 = v2415
		goto L589
	} else {
		goto L591
	}
L591:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2411 - v2413
	v2428 = int32(1)
	goto L589
L592:
	;
	v2431 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2431
	v2433 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v2431 < v2433 {
		goto L544
	} else {
		goto L593
	}
L593:
	;
	v2435 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v2435 {
		goto L544
	} else {
		goto L594
	}
L594:
	;
	v2996 = v2435
	goto L1
L595:
	;
	v2440 = F_slice_del(m, l0)
	mBase = m.M
	if v2440 < int32(0) {
		v2996 = v2440
		goto L1
	} else {
		goto L596
	}
L596:
	;
	v2443 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2443
	v2446 = v2443 - int32(1)
	v2447 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2446 <= v2447 {
		goto L544
	} else {
		goto L597
	}
L597:
	;
	v2449 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2451 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2449+v2446))))
	if base.B2i32(v2451&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v2451)%32)&int32(_a_F_italian_UTF_8_stem_24) == int32(0)) != 0 {
		goto L544
	} else {
		goto L598
	}
L598:
	;
	v2466 = F_find_among_b(m, l0, int32(_a_F_italian_UTF_8_stem_25), int32(3), int32(0))
	mBase = m.M
	v2467 = m.ExcPending
	if v2467 != 0 {
		goto L5
	} else {
		goto L599
	}
L599:
	;
	if v2466 == int32(0) {
		goto L544
	} else {
		goto L600
	}
L600:
	;
	v2470 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2470
	v2472 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v2470 < v2472 {
		goto L544
	} else {
		goto L601
	}
L601:
	;
	v2474 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v2474 {
		goto L544
	} else {
		goto L602
	}
L602:
	;
	v2996 = v2474
	goto L1
L603:
	;
	v2479 = F_slice_del(m, l0)
	mBase = m.M
	if v2479 < int32(0) {
		v2996 = v2479
		goto L1
	} else {
		goto L604
	}
L604:
	;
	v2482 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2482
	v2484 = int32(2)
	v2486 = int32(0)
	v2489 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2482-v2489 < v2484 {
		v2499 = v2486
		goto L606
	} else {
		goto L607
	}
L605:
	;
	if v2499 == int32(0) {
		goto L544
	} else {
		goto L609
	}
L606:
	;
	goto L605
L607:
	;
	v2492 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2495 = F_memcmp(m, v2492+v2482-v2484, int32(_a_F_italian_UTF_8_stem_26), v2484)
	mBase = m.M
	if v2495 != 0 {
		v2499 = v2486
		goto L606
	} else {
		goto L608
	}
L608:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2482 - v2484
	v2499 = int32(1)
	goto L606
L609:
	;
	v2502 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2502
	v2504 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v2502 < v2504 {
		goto L544
	} else {
		goto L610
	}
L610:
	;
	v2506 = F_slice_del(m, l0)
	mBase = m.M
	if v2506 < int32(0) {
		v2996 = v2506
		goto L1
	} else {
		goto L611
	}
L611:
	;
	v2509 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2509
	v2511 = int32(2)
	v2513 = int32(0)
	v2516 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2509-v2516 < v2511 {
		v2526 = v2513
		goto L613
	} else {
		goto L614
	}
L612:
	;
	if v2526 == int32(0) {
		goto L544
	} else {
		goto L616
	}
L613:
	;
	goto L612
L614:
	;
	v2519 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2522 = F_memcmp(m, v2519+v2509-v2511, int32(_a_F_italian_UTF_8_stem_27), v2511)
	mBase = m.M
	if v2522 != 0 {
		v2526 = v2513
		goto L613
	} else {
		goto L615
	}
L615:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2509 - v2511
	v2526 = int32(1)
	goto L613
L616:
	;
	v2529 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2529
	v2531 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v2529 < v2531 {
		goto L544
	} else {
		goto L617
	}
L617:
	;
	v2533 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v2533 {
		goto L544
	} else {
		goto L618
	}
L618:
	;
	v2996 = v2533
	goto L1
L619:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2537
	v2542 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v2539
	v2547 = F_find_among_b(m, l0, int32(_a_F_italian_UTF_8_stem_28), int32(87), int32(0))
	mBase = m.M
	v2548 = m.ExcPending
	if v2548 != 0 {
		goto L5
	} else {
		goto L620
	}
L620:
	;
	if v2547 != 0 {
		goto L621
	} else {
		goto L622
	}
L621:
	;
	v2549 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2549
	v2551 = F_slice_del(m, l0)
	mBase = m.M
	if v2551 < int32(0) {
		v2996 = v2551
		goto L1
	} else {
		goto L624
	}
L622:
	;
	goto L623
L623:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v2542
	goto L544
L624:
	;
	goto L623
L625:
	;
	if v2690 != 0 {
		goto L53
	} else {
		goto L648
	}
L626:
	;
	v2690 = v2683
	goto L625
L627:
	;
	if v2559 <= v2575 {
		v2683 = int32(-1)
		goto L626
	} else {
		goto L629
	}
L628:
	;
	v2683 = int32(0)
	goto L626
L629:
	;
	v2592 = int32(1)
	v2593 = v2559 - v2592
	v2595 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2576+v2593))))
	v2597 = v2595 & int32(255)
	if base.B2i32(v2593 == v2575)|base.B2i32(int32(0) <= v2595) != 0 {
		v2655 = v2597
		v2659 = v2592
		goto L630
	} else {
		goto L631
	}
L630:
	;
	if int32(242) < v2655 {
		goto L638
	} else {
		goto L639
	}
L631:
	;
	v2604 = v2597 & int32(63)
	v2606 = v2559 - int32(2)
	v2608 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2576+v2606))))
	v2610 = v2608 << (uint(int32(6)) % 32)
	if base.B2i32(v2606 != v2575)&base.B2i32(base.Ui32(v2608) < base.Ui32(int32(192))) == int32(0) {
		goto L632
	} else {
		goto L633
	}
L632:
	;
	v2655 = v2610&int32(1984) | v2604
	v2659 = int32(2)
	goto L630
L633:
	;
	goto L634
L634:
	;
	v2623 = v2610&int32(4032) | v2604
	v2625 = v2559 - int32(3)
	v2627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2576+v2625))))
	if base.B2i32(v2625 != v2575)&base.B2i32(base.Ui32(v2627) < base.Ui32(int32(224))) == int32(0) {
		goto L635
	} else {
		goto L636
	}
L635:
	;
	v2655 = v2627<<(uint(int32(12))%32)&int32(_a_F_italian_UTF_8_stem_8) | v2623
	v2659 = int32(3)
	goto L630
L636:
	;
	goto L637
L637:
	;
	v2645 = int32(4)
	v2647 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2559+v2576-v2645))))
	v2655 = v2627<<(uint(int32(12))%32)&int32(_a_F_italian_UTF_8_stem_29) | v2647&int32(7)<<(uint(int32(18))%32) | v2623
	v2659 = v2645
	goto L630
L638:
	;
	v2690 = v2659
	goto L625
L639:
	;
	goto L640
L640:
	;
	v2661 = v2655 - int32(97)
	if v2661 < int32(0) {
		goto L641
	} else {
		goto L642
	}
L641:
	;
	v2690 = v2659
	goto L625
L642:
	;
	goto L643
L643:
	;
	v2667 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v2661)>>(uint(int32(3))%32)))+uint32(_c_F_italian_UTF_8_stem[1]))))
	if int32(base.Ui32(v2667)>>(uint(v2661&int32(7))%32))&int32(1) == int32(0) {
		goto L644
	} else {
		goto L645
	}
L644:
	;
	v2690 = v2659
	goto L625
L645:
	;
	goto L646
L646:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2559 - v2659
	goto L647
L647:
	;
	goto L628
L648:
	;
	v2691 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2691
	v2693 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v2691 < v2693 {
		goto L53
	} else {
		goto L649
	}
L649:
	;
	v2695 = F_slice_del(m, l0)
	mBase = m.M
	if v2695 < int32(0) {
		v2996 = v2695
		goto L1
	} else {
		goto L650
	}
L650:
	;
	v2698 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2698
	v2700 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2698 <= v2700 {
		goto L53
	} else {
		goto L651
	}
L651:
	;
	v2702 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2706 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2702+v2698-int32(1)))))
	if v2706 != int32(105) {
		goto L53
	} else {
		goto L652
	}
L652:
	;
	v2710 = v2698 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2710
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2710
	v2713 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v2698 <= v2713 {
		goto L53
	} else {
		goto L653
	}
L653:
	;
	v2715 = F_slice_del(m, l0)
	mBase = m.M
	if v2715 < int32(0) {
		v2996 = v2715
		goto L1
	} else {
		goto L654
	}
L654:
	;
	v2718 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2729 = v2718
	goto L50
L655:
	;
	if int32(0) <= v2721 {
		goto L51
	} else {
		goto L656
	}
L656:
	;
	v2996 = v2721
	goto L1
L657:
	;
	v2881 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2881
	goto L687
L658:
	;
	v2734 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2738 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2734+v2729-int32(1)))))
	if v2738 != int32(104) {
		goto L657
	} else {
		goto L659
	}
L659:
	;
	v2742 = v2729 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2742
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2742
	v2758 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v2759 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L662
L660:
	;
	if v2873 != 0 {
		goto L657
	} else {
		goto L683
	}
L661:
	;
	v2873 = v2866
	goto L660
L662:
	;
	if v2742 <= v2758 {
		v2866 = int32(-1)
		goto L661
	} else {
		goto L664
	}
L663:
	;
	v2866 = int32(0)
	goto L661
L664:
	;
	v2775 = int32(1)
	v2776 = v2742 - v2775
	v2778 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2759+v2776))))
	v2780 = v2778 & int32(255)
	if base.B2i32(v2776 == v2758)|base.B2i32(int32(0) <= v2778) != 0 {
		v2838 = v2780
		v2842 = v2775
		goto L665
	} else {
		goto L666
	}
L665:
	;
	if int32(103) < v2838 {
		goto L673
	} else {
		goto L674
	}
L666:
	;
	v2787 = v2780 & int32(63)
	v2789 = v2742 - int32(2)
	v2791 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2759+v2789))))
	v2793 = v2791 << (uint(int32(6)) % 32)
	if base.B2i32(v2789 != v2758)&base.B2i32(base.Ui32(v2791) < base.Ui32(int32(192))) == int32(0) {
		goto L667
	} else {
		goto L668
	}
L667:
	;
	v2838 = v2793&int32(1984) | v2787
	v2842 = int32(2)
	goto L665
L668:
	;
	goto L669
L669:
	;
	v2806 = v2793&int32(4032) | v2787
	v2808 = v2742 - int32(3)
	v2810 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2759+v2808))))
	if base.B2i32(v2808 != v2758)&base.B2i32(base.Ui32(v2810) < base.Ui32(int32(224))) == int32(0) {
		goto L670
	} else {
		goto L671
	}
L670:
	;
	v2838 = v2810<<(uint(int32(12))%32)&int32(_a_F_italian_UTF_8_stem_8) | v2806
	v2842 = int32(3)
	goto L665
L671:
	;
	goto L672
L672:
	;
	v2828 = int32(4)
	v2830 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2742+v2759-v2828))))
	v2838 = v2810<<(uint(int32(12))%32)&int32(_a_F_italian_UTF_8_stem_29) | v2830&int32(7)<<(uint(int32(18))%32) | v2806
	v2842 = v2828
	goto L665
L673:
	;
	v2873 = v2842
	goto L660
L674:
	;
	goto L675
L675:
	;
	v2844 = v2838 - int32(99)
	if v2844 < int32(0) {
		goto L676
	} else {
		goto L677
	}
L676:
	;
	v2873 = v2842
	goto L660
L677:
	;
	goto L678
L678:
	;
	v2850 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v2844)>>(uint(int32(3))%32)))+uint32(_c_F_italian_UTF_8_stem[2]))))
	if int32(base.Ui32(v2850)>>(uint(v2844&int32(7))%32))&int32(1) == int32(0) {
		goto L679
	} else {
		goto L680
	}
L679:
	;
	v2873 = v2842
	goto L660
L680:
	;
	goto L681
L681:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2742 - v2842
	goto L682
L682:
	;
	goto L663
L683:
	;
	v2874 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v2875 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v2875 < v2874 {
		goto L657
	} else {
		goto L684
	}
L684:
	;
	v2877 = F_slice_del(m, l0)
	mBase = m.M
	if v2877 < int32(0) {
		v2996 = v2877
		goto L1
	} else {
		goto L685
	}
L685:
	;
	goto L657
L686:
	;
	if v2988 < int32(0) {
		v2996 = v2988
		goto L1
	} else {
		goto L724
	}
L687:
	;
	v2888 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2888
	v2890 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v2890 <= v2888 {
		goto L690
	} else {
		goto L691
	}
L688:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2888
	v2988 = int32(1)
	goto L686
L689:
	;
	v2930 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L703
L690:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2888
	v2928 = v2888
	v2929 = v2890
	goto L689
L691:
	;
	v2892 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2894 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2892+v2888))))
	v2896 = v2894 - int32(73)
	v2897 = int32(0)
	if base.B2i32(v2896 == v2897)|base.B2i32(v2896 == int32(12)) == v2897 {
		goto L690
	} else {
		goto L692
	}
L692:
	;
	v2907 = F_find_among(m, l0, int32(_a_F_italian_UTF_8_stem_30), int32(3), int32(0))
	mBase = m.M
	v2908 = m.ExcPending
	if v2908 != 0 {
		goto L5
	} else {
		goto L693
	}
L693:
	;
	v2909 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2909
	switch v2907 - int32(1) {
	case 0:
		goto L695
	case 1:
		goto L694
	case 2:
		goto L696
	default:
		goto L687
	}
L694:
	;
	v2922 = F_slice_from_s(m, l0, int32(1), int32(_a_F_italian_UTF_8_stem_31))
	mBase = m.M
	v2923 = m.ExcPending
	if v2923 != 0 {
		goto L5
	} else {
		goto L699
	}
L695:
	;
	v2916 = F_slice_from_s(m, l0, int32(1), int32(_a_F_italian_UTF_8_stem_32))
	mBase = m.M
	v2917 = m.ExcPending
	if v2917 != 0 {
		goto L5
	} else {
		goto L697
	}
L696:
	;
	v2913 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2928 = v2909
	v2929 = v2913
	goto L689
L697:
	;
	if int32(0) <= v2916 {
		goto L687
	} else {
		goto L698
	}
L698:
	;
	v2988 = v2916
	goto L686
L699:
	;
	if int32(0) <= v2922 {
		goto L687
	} else {
		goto L700
	}
L700:
	;
	v2988 = v2922
	goto L686
L701:
	;
	if int32(0) <= v2982 {
		goto L721
	} else {
		goto L722
	}
L703:
	;
	goto L704
L704:
	;
	goto L705
L705:
	;
	v2937 = v2928
	v2939 = int32(1)
	goto L708
L707:
	;
	v2982 = v2967
	goto L701
L708:
	;
	if v2929 <= v2937 {
		goto L710
	} else {
		goto L711
	}
L709:
	;
	goto L707
L710:
	;
	v2982 = int32(-1)
	goto L701
L711:
	;
	goto L712
L712:
	;
	v2944 = v2937 + int32(1)
	v2946 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2930+v2937))))
	if base.Ui32(v2946) < base.Ui32(int32(192)) {
		v2967 = v2944
		goto L713
	} else {
		goto L714
	}
L713:
	;
	v2968 = int32(1)
	if v2968 < v2939 {
		v2937 = v2967
		v2939 = v2939 - v2968
		goto L708
	} else {
		goto L720
	}
L714:
	;
	if v2929 <= v2944 {
		v2967 = v2944
		goto L713
	} else {
		goto L715
	}
L715:
	;
	v2953 = v2944
	goto L716
L716:
	;
	v2956 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2930+v2953))))
	if int32(-65) < v2956 {
		v2967 = v2953
		goto L713
	} else {
		goto L718
	}
L717:
	;
	v2967 = v2929
	goto L713
L718:
	;
	v2960 = v2953 + int32(1)
	if v2960 != v2929 {
		v2953 = v2960
		goto L716
	} else {
		goto L719
	}
L719:
	;
	goto L717
L720:
	;
	goto L709
L721:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2982
	goto L687
L722:
	;
	goto L723
L723:
	;
	goto L688
L724:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2881
	v2996 = int32(1)
	goto L1
}
func F_iterate_values_object_field_start(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	if v4&int32(1) != 0 {
		v7 = F_pstrdup(m, l1)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v12 = F_strlen(m, v7)
			mBase = m.M
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			m.T0[v13].(func(*base.Module, int32, int32, int32))(m, v11, v7, v12)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				return int32(0)
			}
		}
	} else {
		return int32(0)
	}
}
func F_ivfflatbuildempty(m *base.Module, l0 int32) {
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v3 = m.G0
	v5 = v3 - int32(176)
	m.G0 = v5
	v8 = F_BuildIndexInfo(m, l0)
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		F_BuildIndex_2(m, int32(0), l0, v8, v5, int32(3))
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			m.G0 = v5 + int32(176)
			return
		}
	}
}
func F_ivfflatcostestimate(m *base.Module, l0 int32, l1 int32, l2 float64, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 float64
	_ = v52
	var v55 float64
	_ = v55
	var v56 float64
	_ = v56
	var v57 float64
	_ = v57
	var v59 float64
	_ = v59
	var v60 float64
	_ = v60
	var v61 float64
	_ = v61
	var v64 float64
	_ = v64
	var v67 float64
	_ = v67
	var v68 float64
	_ = v68
	var v73 float64
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 float64
	_ = v77
	var v88 float64
	_ = v88
	var v93 float64
	_ = v93
	var v95 float64
	_ = v95
	var v98 int64
	_ = v98
	var v102 int64
	_ = v102
	v17 = m.G0
	v19 = v17 - int32(96)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	if v21 != 0 {
		v23 = v19 + int32(24)
		base.MemoryFill(m, v23, int32(0), int32(72))
		F_genericcostestimate(m, l0, l1, l2, v23)
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return
		} else {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
			v32 = F_index_open(m, v30, int32(0))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return
			} else {
				F_IvfflatGetMetaPageInfo(m, v32, v19+int32(20), int32(0))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					F_relation_close(m, v32, int32(0))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return
					} else {
						v43 = *(*int32)(unsafe.Add(mBase, _c_F_ivfflatcostestimate[0]))
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
						v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
						F_get_tablespace_page_costs(m, v46, int32(0), v19+int32(8))
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return
						} else {
							v52 = *(*float64)(unsafe.Add(mBase, uint32(v19)+56))
							v55 = *(*float64)(unsafe.Add(mBase, uint32(v19)+72))
							v56 = *(*float64)(unsafe.Add(mBase, uint32(v19)+8))
							v57 = base.F64_sub(v55, v56)
							v59 = *(*float64)(unsafe.Add(mBase, uint32(v19)+32))
							v60 = base.F64_add(base.F64_mul(base.F64_mul(v52, float64(-0.5)), v57), v59)
							v61 = float64(1)
							v64 = base.F64_div(base.F64_convert_i32_s(v43), base.F64_convert_i32_s(v44))
							if base.F64_gt(v64, v61) != 0 {
								v67 = v61
							} else {
								v67 = v64
							}
							v68 = base.F64_mul(v60, v67)
							if base.F64_lt(v67, float64(0.5)) == int32(0) {
								v88 = v68
							} else {
								v73 = base.F64_mul(v52, v67)
								v74 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
								v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
								v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+124))
								v77 = base.F64_convert_i32_u(v76)
								if base.F64_gt(v73, v77) == int32(0) {
									v88 = v68
								} else {
									v88 = base.F64_add(base.F64_mul(base.F64_sub(v77, v73), v56), base.F64_add(base.F64_mul(base.F64_mul(v73, float64(-0.5)), v57), v68))
								}
							}
							*(*float64)(unsafe.Add(mBase, uint32(l3))) = v88
							*(*float64)(unsafe.Add(mBase, uint32(l4))) = v60
							v93 = *(*float64)(unsafe.Add(mBase, uint32(v19)+40))
							*(*float64)(unsafe.Add(mBase, uint32(l5))) = v93
							v95 = *(*float64)(unsafe.Add(mBase, uint32(v19)+48))
							*(*float64)(unsafe.Add(mBase, uint32(l6))) = v95
							*(*float64)(unsafe.Add(mBase, uint32(l7))) = v52
							m.G0 = v19 + int32(96)
							return
						}
					}
				}
			}
		}
	} else {
		v98 = int64(9218868437227405312)
		*(*int64)(unsafe.Add(mBase, uint32(l3))) = v98
		*(*int64)(unsafe.Add(mBase, uint32(l4))) = v98
		v102 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(l5))) = v102
		*(*int64)(unsafe.Add(mBase, uint32(l6))) = v102
		*(*int64)(unsafe.Add(mBase, uint32(l7))) = v102
		*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = int32(2)
		m.G0 = v19 + int32(96)
		return
	}
}
