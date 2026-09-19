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
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_IncrTupleDescRefCount[0]))
	F_ResourceOwnerEnlarge(m, v3)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v6 + int32(1)
		v11 = *(*int32)(unsafe.Add(mBase, _c_F_IncrTupleDescRefCount[0]))
		F_ResourceOwnerRemember(m, v11, l0, int32(_a_F_IncrTupleDescRefCount_0))
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
	var v108 int64
	_ = v108
	var v109 int64
	_ = v109
	var v112 int64
	_ = v112
	var v113 int64
	_ = v113
	var v114 int64
	_ = v114
	var v117 int64
	_ = v117
	var v118 int64
	_ = v118
	var v119 int64
	_ = v119
	var v124 int64
	_ = v124
	var v129 int64
	_ = v129
	var v134 int64
	_ = v134
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v169 int64
	_ = v169
	var v170 int32
	_ = v170
	var v179 int64
	_ = v179
	var v181 int64
	_ = v181
	var v184 int32
	_ = v184
	var v190 int32
	_ = v190
	v9 = m.G0
	v10 = int32(16)
	v11 = v9 - v10
	m.G0 = v11
	F_gettimeofday(m, v11)
	mBase = m.M
	v14 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	v15 = int64(*(*int32)(unsafe.Add(mBase, uint32(v11)+8)))
	m.G0 = v11 + v10
	v23 = v15 + v14*int64(1000000) - int64(946684800000000)
	goto L1
L1:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_InitProcessGlobals[0])) = v23
	v27 = base.I64_div_s(v23, int64(1000000))
	goto L2
L2:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_InitProcessGlobals[1])) = v27 + int64(946684800)
	v32 = int32(16)
	v33 = int32(0)
	v37 = m.G0
	v39 = v37 - v32
	m.G0 = v39
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = v33
	v45 = F_open(m, int32(_a_F_InitProcessGlobals_0), v33, v39)
	mBase = m.M
	if v45 != int32(-1) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v141 = Fn13966(m, int64(32))
	mBase = m.M
	goto L26
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
	v51 = int32(_a_F_InitProcessGlobals_1)
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
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcessGlobals[2]))
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
	v83 = int32(_a_F_InitProcessGlobals_1)
	v84 = *(*int64)(unsafe.Add(mBase, _c_F_InitProcessGlobals[3]))
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
	v95 = int32(_a_F_InitProcessGlobals_1)
	v97 = int64(*(*int32)(unsafe.Add(mBase, _c_F_InitProcessGlobals[4])))
	v99 = *(*int64)(unsafe.Add(mBase, _c_F_InitProcessGlobals[0]))
	v105 = v97 ^ v99<<(uint(int64(12))%64) ^ int64(base.Ui64(v99)>>(uint(int64(20))%64))
	v108 = v105 + int64(4354685564936845354)
	v109 = int64(30)
	v112 = int64(-4658895280553007687)
	v113 = (int64(base.Ui64(v108)>>(uint(v109)%64)) ^ v108) * v112
	v114 = int64(27)
	v117 = int64(-7723592293110705685)
	v118 = (int64(base.Ui64(v113)>>(uint(v114)%64)) ^ v113) * v117
	v119 = int64(31)
	*(*int64)(unsafe.Add(mBase, _c_F_InitProcessGlobals[5])) = int64(base.Ui64(v118)>>(uint(v119)%64)) ^ v118
	v124 = v105 - int64(7046029254386353131)
	v129 = (int64(base.Ui64(v124)>>(uint(v109)%64)) ^ v124) * v112
	v134 = (int64(base.Ui64(v129)>>(uint(v114)%64)) ^ v129) * v117
	*(*int64)(unsafe.Add(mBase, _c_F_InitProcessGlobals[3])) = int64(base.Ui64(v134)>>(uint(v119)%64)) ^ v134
	goto L25
L20:
	;
	goto L3
L21:
	;
	goto L20
L22:
	;
	v87 = *(*int64)(unsafe.Add(mBase, _c_F_InitProcessGlobals[5]))
	if v87 != int64(0) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_InitProcessGlobals[5])) = int64(1442695040888963407)
	*(*int64)(unsafe.Add(mBase, _c_F_InitProcessGlobals[3])) = int64(6364136223846793005)
	goto L21
L25:
	;
	goto L3
L26:
	;
	v143 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcessGlobals[6]))
	if v143 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	return
L28:
	;
	v147 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcessGlobals[7]))
	*(*int32)(unsafe.Add(mBase, uint32(v147))) = v141
	goto L27
L29:
	;
	goto L30
L30:
	;
	v150 = int32(3)
	if v143 == int32(7) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v155 = v150
	goto L33
L32:
	;
	v155 = int32(1)
	goto L33
L33:
	;
	if v143 == int32(31) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v158 = v150
	goto L36
L35:
	;
	v158 = v155
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InitProcessGlobals[8])) = v158
	v161 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_InitProcessGlobals[9])) = v161
	v164 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcessGlobals[7]))
	if v161 < v143 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v169 = base.I64_extend_i32_u(v141)
	v170 = int32(0)
	goto L40
L38:
	;
	goto L39
L39:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v164)))
	*(*int32)(unsafe.Add(mBase, uint32(v164))) = v190 | int32(1)
	goto L27
L40:
	;
	v179 = v169*int64(6364136223846793005) + int64(1)
	v181 = int64(base.Ui64(v179) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v164+v170<<(uint(int32(2))%32)))) = uint32(v181)
	v184 = v170 + int32(1)
	if v184 != v143 {
		v169 = v179
		v170 = v184
		goto L40
	} else {
		goto L42
	}
L41:
	;
	goto L39
L42:
	;
	goto L41
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
	var v23 int32
	_ = v23
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
		*(*int64)(unsafe.Add(mBase, uint32(v11)+28)) = int64(0)
		*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = l4
		v23 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v11)+36)) = uint8(v23)
		*(*uint8)(unsafe.Add(mBase, uint32(v11)+60)) = uint8(v23)
		*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = l3
		*(*uint8)(unsafe.Add(mBase, uint32(v11)+52)) = uint8(v23)
		*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = l2
		*(*uint8)(unsafe.Add(mBase, uint32(v11)+44)) = uint8(v23)
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
							F_errmsg_internal(m, int32(_a_F_InputFunctionCallSafe_0), v11)
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_InputFunctionCallSafe_1), int32(1618), int32(_a_F_InputFunctionCallSafe_2))
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
							F_errmsg_internal(m, int32(_a_F_InputFunctionCallSafe_3), v9+int32(-48))
							mBase = m.M
							v88 = m.ExcPending
							if v88 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_InputFunctionCallSafe_1), int32(1624), int32(_a_F_InputFunctionCallSafe_2))
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
								F_errmsg_internal(m, int32(_a_F_InputFunctionCallSafe_0), v11)
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_InputFunctionCallSafe_1), int32(1618), int32(_a_F_InputFunctionCallSafe_2))
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
								F_errmsg_internal(m, int32(_a_F_InputFunctionCallSafe_3), v9+int32(-48))
								mBase = m.M
								v88 = m.ExcPending
								if v88 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_InputFunctionCallSafe_1), int32(1624), int32(_a_F_InputFunctionCallSafe_2))
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
									F_errmsg_internal(m, int32(_a_F_InputFunctionCallSafe_0), v11)
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_InputFunctionCallSafe_1), int32(1618), int32(_a_F_InputFunctionCallSafe_2))
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
									F_errmsg_internal(m, int32(_a_F_InputFunctionCallSafe_3), v9+int32(-48))
									mBase = m.M
									v88 = m.ExcPending
									if v88 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_InputFunctionCallSafe_1), int32(1624), int32(_a_F_InputFunctionCallSafe_2))
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
			*(*int64)(unsafe.Add(mBase, uint32(v11)+28)) = int64(0)
			*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = l4
			v23 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v11)+36)) = uint8(v23)
			*(*uint8)(unsafe.Add(mBase, uint32(v11)+60)) = uint8(v23)
			*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = l3
			*(*uint8)(unsafe.Add(mBase, uint32(v11)+52)) = uint8(v23)
			*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = l2
			*(*uint8)(unsafe.Add(mBase, uint32(v11)+44)) = uint8(v23)
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
								F_errmsg_internal(m, int32(_a_F_InputFunctionCallSafe_0), v11)
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_InputFunctionCallSafe_1), int32(1618), int32(_a_F_InputFunctionCallSafe_2))
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
								F_errmsg_internal(m, int32(_a_F_InputFunctionCallSafe_3), v9+int32(-48))
								mBase = m.M
								v88 = m.ExcPending
								if v88 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_InputFunctionCallSafe_1), int32(1624), int32(_a_F_InputFunctionCallSafe_2))
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
									F_errmsg_internal(m, int32(_a_F_InputFunctionCallSafe_0), v11)
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_InputFunctionCallSafe_1), int32(1618), int32(_a_F_InputFunctionCallSafe_2))
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
									F_errmsg_internal(m, int32(_a_F_InputFunctionCallSafe_3), v9+int32(-48))
									mBase = m.M
									v88 = m.ExcPending
									if v88 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_InputFunctionCallSafe_1), int32(1624), int32(_a_F_InputFunctionCallSafe_2))
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
										F_errmsg_internal(m, int32(_a_F_InputFunctionCallSafe_0), v11)
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_InputFunctionCallSafe_1), int32(1618), int32(_a_F_InputFunctionCallSafe_2))
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
										F_errmsg_internal(m, int32(_a_F_InputFunctionCallSafe_3), v9+int32(-48))
										mBase = m.M
										v88 = m.ExcPending
										if v88 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_InputFunctionCallSafe_1), int32(1624), int32(_a_F_InputFunctionCallSafe_2))
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
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
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
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v156 int32
	_ = v156
	var v158 int64
	_ = v158
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v175 int32
	_ = v175
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v221 int32
	_ = v221
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v290 int32
	_ = v290
	var v306 int64
	_ = v306
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v318 int64
	_ = v318
	var v319 int64
	_ = v319
	var v323 int64
	_ = v323
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v374 int64
	_ = v374
	var v378 int32
	_ = v378
	var v390 int64
	_ = v390
	var v394 int32
	_ = v394
	var v406 int32
	_ = v406
	var v411 int64
	_ = v411
	var v415 int64
	_ = v415
	var v420 int32
	_ = v420
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
	v30 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_IssuePendingWritebacks[0])))
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
	v57 = v48
	v59 = int32(0)
	goto L11
L9:
	;
	v290 = v48
	goto L10
L10:
	;
	v306 = int64(0)
	v310 = m.G0
	v312 = v310 - int32(16)
	m.G0 = v312
	if v44 != v306 {
		goto L47
	} else {
		goto L48
	}
L11:
	;
	v73 = v27 + v59*int32(20)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+16))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v73)+12))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v73)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v76
	v82 = int32(1)
	v84 = v59 + v82
	if v57 <= v84 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v290 = v282
	goto L10
L13:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v23)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v156
	v158 = *(*int64)(unsafe.Add(mBase, uint32(v23)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v23)+8)) = v158
	v163 = F_smgropen(m, v23+int32(8), int32(-1))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L28
	} else {
		goto L29
	}
L14:
	;
	v143 = v82
	v144 = v84
	goto L13
L15:
	;
	goto L16
L16:
	;
	v94 = v73
	v97 = v82
	v100 = int32(0)
	goto L17
L17:
	;
	v110 = v84 + v100
	v113 = v27 + v110*int32(20)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)+8))
	if v78 != v114 {
		v143 = v97
		v144 = v110
		goto L13
	} else {
		goto L19
	}
L18:
	;
	v143 = v132
	v144 = v57
	goto L13
L19:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v113)+4))
	if v77 != v116 {
		v143 = v97
		v144 = v110
		goto L13
	} else {
		goto L20
	}
L20:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	if v76 != v118 {
		v143 = v97
		v144 = v110
		goto L13
	} else {
		goto L21
	}
L21:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v94)+12))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v113)+12))
	if v120 != v121 {
		v143 = v97
		v144 = v110
		goto L13
	} else {
		goto L22
	}
L22:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v94)+16))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v113)+16))
	if v123 != v124 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	if v123+int32(1) != v124 {
		v143 = v97
		v144 = v110
		goto L13
	} else {
		goto L26
	}
L24:
	;
	v131 = v94
	v132 = v97
	goto L25
L25:
	;
	v134 = v100 + int32(1)
	if v134 != v57+(v59^int32(-1)) {
		v94 = v131
		v97 = v132
		v100 = v134
		goto L17
	} else {
		goto L27
	}
L26:
	;
	v131 = v113
	v132 = v97 + int32(1)
	goto L25
L27:
	;
	goto L18
L28:
	;
	return
L29:
	;
	v165 = int32(_a_F_IssuePendingWritebacks_0)
	v167 = *(*int32)(unsafe.Add(mBase, _c_F_IssuePendingWritebacks[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_IssuePendingWritebacks[1])) = v167 + int32(1)
	if v143 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v276 = int32(_a_F_IssuePendingWritebacks_0)
	v278 = *(*int32)(unsafe.Add(mBase, _c_F_IssuePendingWritebacks[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_IssuePendingWritebacks[1])) = v278 - int32(1)
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v144 < v282 {
		v57 = v282
		v59 = v144
		goto L11
	} else {
		goto L45
	}
L31:
	;
	v175 = v163 + v75<<(uint(int32(2))%32)
	v187 = v143
	v189 = v74
	goto L32
L32:
	;
	v201 = int32(base.Ui32(v189) >> (uint(int32(17)) % 32))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v175+int32(40))))
	if base.Ui32(v202) <= base.Ui32(v201) {
		goto L30
	} else {
		goto L34
	}
L33:
	;
	goto L30
L34:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v175+int32(56))))
	if v204 == int32(0) {
		goto L30
	} else {
		goto L35
	}
L35:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v204+v201<<(uint(int32(3))%32))))
	v211 = int32(_a_F_IssuePendingWritebacks_1)
	if base.Ui32(v187+v189-int32(1)^v189) < base.Ui32(v211) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v255 = v187 - v221
	if v255 != 0 {
		v187 = v255
		v189 = v221 + v189
		goto L32
	} else {
		goto L44
	}
L37:
	;
	v221 = v187
	goto L39
L38:
	;
	v221 = v211 - v189&int32(_a_F_IssuePendingWritebacks_2)
	goto L39
L39:
	;
	if base.I64_extend_i32_u(v221)<<(uint(int64(13))%64) == int64(0) {
		goto L36
	} else {
		goto L40
	}
L40:
	;
	v228 = *(*int32)(unsafe.Add(mBase, _c_F_IssuePendingWritebacks[2]))
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228+v210*int32(48))+37)))
	if v232&int32(64) != 0 {
		goto L36
	} else {
		goto L41
	}
L41:
	;
	v235 = F_FileAccess(m, v210)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L28
	} else {
		goto L42
	}
L42:
	;
	if v235 < int32(0) {
		goto L36
	} else {
		goto L43
	}
L43:
	;
	v239 = int32(_a_F_IssuePendingWritebacks_3)
	v240 = *(*int32)(unsafe.Add(mBase, _c_F_IssuePendingWritebacks[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v240))) = int32(167772178)
	v244 = *(*int32)(unsafe.Add(mBase, _c_F_IssuePendingWritebacks[2]))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v244+v210*int32(48))))
	v249 = F_fsync(m, v248)
	mBase = m.M
	v251 = *(*int32)(unsafe.Add(mBase, _c_F_IssuePendingWritebacks[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v251))) = int32(0)
	goto L36
L44:
	;
	goto L33
L45:
	;
	goto L12
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(0)
	goto L3
L47:
	;
	F___clock_gettime(m, int32(1), v312)
	mBase = m.M
	v318 = int64(*(*int32)(unsafe.Add(mBase, uint32(v312)+8)))
	v319 = *(*int64)(unsafe.Add(mBase, uint32(v312)))
	v323 = v318 + (v319*int64(1000000000) - v44)
	goto L51
L48:
	;
	goto L49
