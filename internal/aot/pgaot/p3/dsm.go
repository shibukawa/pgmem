package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_dsm_backend_startup(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_dsm_backend_startup[0])))
	if v9 == int32(1) {
		v12 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v12
		v16 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_backend_startup[1]))
		v23 = F_dsm_impl_op(m, int32(1), v16, v12, int32(_a_F_dsm_backend_startup_0), v6+int32(12), int32(_a_F_dsm_backend_startup_1), int32(21))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
			*(*int32)(unsafe.Add(mBase, _c_F_dsm_backend_startup[2])) = v26
			v29 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_backend_startup[3]))
			if base.Ui32(v29) < base.Ui32(int32(12)) {
				v56 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_backend_startup[1]))
				v63 = F_dsm_impl_op(m, int32(2), v56, int32(0), int32(_a_F_dsm_backend_startup_0), v6+int32(12), int32(_a_F_dsm_backend_startup_1), int32(19))
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
					return
				} else {
					F_errstart_cold(m, int32(22), int32(0))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return
					} else {
						F_errcode(m, int32(2600))
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return
						} else {
							F_errmsg(m, int32(_a_F_dsm_backend_startup_2), int32(0))
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_dsm_backend_startup_3), int32(444), int32(_a_F_dsm_backend_startup_4))
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
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
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
				if v32 != int32(-1706017486) {
					v56 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_backend_startup[1]))
					v63 = F_dsm_impl_op(m, int32(2), v56, int32(0), int32(_a_F_dsm_backend_startup_0), v6+int32(12), int32(_a_F_dsm_backend_startup_1), int32(19))
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return
					} else {
						F_errstart_cold(m, int32(22), int32(0))
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return
						} else {
							F_errcode(m, int32(2600))
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return
							} else {
								F_errmsg(m, int32(_a_F_dsm_backend_startup_2), int32(0))
								mBase = m.M
								v75 = m.ExcPending
								if v75 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_dsm_backend_startup_3), int32(444), int32(_a_F_dsm_backend_startup_4))
									mBase = m.M
									v80 = m.ExcPending
									if v80 != 0 {
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
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
					if base.Ui64(base.I64_extend_i32_u(v29)) < base.Ui64(base.I64_extend_i32_u(v36)*int64(24)+int64(12)) {
						v56 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_backend_startup[1]))
						v63 = F_dsm_impl_op(m, int32(2), v56, int32(0), int32(_a_F_dsm_backend_startup_0), v6+int32(12), int32(_a_F_dsm_backend_startup_1), int32(19))
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return
						} else {
							F_errstart_cold(m, int32(22), int32(0))
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return
							} else {
								F_errcode(m, int32(2600))
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return
								} else {
									F_errmsg(m, int32(_a_F_dsm_backend_startup_2), int32(0))
									mBase = m.M
									v75 = m.ExcPending
									if v75 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_dsm_backend_startup_3), int32(444), int32(_a_F_dsm_backend_startup_4))
										mBase = m.M
										v80 = m.ExcPending
										if v80 != 0 {
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
						v43 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
						if base.Ui32(v36) < base.Ui32(v43) {
							v56 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_backend_startup[1]))
							v63 = F_dsm_impl_op(m, int32(2), v56, int32(0), int32(_a_F_dsm_backend_startup_0), v6+int32(12), int32(_a_F_dsm_backend_startup_1), int32(19))
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return
							} else {
								F_errstart_cold(m, int32(22), int32(0))
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return
								} else {
									F_errcode(m, int32(2600))
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
										return
									} else {
										F_errmsg(m, int32(_a_F_dsm_backend_startup_2), int32(0))
										mBase = m.M
										v75 = m.ExcPending
										if v75 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_dsm_backend_startup_3), int32(444), int32(_a_F_dsm_backend_startup_4))
											mBase = m.M
											v80 = m.ExcPending
											if v80 != 0 {
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
							v48 = int32(1)
							*(*uint8)(unsafe.Add(mBase, _c_F_dsm_backend_startup[4])) = uint8(v48)
							m.G0 = v6 + int32(16)
							return
						}
					}
				}
			}
		}
	} else {
		v48 = int32(1)
		*(*uint8)(unsafe.Add(mBase, _c_F_dsm_backend_startup[4])) = uint8(v48)
		m.G0 = v6 + int32(16)
		return
	}
}
func F_dsm_create(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
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
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
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
	var v284 int32
	_ = v284
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	v3 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v3
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_create[0]))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_dsm_create[1])))
	if v22 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_dsm_backend_startup(m)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_create[2]))
	if v30 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L3
