package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CommandCounterIncrement(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	v4 = int32(*(*uint8)(unsafe.Add(mBase, _consts[152])))
	if v4 != 0 {
		v6 = *(*int32)(unsafe.Add(mBase, _consts[37]))
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+72))
		if v7 != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v64 = m.ExcPending
			if v64 != 0 {
				return
			} else {
				F_errcode(m, int32(322))
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
					return
				} else {
					F_errmsg(m, int32(256278), int32(0))
					mBase = m.M
					v71 = m.ExcPending
					if v71 != 0 {
						return
					} else {
						F_errfinish(m, int32(484613), int32(1118), int32(94484))
						mBase = m.M
						v76 = m.ExcPending
						if v76 != 0 {
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
			v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+76)))
			if v8 != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
					return
				} else {
					F_errcode(m, int32(322))
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return
					} else {
						F_errmsg(m, int32(256278), int32(0))
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return
						} else {
							F_errfinish(m, int32(484613), int32(1118), int32(94484))
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
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
				v10 = *(*int32)(unsafe.Add(mBase, _consts[55]))
				if int32(0) <= v10 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return
					} else {
						F_errcode(m, int32(322))
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return
						} else {
							F_errmsg(m, int32(256278), int32(0))
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return
							} else {
								F_errfinish(m, int32(484613), int32(1118), int32(94484))
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
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
					v13 = int32(4365484)
					v15 = *(*int32)(unsafe.Add(mBase, _consts[112]))
					v17 = v15 + int32(1)
					*(*int32)(unsafe.Add(mBase, _consts[112])) = v17
					if v17 == int32(-1) {
						*(*int32)(unsafe.Add(mBase, _consts[112])) = v15
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v82 = m.ExcPending
						if v82 != 0 {
							return
						} else {
							F_errcode(m, int32(261))
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
								return
							} else {
								F_errmsg(m, int32(253504), int32(0))
								mBase = m.M
								v89 = m.ExcPending
								if v89 != 0 {
									return
								} else {
									F_errfinish(m, int32(484613), int32(1126), int32(94484))
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
						}
					} else {
						v22 = int32(0)
						*(*uint8)(unsafe.Add(mBase, _consts[152])) = uint8(v22)
						v25 = int32(*(*uint8)(unsafe.Add(mBase, _consts[153])))
						if v25 != int32(1) {
						} else {
							v29 = *(*int32)(unsafe.Add(mBase, _consts[154]))
							if v29 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v17
							} else {
							}
							v32 = *(*int32)(unsafe.Add(mBase, _consts[155]))
							if v32 == int32(0) {
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v32)+32)) = v17
							}
						}
						v38 = *(*int32)(unsafe.Add(mBase, _consts[116]))
						if v38 != 0 {
							F_merge_map_updates(m, int32(4458672), int32(4460768), int32(1))
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, _consts[116])) = int32(0)
								v48 = *(*int32)(unsafe.Add(mBase, _consts[117]))
								if v48 != 0 {
									F_merge_map_updates(m, int32(4459720), int32(4461292), int32(1))
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, _consts[117])) = int32(0)
										F_CommandEndInvalidationMessages(m)
										mBase = m.M
										v58 = m.ExcPending
										if v58 != 0 {
											return
										} else {
											return
										}
									}
								} else {
									F_CommandEndInvalidationMessages(m)
									mBase = m.M
									v58 = m.ExcPending
									if v58 != 0 {
										return
									} else {
										return
									}
								}
							}
						} else {
							v48 = *(*int32)(unsafe.Add(mBase, _consts[117]))
							if v48 != 0 {
								F_merge_map_updates(m, int32(4459720), int32(4461292), int32(1))
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, _consts[117])) = int32(0)
									F_CommandEndInvalidationMessages(m)
									mBase = m.M
									v58 = m.ExcPending
									if v58 != 0 {
										return
									} else {
										return
									}
								}
							} else {
								F_CommandEndInvalidationMessages(m)
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return
								} else {
									return
								}
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
func F_CommandEndInvalidationMessages(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	v5 = *(*int32)(unsafe.Add(mBase, _consts[876]))
	if v5 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
	if v6 < v7 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v9 = v6
	goto L7
L5:
	;
	goto L6
L6:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
	if v25 < v26 {
		goto L12
	} else {
		goto L13
	}
L7:
	;
	v13 = *(*int32)(unsafe.Add(mBase, _consts[877]))
	F_LocalExecuteInvalidationMessage(m, v13+v9<<(uint(int32(4))%32))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L6
L9:
	;
	return
L10:
	;
	v20 = v9 + int32(1)
	if v20 != v7 {
		v9 = v20
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	v28 = v25
	goto L15
L13:
	;
	goto L14
L14:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	if int32(2) <= v45 {
		goto L19
	} else {
		goto L20
	}
L15:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _consts[878]))
	F_LocalExecuteInvalidationMessage(m, v32+v28<<(uint(int32(4))%32))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L9
	} else {
		goto L17
	}