L49:
	;
	v406 = l1 << (uint(int32(6)) % 32)
	v411 = *(*int64)(unsafe.Add(mBase, uint32(v406)+uint32(_c_F_IssuePendingWritebacks[4])))
	*(*int64)(unsafe.Add(mBase, uint32(v406)+uint32(_c_F_IssuePendingWritebacks[4]))) = v411 + base.I64_extend_i32_u(v290)
	v415 = *(*int64)(unsafe.Add(mBase, uint32(v406)+uint32(_c_F_IssuePendingWritebacks[5])))
	*(*int64)(unsafe.Add(mBase, uint32(v406)+uint32(_c_F_IssuePendingWritebacks[5]))) = v415 + v306
	F_pgstat_count_backend_io_op(m, int32(0), l1, int32(4), v290, v306)
	mBase = m.M
	v420 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_IssuePendingWritebacks[6])) = uint8(v420)
	*(*uint8)(unsafe.Add(mBase, _c_F_IssuePendingWritebacks[7])) = uint8(v420)
	m.G0 = v312 + int32(16)
	goto L46
L50:
	;
	v367 = int32(0)
	v369 = l1 << (uint(int32(6)) % 32)
	v374 = *(*int64)(unsafe.Add(mBase, uint32(v369)+uint32(_c_F_IssuePendingWritebacks[8])))
	*(*int64)(unsafe.Add(mBase, uint32(v369)+uint32(_c_F_IssuePendingWritebacks[8]))) = v374 + v323
	v378 = *(*int32)(unsafe.Add(mBase, _c_F_IssuePendingWritebacks[9]))
	if base.B2i32(base.Ui32(int32(16)) < base.Ui32(v378))|base.B2i32(int32(1)<<(uint(v378)%32)&int32(_a_F_IssuePendingWritebacks_4) == v367) == v367 {
		goto L60
	} else {
		goto L61
	}
L51:
	;
	goto L53
L53:
	;
	goto L54
L54:
	;
	goto L50
L60:
	;
	v390 = *(*int64)(unsafe.Add(mBase, uint32(v369)+uint32(_c_F_IssuePendingWritebacks[10])))
	*(*int64)(unsafe.Add(mBase, uint32(v369)+uint32(_c_F_IssuePendingWritebacks[10]))) = v390 + v323
	v394 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_IssuePendingWritebacks[6])) = uint8(v394)
	*(*uint8)(unsafe.Add(mBase, _c_F_IssuePendingWritebacks[11])) = uint8(v394)
	goto L62
L61:
	;
	goto L62
L62:
	;
	goto L49
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
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13991(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
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
			F_errmsg(m, int32(_a_F_icu_validate_locale_0), int32(0))
			v11 = m.ExcPending
			if v11 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_icu_validate_locale_1), int32(1674), int32(_a_F_icu_validate_locale_2))
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var __phi58 int32
	_ = __phi58
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
	var v66 int32
	_ = v66
	var __phi66 int32
	_ = __phi66
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v106 int64
	_ = v106
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v149 int64
	_ = v149
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int64
	_ = v210
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
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
	v224 = m.ExcPending
	if v224 != 0 {
		goto L40
	} else {
		goto L53
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
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v32 = v3
	v33 = int32(0)
	goto L7
L6:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+56))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+22)))
	v32 = v27 + v28
	v33 = int32(1)
	goto L7
L7:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if int32(0) < v34 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v52 = int32(48)
	__phi58 = base.B2i32(int32(0) < v34)
	__phi59 = v3
	__phi60 = v47
	__phi61 = v32
	__phi62 = v48
	__phi63 = v49
	__phi64 = v33
	__phi66 = v3
	v58 = __phi58
	v59 = __phi59
	v60 = __phi60
	v61 = __phi61
	v62 = __phi62
	v63 = __phi63
	v64 = __phi64
	v66 = __phi66
	goto L13
L9:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+56))
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+22)))
	v47 = v40 + v41
	v48 = base.B2i32(v32 != int32(0))
	v49 = int32(1)
	goto L8
L10:
	;
	goto L11
L11:
	;
	if v32 == int32(0) {
		v241 = v3
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v47 = v3
	v48 = int32(1)
	v49 = v3
	goto L8
L13:
	;
	v72 = v61 + int32(12)
	v74 = base.B2i32(v59 != int32(0))
	v78 = v58
	v80 = v60
	v83 = v63
	goto L16
L15:
	;
	v181 = F_palloc(m, int32(24))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L40
	} else {
		goto L41
	}
L16:
	;
	if v62&v74 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v241 = v66
	goto L1
L18:
	;
	goto L17
L19:
	;
	if v78&v74 == int32(0) {
		v179 = v78
		goto L15
	} else {
		goto L30
	}
L20:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	if v93 != v94 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	if v96 != v97 {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61)+16)))
	if base.Ui32((v99-int32(1))&int32(_a_F_identify_opfamily_groups_0)) <= base.Ui32(int32(62)) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v106 = *(*int64)(unsafe.Add(mBase, uint32(v59)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v59)+8)) = v106 | int64(1)<<(uint(base.I64_extend_i32_u(v99))%64)
	goto L25
L24:
	;
	goto L25
L25:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v64 < v113 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0+v52+v64<<(uint(int32(2))%32))))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)+56))
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+22)))
	v125 = v119 + v120
	v126 = v64 + int32(1)
	goto L28
L27:
	;
	v125 = int32(0)
	v126 = v64
	goto L28
L28:
	;
	v127 = int32(0)
	if v80|v125 != 0 {
		__phi58 = base.B2i32(v80 != v127)
		__phi60 = v80
		__phi61 = v125
		__phi62 = base.B2i32(v125 != v127)
		__phi63 = v83
		__phi64 = v126
		v58 = __phi58
		v60 = __phi60
		v61 = __phi61
		v62 = __phi62
		v63 = __phi63
		v64 = __phi64
		goto L13
	} else {
		goto L29
	}
L29:
	;
	goto L18
L30:
	;
	v135 = int32(1)
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v80)+8))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	if v136 != v137 {
		v179 = v135
		goto L15
	} else {
		goto L31
	}
L31:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	if v139 != v140 {
		v179 = v135
		goto L15
	} else {
		goto L32
	}
L32:
	;
	v142 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v80)+16)))
	if base.Ui32((v142-int32(1))&int32(_a_F_identify_opfamily_groups_0)) <= base.Ui32(int32(62)) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v149 = *(*int64)(unsafe.Add(mBase, uint32(v59)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v59)+16)) = v149 | int64(1)<<(uint(base.I64_extend_i32_u(v142))%64)
	goto L35
L34:
	;
	goto L35
L35:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v83 < v156 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l1+v52+v83<<(uint(int32(2))%32))))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v161)+56))
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+22)))
	v168 = v162 + v163
	v169 = v83 + int32(1)
	goto L38
L37:
	;
	v168 = int32(0)
	v169 = v83
	goto L38
L38:
	;
	v171 = base.B2i32(v168 != int32(0))
	if v171|v62 != 0 {
		v78 = v171
		v80 = v168
		v83 = v169
		goto L16
	} else {
		goto L39
	}
L39:
	;
	goto L18
L40:
	;
	return int32(0)
L41:
	;
	if v62 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L42:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
	v210 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v181)+8)) = v210
	*(*int32)(unsafe.Add(mBase, uint32(v181)+4)) = v209
	*(*int64)(unsafe.Add(mBase, uint32(v181)+16)) = v210
	v215 = int32(0)
	v219 = F_lappend(m, v66, v181)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L40
	} else {
		goto L52
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v181))) = v188
	v208 = v72
	goto L42
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v181))) = v199
	v208 = v80 + int32(12)
	goto L42
L45:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v80)+8))
	v199 = v187
	goto L44
L46:
	;
	goto L47
L47:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
	if v179&int32(1) == int32(0) {
		goto L43
	} else {
		goto L48
	}
L48:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v80)+8))
	if base.Ui32(v188) < base.Ui32(v193) {
		goto L43
	} else {
		goto L49
	}
L49:
	;
	if v193 != v188 {
		v199 = v193
		goto L44
	} else {
		goto L50
	}
L50:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
	if base.Ui32(v196) < base.Ui32(v197) {
		goto L43
	} else {
		goto L51
	}
L51:
	;
	v199 = v188
	goto L44
L52:
	;
	__phi58 = base.B2i32(v80 != v215)
	__phi59 = v181
	__phi60 = v80
	__phi62 = base.B2i32(v61 != v215)
	__phi63 = v83
	__phi66 = v219
	v58 = __phi58
	v59 = __phi59
	v60 = __phi60
	v62 = __phi62
	v63 = __phi63
	v66 = __phi66
	goto L13
L53:
	;
	F_errmsg_internal(m, int32(_a_F_identify_opfamily_groups_1), int32(0))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L40
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_identify_opfamily_groups_2), int32(54), int32(_a_F_identify_opfamily_groups_3))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L40
	} else {
		goto L55
	}
L55:
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
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
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
							F_errmsg(m, int32(_a_F_idx_0), int32(0))
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_idx_1), int32(267), int32(_a_F_idx_2))
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
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
					v15 = F_ArrayGetNItemsSafe(m, v12, v5+int32(16))
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
			v15 = F_ArrayGetNItemsSafe(m, v12, v5+int32(16))
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
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
	var v124 int32
	_ = v124
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v345 int32
	_ = v345
	var v358 int32
	_ = v358
	var v363 int32
	_ = v363
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v420 int32
	_ = v420
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
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v517 int32
	_ = v517
	var v524 int32
	_ = v524
	v2 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v2
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v19 < v10 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if int32(0) <= v59 {
		goto L16
	} else {
		goto L17
	}
L2:
	;
	v21 = v10
	goto L4
L3:
	;
	v21 = v19
	goto L4
L4:
	;
	v28 = v10
	goto L6
L5:
	;
	v59 = v39
	goto L1
L6:
	;
	if v28 == v21 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v59 = int32(-1)
	goto L1
L9:
	;
	goto L10
L10:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32+v28))))
	if int32(117) < v34 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v51 = v28 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v51
	v28 = v51
	goto L6
L12:
	;
	v36 = v34 - int32(97)
	if v36 < int32(0) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v39 = int32(1)
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v36)>>(uint(int32(3))%32)))+uint32(_c_F_indonesian_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v43)>>(uint(v36&int32(7))%32))&v39 != 0 {
		goto L5
	} else {
		goto L14
	}
L14:
	;
	goto L11
L16:
	;
	v63 = v59
	goto L19
L17:
	;
	goto L18
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v10
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)+4))
	if v135 < int32(3) {
		v524 = v2
		goto L37
	} else {
		goto L38
	}
L19:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v68 + v63
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v71)+4)) = v72 + int32(1)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v84 < v83 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L18
L21:
	;
	if int32(0) <= v124 {
		v63 = v124
		goto L19
	} else {
		goto L36
	}
L22:
	;
	v86 = v83
	goto L24
L23:
	;
	v86 = v84
	goto L24
L24:
	;
	v93 = v83
	goto L26
L25:
	;
	v124 = v104
	goto L21
L26:
	;
	if v93 == v86 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v124 = int32(-1)
	goto L21
L29:
	;
	goto L30
L30:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97+v93))))
	if int32(117) < v99 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v116 = v93 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v116
	v93 = v116
	goto L26
L32:
	;
	v101 = v99 - int32(97)
	if v101 < int32(0) {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v104 = int32(1)
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v101)>>(uint(int32(3))%32)))+uint32(_c_F_indonesian_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v108)>>(uint(v101&int32(7))%32))&v104 != 0 {
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
	return v524
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v134))) = int32(0)
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v140
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v142
	if v142-int32(2) <= v140 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v181
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v178)+4))
	if v183 < int32(3) {
		goto L51
	} else {
		goto L52
	}
L40:
	;
	v178 = v134
	v180 = int32(0)
	goto L39
L41:
	;
	goto L42
L42:
	;
	v149 = int32(0)
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150+v142-int32(1)))))
	switch v154 - int32(104) {
	case 0, 6:
		goto L43
	default:
		v178 = v134
		v180 = v149
		goto L39
	}
L43:
	;
	v159 = F_find_among_b(m, l0, int32(_a_F_indonesian_ISO_8859_1_stem_0), int32(3))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	return int32(0)
L45:
	;
	if v159 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v178 = v165
	v180 = v149
	goto L39
L47:
	;
	goto L48
L48:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v166
	v168 = F_slice_del(m, l0)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L44
	} else {
		goto L49
	}
L49:
	;
	if v168 < int32(0) {
		v524 = v168
		goto L37
	} else {
		goto L50
	}
L50:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v172)+4))
	v174 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v172)+4)) = v173 - v174
	v178 = v172
	v180 = v174
	goto L39
L51:
	;
	return int32(0)
L52:
	;
	goto L53
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v181
	v190 = v181 - int32(1)
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v190 <= v191 {
		v219 = v178
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v221
	v223 = int32(0)
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v219)+4))
	if v224 < int32(3) {
		v524 = v223
		goto L37
	} else {
		goto L63
	}
L55:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193+v190))))
	if base.B2i32(v195 != int32(117))&base.B2i32(v195 != int32(97)) != 0 {
		v219 = v178
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v203 = F_find_among_b(m, l0, int32(_a_F_indonesian_ISO_8859_1_stem_1), int32(3))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L44
	} else {
		goto L57
	}
L57:
	;
	if v203 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v219 = v207
	goto L54
L59:
	;
	goto L60
L60:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v208
	v210 = F_slice_del(m, l0)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L44
	} else {
		goto L61
	}
L61:
	;
	if v210 < int32(0) {
		v524 = v210
		goto L37
	} else {
		goto L62
	}
L62:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v214)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v214)+4)) = v215 - int32(1)
	v219 = v214
	goto L54
L63:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v227
	v230 = v227 + int32(1)
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v231 <= v230 {
		v457 = v223
		goto L66
	} else {
		goto L67
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v221
	v524 = int32(1)
	goto L37
L65:
	;
	if v462 != 0 {
		goto L129
	} else {
		goto L130
	}
L66:
	;
	v462 = v457
	goto L65
L67:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233+v230))))
	switch v235 - int32(101) {
	case 0, 4:
		goto L68
	default:
		v457 = v223
		goto L66
	}
L68:
	;
	v240 = F_find_among(m, l0, int32(_a_F_indonesian_ISO_8859_1_stem_2), int32(12))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L44
	} else {
		goto L69
	}
L69:
	;
	if v240 == int32(0) {
		v457 = v223
		goto L66
	} else {
		goto L70
	}
L70:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v244
	v246 = int32(1)
	switch v240 - v246 {
	case 0:
		goto L76
	case 1:
		goto L75
	case 2:
		goto L74
	case 3:
		goto L73
	case 4:
		goto L72
	case 5:
		goto L71
	default:
		v457 = v246
		goto L66
	}
L71:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v378))) = int32(3)
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v378)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v378)+4)) = v381 - int32(1)
	v385 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v395 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v395 < v385 {
		goto L108
	} else {
		goto L109
	}
L72:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v304 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v303))) = v304
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v303)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v303)+4)) = v306 - v304
	v310 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v320 < v310 {
		goto L86
	} else {
		goto L87
	}
L73:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v288))) = int32(3)
	v293 = F_slice_from_s(m, l0, int32(1), int32(_a_F_indonesian_ISO_8859_1_stem_3))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L44
	} else {
		goto L83
	}
L74:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v274 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v273))) = v274
	v278 = F_slice_from_s(m, l0, v274, int32(_a_F_indonesian_ISO_8859_1_stem_4))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L44
	} else {
		goto L81
	}
L75:
	;
	v261 = F_slice_del(m, l0)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L44
	} else {
		goto L79
	}
L76:
	;
	v249 = F_slice_del(m, l0)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L44
	} else {
		goto L77
	}
L77:
	;
	if v249 < int32(0) {
		v457 = v249
		goto L66
	} else {
		goto L78
	}
L78:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v254 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v253))) = v254
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v253)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v253)+4)) = v256 - v254
	v462 = v254
	goto L65
L79:
	;
	if v261 < int32(0) {
		v457 = v261
		goto L66
	} else {
		goto L80
	}
