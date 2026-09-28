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
								F_errfinish(m, int32(_a_F_dsm_backend_startup_3), int32(454), int32(_a_F_dsm_backend_startup_4))
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
									F_errfinish(m, int32(_a_F_dsm_backend_startup_3), int32(454), int32(_a_F_dsm_backend_startup_4))
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
										F_errfinish(m, int32(_a_F_dsm_backend_startup_3), int32(454), int32(_a_F_dsm_backend_startup_4))
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
											F_errfinish(m, int32(_a_F_dsm_backend_startup_3), int32(454), int32(_a_F_dsm_backend_startup_4))
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
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
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
	F_ResourceOwnerRemember(m, v58, base.I64_extend_i32_u(v36), int32(_a_F_dsm_create_1))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L4
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v65 = v36 + int32(28)
	v67 = v36 + int32(24)
	v69 = v36 + int32(20)
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
	v154 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_create[6]))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v154)+4))
	if v155 != 0 {
		goto L37
	} else {
		goto L38
	}
L19:
	;
	v73 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_create[7]))
	v77 = F_LWLockAcquire(m, v73+int32(_a_F_dsm_create_2), int32(0))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L4
	} else {
		goto L22
	}
L20:
	;
	v107 = v3
	goto L21
L21:
	;
	goto L28
L22:
	;
	v79 = int32(12)
	v85 = int32(base.Ui32(l0)>>(uint(v79)%32)) + base.B2i32(l0&int32(4095) != int32(0))
	v88 = F_FreePageManagerGet(m, v20, v85, v15+v79)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	if v88 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v91 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_create[0]))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v93 = int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v67))) = v91 + v92<<(uint(v93)%32)
	*(*int32)(unsafe.Add(mBase, uint32(v65))) = v85 << (uint(v93) % 32)
	v146 = v85
	v150 = int32(1)
	goto L18
L25:
	;
	goto L26
L26:
	;
	v102 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_create[7]))
	F_LWLockRelease(m, v102+int32(_a_F_dsm_create_2))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	v107 = v85
	goto L21
L28:
	;
	v121 = Fn14349(m, int64(32))
	mBase = m.M
	goto L30
L29:
	;
	v134 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_create[7]))
	v138 = F_LWLockAcquire(m, v134+int32(_a_F_dsm_create_2), int32(0))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L4
	} else {
		goto L34
	}
L30:
	;
	v123 = v121 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+12)) = v123
	if v123 == int32(0) {
		goto L28
	} else {
		goto L31
	}
L31:
	;
	v129 = F_dsm_impl_op(m, int32(0), v123, l0, v69, v67, v65, int32(21))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L4
	} else {
		goto L32
	}
L32:
	;
	if v129 == int32(0) {
		goto L28
	} else {
		goto L33
	}
L33:
	;
	goto L29
L34:
	;
	v146 = v107
	v150 = v3
	goto L18
L35:
	;
	m.G0 = v15 + int32(16)
	return v339
L36:
	;
	v332 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_create[7]))
	F_LWLockRelease(m, v332+int32(_a_F_dsm_create_2))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L4
	} else {
		goto L75
	}
L37:
	;
	v159 = int32(0)
	goto L40
L38:
	;
	goto L39
L39:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v154)+8))
	if base.Ui32(v225) <= base.Ui32(v155) {
		goto L50
	} else {
		goto L51
	}
L40:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v154+v159*int32(24))+16))
	if v171 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	goto L39
L42:
	;
	if v150 != 0 {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	goto L44
L44:
	;
	v211 = v159 + int32(1)
	if v211 != v155 {
		v159 = v211
		goto L40
	} else {
		goto L49
	}
L45:
	;
	v177 = Fn14349(m, int64(32))
	mBase = m.M
	goto L48
L46:
	;
	v195 = v154
	goto L47
L47:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
	v201 = v195 + v159*int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v201)+16)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v201)+12)) = v198
	v205 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v201)+28)) = v205
	*(*uint8)(unsafe.Add(mBase, uint32(v201)+32)) = uint8(v205)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+16)) = v159
	goto L36
L48:
	;
	v180 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_create[6]))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+12)) = v159<<(uint(int32(1))%32) | v177<<(uint(int32(32)-base.I32_clz(v181))%32) | int32(1)
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v192 = v180 + v159*int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v192)+24)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v192)+20)) = v189
	v195 = v180
	goto L47
L49:
	;
	goto L41
L50:
	;
	if v150 != 0 {
		goto L54
	} else {
		goto L55
	}
L51:
	;
	goto L52
L52:
	;
	if v150 != 0 {
		goto L71
	} else {
		goto L72
	}
L53:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
	if v248 != 0 {
		goto L61
	} else {
		goto L62
	}
L54:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	F_FreePageManagerPut(m, v20, v227, v146)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L4
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v237 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_create[7]))
	F_LWLockRelease(m, v237+int32(_a_F_dsm_create_2))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L4
	} else {
		goto L59
	}
L57:
	;
	v231 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_create[7]))
	F_LWLockRelease(m, v231+int32(_a_F_dsm_create_2))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L4
	} else {
		goto L58
	}
L58:
	;
	goto L53