L6:
	;
	F_ResourceOwnerEnlarge(m, v30)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L4
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_create[3]))
	v36 = F_MemoryContextAlloc(m, v34, int32(36))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L4
	} else {
		goto L10
	}
L9:
	;
	goto L8
L10:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_create[4]))
	if v39 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v42 = int32(_a_F_dsm_create_0)
	*(*int32)(unsafe.Add(mBase, _c_F_dsm_create[5])) = v42
	v46 = v42
	goto L13
L12:
	;
	v46 = v39
	goto L13
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = int32(_a_F_dsm_create_0)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+4)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v46))) = v36
	*(*int32)(unsafe.Add(mBase, _c_F_dsm_create[4])) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v36)+24)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v36)+16)) = int64(4294967295)
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_create[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+8)) = v58
	if v58 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	F_ResourceOwnerRemember(m, v58, v36, int32(_a_F_dsm_create_1))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L4
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v64 = v36 + int32(28)
	v66 = v36 + int32(24)
	v68 = v36 + int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+32)) = int32(0)
	if v20 != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L16
L18:
	;
	v153 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_create[6]))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v153)+4))
	if v154 != 0 {
		goto L37
	} else {
		goto L38
	}
L19:
	;
	v72 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_create[7]))
	v76 = F_LWLockAcquire(m, v72+int32(_a_F_dsm_create_2), int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L4
	} else {
		goto L22
	}
L20:
	;
	v106 = v3
	goto L21
L21:
	;
	goto L28
L22:
	;
	v78 = int32(12)
	v84 = int32(base.Ui32(l0)>>(uint(v78)%32)) + base.B2i32(l0&int32(4095) != int32(0))
	v87 = F_FreePageManagerGet(m, v20, v84, v15+v78)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	if v87 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_create[0]))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v92 = int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v66))) = v90 + v91<<(uint(v92)%32)
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v84 << (uint(v92) % 32)
	v145 = v84
	v149 = int32(1)
	goto L18
L25:
	;
	goto L26
L26:
	;
	v101 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_create[7]))
	F_LWLockRelease(m, v101+int32(_a_F_dsm_create_2))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	v106 = v84
	goto L21
L28:
	;
	v120 = Fn13986(m, int64(32))
	mBase = m.M
	goto L30
L29:
	;
	v133 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_create[7]))
	v137 = F_LWLockAcquire(m, v133+int32(_a_F_dsm_create_2), int32(0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L4
	} else {
		goto L34
	}
L30:
	;
	v122 = v120 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+12)) = v122
	if v122 == int32(0) {
		goto L28
	} else {
		goto L31
	}
L31:
	;
	v128 = F_dsm_impl_op(m, int32(0), v122, l0, v68, v66, v64, int32(21))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L4
	} else {
		goto L32
	}
L32:
	;
	if v128 == int32(0) {
		goto L28
	} else {
		goto L33
	}
L33:
	;
	goto L29
L34:
	;
	v145 = v106
	v149 = v3
	goto L18
L35:
	;
	m.G0 = v15 + int32(16)
	return v337
L36:
	;
	v330 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_create[7]))
	F_LWLockRelease(m, v330+int32(_a_F_dsm_create_2))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L4
	} else {
		goto L75
	}
L37:
	;
	v158 = int32(0)
	goto L40