L80:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v265))) = int32(3)
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v265)+4))
	v269 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v265)+4)) = v268 - v269
	v462 = v269
	goto L65
L81:
	;
	if v278 < int32(0) {
		v457 = v278
		goto L66
	} else {
		goto L82
	}
L82:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v282)+4))
	v284 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v282)+4)) = v283 - v284
	v462 = v284
	goto L65
L83:
	;
	if v293 < int32(0) {
		v457 = v293
		goto L66
	} else {
		goto L84
	}
L84:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v297)+4))
	v299 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v297)+4)) = v298 - v299
	v462 = v299
	goto L65
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v310
	if v363 == int32(0) {
		goto L100
	} else {
		goto L101
	}
L86:
	;
	v322 = v310
	goto L88
L87:
	;
	v322 = v320
	goto L88
L88:
	;
	goto L90
L89:
	;
	v363 = v358
	goto L85
L90:
	;
	if v310 == v322 {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	v358 = int32(0)
	goto L89
L92:
	;
	v363 = int32(-1)
	goto L85
L93:
	;
	goto L94
L94:
	;
	v334 = int32(1)
	v335 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335+v310))))
	if int32(117) < v337 {
		v358 = v334
		goto L89
	} else {
		goto L95
	}
L95:
	;
	v339 = v337 - int32(97)
	if v339 < int32(0) {
		v358 = v334
		goto L89
	} else {
		goto L96
	}
L96:
	;
	v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v339)>>(uint(int32(3))%32)))+uint32(_c_F_indonesian_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v345)>>(uint(v339&int32(7))%32))&int32(1) == int32(0) {
		v358 = v334
		goto L89
	} else {
		goto L97
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v310 + int32(1)
	goto L98
L98:
	;
	goto L91
L99:
	;
	v462 = v377
	goto L65
L100:
	;
	v369 = F_slice_from_s(m, l0, int32(1), int32(_a_F_indonesian_ISO_8859_1_stem_5))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L44
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	v373 = F_slice_del(m, l0)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L44
	} else {
		goto L105
	}
L103:
	;
	if v369 < int32(0) {
		v377 = v369
		goto L99
	} else {
		goto L104
	}
L104:
	;
	v457 = v246
	goto L66
L105:
	;
	if int32(0) <= v373 {
		v457 = v246
		goto L66
	} else {
		goto L106
	}
L106:
	;
	v377 = v373
	goto L99
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v385
	if v438 == int32(0) {
		goto L122
	} else {
		goto L123
	}
L108:
	;
	v397 = v385
	goto L110
L109:
	;
	v397 = v395
	goto L110
L110:
	;
	goto L112
L111:
	;
	v438 = v433
	goto L107
L112:
	;
	if v385 == v397 {
		goto L114
	} else {
		goto L115
	}
L113:
	;
	v433 = int32(0)
	goto L111
L114:
	;
	v438 = int32(-1)
	goto L107
L115:
	;
	goto L116
L116:
	;
	v409 = int32(1)
	v410 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v410+v385))))
	if int32(117) < v412 {
		v433 = v409
		goto L111
	} else {
		goto L117
	}
L117:
	;
	v414 = v412 - int32(97)
	if v414 < int32(0) {
		v433 = v409
		goto L111
	} else {
		goto L118
	}
L118:
	;
	v420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v414)>>(uint(int32(3))%32)))+uint32(_c_F_indonesian_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v420)>>(uint(v414&int32(7))%32))&int32(1) == int32(0) {
		v433 = v409
		goto L111
	} else {
		goto L119
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v385 + int32(1)
	goto L120
L120:
	;
	goto L113
L121:
	;
	v457 = v454
	goto L66
L122:
	;
	v442 = int32(1)
	v445 = F_slice_from_s(m, l0, v442, int32(_a_F_indonesian_ISO_8859_1_stem_6))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L44
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	v450 = F_slice_del(m, l0)
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L44
	} else {
		goto L127
	}
L125:
	;
	if v445 < int32(0) {
		v454 = v445
		goto L121
	} else {
		goto L126
	}
L126:
	;
	v457 = v442
	goto L66
L127:
	;
	if int32(0) <= v450 {
		v457 = int32(1)
		goto L66
	} else {
		goto L128
	}
L128:
	;
	v454 = v450
	goto L121
L129:
	;
	if v462 < int32(0) {
		v524 = v462
		goto L37
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v221
	v494 = F_r_remove_second_order_prefix_1(m, l0)
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L44
	} else {
		goto L146
	}
L132:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v465)+4))
	if v466 < int32(3) {
		goto L64
	} else {
		goto L133
	}
L133:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v469
	v471 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v471
	v473 = F_r_remove_suffix_1(m, l0)
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L44
	} else {
		goto L134
	}
L134:
	;
	if v473 == int32(0) {
		goto L64
	} else {
		goto L135
	}
L135:
	;
	if v473 < int32(0) {
		v524 = v473
		goto L37
	} else {
		goto L136
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v469
	v480 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v480)+4))
	if v481 < int32(3) {
		goto L64
	} else {
		goto L137
	}
L137:
	;
	v484 = F_r_remove_second_order_prefix_1(m, l0)
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L44
	} else {
		goto L138
	}
L138:
	;
	if int32(0) <= v484 {
		goto L64
	} else {
		goto L139
	}
L139:
	;
	if v484 < int32(0) {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v490 = v484
	goto L142
L141:
	;
	v490 = v180
	goto L142
L142:
	;
	if v484 != 0 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v491 = v490
	goto L145
L144:
	;
	v491 = v180
	goto L145
L145:
	;
	return v491
L146:
	;
	if v494 < int32(0) {
		v524 = v494
		goto L37
	} else {
		goto L147
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v221
	v499 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v499)+4))
	if v500 < int32(3) {
		goto L64
	} else {
		goto L148
	}
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v221
	v504 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v504
	v506 = F_r_remove_suffix_1(m, l0)
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L44
	} else {
		goto L149
	}
L149:
	;
	if base.Ui32(int32(6)) < base.Ui32(int32(base.Ui32(v506)>>(uint(int32(31))%32))-int32(1)) {
		goto L64
	} else {
		goto L150
	}
L150:
	;
	if int32(0) <= v506 {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v517 = int32(1)
	goto L153
L152:
	;
	v517 = v506
	goto L153
L153:
	;
	return v517
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
							F_errmsg(m, int32(_a_F_inetand_0), int32(0))
							mBase = m.M
							v154 = m.ExcPending
							if v154 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_inetand_1), int32(1858), int32(_a_F_inetand_2))
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
							F_errmsg(m, int32(_a_F_inetor_0), int32(0))
							mBase = m.M
							v154 = m.ExcPending
							if v154 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_inetor_1), int32(1890), int32(_a_F_inetor_2))
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
	v140 = F_pg_sprintf(m, v131, int32(_a_F_infix_2_0), int32(0))
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
	v188 = F_pg_sprintf(m, v179, int32(_a_F_infix_2_1), int32(0))
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
	v241 = F_pg_sprintf(m, v233, int32(_a_F_infix_2_0), int32(0))
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
	v329 = F_pg_sprintf(m, v320, int32(_a_F_infix_2_2), v13)
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
	v379 = F_pg_sprintf(m, v370, int32(_a_F_infix_2_1), int32(0))
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
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 float64
	_ = v19
	var v20 float64
	_ = v20
	var v22 float64
	_ = v22
	var v23 int32
	_ = v23
	var v26 float64
	_ = v26
	var v27 float64
	_ = v27
	var v30 float64
	_ = v30
	var v33 float64
	_ = v33
	var v34 float64
	_ = v34
	var v36 float64
	_ = v36
	var v38 float64
	_ = v38
	var v41 float64
	_ = v41
	var v43 float64
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 float64
	_ = v49
	var v51 int32
	_ = v51
	var v57 float64
	_ = v57
	var v61 float64
	_ = v61
	var v64 float64
	_ = v64
	var v66 float64
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v88 int32
	_ = v88
	var v93 float64
	_ = v93
	var v95 int32
	_ = v95
	var v99 float64
	_ = v99
	var v100 float64
	_ = v100
	var v103 float64
	_ = v103
	var v104 int32
	_ = v104
	var v108 float64
	_ = v108
	var v109 float64
	_ = v109
	var v111 int32
	_ = v111
	var v116 float64
	_ = v116
	var v117 float64
	_ = v117
	var v120 float64
	_ = v120
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v142 float64
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 float64
	_ = v152
	var v154 float64
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v165 int32
	_ = v165
	var v175 float64
	_ = v175
	var v177 int32
	_ = v177
	var v181 float64
	_ = v181
	var v182 float64
	_ = v182
	var v185 float64
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 float64
	_ = v213
	var v215 float64
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v226 int32
	_ = v226
	var v234 float64
	_ = v234
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v269 int32
	_ = v269
	var v275 float64
	_ = v275
	var v277 float64
	_ = v277
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v288 int32
	_ = v288
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v304 int32
	_ = v304
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v334 int32
	_ = v334
	var v342 int32
	_ = v342
	var v353 int32
	_ = v353
	var v359 int32
	_ = v359
	var v367 int32
	_ = v367
	var v371 float64
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v382 float64
	_ = v382
	var v384 float64
	_ = v384
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v398 float64
	_ = v398
	var v404 float64
	_ = v404
	var v406 float64
	_ = v406
	var v418 int32
	_ = v418
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v19 = *(*float64)(unsafe.Add(mBase, uint32(l3)+32))
	v20 = *(*float64)(unsafe.Add(mBase, uint32(l2)+32))
	v22 = *(*float64)(unsafe.Add(mBase, _c_F_initial_cost_hashjoin[0]))
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v26 = base.F64_convert_i32_s(v23)
	goto L3
L2:
	;
	v26 = float64(0)
	goto L3
L3:
	;
	v27 = base.F64_mul(v22, v26)
	v30 = *(*float64)(unsafe.Add(mBase, _c_F_initial_cost_hashjoin[1]))
	v33 = *(*float64)(unsafe.Add(mBase, uint32(l2)+56))
	v34 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
	v36 = float64(0)
	v38 = base.F64_add(base.F64_mul(v27, v20), base.F64_add(base.F64_sub(v33, v34), v36))
	v41 = *(*float64)(unsafe.Add(mBase, uint32(l3)+56))
	v43 = base.F64_add(base.F64_mul(base.F64_add(v27, v30), v19), base.F64_add(base.F64_add(v34, v36), v41))
	v45 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_initial_cost_hashjoin[2])))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l3)+40))
	if l4 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	v49 = base.F64_convert_i32_s(v48)
	v51 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_initial_cost_hashjoin[3])))
	if v51 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v66 = v19
	goto L6
L6:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+32))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	v77 = v17 + int32(4)
	v88 = (v69 + int32(7)) & int32(-8)
	v93 = *(*float64)(unsafe.Add(mBase, _c_F_initial_cost_hashjoin[4]))
	v95 = *(*int32)(unsafe.Add(mBase, _c_F_initial_cost_hashjoin[5]))
	v99 = base.F64_mul(base.F64_mul(v93, base.F64_convert_i32_s(v95)), float64(1024))
	v100 = float64(4.294967295e+09)
	if base.F64_lt(v99, v100) != 0 {
		goto L14
	} else {
		goto L15
	}
L7:
	;
	v57 = base.F64_add(base.F64_mul(v49, float64(-0.3)), float64(1))
	if base.F64_gt(v57, float64(0)) != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v64 = v49
	goto L9
L9:
	;
	v66 = base.F64_mul(v19, v64)
	goto L6
L10:
	;
	v61 = v57
	goto L12
L11:
	;
	v61 = math.Float64frombits(uint64(0x8000000000000000))
	goto L12
L12:
	;
	v64 = base.F64_add(v61, v49)
	goto L9
L13:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	if int32(2) <= v367 {
		goto L99
	} else {
		goto L100
	}
L14:
	;
	v103 = v99
	goto L16
L15:
	;
	v103 = v100
	goto L16
L16:
	;
	v104 = base.I32_trunc_sat_f64_u(v103)
	if base.F64_le(v66, float64(0)) != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v108 = float64(1000)
	goto L19
L18:
	;
	v108 = v66
	goto L19
L19:
	;
	v109 = base.F64_mul(v108, base.F64_convert_i32_s(v88+int32(24)))
	v111 = v88 + int32(68)
	if l4 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v116 = base.F64_mul(base.F64_convert_i32_s(v71+int32(1)), base.F64_convert_i32_u(v104))
	v117 = float64(4.294967295e+09)
	if base.F64_lt(v116, v117) != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v123 = v104
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v123
	goto L26
L23:
	;
	v120 = v116
	goto L25
L24:
	;
	v120 = v117
	goto L25
L25:
	;
	v123 = base.I32_trunc_sat_f64_u(v120)
	goto L22
L26:
	;
	v125 = base.I32_div_u_s(v123, v111)
	v126 = int32(50)
	v127 = base.I32_div_u_s(v125, v126)
	if base.Ui32(v126) <= base.Ui32(v125) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v77))) = v127
	v137 = int32(1)
	v142 = base.F64_ceil(v108)
	v144 = int32(268435455)
	v146 = int32(base.Ui32(v133) >> (uint(int32(2)) % 32))
	if base.Ui32(v144) <= base.Ui32(v146) {
		goto L33
	} else {
		goto L34
	}
L29:
	;
	v132 = v127 * v111
	goto L31
L30:
	;
	v132 = int32(0)
	goto L31
L31:
	;
	v133 = v123 - v132
	goto L28
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17+int32(12)))) = v359
	*(*int32)(unsafe.Add(mBase, uint32(v17+int32(8)))) = v353
	goto L13
L33:
	;
	v149 = v144
	goto L35
L34:
	;
	v149 = v146
	goto L35
L35:
	;
	v151 = int32(base.Ui32(int32(-2147483648)) >> (uint(base.I32_clz(v149)) % 32))
	v152 = base.F64_convert_i32_u(v151)
	if base.F64_gt(v152, v142) != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v154 = v142
	goto L38
L37:
	;
	v154 = v152
	goto L38
L38:
	;
	v155 = base.I32_trunc_sat_f64_s(v154)
	if v155 <= int32(1024) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v158 = int32(1024)
	goto L41
L40:
	;
	v158 = v155
	goto L41
L41:
	;
	if v158&(v158-int32(1)) != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v165 = v137 << (uint(int32(32)-base.I32_clz(v158)) % 32)
	goto L44
L43:
	;
	v165 = v158
	goto L44
L44:
	;
	if base.F64_lt(base.F64_convert_i32_u(v133), base.F64_add(v109, base.F64_convert_i32_u(v165<<(uint(int32(2))%32)))) == int32(0) {
		v353 = v137
		v359 = v165
		goto L32
	} else {
		goto L45
	}
L45:
	;
	if l4 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v175 = *(*float64)(unsafe.Add(mBase, _c_F_initial_cost_hashjoin[4]))
	v177 = *(*int32)(unsafe.Add(mBase, _c_F_initial_cost_hashjoin[5]))
	v181 = base.F64_mul(base.F64_mul(v175, base.F64_convert_i32_s(v177)), float64(1024))
	v182 = float64(4.294967295e+09)
	if base.F64_lt(v181, v182) != 0 {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	v234 = v152
	v235 = v133
	v239 = v151
	goto L48
L48:
	;
	v242 = v88 + int32(28)
	if base.Ui32(v242) < base.Ui32(v235) {
		goto L71
	} else {
		goto L72
	}
L49:
	;
	v185 = v181
	goto L51
L50:
	;
	v185 = v182
	goto L51
L51:
	;
	v186 = base.I32_trunc_sat_f64_u(v185)
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v186
	goto L52
L52:
	;
	v188 = base.I32_div_u_s(v186, v111)
	v189 = int32(50)
	v190 = base.I32_div_u_s(v188, v189)
	if base.Ui32(v189) <= base.Ui32(v188) {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v77))) = v190
	v205 = int32(268435455)
	v207 = int32(base.Ui32(v196) >> (uint(int32(2)) % 32))
	if base.Ui32(v205) <= base.Ui32(v207) {
		goto L58
	} else {
		goto L59
	}
