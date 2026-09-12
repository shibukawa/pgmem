package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_IncrTupleDescRefCount(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	v3 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	F_ResourceOwnerEnlarge(m, v3)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v6 + int32(1)
		v11 = *(*int32)(unsafe.Add(mBase, _consts[11]))
		F_ResourceOwnerRemember(m, v11, l0, int32(795152))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			return
		}
	}
}
func F_InitProcessGlobals(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int64
	_ = v14
	var v15 int64
	_ = v15
	var v23 int64
	_ = v23
	var v27 int64
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v84 int64
	_ = v84
	var v87 int64
	_ = v87
	var v95 int32
	_ = v95
	var v97 int64
	_ = v97
	var v99 int64
	_ = v99
	var v105 int64
	_ = v105
	var v109 int64
	_ = v109
	var v110 int64
	_ = v110
	var v113 int64
	_ = v113
	var v114 int64
	_ = v114
	var v115 int64
	_ = v115
	var v118 int64
	_ = v118
	var v119 int64
	_ = v119
	var v120 int64
	_ = v120
	var v125 int64
	_ = v125
	var v130 int64
	_ = v130
	var v135 int64
	_ = v135
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int64
	_ = v152
	var v154 int64
	_ = v154
	var v155 int64
	_ = v155
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v203 int64
	_ = v203
	var v204 int32
	_ = v204
	var v213 int64
	_ = v213
	var v215 int64
	_ = v215
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	v9 = m.G0
	v10 = int32(16)
	v11 = v9 - v10
	m.G0 = v11
	F___gettimeofday(m, v11)
	mBase = m.M
	v14 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	v15 = int64(*(*int32)(unsafe.Add(mBase, uint32(v11)+8)))
	m.G0 = v11 + v10
	v23 = v15 + v14*int64(1000000) - int64(946684800000000)
	goto L1
L1:
	;
	*(*int64)(unsafe.Add(mBase, _consts[446])) = v23
	v27 = base.I64_div_s(v23, int64(1000000))
	goto L2
L2:
	;
	*(*int64)(unsafe.Add(mBase, _consts[447])) = v27 + int64(946684800)
	v32 = int32(16)
	v33 = int32(0)
	v37 = m.G0
	v39 = v37 - v32
	m.G0 = v39
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = v33
	v45 = F_open(m, int32(303280), v33, v39)
	mBase = m.M
	if v45 != int32(-1) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v150 = int32(4645608)
	v151 = int32(4645600)
	v152 = *(*int64)(unsafe.Add(mBase, _consts[58]))
	v154 = *(*int64)(unsafe.Add(mBase, _consts[59]))
	v155 = v152 ^ v154
	*(*int64)(unsafe.Add(mBase, _consts[59])) = base.I64_rotl(v155, int64(37))
	*(*int64)(unsafe.Add(mBase, _consts[58])) = v155<<(uint(int64(16))%64) ^ base.I64_rotl(v152, int64(24)) ^ v155
	v175 = base.I32_wrap_i64(int64(base.Ui64(base.I64_rotl(v152*int64(5), int64(7))*int64(9)) >> (uint(int64(32)) % 64)))
	goto L29
L4:
	;
	if v78 != 0 {
		goto L17
	} else {
		goto L18
	}
L5:
	;
	goto L9
L6:
	;
	v78 = v33
	goto L7
L7:
	;
	m.G0 = v39 + int32(16)
	goto L4
L8:
	;
	v73 = F_close(m, v45)
	mBase = m.M
	v78 = v71
	goto L7
L9:
	;
	v51 = int32(4645600)
	v52 = v32
	goto L10
L10:
	;
	v57 = F_read(m, v45, v51, v52)
	mBase = m.M
	if v57 <= int32(0) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v71 = int32(1)
	goto L8
L12:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _consts[85]))
	if v61 == int32(27) {
		goto L10
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v66 = v52 - v57
	if v66 != 0 {
		v51 = v51 + v57
		v52 = v66
		goto L10
	} else {
		goto L16
	}
L15:
	;
	v71 = int32(0)
	goto L8
L16:
	;
	goto L11
L17:
	;
	v83 = int32(4645600)
	v84 = *(*int64)(unsafe.Add(mBase, _consts[58]))
	if v84 != int64(0) {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	goto L19
L19:
	;
	v95 = int32(4645600)
	v97 = int64(*(*int32)(unsafe.Add(mBase, _consts[448])))
	v99 = *(*int64)(unsafe.Add(mBase, _consts[446]))
	v105 = v97 ^ v99<<(uint(int64(12))%64) ^ int64(base.Ui64(v99)>>(uint(int64(20))%64))
	v109 = v105 + int64(4354685564936845354)
	v110 = int64(30)
	v113 = int64(-4658895280553007687)
	v114 = (int64(base.Ui64(v109)>>(uint(v110)%64)) ^ v109) * v113
	v115 = int64(27)
	v118 = int64(-7723592293110705685)
	v119 = (int64(base.Ui64(v114)>>(uint(v115)%64)) ^ v114) * v118
	v120 = int64(31)
	*(*int64)(unsafe.Add(mBase, _consts[59])) = int64(base.Ui64(v119)>>(uint(v120)%64)) ^ v119
	v125 = v105 - int64(7046029254386353131)
	v130 = (int64(base.Ui64(v125)>>(uint(v110)%64)) ^ v125) * v113
	v135 = (int64(base.Ui64(v130)>>(uint(v115)%64)) ^ v130) * v118
	*(*int64)(unsafe.Add(mBase, _consts[58])) = int64(base.Ui64(v135)>>(uint(v120)%64)) ^ v135
	if v125|v109 == int64(0) {
		goto L26
	} else {
		goto L27
	}
L20:
	;
	goto L3
L21:
	;
	goto L20
L22:
	;
	v87 = *(*int64)(unsafe.Add(mBase, _consts[59]))
	if v87 != int64(0) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	*(*int64)(unsafe.Add(mBase, _consts[59])) = int64(1442695040888963407)
	*(*int64)(unsafe.Add(mBase, _consts[58])) = int64(6364136223846793005)
	goto L21
L25:
	;
	goto L3
L26:
	;
	*(*int64)(unsafe.Add(mBase, _consts[59])) = int64(1442695040888963407)
	*(*int64)(unsafe.Add(mBase, _consts[58])) = int64(6364136223846793005)
	goto L28
L27:
	;
	goto L28
L28:
	;
	goto L25
L29:
	;
	v177 = *(*int32)(unsafe.Add(mBase, _consts[449]))
	if v177 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	return
L31:
	;
	v181 = *(*int32)(unsafe.Add(mBase, _consts[450]))
	*(*int32)(unsafe.Add(mBase, uint32(v181))) = v175
	goto L30
L32:
	;
	goto L33
L33:
	;
	v184 = int32(3)
	if v177 == int32(7) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v189 = v184
	goto L36
L35:
	;
	v189 = int32(1)
	goto L36
L36:
	;
	if v177 == int32(31) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v192 = v184
	goto L39
L38:
	;
	v192 = v189
	goto L39
L39:
	;
	*(*int32)(unsafe.Add(mBase, _consts[451])) = v192
	v195 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[452])) = v195
	v198 = *(*int32)(unsafe.Add(mBase, _consts[450]))
	if v195 < v177 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v203 = base.I64_extend_i32_u(v175)
	v204 = int32(0)
	goto L43
L41:
	;
	goto L42
L42:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
	*(*int32)(unsafe.Add(mBase, uint32(v198))) = v224 | int32(1)
	goto L30
L43:
	;
	v213 = v203*int64(6364136223846793005) + int64(1)
	v215 = int64(base.Ui64(v213) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v198+v204<<(uint(int32(2))%32)))) = uint32(v215)
	v218 = v204 + int32(1)
	if v218 != v177 {
		v203 = v213
		v204 = v218
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
}
func F_InputFunctionCallSafe(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	v9 = m.G0
	v11 = v9 + int32(-64)
	m.G0 = v11
	if l1 != 0 {
		v20 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v11)+36)) = uint8(v20)
		*(*int64)(unsafe.Add(mBase, uint32(v11)+28)) = int64(0)
		*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = l4
		*(*uint8)(unsafe.Add(mBase, uint32(v11)+60)) = uint8(v20)
		*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = l3
		*(*uint8)(unsafe.Add(mBase, uint32(v11)+52)) = uint8(v20)
		*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = l2
		*(*uint8)(unsafe.Add(mBase, uint32(v11)+44)) = uint8(v20)
		*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = l1
		v34 = int32(3)
		*(*uint16)(unsafe.Add(mBase, uint32(v11)+38)) = uint16(v34)
		*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = l0
		v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v40 = m.T0[v39].(func(*base.Module, int32) int32)(m, v9+int32(-44))
		mBase = m.M
		v43 = m.ExcPending
		if v43 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l5))) = v40
			if l4 == int32(0) {
				v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+36)))
				if l1 == int32(0) {
					v56 = int32(1)
					if v53&v56 != 0 {
						v95 = v56
						m.G0 = v11 - int32(-64)
						return v95
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int32(0)
						} else {
							v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v11))) = v63
							F_errmsg_internal(m, int32(559793), v11)
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(520932), int32(1618), int32(431009))
								mBase = m.M
								v72 = m.ExcPending
								if v72 != 0 {
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
					v73 = int32(1)
					if v53&v73 == int32(0) {
						v95 = v73
						m.G0 = v11 - int32(-64)
						return v95
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v81 = m.ExcPending
						if v81 != 0 {
							return int32(0)
						} else {
							v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v82
							F_errmsg_internal(m, int32(559945), v9+int32(-48))
							mBase = m.M
							v88 = m.ExcPending
							if v88 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(520932), int32(1624), int32(431009))
								mBase = m.M
								v93 = m.ExcPending
								if v93 != 0 {
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
			} else {
				v47 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
				if v47 != int32(447) {
					v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+36)))
					if l1 == int32(0) {
						v56 = int32(1)
						if v53&v56 != 0 {
							v95 = v56
							m.G0 = v11 - int32(-64)
							return v95
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return int32(0)
							} else {
								v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v11))) = v63
								F_errmsg_internal(m, int32(559793), v11)
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(520932), int32(1618), int32(431009))
									mBase = m.M
									v72 = m.ExcPending
									if v72 != 0 {
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
						v73 = int32(1)
						if v53&v73 == int32(0) {
							v95 = v73
							m.G0 = v11 - int32(-64)
							return v95
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return int32(0)
							} else {
								v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v82
								F_errmsg_internal(m, int32(559945), v9+int32(-48))
								mBase = m.M
								v88 = m.ExcPending
								if v88 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(520932), int32(1624), int32(431009))
									mBase = m.M
									v93 = m.ExcPending
									if v93 != 0 {
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
				} else {
					v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+4)))
					if v51 != 0 {
						v95 = int32(0)
						m.G0 = v11 - int32(-64)
						return v95
					} else {
						v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+36)))
						if l1 == int32(0) {
							v56 = int32(1)
							if v53&v56 != 0 {
								v95 = v56
								m.G0 = v11 - int32(-64)
								return v95
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return int32(0)
								} else {
									v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v11))) = v63
									F_errmsg_internal(m, int32(559793), v11)
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(520932), int32(1618), int32(431009))
										mBase = m.M
										v72 = m.ExcPending
										if v72 != 0 {
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
							v73 = int32(1)
							if v53&v73 == int32(0) {
								v95 = v73
								m.G0 = v11 - int32(-64)
								return v95
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return int32(0)
								} else {
									v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v82
									F_errmsg_internal(m, int32(559945), v9+int32(-48))
									mBase = m.M
									v88 = m.ExcPending
									if v88 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(520932), int32(1624), int32(431009))
										mBase = m.M
										v93 = m.ExcPending
										if v93 != 0 {
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
			}
		}
	} else {
		v13 = int32(1)
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+10)))
		if v14 != v13 {
			v20 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v11)+36)) = uint8(v20)
			*(*int64)(unsafe.Add(mBase, uint32(v11)+28)) = int64(0)
			*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = l4
			*(*uint8)(unsafe.Add(mBase, uint32(v11)+60)) = uint8(v20)
			*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = l3
			*(*uint8)(unsafe.Add(mBase, uint32(v11)+52)) = uint8(v20)
			*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = l2
			*(*uint8)(unsafe.Add(mBase, uint32(v11)+44)) = uint8(v20)
			*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = l1
			v34 = int32(3)
			*(*uint16)(unsafe.Add(mBase, uint32(v11)+38)) = uint16(v34)
			*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = l0
			v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v40 = m.T0[v39].(func(*base.Module, int32) int32)(m, v9+int32(-44))
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l5))) = v40
				if l4 == int32(0) {
					v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+36)))
					if l1 == int32(0) {
						v56 = int32(1)
						if v53&v56 != 0 {
							v95 = v56
							m.G0 = v11 - int32(-64)
							return v95
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return int32(0)
							} else {
								v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v11))) = v63
								F_errmsg_internal(m, int32(559793), v11)
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(520932), int32(1618), int32(431009))
									mBase = m.M
									v72 = m.ExcPending
									if v72 != 0 {
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
						v73 = int32(1)
						if v53&v73 == int32(0) {
							v95 = v73
							m.G0 = v11 - int32(-64)
							return v95
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return int32(0)
							} else {
								v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v82
								F_errmsg_internal(m, int32(559945), v9+int32(-48))
								mBase = m.M
								v88 = m.ExcPending
								if v88 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(520932), int32(1624), int32(431009))
									mBase = m.M
									v93 = m.ExcPending
									if v93 != 0 {
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
				} else {
					v47 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
					if v47 != int32(447) {
						v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+36)))
						if l1 == int32(0) {
							v56 = int32(1)
							if v53&v56 != 0 {
								v95 = v56
								m.G0 = v11 - int32(-64)
								return v95
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return int32(0)
								} else {
									v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v11))) = v63
									F_errmsg_internal(m, int32(559793), v11)
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(520932), int32(1618), int32(431009))
										mBase = m.M
										v72 = m.ExcPending
										if v72 != 0 {
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
							v73 = int32(1)
							if v53&v73 == int32(0) {
								v95 = v73
								m.G0 = v11 - int32(-64)
								return v95
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return int32(0)
								} else {
									v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v82
									F_errmsg_internal(m, int32(559945), v9+int32(-48))
									mBase = m.M
									v88 = m.ExcPending
									if v88 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(520932), int32(1624), int32(431009))
										mBase = m.M
										v93 = m.ExcPending
										if v93 != 0 {
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
					} else {
						v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+4)))
						if v51 != 0 {
							v95 = int32(0)
							m.G0 = v11 - int32(-64)
							return v95
						} else {
							v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+36)))
							if l1 == int32(0) {
								v56 = int32(1)
								if v53&v56 != 0 {
									v95 = v56
									m.G0 = v11 - int32(-64)
									return v95
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v62 = m.ExcPending
									if v62 != 0 {
										return int32(0)
									} else {
										v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v11))) = v63
										F_errmsg_internal(m, int32(559793), v11)
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(520932), int32(1618), int32(431009))
											mBase = m.M
											v72 = m.ExcPending
											if v72 != 0 {
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
								v73 = int32(1)
								if v53&v73 == int32(0) {
									v95 = v73
									m.G0 = v11 - int32(-64)
									return v95
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
										return int32(0)
									} else {
										v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v82
										F_errmsg_internal(m, int32(559945), v9+int32(-48))
										mBase = m.M
										v88 = m.ExcPending
										if v88 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(520932), int32(1624), int32(431009))
											mBase = m.M
											v93 = m.ExcPending
											if v93 != 0 {
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
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l5))) = int32(0)
			v95 = v13
			m.G0 = v11 - int32(-64)
			return v95
		}
	}
}
func F_IsPreferredType(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	F_get_type_category_preferred(m, l1, v6+int32(15), v6+int32(14))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if l0 != 0 {
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+15)))
			if v17 != l0&int32(255) {
				v22 = int32(0)
			} else {
				v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+14)))
				v22 = v21
			}
		} else {
			v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+14)))
			v22 = v21
		}
		m.G0 = v6 + int32(16)
		return v22 & int32(1)
	}
}
func F_IssuePendingWritebacks(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int64
	_ = v39
	var v40 int64
	_ = v40
	var v44 int64
	_ = v44
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v96 int32
	_ = v96
	var __phi96 int32
	_ = __phi96
	var v98 int32
	_ = v98
	var __phi98 int32
	_ = __phi98
	var v100 int32
	_ = v100
	var __phi100 int32
	_ = __phi100
	var v101 int32
	_ = v101
	var __phi101 int32
	_ = __phi101
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v146 int32
	_ = v146
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int64
	_ = v162
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v197 int32
	_ = v197
	var v213 int64
	_ = v213
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v225 int64
	_ = v225
	var v226 int64
	_ = v226
	var v230 int64
	_ = v230
	var v276 int32
	_ = v276
	var v283 int64
	_ = v283
	var v287 int32
	_ = v287
	var v299 int32
	_ = v299
	var v306 int64
	_ = v306
	var v310 int32
	_ = v310
	var v321 int32
	_ = v321
	var v328 int64
	_ = v328
	var v334 int64
	_ = v334
	var v339 int32
	_ = v339
	v21 = m.G0
	v23 = v21 - int32(32)
	m.G0 = v23
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v27 = l0 + int32(8)
	F_sort_pending_writebacks(m, v27, v25)
	mBase = m.M
	v30 = int32(*(*uint8)(unsafe.Add(mBase, _consts[624])))
	v33 = m.G0
	v35 = v33 - int32(16)
	m.G0 = v35
	if v30 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L3
L3:
	;
	m.G0 = v23 + int32(32)
	return
L4:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) < v48 {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	F___clock_gettime(m, int32(1), v35)
	mBase = m.M
	v39 = int64(*(*int32)(unsafe.Add(mBase, uint32(v35)+8)))
	v40 = *(*int64)(unsafe.Add(mBase, uint32(v35)))
	v44 = v39 + v40*int64(1000000000)
	goto L7
L6:
	;
	v44 = int64(0)
	goto L7
L7:
	;
	m.G0 = v35 + int32(16)
	goto L4
L8:
	;
	v54 = int32(0)
	v58 = v48
	goto L11
L9:
	;
	v197 = v48
	goto L10
L10:
	;
	v213 = int64(0)
	v217 = m.G0
	v219 = v217 - int32(16)
	m.G0 = v219
	if v44 != v213 {
		goto L33
	} else {
		goto L34
	}
L11:
	;
	v74 = v27 + v54*int32(20)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+16))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v77
	v83 = int32(1)
	v85 = v54 + v83
	if v58 <= v85 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v197 = v189
	goto L10
L13:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v23)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v160
	v162 = *(*int64)(unsafe.Add(mBase, uint32(v23)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v23)+8)) = v162
	v167 = F_smgropen(m, v23+int32(8), int32(-1))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L28
	} else {
		goto L29
	}
L14:
	;
	v146 = v83
	v159 = v85
	goto L13
L15:
	;
	goto L16
L16:
	;
	__phi96 = v74
	__phi98 = v85
	__phi100 = v83
	__phi101 = int32(0)
	v96 = __phi96
	v98 = __phi98
	v100 = __phi100
	v101 = __phi101
	goto L17
L17:
	;
	v115 = v27 + v98*int32(20)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+8))
	if v116 != v79 {
		v146 = v100
		v159 = v98
		goto L13
	} else {
		goto L19
	}
L18:
	;
	v146 = v134
	v159 = v58
	goto L13
L19:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	if v118 != v78 {
		v146 = v100
		v159 = v98
		goto L13
	} else {
		goto L20
	}
L20:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	if v120 != v77 {
		v146 = v100
		v159 = v98
		goto L13
	} else {
		goto L21
	}
L21:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v96)+12))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v115)+12))
	if v122 != v123 {
		v146 = v100
		v159 = v98
		goto L13
	} else {
		goto L22
	}
L22:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v96)+16))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v115)+16))
	if v125 != v126 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	if v126 != v125+int32(1) {
		v146 = v100
		v159 = v98
		goto L13
	} else {
		goto L26
	}
L24:
	;
	v133 = v96
	v134 = v100
	goto L25
L25:
	;
	v137 = v101 + int32(1)
	if v137 != v58+(v54^int32(-1)) {
		__phi96 = v133
		__phi98 = v101 + (v54 + int32(2))
		__phi100 = v134
		__phi101 = v137
		v96 = __phi96
		v98 = __phi98
		v100 = __phi100
		v101 = __phi101
		goto L17
	} else {
		goto L27
	}
L26:
	;
	v133 = v115
	v134 = v100 + int32(1)
	goto L25
L27:
	;
	goto L18
L28:
	;
	return
L29:
	;
	v169 = int32(4556748)
	v171 = *(*int32)(unsafe.Add(mBase, _consts[410]))
	*(*int32)(unsafe.Add(mBase, _consts[410])) = v171 + int32(1)
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v167)+36))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v175*int32(80))+uint32(_consts[625])))
	m.T0[v180].(func(*base.Module, int32, int32, int32, int32))(m, v167, v76, v75, v146)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v183 = int32(4556748)
	v185 = *(*int32)(unsafe.Add(mBase, _consts[410]))
	*(*int32)(unsafe.Add(mBase, _consts[410])) = v185 - int32(1)
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v159 < v189 {
		v54 = v159
		v58 = v189
		goto L11
	} else {
		goto L31
	}
L31:
	;
	goto L12
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(0)
	goto L3
L33:
	;
	F___clock_gettime(m, int32(1), v219)
	mBase = m.M
	v225 = int64(*(*int32)(unsafe.Add(mBase, uint32(v219)+8)))
	v226 = *(*int64)(unsafe.Add(mBase, uint32(v219)))
	v230 = v225 + (v226*int64(1000000000) - v44)
	goto L37
L34:
	;
	goto L35
L35:
	;
	v321 = l1 << (uint(int32(6)) % 32)
	v328 = *(*int64)(unsafe.Add(mBase, uint32(v321)+uint32(_consts[626])))
	*(*int64)(unsafe.Add(mBase, uint32(v321)+uint32(_consts[626]))) = v328 + base.I64_extend_i32_u(v197)
	v334 = *(*int64)(unsafe.Add(mBase, uint32(v321)+uint32(_consts[627])))
	*(*int64)(unsafe.Add(mBase, uint32(v321)+uint32(_consts[627]))) = v334 + v213
	F_pgstat_count_backend_io_op(m, int32(0), l1, int32(4), v197, v213)
	mBase = m.M
	v339 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[88])) = uint8(v339)
	*(*uint8)(unsafe.Add(mBase, _consts[605])) = uint8(v339)
	m.G0 = v219 + int32(16)
	goto L32
L36:
	;
	v276 = l1 << (uint(int32(6)) % 32)
	v283 = *(*int64)(unsafe.Add(mBase, uint32(v276)+uint32(_consts[628])))
	*(*int64)(unsafe.Add(mBase, uint32(v276)+uint32(_consts[628]))) = v283 + v230
	v287 = *(*int32)(unsafe.Add(mBase, _consts[20]))
	if base.Ui32(int32(16)) < base.Ui32(v287) {
		goto L46
	} else {
		goto L47
	}
L37:
	;
	goto L39
L39:
	;
	goto L40
L40:
	;
	goto L36
L46:
	;
	goto L35