L16:
	;
	goto L14
L17:
	;
	v39 = v28 + int32(1)
	if v39 != v26 {
		v28 = v39
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	F_LogLogicalInvalidations(m)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L9
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _consts[876]))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v51)+28)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v51))) = v52
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v51)+32)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = v55
	goto L3
L22:
	;
	goto L21
}
func F_EndCommand(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int64
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v70 int64
	_ = v70
	var v72 int32
	_ = v72
	var v76 int64
	_ = v76
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int64
	_ = v89
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v154 int64
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	v7 = m.G0
	v9 = v7 + int32(-64)
	m.G0 = v9
	v11 = int32(2)
	if base.Ui32(l1-v11) <= base.Ui32(v11) {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v17 = v15 << (uint(int32(3)) % 32)
		v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+uint32(_consts[656]))))
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v17)+uint32(_consts[657])))
		if v20 != 0 {
			v24 = F__emscripten_memcpy_bulkmem(m, v9, v23, v20)
			mBase = m.M
			v25 = v24
		} else {
			v25 = v9
		}
		v26 = v20 + v25
		v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+uint32(_consts[658]))))
		if v29&int32(1) != 0 {
			if v15 == int32(158) {
				v34 = int32(12320)
				*(*uint16)(unsafe.Add(mBase, uint32(v26))) = uint16(v34)
				v38 = v26 + int32(2)
			} else {
				v38 = v26
			}
			v39 = int32(32)
			*(*uint8)(unsafe.Add(mBase, uint32(v38))) = uint8(v39)
			v41 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
			v43 = v38 + int32(1)
			v44 = int32(0)
			if v41 == int64(0) {
				v53 = int32(48)
				*(*uint8)(unsafe.Add(mBase, uint32(v43))) = uint8(v53)
				v238 = int32(1)
			} else {
				v60 = int32(1233)
				v65 = int32(base.Ui32((base.I32_wrap_i64(base.I64_clz(v41))^int32(63))*v60+v60) >> (uint(int32(12)) % 32))
				v70 = *(*int64)(unsafe.Add(mBase, uint32(v65<<(uint(int32(3))%32))+uint32(_consts[659])))
				v72 = v65 + base.B2i32(base.Ui64(v70) <= base.Ui64(v41))
				if base.Ui64(v41) < base.Ui64(int64(100000000)) {
					v150 = v44
					v154 = v41
				} else {
					v76 = v41
					v80 = v44
					for {
						v85 = v43 + v72 - v80
						v86 = int32(8)
						v89 = base.I64_div_u_s(v76, int64(100000000))
						v93 = base.I32_wrap_i64(v89*int64(4194967296) + v76)
						v95 = base.I32_div_u_s(v93, int32(1000000))
						v96 = int32(1)
						v100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v95<<(uint(v96)%32))+uint32(_consts[660]))))
						*(*uint16)(unsafe.Add(mBase, uint32(v85-v86))) = uint16(v100)
						v104 = int32(10000)
						v105 = base.I32_div_u_s(v93, v104)
						v106 = int32(100)
						v107 = base.I32_rem_u_s(v105, v106)
						v112 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v107<<(uint(v96)%32))+uint32(_consts[660]))))
						*(*uint16)(unsafe.Add(mBase, uint32(v85-int32(6)))) = uint16(v112)
						v118 = v93 - v105*v104
						v119 = int32(65535)
						v122 = base.I32_div_u_s(v118&v119, v106)
						v127 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v122<<(uint(v96)%32))+uint32(_consts[660]))))
						*(*uint16)(unsafe.Add(mBase, uint32(v85-int32(4)))) = uint16(v127)
						v140 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v118-v122*v106)&v119<<(uint(v96)%32))+uint32(_consts[660]))))
						*(*uint16)(unsafe.Add(mBase, uint32(v85-int32(2)))) = uint16(v140)
						v143 = v80 + v86
						if base.Ui64(int64(9999999999999999)) < base.Ui64(v76) {
							v76 = v89
							v80 = v143
							continue
						} else {
							break
						}
						break
					}
					v150 = v143
					v154 = v89
				}
				v155 = base.I32_wrap_i64(v154)
				if base.Ui64(v154) < base.Ui64(int64(10000)) {
					v189 = v155
					v190 = v150
				} else {
					v159 = v43 + v72 - v150
					v160 = int32(4)
					v163 = base.I32_div_u_s(v155, int32(10000))
					v166 = v163*int32(-10000) + v155
					v167 = int32(100)
					v168 = base.I32_div_u_s(v166, v167)
					v169 = int32(1)
					v173 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v168<<(uint(v169)%32))+uint32(_consts[660]))))
					*(*uint16)(unsafe.Add(mBase, uint32(v159-v160))) = uint16(v173)
					v184 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v166-v168*v167)<<(uint(v169)%32))+uint32(_consts[660]))))
					*(*uint16)(unsafe.Add(mBase, uint32(v159-int32(2)))) = uint16(v184)
					v189 = v163
					v190 = v150 | v160
				}
				if base.Ui32(v189) < base.Ui32(int32(100)) {
					v212 = v189
					v213 = v190
				} else {
					v197 = int32(2)
					v199 = int32(100)
					v200 = base.I32_div_u_s(v189, v199)
					v208 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v189-v200*v199)<<(uint(int32(1))%32))+uint32(_consts[660]))))
					*(*uint16)(unsafe.Add(mBase, uint32(v43+v72-v190-v197))) = uint16(v208)
					v212 = v200
					v213 = v190 + v197
				}
				if base.Ui32(int32(10)) <= base.Ui32(v212) {
					v224 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v212<<(uint(int32(1))%32))+uint32(_consts[660]))))
					*(*uint16)(unsafe.Add(mBase, uint32(v43+v72-v213-int32(2)))) = uint16(v224)
					v238 = v72
				} else {
					v227 = v212 | int32(48)
					*(*uint8)(unsafe.Add(mBase, uint32(v43))) = uint8(v227)
					v238 = v72
				}
			}
			v240 = v238 + v43
		} else {
			v240 = v26
		}
		v241 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v240))) = uint8(v241)
		v248 = *(*int32)(unsafe.Add(mBase, _consts[136]))
		v249 = *(*int32)(unsafe.Add(mBase, uint32(v248)+16))
		v250 = m.T0[v249].(func(*base.Module, int32, int32, int32) int32)(m, int32(67), v9, v240-v25+int32(1))
		mBase = m.M
		v251 = m.ExcPending
		if v251 != 0 {
			return
		} else {
			m.G0 = v9 - int32(-64)
			return
		}
	} else {
		m.G0 = v9 - int32(-64)
		return
	}
}
