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
	var v18 int32
	_ = v18
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
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
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
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
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
	var v119 int32
	_ = v119
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v151 int32
	_ = v151
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _consts[172]))
	if v11 < int32(0) {
		v16 = *(*int32)(unsafe.Add(mBase, _consts[933]))
		v17 = int32(64)
		v18 = int32(0)
		if v16 == v18 {
			v37 = v18
		} else {
			v25 = base.I64_extend_i32_u(v16) * base.I64_extend_i32_u(v17)
			v26 = base.I32_wrap_i64(v25)
			if base.Ui32(v16|v17) < base.Ui32(int32(65536)) {
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
				v49 = F___memset(m, v38, int32(0), v37)
				mBase = m.M
			}
		}
		*(*int32)(unsafe.Add(mBase, _consts[3])) = v38
		v52 = int32(4)
		v53 = int32(0)
		if v16 == v53 {
			v72 = v53
		} else {
			v60 = base.I64_extend_i32_u(v16) * base.I64_extend_i32_u(v52)
			v61 = base.I32_wrap_i64(v60)
			if base.Ui32(v16|v52) < base.Ui32(int32(65536)) {
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
				v84 = F___memset(m, v73, int32(0), v72)
				mBase = m.M
			}
		}
		*(*int32)(unsafe.Add(mBase, _consts[1])) = v73
		v87 = int32(4)
		v88 = int32(0)
		if v16 == v88 {
			v107 = v88
		} else {
			v95 = base.I64_extend_i32_u(v16) * base.I64_extend_i32_u(v87)
			v96 = base.I32_wrap_i64(v95)
			if base.Ui32(v16|v87) < base.Ui32(int32(65536)) {
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
				v119 = F___memset(m, v108, int32(0), v107)
				mBase = m.M
			}
		}
		*(*int32)(unsafe.Add(mBase, _consts[932])) = v108
		if v38 == int32(0) {
			F_errstart_cold(m, int32(22), int32(0))
			mBase = m.M
			v192 = m.ExcPending
			if v192 != 0 {
				return
			} else {
				F_errcode(m, int32(8389))
				mBase = m.M
				v195 = m.ExcPending
				if v195 != 0 {
					return
				} else {
					F_errmsg(m, int32(13891), int32(0))
					mBase = m.M
					v199 = m.ExcPending
					if v199 != 0 {
						return
					} else {
						F_errfinish(m, int32(496488), int32(745), int32(135050))
						mBase = m.M
						v204 = m.ExcPending
						if v204 != 0 {
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
			if v73 == int32(0) {
				F_errstart_cold(m, int32(22), int32(0))
				mBase = m.M
				v192 = m.ExcPending
				if v192 != 0 {
					return
				} else {
					F_errcode(m, int32(8389))
					mBase = m.M
					v195 = m.ExcPending
					if v195 != 0 {
						return
					} else {
						F_errmsg(m, int32(13891), int32(0))
						mBase = m.M
						v199 = m.ExcPending
						if v199 != 0 {
							return
						} else {
							F_errfinish(m, int32(496488), int32(745), int32(135050))
							mBase = m.M
							v204 = m.ExcPending
							if v204 != 0 {
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
				if v108 == int32(0) {
					F_errstart_cold(m, int32(22), int32(0))
					mBase = m.M
					v192 = m.ExcPending
					if v192 != 0 {
						return
					} else {
						F_errcode(m, int32(8389))
						mBase = m.M
						v195 = m.ExcPending
						if v195 != 0 {
							return
						} else {
							F_errmsg(m, int32(13891), int32(0))
							mBase = m.M
							v199 = m.ExcPending
							if v199 != 0 {
								return
							} else {
								F_errfinish(m, int32(496488), int32(745), int32(135050))
								mBase = m.M
								v204 = m.ExcPending
								if v204 != 0 {
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
					v127 = int32(0)
					*(*int32)(unsafe.Add(mBase, _consts[934])) = v127
					if v127 < v16 {
						v133 = v127
						for {
							v139 = *(*int32)(unsafe.Add(mBase, _consts[3]))
							v142 = v139 + v133<<(uint(int32(6))%32)
							*(*int32)(unsafe.Add(mBase, uint32(v142)+20)) = int32(-2) - v133
							*(*int32)(unsafe.Add(mBase, uint32(v142+int32(36)))) = int32(-1)
							v151 = v133 + int32(1)
							if v151 != v16 {
								v133 = v151
								continue
							} else {
								break
							}
							break
						}
					} else {
					}
					*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = int64(103079215124)
					v163 = F_hash_create(m, int32(396140), v16, v8, int32(40))
					mBase = m.M
					v164 = m.ExcPending
					if v164 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[935])) = v163
						if v163 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v208 = m.ExcPending
							if v208 != 0 {
								return
							} else {
								F_errmsg_internal(m, int32(392057), int32(0))
								mBase = m.M
								v212 = m.ExcPending
								if v212 != 0 {
									return
								} else {
									F_errfinish(m, int32(496488), int32(782), int32(135050))
									mBase = m.M
									v217 = m.ExcPending
									if v217 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, _consts[936])) = v16
							m.G0 = v8 + int32(48)
							return
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v176 = m.ExcPending
		if v176 != 0 {
			return
		} else {
			F_errcode(m, int32(322))
			mBase = m.M
			v179 = m.ExcPending
			if v179 != 0 {
				return
			} else {
				F_errmsg(m, int32(259701), int32(0))
				mBase = m.M
				v183 = m.ExcPending
				if v183 != 0 {
					return
				} else {
					F_errfinish(m, int32(496488), int32(736), int32(135050))
					mBase = m.M
					v188 = m.ExcPending
					if v188 != 0 {
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
