package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_InitLocalBuffers(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v25 int64
	_ = v25
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v60 int64
	_ = v60
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v95 int64
	_ = v95
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v121 int32
	_ = v121
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v153 int32
	_ = v153
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_InitLocalBuffers[0]))
	if v11 < int32(0) {
		v16 = *(*int32)(unsafe.Add(mBase, _c_F_InitLocalBuffers[1]))
		v17 = int32(64)
		v19 = int32(0)
		if v16 == v19 {
			v37 = v19
		} else {
			v25 = base.I64_extend_i32_u(v16) * base.I64_extend_i32_u(v17)
			v26 = base.I32_wrap_i64(v25)
			if base.Ui32(v16|v17) < base.Ui32(int32(_a_F_InitLocalBuffers_0)) {
				v37 = v26
			} else {
				if base.I32_wrap_i64(int64(base.Ui64(v25)>>(uint(int64(32))%64))) != 0 {
					v34 = int32(-1)
				} else {
					v34 = v26
				}
				v37 = v34
			}
		}
		v38 = F_emscripten_builtin_malloc(m, v37)
		mBase = m.M
		if v38 == int32(0) {
		} else {
			v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38-int32(4)))))
			if v43&int32(3) == int32(0) {
			} else {
				F___memset(m, v38, int32(0), v37)
				mBase = m.M
			}
		}
		*(*int32)(unsafe.Add(mBase, _c_F_InitLocalBuffers[2])) = v38
		v52 = int32(4)
		v54 = int32(0)
		if v16 == v54 {
			v72 = v54
		} else {
			v60 = base.I64_extend_i32_u(v16) * base.I64_extend_i32_u(v52)
			v61 = base.I32_wrap_i64(v60)
			if base.Ui32(v16|v52) < base.Ui32(int32(_a_F_InitLocalBuffers_0)) {
				v72 = v61
			} else {
				if base.I32_wrap_i64(int64(base.Ui64(v60)>>(uint(int64(32))%64))) != 0 {
					v69 = int32(-1)
				} else {
					v69 = v61
				}
				v72 = v69
			}
		}
		v73 = F_emscripten_builtin_malloc(m, v72)
		mBase = m.M
		if v73 == int32(0) {
		} else {
			v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73-int32(4)))))
			if v78&int32(3) == int32(0) {
			} else {
				F___memset(m, v73, int32(0), v72)
				mBase = m.M
			}
		}
		*(*int32)(unsafe.Add(mBase, _c_F_InitLocalBuffers[3])) = v73
		v87 = int32(4)
		v89 = int32(0)
		if v16 == v89 {
			v107 = v89
		} else {
			v95 = base.I64_extend_i32_u(v16) * base.I64_extend_i32_u(v87)
			v96 = base.I32_wrap_i64(v95)
			if base.Ui32(v16|v87) < base.Ui32(int32(_a_F_InitLocalBuffers_0)) {
				v107 = v96
			} else {
				if base.I32_wrap_i64(int64(base.Ui64(v95)>>(uint(int64(32))%64))) != 0 {
					v104 = int32(-1)
				} else {
					v104 = v96
				}
				v107 = v104
			}
		}
		v108 = F_emscripten_builtin_malloc(m, v107)
		mBase = m.M
		if v108 == int32(0) {
		} else {
			v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108-int32(4)))))
			if v113&int32(3) == int32(0) {
			} else {
				F___memset(m, v108, int32(0), v107)
				mBase = m.M
			}
		}
		*(*int32)(unsafe.Add(mBase, _c_F_InitLocalBuffers[4])) = v108
		v121 = int32(0)
		if base.B2i32(v108 == v121)|(base.B2i32(v38 == v121)|base.B2i32(v73 == v121)) != 0 {
			F_errstart_cold(m, int32(22), int32(0))
			mBase = m.M
			v194 = m.ExcPending
			if v194 != 0 {
				return
			} else {
				F_errcode(m, int32(_a_F_InitLocalBuffers_1))
				mBase = m.M
				v197 = m.ExcPending
				if v197 != 0 {
					return
				} else {
					F_errmsg(m, int32(_a_F_InitLocalBuffers_2), int32(0))
					mBase = m.M
					v201 = m.ExcPending
					if v201 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_InitLocalBuffers_3), int32(745), int32(_a_F_InitLocalBuffers_4))
						mBase = m.M
						v206 = m.ExcPending
						if v206 != 0 {
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
			v129 = int32(0)
			*(*int32)(unsafe.Add(mBase, _c_F_InitLocalBuffers[5])) = v129
			if v129 < v16 {
				v135 = v129
				for {
					v141 = *(*int32)(unsafe.Add(mBase, _c_F_InitLocalBuffers[2]))
					v144 = v141 + v135<<(uint(int32(6))%32)
					*(*int32)(unsafe.Add(mBase, uint32(v144)+20)) = int32(-2) - v135
					*(*int32)(unsafe.Add(mBase, uint32(v144+int32(36)))) = int32(-1)
					v153 = v135 + int32(1)
					if v153 != v16 {
						v135 = v153
						continue
					} else {
						break
					}
					break
				}
			} else {
			}
			*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = int64(103079215124)
			v165 = F_hash_create(m, int32(_a_F_InitLocalBuffers_5), v16, v8, int32(40))
			mBase = m.M
			v166 = m.ExcPending
			if v166 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_InitLocalBuffers[6])) = v165
				if v165 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v210 = m.ExcPending
					if v210 != 0 {
						return
					} else {
						F_errmsg_internal(m, int32(_a_F_InitLocalBuffers_6), int32(0))
						mBase = m.M
						v214 = m.ExcPending
						if v214 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_InitLocalBuffers_3), int32(782), int32(_a_F_InitLocalBuffers_4))
							mBase = m.M
							v219 = m.ExcPending
							if v219 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_InitLocalBuffers[7])) = v16
					m.G0 = v8 + int32(48)
					return
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v178 = m.ExcPending
		if v178 != 0 {
			return
		} else {
			F_errcode(m, int32(322))
			mBase = m.M
			v181 = m.ExcPending
			if v181 != 0 {
				return
			} else {
				F_errmsg(m, int32(_a_F_InitLocalBuffers_7), int32(0))
				mBase = m.M
				v185 = m.ExcPending
				if v185 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_InitLocalBuffers_3), int32(736), int32(_a_F_InitLocalBuffers_4))
					mBase = m.M
					v190 = m.ExcPending
					if v190 != 0 {
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