L55:
	;
	v195 = v190 * v111
	goto L57
L56:
	;
	v195 = int32(0)
	goto L57
L57:
	;
	v196 = v186 - v195
	goto L54
L58:
	;
	v210 = v205
	goto L60
L59:
	;
	v210 = v207
	goto L60
L60:
	;
	v212 = int32(base.Ui32(int32(-2147483648)) >> (uint(base.I32_clz(v210)) % 32))
	v213 = base.F64_convert_i32_u(v212)
	if base.F64_gt(v213, v142) != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v215 = v142
	goto L63
L62:
	;
	v215 = v213
	goto L63
L63:
	;
	v216 = base.I32_trunc_sat_f64_s(v215)
	if v216 <= int32(1024) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v219 = int32(1024)
	goto L66
L65:
	;
	v219 = v216
	goto L66
L66:
	;
	if v219&(v219-int32(1)) != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v226 = int32(1) << (uint(int32(32)-base.I32_clz(v219)) % 32)
	goto L69
L68:
	;
	v226 = v219
	goto L69
L69:
	;
	if base.F64_lt(base.F64_convert_i32_u(v196), base.F64_add(v109, base.F64_convert_i32_u(v226<<(uint(int32(2))%32)))) == int32(0) {
		v353 = v137
		v359 = v226
		goto L32
	} else {
		goto L70
	}
L70:
	;
	v234 = v213
	v235 = v196
	v239 = v212
	goto L48
L71:
	;
	v244 = int32(1)
	v246 = base.I32_div_u_s(v235, v242)
	if v246&(v246-v244) != 0 {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	v257 = int32(1)
	goto L73
L73:
	;
	v258 = int32(1)
	v259 = int32(32)
	if v257&(v257-v258) != 0 {
		goto L81
	} else {
		goto L82
	}
L74:
	;
	v253 = v244 << (uint(int32(32)-base.I32_clz(v246)) % 32)
	goto L76
L75:
	;
	v253 = v246
	goto L76
L76:
	;
	if base.Ui32(v253) < base.Ui32(v239) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v255 = v253
	goto L79
L78:
	;
	v255 = v239
	goto L79
L79:
	;
	v257 = v255
	goto L73
L80:
	;
	v353 = v334
	v359 = v342
	goto L32
L81:
	;
	v269 = v258 << (uint(v259-base.I32_clz(v257)) % 32)
	goto L83
L82:
	;
	v269 = v257
	goto L83
L83:
	;
	v275 = base.F64_ceil(base.F64_div(v109, base.F64_convert_i32_u(v235-v269<<(uint(int32(2))%32))))
	if base.F64_gt(v234, v275) != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v277 = v275
	goto L86
L85:
	;
	v277 = v234
	goto L86
L86:
	;
	v278 = base.I32_trunc_sat_f64_s(v277)
	if v278 <= int32(2) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v281 = int32(2)
	goto L89
L88:
	;
	v281 = v278
	goto L89
L89:
	;
	if v281&(v281-int32(1)) != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v288 = v258 << (uint(v259-base.I32_clz(v281)) % 32)
	goto L92
L91:
	;
	v288 = v281
	goto L92
L92:
	;
	if base.B2i32(v288 < int32(2))|base.B2i32(base.Ui32(int32(134217727)) < base.Ui32(v269)) != 0 {
		v334 = v288
		v342 = v269
		goto L80
	} else {
		goto L93
	}
L93:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v296 = v288
	v298 = v294
	v304 = v269
	goto L94
L94:
	;
	if base.B2i32(v298 < int32(0))|base.B2i32(base.Ui32(v296) < base.Ui32(int32(base.Ui32(v298)>>(uint(int32(13))%32)))) != 0 {
		v334 = v296
		v342 = v304
		goto L80
	} else {
		goto L96
	}
L95:
	;
	v353 = v326
	v359 = v328
	goto L32
L96:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	v318 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v77))) = v317 << (uint(v318) % 32)
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v323 = v321 << (uint(v318) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v323
	v326 = int32(base.Ui32(v296) >> (uint(v318) % 32))
	v328 = v304 << (uint(v318) % 32)
	if base.Ui32(v296) < base.Ui32(int32(4)) {
		v353 = v326
		v359 = v328
		goto L32
	} else {
		goto L97
	}
L97:
	;
	if base.Ui32(v304) < base.Ui32(int32(67108864)) {
		v296 = v326
		v298 = v323
		v304 = v328
		goto L94
	} else {
		goto L98
	}
L98:
	;
	goto L95
L99:
	;
	v371 = *(*float64)(unsafe.Add(mBase, _c_F_initial_cost_hashjoin[6]))
	v372 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v372)+32))
	v374 = int32(7)
	v376 = int32(-8)
	v378 = int32(24)
	v382 = float64(0.0001220703125)
	v384 = base.F64_ceil(base.F64_mul(base.F64_mul(v20, base.F64_convert_i32_u((v373+v374)&v376+v378)), v382))
	v386 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v386)+32))
	v398 = base.F64_ceil(base.F64_mul(base.F64_mul(v19, base.F64_convert_i32_u((v387+v374)&v376+v378)), v382))
	v404 = base.F64_add(base.F64_mul(v371, v398), v43)
	v406 = base.F64_add(base.F64_mul(v371, base.F64_add(base.F64_add(v384, v384), v398)), v38)
	goto L101
L100:
	;
	v404 = v43
	v406 = v38
	goto L101
L101:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+24)) = v406
	*(*float64)(unsafe.Add(mBase, uint32(l0)+8)) = v404
	*(*float64)(unsafe.Add(mBase, uint32(l0)+16)) = base.F64_add(v406, v404)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v47 + (v45 ^ int32(1)) + v46
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+88)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v367
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v418
	m.G0 = v17 + int32(16)
	return
}
func F_initscan(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
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
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int64
	_ = v127
	v4 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v7 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+48))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+118)))
	if v17 == int32(116) {
		v40 = v4
		goto L8
	} else {
		goto L9
	}
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
	v13 = v8
	goto L1
L3:
	;
	goto L4
L4:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = F_RelationGetNumberOfBlocksInFork(m, v9, int32(0))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return
L6:
	;
	v13 = v11
	goto L1
L7:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v49 != 0 {
		goto L19
	} else {
		goto L20
	}
L8:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	if v41 != 0 {
		goto L14
	} else {
		goto L15
	}
L9:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_initscan[1]))
	v23 = base.I32_div_s(v21, int32(4))
	if base.Ui32(v13) <= base.Ui32(v23) {
		v40 = v4
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v29 = int32(base.Ui32(v25&int32(128)) >> (uint(int32(7)) % 32))
	if v25&int32(64) == int32(0) {
		v40 = v29
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	if v34 != 0 {
		v47 = v29
		goto L7
	} else {
		goto L12
	}
L12:
	;
	v36 = F_GetAccessStrategy(m, int32(1))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v36
	v47 = v29
	goto L7
L14:
	;
	F_bms_free(m, v41)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
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
	v47 = v40
	goto L7
L17:
	;
	goto L16
L18:
	;
	v82 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v82
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)) = uint8(v82)
	v86 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v86
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+72)) = uint16(v82)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v86
	*(*int64)(unsafe.Add(mBase, uint32(l0)+52)) = int64(4294967295)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+100)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+88)) = int64(-4294967295)
	if l1 == v82 {
		goto L35
	} else {
		goto L36
	}
L19:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+12)))
	if v50 == int32(1) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_initscan[0])))
	v61 = v47 & v60
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
	if v61 != 0 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	if v61 != 0 {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v48 | int32(128)
	goto L18
L29:
	;
	goto L30
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v48 & int32(-129)
	goto L18
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v48 | int32(128)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v73 = F_ss_get_location(m, v71, v72)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L5
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v48 & int32(-129)
	goto L18
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v73
	goto L18
L35:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v110&int32(1) == int32(0) {
		goto L39
	} else {
		goto L40
	}
L36:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v100 <= int32(0) {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v104 = v100 * int32(48)
	if v104 == int32(0) {
		goto L35
	} else {
		goto L38
	}
L38:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	base.MemoryCopy(m, v107, l1, v104)
	goto L35
L39:
	;
	return
L40:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+272))
	if v116 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+268)))
	if v119 != int32(1) {
		goto L39
	} else {
		goto L44
	}
L42:
	;
	v126 = v116
	goto L43
L43:
	;
	v127 = *(*int64)(unsafe.Add(mBase, uint32(v126)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v126)+16)) = v127 + int64(1)
	goto L39
L44:
	;
	F_pgstat_assoc_relation(m, v115)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L5
	} else {
		goto L45
	}
L45:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)+272))
	v126 = v125
	goto L43
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
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int64
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int64
	_ = v64
	var v66 int64
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int64
	_ = v80
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
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
		v24 = v22
	} else {
		v24 = int32(6)
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v24
	v27 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_inittapes[0])))
	if v27 != int32(1) {
		v55 = v24
		v58 = base.I64_extend_i32_s(v55) << (uint(int64(13)) % 64)
		v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
		v60 = F_GetMemoryChunkSpace(m, v59)
		mBase = m.M
		v61 = m.ExcPending
		if v61 != 0 {
			return
		} else {
			v64 = *(*int64)(unsafe.Add(mBase, uint32(l0)+96))
			if v58+base.I64_extend_i32_u(v60) < v64 {
				v66 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+88)) = v66 - v58
			} else {
			}
			F_PrepareTempTablespaces(m)
			mBase = m.M
			v70 = m.ExcPending
			if v70 != 0 {
				return
			} else {
				v71 = int32(0)
				v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+236))
				if v72 != 0 {
					v76 = v72 + int32(12)
				} else {
					v76 = v71
				}
				v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
				v78 = F_LogicalTapeSetCreate(m, v71, v76, v77)
				mBase = m.M
				v79 = m.ExcPending
				if v79 != 0 {
					return
				} else {
					v80 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(l0)+168)) = v80
					*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v78
					*(*int64)(unsafe.Add(mBase, uint32(l0)+176)) = v80
					v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
					v88 = F_palloc0(m, v85<<(uint(int32(2))%32))
					mBase = m.M
					v89 = m.ExcPending
					if v89 != 0 {
						return
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(l0)+188)) = int64(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = v88
						*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = int32(2)
						v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
						v96 = F_LogicalTapeCreate(m, v95)
						mBase = m.M
						v97 = m.ExcPending
						if v97 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+196)) = v96
							v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
							v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
							*(*int32)(unsafe.Add(mBase, uint32(v99+v100<<(uint(int32(2))%32)))) = v96
							v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
							v106 = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = v105 + v106
							v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = v109 + v106
							m.G0 = v8 + int32(16)
							return
						}
					}
				}
			}
		}
	} else {
		v32 = F_errstart(m, int32(15), int32(0))
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return
		} else {
			v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
			if v32 == int32(0) {
				v55 = v34
				v58 = base.I64_extend_i32_s(v55) << (uint(int64(13)) % 64)
				v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
				v60 = F_GetMemoryChunkSpace(m, v59)
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
					return
				} else {
					v64 = *(*int64)(unsafe.Add(mBase, uint32(l0)+96))
					if v58+base.I64_extend_i32_u(v60) < v64 {
						v66 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
						*(*int64)(unsafe.Add(mBase, uint32(l0)+88)) = v66 - v58
					} else {
					}
					F_PrepareTempTablespaces(m)
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return
					} else {
						v71 = int32(0)
						v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+236))
						if v72 != 0 {
							v76 = v72 + int32(12)
						} else {
							v76 = v71
						}
						v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
						v78 = F_LogicalTapeSetCreate(m, v71, v76, v77)
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
							return
						} else {
							v80 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(l0)+168)) = v80
							*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v78
							*(*int64)(unsafe.Add(mBase, uint32(l0)+176)) = v80
							v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
							v88 = F_palloc0(m, v85<<(uint(int32(2))%32))
							mBase = m.M
							v89 = m.ExcPending
							if v89 != 0 {
								return
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(l0)+188)) = int64(0)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = v88
								*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = int32(2)
								v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
								v96 = F_LogicalTapeCreate(m, v95)
								mBase = m.M
								v97 = m.ExcPending
								if v97 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+196)) = v96
									v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
									v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
									*(*int32)(unsafe.Add(mBase, uint32(v99+v100<<(uint(int32(2))%32)))) = v96
									v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
									v106 = int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = v105 + v106
									v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = v109 + v106
									m.G0 = v8 + int32(16)
									return
								}
							}
						}
					}
				}
			} else {
				v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
				v40 = F_pg_rusage_show(m, l0+int32(256))
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v40
					*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v34
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v37
					F_errmsg_internal(m, int32(_a_F_inittapes_0), v8)
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_inittapes_1), int32(1883), int32(_a_F_inittapes_2))
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return
						} else {
							v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
							v55 = v53
							v58 = base.I64_extend_i32_s(v55) << (uint(int64(13)) % 64)
							v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
							v60 = F_GetMemoryChunkSpace(m, v59)
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return
							} else {
								v64 = *(*int64)(unsafe.Add(mBase, uint32(l0)+96))
								if v58+base.I64_extend_i32_u(v60) < v64 {
									v66 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
									*(*int64)(unsafe.Add(mBase, uint32(l0)+88)) = v66 - v58
								} else {
								}
								F_PrepareTempTablespaces(m)
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return
								} else {
									v71 = int32(0)
									v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+236))
									if v72 != 0 {
										v76 = v72 + int32(12)
									} else {
										v76 = v71
									}
									v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
									v78 = F_LogicalTapeSetCreate(m, v71, v76, v77)
									mBase = m.M
									v79 = m.ExcPending
									if v79 != 0 {
										return
									} else {
										v80 = int64(0)
										*(*int64)(unsafe.Add(mBase, uint32(l0)+168)) = v80
										*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v78
										*(*int64)(unsafe.Add(mBase, uint32(l0)+176)) = v80
										v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
										v88 = F_palloc0(m, v85<<(uint(int32(2))%32))
										mBase = m.M
										v89 = m.ExcPending
										if v89 != 0 {
											return
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(l0)+188)) = int64(0)
											*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = v88
											*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = int32(2)
											v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
											v96 = F_LogicalTapeCreate(m, v95)
											mBase = m.M
											v97 = m.ExcPending
											if v97 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+196)) = v96
												v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
												v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
												*(*int32)(unsafe.Add(mBase, uint32(v99+v100<<(uint(int32(2))%32)))) = v96
												v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
												v106 = int32(1)
												*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = v105 + v106
												v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = v109 + v106
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
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
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
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v180 int32
	_ = v180
	var v191 int32
	_ = v191
	var v199 int32
	_ = v199
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
		goto L60
	} else {
		goto L61
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v180
	goto L1
L3:
	;
	v180 = v28
	goto L2
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L11
	} else {
		goto L54
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L11
	} else {
		goto L49
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L11
	} else {
		goto L44
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
		v180 = v28
		goto L2
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	if base.B2i32(v17 == int32(0))|base.B2i32(v28 != int32(1)) != 0 {
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
	F_errmsg(m, int32(_a_F_insertSelectOptions_0), int32(0))
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
	F_errfinish(m, int32(_a_F_insertSelectOptions_1), int32(_a_F_insertSelectOptions_2), int32(_a_F_insertSelectOptions_3))
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
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v59 <= int32(0) {
		v180 = int32(1)
		goto L2
	} else {
		goto L32
	}
L32:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v65 = int32(0)
	goto L33
L33:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v62+v65<<(uint(int32(2))%32))))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+12))
	if v77 != int32(1) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L11
	} else {
		goto L39
	}
L35:
	;
	v81 = v65 + int32(1)
	if v59 != v81 {
		v65 = v81
		goto L33
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	goto L34
L38:
	;
	goto L3
L39:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L11
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = int32(_a_F_insertSelectOptions_4)
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(_a_F_insertSelectOptions_5)
	F_errmsg(m, int32(_a_F_insertSelectOptions_6), v12)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L11
	} else {
		goto L41
	}