L47:
	;
	if int32(1)<<(uint(v287)%32)&int32(115186) == int32(0) {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v299 = l1 << (uint(int32(6)) % 32)
	v306 = *(*int64)(unsafe.Add(mBase, uint32(v299)+uint32(_consts[629])))
	*(*int64)(unsafe.Add(mBase, uint32(v299)+uint32(_consts[629]))) = v306 + v230
	v310 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[88])) = uint8(v310)
	*(*uint8)(unsafe.Add(mBase, _consts[630])) = uint8(v310)
	goto L46
}
func F_i2toi4(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+20)))
	return v2
}
func F_icnlikesel(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 float64
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = F_get_negator(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if v11 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(221149), int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(517715), int32(773), int32(322473))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v31 = int32(1)
			v33 = F_patternsel_common(m, v9, v11, int32(0), v7, v6, v8, v31, v31)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				v35 = F_Float8GetDatum(m, v33)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					return v35
				}
			}
		}
	}
}
func F_icu_validate_locale(m *base.Module) {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	F_errstart_cold(m, int32(21), int32(0))
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		F_errcode(m, int32(1088))
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			F_errmsg(m, int32(453231), int32(0))
			v11 = m.ExcPending
			if v11 != 0 {
				return
			} else {
				F_errfinish(m, int32(525445), int32(1674), int32(417944))
				v16 = m.ExcPending
				if v16 != 0 {
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
func F_identify_opfamily_groups(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
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
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var __phi59 int32
	_ = __phi59
	var v60 int32
	_ = v60
	var __phi60 int32
	_ = __phi60
	var v61 int32
	_ = v61
	var __phi61 int32
	_ = __phi61
	var v62 int32
	_ = v62
	var __phi62 int32
	_ = __phi62
	var v63 int32
	_ = v63
	var __phi63 int32
	_ = __phi63
	var v64 int32
	_ = v64
	var __phi64 int32
	_ = __phi64
	var v65 int32
	_ = v65
	var __phi65 int32
	_ = __phi65
	var v66 int32
	_ = v66
	var __phi66 int32
	_ = __phi66
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v107 int64
	_ = v107
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v150 int64
	_ = v150
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int64
	_ = v208
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v241 int32
	_ = v241
	v3 = int32(0)
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+37)))
	if v16 != int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v241
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L39
	} else {
		goto L52
	}
L3:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+37)))
	if v19 == int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v22 <= int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v34 <= int32(0) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v32 = v3
	v33 = int32(0)
	goto L5
L7:
	;
	goto L8
L8:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+56))
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+22)))
	v32 = int32(1)
	v33 = v28 + v29
	goto L5
L9:
	;
	v43 = v3
	v44 = int32(0)
	goto L11
L10:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+56))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+22)))
	v43 = v39 + v40
	v44 = int32(1)
	goto L11
L11:
	;
	if v43|v33 == int32(0) {
		v241 = v3
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v48 = int32(0)
	v52 = int32(48)
	__phi59 = v43
	__phi60 = v48
	__phi61 = base.B2i32(v43 != v48)
	__phi62 = v33
	__phi63 = base.B2i32(v33 != v48)
	__phi64 = v32
	__phi65 = v44
	__phi66 = v3
	v59 = __phi59
	v60 = __phi60
	v61 = __phi61
	v62 = __phi62
	v63 = __phi63
	v64 = __phi64
	v65 = __phi65
	v66 = __phi66
	goto L13
L13:
	;
	v73 = v62 + int32(12)
	v75 = base.B2i32(v60 != int32(0))
	v79 = v59
	v81 = v61
	v85 = v65
	goto L15
L15:
	;
	if v63&v75 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v177 = F_palloc(m, int32(24))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L39
	} else {
		goto L40
	}
L17:
	;
	if v81&v75 == int32(0) {
		v175 = v81
		goto L28
	} else {
		goto L29
	}
L18:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v94 != v95 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	if v97 != v98 {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	v100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v62)+16)))
	if base.Ui32((v100-int32(1))&int32(65535)) <= base.Ui32(int32(62)) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v107 = *(*int64)(unsafe.Add(mBase, uint32(v60)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v60)+8)) = v107 | int64(1)<<(uint(base.I64_extend_i32_u(v100))%64)
	goto L23
L22:
	;
	goto L23
L23:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v64 < v114 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0+v52+v64<<(uint(int32(2))%32))))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+56))
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+22)))
	v126 = v120 + v121
	v127 = v64 + int32(1)
	goto L26
L25:
	;
	v126 = int32(0)
	v127 = v64
	goto L26
L26:
	;
	v128 = int32(0)
	if v79|v126 != 0 {
		__phi59 = v79
		__phi61 = base.B2i32(v79 != v128)
		__phi62 = v126
		__phi63 = base.B2i32(v126 != v128)
		__phi64 = v127
		__phi65 = v85
		v59 = __phi59
		v61 = __phi61
		v62 = __phi62
		v63 = __phi63
		v64 = __phi64
		v65 = __phi65
		goto L13
	} else {
		goto L27
	}
L27:
	;
	v241 = v66
	goto L1
L28:
	;
	goto L16
L29:
	;
	v136 = int32(1)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v137 != v138 {
		v175 = v136
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	if v140 != v141 {
		v175 = v136
		goto L28
	} else {
		goto L31
	}
L31:
	;
	v143 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v79)+16)))
	if base.Ui32((v143-int32(1))&int32(65535)) <= base.Ui32(int32(62)) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v150 = *(*int64)(unsafe.Add(mBase, uint32(v60)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v60)+16)) = v150 | int64(1)<<(uint(base.I64_extend_i32_u(v143))%64)
	goto L34
L33:
	;
	goto L34
L34:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v85 < v157 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l1+v52+v85<<(uint(int32(2))%32))))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)+56))
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163)+22)))
	v168 = v163 + v164
	v169 = v85 + int32(1)
	goto L37
L36:
	;
	v168 = int32(0)
	v169 = v85
	goto L37
L37:
	;
	v171 = base.B2i32(v168 != int32(0))
	if (v63|v171)&int32(1) != 0 {
		v79 = v168
		v81 = v171
		v85 = v169
		goto L15
	} else {
		goto L38
	}
L38:
	;
	v241 = v66
	goto L1
L39:
	;
	return int32(0)
L40:
	;
	if v63&int32(1) == int32(0) {
		goto L44
	} else {
		goto L45
	}
L41:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
	v208 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v177)+8)) = v208
	*(*int32)(unsafe.Add(mBase, uint32(v177)+4)) = v207
	*(*int64)(unsafe.Add(mBase, uint32(v177)+16)) = v208
	v213 = int32(0)
	v217 = F_lappend(m, v66, v177)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L39
	} else {
		goto L51
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v177))) = v186
	v206 = v73
	goto L41
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v177))) = v197
	v206 = v79 + int32(12)
	goto L41
L44:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
	v197 = v185
	goto L43
L45:
	;
	goto L46
L46:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
	if v175&int32(1) == int32(0) {
		goto L42
	} else {
		goto L47
	}
L47:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
	if base.Ui32(v186) < base.Ui32(v191) {
		goto L42
	} else {
		goto L48
	}
L48:
	;
	if v191 != v186 {
		v197 = v191
		goto L43
	} else {
		goto L49
	}
L49:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
	if base.Ui32(v194) < base.Ui32(v195) {
		goto L42
	} else {
		goto L50
	}
L50:
	;
	v197 = v186
	goto L43
L51:
	;
	__phi59 = v79
	__phi60 = v177
	__phi61 = base.B2i32(v79 != v213)
	__phi63 = base.B2i32(v62 != v213)
	__phi65 = v85
	__phi66 = v217
	v59 = __phi59
	v60 = __phi60
	v61 = __phi61
	v63 = __phi63
	v65 = __phi65
	v66 = __phi66
	goto L13
L52:
	;
	F_errmsg_internal(m, int32(532308), int32(0))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L39
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(524885), int32(54), int32(145377))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L39
	} else {
		goto L54
	}
L54:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_idx(m *base.Module, l0 int32) int32 {
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_pg_detoast_datum(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
		if v9 != 0 {
			v10 = F_array_contains_nulls(m, v5)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				if v10 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(67108994))
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(162746), int32(0))
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(521579), int32(267), int32(30597))
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
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
					v12 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
					v15 = F_ArrayGetNItems(m, v12, v5+int32(16))
					mBase = m.M
					v16 = m.ExcPending
					if v16 != 0 {
						return int32(0)
					} else {
						if v15 != 0 {
							v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							v18 = F_intarray_match_first(m, v5, v17)
							mBase = m.M
							v19 = m.ExcPending
							if v19 != 0 {
								return int32(0)
							} else {
								v21 = v18
								v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								if v22 != v5 {
									F_pfree(m, v5)
									mBase = m.M
									v25 = m.ExcPending
									if v25 != 0 {
										return int32(0)
									} else {
										return v21
									}
								} else {
									return v21
								}
							}
						} else {
							v21 = int32(0)
							v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							if v22 != v5 {
								F_pfree(m, v5)
								mBase = m.M
								v25 = m.ExcPending
								if v25 != 0 {
									return int32(0)
								} else {
									return v21
								}
							} else {
								return v21
							}
						}
					}
				}
			}
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
			v15 = F_ArrayGetNItems(m, v12, v5+int32(16))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				if v15 != 0 {
					v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					v18 = F_intarray_match_first(m, v5, v17)
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return int32(0)
					} else {
						v21 = v18
						v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						if v22 != v5 {
							F_pfree(m, v5)
							mBase = m.M
							v25 = m.ExcPending
							if v25 != 0 {
								return int32(0)
							} else {
								return v21
							}
						} else {
							return v21
						}
					}
				} else {
					v21 = int32(0)
					v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v22 != v5 {
						F_pfree(m, v5)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return int32(0)
						} else {
							return v21
						}
					} else {
						return v21
					}
				}
			}
		}
	}
}
func F_indonesian_ISO_8859_1_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v114 int32
	_ = v114
	var v122 int32
	_ = v122
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v326 int32
	_ = v326
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v400 int32
	_ = v400
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v496 int32
	_ = v496
	v2 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v2
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v18 < v9 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if int32(0) <= v58 {
		goto L16
	} else {
		goto L17
	}
L2:
	;
	v20 = v9
	goto L4
L3:
	;
	v20 = v18
	goto L4
L4:
	;
	v27 = v9
	goto L6
L5:
	;
	v58 = v38
	goto L1
L6:
	;
	if v27 == v20 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v58 = int32(-1)
	goto L1
L9:
	;
	goto L10
L10:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31+v27))))
	if int32(117) < v33 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v50 = v27 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v50
	v27 = v50
	goto L6
L12:
	;
	v35 = v33 - int32(97)
	if v35 < int32(0) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v38 = int32(1)
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v35)>>(uint(int32(3))%32)))+uint32(_consts[1435]))))
	if int32(base.Ui32(v42)>>(uint(v35&int32(7))%32))&v38 != 0 {
		goto L5
	} else {
		goto L14
	}
L14:
	;
	goto L11
L16:
	;
	v62 = v58
	goto L19
L17:
	;
	goto L18
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v9
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+4))
	if v132 < int32(3) {
		v496 = v2
		goto L37
	} else {
		goto L38
	}
L19:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v66 + v62
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v69)+4)) = v70 + int32(1)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v82 < v81 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L18
L21:
	;
	if int32(0) <= v122 {
		v62 = v122
		goto L19
	} else {
		goto L36
	}
L22:
	;
	v84 = v81
	goto L24
L23:
	;
	v84 = v82
	goto L24
L24:
	;
	v91 = v81
	goto L26
L25:
	;
	v122 = v102
	goto L21
L26:
	;
	if v91 == v84 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v122 = int32(-1)
	goto L21
L29:
	;
	goto L30
L30:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95+v91))))
	if int32(117) < v97 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v114 = v91 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v114
	v91 = v114
	goto L26
L32:
	;
	v99 = v97 - int32(97)
	if v99 < int32(0) {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v102 = int32(1)
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v99)>>(uint(int32(3))%32)))+uint32(_consts[1435]))))
	if int32(base.Ui32(v106)>>(uint(v99&int32(7))%32))&v102 != 0 {
		goto L25
	} else {
		goto L34
	}
L34:
	;
	goto L31
L36:
	;
	goto L20
L37:
	;
	return v496
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v131))) = int32(0)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v137
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v139
	if v139-int32(2) <= v137 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v178
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v175)+4))
	if v180 < int32(3) {
		goto L51
	} else {
		goto L52
	}
L40:
	;
	v175 = v131
	v177 = int32(0)
	goto L39
L41:
	;
	goto L42
L42:
	;
	v146 = int32(0)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147+v139-int32(1)))))
	switch v151 - int32(104) {
	case 0, 6:
		goto L43
	default:
		v175 = v131
		v177 = v146
		goto L39
	}
L43:
	;
	v156 = F_find_among_b(m, l0, int32(4249200), int32(3))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	return int32(0)
L45:
	;
	if v156 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v175 = v162
	v177 = v146
	goto L39
L47:
	;
	goto L48
L48:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v163
	v165 = F_slice_del(m, l0)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L44
	} else {
		goto L49
	}
L49:
	;
	if v165 < int32(0) {
		v496 = v165
		goto L37
	} else {
		goto L50
	}
L50:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v169)+4))
	v171 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v169)+4)) = v170 - v171
	v175 = v169
	v177 = v171
	goto L39
L51:
	;
	return int32(0)
L52:
	;
	goto L53
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v178
	v187 = v178 - int32(1)
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v187 <= v188 {
		v216 = v175
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v218
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v216)+4))
	if v221 < int32(3) {
		v496 = int32(0)
		goto L37
	} else {
		goto L63
	}
L55:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190+v187))))
	if base.B2i32(v192 != int32(117))&base.B2i32(v192 != int32(97)) != 0 {
		v216 = v175
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v200 = F_find_among_b(m, l0, int32(4249264), int32(3))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L44
	} else {
		goto L57
	}
L57:
	;
	if v200 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v216 = v204
	goto L54
L59:
	;
	goto L60
L60:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v205
	v207 = F_slice_del(m, l0)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L44
	} else {
		goto L61
	}
L61:
	;
	if v207 < int32(0) {
		v496 = v207
		goto L37
	} else {
		goto L62
	}
L62:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v211)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v211)+4)) = v212 - int32(1)
	v216 = v211
	goto L54
L63:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v218
	v227 = v218 + int32(1)
	if v224 <= v227 {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v218
	v496 = int32(1)
	goto L37
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v218
	v468 = F_r_remove_second_order_prefix_1(m, l0)
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L44
	} else {
		goto L146
	}
L66:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229+v227))))
	switch v231 - int32(101) {
	case 0, 4:
		goto L67
	default:
		goto L65
	}
L67:
	;
	v236 = F_find_among(m, l0, int32(4249328), int32(12))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L44
	} else {
		goto L68
	}
L68:
	;
	if v236 == int32(0) {
		goto L65
	} else {
		goto L69
	}
L69:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v240
	switch v236 - int32(1) {
	case 0:
		goto L78
	case 1:
		goto L77
	case 2:
		goto L76
	case 3:
		goto L75
	case 4:
		goto L72
	case 5:
		goto L71
	default:
		goto L70
	}
L70:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v434)+4))
	if v435 < int32(3) {
		goto L64
	} else {
		goto L129
	}
L71:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v358))) = int32(3)
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v358)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v358)+4)) = v361 - int32(1)
	v365 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v375 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v375 < v365 {
		goto L109
	} else {
		goto L110
	}
L72:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v285 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v284))) = v285
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v284)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v284)+4)) = v287 - v285
	v291 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v301 < v291 {
		goto L88
	} else {
		goto L89
	}
L73:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v278)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v278)+4)) = v280 - int32(1)
	goto L70
L74:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v278 = v277
	goto L73
L75:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v267))) = int32(3)
	v272 = F_slice_from_s(m, l0, int32(1), int32(2220507))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L44
	} else {
		goto L85
	}
L76:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v259 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v258))) = v259
	v263 = F_slice_from_s(m, l0, v259, int32(2220506))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L44
	} else {
		goto L83
	}
L77:
	;
	v251 = F_slice_del(m, l0)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L44
	} else {
		goto L81
	}
L78:
	;
	v244 = F_slice_del(m, l0)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L44
	} else {
		goto L79
	}
L79:
	;
	if v244 < int32(0) {
		v496 = v244
		goto L37
	} else {
		goto L80
	}
L80:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v248))) = int32(1)
	v278 = v248
	goto L73
L81:
	;
	if v251 < int32(0) {
		v496 = v251
		goto L37
	} else {
		goto L82
	}
L82:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v255))) = int32(3)
	v278 = v255
	goto L73
L83:
	;
	if int32(0) <= v263 {
		goto L74
	} else {
		goto L84
	}
L84:
	;
	v496 = v263
	goto L37
L85:
	;
	if v272 < int32(0) {
		v496 = v272
		goto L37
	} else {
		goto L86
	}
L86:
	;
	goto L74
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v291
	if v344 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L88:
	;
	v303 = v291
	goto L90
L89:
	;
	v303 = v301
	goto L90
L90:
	;
	goto L92
L91:
	;
	v344 = v340
	goto L87
L92:
	;
	if v291 == v303 {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	v340 = int32(0)
	goto L91
L94:
	;
	v344 = int32(-1)
	goto L87
L95:
	;
	goto L96
L96:
	;
	v315 = int32(1)
	v316 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v316+v291))))
	if int32(117) < v318 {
		v340 = v315
		goto L91
	} else {
		goto L97
	}
L97:
	;
	v320 = v318 - int32(97)
	if v320 < int32(0) {
		v340 = v315
		goto L91
	} else {
		goto L98
	}
L98:
	;
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v320)>>(uint(int32(3))%32)))+uint32(_consts[1435]))))
	if int32(base.Ui32(v326)>>(uint(v320&int32(7))%32))&int32(1) == int32(0) {
		v340 = v315
		goto L91
	} else {
		goto L99
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v291 + int32(1)
	goto L100
L100:
	;
	goto L93
L101:
	;
	v350 = F_slice_from_s(m, l0, int32(1), int32(2220508))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L44
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	v354 = F_slice_del(m, l0)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L44
	} else {
		goto L106
	}
L104:
	;
	if int32(0) <= v350 {
		goto L70
	} else {
		goto L105
	}
L105:
	;
	v496 = v350
	goto L37
L106:
	;
	if int32(0) <= v354 {
		goto L70
	} else {
		goto L107
	}
L107:
	;
	v496 = v354
	goto L37
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v365
	if v418 == int32(0) {
		goto L122
	} else {
		goto L123
	}
L109:
	;
	v377 = v365
	goto L111
L110:
	;
	v377 = v375
	goto L111
L111:
	;
	goto L113
L112:
	;
	v418 = v414
	goto L108
L113:
	;
	if v365 == v377 {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	v414 = int32(0)
	goto L112
L115:
	;
	v418 = int32(-1)
	goto L108
L116:
	;
	goto L117
L117:
	;
	v389 = int32(1)
	v390 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v390+v365))))
	if int32(117) < v392 {
		v414 = v389
		goto L112
	} else {
		goto L118
	}
L118:
	;
	v394 = v392 - int32(97)
	if v394 < int32(0) {
		v414 = v389
		goto L112
	} else {
		goto L119
	}
L119:
	;
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v394)>>(uint(int32(3))%32)))+uint32(_consts[1435]))))
	if int32(base.Ui32(v400)>>(uint(v394&int32(7))%32))&int32(1) == int32(0) {
		v414 = v389
		goto L112
	} else {
		goto L120
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v365 + int32(1)
	goto L121
L121:
	;
	goto L114
L122:
	;
	v424 = F_slice_from_s(m, l0, int32(1), int32(2220509))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L44
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	v428 = F_slice_del(m, l0)
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L44
	} else {
		goto L127
	}
L125:
	;
	if int32(0) <= v424 {
		goto L70
	} else {
		goto L126
	}
L126:
	;
	v496 = v424
	goto L37
L127:
	;
	if v428 < int32(0) {
		v496 = v428
		goto L37
	} else {
		goto L128
	}
L128:
	;
	goto L70
L129:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v438
	v440 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v440
	v442 = F_r_remove_suffix_1(m, l0)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L44
	} else {
		goto L130
	}
L130:
	;
	if v442 == int32(0) {
		goto L64
	} else {
		goto L131
	}
L131:
	;
	if v442 < int32(0) {
		v496 = v442
		goto L37
	} else {
		goto L132
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v438
	v449 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v449)+4))
	if v450 < int32(3) {
		goto L64
	} else {
		goto L133
	}
L133:
	;
	v453 = F_r_remove_second_order_prefix_1(m, l0)
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L44
	} else {
		goto L135
	}
L134:
	;
	if int32(0) <= v453 {
		goto L64
	} else {
		goto L139
	}
L135:
	;
	if v453 != 0 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v458 = int32(base.Ui32(v453) >> (uint(int32(31)) % 32))
	goto L138
L137:
	;
	v458 = int32(6)
	goto L138
L138:
	;
	switch v458 {
	case 0, 6:
		goto L64
	default:
		goto L134
	}
L139:
	;
	if v453 < int32(0) {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v463 = v453
	goto L142
L141:
	;
	v463 = v177
	goto L142
L142:
	;
	if v453 != 0 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v464 = v463
	goto L145
L144:
	;
	v464 = v177
	goto L145
L145:
	;
	return v464
L146:
	;
	if v468 < int32(0) {
		v496 = v468
		goto L37
	} else {
		goto L147
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v218
	v473 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v473)+4))
	if v474 < int32(3) {
		goto L64
	} else {
		goto L148
	}
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v218
	v478 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v478
	v480 = F_r_remove_suffix_1(m, l0)
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L44
	} else {
		goto L150
	}
L149:
	;
	if int32(0) <= v480 {
		goto L154
	} else {
		goto L155
	}