L38:
	;
	goto L39
L39:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v153)+8))
	if base.Ui32(v224) <= base.Ui32(v154) {
		goto L50
	} else {
		goto L51
	}
L40:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v153+v158*int32(24))+16))
	if v170 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	goto L39
L42:
	;
	if v149 != 0 {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	goto L44
L44:
	;
	v210 = v158 + int32(1)
	if v210 != v154 {
		v158 = v210
		goto L40
	} else {
		goto L49
	}
L45:
	;
	v176 = Fn13986(m, int64(32))
	mBase = m.M
	goto L48
L46:
	;
	v194 = v153
	goto L47
L47:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
	v200 = v194 + v158*int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v200)+16)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v200)+12)) = v197
	v204 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v200)+28)) = v204
	*(*uint8)(unsafe.Add(mBase, uint32(v200)+32)) = uint8(v204)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+16)) = v158
	goto L36
L48:
	;
	v179 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_create[6]))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v179)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+12)) = v158<<(uint(int32(1))%32) | v176<<(uint(int32(32)-base.I32_clz(v180))%32) | int32(1)
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v191 = v179 + v158*int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v191)+24)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v191)+20)) = v188
	v194 = v179
	goto L47
L49:
	;
	goto L41
L50:
	;
	if v149 != 0 {
		goto L54
	} else {
		goto L55
	}
L51:
	;
	goto L52
L52:
	;
	if v149 != 0 {
		goto L71
	} else {
		goto L72
	}
L53:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
	if v247 != 0 {
		goto L61
	} else {
		goto L62
	}
L54:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	F_FreePageManagerPut(m, v20, v226, v145)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L4
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v236 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_create[7]))
	F_LWLockRelease(m, v236+int32(_a_F_dsm_create_2))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L4
	} else {
		goto L59
	}
L57:
	;
	v230 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_create[7]))
	F_LWLockRelease(m, v230+int32(_a_F_dsm_create_2))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L4
	} else {
		goto L58
	}
L58:
	;
	goto L53
L59:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
	v245 = F_dsm_impl_op(m, int32(3), v242, int32(0), v68, v66, v64, int32(19))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L4
	} else {
		goto L60
	}
L60:
	;
	goto L53
L61:
	;
	F_ResourceOwnerForget(m, v247, v36, int32(_a_F_dsm_create_1))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L4
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v251)+4)) = v252
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	*(*int32)(unsafe.Add(mBase, uint32(v252))) = v254
	F_pfree(m, v36)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L4
	} else {
		goto L65
	}
L64:
	;
	goto L63
L65:
	;
	if l1&int32(1) != 0 {
		v337 = int32(0)
		goto L35
	} else {
		goto L66
	}
L66:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L4
	} else {
		goto L67
	}
L67:
	;
	F_errcode(m, int32(197))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L4
	} else {
		goto L68
	}
L68:
	;
	F_errmsg(m, int32(_a_F_dsm_create_3), int32(0))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L4
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(_a_F_dsm_create_4), int32(626), int32(_a_F_dsm_create_5))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L4
	} else {
		goto L70
	}
L70:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L71:
	;
	v280 = Fn13986(m, int64(32))
	mBase = m.M
	goto L74
L72:
	;
	v298 = v153
	goto L73
L73:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
	v304 = v298 + v154*int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v304)+16)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v304)+12)) = v301
	v308 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v304)+28)) = v308
	*(*uint8)(unsafe.Add(mBase, uint32(v304)+32)) = uint8(v308)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+16)) = v154
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v298)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v298)+4)) = v313 + int32(1)
	goto L36
L74:
	;
	v283 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_create[6]))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v283)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+12)) = v154<<(uint(int32(1))%32) | v280<<(uint(int32(32)-base.I32_clz(v284))%32) | int32(1)
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v295 = v283 + v154*int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v295)+24)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v295)+20)) = v292
	v298 = v283
	goto L73
L75:
	;
	v337 = v36
	goto L35
}