L41:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	F_scanner_errposition(m, v97, l5)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L11
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(_a_F_insertSelectOptions_1), int32(_a_F_insertSelectOptions_7), int32(_a_F_insertSelectOptions_3))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L11
	} else {
		goto L43
	}
L43:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L44:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L11
	} else {
		goto L45
	}
L45:
	;
	F_errmsg(m, int32(_a_F_insertSelectOptions_8), int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L11
	} else {
		goto L46
	}
L46:
	;
	v116 = F_exprLocation(m, l1)
	mBase = m.M
	F_scanner_errposition(m, v116, l5)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L11
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_insertSelectOptions_1), int32(_a_F_insertSelectOptions_9), int32(_a_F_insertSelectOptions_3))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L11
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L49:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L11
	} else {
		goto L50
	}
L50:
	;
	F_errmsg(m, int32(_a_F_insertSelectOptions_10), int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L11
	} else {
		goto L51
	}
L51:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	F_scanner_errposition(m, v135, l5)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L11
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(_a_F_insertSelectOptions_1), int32(_a_F_insertSelectOptions_11), int32(_a_F_insertSelectOptions_3))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L11
	} else {
		goto L53
	}
L53:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L54:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L11
	} else {
		goto L55
	}
L55:
	;
	F_errmsg(m, int32(_a_F_insertSelectOptions_12), int32(0))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L11
	} else {
		goto L56
	}
L56:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	F_scanner_errposition(m, v154, l5)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L11
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(_a_F_insertSelectOptions_1), int32(_a_F_insertSelectOptions_13), int32(_a_F_insertSelectOptions_3))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L11
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
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L11
	} else {
		goto L64
	}
L60:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v191 != 0 {
		goto L59
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	m.G0 = v12 + int32(16)
	return
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = l4
	goto L62
L64:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L11
	} else {
		goto L65
	}
L65:
	;
	F_errmsg(m, int32(_a_F_insertSelectOptions_14), int32(0))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L11
	} else {
		goto L66
	}
L66:
	;
	v207 = F_exprLocation(m, l4)
	mBase = m.M
	F_scanner_errposition(m, v207, l5)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L11
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(_a_F_insertSelectOptions_1), int32(_a_F_insertSelectOptions_15), int32(_a_F_insertSelectOptions_3))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L11
	} else {
		goto L68
	}
L68:
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
	v3 = int32(_a_F_int2eqfast_0)
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
				F_errmsg(m, int32(_a_F_int2pl_0), int32(0))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_int2pl_1), int32(944), int32(_a_F_int2pl_2))
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
				F_errmsg(m, int32(_a_F_int4mi_0), int32(0))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_int4mi_1), int32(843), int32(_a_F_int4mi_2))
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
				F_errmsg(m, int32(_a_F_int4mul_0), int32(0))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_int4mul_1), int32(857), int32(_a_F_int4mul_2))
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
				F_errmsg(m, int32(_a_F_int4um_0), int32(0))
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_int4um_1), int32(807), int32(_a_F_int4um_2))
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
				F_errmsg(m, int32(_a_F_int82_0), int32(0))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_int82_1), int32(1277), int32(_a_F_int82_2))
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
				F_errmsg(m, int32(_a_F_int82mi_0), int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_int82mi_1), int32(1055), int32(_a_F_int82mi_2))
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
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = Fn13863(m, l0, int32(_a_F_int84mul_0), int32(927), int32(_a_F_int84mul_1), int32(_a_F_int84mul_2))
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
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
	v9 = v4 + v8
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
				F_errmsg(m, int32(_a_F_int84pl_0), int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_int84pl_1), int32(899), int32(_a_F_int84pl_2))
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
	var v38 int64
	_ = v38
	var __phi38 int64
	_ = __phi38
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
					F_errmsg(m, int32(_a_F_int8gcd_0), int32(0))
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_int8gcd_1), int32(636), int32(_a_F_int8gcd_2))
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
				__phi38 = v19
				v37 = __phi37
				v38 = __phi38
				for {
					v40 = base.I64_rem_s(v38, v37)
					if v40 != int64(0) {
						__phi37 = v40
						__phi38 = v37
						v37 = __phi37
						v38 = __phi38
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
			__phi38 = v19
			v37 = __phi37
			v38 = __phi38
			for {
				v40 = base.I64_rem_s(v38, v37)
				if v40 != int64(0) {
					__phi37 = v40
					__phi38 = v37
					v37 = __phi37
					v38 = __phi38
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
	var v25 int32
	_ = v25
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
		v25 = v3
	} else {
		v25 = int32(0)
	}
	return v25
}
func F_int8mul(m *base.Module, l0 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = Fn13864(m, l0, int32(_a_F_int8mul_0), int32(499), int32(_a_F_int8mul_1), int32(_a_F_int8mul_2))
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_int8out(m *base.Module, l0 int32) int32 {
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
	var v14 int32
	_ = v14
	var v19 int64
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
	if int64(0) <= v9 {
		v19 = v9
		v20 = int32(0)
	} else {
		v14 = int32(45)
		*(*uint8)(unsafe.Add(mBase, uint32(v6))) = uint8(v14)
		v19 = int64(0) - v9
		v20 = int32(1)
	}
	v22 = F_pg_ulltoa_n(m, v19, v6+v20)
	mBase = m.M
	v23 = v22 + v20
	v25 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v6+v23))) = uint8(v25)
	v28 = v23 + int32(1)
	v29 = F_palloc(m, v28)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		return int32(0)
	} else {
		if v28 != 0 {
			base.MemoryCopy(m, v29, v6, v28)
		} else {
		}
		m.G0 = v6 + int32(32)
		return v29
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
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
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
		goto L32
	case 1:
		goto L31
	default:
		goto L30
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
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L23
	}
L7:
	;
	switch v21 - int32(1024) {
	case 0, 8:
		goto L4
	case 1, 2, 3, 4, 5, 6, 7:
		goto L6
	default:
		goto L21
	}
L8:
	;
	if int32(1)<<(uint(v21)%32)&int32(340) != 0 {
		goto L4
	} else {
		goto L20
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
	if v21 <= int32(_a_F_intervaltypmodin_7) {
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
	switch v21 - int32(_a_F_intervaltypmodin_9) {
	case 0, 8:
		goto L4
	case 1, 2, 3, 4, 5, 6, 7:
		goto L6
	default:
		goto L18
	}
L16:
	;
	if v21 != int32(_a_F_intervaltypmodin_8) {
		goto L6
	} else {
		goto L17
	}
L17:
	;
	goto L4
L18:
	;
	if base.B2i32(v21 == int32(_a_F_intervaltypmodin_10))|base.B2i32(v21 == int32(_a_F_intervaltypmodin_0)) != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	goto L6
L20:
	;
	goto L7
L21:
	;
	if v21 == int32(2048) {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	goto L6
L23:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	F_errmsg(m, int32(_a_F_intervaltypmodin_6), int32(0))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(_a_F_intervaltypmodin_3), int32(1085), int32(_a_F_intervaltypmodin_4))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L47
	}
L28:
	;
	m.G0 = v7 + int32(32)
	return v132
L29:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v132 = v126<<(uint(int32(16))%32)&int32(2147418112) | v76
	goto L28
L30:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L43
	}
L31:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v76 < int32(0) {
		goto L27
	} else {
		goto L34
	}
L32:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v67 == int32(_a_F_intervaltypmodin_0) {
		v132 = int32(-1)
		goto L28
	} else {
		goto L33
	}
L33:
	;
	v132 = v67<<(uint(int32(16))%32)&int32(2147418112) | int32(_a_F_intervaltypmodin_1)
	goto L28
L34:
	;
	if base.Ui32(v76) < base.Ui32(int32(7)) {
		goto L29
	} else {
		goto L35
	}
L35:
	;
	v83 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	if v83 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v132 = v103<<(uint(int32(16))%32)&int32(2147418112) | int32(6)
	goto L28
L40:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = int32(6)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v88
	F_errmsg(m, int32(_a_F_intervaltypmodin_5), v7+int32(16))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_intervaltypmodin_3), int32(1108), int32(_a_F_intervaltypmodin_4))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	goto L39
L43:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	F_errmsg(m, int32(_a_F_intervaltypmodin_6), int32(0))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_intervaltypmodin_3), int32(1118), int32(_a_F_intervaltypmodin_4))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
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
	F_errcode(m, int32(50856066))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v145
	F_errmsg(m, int32(_a_F_intervaltypmodin_2), v7)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(_a_F_intervaltypmodin_3), int32(1102), int32(_a_F_intervaltypmodin_4))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
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
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
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
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	v4 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v19 = v18
	goto L3
L2:
	;
	v19 = v4
	goto L3
L3:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if int32(0) < v21 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L19
	} else {
		goto L52
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L19
	} else {
		goto L48
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L19
	} else {
		goto L42
	}
L7:
	;
	v25 = v21
	v27 = v19
	v30 = v4
	v32 = v4
	goto L10
L8:
	;
	v84 = v19
	v89 = v4
	goto L9
L9:
	;
	if v84 != 0 {
		goto L5
	} else {
		goto L28
	}
L10:
	;
	v40 = l2 + v25<<(uint(int32(4))%32) + v30*int32(100)
	v42 = v40 + int32(20)
	if v27 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v84 = v62
	v89 = v75
	goto L9
L12:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v42)+68))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v42)+76))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v42)+96))
	v66 = F_makeColumnDef(m, v60, v63, v64, v65)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L19
	} else {
		goto L20
	}
L13:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v46 = v27 + int32(4)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	if base.Ui32(v46) < base.Ui32(v49+v50<<(uint(int32(2))%32)) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	v60 = v40 + int32(24)
	v62 = int32(0)
	goto L12
L16:
	;
	v55 = v46
	goto L18
L17:
	;
	v55 = int32(0)
	goto L18
L18:
	;
	v60 = v44
	v62 = v55
	goto L12
L19:
	;
	return
L20:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v66)+52))
	if v68 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v66)+8))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+8))
	v73 = F_type_is_collatable(m, v72)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L19
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v75 = F_lappend(m, v32, v66)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L19
	} else {
		goto L26
	}
L24:
	;
	if v73 != 0 {
		goto L6
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v78 = v30 + int32(1)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v78 < v79 {
		v25 = v79
		v27 = v62
		v30 = v78
		v32 = v75
		goto L10
	} else {
		goto L27
	}
L27:
	;
	goto L11
L28:
	;
	F_create_ctas_internal(m, v14+int32(20), v89, v16)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L19
	} else {
		goto L29
	}
L29:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
	v100 = F_table_open(m, v98, int32(8))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L19
	} else {
		goto L30
	}
L30:
	;
	v102 = int32(0)
	v104 = F_check_enable_rls(m, v98, v102, v102)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L19
	} else {
		goto L31
	}
L31:
	;
	if v104 == int32(2) {
		goto L4
	} else {
		goto L32
	}
L32:
	;
	if v20 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v97
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v100
	v119 = F_GetCurrentCommandId(m, int32(1))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L19
	} else {
		goto L37
	}
L34:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+32)))
	if v110 != 0 {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	F_SetMatViewPopulatedState(m, v100, int32(1))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v119
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+32)))
	if v124 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v128 = int32(0)
	goto L40
L39:
	;
	v126 = F_GetBulkInsertState(m)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L19
	} else {
		goto L41
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v128
	m.G0 = v14 + int32(32)
	return
L41:
	;
	v128 = v126
	goto L40
L42:
	;
	F_errcode(m, int32(34209924))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L19
	} else {
		goto L43
	}
L43:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v66)+8))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)+8))
	v143 = F_format_type_be(m, v142)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L19
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v140
	F_errmsg(m, int32(_a_F_intorel_startup_0), v14)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L19
	} else {
		goto L45
	}
L45:
	;
	F_errhint(m, int32(_a_F_intorel_startup_1), int32(0))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L19
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F_intorel_startup_2), int32(515), int32(_a_F_intorel_startup_3))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
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
	v165 = m.ExcPending
	if v165 != 0 {
		goto L19
	} else {
		goto L49
	}
L49:
	;
	F_errmsg(m, int32(_a_F_intorel_startup_4), int32(0))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L19
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(_a_F_intorel_startup_2), int32(523), int32(_a_F_intorel_startup_3))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
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
	v181 = m.ExcPending
	if v181 != 0 {
		goto L19
	} else {
		goto L53
	}
L53:
	;
	F_errmsg(m, int32(_a_F_intorel_startup_5), int32(0))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L19
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_intorel_startup_2), int32(546), int32(_a_F_intorel_startup_3))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
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
					v24 = F_ArrayGetNItemsSafe(m, v21, v15+int32(16))
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
				v24 = F_ArrayGetNItemsSafe(m, v21, v15+int32(16))
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
	var v25 int32
	_ = v25
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
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
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
	return v76
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
	v69 = F_expression_tree_walker_impl(m, l0, int32(497), l1)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L16
	} else {
		goto L22
	}
L8:
	;
	v65 = F_query_tree_walker_impl(m, l0, int32(497), l1, int32(4))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
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
	v25 = v3
	goto L11
L11:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28+v25<<(uint(int32(2))%32))))
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
	v53 = v25 + int32(1)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v53 < v54 {
		v25 = v53
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
	F_relation_close(m, v39, int32(1))
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
		v76 = v36
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
	return v65
L22:
	;
	v76 = v69
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
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	v5 = int32(0)
	goto L1
L1:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v5*int32(40))+uint32(_c_F_is_objectclass_supported[0])))
	v10 = base.B2i32(v9 == l0)
	if v10 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	return v10