L150:
	;
	if v480 != 0 {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v485 = int32(base.Ui32(v480) >> (uint(int32(31)) % 32))
	goto L153
L152:
	;
	v485 = int32(8)
	goto L153
L153:
	;
	switch v485 {
	case 0, 8:
		goto L64
	default:
		goto L149
	}
L154:
	;
	v489 = int32(1)
	goto L156
L155:
	;
	v489 = v480
	goto L156
L156:
	;
	return v489
}
func F_inetand(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = F_pg_detoast_datum_packed(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v21 = F_pg_detoast_datum_packed(m, v20)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			v24 = F_palloc0(m, int32(22))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				v26 = int32(1)
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
				v30 = v28 & v26
				if v30 != 0 {
					v31 = v26
				} else {
					v31 = int32(4)
				}
				v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+v31))))
				v34 = int32(1)
				v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
				v38 = v36 & v34
				if v38 != 0 {
					v39 = v34
				} else {
					v39 = int32(4)
				}
				v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21+v39))))
				if v33 == v41 {
					v43 = int32(1)
					v44 = v24 + v43
					v46 = v24 + int32(4)
					v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
					if v47&v43 != 0 {
						v50 = v44
					} else {
						v50 = v46
					}
					v51 = int32(2)
					v52 = v50 + v51
					if v33 == v51 {
						v57 = int32(3)
					} else {
						v57 = int32(15)
					}
					v59 = v21 + int32(1)
					v61 = v21 + int32(4)
					if v38 != 0 {
						v62 = v59
					} else {
						v62 = v61
					}
					v64 = v62 + int32(2)
					v66 = v16 + int32(1)
					v68 = v16 + int32(4)
					if v30 != 0 {
						v69 = v66
					} else {
						v69 = v68
					}
					v71 = v69 + int32(2)
					v72 = v57
					for {
						v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72+v64))))
						v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72+v71))))
						v91 = v88 & v90
						*(*uint8)(unsafe.Add(mBase, uint32(v72+v52))) = uint8(v91)
						v94 = v72 - int32(1)
						v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94+v64))))
						v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94+v71))))
						v100 = v97 & v99
						*(*uint8)(unsafe.Add(mBase, uint32(v52+v94))) = uint8(v100)
						if v94 != 0 {
							v72 = v72 - int32(2)
							continue
						} else {
							break
						}
						break
					}
					v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
					if v104&int32(1) != 0 {
						v107 = v44
					} else {
						v107 = v46
					}
					v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
					if v108&int32(1) != 0 {
						v111 = v66
					} else {
						v111 = v68
					}
					v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+1)))
					v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
					if v113&int32(1) != 0 {
						v116 = v59
					} else {
						v116 = v61
					}
					v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+1)))
					if base.Ui32(v117) < base.Ui32(v112) {
						v119 = v112
					} else {
						v119 = v117
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v107)+1)) = uint8(v119)
					v121 = int32(1)
					v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
					if v123&v121 != 0 {
						v126 = v121
					} else {
						v126 = int32(4)
					}
					v128 = int32(1)
					v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
					if v130&v128 != 0 {
						v133 = v128
					} else {
						v133 = int32(4)
					}
					v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+v133))))
					*(*uint8)(unsafe.Add(mBase, uint32(v24+v126))) = uint8(v135)
					if v135 == int32(2) {
						v141 = int32(40)
					} else {
						v141 = int32(88)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v24))) = v141
					return v24
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v147 = m.ExcPending
					if v147 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v150 = m.ExcPending
						if v150 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(167660), int32(0))
							mBase = m.M
							v154 = m.ExcPending
							if v154 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(523897), int32(1858), int32(449375))
								mBase = m.M
								v159 = m.ExcPending
								if v159 != 0 {
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
	}
}
func F_inetor(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = F_pg_detoast_datum_packed(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v21 = F_pg_detoast_datum_packed(m, v20)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			v24 = F_palloc0(m, int32(22))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				v26 = int32(1)
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
				v30 = v28 & v26
				if v30 != 0 {
					v31 = v26
				} else {
					v31 = int32(4)
				}
				v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+v31))))
				v34 = int32(1)
				v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
				v38 = v36 & v34
				if v38 != 0 {
					v39 = v34
				} else {
					v39 = int32(4)
				}
				v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21+v39))))
				if v33 == v41 {
					v43 = int32(1)
					v44 = v24 + v43
					v46 = v24 + int32(4)
					v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
					if v47&v43 != 0 {
						v50 = v44
					} else {
						v50 = v46
					}
					v51 = int32(2)
					v52 = v50 + v51
					if v33 == v51 {
						v57 = int32(3)
					} else {
						v57 = int32(15)
					}
					v59 = v21 + int32(1)
					v61 = v21 + int32(4)
					if v38 != 0 {
						v62 = v59
					} else {
						v62 = v61
					}
					v64 = v62 + int32(2)
					v66 = v16 + int32(1)
					v68 = v16 + int32(4)
					if v30 != 0 {
						v69 = v66
					} else {
						v69 = v68
					}
					v71 = v69 + int32(2)
					v72 = v57
					for {
						v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72+v64))))
						v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72+v71))))
						v91 = v88 | v90
						*(*uint8)(unsafe.Add(mBase, uint32(v72+v52))) = uint8(v91)
						v94 = v72 - int32(1)
						v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94+v64))))
						v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94+v71))))
						v100 = v97 | v99
						*(*uint8)(unsafe.Add(mBase, uint32(v52+v94))) = uint8(v100)
						if v94 != 0 {
							v72 = v72 - int32(2)
							continue
						} else {
							break
						}
						break
					}
					v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
					if v104&int32(1) != 0 {
						v107 = v44
					} else {
						v107 = v46
					}
					v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
					if v108&int32(1) != 0 {
						v111 = v66
					} else {
						v111 = v68
					}
					v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+1)))
					v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
					if v113&int32(1) != 0 {
						v116 = v59
					} else {
						v116 = v61
					}
					v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+1)))
					if base.Ui32(v117) < base.Ui32(v112) {
						v119 = v112
					} else {
						v119 = v117
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v107)+1)) = uint8(v119)
					v121 = int32(1)
					v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
					if v123&v121 != 0 {
						v126 = v121
					} else {
						v126 = int32(4)
					}
					v128 = int32(1)
					v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
					if v130&v128 != 0 {
						v133 = v128
					} else {
						v133 = int32(4)
					}
					v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+v133))))
					*(*uint8)(unsafe.Add(mBase, uint32(v24+v126))) = uint8(v135)
					if v135 == int32(2) {
						v141 = int32(40)
					} else {
						v141 = int32(88)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v24))) = v141
					return v24
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v147 = m.ExcPending
					if v147 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v150 = m.ExcPending
						if v150 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(167619), int32(0))
							mBase = m.M
							v154 = m.ExcPending
							if v154 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(523897), int32(1890), int32(219311))
								mBase = m.M
								v159 = m.ExcPending
								if v159 != 0 {
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
	}
}
func F_infix_2(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
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
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
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
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
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
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v16 = l1
	goto L5
L1:
	;
	m.G0 = v13 + int32(32)
	return
L2:
	;
	v420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	if v420 != 0 {
		goto L75
	} else {
		goto L76
	}
L3:
	;
	v389 = v36
	v391 = v34
	goto L71
L4:
	;
	v195 = v27 + int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v195
	v199 = v16 | base.B2i32(v45 != int32(124))
	if v199&int32(1) != 0 {
		goto L39
	} else {
		goto L40
	}
L5:
	;
	F_check_stack_depth(m)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v103 = v101 - v102
	v105 = v103 + int32(3)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v106 <= v105 {
		goto L22
	} else {
		goto L23
	}
L7:
	;
	return
L8:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27))))
	if v28 == int32(2) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v32 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27)+10)))
	v33 = v31 + v32
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v37 = v35 - v36
	v39 = v37 + int32(6)
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+9)))
	if v34 <= v39+v40<<(uint(int32(1))%32) {
		goto L3
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if v45 != int32(33) {
		goto L4
	} else {
		goto L13
	}
L12:
	;
	v412 = v27
	v413 = v35
	goto L2
L13:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v50 = v48 - v49
	v52 = v50 + int32(2)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v53 <= v52 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v56 = v53
	v57 = v49
	goto L17
L15:
	;
	v78 = v48
	goto L16
L16:
	;
	v85 = int32(33)
	*(*uint8)(unsafe.Add(mBase, uint32(v78))) = uint8(v85)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v87 + int32(1)
	v91 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v87)+1)) = uint8(v91)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v94 + int32(12)
	v98 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v94)+12)))
	if v98 != int32(3) {
		v16 = v91
		goto L5
	} else {
		goto L21
	}
L17:
	;
	v66 = v56 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v66
	v68 = F_repalloc(m, v57, v66)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L7
	} else {
		goto L19
	}
L18:
	;
	v78 = v71
	goto L16
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v68
	v71 = v68 + v50
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v71
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v73 <= v52 {
		v56 = v73
		v57 = v68
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	goto L6
L22:
	;
	v109 = v106
	v110 = v102
	goto L25
L23:
	;
	v131 = v101
	goto L24
L24:
	;
	v140 = F_pg_sprintf(m, v131, int32(783364), int32(0))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L7
	} else {
		goto L29
	}
L25:
	;
	v119 = v109 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v119
	v121 = F_repalloc(m, v110, v119)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L7
	} else {
		goto L27
	}
L26:
	;
	v131 = v124
	goto L24
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v121
	v124 = v121 + v103
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v124
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v126 <= v105 {
		v109 = v126
		v110 = v121
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v143 = F_strlen(m, v142)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v143 + v142
	F_infix_2(m, l0, int32(1))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L7
	} else {
		goto L30
	}
L30:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v151 = v149 - v150
	v153 = v151 + int32(3)
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v154 <= v153 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v157 = v154
	v158 = v150
	goto L34
L32:
	;
	v179 = v149
	goto L33
L33:
	;
	v188 = F_pg_sprintf(m, v179, int32(719653), int32(0))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L7
	} else {
		goto L38
	}
L34:
	;
	v167 = v157 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v167
	v169 = F_repalloc(m, v158, v167)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L7
	} else {
		goto L36
	}
L35:
	;
	v179 = v172
	goto L33
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v169
	v172 = v169 + v151
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v172
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v174 <= v153 {
		v157 = v174
		v158 = v169
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v191 = F_strlen(m, v190)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v191 + v190
	goto L1
L39:
	;
	v258 = v195
	goto L41
L40:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v204 = v202 - v203
	v206 = v204 + int32(3)
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v207 <= v206 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v258
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v261 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v261
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v260
	v265 = F_palloc(m, v261)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L7
	} else {
		goto L50
	}
L42:
	;
	v210 = v207
	v211 = v203
	goto L45
L43:
	;
	v233 = v202
	goto L44
L44:
	;
	v241 = F_pg_sprintf(m, v233, int32(783364), int32(0))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L7
	} else {
		goto L49
	}
L45:
	;
	v220 = v210 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v220
	v222 = F_repalloc(m, v211, v220)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L7
	} else {
		goto L47
	}
L46:
	;
	v233 = v225
	goto L44
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v222
	v225 = v222 + v204
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v225
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v227 <= v206 {
		v210 = v227
		v211 = v222
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v244 = F_strlen(m, v243)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v244 + v243
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v258 = v247
	goto L41
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v265
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v265
	F_infix_2(m, v13+int32(12), int32(0))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L7
	} else {
		goto L51
	}
L51:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v274
	F_infix_2(m, l0, int32(0))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L7
	} else {
		goto L52
	}
L52:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v281 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v282 = v280 - v281
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	if v279 <= v282+v283-v285+int32(4) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v293 = v279
	v294 = v281
	goto L56
L54:
	;
	v320 = v280
	v322 = v285
	goto L55
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v322
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v45
	v329 = F_pg_sprintf(m, v320, int32(208531), v13)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L7
	} else {
		goto L60
	}
L56:
	;
	v303 = v293 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v303
	v305 = F_repalloc(m, v294, v303)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L7
	} else {
		goto L58
	}
L57:
	;
	v320 = v308
	v322 = v313
	goto L55
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v305
	v308 = v305 + v282
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v308
	v310 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	if v310 <= v282+int32(4)+v311-v313 {
		v293 = v310
		v294 = v305
		goto L56
	} else {
		goto L59
	}
L59:
	;
	goto L57
L60:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v332 = F_strlen(m, v331)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v332 + v331
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	F_pfree(m, v335)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L7
	} else {
		goto L61
	}
L61:
	;
	if v199&int32(1) != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v342 = v340 - v341
	v344 = v342 + int32(3)
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v345 <= v344 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v348 = v345
	v349 = v341
	goto L66
L64:
	;
	v370 = v340
	goto L65
L65:
	;
	v379 = F_pg_sprintf(m, v370, int32(719653), int32(0))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L7
	} else {
		goto L70
	}
L66:
	;
	v358 = v348 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v358
	v360 = F_repalloc(m, v349, v358)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L7
	} else {
		goto L68
	}
L67:
	;
	v370 = v363
	goto L65
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v360
	v363 = v360 + v342
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v363
	v365 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v365 <= v344 {
		v348 = v365
		v349 = v360
		goto L66
	} else {
		goto L69
	}
L69:
	;
	goto L67
L70:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v382 = F_strlen(m, v381)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v382 + v381
	goto L1
L71:
	;
	v396 = v391 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v396
	v398 = F_repalloc(m, v389, v396)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L7
	} else {
		goto L73
	}
L72:
	;
	v412 = v404
	v413 = v401
	goto L2
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v398
	v401 = v398 + v37
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v401
	v403 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v405 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v404)+9)))
	if v403 <= v39+v405<<(uint(int32(1))%32) {
		v389 = v398
		v391 = v403
		goto L71
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	v422 = v33
	v424 = v413
	v425 = v420
	goto L78
L76:
	;
	v443 = v413
	v450 = v412
	goto L77
L77:
	;
	v451 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v450)+8)))
	if v451&int32(4) != 0 {
		goto L81
	} else {
		goto L82
	}
L78:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v424))) = uint8(v425)
	v432 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v433 = int32(1)
	v434 = v432 + v433
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v434
	v436 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v422)+1)))
	if v436 != 0 {
		v422 = v422 + v433
		v424 = v434
		v425 = v436
		goto L78
	} else {
		goto L80
	}
L79:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v443 = v434
	v450 = v439
	goto L77
L80:
	;
	goto L79
L81:
	;
	v454 = int32(37)
	*(*uint8)(unsafe.Add(mBase, uint32(v443))) = uint8(v454)
	v456 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v458 = v456 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v458
	v460 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v460)+8)))
	v462 = v461
	v463 = v458
	goto L83
L82:
	;
	v462 = v451
	v463 = v443
	goto L83
L83:
	;
	if v462&int32(2) != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v466 = int32(64)
	*(*uint8)(unsafe.Add(mBase, uint32(v463))) = uint8(v466)
	v468 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v470 = v468 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v470
	v472 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v473 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v472)+8)))
	v474 = v470
	v475 = v473
	goto L86
L85:
	;
	v474 = v463
	v475 = v462
	goto L86
L86:
	;
	if v475&int32(1) != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v478 = int32(42)
	*(*uint8)(unsafe.Add(mBase, uint32(v474))) = uint8(v478)
	v480 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v482 = v480 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v482
	v484 = v482
	goto L89
L88:
	;
	v484 = v474
	goto L89
L89:
	;
	v485 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v484))) = uint8(v485)
	v487 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v487 + int32(12)
	goto L1
}
func F_initial_cost_hashjoin(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 float64
	_ = v21
	var v22 float64
	_ = v22
	var v24 float64
	_ = v24
	var v25 int32
	_ = v25
	var v28 float64
	_ = v28
	var v29 float64
	_ = v29
	var v31 float64
	_ = v31
	var v32 float64
	_ = v32
	var v34 float64
	_ = v34
	var v37 float64
	_ = v37
	var v42 float64
	_ = v42
	var v44 int32
	_ = v44
	var v45 float64
	_ = v45
	var v47 int32
	_ = v47
	var v53 float64
	_ = v53
	var v57 float64
	_ = v57
	var v60 float64
	_ = v60
	var v63 float64
	_ = v63
	var v64 float64
	_ = v64
	var v65 float64
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v93 int32
	_ = v93
	var v97 float64
	_ = v97
	var v99 int32
	_ = v99
	var v103 float64
	_ = v103
	var v104 float64
	_ = v104
	var v107 float64
	_ = v107
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v122 float64
	_ = v122
	var v123 float64
	_ = v123
	var v126 float64
	_ = v126
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v137 float64
	_ = v137
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v158 float64
	_ = v158
	var v162 float64
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 float64
	_ = v172
	var v174 float64
	_ = v174
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v190 int32
	_ = v190
	var v200 float64
	_ = v200
	var v202 int32
	_ = v202
	var v206 float64
	_ = v206
	var v207 float64
	_ = v207
	var v210 float64
	_ = v210
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 float64
	_ = v247
	var v249 float64
	_ = v249
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v265 int32
	_ = v265
	var v273 float64
	_ = v273
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v308 int32
	_ = v308
	var v314 float64
	_ = v314
	var v316 float64
	_ = v316
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v377 int32
	_ = v377
	var v385 int32
	_ = v385
	var v397 int32
	_ = v397
	var v403 int32
	_ = v403
	var v412 int32
	_ = v412
	var v416 float64
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v427 float64
	_ = v427
	var v429 float64
	_ = v429
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v443 float64
	_ = v443
	var v449 float64
	_ = v449
	var v450 float64
	_ = v450
	var v465 int32
	_ = v465
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	v21 = *(*float64)(unsafe.Add(mBase, uint32(l3)+32))
	v22 = *(*float64)(unsafe.Add(mBase, uint32(l2)+32))
	v24 = *(*float64)(unsafe.Add(mBase, _consts[383]))
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v28 = base.F64_convert_i32_s(v25)
	goto L3
L2:
	;
	v28 = float64(0)
	goto L3
L3:
	;
	v29 = base.F64_mul(v24, v28)
	v31 = *(*float64)(unsafe.Add(mBase, uint32(l2)+56))
	v32 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
	v34 = float64(0)
	v37 = *(*float64)(unsafe.Add(mBase, _consts[384]))
	v42 = *(*float64)(unsafe.Add(mBase, uint32(l3)+56))
	if l4 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	v45 = base.F64_convert_i32_s(v44)
	v47 = int32(*(*uint8)(unsafe.Add(mBase, _consts[338])))
	if v47 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v63 = v21
	goto L6
L6:
	;
	v64 = base.F64_add(base.F64_mul(v29, v22), base.F64_add(base.F64_sub(v31, v32), v34))
	v65 = base.F64_add(base.F64_mul(base.F64_add(v29, v37), v21), base.F64_add(base.F64_add(v32, v34), v42))
	v67 = int32(*(*uint8)(unsafe.Add(mBase, _consts[387])))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l3)+40))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+32))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	v79 = v19 + int32(4)
	v93 = (v71 + int32(7)) & int32(-8)
	v97 = *(*float64)(unsafe.Add(mBase, _consts[337]))
	v99 = *(*int32)(unsafe.Add(mBase, _consts[21]))
	v103 = base.F64_mul(base.F64_mul(v97, base.F64_convert_i32_s(v99)), float64(1024))
	v104 = float64(4.294967295e+09)
	if base.F64_lt(v103, v104) != 0 {
		goto L16
	} else {
		goto L17
	}
L7:
	;
	v53 = base.F64_add(base.F64_mul(v45, float64(-0.3)), float64(1))
	if base.F64_gt(v53, float64(0)) != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v60 = v45
	goto L9
L9:
	;
	v63 = base.F64_mul(v21, v60)
	goto L6
L10:
	;
	v57 = v53
	goto L12
L11:
	;
	v57 = math.Float64frombits(uint64(0x8000000000000000))
	goto L12
L12:
	;
	v60 = base.F64_add(v57, v45)
	goto L9
L13:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	if int32(2) <= v412 {
		goto L123
	} else {
		goto L124
	}
L14:
	;
	if base.F64_le(v63, float64(0)) != 0 {
		goto L29
	} else {
		goto L30
	}
L15:
	;
	if l4 == int32(0) {
		v135 = v115
		goto L14
	} else {
		goto L22
	}
L16:
	;
	v107 = v103
	goto L18
L17:
	;
	v107 = v104
	goto L18
L18:
	;
	if base.F64_lt(v107, float64(4.294967296e+09))&base.F64_ge(v107, float64(0)) != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v113 = base.I32_trunc_f64_u(v107)
	v115 = v113
	goto L15
L20:
	;
	goto L21
L21:
	;
	v115 = int32(0)
	goto L15
L22:
	;
	v122 = base.F64_mul(base.F64_convert_i32_s(v73+int32(1)), base.F64_convert_i32_u(v115))
	v123 = float64(4.294967295e+09)
	if base.F64_lt(v122, v123) != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v126 = v122
	goto L25
L24:
	;
	v126 = v123
	goto L25
L25:
	;
	if base.F64_lt(v126, float64(4.294967296e+09))&base.F64_ge(v126, float64(0)) != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v132 = base.I32_trunc_f64_u(v126)
	v135 = v132
	goto L14
L27:
	;
	goto L28
L28:
	;
	v135 = int32(0)
	goto L14
L29:
	;
	v137 = float64(1000)
	goto L31
L30:
	;
	v137 = v63
	goto L31
L31:
	;
	v140 = v93 + int32(68)
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v135
	goto L32
L32:
	;
	v143 = base.I32_div_u_s(v135, v140)
	v144 = int32(50)
	v145 = base.I32_div_u_s(v143, v144)
	if base.Ui32(v144) <= base.Ui32(v143) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v79))) = v145
	v156 = int32(1)
	v158 = base.F64_mul(v137, base.F64_convert_i32_s(v93+int32(24)))
	v162 = base.F64_ceil(v137)
	v164 = int32(268435455)
	v166 = int32(base.Ui32(v151) >> (uint(int32(2)) % 32))
	if base.Ui32(v164) <= base.Ui32(v166) {
		goto L40
	} else {
		goto L41
	}
L35:
	;
	v150 = v145 * v140
	goto L37
L36:
	;
	v150 = int32(0)
	goto L37
L37:
	;
	v151 = v135 - v150
	goto L34
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19+int32(12)))) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v19+int32(8)))) = v403
	goto L13
L39:
	;
	if v180 <= int32(1024) {
		goto L49
	} else {
		goto L50
	}
L40:
	;
	v169 = v164
	goto L42
L41:
	;
	v169 = v166
	goto L42
L42:
	;
	v171 = int32(base.Ui32(int32(-2147483648)) >> (uint(base.I32_clz(v169)) % 32))
	v172 = base.F64_convert_i32_u(v171)
	if base.F64_gt(v172, v162) != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v174 = v162
	goto L45
L44:
	;
	v174 = v172
	goto L45
L45:
	;
	if base.F64_lt(base.F64_abs(v174), float64(2.147483648e+09)) != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v178 = base.I32_trunc_f64_s(v174)
	v180 = v178
	goto L39
L47:
	;
	goto L48
L48:
	;
	v180 = int32(-2147483648)
	goto L39
L49:
	;
	v183 = int32(1024)
	goto L51
L50:
	;
	v183 = v180
	goto L51
L51:
	;
	if v183&(v183-int32(1)) != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v190 = v156 << (uint(int32(32)-base.I32_clz(v183)) % 32)
	goto L54
L53:
	;
	v190 = v183
	goto L54
L54:
	;
	if base.F64_lt(base.F64_convert_i32_u(v151), base.F64_add(v158, base.F64_convert_i32_u(v190<<(uint(int32(2))%32)))) == int32(0) {
		v397 = v190
		v403 = v156
		goto L38
	} else {
		goto L55
	}
L55:
	;
	if l4 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v200 = *(*float64)(unsafe.Add(mBase, _consts[337]))
	v202 = *(*int32)(unsafe.Add(mBase, _consts[21]))
	v206 = base.F64_mul(base.F64_mul(v200, base.F64_convert_i32_s(v202)), float64(1024))
	v207 = float64(4.294967295e+09)
	if base.F64_lt(v206, v207) != 0 {
		goto L60
	} else {
		goto L61
	}
L57:
	;
	v273 = v172
	v274 = v151
	v278 = v171
	goto L58
L58:
	;
	v280 = int32(1)
	v285 = v93 + int32(28)
	if base.Ui32(v285) < base.Ui32(v274) {
		goto L91
	} else {
		goto L92
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v218
	goto L66
L60:
	;
	v210 = v206
	goto L62
L61:
	;
	v210 = v207
	goto L62
L62:
	;
	if base.F64_lt(v210, float64(4.294967296e+09))&base.F64_ge(v210, float64(0)) != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v216 = base.I32_trunc_f64_u(v210)
	v218 = v216
	goto L59
L64:
	;
	goto L65
L65:
	;
	v218 = int32(0)
	goto L59
L66:
	;
	v220 = base.I32_div_u_s(v218, v140)
	v221 = int32(50)
	v222 = base.I32_div_u_s(v220, v221)
	if base.Ui32(v221) <= base.Ui32(v220) {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v79))) = v222
	v233 = int32(1)
	v239 = int32(268435455)
	v241 = int32(base.Ui32(v228) >> (uint(int32(2)) % 32))
	if base.Ui32(v239) <= base.Ui32(v241) {
		goto L73
	} else {
		goto L74
	}
L69:
	;
	v227 = v222 * v140
	goto L71
L70:
	;
	v227 = int32(0)
	goto L71
L71:
	;
	v228 = v218 - v227
	goto L68
L72:
	;
	if v255 <= int32(1024) {
		goto L82
	} else {
		goto L83
	}
L73:
	;
	v244 = v239
	goto L75
L74:
	;
	v244 = v241
	goto L75
L75:
	;
	v246 = int32(base.Ui32(int32(-2147483648)) >> (uint(base.I32_clz(v244)) % 32))
	v247 = base.F64_convert_i32_u(v246)
	if base.F64_gt(v247, v162) != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v249 = v162
	goto L78
L77:
	;
	v249 = v247
	goto L78
L78:
	;
	if base.F64_lt(base.F64_abs(v249), float64(2.147483648e+09)) != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v253 = base.I32_trunc_f64_s(v249)
	v255 = v253
	goto L72
L80:
	;
	goto L81
L81:
	;
	v255 = int32(-2147483648)
	goto L72
L82:
	;
	v258 = int32(1024)
	goto L84
L83:
	;
	v258 = v255
	goto L84
L84:
	;
	if v258&(v258-int32(1)) != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v265 = v233 << (uint(int32(32)-base.I32_clz(v258)) % 32)
	goto L87
L86:
	;
	v265 = v258
	goto L87
L87:
	;
	if base.F64_lt(base.F64_convert_i32_u(v228), base.F64_add(v158, base.F64_convert_i32_u(v265<<(uint(int32(2))%32)))) == int32(0) {
		v397 = v265
		v403 = v233
		goto L38
	} else {
		goto L88
	}
L88:
	;
	v273 = v247
	v274 = v228
	v278 = v246
	goto L58
L89:
	;
	v397 = v385
	v403 = v377
	goto L38
L90:
	;
	if v322 <= int32(2) {
		goto L109
	} else {
		goto L110
	}