L59:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
	v246 = F_dsm_impl_op(m, int32(3), v243, int32(0), v69, v67, v65, int32(19))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L4
	} else {
		goto L60
	}
L60:
	;
	goto L53
L61:
	;
	F_ResourceOwnerForget(m, v248, base.I64_extend_i32_u(v36), int32(_a_F_dsm_create_1))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L4
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v253)+4)) = v254
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	*(*int32)(unsafe.Add(mBase, uint32(v254))) = v256
	F_pfree(m, v36)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
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
		v339 = int32(0)
		goto L35
	} else {
		goto L66
	}
L66:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L4
	} else {
		goto L67
	}
L67:
	;
	F_errcode(m, int32(197))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L4
	} else {
		goto L68
	}
L68:
	;
	F_errmsg(m, int32(_a_F_dsm_create_3), int32(0))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L4
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(_a_F_dsm_create_4), int32(634), int32(_a_F_dsm_create_5))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
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
	v282 = Fn14349(m, int64(32))
	mBase = m.M
	goto L74
L72:
	;
	v300 = v154
	goto L73
L73:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
	v306 = v300 + v155*int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v306)+16)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v306)+12)) = v303
	v310 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v306)+28)) = v310
	*(*uint8)(unsafe.Add(mBase, uint32(v306)+32)) = uint8(v310)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+16)) = v155
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v300)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v300)+4)) = v315 + int32(1)
	goto L36
L74:
	;
	v285 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_create[6]))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v285)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+12)) = v155<<(uint(int32(1))%32) | v282<<(uint(int32(32)-base.I32_clz(v286))%32) | int32(1)
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v297 = v285 + v155*int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v297)+24)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v297)+20)) = v294
	v300 = v285
	goto L73
L75:
	;
	v339 = v36
	goto L35
}
func F_dsm_find_mapping(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_find_mapping[0]))
	if base.B2i32(v4 == v2)|base.B2i32(v4 == int32(_a_F_dsm_find_mapping_0)) == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v13 = v4
	goto L4
L2:
	;
	goto L3
L3:
	;
	return int32(0)
L4:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if l0 == v14 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	return v13
L7:
	;
	goto L8
L8:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v17 != int32(_a_F_dsm_find_mapping_0) {
		v13 = v17
		goto L4
	} else {
		goto L9
	}
L9:
	;
	goto L5
}
func F_init_dsm_registry(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
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
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_init_dsm_registry[0]))
	if v3 == int32(0) {
		v7 = *(*int32)(unsafe.Add(mBase, _c_F_init_dsm_registry[1]))
		v11 = F_LWLockAcquire(m, v7+int32(_a_F_init_dsm_registry_0), int32(0))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, _c_F_init_dsm_registry[2]))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
			if v15 == int32(0) {
				v22 = F_dsa_create_ext(m, int32(88), int32(_a_F_init_dsm_registry_1), int32(134217728))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_init_dsm_registry[3])) = v22
					v28 = F_dshash_create(m, v22, int32(_a_F_init_dsm_registry_2), int32(0))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_init_dsm_registry[0])) = v28
						v32 = *(*int32)(unsafe.Add(mBase, _c_F_init_dsm_registry[3]))
						F_dsa_pin(m, v32)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return
						} else {
							v36 = *(*int32)(unsafe.Add(mBase, _c_F_init_dsm_registry[3]))
							F_dsa_pin_mapping(m, v36)
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return
							} else {
								v40 = *(*int32)(unsafe.Add(mBase, _c_F_init_dsm_registry[3]))
								v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
								v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+28))
								v44 = *(*int32)(unsafe.Add(mBase, _c_F_init_dsm_registry[2]))
								*(*int32)(unsafe.Add(mBase, uint32(v44))) = v42
								v47 = *(*int32)(unsafe.Add(mBase, _c_F_init_dsm_registry[0]))
								v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+32))
								v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
								v51 = *(*int32)(unsafe.Add(mBase, _c_F_init_dsm_registry[2]))
								*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = v49
								v73 = *(*int32)(unsafe.Add(mBase, _c_F_init_dsm_registry[1]))
								F_LWLockRelease(m, v73+int32(_a_F_init_dsm_registry_0))
								mBase = m.M
								v77 = m.ExcPending
								if v77 != 0 {
									return
								} else {
									return
								}
							}
						}
					}
				}
			} else {
				v54 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
				v55 = F_dsa_attach(m, v54)
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_init_dsm_registry[3])) = v55
					F_dsa_pin_mapping(m, v55)
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return
					} else {
						v62 = *(*int32)(unsafe.Add(mBase, _c_F_init_dsm_registry[3]))
						v65 = *(*int32)(unsafe.Add(mBase, _c_F_init_dsm_registry[2]))
						v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
						v68 = F_dshash_attach(m, v62, int32(_a_F_init_dsm_registry_2), v66, int32(0))
						mBase = m.M
						v69 = m.ExcPending
						if v69 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_init_dsm_registry[0])) = v68
							v73 = *(*int32)(unsafe.Add(mBase, _c_F_init_dsm_registry[1]))
							F_LWLockRelease(m, v73+int32(_a_F_init_dsm_registry_0))
							mBase = m.M
							v77 = m.ExcPending
							if v77 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			}
		}
	} else {
		return
	}
}