L3:
	;
	v14 = v5 + int32(1)
	if v14 != int32(37) {
		v5 = v14
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
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13928(m, l0, int32(4))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
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
	var v92 int32
	_ = v92
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
	var v141 int32
	_ = v141
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
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
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
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v384 int32
	_ = v384
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
	var v417 int32
	_ = v417
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
	var v455 int32
	_ = v455
	var v466 int32
	_ = v466
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
	var v512 int32
	_ = v512
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
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
	var v605 int32
	_ = v605
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
	var v662 int32
	_ = v662
	var v680 int32
	_ = v680
	var v691 int32
	_ = v691
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
	v141 = v38
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
	v92 = v79
	goto L14
L14:
	;
	v106 = v92 - int32(4)
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
	*(*int32)(unsafe.Add(mBase, uint32(v92))) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v106))) = v83
	if base.Ui32(v19) < base.Ui32(v106) {
		v92 = v106
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
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v141-int32(4))))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
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
	v162 = v141 + int32(4)
	if base.Ui32(v162) < base.Ui32(v59) {
		v141 = v162
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
	v311 = v38
	v312 = v306
	v313 = v38
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
	v210 = v168 + v178
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
	if v187 < v186 {
		v206 = v179
		goto L42
	} else {
		goto L58
	}
L45:
	;
	if v188 < v187 {
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
	if v186 < v187 {
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
	v195 = v19
	goto L52
L51:
	;
	v195 = v182
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
	v202 = v182
	goto L61
L60:
	;
	v202 = v19
	goto L61
L61:
	;
	v208 = v202
	goto L41
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
	v206 = v205
	goto L42
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
	if v215 < v214 {
		v234 = v168
		goto L67
	} else {
		goto L83
	}
L70:
	;
	if v216 < v215 {
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
	if v214 < v215 {
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
	v223 = v209
	goto L77
L76:
	;
	v223 = v210
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
	v230 = v210
	goto L86
L85:
	;
	v230 = v209
	goto L86
L86:
	;
	v236 = v230
	goto L66
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
	v234 = v233
	goto L67
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
	if v243 < v242 {
		v262 = v238
		goto L92
	} else {
		goto L108
	}
L95:
	;
	if v244 < v243 {
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
	if v242 < v243 {
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
	v251 = v237
	goto L102
L101:
	;
	v251 = v172
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
	v258 = v172
	goto L111
L110:
	;
	v258 = v237
	goto L111
L111:
	;
	v264 = v258
	goto L91
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
	v262 = v261
	goto L92
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
	if v274 < v273 {
		v293 = v267
		goto L117
	} else {
		goto L133
	}
L120:
	;
	if v275 < v274 {
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
	if v273 < v274 {
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
	v282 = v266
	goto L127
L126:
	;
	v282 = v268
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
	v289 = v268
	goto L136
L135:
	;
	v289 = v266
	goto L136
L136:
	;
	v295 = v289
	goto L116
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
	v293 = v292
	goto L117
L141:
	;
	if base.Ui32(v312) < base.Ui32(v311) {
		v364 = v311
		v366 = v313
		goto L143
	} else {
		goto L144
	}
L143:
	;
	if base.Ui32(v364) <= base.Ui32(v312) {
		goto L158
	} else {
		goto L159
	}
L144:
	;
	v330 = v311
	v332 = v313
	goto L145
L145:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v330)))
	if v136 != 0 {
		goto L149
	} else {
		goto L150
	}
L146:
	;
	v364 = v358
	v366 = v356
	goto L143
L147:
	;
	v358 = v330 + int32(4)
	if base.Ui32(v358) <= base.Ui32(v312) {
		v330 = v358
		v332 = v356
		goto L145
	} else {
		goto L156
	}
L148:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v332)))
	*(*int32)(unsafe.Add(mBase, uint32(v332))) = v345
	*(*int32)(unsafe.Add(mBase, uint32(v330))) = v350
	v356 = v332 + int32(4)
	goto L147
L149:
	;
	if v345 < v344 {
		v356 = v332
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
		v356 = v332
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
	v364 = v330
	v366 = v332
	goto L143
L154:
	;
	if v345 < v344 {
		v364 = v330
		v366 = v332
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
	v741 = *(*int32)(unsafe.Add(mBase, uint32(v364)))
	*(*int32)(unsafe.Add(mBase, uint32(v364))) = v398
	*(*int32)(unsafe.Add(mBase, uint32(v384))) = v741
	v744 = int32(4)
	v311 = v364 + v744
	v312 = v384 - v744
	v313 = v366
	v317 = v389
	goto L141
L158:
	;
	v384 = v312
	v389 = v317
	goto L161
L159:
	;
	v417 = v312
	v422 = v317
	goto L160
L160:
	;
	v431 = int32(2)
	v432 = (v366 - v19) >> (uint(v431) % 32)
	v435 = (v364 - v366) >> (uint(v431) % 32)
	if v432 < v435 {
		goto L174
	} else {
		goto L175
	}
L161:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v384)))
	if v136 != 0 {
		goto L165
	} else {
		goto L166
	}
L162:
	;
	v417 = v410
	v422 = v408
	goto L160
L163:
	;
	v410 = v384 - int32(4)
	if base.Ui32(v364) <= base.Ui32(v410) {
		v384 = v410
		v389 = v408
		goto L161
	} else {
		goto L172
	}
L164:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v389)))
	*(*int32)(unsafe.Add(mBase, uint32(v384))) = v403
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
	v578 = (v422 - v417) >> (uint(v577) % 32)
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
	v442 = v364 - v437<<(uint(int32(2))%32)
	v444 = v437 & int32(3)
	v445 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v437) {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v455 = v445
	v466 = int32(0)
	goto L181
L179:
	;
	v512 = v445
	goto L180
L180:
	;
	v530 = v512
	v533 = v445
	goto L185
L181:
	;
	v471 = v455 << (uint(int32(2)) % 32)
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
	v482 = v479 + v442
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v482)))
	*(*int32)(unsafe.Add(mBase, uint32(v480))) = v483
	*(*int32)(unsafe.Add(mBase, uint32(v482))) = v481
	v487 = v471 | int32(8)
	v488 = v19 + v487
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v488)))
	v490 = v487 + v442
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
	v503 = v455 + v478
	v505 = v466 + v478
	if v505 != v437&int32(-4) {
		v455 = v503
		v466 = v505
		goto L181
	} else {
		goto L183
	}
L182:
	;
	if v444 == int32(0) {
		goto L173
	} else {
		goto L184
	}
L183:
	;
	goto L182
L184:
	;
	v512 = v503
	goto L180
L185:
	;
	v546 = v530 << (uint(int32(2)) % 32)
	v547 = v19 + v546
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v547)))
	v549 = v442 + v546
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v549)))
	*(*int32)(unsafe.Add(mBase, uint32(v547))) = v550
	*(*int32)(unsafe.Add(mBase, uint32(v549))) = v548
	v553 = int32(1)
	v556 = v533 + v553
	if v556 != v444 {
		v530 = v530 + v553
		v533 = v556
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
	v605 = v593
	v610 = int32(0)
	goto L196
L194:
	;
	v662 = v593
	goto L195
L195:
	;
	v680 = v662
	v691 = v593
	goto L200
L196:
	;
	v619 = v605 << (uint(int32(2)) % 32)
	v620 = v364 + v619
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v620)))
	v622 = v590 + v619
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v622)))
	*(*int32)(unsafe.Add(mBase, uint32(v620))) = v623
	*(*int32)(unsafe.Add(mBase, uint32(v622))) = v621
	v626 = int32(4)
	v627 = v619 | v626
	v628 = v364 + v627
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v628)))
	v630 = v590 + v627
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v630)))
	*(*int32)(unsafe.Add(mBase, uint32(v628))) = v631
	*(*int32)(unsafe.Add(mBase, uint32(v630))) = v629
	v635 = v619 | int32(8)
	v636 = v364 + v635
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v636)))
	v638 = v590 + v635
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v638)))
	*(*int32)(unsafe.Add(mBase, uint32(v636))) = v639
	*(*int32)(unsafe.Add(mBase, uint32(v638))) = v637
	v643 = v619 | int32(12)
	v644 = v364 + v643
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v644)))
	v646 = v590 + v643
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v646)))
	*(*int32)(unsafe.Add(mBase, uint32(v644))) = v647
	*(*int32)(unsafe.Add(mBase, uint32(v646))) = v645
	v651 = v605 + v626
	v653 = v610 + v626
	if v653 != v585&int32(-4) {
		v605 = v651
		v610 = v653
		goto L196
	} else {
		goto L198
	}
L197:
	;
	if v592 == int32(0) {
		goto L188
	} else {
		goto L199
	}
L198:
	;
	goto L197
L199:
	;
	v662 = v651
	goto L195
L200:
	;
	v694 = v680 << (uint(int32(2)) % 32)
	v695 = v364 + v694
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v695)))
	v697 = v590 + v694
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v697)))
	*(*int32)(unsafe.Add(mBase, uint32(v695))) = v698
	*(*int32)(unsafe.Add(mBase, uint32(v697))) = v696
	v701 = int32(1)
	v704 = v691 + v701
	if v704 != v592 {
		v680 = v680 + v701
		v691 = v704
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v11 == int32(1) {
		if v10 < v9 {
			if v9 < v8 {
				v31 = l1
				return v31
			} else {
				if v10 < v8 {
					v26 = l2
				} else {
					v26 = l0
				}
				return v26
			}
		} else {
			if v8 < v9 {
				v31 = l1
				return v31
			} else {
				if v10 < v8 {
					v17 = l0
				} else {
					v17 = l2
				}
				return v17
			}
		}
	} else {
		if v9 < v10 {
			if v8 < v9 {
				v31 = l1
			} else {
				if v8 < v10 {
					v30 = l2
				} else {
					v30 = l0
				}
				v31 = v30
			}
			return v31
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
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13871(m, l0, int32(5))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_iswspace(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v41 int32
	_ = v41
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
	return base.B2i32(v41 != int32(0))
L5:
	;
	v11 = int32(_a_F_iswspace_0)
	goto L8
L6:
	;
	goto L7
L7:
	;
	v21 = int32(_a_F_iswspace_0)
	v25 = v21
	goto L18
L8:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
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
		v11 = v11 + int32(4)
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
	v20 = v11
	goto L16
L15:
	;
	v20 = int32(0)
	goto L16
L16:
	;
	v41 = v20
	goto L4
L17:
	;
	v41 = (v25-v21)>>(uint(int32(2))%32)<<(uint(int32(2))%32) + int32(_a_F_iswspace_0)
	goto L4
L18:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	if v29 != 0 {
		v25 = v25 + int32(4)
		goto L18
	} else {
		goto L20
	}
L19:
	;
	goto L17
L20:
	;
	goto L19
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
	var v103 int32
	_ = v103
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
	var v152 int32
	_ = v152
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
	var v218 int32
	_ = v218
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
	var v292 int32
	_ = v292
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
	var v361 int32
	_ = v361
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
	var v519 int32
	_ = v519
	var v524 int32
	_ = v524
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v542 int32
	_ = v542
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
	var v786 int32
	_ = v786
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
	var v876 int32
	_ = v876
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
	var v990 int32
	_ = v990
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
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1060 int32
	_ = v1060
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1066 int32
	_ = v1066
	var v1068 int32
	_ = v1068
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1095 int32
	_ = v1095
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1123 int32
	_ = v1123
	var v1125 int32
	_ = v1125
	var v1127 int32
	_ = v1127
	var v1130 int32
	_ = v1130
	var v1133 int32
	_ = v1133
	var v1136 int32
	_ = v1136
	var v1140 int32
	_ = v1140
	var v1143 int32
	_ = v1143
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1161 int32
	_ = v1161
	var v1162 int32
	_ = v1162
	var v1166 int32
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1179 int32
	_ = v1179
	var v1180 int32
	_ = v1180
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1193 int32
	_ = v1193
	var v1196 int32
	_ = v1196
	var v1197 int32
	_ = v1197
	var v1199 int32
	_ = v1199
	var v1201 int32
	_ = v1201
	var v1215 int32
	_ = v1215
	var v1216 int32
	_ = v1216
	var v1219 int32
	_ = v1219
	var v1221 int32
	_ = v1221
	var v1222 int32
	_ = v1222
	var v1224 int32
	_ = v1224
	var v1225 int32
	_ = v1225
	var v1230 int32
	_ = v1230
	var v1232 int32
	_ = v1232
	var v1234 int32
	_ = v1234
	var v1237 int32
	_ = v1237
	var v1240 int32
	_ = v1240
	var v1243 int32
	_ = v1243
	var v1247 int32
	_ = v1247
	var v1250 int32
	_ = v1250
	var v1252 int32
	_ = v1252
	var v1253 int32
	_ = v1253
	var v1255 int32
	_ = v1255
	var v1256 int32
	_ = v1256
	var v1259 int32
	_ = v1259
	var v1260 int32
	_ = v1260
	var v1262 int32
	_ = v1262
	var v1263 int32
	_ = v1263
	var v1266 int32
	_ = v1266
	var v1269 int32
	_ = v1269
	var v1270 int32
	_ = v1270
	var v1272 int32
	_ = v1272
	var v1274 int32
	_ = v1274
	var v1288 int32
	_ = v1288
	var v1289 int32
	_ = v1289
	var v1292 int32
	_ = v1292
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1297 int32
	_ = v1297
	var v1298 int32
	_ = v1298
	var v1301 int32
	_ = v1301
	var v1302 int32
	_ = v1302
	var v1304 int32
	_ = v1304
	var v1305 int32
	_ = v1305
	var v1308 int32
	_ = v1308
	var v1310 int32
	_ = v1310
	var v1312 int32
	_ = v1312
	var v1315 int32
	_ = v1315
	var v1318 int32
	_ = v1318
	var v1321 int32
	_ = v1321
	var v1325 int32
	_ = v1325
	var v1328 int32
	_ = v1328
	var v1330 int32
	_ = v1330
	var v1331 int32
	_ = v1331
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
	var v1337 int32
	_ = v1337
	var v1339 int32
	_ = v1339
	var v1341 int32
	_ = v1341
	var v1344 int32
	_ = v1344
	var v1347 int32
	_ = v1347
	var v1350 int32
	_ = v1350
	var v1354 int32
	_ = v1354
	var v1357 int32
	_ = v1357
	var v1359 int32
	_ = v1359
	var v1360 int32
	_ = v1360
	var v1362 int32
	_ = v1362
	var v1363 int32
	_ = v1363
	var v1368 int32
	_ = v1368
	var v1369 int32
	_ = v1369
	var v1373 int32
	_ = v1373
	var v1374 int32
	_ = v1374
	var v1376 int32
	_ = v1376
	var v1379 int32
	_ = v1379
	var v1383 int32
	_ = v1383
	var v1384 int32
	_ = v1384
	var v1385 int32
	_ = v1385
	var v1387 int32
	_ = v1387
	var v1388 int32
	_ = v1388
	var v1397 int32
	_ = v1397
	var v1409 int32
	_ = v1409
	var v1421 int32
	_ = v1421
	var v1422 int32
	_ = v1422
	var v1426 int32
	_ = v1426
	var v1428 int32
	_ = v1428
	var v1434 int32
	_ = v1434
	var v1448 int32
	_ = v1448
	var v1452 int32
	_ = v1452
	var v1453 int32
	_ = v1453
	var v1455 int32
	_ = v1455
	var v1456 int32
	_ = v1456
	var v1458 int32
	_ = v1458
	var v1459 int32
	_ = v1459
	var v1462 int32
	_ = v1462
	var v1464 int32
	_ = v1464
	var v1466 int32
	_ = v1466
	var v1470 int32
	_ = v1470
	var v1474 int32
	_ = v1474
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1480 int32
	_ = v1480
	var v1481 int32
	_ = v1481
	var v1484 int32
	_ = v1484
	var v1485 int32
	_ = v1485
	var v1488 int32
	_ = v1488
	var v1490 int32
	_ = v1490
	var v1492 int32
	_ = v1492
	var v1494 int32
	_ = v1494
	var v1496 int32
	_ = v1496
	var v1500 int32
	_ = v1500
	var v1504 int32
	_ = v1504
	var v1516 int32
	_ = v1516
	var v1528 int32
	_ = v1528
	var v1529 int32
	_ = v1529
	var v1533 int32
	_ = v1533
	var v1535 int32
	_ = v1535
	var v1541 int32
	_ = v1541
	var v1555 int32
	_ = v1555
	var v1559 int32
	_ = v1559
	var v1560 int32
	_ = v1560
	var v1561 int32
	_ = v1561
	var v1562 int32
	_ = v1562
	var v1564 int32
	_ = v1564
	var v1565 int32
	_ = v1565
	var v1569 int32
	_ = v1569
	var v1572 int32
	_ = v1572
	var v1573 int32
	_ = v1573
	var v1576 int32
	_ = v1576
	var v1578 int32
	_ = v1578
	var v1582 int32
	_ = v1582
	var v1584 int32
	_ = v1584
	var v1586 int32
	_ = v1586
	var v1587 int32
	_ = v1587
	var v1596 int32
	_ = v1596
	var v1597 int32
	_ = v1597
	var v1598 int32
	_ = v1598
	var v1602 int32
	_ = v1602
	var v1605 int32
	_ = v1605
	var v1606 int32
	_ = v1606
	var v1611 int32
	_ = v1611
	var v1612 int32
	_ = v1612
	var v1617 int32
	_ = v1617
	var v1618 int32
	_ = v1618
	var v1627 int32
	_ = v1627
	var v1628 int32
	_ = v1628
	var v1632 int32
	_ = v1632
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
	v18 = F_memcmp(m, v16+v6, int32(_a_F_italian_ISO_8859_1_stem_0), v8)
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
	v32 = F_slice_from_s(m, l0, int32(5), int32(_a_F_italian_ISO_8859_1_stem_1))
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
	return v1632
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v43
	v50 = F_find_among(m, l0, int32(_a_F_italian_ISO_8859_1_stem_2), int32(7))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L8
	} else {
		goto L17
	}
L15:
	;
	v103 = v6
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
	v88 = F_slice_from_s(m, l0, int32(2), int32(_a_F_italian_ISO_8859_1_stem_3))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L8
	} else {
		goto L36
	}
L21:
	;
	v82 = F_slice_from_s(m, l0, int32(1), int32(_a_F_italian_ISO_8859_1_stem_4))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L8
	} else {
		goto L34
	}
L22:
	;
	v76 = F_slice_from_s(m, l0, int32(1), int32(_a_F_italian_ISO_8859_1_stem_5))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L8
	} else {
		goto L32
	}