L91:
	;
	v287 = int32(1)
	v289 = base.I32_div_u_s(v274, v285)
	if v289&(v289-v287) != 0 {
		goto L94
	} else {
		goto L95
	}
L92:
	;
	v297 = v280
	goto L93
L93:
	;
	if base.Ui32(v297) < base.Ui32(v278) {
		goto L97
	} else {
		goto L98
	}
L94:
	;
	v296 = v287 << (uint(int32(32)-base.I32_clz(v289)) % 32)
	goto L96
L95:
	;
	v296 = v289
	goto L96
L96:
	;
	v297 = v296
	goto L93
L97:
	;
	v301 = v297
	goto L99
L98:
	;
	v301 = v278
	goto L99
L99:
	;
	if v301&(v301-int32(1)) != 0 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v308 = int32(1) << (uint(int32(32)-base.I32_clz(v301)) % 32)
	goto L102
L101:
	;
	v308 = v301
	goto L102
L102:
	;
	v314 = base.F64_ceil(base.F64_div(v158, base.F64_convert_i32_u(v274-v308<<(uint(int32(2))%32))))
	if base.F64_gt(v273, v314) != 0 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v316 = v314
	goto L105
L104:
	;
	v316 = v273
	goto L105
L105:
	;
	if base.F64_lt(base.F64_abs(v316), float64(2.147483648e+09)) != 0 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v320 = base.I32_trunc_f64_s(v316)
	v322 = v320
	goto L90
L107:
	;
	goto L108
L108:
	;
	v322 = int32(-2147483648)
	goto L90
L109:
	;
	v325 = int32(2)
	goto L111
L110:
	;
	v325 = v322
	goto L111
L111:
	;
	if v325&(v325-int32(1)) != 0 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v332 = v280 << (uint(int32(32)-base.I32_clz(v325)) % 32)
	goto L114
L113:
	;
	v332 = v325
	goto L114
L114:
	;
	if v332 < int32(2) {
		v377 = v332
		v385 = v308
		goto L89
	} else {
		goto L115
	}
L115:
	;
	if base.Ui32(int32(134217727)) < base.Ui32(v308) {
		v377 = v332
		v385 = v308
		goto L89
	} else {
		goto L116
	}
L116:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v339 = v332
	v347 = v308
	v349 = v337
	goto L117
L117:
	;
	if v349 < int32(0) {
		v377 = v339
		v385 = v347
		goto L89
	} else {
		goto L119
	}
L118:
	;
	v397 = v371
	v403 = v369
	goto L38
L119:
	;
	if base.Ui32(v339) < base.Ui32(int32(base.Ui32(v349)>>(uint(int32(13))%32))) {
		v377 = v339
		v385 = v347
		goto L89
	} else {
		goto L120
	}
L120:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	v361 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v79))) = v360 << (uint(v361) % 32)
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v366 = v364 << (uint(v361) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v366
	v369 = int32(base.Ui32(v339) >> (uint(v361) % 32))
	v371 = v347 << (uint(v361) % 32)
	if base.Ui32(v339) < base.Ui32(int32(4)) {
		v397 = v371
		v403 = v369
		goto L38
	} else {
		goto L121
	}
L121:
	;
	if base.Ui32(v347) < base.Ui32(int32(67108864)) {
		v339 = v369
		v347 = v371
		v349 = v366
		goto L117
	} else {
		goto L122
	}
L122:
	;
	goto L118
L123:
	;
	v416 = *(*float64)(unsafe.Add(mBase, _consts[385]))
	v417 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v417)+32))
	v419 = int32(7)
	v421 = int32(-8)
	v423 = int32(24)
	v427 = float64(0.0001220703125)
	v429 = base.F64_ceil(base.F64_mul(base.F64_mul(v22, base.F64_convert_i32_u((v418+v419)&v421+v423)), v427))
	v431 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v431)+32))
	v443 = base.F64_ceil(base.F64_mul(base.F64_mul(v21, base.F64_convert_i32_u((v432+v419)&v421+v423)), v427))
	v449 = base.F64_add(base.F64_mul(v416, base.F64_add(base.F64_add(v429, v429), v443)), v64)
	v450 = base.F64_add(base.F64_mul(v416, v443), v65)
	goto L125
L124:
	;
	v449 = v64
	v450 = v65
	goto L125
L125:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+24)) = v449
	*(*float64)(unsafe.Add(mBase, uint32(l0)+8)) = v450
	*(*float64)(unsafe.Add(mBase, uint32(l0)+16)) = base.F64_add(v449, v450)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v69 + (v67^int32(1))&int32(255) + v68
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+88)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v412
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v465
	m.G0 = v19 + int32(16)
	return
}
func F_initscan(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int64
	_ = v138
	v4 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v6 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+118)))
	if v16 == int32(116) {
		v39 = v4
		goto L8
	} else {
		goto L9
	}
L2:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
	v12 = v7
	goto L1
L3:
	;
	goto L4
L4:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v10 = F_RelationGetNumberOfBlocksInFork(m, v8, int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return
L6:
	;
	v12 = v10
	goto L1
L7:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v47 != 0 {
		goto L19
	} else {
		goto L20
	}
L8:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	if v40 != 0 {
		goto L14
	} else {
		goto L15
	}
L9:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _consts[35]))
	v22 = base.I32_div_s(v20, int32(4))
	if base.Ui32(v12) <= base.Ui32(v22) {
		v39 = v4
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v28 = int32(base.Ui32(v24&int32(128)) >> (uint(int32(7)) % 32))
	if v24&int32(64) == int32(0) {
		v39 = v28
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	if v33 != 0 {
		v46 = v28
		goto L7
	} else {
		goto L12
	}
L12:
	;
	v35 = F_GetAccessStrategy(m, int32(1))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v35
	v46 = v28
	goto L7
L14:
	;
	F_bms_free(m, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L5
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = int32(0)
	v46 = v39
	goto L7
L17:
	;
	goto L16
L18:
	;
	v94 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v94
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)) = uint8(v94)
	v98 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v98
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+72)) = uint16(v94)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v98
	*(*int64)(unsafe.Add(mBase, uint32(l0)+52)) = int64(4294967295)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+100)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+88)) = int64(-4294967295)
	if l1 == v94 {
		goto L35
	} else {
		goto L36
	}
L19:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+12)))
	if v49 == int32(1) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	if l2 != 0 {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v48 | int32(128)
	goto L18
L23:
	;
	goto L24
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v48 & int32(-129)
	goto L18
L25:
	;
	if v46 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	if v46 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v68 & int32(-129)
	goto L18
L29:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, _consts[34])))
	if v61 != int32(1) {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v64 | int32(128)
	goto L18
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(0)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v89 & int32(-129)
	goto L18
L32:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, _consts[34])))
	if v75 != int32(1) {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v78 | int32(128)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v84 = F_ss_get_location(m, v82, v83)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L5
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v84
	goto L18
L35:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v121&int32(1) == int32(0) {
		goto L42
	} else {
		goto L43
	}
L36:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v112 <= int32(0) {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v117 = v112 * int32(48)
	if v117 != 0 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	goto L35
L39:
	;
	v118 = F__emscripten_memcpy_bulkmem(m, v115, l1, v117)
	mBase = m.M
	goto L41
L40:
	;
	goto L41
L41:
	;
	goto L38
L42:
	;
	return
L43:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+272))
	if v127 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+268)))
	if v130 != int32(1) {
		goto L42
	} else {
		goto L47
	}
L45:
	;
	v137 = v127
	goto L46
L46:
	;
	v138 = *(*int64)(unsafe.Add(mBase, uint32(v137)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v137)+16)) = v138 + int64(1)
	goto L42
L47:
	;
	F_pgstat_assoc_relation(m, v126)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L5
	} else {
		goto L48
	}
L48:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)+272))
	v137 = v136
	goto L46
}
func F_inittapes(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v15 int64
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
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
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int64
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int64
	_ = v63
	var v65 int64
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int64
	_ = v79
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	if l1 != 0 {
		v12 = int32(6)
		v13 = *(*int64)(unsafe.Add(mBase, uint32(l0)+96))
		v15 = base.I64_div_s(v13, int64(278528))
		v16 = base.I32_wrap_i64(v15)
		if v16 <= v12 {
			v19 = v12
		} else {
			v19 = v16
		}
		if int32(500) <= v19 {
			v22 = int32(500)
		} else {
			v22 = v19
		}
		v23 = v22
	} else {
		v23 = int32(6)
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v23
	v26 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1107])))
	if v26 != int32(1) {
		v54 = v23
		v57 = base.I64_extend_i32_s(v54) << (uint(int64(13)) % 64)
		v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
		v59 = F_GetMemoryChunkSpace(m, v58)
		mBase = m.M
		v60 = m.ExcPending
		if v60 != 0 {
			return
		} else {
			v63 = *(*int64)(unsafe.Add(mBase, uint32(l0)+96))
			if v57+base.I64_extend_i32_u(v59) < v63 {
				v65 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+88)) = v65 - v57
			} else {
			}
			F_PrepareTempTablespaces(m)
			mBase = m.M
			v69 = m.ExcPending
			if v69 != 0 {
				return
			} else {
				v70 = int32(0)
				v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+236))
				if v71 != 0 {
					v75 = v71 + int32(12)
				} else {
					v75 = v70
				}
				v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
				v77 = F_LogicalTapeSetCreate(m, v70, v75, v76)
				mBase = m.M
				v78 = m.ExcPending
				if v78 != 0 {
					return
				} else {
					v79 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(l0)+168)) = v79
					*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v77
					*(*int64)(unsafe.Add(mBase, uint32(l0)+176)) = v79
					v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
					v87 = F_palloc0(m, v84<<(uint(int32(2))%32))
					mBase = m.M
					v88 = m.ExcPending
					if v88 != 0 {
						return
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(l0)+188)) = int64(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = v87
						*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = int32(2)
						v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
						v95 = F_LogicalTapeCreate(m, v94)
						mBase = m.M
						v96 = m.ExcPending
						if v96 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+196)) = v95
							v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
							v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
							*(*int32)(unsafe.Add(mBase, uint32(v98+v99<<(uint(int32(2))%32)))) = v95
							v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
							v105 = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = v104 + v105
							v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = v108 + v105
							m.G0 = v8 + int32(16)
							return
						}
					}
				}
			}
		}
	} else {
		v31 = F_errstart(m, int32(15), int32(0))
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return
		} else {
			v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
			if v31 == int32(0) {
				v54 = v33
				v57 = base.I64_extend_i32_s(v54) << (uint(int64(13)) % 64)
				v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
				v59 = F_GetMemoryChunkSpace(m, v58)
				mBase = m.M
				v60 = m.ExcPending
				if v60 != 0 {
					return
				} else {
					v63 = *(*int64)(unsafe.Add(mBase, uint32(l0)+96))
					if v57+base.I64_extend_i32_u(v59) < v63 {
						v65 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
						*(*int64)(unsafe.Add(mBase, uint32(l0)+88)) = v65 - v57
					} else {
					}
					F_PrepareTempTablespaces(m)
					mBase = m.M
					v69 = m.ExcPending
					if v69 != 0 {
						return
					} else {
						v70 = int32(0)
						v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+236))
						if v71 != 0 {
							v75 = v71 + int32(12)
						} else {
							v75 = v70
						}
						v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
						v77 = F_LogicalTapeSetCreate(m, v70, v75, v76)
						mBase = m.M
						v78 = m.ExcPending
						if v78 != 0 {
							return
						} else {
							v79 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(l0)+168)) = v79
							*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v77
							*(*int64)(unsafe.Add(mBase, uint32(l0)+176)) = v79
							v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
							v87 = F_palloc0(m, v84<<(uint(int32(2))%32))
							mBase = m.M
							v88 = m.ExcPending
							if v88 != 0 {
								return
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(l0)+188)) = int64(0)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = v87
								*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = int32(2)
								v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
								v95 = F_LogicalTapeCreate(m, v94)
								mBase = m.M
								v96 = m.ExcPending
								if v96 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+196)) = v95
									v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
									v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
									*(*int32)(unsafe.Add(mBase, uint32(v98+v99<<(uint(int32(2))%32)))) = v95
									v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
									v105 = int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = v104 + v105
									v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = v108 + v105
									m.G0 = v8 + int32(16)
									return
								}
							}
						}
					}
				}
			} else {
				v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
				v39 = F_pg_rusage_show(m, l0+int32(256))
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v39
					*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v33
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v36
					F_errmsg_internal(m, int32(211633), v8)
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return
					} else {
						F_errfinish(m, int32(517689), int32(1883), int32(173265))
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return
						} else {
							v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
							v54 = v52
							v57 = base.I64_extend_i32_s(v54) << (uint(int64(13)) % 64)
							v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
							v59 = F_GetMemoryChunkSpace(m, v58)
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return
							} else {
								v63 = *(*int64)(unsafe.Add(mBase, uint32(l0)+96))
								if v57+base.I64_extend_i32_u(v59) < v63 {
									v65 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
									*(*int64)(unsafe.Add(mBase, uint32(l0)+88)) = v65 - v57
								} else {
								}
								F_PrepareTempTablespaces(m)
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return
								} else {
									v70 = int32(0)
									v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+236))
									if v71 != 0 {
										v75 = v71 + int32(12)
									} else {
										v75 = v70
									}
									v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
									v77 = F_LogicalTapeSetCreate(m, v70, v75, v76)
									mBase = m.M
									v78 = m.ExcPending
									if v78 != 0 {
										return
									} else {
										v79 = int64(0)
										*(*int64)(unsafe.Add(mBase, uint32(l0)+168)) = v79
										*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v77
										*(*int64)(unsafe.Add(mBase, uint32(l0)+176)) = v79
										v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
										v87 = F_palloc0(m, v84<<(uint(int32(2))%32))
										mBase = m.M
										v88 = m.ExcPending
										if v88 != 0 {
											return
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(l0)+188)) = int64(0)
											*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = v87
											*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = int32(2)
											v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
											v95 = F_LogicalTapeCreate(m, v94)
											mBase = m.M
											v96 = m.ExcPending
											if v96 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+196)) = v95
												v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
												v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
												*(*int32)(unsafe.Add(mBase, uint32(v98+v99<<(uint(int32(2))%32)))) = v95
												v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
												v105 = int32(1)
												*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = v104 + v105
												v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = v108 + v105
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
	}
}
func F_insertSelectOptions(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
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
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v179 int32
	_ = v179
	var v190 int32
	_ = v190
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	if l1 != 0 {
		goto L7
	} else {
		goto L8
	}
L1:
	;
	if l4 != 0 {
		goto L61
	} else {
		goto L62
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v179
	goto L1
L3:
	;
	v179 = v28
	goto L2
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L11
	} else {
		goto L55
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L11
	} else {
		goto L50
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L11
	} else {
		goto L45
	}
L7:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v14 != 0 {
		goto L6
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v17 = F_list_concat(m, v16, l2)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = l1
	goto L9
L11:
	;
	return
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v17
	if l3 == int32(0) {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v22 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v23 != 0 {
		goto L5
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v25 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v22
	goto L16
L18:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v26 != 0 {
		goto L4
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v29 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v25
	goto L20
L22:
	;
	if v28 != int32(1) {
		v179 = v28
		goto L2
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	if v28 != int32(1) {
		goto L3
	} else {
		goto L31
	}
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L11
	} else {
		goto L26
	}
L26:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L11
	} else {
		goto L27
	}
L27:
	;
	F_errmsg(m, int32(377424), int32(0))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L11
	} else {
		goto L28
	}
L28:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	F_scanner_errposition(m, v45, l5)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L11
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(27597), int32(19075), int32(147554))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L11
	} else {
		goto L30
	}
L30:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L31:
	;
	if v17 == int32(0) {
		goto L3
	} else {
		goto L32
	}
L32:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v58 <= int32(0) {
		v179 = int32(1)
		goto L2
	} else {
		goto L33
	}
L33:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v64 = int32(0)
	goto L34
L34:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v61+v64<<(uint(int32(2))%32))))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
	if v76 != int32(1) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L11
	} else {
		goto L40
	}
L36:
	;
	v80 = v64 + int32(1)
	if v58 != v80 {
		v64 = v80
		goto L34
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	goto L35
L39:
	;
	goto L3
L40:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L11
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = int32(550896)
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(571913)
	F_errmsg(m, int32(234550), v12)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L11
	} else {
		goto L42
	}
L42:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	F_scanner_errposition(m, v96, l5)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L11
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(27597), int32(19089), int32(147554))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L11
	} else {
		goto L44
	}
L44:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L45:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L11
	} else {
		goto L46
	}
L46:
	;
	F_errmsg(m, int32(461721), int32(0))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L11
	} else {
		goto L47
	}
L47:
	;
	v115 = F_exprLocation(m, l1)
	mBase = m.M
	F_scanner_errposition(m, v115, l5)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L11
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(27597), int32(19044), int32(147554))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L11
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L50:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L11
	} else {
		goto L51
	}
L51:
	;
	F_errmsg(m, int32(461794), int32(0))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L11
	} else {
		goto L52
	}
L52:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	F_scanner_errposition(m, v134, l5)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L11
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(27597), int32(19055), int32(147554))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L11
	} else {
		goto L54
	}
L54:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L55:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L11
	} else {
		goto L56
	}
L56:
	;
	F_errmsg(m, int32(461759), int32(0))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L11
	} else {
		goto L57
	}
L57:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	F_scanner_errposition(m, v153, l5)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L11
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(27597), int32(19064), int32(147554))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L11
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
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L11
	} else {
		goto L65
	}
L61:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v190 != 0 {
		goto L60
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	m.G0 = v12 + int32(16)
	return
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = l4
	goto L63
L65:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L11
	} else {
		goto L66
	}
L66:
	;
	F_errmsg(m, int32(461830), int32(0))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L11
	} else {
		goto L67
	}
L67:
	;
	v206 = F_exprLocation(m, l4)
	mBase = m.M
	F_scanner_errposition(m, v206, l5)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L11
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(27597), int32(19100), int32(147554))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L11
	} else {
		goto L69
	}
L69:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_int24eq(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+20)))
	return base.B2i32(v2 == v3)
}
func F_int28(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v2 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+20)))
	v3 = F_Int64GetDatum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_int28eq(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 int64
	_ = v4
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v4 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+20)))
	return base.B2i32(v3 == v4)
}
func F_int2eqfast(m *base.Module, l0 int32, l1 int32) int32 {
	var v3 int32
	_ = v3
	v3 = int32(65535)
	return base.B2i32(l0&v3 == l1&v3)
}
func F_int2ge(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+20)))
	v3 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+28)))
	return base.B2i32(v3 <= v2)
}
func F_int2out(m *base.Module, l0 int32) int32 {
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
	v2 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+20)))
	v4 = F_palloc(m, int32(7))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
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
		return v4
	}
}
func F_int2pl(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	v3 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+20)))
	v4 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+28)))
	v5 = v3 + v4
	v6 = base.I32_extend16_s(v5)
	if v6 != v5 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(422476), int32(0))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(518199), int32(944), int32(316640))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
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
		return v6
	}
}
func F_int42eq(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+28)))
	return base.B2i32(v2 == v3)
}
func F_int48(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v2 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+20)))
	v3 = F_Int64GetDatum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_int48eq(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 int64
	_ = v4
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v4 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+20)))
	return base.B2i32(v3 == v4)
}
func F_int48ge(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 int64
	_ = v4
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v4 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+20)))
	return base.B2i32(v3 <= v4)
}
func F_int4mi(m *base.Module, l0 int32) int32 {
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
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = v6 - v3
	if base.B2i32(int32(0) < v3) != base.B2i32(v7 < v6) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(422822), int32(0))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(518199), int32(843), int32(336763))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
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
		return v7
	}
}
func F_int4mul(m *base.Module, l0 int32) int32 {
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
	v3 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+20)))
	v4 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+28)))
	v5 = v3 * v4
	v9 = base.I32_wrap_i64(v5)
	if base.I32_wrap_i64(int64(base.Ui64(v5)>>(uint(int64(32))%64))) != v9>>(uint(int32(31))%32) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(422822), int32(0))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(518199), int32(857), int32(315181))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
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
		return v9
	}
}
func F_int4um(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v2 == int32(-2147483648) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(422822), int32(0))
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(518199), int32(807), int32(302192))
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
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
		return int32(0) - v2
	}
}
func F_int82(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	if base.Ui64(v4-int64(32768)) <= base.Ui64(int64(-65537)) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(422476), int32(0))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(527193), int32(1277), int32(589748))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
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
		return base.I32_wrap_i64(v4)
	}
}
func F_int82mi(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	var v9 int64
	_ = v9
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	v4 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+28)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
	v9 = v8 - v4
	if base.B2i32(int64(0) < v4) != base.B2i32(v9 < v8) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(422498), int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(527193), int32(1055), int32(336812))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
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
		v30 = F_Int64GetDatum(m, v9)
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			return v30
		}
	}
}
func F_int84ge(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 int64
	_ = v4
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v4 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+28)))
	return base.B2i32(v4 <= v3)
}
func F_int84mul(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v12 int64
	_ = v12
	var v19 int64
	_ = v19
	var v20 int64
	_ = v20
	var v22 int64
	_ = v22
	var v25 int64
	_ = v25
	var v26 int64
	_ = v26
	var v28 int64
	_ = v28
	var v29 int64
	_ = v29
	var v33 int64
	_ = v33
	var v40 int64
	_ = v40
	var v51 int64
	_ = v51
	var v52 int64
	_ = v52
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
	v10 = int64(63)
	v12 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+28)))
	v19 = int64(32)
	v20 = int64(base.Ui64(v12) >> (uint(v19) % 64))
	v22 = int64(base.Ui64(v9) >> (uint(v19) % 64))
	v25 = int64(4294967295)
	v26 = v12 & v25
	v28 = v9 & v25
	v29 = v26 * v28
	v33 = int64(base.Ui64(v29)>>(uint(v19)%64)) + v26*v22
	v40 = v28*v20 + v33&v25
	*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = v9*(v12>>(uint(v10)%64)) + v9>>(uint(v10)%64)*v12 + v20*v22 + int64(base.Ui64(v33)>>(uint(v19)%64)) + int64(base.Ui64(v40)>>(uint(v19)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v6))) = v29&v25 | v40<<(uint(v19)%64)
	v51 = *(*int64)(unsafe.Add(mBase, uint32(v6)+8))
	v52 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
	if v51 != v52>>(uint(int64(63))%64) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v61 = m.ExcPending
		if v61 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v64 = m.ExcPending
			if v64 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(422498), int32(0))
				mBase = m.M
				v68 = m.ExcPending
				if v68 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(527193), int32(927), int32(315199))
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
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
		v74 = F_Int64GetDatum(m, v52)
		mBase = m.M
		v75 = m.ExcPending
		if v75 != 0 {
			return int32(0)
		} else {
			m.G0 = v6 + int32(16)
			return v74
		}
	}
}
func F_int84pl(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	var v9 int64
	_ = v9
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	v4 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+28)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
	v9 = v8 + v4
	if base.B2i32(v4 < int64(0)) != base.B2i32(v9 < v8) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(422498), int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(527193), int32(899), int32(316614))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
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
		v30 = F_Int64GetDatum(m, v9)
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			return v30
		}
	}
}
func F_int8gcd(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v14 int64
	_ = v14
	var v17 int32
	_ = v17
	var v18 int64
	_ = v18
	var v19 int64
	_ = v19
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v37 int64
	_ = v37
	var __phi37 int64
	_ = __phi37
	var v39 int64
	_ = v39
	var __phi39 int64
	_ = __phi39
	var v40 int64
	_ = v40
	var v45 int64
	_ = v45
	var v48 int64
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
	v9 = int64(63)
	v10 = v6 >> (uint(v9) % 64)
	v14 = v8 >> (uint(v9) % 64)
	v17 = base.B2i32(v14-(v14^v8) < v10-(v10^v6))
	if v14-(v14^v8) < v10-(v10^v6) {
		v18 = v6
	} else {
		v18 = v8
	}
	if v14-(v14^v8) < v10-(v10^v6) {
		v19 = v8
	} else {
		v19 = v6
	}
	if v19 == int64(-9223372036854775807-1) {
		if v18&int64(9223372036854775807) == int64(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v57 = m.ExcPending
			if v57 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50331778))
				mBase = m.M
				v60 = m.ExcPending
				if v60 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(422498), int32(0))
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(527193), int32(636), int32(328542))
						mBase = m.M
						v69 = m.ExcPending
						if v69 != 0 {
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
			if v18 != int64(-1) {
				__phi37 = v18
				__phi39 = v19
				v37 = __phi37
				v39 = __phi39
				for {
					v40 = base.I64_rem_s(v39, v37)
					if v40 != int64(0) {
						__phi37 = v40
						__phi39 = v37
						v37 = __phi37
						v39 = __phi39
						continue
					} else {
						break
					}
					break
				}
				v45 = v37
				v48 = v45 >> (uint(int64(63)) % 64)
				v51 = F_Int64GetDatum(m, v45^v48-v48)
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return int32(0)
				} else {
					return v51
				}
			} else {
				v29 = F_Int64GetDatum(m, int64(1))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					return v29
				}
			}
		}
	} else {
		if v18 != int64(0) {
			__phi37 = v18
			__phi39 = v19
			v37 = __phi37
			v39 = __phi39
			for {
				v40 = base.I64_rem_s(v39, v37)
				if v40 != int64(0) {
					__phi37 = v40
					__phi39 = v37
					v37 = __phi37
					v39 = __phi39
					continue
				} else {
					break
				}
				break
			}
			v45 = v37
		} else {
			v45 = v19
		}
		v48 = v45 >> (uint(int64(63)) % 64)
		v51 = F_Int64GetDatum(m, v45^v48-v48)
		mBase = m.M
		v52 = m.ExcPending
		if v52 != 0 {
			return int32(0)
		} else {
			return v51
		}
	}
}
func F_int8inc_support(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
	if v4 == int32(462) {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v3)+8))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
		if v8 != 0 {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
			v21 = int32(base.Ui32(v9)>>(uint(int32(5))%32))&int32(1) | int32(base.Ui32(v9)>>(uint(int32(7))%32))&int32(2)
		} else {
			v21 = int32(3)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v3)+12)) = v21
		v24 = v3
	} else {
		v24 = int32(0)
	}
	return v24
}
func F_int8mul(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v20 int64
	_ = v20
	var v21 int64
	_ = v21
	var v23 int64
	_ = v23
	var v26 int64
	_ = v26
	var v27 int64
	_ = v27
	var v29 int64
	_ = v29
	var v30 int64
	_ = v30
	var v34 int64
	_ = v34
	var v41 int64
	_ = v41
	var v52 int64
	_ = v52
	var v53 int64
	_ = v53
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
	v10 = int64(63)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v13 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
	v20 = int64(32)
	v21 = int64(base.Ui64(v13) >> (uint(v20) % 64))
	v23 = int64(base.Ui64(v9) >> (uint(v20) % 64))
	v26 = int64(4294967295)
	v27 = v13 & v26
	v29 = v9 & v26
	v30 = v27 * v29
	v34 = int64(base.Ui64(v30)>>(uint(v20)%64)) + v27*v23
	v41 = v29*v21 + v34&v26
	*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = v9*(v13>>(uint(v10)%64)) + v9>>(uint(v10)%64)*v13 + v21*v23 + int64(base.Ui64(v34)>>(uint(v20)%64)) + int64(base.Ui64(v41)>>(uint(v20)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v6))) = v30&v26 | v41<<(uint(v20)%64)
	v52 = *(*int64)(unsafe.Add(mBase, uint32(v6)+8))
	v53 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
	if v52 != v53>>(uint(int64(63))%64) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v62 = m.ExcPending
		if v62 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v65 = m.ExcPending
			if v65 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(422498), int32(0))
				mBase = m.M
				v69 = m.ExcPending
				if v69 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(527193), int32(499), int32(315134))
					mBase = m.M
					v74 = m.ExcPending
					if v74 != 0 {
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
		v75 = F_Int64GetDatum(m, v53)
		mBase = m.M
		v76 = m.ExcPending
		if v76 != 0 {
			return int32(0)
		} else {
			m.G0 = v6 + int32(16)
			return v75
		}
	}
}
func F_int8out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	var v13 int32
	_ = v13
	var v18 int64
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	v3 = m.G0
	v5 = v3 - int32(32)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
	if int64(0) <= v8 {
		v18 = v8
		v19 = int32(0)
	} else {
		v13 = int32(45)
		*(*uint8)(unsafe.Add(mBase, uint32(v5))) = uint8(v13)
		v18 = int64(0) - v8
		v19 = int32(1)
	}
	v21 = F_pg_ulltoa_n(m, v18, v5+v19)
	mBase = m.M
	v22 = v21 + v19
	v24 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v5+v22))) = uint8(v24)
	v27 = v22 + int32(1)
	v28 = F_palloc(m, v27)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		return int32(0)
	} else {
		if v27 != 0 {
			v32 = F__emscripten_memcpy_bulkmem(m, v28, v5, v27)
			mBase = m.M
			v33 = v32
		} else {
			v33 = v28
		}
		m.G0 = v5 + int32(32)
		return v33
	}
}
func F_intervaltypmodin(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v16 = F_ArrayGetIntegerTypmods(m, v10, v7+int32(28))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
	if v18 <= int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	switch v18 - int32(1) {
	case 0:
		goto L33
	case 1:
		goto L32
	default:
		goto L31
	}
L5:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v21 <= int32(3071) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L24
	}
L7:
	;
	switch v21 - int32(1024) {
	case 0, 8:
		goto L4
	case 1, 2, 3, 4, 5, 6, 7:
		goto L6
	default:
		goto L22
	}
L8:
	;
	if int32(1)<<(uint(v21)%32)&int32(340) != 0 {
		goto L4
	} else {
		goto L21
	}
L9:
	;
	if base.Ui32(v21) <= base.Ui32(int32(8)) {
		goto L8
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	if v21 <= int32(6143) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L7
L13:
	;
	switch v21 - int32(3072) {
	case 0, 8:
		goto L4
	case 1, 2, 3, 4, 5, 6, 7:
		goto L6
	default:
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	switch v21 - int32(7168) {
	case 0, 8:
		goto L4
	case 1, 2, 3, 4, 5, 6, 7:
		goto L6
	default:
		goto L18
	}
L16:
	;
	if v21 != int32(4096) {
		goto L6
	} else {
		goto L17
	}
L17:
	;
	goto L4
L18:
	;
	if v21 == int32(6144) {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	if v21 == int32(32767) {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	goto L6
L21:
	;
	goto L7
L22:
	;
	if v21 == int32(2048) {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	goto L6
L24:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	F_errmsg(m, int32(234404), int32(0))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(521665), int32(1085), int32(291317))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L48
	}
L29:
	;
	m.G0 = v7 + int32(32)
	return v134
L30:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v134 = v127<<(uint(int32(16))%32)&int32(2147418112) | v77
	goto L29
L31:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L44
	}
L32:
	;
	v76 = v16 + int32(4)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	if v77 < int32(0) {
		goto L28
	} else {
		goto L35
	}
L33:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v66 == int32(32767) {
		v134 = int32(-1)
		goto L29
	} else {
		goto L34
	}
L34:
	;
	v134 = v66<<(uint(int32(16))%32)&int32(2147418112) | int32(65535)
	goto L29
L35:
	;
	if base.Ui32(v77) < base.Ui32(int32(7)) {
		goto L30
	} else {
		goto L36
	}
L36:
	;
	v84 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	if v84 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v134 = v104<<(uint(int32(16))%32)&int32(2147418112) | int32(6)
	goto L29
L41:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = int32(6)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v89
	F_errmsg(m, int32(512495), v7+int32(16))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(521665), int32(1108), int32(291317))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	goto L40
L44:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	F_errmsg(m, int32(234404), int32(0))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(521665), int32(1118), int32(291317))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L48:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v147
	F_errmsg(m, int32(361986), v7)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(521665), int32(1102), int32(291317))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_intorel_receive(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+32)))
	if v5 == int32(0) {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v8)+188))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+80))
		m.T0[v13].(func(*base.Module, int32, int32, int32, int32, int32))(m, v8, l0, v9, v10, v11)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			return int32(1)
		}
	} else {
		return int32(1)
	}
}
func F_intorel_startup(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
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
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
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
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	v4 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(32)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v20 = v19
	goto L3
L2:
	;
	v20 = v4
	goto L3
L3:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if int32(0) < v22 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L19
	} else {
		goto L52
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L19
	} else {
		goto L48
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L19
	} else {
		goto L42
	}
L7:
	;
	v28 = v22
	v30 = v20
	v32 = v4
	v34 = v4
	goto L10
L8:
	;
	v86 = v20
	v88 = v4
	goto L9
L9:
	;
	if v86 != 0 {
		goto L5
	} else {
		goto L28
	}
L10:
	;
	v44 = l2 + int32(20) + v28<<(uint(int32(4))%32) + v34*int32(100)
	if v30 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v86 = v64
	v88 = v77
	goto L9
L12:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v44)+68))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v44)+76))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v44)+96))
	v68 = F_makeColumnDef(m, v62, v65, v66, v67)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L19
	} else {
		goto L20
	}
L13:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	v48 = v30 + int32(4)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	if base.Ui32(v48) < base.Ui32(v51+v52<<(uint(int32(2))%32)) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	v62 = v44 + int32(4)
	v64 = int32(0)
	goto L12
L16:
	;
	v57 = v48
	goto L18
L17:
	;
	v57 = int32(0)
	goto L18
L18:
	;
	v62 = v46
	v64 = v57
	goto L12
L19:
	;
	return
L20:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v68)+52))
	if v70 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v68)+8))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+8))
	v75 = F_type_is_collatable(m, v74)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L19
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v77 = F_lappend(m, v32, v68)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L19
	} else {
		goto L26
	}
L24:
	;
	if v75 != 0 {
		goto L6
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v80 = v34 + int32(1)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v80 < v81 {
		v28 = v81
		v30 = v64
		v32 = v77
		v34 = v80
		goto L10
	} else {
		goto L27
	}
L27:
	;
	goto L11
L28:
	;
	F_create_ctas_internal(m, v15+int32(20), v88, v17)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L19
	} else {
		goto L29
	}
L29:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
	v103 = F_table_open(m, v101, int32(8))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L19
	} else {
		goto L30
	}
L30:
	;
	v105 = int32(0)
	v107 = F_check_enable_rls(m, v101, v105, v105)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L19
	} else {
		goto L31
	}
L31:
	;
	if v107 == int32(2) {
		goto L4
	} else {
		goto L32
	}
L32:
	;
	if v21 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v100
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v103
	v122 = F_GetCurrentCommandId(m, int32(1))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L19
	} else {
		goto L37
	}
L34:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+32)))
	if v113 != 0 {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	F_SetMatViewPopulatedState(m, v103, int32(1))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L19
	} else {
		goto L36
	}
L36:
	;
	goto L33
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v122
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+32)))
	if v128 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v131 = int32(0)
	goto L40
L39:
	;
	v129 = F_GetBulkInsertState(m)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L19
	} else {
		goto L41
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v131
	m.G0 = v15 + int32(32)
	return
L41:
	;
	v131 = v129
	goto L40
L42:
	;
	F_errcode(m, int32(34209924))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L19
	} else {
		goto L43
	}
L43:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v68)+8))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)+8))
	v146 = F_format_type_be(m, v145)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L19
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v143
	F_errmsg(m, int32(203969), v15)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L19
	} else {
		goto L45
	}
L45:
	;
	F_errhint(m, int32(604209), int32(0))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L19
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(520686), int32(515), int32(244150))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L19
	} else {
		goto L47
	}
L47:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L48:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L19
	} else {
		goto L49
	}
L49:
	;
	F_errmsg(m, int32(480614), int32(0))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L19
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(520686), int32(523), int32(244150))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L19
	} else {
		goto L51
	}
L51:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L52:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L19
	} else {
		goto L53
	}
L53:
	;
	F_errmsg(m, int32(450016), int32(0))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L19
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(520686), int32(546), int32(244150))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L19
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_intset(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_int_to_intset(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_intset_union_elem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
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
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v15 = F_intarray_add_elem(m, v10, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v17 != v10 {
				F_pfree(m, v10)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
					v24 = F_ArrayGetNItems(m, v21, v15+int32(16))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						v26 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)) = uint8(v26)
						v28 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
						if v28 != 0 {
							v36 = v28
						} else {
							v29 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
							v36 = (v29<<(uint(int32(3))%32) + int32(23)) & int32(-8)
						}
						F_isort(m, v36+v15, v24, v7+int32(15))
						mBase = m.M
						v41 = F__int_unique(m, v15)
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							m.G0 = v7 + int32(16)
							return v41
						}
					}
				}
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
				v24 = F_ArrayGetNItems(m, v21, v15+int32(16))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v26 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)) = uint8(v26)
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
					if v28 != 0 {
						v36 = v28
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
						v36 = (v29<<(uint(int32(3))%32) + int32(23)) & int32(-8)
					}
					F_isort(m, v36+v15, v24, v7+int32(15))
					mBase = m.M
					v41 = F__int_unique(m, v15)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						m.G0 = v7 + int32(16)
						return v41
					}
				}
			}
		}
	}
}
func F_isQueryUsingTempRelation_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	v3 = int32(0)
	if l0 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v12 == int32(67) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	return v77
L5:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v15 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v70 = F_expression_tree_walker_impl(m, l0, int32(497), l1)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L16
	} else {
		goto L22
	}
L8:
	;
	v66 = F_query_tree_walker_impl(m, l0, int32(497), l1, int32(4))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L16
	} else {
		goto L21
	}
L9:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v18 <= int32(0) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v24 = v3
	goto L11
L11:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28+v24<<(uint(int32(2))%32))))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	if v33 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L8
L13:
	;
	v36 = int32(1)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	v39 = F_table_open(m, v37, v36)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	v54 = v24 + int32(1)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v54 < v55 {
		v24 = v54
		goto L11
	} else {
		goto L20
	}
L16:
	;
	return int32(0)
L17:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+118)))
	F_sequence_close(m, v39, int32(1))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	if v44 == int32(116) {
		v77 = v36
		goto L4
	} else {
		goto L19
	}
L19:
	;
	goto L15
L20:
	;
	goto L12
L21:
	;
	return v66
L22:
	;
	v77 = v70
	goto L4
}
func F_is_complex_array(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	v2 = F_get_element_type(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		if v2 == int32(0) {
			return int32(0)
		} else {
			v10 = F_typeOrDomainTypeRelid(m, v2)
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				return base.B2i32(v10 != int32(0))
			}
		}
	}
}
func F_is_objectclass_supported(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	v5 = int32(0)
	goto L1
L1:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v5*int32(40))+uint32(_consts[252])))
	v12 = base.B2i32(v11 == l0)
	if v12 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	return v12
L3:
	;
	v16 = v5 + int32(1)
	if v16 != int32(37) {
		v5 = v16
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	goto L2
L6:
	;
	goto L5
}
func F_isalnum(m *base.Module, l0 int32) int32 {
	return base.B2i32(base.Ui32(l0-int32(48)) < base.Ui32(int32(10))) | base.B2i32(base.Ui32(l0|int32(32)-int32(97)) < base.Ui32(int32(26)))
}
func F_isdigit(m *base.Module, l0 int32) int32 {
	return base.B2i32(base.Ui32(l0-int32(48)) < base.Ui32(int32(10)))
}
func F_islower(m *base.Module, l0 int32) int32 {
	return base.B2i32(base.Ui32(l0-int32(97)) < base.Ui32(int32(26)))
}
func F_ismn_cast_from_ean13(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
	F_ean2isn(m, v8, v5+int32(8), int32(4))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int64)(unsafe.Add(mBase, uint32(v5)+8))
		v17 = F_Int64GetDatum(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			m.G0 = v5 + int32(16)
			return v17
		}
	}
}
func F_isort(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v59 int32
	_ = v59
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v93 int32
	_ = v93
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v403 int32
	_ = v403
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v456 int32
	_ = v456
	var v463 int32
	_ = v463
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v511 int32
	_ = v511
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v606 int32
	_ = v606
	var v610 int32
	_ = v610
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
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
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v661 int32
	_ = v661
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v741 int32
	_ = v741
	var v744 int32
	_ = v744
	v19 = l0
	v20 = l1
	goto L1
L1:
	;
	v38 = v19 + int32(4)
	v40 = v20
	goto L3
L2:
	;
	return
L3:
	;
	v59 = v19 + v40<<(uint(int32(2))%32)
	if base.Ui32(v40) <= base.Ui32(int32(6)) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L2
L5:
	;
	goto L4
L6:
	;
	if base.Ui32(v40) < base.Ui32(int32(2)) {
		goto L5
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	v136 = v134 & int32(1)
	v142 = v38
	goto L24
L9:
	;
	v79 = v38
	goto L10
L10:
	;
	if base.Ui32(v79) <= base.Ui32(v19) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L5
L12:
	;
	v132 = v79 + int32(4)
	if base.Ui32(v132) < base.Ui32(v59) {
		v79 = v132
		goto L10
	} else {
		goto L23
	}
L13:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	v93 = v79
	goto L14
L14:
	;
	v106 = v93 - int32(4)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
	if v84&int32(1) != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L12
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93))) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v106))) = v83
	if base.Ui32(v19) < base.Ui32(v106) {
		v93 = v106
		goto L14
	} else {
		goto L22
	}
L17:
	;
	if v83 < v107 {
		goto L16
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	if v83 <= v107 {
		goto L12
	} else {
		goto L21
	}
L20:
	;
	goto L12
L21:
	;
	goto L16
L22:
	;
	goto L15
L23:
	;
	goto L11
L24:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v142-int32(4))))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	if v136 != 0 {
		goto L28
	} else {
		goto L29
	}
L25:
	;
	v168 = v19 + v40<<(uint(int32(1))%32)&int32(-4)
	if v40 != int32(7) {
		goto L34
	} else {
		goto L35
	}
L26:
	;
	goto L25
L27:
	;
	v162 = v142 + int32(4)
	if base.Ui32(v162) < base.Ui32(v59) {
		v142 = v162
		goto L24
	} else {
		goto L33
	}
L28:
	;
	if v157 <= v158 {
		goto L27
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	if v157 < v158 {
		goto L26
	} else {
		goto L32
	}
L31:
	;
	goto L26
L32:
	;
	goto L27
L33:
	;
	goto L5
L34:
	;
	v172 = v59 - int32(4)
	if base.Ui32(v40) < base.Ui32(int32(41)) {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v298 = v168
	goto L36
L36:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v298)))
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v302
	*(*int32)(unsafe.Add(mBase, uint32(v298))) = v301
	v306 = v59 - int32(4)
	v312 = v38
	v313 = v306
	v315 = v38
	v317 = v306
	goto L141
L37:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v268)))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v267)))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v266)))
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v276 == int32(1) {
		goto L120
	} else {
		goto L121
	}
L38:
	;
	v266 = v19
	v267 = v168
	v268 = v172
	goto L37
L39:
	;
	goto L40
L40:
	;
	v175 = int32(1)
	v178 = int32(base.Ui32(v40)>>(uint(v175)%32)) & int32(2147483644)
	v179 = v19 + v178
	v181 = v40 & int32(-8)
	v182 = v19 + v181
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v182)))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v179)))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v189 == v175 {
		goto L45
	} else {
		goto L46
	}
L41:
	;
	v209 = v168 - v178
	v210 = v178 + v168
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v168)))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v217 == int32(1) {
		goto L70
	} else {
		goto L71
	}
L42:
	;
	v208 = v206
	goto L41
L43:
	;
	if v186 < v187 {
		v206 = v179
		goto L42
	} else {
		goto L62
	}
L44:
	;
	if v186 < v187 {
		v206 = v179
		goto L42
	} else {
		goto L58
	}
L45:
	;
	if v187 <= v188 {
		goto L44
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	if v187 < v188 {
		goto L43
	} else {
		goto L53
	}
L48:
	;
	if v187 < v186 {
		v206 = v179
		goto L42
	} else {
		goto L49
	}
L49:
	;
	if v188 < v186 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v195 = v182
	goto L52
L51:
	;
	v195 = v19
	goto L52
L52:
	;
	v208 = v195
	goto L41
L53:
	;
	if v187 < v186 {
		v206 = v179
		goto L42
	} else {
		goto L54
	}
L54:
	;
	if v186 < v188 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v199 = v19
	goto L57
L56:
	;
	v199 = v182
	goto L57
L57:
	;
	v208 = v199
	goto L41
L58:
	;
	if v188 < v186 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v202 = v19
	goto L61
L60:
	;
	v202 = v182
	goto L61
L61:
	;
	v206 = v202
	goto L42
L62:
	;
	if v186 < v188 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v205 = v182
	goto L65
L64:
	;
	v205 = v19
	goto L65
L65:
	;
	v208 = v205
	goto L41
L66:
	;
	v237 = v172 - v181
	v238 = v172 - v178
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v172)))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v238)))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v237)))
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v245 == int32(1) {
		goto L95
	} else {
		goto L96
	}
L67:
	;
	v236 = v234
	goto L66
L68:
	;
	if v214 < v215 {
		v234 = v168
		goto L67
	} else {
		goto L87
	}
L69:
	;
	if v214 < v215 {
		v234 = v168
		goto L67
	} else {
		goto L83
	}
L70:
	;
	if v215 <= v216 {
		goto L69
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	if v215 < v216 {
		goto L68
	} else {
		goto L78
	}
L73:
	;
	if v215 < v214 {
		v234 = v168
		goto L67
	} else {
		goto L74
	}
L74:
	;
	if v216 < v214 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v223 = v210
	goto L77
L76:
	;
	v223 = v209
	goto L77
L77:
	;
	v236 = v223
	goto L66
L78:
	;
	if v215 < v214 {
		v234 = v168
		goto L67
	} else {
		goto L79
	}
L79:
	;
	if v214 < v216 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v227 = v209
	goto L82
L81:
	;
	v227 = v210
	goto L82
L82:
	;
	v236 = v227
	goto L66
L83:
	;
	if v216 < v214 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v230 = v209
	goto L86
L85:
	;
	v230 = v210
	goto L86
L86:
	;
	v234 = v230
	goto L67
L87:
	;
	if v214 < v216 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v233 = v210
	goto L90
L89:
	;
	v233 = v209
	goto L90
L90:
	;
	v236 = v233
	goto L66
L91:
	;
	v266 = v208
	v267 = v236
	v268 = v264
	goto L37
L92:
	;
	v264 = v262
	goto L91
L93:
	;
	if v242 < v243 {
		v262 = v238
		goto L92
	} else {
		goto L112
	}
L94:
	;
	if v242 < v243 {
		v262 = v238
		goto L92
	} else {
		goto L108
	}
L95:
	;
	if v243 <= v244 {
		goto L94
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	if v243 < v244 {
		goto L93
	} else {
		goto L103
	}
L98:
	;
	if v243 < v242 {
		v262 = v238
		goto L92
	} else {
		goto L99
	}
L99:
	;
	if v244 < v242 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v251 = v172
	goto L102
L101:
	;
	v251 = v237
	goto L102
L102:
	;
	v264 = v251
	goto L91
L103:
	;
	if v243 < v242 {
		v262 = v238
		goto L92
	} else {
		goto L104
	}
L104:
	;
	if v242 < v244 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v255 = v237
	goto L107
L106:
	;
	v255 = v172
	goto L107
L107:
	;
	v264 = v255
	goto L91
L108:
	;
	if v244 < v242 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v258 = v237
	goto L111
L110:
	;
	v258 = v172
	goto L111
L111:
	;
	v262 = v258
	goto L92
L112:
	;
	if v242 < v244 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v261 = v172
	goto L115
L114:
	;
	v261 = v237
	goto L115
L115:
	;
	v264 = v261
	goto L91
L116:
	;
	v298 = v295
	goto L36
L117:
	;
	v295 = v293
	goto L116
L118:
	;
	if v273 < v274 {
		v293 = v267
		goto L117
	} else {
		goto L137
	}
L119:
	;
	if v273 < v274 {
		v293 = v267
		goto L117
	} else {
		goto L133
	}
L120:
	;
	if v274 <= v275 {
		goto L119
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	if v274 < v275 {
		goto L118
	} else {
		goto L128
	}
L123:
	;
	if v274 < v273 {
		v293 = v267
		goto L117
	} else {
		goto L124
	}
L124:
	;
	if v275 < v273 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v282 = v268
	goto L127
L126:
	;
	v282 = v266
	goto L127
L127:
	;
	v295 = v282
	goto L116
L128:
	;
	if v274 < v273 {
		v293 = v267
		goto L117
	} else {
		goto L129
	}
L129:
	;
	if v273 < v275 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v286 = v266
	goto L132
L131:
	;
	v286 = v268
	goto L132
L132:
	;
	v295 = v286
	goto L116
L133:
	;
	if v275 < v273 {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v289 = v266
	goto L136
L135:
	;
	v289 = v268
	goto L136
L136:
	;
	v293 = v289
	goto L117
L137:
	;
	if v273 < v275 {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v292 = v268
	goto L140
L139:
	;
	v292 = v266
	goto L140
L140:
	;
	v295 = v292
	goto L116
L141:
	;
	if base.Ui32(v313) < base.Ui32(v312) {
		v365 = v312
		v368 = v315
		goto L143
	} else {
		goto L144
	}
L143:
	;
	if base.Ui32(v365) <= base.Ui32(v313) {
		goto L158
	} else {
		goto L159
	}
L144:
	;
	v331 = v312
	v334 = v315
	goto L145
L145:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v331)))
	if v136 != 0 {
		goto L149
	} else {
		goto L150
	}
L146:
	;
	v365 = v358
	v368 = v356
	goto L143
L147:
	;
	v358 = v331 + int32(4)
	if base.Ui32(v358) <= base.Ui32(v313) {
		v331 = v358
		v334 = v356
		goto L145
	} else {
		goto L156
	}
L148:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v334)))
	*(*int32)(unsafe.Add(mBase, uint32(v334))) = v345
	*(*int32)(unsafe.Add(mBase, uint32(v331))) = v350
	v356 = v334 + int32(4)
	goto L147