L23:
	;
	v70 = F_slice_from_s(m, l0, int32(1), int32(_a_F_italian_ISO_8859_1_stem_6))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L8
	} else {
		goto L30
	}
L24:
	;
	v64 = F_slice_from_s(m, l0, int32(1), int32(_a_F_italian_ISO_8859_1_stem_7))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L8
	} else {
		goto L28
	}
L25:
	;
	v58 = F_slice_from_s(m, l0, int32(1), int32(_a_F_italian_ISO_8859_1_stem_8))
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
	v1632 = v58
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
	v1632 = v64
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
	v1632 = v70
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
	v1632 = v76
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
	v1632 = v82
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
	v1632 = v88
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
	v1397 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1397
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1397
	v1409 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L431
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v103
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v114 < v103 {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	v1374 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1374
	v1376 = *(*int32)(unsafe.Add(mBase, uint32(v1373)+8))
	if v1374 < v1376 {
		goto L41
	} else {
		goto L420
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
	v116 = v103
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
	v157 = v152
	goto L45
L50:
	;
	if v103 == v116 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v152 = int32(0)
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
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129+v103))))
	if int32(249) < v131 {
		v152 = v128
		goto L49
	} else {
		goto L55
	}
L55:
	;
	v133 = v131 - int32(97)
	if v133 < int32(0) {
		v152 = v128
		goto L49
	} else {
		goto L56
	}
L56:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v133)>>(uint(int32(3))%32)))+uint32(_c_F_italian_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v139)>>(uint(v133&int32(7))%32))&int32(1) == int32(0) {
		v152 = v128
		goto L49
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v103 + int32(1)
	goto L58
L58:
	;
	goto L51
L59:
	;
	v1368 = F_slice_from_s(m, l0, int32(1), int32(_a_F_italian_ISO_8859_1_stem_9))
	mBase = m.M
	v1369 = m.ExcPending
	if v1369 != 0 {
		goto L8
	} else {
		goto L418
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v103
	if v103 < v302 {
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
	v223 = v218
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
	v218 = int32(0)
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
		v218 = v194
		goto L69
	} else {
		goto L75
	}
L75:
	;
	v199 = v197 - int32(97)
	if v199 < int32(0) {
		v218 = v194
		goto L69
	} else {
		goto L76
	}
L76:
	;
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v199)>>(uint(int32(3))%32)))+uint32(_c_F_italian_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v205)>>(uint(v199&int32(7))%32))&int32(1) == int32(0) {
		v218 = v194
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
	v228 = F_slice_from_s(m, l0, int32(1), int32(_a_F_italian_ISO_8859_1_stem_10))
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
		v1632 = v228
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
	v297 = v292
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
	v292 = int32(0)
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
		v292 = v268
		goto L92
	} else {
		goto L98
	}
L98:
	;
	v273 = v271 - int32(97)
	if v273 < int32(0) {
		v292 = v268
		goto L92
	} else {
		goto L99
	}
L99:
	;
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v273)>>(uint(int32(3))%32)))+uint32(_c_F_italian_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v279)>>(uint(v273&int32(7))%32))&int32(1) == int32(0) {
		v292 = v268
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
	v103 = v103 + int32(1)
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
	v366 = v361
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
	v361 = int32(0)
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
		v361 = v337
		goto L113
	} else {
		goto L119
	}
L119:
	;
	v342 = v340 - int32(97)
	if v342 < int32(0) {
		v361 = v337
		goto L113
	} else {
		goto L120
	}
L120:
	;
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v342)>>(uint(int32(3))%32)))+uint32(_c_F_italian_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v348)>>(uint(v342&int32(7))%32))&int32(1) == int32(0) {
		v361 = v337
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
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v393)>>(uint(int32(3))%32)))+uint32(_c_F_italian_ISO_8859_1_stem[0]))))
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
	v449 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v442)>>(uint(int32(3))%32)))+uint32(_c_F_italian_ISO_8859_1_stem[0]))))
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
	v524 = v519
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
	v519 = int32(0)
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
		v519 = v495
		goto L161
	} else {
		goto L167
	}
L167:
	;
	v500 = v498 - int32(97)
	if v500 < int32(0) {
		v519 = v495
		goto L161
	} else {
		goto L168
	}
L168:
	;
	v506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v500)>>(uint(int32(3))%32)))+uint32(_c_F_italian_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v506)>>(uint(v500&int32(7))%32))&int32(1) == int32(0) {
		v519 = v495
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
	v542 = v533
	goto L177
L176:
	;
	v577 = int32(1)
	goto L172
L177:
	;
	if v542 == v536 {
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
	v551 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v549+v542))))
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
	v559 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v553)>>(uint(int32(3))%32)))+uint32(_c_F_italian_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v559)>>(uint(v553&int32(7))%32))&int32(1) == int32(0) {
		goto L176
	} else {
		goto L184
	}
L184:
	;
	v568 = v542 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v568
	v542 = v568
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
	v617 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v610)>>(uint(int32(3))%32)))+uint32(_c_F_italian_ISO_8859_1_stem[0]))))
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
	v667 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v660)>>(uint(int32(3))%32)))+uint32(_c_F_italian_ISO_8859_1_stem[0]))))
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
	v716 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v709)>>(uint(int32(3))%32)))+uint32(_c_F_italian_ISO_8859_1_stem[0]))))
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
	v791 = v786
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
	v786 = int32(0)
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
		v786 = v762
		goto L240
	} else {
		goto L246
	}
L246:
	;
	v767 = v765 - int32(97)
	if v767 < int32(0) {
		v786 = v762
		goto L240
	} else {
		goto L247
	}
L247:
	;
	v773 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v767)>>(uint(int32(3))%32)))+uint32(_c_F_italian_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v773)>>(uint(v767&int32(7))%32))&int32(1) == int32(0) {
		v786 = v762
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
	v837 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v830)>>(uint(int32(3))%32)))+uint32(_c_F_italian_ISO_8859_1_stem[0]))))
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
	v876 = v857
	goto L274
L273:
	;
	v911 = int32(1)
	goto L269
L274:
	;
	if v876 == v870 {
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
	v885 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v883+v876))))
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
	v893 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v887)>>(uint(int32(3))%32)))+uint32(_c_F_italian_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v893)>>(uint(v887&int32(7))%32))&int32(1) == int32(0) {
		goto L273
	} else {
		goto L281
	}
L281:
	;
	v902 = v876 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v902
	v876 = v902
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
	v951 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v944)>>(uint(int32(3))%32)))+uint32(_c_F_italian_ISO_8859_1_stem[0]))))
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
	v990 = v971
	goto L305
L304:
	;
	v1025 = int32(1)
	goto L300
L305:
	;
	if v990 == v984 {
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
	v999 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v997+v990))))
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
	v1007 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1001)>>(uint(int32(3))%32)))+uint32(_c_F_italian_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v1007)>>(uint(v1001&int32(7))%32))&int32(1) == int32(0) {
		goto L304
	} else {
		goto L312
	}
L312:
	;
	v1016 = v990 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1016
	v990 = v1016
	goto L305
L314:
	;
	v1028 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1029 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1028))) = v1029 + v1025
	goto L252
L315:
	;
	v1095 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1095
	v1100 = F_find_among_b(m, l0, int32(_a_F_italian_ISO_8859_1_stem_11), int32(51))
	mBase = m.M
	v1101 = m.ExcPending
	if v1101 != 0 {
		goto L8
	} else {
		goto L331
	}
L316:
	;
	v1040 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1042 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1040+v1038))))
	if base.B2i32(v1042&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v1042)%32)&int32(_a_F_italian_ISO_8859_1_stem_12) == int32(0)) != 0 {
		goto L315
	} else {
		goto L317
	}
L317:
	;
	v1056 = F_find_among_b(m, l0, int32(_a_F_italian_ISO_8859_1_stem_13), int32(37))
	mBase = m.M
	v1057 = m.ExcPending
	if v1057 != 0 {
		goto L8
	} else {
		goto L318
	}
L318:
	;
	if v1056 == int32(0) {
		goto L315
	} else {
		goto L319
	}
L319:
	;
	v1060 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1060
	v1063 = v1060 - int32(1)
	v1064 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1063 <= v1064 {
		goto L315
	} else {
		goto L320
	}
L320:
	;
	v1066 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1068 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1066+v1063))))
	switch v1068 - int32(111) {
	case 0, 3:
		goto L321
	default:
		goto L315
	}
L321:
	;
	v1073 = F_find_among_b(m, l0, int32(_a_F_italian_ISO_8859_1_stem_14), int32(5))
	mBase = m.M
	v1074 = m.ExcPending
	if v1074 != 0 {
		goto L8
	} else {
		goto L322
	}
L322:
	;
	if v1073 == int32(0) {
		goto L315
	} else {
		goto L323
	}
L323:
	;
	v1077 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1078 = *(*int32)(unsafe.Add(mBase, uint32(v1077)+8))
	v1079 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v1079 < v1078 {
		goto L315
	} else {
		goto L324
	}
L324:
	;
	switch v1073 - int32(1) {
	case 0:
		goto L326
	case 1:
		goto L325
	default:
		goto L315
	}
L325:
	;
	v1089 = F_slice_from_s(m, l0, int32(1), int32(_a_F_italian_ISO_8859_1_stem_15))
	mBase = m.M
	v1090 = m.ExcPending
	if v1090 != 0 {
		goto L8
	} else {
		goto L329
	}
L326:
	;
	v1083 = F_slice_del(m, l0)
	mBase = m.M
	v1084 = m.ExcPending
	if v1084 != 0 {
		goto L8
	} else {
		goto L327
	}
L327:
	;
	if int32(0) <= v1083 {
		goto L315
	} else {
		goto L328
	}
L328:
	;
	v1632 = v1083
	goto L13
L329:
	;
	if v1089 < int32(0) {
		v1632 = v1089
		goto L13
	} else {
		goto L330
	}
L330:
	;
	goto L315
L331:
	;
	if v1100 == int32(0) {
		goto L332
	} else {
		goto L333
	}
L332:
	;
	v1104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1373 = v1104
	goto L44
L333:
	;
	goto L334
L334:
	;
	v1105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1105
	switch v1100 - int32(1) {
	case 0:
		goto L343
	case 1:
		goto L342
	case 2:
		goto L341
	case 3:
		goto L340
	case 4:
		goto L339
	case 5:
		goto L338
	case 6:
		goto L337
	case 7:
		goto L336
	case 8:
		goto L335
	default:
		goto L41
	}
L335:
	;
	v1301 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1302 = *(*int32)(unsafe.Add(mBase, uint32(v1301)))
	if v1105 < v1302 {
		v1373 = v1301
		goto L44
	} else {
		goto L399
	}
L336:
	;
	v1259 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1260 = *(*int32)(unsafe.Add(mBase, uint32(v1259)))
	if v1105 < v1260 {
		v1373 = v1259
		goto L44
	} else {
		goto L389
	}
L337:
	;
	v1186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1187 = *(*int32)(unsafe.Add(mBase, uint32(v1186)+4))
	if v1105 < v1187 {
		v1373 = v1186
		goto L44
	} else {
		goto L370
	}
L338:
	;
	v1179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1180 = *(*int32)(unsafe.Add(mBase, uint32(v1179)+8))
	if v1105 < v1180 {
		v1373 = v1179
		goto L44
	} else {
		goto L367
	}
L339:
	;
	v1170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1171 = *(*int32)(unsafe.Add(mBase, uint32(v1170)))
	if v1105 < v1171 {
		v1373 = v1170
		goto L44
	} else {
		goto L364
	}
L340:
	;
	v1161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1162 = *(*int32)(unsafe.Add(mBase, uint32(v1161)))
	if v1105 < v1162 {
		v1373 = v1161
		goto L44
	} else {
		goto L361
	}
L341:
	;
	v1152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1153 = *(*int32)(unsafe.Add(mBase, uint32(v1152)))
	if v1105 < v1153 {
		v1373 = v1152
		goto L44
	} else {
		goto L358
	}
L342:
	;
	v1116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1117 = *(*int32)(unsafe.Add(mBase, uint32(v1116)))
	if v1105 < v1117 {
		v1373 = v1116
		goto L44
	} else {
		goto L347
	}
L343:
	;
	v1109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1110 = *(*int32)(unsafe.Add(mBase, uint32(v1109)))
	if v1105 < v1110 {
		v1373 = v1109
		goto L44
	} else {
		goto L344
	}
L344:
	;
	v1112 = F_slice_del(m, l0)
	mBase = m.M
	v1113 = m.ExcPending
	if v1113 != 0 {
		goto L8
	} else {
		goto L345
	}
L345:
	;
	if int32(0) <= v1112 {
		goto L41
	} else {
		goto L346
	}
L346:
	;
	v1632 = v1112
	goto L13
L347:
	;
	v1119 = F_slice_del(m, l0)
	mBase = m.M
	v1120 = m.ExcPending
	if v1120 != 0 {
		goto L8
	} else {
		goto L348
	}
L348:
	;
	if v1119 < int32(0) {
		v1632 = v1119
		goto L13
	} else {
		goto L349
	}
L349:
	;
	v1123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1123
	v1125 = int32(2)
	v1127 = int32(0)
	v1130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1123-v1130 < v1125 {
		v1140 = v1127
		goto L351
	} else {
		goto L352
	}
L350:
	;
	if v1140 == int32(0) {
		goto L41
	} else {
		goto L354
	}
L351:
	;
	goto L350
L352:
	;
	v1133 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1136 = F_memcmp(m, v1133+v1123-v1125, int32(_a_F_italian_ISO_8859_1_stem_16), v1125)
	mBase = m.M
	if v1136 != 0 {
		v1140 = v1127
		goto L351
	} else {
		goto L353
	}
L353:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1123 - v1125
	v1140 = int32(1)
	goto L351
L354:
	;
	v1143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1143
	v1145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1146 = *(*int32)(unsafe.Add(mBase, uint32(v1145)))
	if v1143 < v1146 {
		goto L41
	} else {
		goto L355
	}
L355:
	;
	v1148 = F_slice_del(m, l0)
	mBase = m.M
	v1149 = m.ExcPending
	if v1149 != 0 {
		goto L8
	} else {
		goto L356
	}
L356:
	;
	if int32(0) <= v1148 {
		goto L41
	} else {
		goto L357
	}
L357:
	;
	v1632 = v1148
	goto L13
L358:
	;
	v1157 = F_slice_from_s(m, l0, int32(3), int32(_a_F_italian_ISO_8859_1_stem_17))
	mBase = m.M
	v1158 = m.ExcPending
	if v1158 != 0 {
		goto L8
	} else {
		goto L359
	}
L359:
	;
	if int32(0) <= v1157 {
		goto L41
	} else {
		goto L360
	}
L360:
	;
	v1632 = v1157
	goto L13
L361:
	;
	v1166 = F_slice_from_s(m, l0, int32(1), int32(_a_F_italian_ISO_8859_1_stem_18))
	mBase = m.M
	v1167 = m.ExcPending
	if v1167 != 0 {
		goto L8
	} else {
		goto L362
	}
L362:
	;
	if int32(0) <= v1166 {
		goto L41
	} else {
		goto L363
	}
L363:
	;
	v1632 = v1166
	goto L13
L364:
	;
	v1175 = F_slice_from_s(m, l0, int32(4), int32(_a_F_italian_ISO_8859_1_stem_19))
	mBase = m.M
	v1176 = m.ExcPending
	if v1176 != 0 {
		goto L8
	} else {
		goto L365
	}
L365:
	;
	if int32(0) <= v1175 {
		goto L41
	} else {
		goto L366
	}
L366:
	;
	v1632 = v1175
	goto L13
L367:
	;
	v1182 = F_slice_del(m, l0)
	mBase = m.M
	v1183 = m.ExcPending
	if v1183 != 0 {
		goto L8
	} else {
		goto L368
	}
L368:
	;
	if int32(0) <= v1182 {
		goto L41
	} else {
		goto L369
	}