L149:
	;
	if v345 < v344 {
		v356 = v334
		goto L147
	} else {
		goto L152
	}
L150:
	;
	goto L151
L151:
	;
	if v344 < v345 {
		v356 = v334
		goto L147
	} else {
		goto L154
	}
L152:
	;
	if v345 <= v344 {
		goto L148
	} else {
		goto L153
	}
L153:
	;
	v365 = v331
	v368 = v334
	goto L143
L154:
	;
	if v345 < v344 {
		v365 = v331
		v368 = v334
		goto L143
	} else {
		goto L155
	}
L155:
	;
	goto L148
L156:
	;
	goto L146
L157:
	;
	v741 = *(*int32)(unsafe.Add(mBase, uint32(v365)))
	*(*int32)(unsafe.Add(mBase, uint32(v365))) = v398
	*(*int32)(unsafe.Add(mBase, uint32(v385))) = v741
	v744 = int32(4)
	v312 = v365 + v744
	v313 = v385 - v744
	v315 = v368
	v317 = v389
	goto L141
L158:
	;
	v385 = v313
	v389 = v317
	goto L161
L159:
	;
	v418 = v313
	v422 = v317
	goto L160
L160:
	;
	v431 = int32(2)
	v432 = (v368 - v19) >> (uint(v431) % 32)
	v435 = (v365 - v368) >> (uint(v431) % 32)
	if v432 < v435 {
		goto L174
	} else {
		goto L175
	}
L161:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v385)))
	if v136 != 0 {
		goto L165
	} else {
		goto L166
	}
L162:
	;
	v418 = v410
	v422 = v408
	goto L160
L163:
	;
	v410 = v385 - int32(4)
	if base.Ui32(v365) <= base.Ui32(v410) {
		v385 = v410
		v389 = v408
		goto L161
	} else {
		goto L172
	}
L164:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v389)))
	*(*int32)(unsafe.Add(mBase, uint32(v385))) = v403
	*(*int32)(unsafe.Add(mBase, uint32(v389))) = v398
	v408 = v389 - int32(4)
	goto L163
L165:
	;
	if v398 < v397 {
		goto L157
	} else {
		goto L168
	}
L166:
	;
	goto L167
L167:
	;
	if v397 < v398 {
		goto L157
	} else {
		goto L170
	}
L168:
	;
	if v398 <= v397 {
		goto L164
	} else {
		goto L169
	}
L169:
	;
	v408 = v389
	goto L163
L170:
	;
	if v398 < v397 {
		v408 = v389
		goto L163
	} else {
		goto L171
	}
L171:
	;
	goto L164
L172:
	;
	goto L162
L173:
	;
	v577 = int32(2)
	v578 = (v422 - v418) >> (uint(v577) % 32)
	v583 = (v59-v422)>>(uint(v577)%32) - int32(1)
	if v578 < v583 {
		goto L189
	} else {
		goto L190
	}
L174:
	;
	v437 = v432
	goto L176
L175:
	;
	v437 = v435
	goto L176
L176:
	;
	if v437 == int32(0) {
		goto L173
	} else {
		goto L177
	}
L177:
	;
	v442 = v365 - v437<<(uint(int32(2))%32)
	v444 = v437 & int32(3)
	v445 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v437) {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v456 = v445
	v463 = int32(0)
	goto L181
L179:
	;
	v511 = v445
	goto L180
L180:
	;
	if v444 == int32(0) {
		goto L173
	} else {
		goto L184
	}
L181:
	;
	v471 = v456 << (uint(int32(2)) % 32)
	v472 = v19 + v471
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v472)))
	v474 = v442 + v471
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v474)))
	*(*int32)(unsafe.Add(mBase, uint32(v472))) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v474))) = v473
	v478 = int32(4)
	v479 = v471 | v478
	v480 = v19 + v479
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v480)))
	v482 = v442 + v479
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v482)))
	*(*int32)(unsafe.Add(mBase, uint32(v480))) = v483
	*(*int32)(unsafe.Add(mBase, uint32(v482))) = v481
	v487 = v471 | int32(8)
	v488 = v19 + v487
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v488)))
	v490 = v442 + v487
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v490)))
	*(*int32)(unsafe.Add(mBase, uint32(v488))) = v491
	*(*int32)(unsafe.Add(mBase, uint32(v490))) = v489
	v495 = v471 | int32(12)
	v496 = v19 + v495
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v496)))
	v498 = v442 + v495
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v498)))
	*(*int32)(unsafe.Add(mBase, uint32(v496))) = v499
	*(*int32)(unsafe.Add(mBase, uint32(v498))) = v497
	v503 = v456 + v478
	v505 = v463 + v478
	if v505 != v437&int32(-4) {
		v456 = v503
		v463 = v505
		goto L181
	} else {
		goto L183
	}
L182:
	;
	v511 = v503
	goto L180
L183:
	;
	goto L182
L184:
	;
	v531 = v511
	v535 = v445
	goto L185
L185:
	;
	v546 = v531 << (uint(int32(2)) % 32)
	v547 = v19 + v546
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v547)))
	v549 = v442 + v546
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v549)))
	*(*int32)(unsafe.Add(mBase, uint32(v547))) = v550
	*(*int32)(unsafe.Add(mBase, uint32(v549))) = v548
	v553 = int32(1)
	v556 = v535 + v553
	if v556 != v444 {
		v531 = v531 + v553
		v535 = v556
		goto L185
	} else {
		goto L187
	}
L186:
	;
	goto L173
L187:
	;
	goto L186
L188:
	;
	if base.Ui32(v435) <= base.Ui32(v578) {
		goto L203
	} else {
		goto L204
	}
L189:
	;
	v585 = v578
	goto L191
L190:
	;
	v585 = v583
	goto L191
L191:
	;
	if v585 == int32(0) {
		goto L188
	} else {
		goto L192
	}
L192:
	;
	v590 = v59 - v585<<(uint(int32(2))%32)
	v592 = v585 & int32(3)
	v593 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v585) {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v606 = v593
	v610 = int32(0)
	goto L196
L194:
	;
	v661 = v593
	goto L195
L195:
	;
	if v592 == int32(0) {
		goto L188
	} else {
		goto L199
	}
L196:
	;
	v619 = v606 << (uint(int32(2)) % 32)
	v620 = v365 + v619
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v620)))
	v622 = v619 + v590
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v622)))
	*(*int32)(unsafe.Add(mBase, uint32(v620))) = v623
	*(*int32)(unsafe.Add(mBase, uint32(v622))) = v621
	v626 = int32(4)
	v627 = v619 | v626
	v628 = v365 + v627
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v628)))
	v630 = v590 + v627
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v630)))
	*(*int32)(unsafe.Add(mBase, uint32(v628))) = v631
	*(*int32)(unsafe.Add(mBase, uint32(v630))) = v629
	v635 = v619 | int32(8)
	v636 = v365 + v635
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v636)))
	v638 = v590 + v635
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v638)))
	*(*int32)(unsafe.Add(mBase, uint32(v636))) = v639
	*(*int32)(unsafe.Add(mBase, uint32(v638))) = v637
	v643 = v619 | int32(12)
	v644 = v365 + v643
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v644)))
	v646 = v643 + v590
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v646)))
	*(*int32)(unsafe.Add(mBase, uint32(v644))) = v647
	*(*int32)(unsafe.Add(mBase, uint32(v646))) = v645
	v651 = v606 + v626
	v653 = v610 + v626
	if v653 != v585&int32(-4) {
		v606 = v651
		v610 = v653
		goto L196
	} else {
		goto L198
	}
L197:
	;
	v661 = v651
	goto L195
L198:
	;
	goto L197
L199:
	;
	v681 = v661
	v682 = v593
	goto L200
L200:
	;
	v694 = v681 << (uint(int32(2)) % 32)
	v695 = v365 + v694
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v695)))
	v697 = v694 + v590
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v697)))
	*(*int32)(unsafe.Add(mBase, uint32(v695))) = v698
	*(*int32)(unsafe.Add(mBase, uint32(v697))) = v696
	v701 = int32(1)
	v704 = v682 + v701
	if v704 != v592 {
		v681 = v681 + v701
		v682 = v704
		goto L200
	} else {
		goto L202
	}
L201:
	;
	goto L188
L202:
	;
	goto L201
L203:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v435) {
		goto L206
	} else {
		goto L207
	}
L204:
	;
	goto L205
L205:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v578) {
		goto L210
	} else {
		goto L211
	}
L206:
	;
	F_isort(m, v19, v435, l2)
	mBase = m.M
	goto L208
L207:
	;
	goto L208
L208:
	;
	if base.Ui32(v578) < base.Ui32(int32(2)) {
		goto L5
	} else {
		goto L209
	}
L209:
	;
	v19 = v59 - v578<<(uint(int32(2))%32)
	v20 = v578
	goto L1
L210:
	;
	F_isort(m, v59-v578<<(uint(int32(2))%32), v578, l2)
	mBase = m.M
	goto L212
L211:
	;
	goto L212
L212:
	;
	if base.Ui32(int32(1)) < base.Ui32(v435) {
		v40 = v435
		goto L3
	} else {
		goto L213
	}
L213:
	;
	goto L5
}
func F_isort_med3(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v11 == int32(1) {
		if v9 <= v10 {
			if v8 < v9 {
				v31 = l1
			} else {
				if v10 < v8 {
					v26 = l0
				} else {
					v26 = l2
				}
				v31 = v26
			}
			return v31
		} else {
			if v9 < v8 {
				v31 = l1
				return v31
			} else {
				if v10 < v8 {
					v17 = l2
				} else {
					v17 = l0
				}
				return v17
			}
		}
	} else {
		if v9 < v10 {
			if v8 < v9 {
				v31 = l1
				return v31
			} else {
				if v8 < v10 {
					v29 = l2
				} else {
					v29 = l0
				}
				return v29
			}
		} else {
			if v9 < v8 {
				v31 = l1
				return v31
			} else {
				if v8 < v10 {
					v22 = l0
				} else {
					v22 = l2
				}
				return v22
			}
		}
	}
}
func F_isprint(m *base.Module, l0 int32) int32 {
	return base.B2i32(base.Ui32(l0-int32(32)) < base.Ui32(int32(95)))
}
func F_issn_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v12 = F_string2ean(m, v7, v8, v5+int32(8), int32(5))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if v12 == int32(0) {
			v18 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v18)
			v24 = int32(0)
			m.G0 = v5 + int32(16)
			return v24
		} else {
			v21 = *(*int64)(unsafe.Add(mBase, uint32(v5)+8))
			v22 = F_Int64GetDatum(m, v21)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v24 = v22
				m.G0 = v5 + int32(16)
				return v24
			}
		}
	}
}
func F_iswspace(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v39 int32
	_ = v39
	if l0 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	if l0 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	return base.B2i32(v39 != int32(0))
L5:
	;
	v12 = int32(4134704)
	goto L8
L6:
	;
	goto L7
L7:
	;
	v22 = int32(4134704)
	goto L17
L8:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	if v14 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	if v14 != 0 {
		goto L14
	} else {
		goto L15
	}
L10:
	;
	if l0 != v14 {
		v12 = v12 + int32(4)
		goto L8
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	goto L9
L13:
	;
	goto L12
L14:
	;
	v20 = v12
	goto L16
L15:
	;
	v20 = int32(0)
	goto L16
L16:
	;
	v39 = v20
	goto L4
L17:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	if v28 != 0 {
		v22 = v22 + int32(4)
		goto L17
	} else {
		goto L19
	}
L18:
	;
	v29 = int32(4134704)
	v39 = (v22-v29)&int32(-4) + v29
	goto L4
L19:
	;
	goto L18
}
func F_italian_ISO_8859_1_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v279 int32
	_ = v279
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v348 int32
	_ = v348
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v457 int32
	_ = v457
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v506 int32
	_ = v506
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v543 int32
	_ = v543
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v559 int32
	_ = v559
	var v568 int32
	_ = v568
	var v577 int32
	_ = v577
	var v580 int32
	_ = v580
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v617 int32
	_ = v617
	var v630 int32
	_ = v630
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v663 int32
	_ = v663
	var v667 int32
	_ = v667
	var v680 int32
	_ = v680
	var v683 int32
	_ = v683
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v701 int32
	_ = v701
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v716 int32
	_ = v716
	var v724 int32
	_ = v724
	var v732 int32
	_ = v732
	var v735 int32
	_ = v735
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v765 int32
	_ = v765
	var v767 int32
	_ = v767
	var v773 int32
	_ = v773
	var v787 int32
	_ = v787
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v813 int32
	_ = v813
	var v815 int32
	_ = v815
	var v822 int32
	_ = v822
	var v826 int32
	_ = v826
	var v828 int32
	_ = v828
	var v830 int32
	_ = v830
	var v833 int32
	_ = v833
	var v837 int32
	_ = v837
	var v845 int32
	_ = v845
	var v853 int32
	_ = v853
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v868 int32
	_ = v868
	var v870 int32
	_ = v870
	var v877 int32
	_ = v877
	var v883 int32
	_ = v883
	var v885 int32
	_ = v885
	var v887 int32
	_ = v887
	var v893 int32
	_ = v893
	var v902 int32
	_ = v902
	var v911 int32
	_ = v911
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v917 int32
	_ = v917
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v929 int32
	_ = v929
	var v936 int32
	_ = v936
	var v940 int32
	_ = v940
	var v942 int32
	_ = v942
	var v944 int32
	_ = v944
	var v947 int32
	_ = v947
	var v951 int32
	_ = v951
	var v959 int32
	_ = v959
	var v967 int32
	_ = v967
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v982 int32
	_ = v982
	var v984 int32
	_ = v984
	var v991 int32
	_ = v991
	var v997 int32
	_ = v997
	var v999 int32
	_ = v999
	var v1001 int32
	_ = v1001
	var v1007 int32
	_ = v1007
	var v1016 int32
	_ = v1016
	var v1025 int32
	_ = v1025
	var v1028 int32
	_ = v1028
	var v1029 int32
	_ = v1029
	var v1034 int32
	_ = v1034
	var v1038 int32
	_ = v1038
	var v1040 int32
	_ = v1040
	var v1042 int32
	_ = v1042
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1059 int32
	_ = v1059
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1065 int32
	_ = v1065
	var v1067 int32
	_ = v1067
	var v1072 int32
	_ = v1072
	var v1073 int32
	_ = v1073
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1093 int32
	_ = v1093
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1102 int32
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1121 int32
	_ = v1121
	var v1123 int32
	_ = v1123
	var v1125 int32
	_ = v1125
	var v1128 int32
	_ = v1128
	var v1131 int32
	_ = v1131
	var v1134 int32
	_ = v1134
	var v1138 int32
	_ = v1138
	var v1141 int32
	_ = v1141
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1155 int32
	_ = v1155
	var v1156 int32
	_ = v1156
	var v1159 int32
	_ = v1159
	var v1160 int32
	_ = v1160
	var v1164 int32
	_ = v1164
	var v1165 int32
	_ = v1165
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1177 int32
	_ = v1177
	var v1178 int32
	_ = v1178
	var v1180 int32
	_ = v1180
	var v1181 int32
	_ = v1181
	var v1184 int32
	_ = v1184
	var v1185 int32
	_ = v1185
	var v1187 int32
	_ = v1187
	var v1188 int32
	_ = v1188
	var v1191 int32
	_ = v1191
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1197 int32
	_ = v1197
	var v1199 int32
	_ = v1199
	var v1212 int32
	_ = v1212
	var v1213 int32
	_ = v1213
	var v1216 int32
	_ = v1216
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1221 int32
	_ = v1221
	var v1222 int32
	_ = v1222
	var v1227 int32
	_ = v1227
	var v1229 int32
	_ = v1229
	var v1231 int32
	_ = v1231
	var v1234 int32
	_ = v1234
	var v1237 int32
	_ = v1237
	var v1240 int32
	_ = v1240
	var v1244 int32
	_ = v1244
	var v1247 int32
	_ = v1247
	var v1249 int32
	_ = v1249
	var v1250 int32
	_ = v1250
	var v1252 int32
	_ = v1252
	var v1253 int32
	_ = v1253
	var v1256 int32
	_ = v1256
	var v1257 int32
	_ = v1257
	var v1259 int32
	_ = v1259
	var v1260 int32
	_ = v1260
	var v1263 int32
	_ = v1263
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1269 int32
	_ = v1269
	var v1271 int32
	_ = v1271
	var v1284 int32
	_ = v1284
	var v1285 int32
	_ = v1285
	var v1288 int32
	_ = v1288
	var v1290 int32
	_ = v1290
	var v1291 int32
	_ = v1291
	var v1293 int32
	_ = v1293
	var v1294 int32
	_ = v1294
	var v1297 int32
	_ = v1297
	var v1298 int32
	_ = v1298
	var v1300 int32
	_ = v1300
	var v1301 int32
	_ = v1301
	var v1304 int32
	_ = v1304
	var v1306 int32
	_ = v1306
	var v1308 int32
	_ = v1308
	var v1311 int32
	_ = v1311
	var v1314 int32
	_ = v1314
	var v1317 int32
	_ = v1317
	var v1321 int32
	_ = v1321
	var v1324 int32
	_ = v1324
	var v1326 int32
	_ = v1326
	var v1327 int32
	_ = v1327
	var v1329 int32
	_ = v1329
	var v1330 int32
	_ = v1330
	var v1333 int32
	_ = v1333
	var v1335 int32
	_ = v1335
	var v1337 int32
	_ = v1337
	var v1340 int32
	_ = v1340
	var v1343 int32
	_ = v1343
	var v1346 int32
	_ = v1346
	var v1350 int32
	_ = v1350
	var v1353 int32
	_ = v1353
	var v1355 int32
	_ = v1355
	var v1356 int32
	_ = v1356
	var v1358 int32
	_ = v1358
	var v1359 int32
	_ = v1359
	var v1364 int32
	_ = v1364
	var v1365 int32
	_ = v1365
	var v1368 int32
	_ = v1368
	var v1370 int32
	_ = v1370
	var v1372 int32
	_ = v1372
	var v1375 int32
	_ = v1375
	var v1379 int32
	_ = v1379
	var v1380 int32
	_ = v1380
	var v1381 int32
	_ = v1381
	var v1383 int32
	_ = v1383
	var v1384 int32
	_ = v1384
	var v1392 int32
	_ = v1392
	var v1404 int32
	_ = v1404
	var v1416 int32
	_ = v1416
	var v1417 int32
	_ = v1417
	var v1421 int32
	_ = v1421
	var v1423 int32
	_ = v1423
	var v1429 int32
	_ = v1429
	var v1443 int32
	_ = v1443
	var v1447 int32
	_ = v1447
	var v1448 int32
	_ = v1448
	var v1450 int32
	_ = v1450
	var v1451 int32
	_ = v1451
	var v1453 int32
	_ = v1453
	var v1454 int32
	_ = v1454
	var v1457 int32
	_ = v1457
	var v1459 int32
	_ = v1459
	var v1461 int32
	_ = v1461
	var v1465 int32
	_ = v1465
	var v1469 int32
	_ = v1469
	var v1472 int32
	_ = v1472
	var v1473 int32
	_ = v1473
	var v1475 int32
	_ = v1475
	var v1476 int32
	_ = v1476
	var v1479 int32
	_ = v1479
	var v1480 int32
	_ = v1480
	var v1483 int32
	_ = v1483
	var v1485 int32
	_ = v1485
	var v1487 int32
	_ = v1487
	var v1489 int32
	_ = v1489
	var v1491 int32
	_ = v1491
	var v1495 int32
	_ = v1495
	var v1499 int32
	_ = v1499
	var v1511 int32
	_ = v1511
	var v1523 int32
	_ = v1523
	var v1524 int32
	_ = v1524
	var v1528 int32
	_ = v1528
	var v1530 int32
	_ = v1530
	var v1536 int32
	_ = v1536
	var v1550 int32
	_ = v1550
	var v1554 int32
	_ = v1554
	var v1555 int32
	_ = v1555
	var v1556 int32
	_ = v1556
	var v1557 int32
	_ = v1557
	var v1559 int32
	_ = v1559
	var v1560 int32
	_ = v1560
	var v1564 int32
	_ = v1564
	var v1566 int32
	_ = v1566
	var v1567 int32
	_ = v1567
	var v1570 int32
	_ = v1570
	var v1571 int32
	_ = v1571
	var v1576 int32
	_ = v1576
	var v1578 int32
	_ = v1578
	var v1583 int32
	_ = v1583
	var v1584 int32
	_ = v1584
	var v1585 int32
	_ = v1585
	var v1589 int32
	_ = v1589
	var v1592 int32
	_ = v1592
	var v1593 int32
	_ = v1593
	var v1598 int32
	_ = v1598
	var v1599 int32
	_ = v1599
	var v1603 int32
	_ = v1603
	var v1604 int32
	_ = v1604
	var v1611 int32
	_ = v1611
	var v1612 int32
	_ = v1612
	var v1616 int32
	_ = v1616
	v2 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v6
	v8 = int32(6)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v12-v6 < v8 {
		v22 = v2
		goto L3
	} else {
		goto L4
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v6
	v43 = v6
	goto L14
L2:
	;
	if v22 == int32(0) {
		goto L1
	} else {
		goto L6
	}
L3:
	;
	goto L2
L4:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v18 = F_memcmp(m, v16+v6, int32(2220979), v8)
	mBase = m.M
	if v18 != 0 {
		v22 = v2
		goto L3
	} else {
		goto L5
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v8 + v6
	v22 = int32(1)
	goto L3
L6:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v25 < v26 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v25
	v32 = F_slice_from_s(m, l0, int32(5), int32(2220985))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return int32(0)
L9:
	;
	if int32(0) <= v32 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v38 = int32(1)
	goto L12
L11:
	;
	v38 = v32
	goto L12
L12:
	;
	return v38
L13:
	;
	return v1616
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v43
	v50 = F_find_among(m, l0, int32(4251312), int32(7))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L8
	} else {
		goto L17
	}
L15:
	;
	v102 = v6
	goto L42
L16:
	;
	goto L15
L17:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v52
	switch v50 - int32(1) {
	case 0:
		goto L25
	case 1:
		goto L24
	case 2:
		goto L23
	case 3:
		goto L22
	case 4:
		goto L21
	case 5:
		goto L20
	case 6:
		goto L19
	default:
		goto L18
	}
L18:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v43 = v98
	goto L14
L19:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v92 <= v52 {
		goto L38
	} else {
		goto L39
	}
L20:
	;
	v88 = F_slice_from_s(m, l0, int32(2), int32(2220995))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L8
	} else {
		goto L36
	}
L21:
	;
	v82 = F_slice_from_s(m, l0, int32(1), int32(2220994))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L8
	} else {
		goto L34
	}
L22:
	;
	v76 = F_slice_from_s(m, l0, int32(1), int32(2220993))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L8
	} else {
		goto L32
	}
L23:
	;
	v70 = F_slice_from_s(m, l0, int32(1), int32(2220992))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L8
	} else {
		goto L30
	}
L24:
	;
	v64 = F_slice_from_s(m, l0, int32(1), int32(2220991))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L8
	} else {
		goto L28
	}
L25:
	;
	v58 = F_slice_from_s(m, l0, int32(1), int32(2220990))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L8
	} else {
		goto L26
	}
L26:
	;
	if int32(0) <= v58 {
		goto L18
	} else {
		goto L27
	}
L27:
	;
	v1616 = v58
	goto L13
L28:
	;
	if int32(0) <= v64 {
		goto L18
	} else {
		goto L29
	}
L29:
	;
	v1616 = v64
	goto L13
L30:
	;
	if int32(0) <= v70 {
		goto L18
	} else {
		goto L31
	}
L31:
	;
	v1616 = v70
	goto L13
L32:
	;
	if int32(0) <= v76 {
		goto L18
	} else {
		goto L33
	}
L33:
	;
	v1616 = v76
	goto L13
L34:
	;
	if int32(0) <= v82 {
		goto L18
	} else {
		goto L35
	}
L35:
	;
	v1616 = v82
	goto L13
L36:
	;
	if int32(0) <= v88 {
		goto L18
	} else {
		goto L37
	}
L37:
	;
	v1616 = v88
	goto L13
L38:
	;
	goto L16
L39:
	;
	goto L40
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v52 + int32(1)
	goto L18
L41:
	;
	v1392 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1392
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1392
	v1404 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L434
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v102
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v114 < v102 {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	v1370 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1370
	v1372 = *(*int32)(unsafe.Add(mBase, uint32(v1368)+8))
	if v1370 < v1372 {
		goto L41
	} else {
		goto L423
	}
L44:
	;
	goto L43
L45:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v157 != 0 {
		v302 = v158
		goto L60
	} else {
		goto L61
	}
L46:
	;
	v116 = v102
	goto L48
L47:
	;
	v116 = v114
	goto L48
L48:
	;
	goto L50
L49:
	;
	v157 = v153
	goto L45
L50:
	;
	if v102 == v116 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v153 = int32(0)
	goto L49
L52:
	;
	v157 = int32(-1)
	goto L45
L53:
	;
	goto L54
L54:
	;
	v128 = int32(1)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129+v102))))
	if int32(249) < v131 {
		v153 = v128
		goto L49
	} else {
		goto L55
	}
L55:
	;
	v133 = v131 - int32(97)
	if v133 < int32(0) {
		v153 = v128
		goto L49
	} else {
		goto L56
	}
L56:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v133)>>(uint(int32(3))%32)))+uint32(_consts[1436]))))
	if int32(base.Ui32(v139)>>(uint(v133&int32(7))%32))&int32(1) == int32(0) {
		v153 = v128
		goto L49
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v102 + int32(1)
	goto L58
L58:
	;
	goto L51
L59:
	;
	v1364 = F_slice_from_s(m, l0, int32(1), int32(2221029))
	mBase = m.M
	v1365 = m.ExcPending
	if v1365 != 0 {
		goto L8
	} else {
		goto L421
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v102
	if v102 < v302 {
		goto L103
	} else {
		goto L104
	}
L61:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v159
	if v159 == v158 {
		v233 = v158
		goto L62
	} else {
		goto L63
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v159
	if v159 == v233 {
		goto L84
	} else {
		goto L85
	}
L63:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162+v159))))
	if v164 != int32(117) {
		v233 = v158
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v168 = v159 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v168
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v168
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v180 < v168 {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	if v223 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L66:
	;
	v182 = v168
	goto L68
L67:
	;
	v182 = v180
	goto L68
L68:
	;
	goto L70
L69:
	;
	v223 = v219
	goto L65
L70:
	;
	if v168 == v182 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	v219 = int32(0)
	goto L69
L72:
	;
	v223 = int32(-1)
	goto L65
L73:
	;
	goto L74
L74:
	;
	v194 = int32(1)
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195+v168))))
	if int32(249) < v197 {
		v219 = v194
		goto L69
	} else {
		goto L75
	}
L75:
	;
	v199 = v197 - int32(97)
	if v199 < int32(0) {
		v219 = v194
		goto L69
	} else {
		goto L76
	}
L76:
	;
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v199)>>(uint(int32(3))%32)))+uint32(_consts[1436]))))
	if int32(base.Ui32(v205)>>(uint(v199&int32(7))%32))&int32(1) == int32(0) {
		v219 = v194
		goto L69
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v159 + int32(2)
	goto L78
L78:
	;
	goto L71
L79:
	;
	v228 = F_slice_from_s(m, l0, int32(1), int32(2221028))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L8
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v233 = v232
	goto L62
L82:
	;
	if v228 < int32(0) {
		v1616 = v228
		goto L13
	} else {
		goto L83
	}
L83:
	;
	goto L42
L84:
	;
	v302 = v159
	goto L60
L85:
	;
	goto L86
L86:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v236+v159))))
	if v238 != int32(105) {
		v302 = v233
		goto L60
	} else {
		goto L87
	}
L87:
	;
	v242 = v159 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v242
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v242
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v254 < v242 {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	if v297 == int32(0) {
		goto L59
	} else {
		goto L102
	}
L89:
	;
	v256 = v242
	goto L91
L90:
	;
	v256 = v254
	goto L91
L91:
	;
	goto L93
L92:
	;
	v297 = v293
	goto L88
L93:
	;
	if v242 == v256 {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	v293 = int32(0)
	goto L92
L95:
	;
	v297 = int32(-1)
	goto L88
L96:
	;
	goto L97
L97:
	;
	v268 = int32(1)
	v269 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269+v242))))
	if int32(249) < v271 {
		v293 = v268
		goto L92
	} else {
		goto L98
	}
L98:
	;
	v273 = v271 - int32(97)
	if v273 < int32(0) {
		v293 = v268
		goto L92
	} else {
		goto L99
	}
L99:
	;
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v273)>>(uint(int32(3))%32)))+uint32(_consts[1436]))))
	if int32(base.Ui32(v279)>>(uint(v273&int32(7))%32))&int32(1) == int32(0) {
		v293 = v268
		goto L92
	} else {
		goto L100
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v159 + int32(2)
	goto L101
L101:
	;
	goto L94
L102:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v302 = v300
	goto L60
L103:
	;
	v102 = v102 + int32(1)
	goto L42
L104:
	;
	goto L105
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v6
	v308 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v308)+4)) = v302
	*(*int32)(unsafe.Add(mBase, uint32(v308)+8)) = v302
	v311 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v308))) = v311
	v313 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v323 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v323 < v313 {
		goto L110
	} else {
		goto L111
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v313
	v813 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v813 < v313 {
		goto L254
	} else {
		goto L255
	}
L107:
	;
	v800 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v800)+8)) = v799
	goto L106
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v313
	v593 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v593 < v313 {
		goto L188
	} else {
		goto L189
	}
L109:
	;
	if v366 != 0 {
		goto L108
	} else {
		goto L123
	}
L110:
	;
	v325 = v313
	goto L112
L111:
	;
	v325 = v323
	goto L112
L112:
	;
	goto L114
L113:
	;
	v366 = v362
	goto L109
L114:
	;
	if v313 == v325 {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	v362 = int32(0)
	goto L113
L116:
	;
	v366 = int32(-1)
	goto L109
L117:
	;
	goto L118
L118:
	;
	v337 = int32(1)
	v338 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v338+v313))))
	if int32(249) < v340 {
		v362 = v337
		goto L113
	} else {
		goto L119
	}
L119:
	;
	v342 = v340 - int32(97)
	if v342 < int32(0) {
		v362 = v337
		goto L113
	} else {
		goto L120
	}
L120:
	;
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v342)>>(uint(int32(3))%32)))+uint32(_consts[1436]))))
	if int32(base.Ui32(v348)>>(uint(v342&int32(7))%32))&int32(1) == int32(0) {
		v362 = v337
		goto L113
	} else {
		goto L121
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v313 + int32(1)
	goto L122
L122:
	;
	goto L115
L123:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v376 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v376 < v367 {
		goto L126
	} else {
		goto L127
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v367
	v481 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v481 < v367 {
		goto L158
	} else {
		goto L159
	}
L125:
	;
	if v416 != 0 {
		goto L124
	} else {
		goto L140
	}
L126:
	;
	v378 = v367
	goto L128
L127:
	;
	v378 = v376
	goto L128
L128:
	;
	goto L130
L129:
	;
	v416 = v413
	goto L125
L130:
	;
	if v367 == v378 {
		goto L132
	} else {
		goto L133
	}
L131:
	;
	v413 = int32(0)
	goto L129
L132:
	;
	v416 = int32(-1)
	goto L125
L133:
	;
	goto L134
L134:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v389+v367))))
	if int32(249) < v391 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v367 + int32(1)
	goto L139
L136:
	;
	v393 = v391 - int32(97)
	if v393 < int32(0) {
		goto L135
	} else {
		goto L137
	}
L137:
	;
	v396 = int32(1)
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v393)>>(uint(int32(3))%32)))+uint32(_consts[1436]))))
	if int32(base.Ui32(v400)>>(uint(v393&int32(7))%32))&v396 != 0 {
		v413 = v396
		goto L129
	} else {
		goto L138
	}
L138:
	;
	goto L135
L139:
	;
	goto L131
L140:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v425 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v425 < v424 {
		goto L142
	} else {
		goto L143
	}
L141:
	;
	if v465 < int32(0) {
		goto L124
	} else {
		goto L156
	}
L142:
	;
	v427 = v424
	goto L144
L143:
	;
	v427 = v425
	goto L144
L144:
	;
	v434 = v424
	goto L146
L145:
	;
	v465 = v445
	goto L141
L146:
	;
	if v434 == v427 {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v465 = int32(-1)
	goto L141
L149:
	;
	goto L150
L150:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438+v434))))
	if int32(249) < v440 {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v457 = v434 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v457
	v434 = v457
	goto L146
L152:
	;
	v442 = v440 - int32(97)
	if v442 < int32(0) {
		goto L151
	} else {
		goto L153
	}
L153:
	;
	v445 = int32(1)
	v449 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v442)>>(uint(int32(3))%32)))+uint32(_consts[1436]))))
	if int32(base.Ui32(v449)>>(uint(v442&int32(7))%32))&v445 != 0 {
		goto L145
	} else {
		goto L154
	}
L154:
	;
	goto L151
L156:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v799 = v468 + v465
	goto L107
L157:
	;
	if v524 != 0 {
		goto L108
	} else {
		goto L171
	}
L158:
	;
	v483 = v367
	goto L160
L159:
	;
	v483 = v481
	goto L160
L160:
	;
	goto L162
L161:
	;
	v524 = v520
	goto L157
L162:
	;
	if v367 == v483 {
		goto L164
	} else {
		goto L165
	}
L163:
	;
	v520 = int32(0)
	goto L161
L164:
	;
	v524 = int32(-1)
	goto L157
L165:
	;
	goto L166
L166:
	;
	v495 = int32(1)
	v496 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v498 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496+v367))))
	if int32(249) < v498 {
		v520 = v495
		goto L161
	} else {
		goto L167
	}
L167:
	;
	v500 = v498 - int32(97)
	if v500 < int32(0) {
		v520 = v495
		goto L161
	} else {
		goto L168
	}
L168:
	;
	v506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v500)>>(uint(int32(3))%32)))+uint32(_consts[1436]))))
	if int32(base.Ui32(v506)>>(uint(v500&int32(7))%32))&int32(1) == int32(0) {
		v520 = v495
		goto L161
	} else {
		goto L169
	}
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v367 + int32(1)
	goto L170
L170:
	;
	goto L163
L171:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v534 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v534 < v533 {
		goto L173
	} else {
		goto L174
	}
L172:
	;
	if v577 < int32(0) {
		goto L108
	} else {
		goto L186
	}
L173:
	;
	v536 = v533
	goto L175
L174:
	;
	v536 = v534
	goto L175
L175:
	;
	v543 = v533
	goto L177
L176:
	;
	v577 = int32(1)
	goto L172
L177:
	;
	if v543 == v536 {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	v577 = int32(-1)
	goto L172
L180:
	;
	goto L181
L181:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v551 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v549+v543))))
	if int32(249) < v551 {
		goto L176
	} else {
		goto L182
	}
L182:
	;
	v553 = v551 - int32(97)
	if v553 < int32(0) {
		goto L176
	} else {
		goto L183
	}
L183:
	;
	v559 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v553)>>(uint(int32(3))%32)))+uint32(_consts[1436]))))
	if int32(base.Ui32(v559)>>(uint(v553&int32(7))%32))&int32(1) == int32(0) {
		goto L176
	} else {
		goto L184
	}
L184:
	;
	v568 = v543 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v568
	v543 = v568
	goto L177
L186:
	;
	v580 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v799 = v580 + v577
	goto L107
L187:
	;
	if v633 != 0 {
		goto L106
	} else {
		goto L202
	}
L188:
	;
	v595 = v313
	goto L190
L189:
	;
	v595 = v593
	goto L190
L190:
	;
	goto L192
L191:
	;
	v633 = v630
	goto L187
L192:
	;
	if v313 == v595 {
		goto L194
	} else {
		goto L195
	}
L193:
	;
	v630 = int32(0)
	goto L191
L194:
	;
	v633 = int32(-1)
	goto L187
L195:
	;
	goto L196
L196:
	;
	v606 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v608 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v606+v313))))
	if int32(249) < v608 {
		goto L197
	} else {
		goto L198
	}
L197:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v313 + int32(1)
	goto L201
L198:
	;
	v610 = v608 - int32(97)
	if v610 < int32(0) {
		goto L197
	} else {
		goto L199
	}
L199:
	;
	v613 = int32(1)
	v617 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v610)>>(uint(int32(3))%32)))+uint32(_consts[1436]))))
	if int32(base.Ui32(v617)>>(uint(v610&int32(7))%32))&v613 != 0 {
		v630 = v613
		goto L191
	} else {
		goto L200
	}
L200:
	;
	goto L197
L201:
	;
	goto L193
L202:
	;
	v634 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v643 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v643 < v634 {
		goto L205
	} else {
		goto L206
	}
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v634
	v748 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v748 < v634 {
		goto L237
	} else {
		goto L238
	}
L204:
	;
	if v683 != 0 {
		goto L203
	} else {
		goto L219
	}
L205:
	;
	v645 = v634
	goto L207
L206:
	;
	v645 = v643
	goto L207
L207:
	;
	goto L209
L208:
	;
	v683 = v680
	goto L204
L209:
	;
	if v634 == v645 {
		goto L211
	} else {
		goto L212
	}
L210:
	;
	v680 = int32(0)
	goto L208
L211:
	;
	v683 = int32(-1)
	goto L204
L212:
	;
	goto L213
L213:
	;
	v656 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v658 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v656+v634))))
	if int32(249) < v658 {
		goto L214
	} else {
		goto L215
	}
L214:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v634 + int32(1)
	goto L218
L215:
	;
	v660 = v658 - int32(97)
	if v660 < int32(0) {
		goto L214
	} else {
		goto L216
	}
L216:
	;
	v663 = int32(1)
	v667 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v660)>>(uint(int32(3))%32)))+uint32(_consts[1436]))))
	if int32(base.Ui32(v667)>>(uint(v660&int32(7))%32))&v663 != 0 {
		v680 = v663
		goto L208
	} else {
		goto L217
	}
L217:
	;
	goto L214
L218:
	;
	goto L210
L219:
	;
	v691 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v692 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v692 < v691 {
		goto L221
	} else {
		goto L222
	}
L220:
	;
	if v732 < int32(0) {
		goto L203
	} else {
		goto L235
	}
L221:
	;
	v694 = v691
	goto L223
L222:
	;
	v694 = v692
	goto L223
L223:
	;
	v701 = v691
	goto L225
L224:
	;
	v732 = v712
	goto L220
L225:
	;
	if v701 == v694 {
		goto L227
	} else {
		goto L228
	}
L227:
	;
	v732 = int32(-1)
	goto L220
L228:
	;
	goto L229
L229:
	;
	v705 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v707 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v705+v701))))
	if int32(249) < v707 {
		goto L230
	} else {
		goto L231
	}
L230:
	;
	v724 = v701 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v724
	v701 = v724
	goto L225
L231:
	;
	v709 = v707 - int32(97)
	if v709 < int32(0) {
		goto L230
	} else {
		goto L232
	}
L232:
	;
	v712 = int32(1)
	v716 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v709)>>(uint(int32(3))%32)))+uint32(_consts[1436]))))
	if int32(base.Ui32(v716)>>(uint(v709&int32(7))%32))&v712 != 0 {
		goto L224
	} else {
		goto L233
	}
L233:
	;
	goto L230
L235:
	;
	v735 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v799 = v735 + v732
	goto L107
L236:
	;
	if v791 != 0 {
		goto L106
	} else {
		goto L250
	}
L237:
	;
	v750 = v634
	goto L239
L238:
	;
	v750 = v748
	goto L239
L239:
	;
	goto L241
L240:
	;
	v791 = v787
	goto L236
L241:
	;
	if v634 == v750 {
		goto L243
	} else {
		goto L244
	}
L242:
	;
	v787 = int32(0)
	goto L240
L243:
	;
	v791 = int32(-1)
	goto L236
L244:
	;
	goto L245
L245:
	;
	v762 = int32(1)
	v763 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v765 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v763+v634))))
	if int32(249) < v765 {
		v787 = v762
		goto L240
	} else {
		goto L246
	}
L246:
	;
	v767 = v765 - int32(97)
	if v767 < int32(0) {
		v787 = v762
		goto L240
	} else {
		goto L247
	}
L247:
	;
	v773 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v767)>>(uint(int32(3))%32)))+uint32(_consts[1436]))))
	if int32(base.Ui32(v773)>>(uint(v767&int32(7))%32))&int32(1) == int32(0) {
		v787 = v762
		goto L240
	} else {
		goto L248
	}
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v634 + int32(1)
	goto L249
L249:
	;
	goto L242
L250:
	;
	v792 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v793 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v793 <= v792 {
		goto L106
	} else {
		goto L251
	}
L251:
	;
	v799 = v792 + int32(1)
	goto L107
L252:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v313
	v1034 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1034
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1034
	v1038 = v1034 - int32(1)
	if v1038 <= v313 {
		goto L315
	} else {
		goto L316
	}
L253:
	;
	if v853 < int32(0) {
		goto L252
	} else {
		goto L268
	}
L254:
	;
	v815 = v313
	goto L256
L255:
	;
	v815 = v813
	goto L256
L256:
	;
	v822 = v313
	goto L258
L257:
	;
	v853 = v833
	goto L253
L258:
	;
	if v822 == v815 {
		goto L260
	} else {
		goto L261
	}
L260:
	;
	v853 = int32(-1)
	goto L253
L261:
	;
	goto L262
L262:
	;
	v826 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v828 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v826+v822))))
	if int32(249) < v828 {
		goto L263
	} else {
		goto L264
	}
L263:
	;
	v845 = v822 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v845
	v822 = v845
	goto L258
L264:
	;
	v830 = v828 - int32(97)
	if v830 < int32(0) {
		goto L263
	} else {
		goto L265
	}
L265:
	;
	v833 = int32(1)
	v837 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v830)>>(uint(int32(3))%32)))+uint32(_consts[1436]))))
	if int32(base.Ui32(v837)>>(uint(v830&int32(7))%32))&v833 != 0 {
		goto L257
	} else {
		goto L266
	}
L266:
	;
	goto L263
L268:
	;
	v856 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v857 = v856 + v853
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v857
	v868 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v868 < v857 {
		goto L270
	} else {
		goto L271
	}
L269:
	;
	if v911 < int32(0) {
		goto L252
	} else {
		goto L283
	}
L270:
	;
	v870 = v857
	goto L272
L271:
	;
	v870 = v868
	goto L272
L272:
	;
	v877 = v857
	goto L274
L273:
	;
	v911 = int32(1)
	goto L269
L274:
	;
	if v877 == v870 {
		goto L276
	} else {
		goto L277
	}
L276:
	;
	v911 = int32(-1)
	goto L269
L277:
	;
	goto L278
L278:
	;
	v883 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v885 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v883+v877))))
	if int32(249) < v885 {
		goto L273
	} else {
		goto L279
	}
L279:
	;
	v887 = v885 - int32(97)
	if v887 < int32(0) {
		goto L273
	} else {
		goto L280
	}
L280:
	;
	v893 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v887)>>(uint(int32(3))%32)))+uint32(_consts[1436]))))
	if int32(base.Ui32(v893)>>(uint(v887&int32(7))%32))&int32(1) == int32(0) {
		goto L273
	} else {
		goto L281
	}
L281:
	;
	v902 = v877 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v902
	v877 = v902
	goto L274
L283:
	;
	v914 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v915 = v914 + v911
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v915
	v917 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v917)+4)) = v915
	v926 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v927 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v927 < v926 {
		goto L285
	} else {
		goto L286
	}
L284:
	;
	if v967 < int32(0) {
		goto L252
	} else {
		goto L299
	}
L285:
	;
	v929 = v926
	goto L287
L286:
	;
	v929 = v927
	goto L287
L287:
	;
	v936 = v926
	goto L289
L288:
	;
	v967 = v947
	goto L284
L289:
	;
	if v936 == v929 {
		goto L291
	} else {
		goto L292
	}
L291:
	;
	v967 = int32(-1)
	goto L284
L292:
	;
	goto L293
L293:
	;
	v940 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v942 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v940+v936))))
	if int32(249) < v942 {
		goto L294
	} else {
		goto L295
	}
L294:
	;
	v959 = v936 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v959
	v936 = v959
	goto L289
L295:
	;
	v944 = v942 - int32(97)
	if v944 < int32(0) {
		goto L294
	} else {
		goto L296
	}
L296:
	;
	v947 = int32(1)
	v951 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v944)>>(uint(int32(3))%32)))+uint32(_consts[1436]))))
	if int32(base.Ui32(v951)>>(uint(v944&int32(7))%32))&v947 != 0 {
		goto L288
	} else {
		goto L297
	}
L297:
	;
	goto L294
L299:
	;
	v970 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v971 = v970 + v967
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v971
	v982 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v982 < v971 {
		goto L301
	} else {
		goto L302
	}
L300:
	;
	if v1025 < int32(0) {
		goto L252
	} else {
		goto L314
	}
L301:
	;
	v984 = v971
	goto L303
L302:
	;
	v984 = v982
	goto L303
L303:
	;
	v991 = v971
	goto L305
L304:
	;
	v1025 = int32(1)
	goto L300
L305:
	;
	if v991 == v984 {
		goto L307
	} else {
		goto L308
	}
L307:
	;
	v1025 = int32(-1)
	goto L300
L308:
	;
	goto L309
L309:
	;
	v997 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v999 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v997+v991))))
	if int32(249) < v999 {
		goto L304
	} else {
		goto L310
	}
L310:
	;
	v1001 = v999 - int32(97)
	if v1001 < int32(0) {
		goto L304
	} else {
		goto L311
	}
L311:
	;
	v1007 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1001)>>(uint(int32(3))%32)))+uint32(_consts[1436]))))
	if int32(base.Ui32(v1007)>>(uint(v1001&int32(7))%32))&int32(1) == int32(0) {
		goto L304
	} else {
		goto L312
	}
L312:
	;
	v1016 = v991 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1016
	v991 = v1016
	goto L305