L369:
	;
	v1632 = v1182
	goto L13
L370:
	;
	v1189 = F_slice_del(m, l0)
	mBase = m.M
	v1190 = m.ExcPending
	if v1190 != 0 {
		goto L8
	} else {
		goto L371
	}
L371:
	;
	if v1189 < int32(0) {
		v1632 = v1189
		goto L13
	} else {
		goto L372
	}
L372:
	;
	v1193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1193
	v1196 = v1193 - int32(1)
	v1197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1196 <= v1197 {
		goto L41
	} else {
		goto L373
	}
L373:
	;
	v1199 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1199+v1196))))
	if base.B2i32(v1201&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v1201)%32)&int32(_a_F_italian_ISO_8859_1_stem_20) == int32(0)) != 0 {
		goto L41
	} else {
		goto L374
	}
L374:
	;
	v1215 = F_find_among_b(m, l0, int32(_a_F_italian_ISO_8859_1_stem_21), int32(4))
	mBase = m.M
	v1216 = m.ExcPending
	if v1216 != 0 {
		goto L8
	} else {
		goto L375
	}
L375:
	;
	if v1215 == int32(0) {
		goto L41
	} else {
		goto L376
	}
L376:
	;
	v1219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1219
	v1221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1222 = *(*int32)(unsafe.Add(mBase, uint32(v1221)))
	if v1219 < v1222 {
		goto L41
	} else {
		goto L377
	}
L377:
	;
	v1224 = F_slice_del(m, l0)
	mBase = m.M
	v1225 = m.ExcPending
	if v1225 != 0 {
		goto L8
	} else {
		goto L378
	}
L378:
	;
	if v1224 < int32(0) {
		v1632 = v1224
		goto L13
	} else {
		goto L379
	}
L379:
	;
	if v1215 != int32(1) {
		goto L41
	} else {
		goto L380
	}
L380:
	;
	v1230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1230
	v1232 = int32(2)
	v1234 = int32(0)
	v1237 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1230-v1237 < v1232 {
		v1247 = v1234
		goto L382
	} else {
		goto L383
	}
L381:
	;
	if v1247 == int32(0) {
		goto L41
	} else {
		goto L385
	}
L382:
	;
	goto L381
L383:
	;
	v1240 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1243 = F_memcmp(m, v1240+v1230-v1232, int32(_a_F_italian_ISO_8859_1_stem_22), v1232)
	mBase = m.M
	if v1243 != 0 {
		v1247 = v1234
		goto L382
	} else {
		goto L384
	}
L384:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1230 - v1232
	v1247 = int32(1)
	goto L382
L385:
	;
	v1250 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1250
	v1252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1253 = *(*int32)(unsafe.Add(mBase, uint32(v1252)))
	if v1250 < v1253 {
		goto L41
	} else {
		goto L386
	}
L386:
	;
	v1255 = F_slice_del(m, l0)
	mBase = m.M
	v1256 = m.ExcPending
	if v1256 != 0 {
		goto L8
	} else {
		goto L387
	}
L387:
	;
	if int32(0) <= v1255 {
		goto L41
	} else {
		goto L388
	}
L388:
	;
	v1632 = v1255
	goto L13
L389:
	;
	v1262 = F_slice_del(m, l0)
	mBase = m.M
	v1263 = m.ExcPending
	if v1263 != 0 {
		goto L8
	} else {
		goto L390
	}
L390:
	;
	if v1262 < int32(0) {
		v1632 = v1262
		goto L13
	} else {
		goto L391
	}
L391:
	;
	v1266 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1266
	v1269 = v1266 - int32(1)
	v1270 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1269 <= v1270 {
		goto L41
	} else {
		goto L392
	}
L392:
	;
	v1272 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1272+v1269))))
	if base.B2i32(v1274&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v1274)%32)&int32(_a_F_italian_ISO_8859_1_stem_23) == int32(0)) != 0 {
		goto L41
	} else {
		goto L393
	}
L393:
	;
	v1288 = F_find_among_b(m, l0, int32(_a_F_italian_ISO_8859_1_stem_24), int32(3))
	mBase = m.M
	v1289 = m.ExcPending
	if v1289 != 0 {
		goto L8
	} else {
		goto L394
	}
L394:
	;
	if v1288 == int32(0) {
		goto L41
	} else {
		goto L395
	}
L395:
	;
	v1292 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1292
	v1294 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1295 = *(*int32)(unsafe.Add(mBase, uint32(v1294)))
	if v1292 < v1295 {
		goto L41
	} else {
		goto L396
	}
L396:
	;
	v1297 = F_slice_del(m, l0)
	mBase = m.M
	v1298 = m.ExcPending
	if v1298 != 0 {
		goto L8
	} else {
		goto L397
	}
L397:
	;
	if int32(0) <= v1297 {
		goto L41
	} else {
		goto L398
	}
L398:
	;
	v1632 = v1297
	goto L13
L399:
	;
	v1304 = F_slice_del(m, l0)
	mBase = m.M
	v1305 = m.ExcPending
	if v1305 != 0 {
		goto L8
	} else {
		goto L400
	}
L400:
	;
	if v1304 < int32(0) {
		v1632 = v1304
		goto L13
	} else {
		goto L401
	}
L401:
	;
	v1308 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1308
	v1310 = int32(2)
	v1312 = int32(0)
	v1315 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1308-v1315 < v1310 {
		v1325 = v1312
		goto L403
	} else {
		goto L404
	}
L402:
	;
	if v1325 == int32(0) {
		goto L41
	} else {
		goto L406
	}
L403:
	;
	goto L402
L404:
	;
	v1318 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1321 = F_memcmp(m, v1318+v1308-v1310, int32(_a_F_italian_ISO_8859_1_stem_25), v1310)
	mBase = m.M
	if v1321 != 0 {
		v1325 = v1312
		goto L403
	} else {
		goto L405
	}
L405:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1308 - v1310
	v1325 = int32(1)
	goto L403
L406:
	;
	v1328 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1328
	v1330 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1331 = *(*int32)(unsafe.Add(mBase, uint32(v1330)))
	if v1328 < v1331 {
		goto L41
	} else {
		goto L407
	}
L407:
	;
	v1333 = F_slice_del(m, l0)
	mBase = m.M
	v1334 = m.ExcPending
	if v1334 != 0 {
		goto L8
	} else {
		goto L408
	}
L408:
	;
	if v1333 < int32(0) {
		v1632 = v1333
		goto L13
	} else {
		goto L409
	}
L409:
	;
	v1337 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1337
	v1339 = int32(2)
	v1341 = int32(0)
	v1344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1337-v1344 < v1339 {
		v1354 = v1341
		goto L411
	} else {
		goto L412
	}
L410:
	;
	if v1354 == int32(0) {
		goto L41
	} else {
		goto L414
	}
L411:
	;
	goto L410
L412:
	;
	v1347 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1350 = F_memcmp(m, v1347+v1337-v1339, int32(_a_F_italian_ISO_8859_1_stem_26), v1339)
	mBase = m.M
	if v1350 != 0 {
		v1354 = v1341
		goto L411
	} else {
		goto L413
	}
L413:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1337 - v1339
	v1354 = int32(1)
	goto L411
L414:
	;
	v1357 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1357
	v1359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1360 = *(*int32)(unsafe.Add(mBase, uint32(v1359)))
	if v1357 < v1360 {
		goto L41
	} else {
		goto L415
	}
L415:
	;
	v1362 = F_slice_del(m, l0)
	mBase = m.M
	v1363 = m.ExcPending
	if v1363 != 0 {
		goto L8
	} else {
		goto L416
	}
L416:
	;
	if v1362 < int32(0) {
		v1632 = v1362
		goto L13
	} else {
		goto L417
	}
L417:
	;
	goto L41
L418:
	;
	if int32(0) <= v1368 {
		goto L42
	} else {
		goto L419
	}
L419:
	;
	v1632 = v1368
	goto L13
L420:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1374
	v1379 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1376
	v1383 = F_find_among_b(m, l0, int32(_a_F_italian_ISO_8859_1_stem_27), int32(87))
	mBase = m.M
	v1384 = m.ExcPending
	if v1384 != 0 {
		goto L8
	} else {
		goto L421
	}
L421:
	;
	if v1383 != 0 {
		goto L422
	} else {
		goto L423
	}
L422:
	;
	v1385 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1385
	v1387 = F_slice_del(m, l0)
	mBase = m.M
	v1388 = m.ExcPending
	if v1388 != 0 {
		goto L8
	} else {
		goto L425
	}
L423:
	;
	goto L424
L424:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1379
	goto L41
L425:
	;
	if v1387 < int32(0) {
		v1632 = v1387
		goto L13
	} else {
		goto L426
	}
L426:
	;
	goto L424
L427:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1490
	v1494 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1490 <= v1494 {
		v1572 = v1492
		goto L449
	} else {
		goto L450
	}
L428:
	;
	v1488 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1488
	v1490 = v1488
	v1492 = v1488
	goto L427
L429:
	;
	if v1452 != 0 {
		goto L428
	} else {
		goto L440
	}
L430:
	;
	v1452 = v1448
	goto L429
L431:
	;
	if v1397 <= v1409 {
		goto L433
	} else {
		goto L434
	}
L432:
	;
	v1448 = int32(0)
	goto L430
L433:
	;
	v1452 = int32(-1)
	goto L429
L434:
	;
	goto L435
L435:
	;
	v1421 = int32(1)
	v1422 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1426 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1422+v1397-v1421))))
	if int32(242) < v1426 {
		v1448 = v1421
		goto L430
	} else {
		goto L436
	}
L436:
	;
	v1428 = v1426 - int32(97)
	if v1428 < int32(0) {
		v1448 = v1421
		goto L430
	} else {
		goto L437
	}
L437:
	;
	v1434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1428)>>(uint(int32(3))%32)))+uint32(_c_F_italian_ISO_8859_1_stem[1]))))
	if int32(base.Ui32(v1434)>>(uint(v1428&int32(7))%32))&int32(1) == int32(0) {
		v1448 = v1421
		goto L430
	} else {
		goto L438
	}
L438:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1397 - int32(1)
	goto L439
L439:
	;
	goto L432
L440:
	;
	v1453 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1453
	v1455 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1456 = *(*int32)(unsafe.Add(mBase, uint32(v1455)+8))
	if v1453 < v1456 {
		goto L428
	} else {
		goto L441
	}
L441:
	;
	v1458 = F_slice_del(m, l0)
	mBase = m.M
	v1459 = m.ExcPending
	if v1459 != 0 {
		goto L8
	} else {
		goto L442
	}
L442:
	;
	if v1458 < int32(0) {
		v1632 = v1458
		goto L13
	} else {
		goto L443
	}
L443:
	;
	v1462 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1462
	v1464 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1462 <= v1464 {
		goto L428
	} else {
		goto L444
	}
L444:
	;
	v1466 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1470 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1466+v1462-int32(1)))))
	if v1470 != int32(105) {
		goto L428
	} else {
		goto L445
	}
L445:
	;
	v1474 = v1462 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1474
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1474
	v1477 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1478 = *(*int32)(unsafe.Add(mBase, uint32(v1477)+8))
	if v1462 <= v1478 {
		goto L428
	} else {
		goto L446
	}
L446:
	;
	v1480 = F_slice_del(m, l0)
	mBase = m.M
	v1481 = m.ExcPending
	if v1481 != 0 {
		goto L8
	} else {
		goto L447
	}
L447:
	;
	if v1480 < int32(0) {
		v1632 = v1480
		goto L13
	} else {
		goto L448
	}
L448:
	;
	v1484 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1485 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1490 = v1484
	v1492 = v1485
	goto L427
L449:
	;
	v1573 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1573
	v1576 = v1573
	v1578 = v1572
	goto L468
L450:
	;
	v1496 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1500 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1496+v1490-int32(1)))))
	if v1500 != int32(104) {
		v1572 = v1492
		goto L449
	} else {
		goto L451
	}
L451:
	;
	v1504 = v1490 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1504
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1504
	v1516 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L455
L452:
	;
	v1569 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1572 = v1569
	goto L449
L453:
	;
	if v1559 != 0 {
		goto L452
	} else {
		goto L464
	}
L454:
	;
	v1559 = v1555
	goto L453
L455:
	;
	if v1504 <= v1516 {
		goto L457
	} else {
		goto L458
	}
L456:
	;
	v1555 = int32(0)
	goto L454
L457:
	;
	v1559 = int32(-1)
	goto L453
L458:
	;
	goto L459
L459:
	;
	v1528 = int32(1)
	v1529 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1533 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1529+v1504-v1528))))
	if int32(103) < v1533 {
		v1555 = v1528
		goto L454
	} else {
		goto L460
	}
L460:
	;
	v1535 = v1533 - int32(99)
	if v1535 < int32(0) {
		v1555 = v1528
		goto L454
	} else {
		goto L461
	}
L461:
	;
	v1541 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1535)>>(uint(int32(3))%32)))+uint32(_c_F_italian_ISO_8859_1_stem[2]))))
	if int32(base.Ui32(v1541)>>(uint(v1535&int32(7))%32))&int32(1) == int32(0) {
		v1555 = v1528
		goto L454
	} else {
		goto L462
	}
L462:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1504 - int32(1)
	goto L463
L463:
	;
	goto L456
L464:
	;
	v1560 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1561 = *(*int32)(unsafe.Add(mBase, uint32(v1560)+8))
	v1562 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v1562 < v1561 {
		goto L452
	} else {
		goto L465
	}
L465:
	;
	v1564 = F_slice_del(m, l0)
	mBase = m.M
	v1565 = m.ExcPending
	if v1565 != 0 {
		goto L8
	} else {
		goto L466
	}
L466:
	;
	if v1564 < int32(0) {
		v1632 = v1564
		goto L13
	} else {
		goto L467
	}
L467:
	;
	goto L452
L468:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1576
	if v1578 <= v1576 {
		goto L473
	} else {
		goto L474
	}
L469:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1573
	v1632 = int32(1)
	goto L13
L470:
	;
	goto L469
L471:
	;
	v1627 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1628 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1576 = v1628
	v1578 = v1627
	goto L468
L472:
	;
	if v1618 <= v1617 {
		goto L470
	} else {
		goto L484
	}
L473:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1576
	v1617 = v1576
	v1618 = v1578
	goto L472
L474:
	;
	v1582 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1584 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1582+v1576))))
	v1586 = v1584 - int32(73)
	v1587 = int32(0)
	if base.B2i32(v1586 == v1587)|base.B2i32(v1586 == int32(12)) == v1587 {
		goto L473
	} else {
		goto L475
	}
L475:
	;
	v1596 = F_find_among(m, l0, int32(_a_F_italian_ISO_8859_1_stem_28), int32(3))
	mBase = m.M
	v1597 = m.ExcPending
	if v1597 != 0 {
		goto L8
	} else {
		goto L476
	}
L476:
	;
	v1598 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1598
	switch v1596 - int32(1) {
	case 0:
		goto L478
	case 1:
		goto L477
	case 2:
		goto L479
	default:
		goto L471
	}
L477:
	;
	v1611 = F_slice_from_s(m, l0, int32(1), int32(_a_F_italian_ISO_8859_1_stem_29))
	mBase = m.M
	v1612 = m.ExcPending
	if v1612 != 0 {
		goto L8
	} else {
		goto L482
	}
L478:
	;
	v1605 = F_slice_from_s(m, l0, int32(1), int32(_a_F_italian_ISO_8859_1_stem_30))
	mBase = m.M
	v1606 = m.ExcPending
	if v1606 != 0 {
		goto L8
	} else {
		goto L480
	}
L479:
	;
	v1602 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1617 = v1598
	v1618 = v1602
	goto L472
L480:
	;
	if int32(0) <= v1605 {
		goto L471
	} else {
		goto L481
	}
L481:
	;
	v1632 = v1605
	goto L13
L482:
	;
	if int32(0) <= v1611 {
		goto L471
	} else {
		goto L483
	}
L483:
	;
	v1632 = v1611
	goto L13
L484:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1617 + int32(1)
	goto L471
}