L314:
	;
	v1028 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1029 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1028))) = v1029 + v1025
	goto L252
L315:
	;
	v1093 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1093
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1093
	v1098 = F_find_among_b(m, l0, int32(4252320), int32(51))
	mBase = m.M
	v1099 = m.ExcPending
	if v1099 != 0 {
		goto L8
	} else {
		goto L332
	}
L316:
	;
	v1040 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1042 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1040+v1038))))
	if v1042&int32(224) != int32(96) {
		goto L315
	} else {
		goto L317
	}
L317:
	;
	if int32(1)<<(uint(v1042)%32)&int32(33314) == int32(0) {
		goto L315
	} else {
		goto L318
	}
L318:
	;
	v1055 = F_find_among_b(m, l0, int32(4251456), int32(37))
	mBase = m.M
	v1056 = m.ExcPending
	if v1056 != 0 {
		goto L8
	} else {
		goto L319
	}
L319:
	;
	if v1055 == int32(0) {
		goto L315
	} else {
		goto L320
	}
L320:
	;
	v1059 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1059
	v1062 = v1059 - int32(1)
	v1063 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1062 <= v1063 {
		goto L315
	} else {
		goto L321
	}
L321:
	;
	v1065 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1067 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1065+v1062))))
	switch v1067 - int32(111) {
	case 0, 3:
		goto L322
	default:
		goto L315
	}
L322:
	;
	v1072 = F_find_among_b(m, l0, int32(4252208), int32(5))
	mBase = m.M
	v1073 = m.ExcPending
	if v1073 != 0 {
		goto L8
	} else {
		goto L323
	}
L323:
	;
	if v1072 == int32(0) {
		goto L315
	} else {
		goto L324
	}
L324:
	;
	v1076 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1077 = *(*int32)(unsafe.Add(mBase, uint32(v1076)+8))
	v1078 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v1078 < v1077 {
		goto L315
	} else {
		goto L325
	}
L325:
	;
	switch v1072 - int32(1) {
	case 0:
		goto L327
	case 1:
		goto L326
	default:
		goto L315
	}
L326:
	;
	v1088 = F_slice_from_s(m, l0, int32(1), int32(2221037))
	mBase = m.M
	v1089 = m.ExcPending
	if v1089 != 0 {
		goto L8
	} else {
		goto L330
	}
L327:
	;
	v1082 = F_slice_del(m, l0)
	mBase = m.M
	v1083 = m.ExcPending
	if v1083 != 0 {
		goto L8
	} else {
		goto L328
	}
L328:
	;
	if int32(0) <= v1082 {
		goto L315
	} else {
		goto L329
	}
L329:
	;
	v1616 = v1082
	goto L13
L330:
	;
	if v1088 < int32(0) {
		v1616 = v1088
		goto L13
	} else {
		goto L331
	}
L331:
	;
	goto L315
L332:
	;
	if v1098 == int32(0) {
		goto L333
	} else {
		goto L334
	}
L333:
	;
	v1102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1368 = v1102
	goto L44
L334:
	;
	goto L335
L335:
	;
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1103
	switch v1098 - int32(1) {
	case 0:
		goto L344
	case 1:
		goto L343
	case 2:
		goto L342
	case 3:
		goto L341
	case 4:
		goto L340
	case 5:
		goto L339
	case 6:
		goto L338
	case 7:
		goto L337
	case 8:
		goto L336
	default:
		goto L41
	}
L336:
	;
	v1297 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1298 = *(*int32)(unsafe.Add(mBase, uint32(v1297)))
	if v1103 < v1298 {
		v1368 = v1297
		goto L44
	} else {
		goto L402
	}
L337:
	;
	v1256 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1257 = *(*int32)(unsafe.Add(mBase, uint32(v1256)))
	if v1103 < v1257 {
		v1368 = v1256
		goto L44
	} else {
		goto L391
	}
L338:
	;
	v1184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1185 = *(*int32)(unsafe.Add(mBase, uint32(v1184)+4))
	if v1103 < v1185 {
		v1368 = v1184
		goto L44
	} else {
		goto L371
	}
L339:
	;
	v1177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1178 = *(*int32)(unsafe.Add(mBase, uint32(v1177)+8))
	if v1103 < v1178 {
		v1368 = v1177
		goto L44
	} else {
		goto L368
	}
L340:
	;
	v1168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1169 = *(*int32)(unsafe.Add(mBase, uint32(v1168)))
	if v1103 < v1169 {
		v1368 = v1168
		goto L44
	} else {
		goto L365
	}
L341:
	;
	v1159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1160 = *(*int32)(unsafe.Add(mBase, uint32(v1159)))
	if v1103 < v1160 {
		v1368 = v1159
		goto L44
	} else {
		goto L362
	}
L342:
	;
	v1150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1151 = *(*int32)(unsafe.Add(mBase, uint32(v1150)))
	if v1103 < v1151 {
		v1368 = v1150
		goto L44
	} else {
		goto L359
	}
L343:
	;
	v1114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1115 = *(*int32)(unsafe.Add(mBase, uint32(v1114)))
	if v1103 < v1115 {
		v1368 = v1114
		goto L44
	} else {
		goto L348
	}
L344:
	;
	v1107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1108 = *(*int32)(unsafe.Add(mBase, uint32(v1107)))
	if v1103 < v1108 {
		v1368 = v1107
		goto L44
	} else {
		goto L345
	}
L345:
	;
	v1110 = F_slice_del(m, l0)
	mBase = m.M
	v1111 = m.ExcPending
	if v1111 != 0 {
		goto L8
	} else {
		goto L346
	}
L346:
	;
	if int32(0) <= v1110 {
		goto L41
	} else {
		goto L347
	}
L347:
	;
	v1616 = v1110
	goto L13
L348:
	;
	v1117 = F_slice_del(m, l0)
	mBase = m.M
	v1118 = m.ExcPending
	if v1118 != 0 {
		goto L8
	} else {
		goto L349
	}
L349:
	;
	if v1117 < int32(0) {
		v1616 = v1117
		goto L13
	} else {
		goto L350
	}
L350:
	;
	v1121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1121
	v1123 = int32(2)
	v1125 = int32(0)
	v1128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1121-v1128 < v1123 {
		v1138 = v1125
		goto L352
	} else {
		goto L353
	}
L351:
	;
	if v1138 == int32(0) {
		goto L41
	} else {
		goto L355
	}
L352:
	;
	goto L351
L353:
	;
	v1131 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1134 = F_memcmp(m, v1131+v1121-v1123, int32(2221189), v1123)
	mBase = m.M
	if v1134 != 0 {
		v1138 = v1125
		goto L352
	} else {
		goto L354
	}
L354:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1121 - v1123
	v1138 = int32(1)
	goto L352
L355:
	;
	v1141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1141
	v1143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1144 = *(*int32)(unsafe.Add(mBase, uint32(v1143)))
	if v1141 < v1144 {
		goto L41
	} else {
		goto L356
	}
L356:
	;
	v1146 = F_slice_del(m, l0)
	mBase = m.M
	v1147 = m.ExcPending
	if v1147 != 0 {
		goto L8
	} else {
		goto L357
	}
L357:
	;
	if int32(0) <= v1146 {
		goto L41
	} else {
		goto L358
	}
L358:
	;
	v1616 = v1146
	goto L13
L359:
	;
	v1155 = F_slice_from_s(m, l0, int32(3), int32(2221191))
	mBase = m.M
	v1156 = m.ExcPending
	if v1156 != 0 {
		goto L8
	} else {
		goto L360
	}
L360:
	;
	if int32(0) <= v1155 {
		goto L41
	} else {
		goto L361
	}
L361:
	;
	v1616 = v1155
	goto L13
L362:
	;
	v1164 = F_slice_from_s(m, l0, int32(1), int32(2221194))
	mBase = m.M
	v1165 = m.ExcPending
	if v1165 != 0 {
		goto L8
	} else {
		goto L363
	}
L363:
	;
	if int32(0) <= v1164 {
		goto L41
	} else {
		goto L364
	}
L364:
	;
	v1616 = v1164
	goto L13
L365:
	;
	v1173 = F_slice_from_s(m, l0, int32(4), int32(2221195))
	mBase = m.M
	v1174 = m.ExcPending
	if v1174 != 0 {
		goto L8
	} else {
		goto L366
	}
L366:
	;
	if int32(0) <= v1173 {
		goto L41
	} else {
		goto L367
	}
L367:
	;
	v1616 = v1173
	goto L13
L368:
	;
	v1180 = F_slice_del(m, l0)
	mBase = m.M
	v1181 = m.ExcPending
	if v1181 != 0 {
		goto L8
	} else {
		goto L369
	}
L369:
	;
	if int32(0) <= v1180 {
		goto L41
	} else {
		goto L370
	}
L370:
	;
	v1616 = v1180
	goto L13
L371:
	;
	v1187 = F_slice_del(m, l0)
	mBase = m.M
	v1188 = m.ExcPending
	if v1188 != 0 {
		goto L8
	} else {
		goto L372
	}
L372:
	;
	if v1187 < int32(0) {
		v1616 = v1187
		goto L13
	} else {
		goto L373
	}
L373:
	;
	v1191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1191
	v1194 = v1191 - int32(1)
	v1195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1194 <= v1195 {
		goto L41
	} else {
		goto L374
	}
L374:
	;
	v1197 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1197+v1194))))
	if v1199&int32(224) != int32(96) {
		goto L41
	} else {
		goto L375
	}
L375:
	;
	if int32(1)<<(uint(v1199)%32)&int32(4722696) == int32(0) {
		goto L41
	} else {
		goto L376
	}
L376:
	;
	v1212 = F_find_among_b(m, l0, int32(4253344), int32(4))
	mBase = m.M
	v1213 = m.ExcPending
	if v1213 != 0 {
		goto L8
	} else {
		goto L377
	}
L377:
	;
	if v1212 == int32(0) {
		goto L41
	} else {
		goto L378
	}
L378:
	;
	v1216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1216
	v1218 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1219 = *(*int32)(unsafe.Add(mBase, uint32(v1218)))
	if v1216 < v1219 {
		goto L41
	} else {
		goto L379
	}
L379:
	;
	v1221 = F_slice_del(m, l0)
	mBase = m.M
	v1222 = m.ExcPending
	if v1222 != 0 {
		goto L8
	} else {
		goto L380
	}
L380:
	;
	if v1221 < int32(0) {
		v1616 = v1221
		goto L13
	} else {
		goto L381
	}
L381:
	;
	if v1212 != int32(1) {
		goto L41
	} else {
		goto L382
	}
L382:
	;
	v1227 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1227
	v1229 = int32(2)
	v1231 = int32(0)
	v1234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1227-v1234 < v1229 {
		v1244 = v1231
		goto L384
	} else {
		goto L385
	}
L383:
	;
	if v1244 == int32(0) {
		goto L41
	} else {
		goto L387
	}
L384:
	;
	goto L383
L385:
	;
	v1237 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1240 = F_memcmp(m, v1237+v1227-v1229, int32(2221199), v1229)
	mBase = m.M
	if v1240 != 0 {
		v1244 = v1231
		goto L384
	} else {
		goto L386
	}
L386:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1227 - v1229
	v1244 = int32(1)
	goto L384
L387:
	;
	v1247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1247
	v1249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1250 = *(*int32)(unsafe.Add(mBase, uint32(v1249)))
	if v1247 < v1250 {
		goto L41
	} else {
		goto L388
	}
L388:
	;
	v1252 = F_slice_del(m, l0)
	mBase = m.M
	v1253 = m.ExcPending
	if v1253 != 0 {
		goto L8
	} else {
		goto L389
	}
L389:
	;
	if int32(0) <= v1252 {
		goto L41
	} else {
		goto L390
	}
L390:
	;
	v1616 = v1252
	goto L13
L391:
	;
	v1259 = F_slice_del(m, l0)
	mBase = m.M
	v1260 = m.ExcPending
	if v1260 != 0 {
		goto L8
	} else {
		goto L392
	}
L392:
	;
	if v1259 < int32(0) {
		v1616 = v1259
		goto L13
	} else {
		goto L393
	}
L393:
	;
	v1263 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1263
	v1266 = v1263 - int32(1)
	v1267 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1266 <= v1267 {
		goto L41
	} else {
		goto L394
	}
L394:
	;
	v1269 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1269+v1266))))
	if v1271&int32(224) != int32(96) {
		goto L41
	} else {
		goto L395
	}
L395:
	;
	if int32(1)<<(uint(v1271)%32)&int32(4198408) == int32(0) {
		goto L41
	} else {
		goto L396
	}
L396:
	;
	v1284 = F_find_among_b(m, l0, int32(4253424), int32(3))
	mBase = m.M
	v1285 = m.ExcPending
	if v1285 != 0 {
		goto L8
	} else {
		goto L397
	}
L397:
	;
	if v1284 == int32(0) {
		goto L41
	} else {
		goto L398
	}
L398:
	;
	v1288 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1288
	v1290 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1291 = *(*int32)(unsafe.Add(mBase, uint32(v1290)))
	if v1288 < v1291 {
		goto L41
	} else {
		goto L399
	}
L399:
	;
	v1293 = F_slice_del(m, l0)
	mBase = m.M
	v1294 = m.ExcPending
	if v1294 != 0 {
		goto L8
	} else {
		goto L400
	}
L400:
	;
	if int32(0) <= v1293 {
		goto L41
	} else {
		goto L401
	}
L401:
	;
	v1616 = v1293
	goto L13
L402:
	;
	v1300 = F_slice_del(m, l0)
	mBase = m.M
	v1301 = m.ExcPending
	if v1301 != 0 {
		goto L8
	} else {
		goto L403
	}
L403:
	;
	if v1300 < int32(0) {
		v1616 = v1300
		goto L13
	} else {
		goto L404
	}
L404:
	;
	v1304 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1304
	v1306 = int32(2)
	v1308 = int32(0)
	v1311 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1304-v1311 < v1306 {
		v1321 = v1308
		goto L406
	} else {
		goto L407
	}
L405:
	;
	if v1321 == int32(0) {
		goto L41
	} else {
		goto L409
	}
L406:
	;
	goto L405
L407:
	;
	v1314 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1317 = F_memcmp(m, v1314+v1304-v1306, int32(2221201), v1306)
	mBase = m.M
	if v1317 != 0 {
		v1321 = v1308
		goto L406
	} else {
		goto L408
	}
L408:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1304 - v1306
	v1321 = int32(1)
	goto L406
L409:
	;
	v1324 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1324
	v1326 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1327 = *(*int32)(unsafe.Add(mBase, uint32(v1326)))
	if v1324 < v1327 {
		goto L41
	} else {
		goto L410
	}
L410:
	;
	v1329 = F_slice_del(m, l0)
	mBase = m.M
	v1330 = m.ExcPending
	if v1330 != 0 {
		goto L8
	} else {
		goto L411
	}
L411:
	;
	if v1329 < int32(0) {
		v1616 = v1329
		goto L13
	} else {
		goto L412
	}
L412:
	;
	v1333 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1333
	v1335 = int32(2)
	v1337 = int32(0)
	v1340 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1333-v1340 < v1335 {
		v1350 = v1337
		goto L414
	} else {
		goto L415
	}
L413:
	;
	if v1350 == int32(0) {
		goto L41
	} else {
		goto L417
	}
L414:
	;
	goto L413
L415:
	;
	v1343 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1346 = F_memcmp(m, v1343+v1333-v1335, int32(2221203), v1335)
	mBase = m.M
	if v1346 != 0 {
		v1350 = v1337
		goto L414
	} else {
		goto L416
	}
L416:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1333 - v1335
	v1350 = int32(1)
	goto L414
L417:
	;
	v1353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1353
	v1355 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1356 = *(*int32)(unsafe.Add(mBase, uint32(v1355)))
	if v1353 < v1356 {
		goto L41
	} else {
		goto L418
	}
L418:
	;
	v1358 = F_slice_del(m, l0)
	mBase = m.M
	v1359 = m.ExcPending
	if v1359 != 0 {
		goto L8
	} else {
		goto L419
	}
L419:
	;
	if v1358 < int32(0) {
		v1616 = v1358
		goto L13
	} else {
		goto L420
	}
L420:
	;
	goto L41
L421:
	;
	if int32(0) <= v1364 {
		goto L42
	} else {
		goto L422
	}
L422:
	;
	v1616 = v1364
	goto L13
L423:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1370
	v1375 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1372
	v1379 = F_find_among_b(m, l0, int32(4253488), int32(87))
	mBase = m.M
	v1380 = m.ExcPending
	if v1380 != 0 {
		goto L8
	} else {
		goto L424
	}
L424:
	;
	if v1379 != 0 {
		goto L425
	} else {
		goto L426
	}
L425:
	;
	v1381 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1381
	v1383 = F_slice_del(m, l0)
	mBase = m.M
	v1384 = m.ExcPending
	if v1384 != 0 {
		goto L8
	} else {
		goto L428
	}
L426:
	;
	goto L427
L427:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1375
	goto L41
L428:
	;
	if v1383 < int32(0) {
		v1616 = v1383
		goto L13
	} else {
		goto L429
	}
L429:
	;
	goto L427
L430:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1485
	v1489 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1485 <= v1489 {
		v1566 = v1487
		goto L452
	} else {
		goto L453
	}
L431:
	;
	v1483 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1483
	v1485 = v1483
	v1487 = v1483
	goto L430
L432:
	;
	if v1447 != 0 {
		goto L431
	} else {
		goto L443
	}
L433:
	;
	v1447 = v1443
	goto L432
L434:
	;
	if v1392 <= v1404 {
		goto L436
	} else {
		goto L437
	}
L435:
	;
	v1443 = int32(0)
	goto L433
L436:
	;
	v1447 = int32(-1)
	goto L432
L437:
	;
	goto L438
L438:
	;
	v1416 = int32(1)
	v1417 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1421 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1417+v1392-v1416))))
	if int32(242) < v1421 {
		v1443 = v1416
		goto L433
	} else {
		goto L439
	}
L439:
	;
	v1423 = v1421 - int32(97)
	if v1423 < int32(0) {
		v1443 = v1416
		goto L433
	} else {
		goto L440
	}
L440:
	;
	v1429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1423)>>(uint(int32(3))%32)))+uint32(_consts[1437]))))
	if int32(base.Ui32(v1429)>>(uint(v1423&int32(7))%32))&int32(1) == int32(0) {
		v1443 = v1416
		goto L433
	} else {
		goto L441
	}
L441:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1392 - int32(1)
	goto L442
L442:
	;
	goto L435
L443:
	;
	v1448 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1448
	v1450 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1451 = *(*int32)(unsafe.Add(mBase, uint32(v1450)+8))
	if v1448 < v1451 {
		goto L431
	} else {
		goto L444
	}
L444:
	;
	v1453 = F_slice_del(m, l0)
	mBase = m.M
	v1454 = m.ExcPending
	if v1454 != 0 {
		goto L8
	} else {
		goto L445
	}
L445:
	;
	if v1453 < int32(0) {
		v1616 = v1453
		goto L13
	} else {
		goto L446
	}
L446:
	;
	v1457 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1457
	v1459 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1457 <= v1459 {
		goto L431
	} else {
		goto L447
	}
L447:
	;
	v1461 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1465 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1461+v1457-int32(1)))))
	if v1465 != int32(105) {
		goto L431
	} else {
		goto L448
	}
L448:
	;
	v1469 = v1457 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1469
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1469
	v1472 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1473 = *(*int32)(unsafe.Add(mBase, uint32(v1472)+8))
	if v1457 <= v1473 {
		goto L431
	} else {
		goto L449
	}
L449:
	;
	v1475 = F_slice_del(m, l0)
	mBase = m.M
	v1476 = m.ExcPending
	if v1476 != 0 {
		goto L8
	} else {
		goto L450
	}
L450:
	;
	if v1475 < int32(0) {
		v1616 = v1475
		goto L13
	} else {
		goto L451
	}
L451:
	;
	v1479 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1480 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1485 = v1479
	v1487 = v1480
	goto L430
L452:
	;
	v1567 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1567
	v1570 = v1567
	v1571 = v1566
	goto L471
L453:
	;
	v1491 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1495 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1491+v1485-int32(1)))))
	if v1495 != int32(104) {
		v1566 = v1487
		goto L452
	} else {
		goto L454
	}
L454:
	;
	v1499 = v1485 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1499
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1499
	v1511 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L458
L455:
	;
	v1564 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1566 = v1564
	goto L452
L456:
	;
	if v1554 != 0 {
		goto L455
	} else {
		goto L467
	}
L457:
	;
	v1554 = v1550
	goto L456
L458:
	;
	if v1499 <= v1511 {
		goto L460
	} else {
		goto L461
	}
L459:
	;
	v1550 = int32(0)
	goto L457
L460:
	;
	v1554 = int32(-1)
	goto L456
L461:
	;
	goto L462
L462:
	;
	v1523 = int32(1)
	v1524 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1528 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1524+v1499-v1523))))
	if int32(103) < v1528 {
		v1550 = v1523
		goto L457
	} else {
		goto L463
	}
L463:
	;
	v1530 = v1528 - int32(99)
	if v1530 < int32(0) {
		v1550 = v1523
		goto L457
	} else {
		goto L464
	}
L464:
	;
	v1536 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1530)>>(uint(int32(3))%32)))+uint32(_consts[1438]))))
	if int32(base.Ui32(v1536)>>(uint(v1530&int32(7))%32))&int32(1) == int32(0) {
		v1550 = v1523
		goto L457
	} else {
		goto L465
	}
L465:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1499 - int32(1)
	goto L466
L466:
	;
	goto L459
L467:
	;
	v1555 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1556 = *(*int32)(unsafe.Add(mBase, uint32(v1555)+8))
	v1557 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v1557 < v1556 {
		goto L455
	} else {
		goto L468
	}
L468:
	;
	v1559 = F_slice_del(m, l0)
	mBase = m.M
	v1560 = m.ExcPending
	if v1560 != 0 {
		goto L8
	} else {
		goto L469
	}
L469:
	;
	if v1559 < int32(0) {
		v1616 = v1559
		goto L13
	} else {
		goto L470
	}
L470:
	;
	goto L455
L471:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1570
	if v1571 <= v1570 {
		goto L476
	} else {
		goto L477
	}
L472:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1567
	v1616 = int32(1)
	goto L13
L473:
	;
	goto L472
L474:
	;
	v1611 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1612 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1570 = v1612
	v1571 = v1611
	goto L471
L475:
	;
	if v1604 <= v1603 {
		goto L473
	} else {
		goto L487
	}
L476:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1570
	v1603 = v1570
	v1604 = v1571
	goto L475
L477:
	;
	v1576 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1578 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1576+v1570))))
	switch v1578 - int32(73) {
	case 0, 12:
		goto L478
	default:
		goto L476
	}
L478:
	;
	v1583 = F_find_among(m, l0, int32(4255232), int32(3))
	mBase = m.M
	v1584 = m.ExcPending
	if v1584 != 0 {
		goto L8
	} else {
		goto L479
	}
L479:
	;
	v1585 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1585
	switch v1583 - int32(1) {
	case 0:
		goto L481
	case 1:
		goto L480
	case 2:
		goto L482
	default:
		goto L474
	}
L480:
	;
	v1598 = F_slice_from_s(m, l0, int32(1), int32(2221845))
	mBase = m.M
	v1599 = m.ExcPending
	if v1599 != 0 {
		goto L8
	} else {
		goto L485
	}
L481:
	;
	v1592 = F_slice_from_s(m, l0, int32(1), int32(2221844))
	mBase = m.M
	v1593 = m.ExcPending
	if v1593 != 0 {
		goto L8
	} else {
		goto L483
	}
L482:
	;
	v1589 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1603 = v1585
	v1604 = v1589
	goto L475
L483:
	;
	if int32(0) <= v1592 {
		goto L474
	} else {
		goto L484
	}
L484:
	;
	v1616 = v1592
	goto L13
L485:
	;
	if int32(0) <= v1598 {
		goto L474
	} else {
		goto L486
	}
L486:
	;
	v1616 = v1598
	goto L13
L487:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1603 + int32(1)
	goto L474
}
