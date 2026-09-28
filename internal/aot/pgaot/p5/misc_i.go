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
	var v15 int32
	_ = v15
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
		F_ResourceOwnerRemember(m, v11, base.I64_extend_i32_u(l0), int32(_a_F_IncrTupleDescRefCount_0))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
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
	var v36 int64
	_ = v36
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int64
	_ = v97
	var v100 int64
	_ = v100
	var v108 int32
	_ = v108
	var v110 int64
	_ = v110
	var v112 int64
	_ = v112
	var v118 int64
	_ = v118
	var v121 int64
	_ = v121
	var v122 int64
	_ = v122
	var v125 int64
	_ = v125
	var v126 int64
	_ = v126
	var v127 int64
	_ = v127
	var v130 int64
	_ = v130
	var v131 int64
	_ = v131
	var v132 int64
	_ = v132
	var v137 int64
	_ = v137
	var v142 int64
	_ = v142
	var v147 int64
	_ = v147
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v182 int64
	_ = v182
	var v183 int32
	_ = v183
	var v192 int64
	_ = v192
	var v194 int64
	_ = v194
	var v197 int32
	_ = v197
	var v203 int32
	_ = v203
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
	v32 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitProcessGlobals[2])))
	if v32 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v45 = int32(16)
	v46 = int32(0)
	v50 = m.G0
	v52 = v50 - v45
	m.G0 = v52
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = v46
	v58 = F_open(m, int32(_a_F_InitProcessGlobals_0), v46, v52)
	mBase = m.M
	if v58 != int32(-1) {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	v36 = int64(0)
	*(*int64)(unsafe.Add(mBase, _c_F_InitProcessGlobals[3])) = v36
	*(*int64)(unsafe.Add(mBase, _c_F_InitProcessGlobals[4])) = v36
	v42 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_InitProcessGlobals[2])) = uint8(v42)
	goto L6
L5:
	;
	goto L6
L6:
	;
	goto L3
L7:
	;
	v154 = Fn14349(m, int64(32))
	mBase = m.M
	goto L30
L8:
	;
	if v91 != 0 {
		goto L21
	} else {
		goto L22
	}
L9:
	;
	goto L13
L10:
	;
	v91 = v46
	goto L11
L11:
	;
	m.G0 = v52 + int32(16)
	goto L8
L12:
	;
	v86 = F_close(m, v58)
	mBase = m.M
	v91 = v84
	goto L11
L13:
	;
	v64 = int32(_a_F_InitProcessGlobals_1)
	v65 = v45
	goto L14
L14:
	;
	v70 = F_read(m, v58, v64, v65)
	mBase = m.M
	if v70 <= int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v84 = int32(1)
	goto L12
L16:
	;
	v74 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcessGlobals[5]))
	if v74 == int32(27) {
		goto L14
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v79 = v65 - v70
	if v79 != 0 {
		v64 = v64 + v70
		v65 = v79
		goto L14
	} else {
		goto L20
	}
L19:
	;
	v84 = int32(0)
	goto L12
L20:
	;
	goto L15
L21:
	;
	v96 = int32(_a_F_InitProcessGlobals_1)
	v97 = *(*int64)(unsafe.Add(mBase, _c_F_InitProcessGlobals[6]))
	if v97 != int64(0) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	goto L23
L23:
	;
	v108 = int32(_a_F_InitProcessGlobals_1)
	v110 = int64(*(*int32)(unsafe.Add(mBase, _c_F_InitProcessGlobals[7])))
	v112 = *(*int64)(unsafe.Add(mBase, _c_F_InitProcessGlobals[0]))
	v118 = v110 ^ v112<<(uint(int64(12))%64) ^ int64(base.Ui64(v112)>>(uint(int64(20))%64))
	v121 = v118 + int64(4354685564936845354)
	v122 = int64(30)
	v125 = int64(-4658895280553007687)
	v126 = (int64(base.Ui64(v121)>>(uint(v122)%64)) ^ v121) * v125
	v127 = int64(27)
	v130 = int64(-7723592293110705685)
	v131 = (int64(base.Ui64(v126)>>(uint(v127)%64)) ^ v126) * v130
	v132 = int64(31)
	*(*int64)(unsafe.Add(mBase, _c_F_InitProcessGlobals[8])) = int64(base.Ui64(v131)>>(uint(v132)%64)) ^ v131
	v137 = v118 - int64(7046029254386353131)
	v142 = (int64(base.Ui64(v137)>>(uint(v122)%64)) ^ v137) * v125
	v147 = (int64(base.Ui64(v142)>>(uint(v127)%64)) ^ v142) * v130
	*(*int64)(unsafe.Add(mBase, _c_F_InitProcessGlobals[6])) = int64(base.Ui64(v147)>>(uint(v132)%64)) ^ v147
	goto L29
L24:
	;
	goto L7
L25:
	;
	goto L24
L26:
	;
	v100 = *(*int64)(unsafe.Add(mBase, _c_F_InitProcessGlobals[8]))
	if v100 != int64(0) {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_InitProcessGlobals[8])) = int64(1442695040888963407)
	*(*int64)(unsafe.Add(mBase, _c_F_InitProcessGlobals[6])) = int64(6364136223846793005)
	goto L25
L29:
	;
	goto L7
L30:
	;
	v156 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcessGlobals[9]))
	if v156 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	return
L32:
	;
	v160 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcessGlobals[10]))
	*(*int32)(unsafe.Add(mBase, uint32(v160))) = v154
	goto L31
L33:
	;
	goto L34
L34:
	;
	v163 = int32(3)
	if v156 == int32(7) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v168 = v163
	goto L37
L36:
	;
	v168 = int32(1)
	goto L37
L37:
	;
	if v156 == int32(31) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v171 = v163
	goto L40
L39:
	;
	v171 = v168
	goto L40
L40:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InitProcessGlobals[11])) = v171
	v174 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_InitProcessGlobals[12])) = v174
	v177 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcessGlobals[10]))
	if v174 < v156 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v182 = base.I64_extend_i32_u(v154)
	v183 = int32(0)
	goto L44
L42:
	;
	goto L43
L43:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v177)))
	*(*int32)(unsafe.Add(mBase, uint32(v177))) = v203 | int32(1)
	goto L31
L44:
	;
	v192 = v182*int64(6364136223846793005) + int64(1)
	v194 = int64(base.Ui64(v192) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v177+v183<<(uint(int32(2))%32)))) = uint32(v194)
	v197 = v183 + int32(1)
	if v197 != v156 {
		v182 = v192
		v183 = v197
		goto L44
	} else {
		goto L46
	}
L45:
	;
	goto L43
L46:
	;
	goto L45
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
	var v31 int32
	_ = v31
	var v42 int32
	_ = v42
	var v43 int64
	_ = v43
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	v9 = m.G0
	v11 = v9 - int32(96)
	m.G0 = v11
	if l1 != 0 {
		*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = int64(0)
		*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = l4
		v23 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v11)+40)) = uint8(v23)
		*(*uint8)(unsafe.Add(mBase, uint32(v11)+88)) = uint8(v23)
		*(*uint8)(unsafe.Add(mBase, uint32(v11)+72)) = uint8(v23)
		*(*uint8)(unsafe.Add(mBase, uint32(v11)+56)) = uint8(v23)
		v31 = int32(3)
		*(*uint16)(unsafe.Add(mBase, uint32(v11)+42)) = uint16(v31)
		*(*int64)(unsafe.Add(mBase, uint32(v11)+80)) = base.I64_extend_i32_s(l3)
		*(*int64)(unsafe.Add(mBase, uint32(v11)+64)) = base.I64_extend_i32_u(l2)
		*(*int64)(unsafe.Add(mBase, uint32(v11)+48)) = base.I64_extend_i32_u(l1)
		*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = l0
		v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v43 = m.T0[v42].(func(*base.Module, int32) int64)(m, v11+int32(24))
		mBase = m.M
		v46 = m.ExcPending
		if v46 != 0 {
			return int32(0)
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(l5))) = v43
			if l4 == int32(0) {
				v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+40)))
				if l1 == int32(0) {
					v59 = int32(1)
					if v56&v59 != 0 {
						v98 = v59
						m.G0 = v11 + int32(96)
						return v98
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return int32(0)
						} else {
							v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v11))) = v66
							F_errmsg_internal(m, int32(_a_F_InputFunctionCallSafe_0), v11)
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_InputFunctionCallSafe_1), int32(1619), int32(_a_F_InputFunctionCallSafe_2))
								mBase = m.M
								v75 = m.ExcPending
								if v75 != 0 {
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
					v76 = int32(1)
					if v56&v76 == int32(0) {
						v98 = v76
						m.G0 = v11 + int32(96)
						return v98
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v84 = m.ExcPending
						if v84 != 0 {
							return int32(0)
						} else {
							v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v85
							F_errmsg_internal(m, int32(_a_F_InputFunctionCallSafe_3), v11+int32(16))
							mBase = m.M
							v91 = m.ExcPending
							if v91 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_InputFunctionCallSafe_1), int32(1625), int32(_a_F_InputFunctionCallSafe_2))
								mBase = m.M
								v96 = m.ExcPending
								if v96 != 0 {
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
				v50 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
				if v50 != int32(453) {
					v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+40)))
					if l1 == int32(0) {
						v59 = int32(1)
						if v56&v59 != 0 {
							v98 = v59
							m.G0 = v11 + int32(96)
							return v98
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return int32(0)
							} else {
								v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v11))) = v66
								F_errmsg_internal(m, int32(_a_F_InputFunctionCallSafe_0), v11)
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_InputFunctionCallSafe_1), int32(1619), int32(_a_F_InputFunctionCallSafe_2))
									mBase = m.M
									v75 = m.ExcPending
									if v75 != 0 {
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
						v76 = int32(1)
						if v56&v76 == int32(0) {
							v98 = v76
							m.G0 = v11 + int32(96)
							return v98
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
								return int32(0)
							} else {
								v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v85
								F_errmsg_internal(m, int32(_a_F_InputFunctionCallSafe_3), v11+int32(16))
								mBase = m.M
								v91 = m.ExcPending
								if v91 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_InputFunctionCallSafe_1), int32(1625), int32(_a_F_InputFunctionCallSafe_2))
									mBase = m.M
									v96 = m.ExcPending
									if v96 != 0 {
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
					v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+4)))
					if v54 != 0 {
						v98 = int32(0)
						m.G0 = v11 + int32(96)
						return v98
					} else {
						v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+40)))
						if l1 == int32(0) {
							v59 = int32(1)
							if v56&v59 != 0 {
								v98 = v59
								m.G0 = v11 + int32(96)
								return v98
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return int32(0)
								} else {
									v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v11))) = v66
									F_errmsg_internal(m, int32(_a_F_InputFunctionCallSafe_0), v11)
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_InputFunctionCallSafe_1), int32(1619), int32(_a_F_InputFunctionCallSafe_2))
										mBase = m.M
										v75 = m.ExcPending
										if v75 != 0 {
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
							v76 = int32(1)
							if v56&v76 == int32(0) {
								v98 = v76
								m.G0 = v11 + int32(96)
								return v98
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v84 = m.ExcPending
								if v84 != 0 {
									return int32(0)
								} else {
									v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v85
									F_errmsg_internal(m, int32(_a_F_InputFunctionCallSafe_3), v11+int32(16))
									mBase = m.M
									v91 = m.ExcPending
									if v91 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_InputFunctionCallSafe_1), int32(1625), int32(_a_F_InputFunctionCallSafe_2))
										mBase = m.M
										v96 = m.ExcPending
										if v96 != 0 {
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
			*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = int64(0)
			*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = l4
			v23 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v11)+40)) = uint8(v23)
			*(*uint8)(unsafe.Add(mBase, uint32(v11)+88)) = uint8(v23)
			*(*uint8)(unsafe.Add(mBase, uint32(v11)+72)) = uint8(v23)
			*(*uint8)(unsafe.Add(mBase, uint32(v11)+56)) = uint8(v23)
			v31 = int32(3)
			*(*uint16)(unsafe.Add(mBase, uint32(v11)+42)) = uint16(v31)
			*(*int64)(unsafe.Add(mBase, uint32(v11)+80)) = base.I64_extend_i32_s(l3)
			*(*int64)(unsafe.Add(mBase, uint32(v11)+64)) = base.I64_extend_i32_u(l2)
			*(*int64)(unsafe.Add(mBase, uint32(v11)+48)) = base.I64_extend_i32_u(l1)
			*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = l0
			v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v43 = m.T0[v42].(func(*base.Module, int32) int64)(m, v11+int32(24))
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return int32(0)
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(l5))) = v43
				if l4 == int32(0) {
					v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+40)))
					if l1 == int32(0) {
						v59 = int32(1)
						if v56&v59 != 0 {
							v98 = v59
							m.G0 = v11 + int32(96)
							return v98
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return int32(0)
							} else {
								v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v11))) = v66
								F_errmsg_internal(m, int32(_a_F_InputFunctionCallSafe_0), v11)
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_InputFunctionCallSafe_1), int32(1619), int32(_a_F_InputFunctionCallSafe_2))
									mBase = m.M
									v75 = m.ExcPending
									if v75 != 0 {
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
						v76 = int32(1)
						if v56&v76 == int32(0) {
							v98 = v76
							m.G0 = v11 + int32(96)
							return v98
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
								return int32(0)
							} else {
								v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v85
								F_errmsg_internal(m, int32(_a_F_InputFunctionCallSafe_3), v11+int32(16))
								mBase = m.M
								v91 = m.ExcPending
								if v91 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_InputFunctionCallSafe_1), int32(1625), int32(_a_F_InputFunctionCallSafe_2))
									mBase = m.M
									v96 = m.ExcPending
									if v96 != 0 {
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
					v50 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
					if v50 != int32(453) {
						v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+40)))
						if l1 == int32(0) {
							v59 = int32(1)
							if v56&v59 != 0 {
								v98 = v59
								m.G0 = v11 + int32(96)
								return v98
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return int32(0)
								} else {
									v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v11))) = v66
									F_errmsg_internal(m, int32(_a_F_InputFunctionCallSafe_0), v11)
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_InputFunctionCallSafe_1), int32(1619), int32(_a_F_InputFunctionCallSafe_2))
										mBase = m.M
										v75 = m.ExcPending
										if v75 != 0 {
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
							v76 = int32(1)
							if v56&v76 == int32(0) {
								v98 = v76
								m.G0 = v11 + int32(96)
								return v98
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v84 = m.ExcPending
								if v84 != 0 {
									return int32(0)
								} else {
									v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v85
									F_errmsg_internal(m, int32(_a_F_InputFunctionCallSafe_3), v11+int32(16))
									mBase = m.M
									v91 = m.ExcPending
									if v91 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_InputFunctionCallSafe_1), int32(1625), int32(_a_F_InputFunctionCallSafe_2))
										mBase = m.M
										v96 = m.ExcPending
										if v96 != 0 {
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
						v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+4)))
						if v54 != 0 {
							v98 = int32(0)
							m.G0 = v11 + int32(96)
							return v98
						} else {
							v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+40)))
							if l1 == int32(0) {
								v59 = int32(1)
								if v56&v59 != 0 {
									v98 = v59
									m.G0 = v11 + int32(96)
									return v98
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return int32(0)
									} else {
										v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v11))) = v66
										F_errmsg_internal(m, int32(_a_F_InputFunctionCallSafe_0), v11)
										mBase = m.M
										v70 = m.ExcPending
										if v70 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_InputFunctionCallSafe_1), int32(1619), int32(_a_F_InputFunctionCallSafe_2))
											mBase = m.M
											v75 = m.ExcPending
											if v75 != 0 {
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
								v76 = int32(1)
								if v56&v76 == int32(0) {
									v98 = v76
									m.G0 = v11 + int32(96)
									return v98
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v84 = m.ExcPending
									if v84 != 0 {
										return int32(0)
									} else {
										v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v85
										F_errmsg_internal(m, int32(_a_F_InputFunctionCallSafe_3), v11+int32(16))
										mBase = m.M
										v91 = m.ExcPending
										if v91 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_InputFunctionCallSafe_1), int32(1625), int32(_a_F_InputFunctionCallSafe_2))
											mBase = m.M
											v96 = m.ExcPending
											if v96 != 0 {
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
			*(*int64)(unsafe.Add(mBase, uint32(l5))) = int64(0)
			v98 = v13
			m.G0 = v11 + int32(96)
			return v98
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
	*(*int32)(unsafe.Add(mBase, uint32(v240))) = int32(167772180)
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
func F_i2toi4(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	v2 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+24)))
	return v2
}
func F_icnlikesel(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14371(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
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
				F_errfinish(m, int32(_a_F_icu_validate_locale_1), int32(1968), int32(_a_F_icu_validate_locale_2))
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
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+53)))
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
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+53)))
	if v19 == int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
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
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+72))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+22)))
	v32 = v27 + v28
	v33 = int32(1)
	goto L7
L7:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	if int32(0) < v34 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v52 = int32(-64)
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
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+72))
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
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v64 < v113 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0-v52+v64<<(uint(int32(2))%32))))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)+72))
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
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	if v83 < v156 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l1-v52+v83<<(uint(int32(2))%32))))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v161)+72))
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
func F_idx(m *base.Module, l0 int32) int64 {
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
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = F_pg_detoast_datum(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
		if v9 != 0 {
			v10 = F_array_contains_nulls(m, v5)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int64(0)
			} else {
				if v10 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(67108994))
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_idx_0), int32(0))
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_idx_1), int32(268), int32(_a_F_idx_2))
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
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
					v12 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
					v15 = F_ArrayGetNItemsSafe(m, v12, v5+int32(16))
					mBase = m.M
					v16 = m.ExcPending
					if v16 != 0 {
						return int64(0)
					} else {
						if v15 != 0 {
							v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							v18 = F_intarray_match_first(m, v5, v17)
							mBase = m.M
							v19 = m.ExcPending
							if v19 != 0 {
								return int64(0)
							} else {
								v22 = base.I64_extend_i32_s(v18)
								v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								if v23 != v5 {
									F_pfree(m, v5)
									mBase = m.M
									v26 = m.ExcPending
									if v26 != 0 {
										return int64(0)
									} else {
										return v22
									}
								} else {
									return v22
								}
							}
						} else {
							v22 = int64(0)
							v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							if v23 != v5 {
								F_pfree(m, v5)
								mBase = m.M
								v26 = m.ExcPending
								if v26 != 0 {
									return int64(0)
								} else {
									return v22
								}
							} else {
								return v22
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
				return int64(0)
			} else {
				if v15 != 0 {
					v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					v18 = F_intarray_match_first(m, v5, v17)
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return int64(0)
					} else {
						v22 = base.I64_extend_i32_s(v18)
						v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						if v23 != v5 {
							F_pfree(m, v5)
							mBase = m.M
							v26 = m.ExcPending
							if v26 != 0 {
								return int64(0)
							} else {
								return v22
							}
						} else {
							return v22
						}
					}
				} else {
					v22 = int64(0)
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					if v23 != v5 {
						F_pfree(m, v5)
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return int64(0)
						} else {
							return v22
						}
					} else {
						return v22
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
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
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
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v308 int32
	_ = v308
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v408 int32
	_ = v408
	var v421 int32
	_ = v421
	var v426 int32
	_ = v426
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v444 int32
	_ = v444
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v494 int32
	_ = v494
	var v507 int32
	_ = v507
	var v512 int32
	_ = v512
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v566 int32
	_ = v566
	var v579 int32
	_ = v579
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v596 int32
	_ = v596
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v624 int32
	_ = v624
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v645 int32
	_ = v645
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v658 int32
	_ = v658
	var v664 int32
	_ = v664
	v2 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v2
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
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v35)>>(uint(int32(3))%32)))+uint32(_c_F_indonesian_ISO_8859_1_stem[0]))))
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
	v64 = v58
	goto L19
L17:
	;
	goto L18
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v9
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v132 < int32(3) {
		v664 = v2
		goto L37
	} else {
		goto L38
	}
L19:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v68 = v67 + v64
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v68
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v70 + int32(1)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v82 < v68 {
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
		v64 = v122
		goto L19
	} else {
		goto L36
	}
L22:
	;
	v84 = v68
	goto L24
L23:
	;
	v84 = v82
	goto L24
L24:
	;
	v91 = v68
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
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v99)>>(uint(int32(3))%32)))+uint32(_c_F_indonesian_ISO_8859_1_stem[0]))))
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
	return v664
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v9
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v138
	if v138-int32(2) <= v9 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v184
	v188 = v184 - int32(1)
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v188 <= v189 {
		goto L52
	} else {
		goto L53
	}
L40:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v180
	v184 = v180
	v185 = v2
	goto L39
L41:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144+v138-int32(1)))))
	switch v148 - int32(104) {
	case 0, 6:
		goto L42
	default:
		goto L40
	}
L42:
	;
	v154 = F_find_among_b(m, l0, int32(_a_F_indonesian_ISO_8859_1_stem_0), int32(3), int32(0))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v174
	if int32(3) <= v172 {
		v184 = v174
		v185 = v173
		goto L39
	} else {
		goto L50
	}
L44:
	;
	return int32(0)
L45:
	;
	if v154 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v172 = v160
	v173 = v2
	goto L43
L47:
	;
	goto L48
L48:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v161
	v163 = F_slice_del(m, l0)
	mBase = m.M
	if v163 < int32(0) {
		v664 = v163
		goto L37
	} else {
		goto L49
	}
L49:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v167 = int32(1)
	v168 = v166 - v167
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v168
	v172 = v168
	v173 = v167
	goto L43
L50:
	;
	return int32(0)
L51:
	;
	v230 = int32(0)
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v231
	v234 = v231 + int32(1)
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v235 <= v234 {
		v601 = v230
		goto L64
	} else {
		goto L65
	}
L52:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v225
	v229 = v225
	goto L51
L53:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191+v188))))
	if base.B2i32(v193 != int32(117))&base.B2i32(v193 != int32(97)) != 0 {
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v202 = F_find_among_b(m, l0, int32(_a_F_indonesian_ISO_8859_1_stem_1), int32(3), int32(0))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L44
	} else {
		goto L56
	}
L55:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v218
	if int32(3) <= v217 {
		v229 = v218
		goto L51
	} else {
		goto L61
	}
L56:
	;
	if v202 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v217 = v206
	goto L55
L58:
	;
	goto L59
L59:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v207
	v209 = F_slice_del(m, l0)
	mBase = m.M
	if v209 < int32(0) {
		v664 = v209
		goto L37
	} else {
		goto L60
	}
L60:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v214 = v212 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v214
	v217 = v214
	goto L55
L61:
	;
	return int32(0)
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v229
	v664 = int32(1)
	goto L37
L63:
	;
	if v607 != 0 {
		goto L161
	} else {
		goto L162
	}
L64:
	;
	v607 = v601
	goto L63
L65:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237+v234))))
	switch v239 - int32(101) {
	case 0, 4:
		goto L66
	default:
		v601 = v230
		goto L64
	}
L66:
	;
	v245 = F_find_among(m, l0, int32(_a_F_indonesian_ISO_8859_1_stem_2), int32(10), int32(0))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L44
	} else {
		goto L67
	}
L67:
	;
	if v245 == int32(0) {
		v601 = v230
		goto L64
	} else {
		goto L68
	}
L68:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v249
	v251 = int32(1)
	switch v245 - v251 {
	case 0:
		goto L74
	case 1:
		goto L73
	case 2:
		goto L72
	case 3:
		goto L71
	case 4:
		goto L70
	case 5:
		goto L69
	default:
		v601 = v251
		goto L64
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(3)
	v528 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v528 - int32(1)
	v540 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v541 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v541 < v540 {
		goto L141
	} else {
		goto L142
	}
L70:
	;
	v454 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v454
	v456 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v456 - v454
	v468 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v469 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v469 < v468 {
		goto L120
	} else {
		goto L121
	}
L71:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v249 == v364 {
		goto L98
	} else {
		goto L99
	}
L72:
	;
	v354 = F_slice_del(m, l0)
	mBase = m.M
	if v354 < int32(0) {
		v601 = v354
		goto L64
	} else {
		goto L97
	}
L73:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v249 == v264 {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	v254 = F_slice_del(m, l0)
	mBase = m.M
	if v254 < int32(0) {
		v601 = v254
		goto L64
	} else {
		goto L75
	}
L75:
	;
	v257 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v257
	v259 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v259 - v257
	v607 = v257
	goto L63
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v249
	v344 = F_slice_del(m, l0)
	mBase = m.M
	if v344 < int32(0) {
		v601 = v344
		goto L64
	} else {
		goto L96
	}
L77:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v266+v249))))
	if v268 != int32(121) {
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v272 = v249 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v272
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v283 < v272 {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	if v326 != 0 {
		goto L76
	} else {
		goto L93
	}
L80:
	;
	v285 = v272
	goto L82
L81:
	;
	v285 = v283
	goto L82
L82:
	;
	goto L84
L83:
	;
	v326 = v321
	goto L79
L84:
	;
	if v272 == v285 {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v321 = int32(0)
	goto L83
L86:
	;
	v326 = int32(-1)
	goto L79
L87:
	;
	goto L88
L88:
	;
	v297 = int32(1)
	v298 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v298+v272))))
	if int32(117) < v300 {
		v321 = v297
		goto L83
	} else {
		goto L89
	}
L89:
	;
	v302 = v300 - int32(97)
	if v302 < int32(0) {
		v321 = v297
		goto L83
	} else {
		goto L90
	}
L90:
	;
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v302)>>(uint(int32(3))%32)))+uint32(_c_F_indonesian_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v308)>>(uint(v302&int32(7))%32))&int32(1) == int32(0) {
		v321 = v297
		goto L83
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v249 + int32(2)
	goto L92
L92:
	;
	goto L85
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v272
	v331 = F_slice_from_s(m, l0, int32(1), int32(_a_F_indonesian_ISO_8859_1_stem_3))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L44
	} else {
		goto L94
	}
L94:
	;
	if v331 < int32(0) {
		v601 = v331
		goto L64
	} else {
		goto L95
	}
L95:
	;
	v335 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v335
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v337 - v335
	v607 = v335
	goto L63
L96:
	;
	v347 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v347
	v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v349 - v347
	v607 = v347
	goto L63
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(3)
	v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v360 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v359 - v360
	v607 = v360
	goto L63
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v249
	v444 = F_slice_del(m, l0)
	mBase = m.M
	if v444 < int32(0) {
		v601 = v444
		goto L64
	} else {
		goto L118
	}
L99:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v366+v249))))
	if v368 != int32(121) {
		goto L98
	} else {
		goto L100
	}
L100:
	;
	v372 = v249 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v372
	v383 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v383 < v372 {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	if v426 != 0 {
		goto L98
	} else {
		goto L115
	}
L102:
	;
	v385 = v372
	goto L104
L103:
	;
	v385 = v383
	goto L104
L104:
	;
	goto L106
L105:
	;
	v426 = v421
	goto L101
L106:
	;
	if v372 == v385 {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	v421 = int32(0)
	goto L105
L108:
	;
	v426 = int32(-1)
	goto L101
L109:
	;
	goto L110
L110:
	;
	v397 = int32(1)
	v398 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v398+v372))))
	if int32(117) < v400 {
		v421 = v397
		goto L105
	} else {
		goto L111
	}
L111:
	;
	v402 = v400 - int32(97)
	if v402 < int32(0) {
		v421 = v397
		goto L105
	} else {
		goto L112
	}
L112:
	;
	v408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v402)>>(uint(int32(3))%32)))+uint32(_c_F_indonesian_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v408)>>(uint(v402&int32(7))%32))&int32(1) == int32(0) {
		v421 = v397
		goto L105
	} else {
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v249 + int32(2)
	goto L114
L114:
	;
	goto L107
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v372
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v372
	v431 = F_slice_from_s(m, l0, int32(1), int32(_a_F_indonesian_ISO_8859_1_stem_4))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L44
	} else {
		goto L116
	}
L116:
	;
	if v431 < int32(0) {
		v601 = v431
		goto L64
	} else {
		goto L117
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(3)
	v437 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v438 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v437 - v438
	v607 = v438
	goto L63
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(3)
	v449 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v450 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v449 - v450
	v607 = v450
	goto L63
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v249
	if v512 == int32(0) {
		goto L134
	} else {
		goto L135
	}
L120:
	;
	v471 = v468
	goto L122
L121:
	;
	v471 = v469
	goto L122
L122:
	;
	goto L124
L123:
	;
	v512 = v507
	goto L119
L124:
	;
	if v468 == v471 {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	v507 = int32(0)
	goto L123
L126:
	;
	v512 = int32(-1)
	goto L119
L127:
	;
	goto L128
L128:
	;
	v483 = int32(1)
	v484 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v486 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v484+v468))))
	if int32(117) < v486 {
		v507 = v483
		goto L123
	} else {
		goto L129
	}
L129:
	;
	v488 = v486 - int32(97)
	if v488 < int32(0) {
		v507 = v483
		goto L123
	} else {
		goto L130
	}
L130:
	;
	v494 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v488)>>(uint(int32(3))%32)))+uint32(_c_F_indonesian_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v494)>>(uint(v488&int32(7))%32))&int32(1) == int32(0) {
		v507 = v483
		goto L123
	} else {
		goto L131
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v468 + int32(1)
	goto L132
L132:
	;
	goto L125
L133:
	;
	v607 = v525
	goto L63
L134:
	;
	v518 = F_slice_from_s(m, l0, int32(1), int32(_a_F_indonesian_ISO_8859_1_stem_5))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L44
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	v522 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v522 {
		v601 = v251
		goto L64
	} else {
		goto L139
	}
L137:
	;
	if v518 < int32(0) {
		v525 = v518
		goto L133
	} else {
		goto L138
	}
L138:
	;
	v601 = v251
	goto L64
L139:
	;
	v525 = v522
	goto L133
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v249
	if v584 == int32(0) {
		goto L155
	} else {
		goto L156
	}
L141:
	;
	v543 = v540
	goto L143
L142:
	;
	v543 = v541
	goto L143
L143:
	;
	goto L145
L144:
	;
	v584 = v579
	goto L140
L145:
	;
	if v540 == v543 {
		goto L147
	} else {
		goto L148
	}
L146:
	;
	v579 = int32(0)
	goto L144
L147:
	;
	v584 = int32(-1)
	goto L140
L148:
	;
	goto L149
L149:
	;
	v555 = int32(1)
	v556 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v558 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556+v540))))
	if int32(117) < v558 {
		v579 = v555
		goto L144
	} else {
		goto L150
	}
L150:
	;
	v560 = v558 - int32(97)
	if v560 < int32(0) {
		v579 = v555
		goto L144
	} else {
		goto L151
	}
L151:
	;
	v566 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v560)>>(uint(int32(3))%32)))+uint32(_c_F_indonesian_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v566)>>(uint(v560&int32(7))%32))&int32(1) == int32(0) {
		v579 = v555
		goto L144
	} else {
		goto L152
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v540 + int32(1)
	goto L153
L153:
	;
	goto L146
L154:
	;
	v601 = v600
	goto L64
L155:
	;
	v588 = int32(1)
	v591 = F_slice_from_s(m, l0, v588, int32(_a_F_indonesian_ISO_8859_1_stem_6))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L44
	} else {
		goto L158
	}
L156:
	;
	goto L157
L157:
	;
	v596 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v596 {
		v601 = int32(1)
		goto L64
	} else {
		goto L160
	}
L158:
	;
	if v591 < int32(0) {
		v600 = v591
		goto L154
	} else {
		goto L159
	}
L159:
	;
	v601 = v588
	goto L64
L160:
	;
	v600 = v596
	goto L154
L161:
	;
	if v607 < int32(0) {
		v664 = v607
		goto L37
	} else {
		goto L164
	}
L162:
	;
	goto L163
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v229
	v637 = F_r_remove_second_order_prefix_1(m, l0)
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L44
	} else {
		goto L178
	}
L164:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v610 < int32(3) {
		goto L62
	} else {
		goto L165
	}
L165:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v613
	v615 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v615
	v617 = F_r_remove_suffix_1(m, l0)
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L44
	} else {
		goto L166
	}
L166:
	;
	if v617 == int32(0) {
		goto L62
	} else {
		goto L167
	}
L167:
	;
	if v617 < int32(0) {
		v664 = v617
		goto L37
	} else {
		goto L168
	}
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v613
	v624 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v624 < int32(3) {
		goto L62
	} else {
		goto L169
	}
L169:
	;
	v627 = F_r_remove_second_order_prefix_1(m, l0)
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L44
	} else {
		goto L170
	}
L170:
	;
	if int32(0) <= v627 {
		goto L62
	} else {
		goto L171
	}
L171:
	;
	if v627 < int32(0) {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v633 = v627
	goto L174
L173:
	;
	v633 = v185
	goto L174
L174:
	;
	if v627 != 0 {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v634 = v633
	goto L177
L176:
	;
	v634 = v185
	goto L177
L177:
	;
	return v634
L178:
	;
	if v637 < int32(0) {
		v664 = v637
		goto L37
	} else {
		goto L179
	}
L179:
	;
	v641 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v641 < int32(3) {
		goto L62
	} else {
		goto L180
	}
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v229
	v645 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v645
	v647 = F_r_remove_suffix_1(m, l0)
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L44
	} else {
		goto L181
	}
L181:
	;
	if base.Ui32(int32(7)) < base.Ui32(int32(base.Ui32(v647)>>(uint(int32(31))%32))-int32(1)) {
		goto L62
	} else {
		goto L182
	}
L182:
	;
	if int32(0) <= v647 {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v658 = int32(1)
	goto L185
L184:
	;
	v658 = v647
	goto L185
L185:
	;
	return v658
}
func F_inetand(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
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
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
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
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = F_pg_detoast_datum_packed(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v16 = F_pg_detoast_datum_packed(m, v15)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int64(0)
		} else {
			v19 = F_palloc0(m, int32(22))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int64(0)
			} else {
				v21 = int32(1)
				v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
				if v23&v21 != 0 {
					v26 = v21
				} else {
					v26 = int32(4)
				}
				v27 = v11 + v26
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
				v29 = int32(1)
				v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
				if v31&v29 != 0 {
					v34 = v29
				} else {
					v34 = int32(4)
				}
				v35 = v16 + v34
				v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
				if v28 == v36 {
					if v28 == int32(2) {
						v42 = int32(3)
					} else {
						v42 = int32(15)
					}
					v43 = int32(2)
					v44 = v35 + v43
					v46 = v27 + v43
					v47 = int32(1)
					v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
					if v49&v47 != 0 {
						v52 = v47
					} else {
						v52 = int32(4)
					}
					v53 = v19 + v52
					v55 = v53 + int32(2)
					v56 = v42
					for {
						v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56+v44))))
						v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56+v46))))
						v70 = v67 & v69
						*(*uint8)(unsafe.Add(mBase, uint32(v56+v55))) = uint8(v70)
						v73 = v56 - int32(1)
						v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73+v44))))
						v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73+v46))))
						v79 = v76 & v78
						*(*uint8)(unsafe.Add(mBase, uint32(v55+v73))) = uint8(v79)
						if v73 != 0 {
							v56 = v56 - int32(2)
							continue
						} else {
							break
						}
						break
					}
					v83 = int32(1)
					v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
					if v85&v83 != 0 {
						v88 = v83
					} else {
						v88 = int32(4)
					}
					v89 = v11 + v88
					v90 = int32(1)
					v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
					if v92&v90 != 0 {
						v95 = v90
					} else {
						v95 = int32(4)
					}
					v96 = v16 + v95
					v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+1)))
					v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+1)))
					if base.Ui32(v98) < base.Ui32(v97) {
						v100 = v89
					} else {
						v100 = v96
					}
					v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+1)))
					*(*uint8)(unsafe.Add(mBase, uint32(v53)+1)) = uint8(v101)
					v103 = int32(1)
					v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
					if v105&v103 != 0 {
						v108 = v103
					} else {
						v108 = int32(4)
					}
					v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11+v108))))
					*(*uint8)(unsafe.Add(mBase, uint32(v53))) = uint8(v110)
					if v110 == int32(2) {
						v116 = int32(40)
					} else {
						v116 = int32(88)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v19))) = v116
					return base.I64_extend_i32_u(v19)
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v123 = m.ExcPending
					if v123 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v126 = m.ExcPending
						if v126 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_inetand_0), int32(0))
							mBase = m.M
							v130 = m.ExcPending
							if v130 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_inetand_1), int32(1826), int32(_a_F_inetand_2))
								mBase = m.M
								v135 = m.ExcPending
								if v135 != 0 {
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
			}
		}
	}
}
func F_inetor(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
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
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
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
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = F_pg_detoast_datum_packed(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v16 = F_pg_detoast_datum_packed(m, v15)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int64(0)
		} else {
			v19 = F_palloc0(m, int32(22))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int64(0)
			} else {
				v21 = int32(1)
				v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
				if v23&v21 != 0 {
					v26 = v21
				} else {
					v26 = int32(4)
				}
				v27 = v11 + v26
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
				v29 = int32(1)
				v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
				if v31&v29 != 0 {
					v34 = v29
				} else {
					v34 = int32(4)
				}
				v35 = v16 + v34
				v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
				if v28 == v36 {
					if v28 == int32(2) {
						v42 = int32(3)
					} else {
						v42 = int32(15)
					}
					v43 = int32(2)
					v44 = v35 + v43
					v46 = v27 + v43
					v47 = int32(1)
					v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
					if v49&v47 != 0 {
						v52 = v47
					} else {
						v52 = int32(4)
					}
					v53 = v19 + v52
					v55 = v53 + int32(2)
					v56 = v42
					for {
						v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56+v44))))
						v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56+v46))))
						v70 = v67 | v69
						*(*uint8)(unsafe.Add(mBase, uint32(v56+v55))) = uint8(v70)
						v73 = v56 - int32(1)
						v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73+v44))))
						v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73+v46))))
						v79 = v76 | v78
						*(*uint8)(unsafe.Add(mBase, uint32(v55+v73))) = uint8(v79)
						if v73 != 0 {
							v56 = v56 - int32(2)
							continue
						} else {
							break
						}
						break
					}
					v83 = int32(1)
					v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
					if v85&v83 != 0 {
						v88 = v83
					} else {
						v88 = int32(4)
					}
					v89 = v11 + v88
					v90 = int32(1)
					v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
					if v92&v90 != 0 {
						v95 = v90
					} else {
						v95 = int32(4)
					}
					v96 = v16 + v95
					v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+1)))
					v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+1)))
					if base.Ui32(v98) < base.Ui32(v97) {
						v100 = v89
					} else {
						v100 = v96
					}
					v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+1)))
					*(*uint8)(unsafe.Add(mBase, uint32(v53)+1)) = uint8(v101)
					v103 = int32(1)
					v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
					if v105&v103 != 0 {
						v108 = v103
					} else {
						v108 = int32(4)
					}
					v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11+v108))))
					*(*uint8)(unsafe.Add(mBase, uint32(v53))) = uint8(v110)
					if v110 == int32(2) {
						v116 = int32(40)
					} else {
						v116 = int32(88)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v19))) = v116
					return base.I64_extend_i32_u(v19)
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v123 = m.ExcPending
					if v123 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v126 = m.ExcPending
						if v126 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_inetor_0), int32(0))
							mBase = m.M
							v130 = m.ExcPending
							if v130 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_inetor_1), int32(1858), int32(_a_F_inetor_2))
								mBase = m.M
								v135 = m.ExcPending
								if v135 != 0 {
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
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
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
	var v286 int32
	_ = v286
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v444 int32
	_ = v444
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
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
	v421 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	if v421 != 0 {
		goto L75
	} else {
		goto L76
	}
L3:
	;
	v390 = v36
	v392 = v34
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
	v413 = v27
	v414 = v35
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
	v266 = F_palloc_mul(m, int32(1), v261)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v266
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v266
	F_infix_2(m, v13+int32(12), int32(0))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L7
	} else {
		goto L51
	}
L51:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v275
	F_infix_2(m, l0, int32(0))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L7
	} else {
		goto L52
	}
L52:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v281 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v283 = v281 - v282
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	if v280 <= v283+v284-v286+int32(4) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v294 = v280
	v295 = v282
	goto L56
L54:
	;
	v321 = v281
	v323 = v286
	goto L55
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v323
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v45
	v330 = F_pg_sprintf(m, v321, int32(_a_F_infix_2_2), v13)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L7
	} else {
		goto L60
	}
L56:
	;
	v304 = v294 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v304
	v306 = F_repalloc(m, v295, v304)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L7
	} else {
		goto L58
	}
L57:
	;
	v321 = v309
	v323 = v314
	goto L55
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v306
	v309 = v306 + v283
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v309
	v311 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	if v311 <= v283+int32(4)+v312-v314 {
		v294 = v311
		v295 = v306
		goto L56
	} else {
		goto L59
	}
L59:
	;
	goto L57
L60:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v333 = F_strlen(m, v332)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v333 + v332
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	F_pfree(m, v336)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
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
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v343 = v341 - v342
	v345 = v343 + int32(3)
	v346 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v346 <= v345 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v349 = v346
	v350 = v342
	goto L66
L64:
	;
	v371 = v341
	goto L65
L65:
	;
	v380 = F_pg_sprintf(m, v371, int32(_a_F_infix_2_1), int32(0))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L7
	} else {
		goto L70
	}
L66:
	;
	v359 = v349 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v359
	v361 = F_repalloc(m, v350, v359)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L7
	} else {
		goto L68
	}
L67:
	;
	v371 = v364
	goto L65
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v361
	v364 = v361 + v343
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v364
	v366 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v366 <= v345 {
		v349 = v366
		v350 = v361
		goto L66
	} else {
		goto L69
	}
L69:
	;
	goto L67
L70:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v383 = F_strlen(m, v382)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v383 + v382
	goto L1
L71:
	;
	v397 = v392 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v397
	v399 = F_repalloc(m, v390, v397)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L7
	} else {
		goto L73
	}
L72:
	;
	v413 = v405
	v414 = v402
	goto L2
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v399
	v402 = v399 + v37
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v402
	v404 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v405 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v405)+9)))
	if v404 <= v39+v406<<(uint(int32(1))%32) {
		v390 = v399
		v392 = v404
		goto L71
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	v423 = v33
	v425 = v414
	v426 = v421
	goto L78
L76:
	;
	v444 = v414
	v451 = v413
	goto L77
L77:
	;
	v452 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v451)+8)))
	if v452&int32(4) != 0 {
		goto L81
	} else {
		goto L82
	}
L78:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v425))) = uint8(v426)
	v433 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v434 = int32(1)
	v435 = v433 + v434
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v435
	v437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v423)+1)))
	if v437 != 0 {
		v423 = v423 + v434
		v425 = v435
		v426 = v437
		goto L78
	} else {
		goto L80
	}
L79:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v444 = v435
	v451 = v440
	goto L77
L80:
	;
	goto L79
L81:
	;
	v455 = int32(37)
	*(*uint8)(unsafe.Add(mBase, uint32(v444))) = uint8(v455)
	v457 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v459 = v457 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v459
	v461 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v462 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v461)+8)))
	v463 = v462
	v464 = v459
	goto L83
L82:
	;
	v463 = v452
	v464 = v444
	goto L83
L83:
	;
	if v463&int32(2) != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v467 = int32(64)
	*(*uint8)(unsafe.Add(mBase, uint32(v464))) = uint8(v467)
	v469 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v471 = v469 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v471
	v473 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v474 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v473)+8)))
	v475 = v471
	v476 = v474
	goto L86
L85:
	;
	v475 = v464
	v476 = v463
	goto L86
L86:
	;
	if v476&int32(1) != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v479 = int32(42)
	*(*uint8)(unsafe.Add(mBase, uint32(v475))) = uint8(v479)
	v481 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v483 = v481 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v483
	v485 = v483
	goto L89
L88:
	;
	v485 = v475
	goto L89
L89:
	;
	v486 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v485))) = uint8(v486)
	v488 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v488 + int32(12)
	goto L1
}
func F_initial_cost_hashjoin(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
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
	var v32 float64
	_ = v32
	var v35 int32
	_ = v35
	var v36 float64
	_ = v36
	var v37 float64
	_ = v37
	var v39 float64
	_ = v39
	var v41 float64
	_ = v41
	var v44 float64
	_ = v44
	var v46 float64
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int64
	_ = v49
	var v50 int32
	_ = v50
	var v51 float64
	_ = v51
	var v53 int32
	_ = v53
	var v59 float64
	_ = v59
	var v63 float64
	_ = v63
	var v66 float64
	_ = v66
	var v68 float64
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v89 int32
	_ = v89
	var v94 float64
	_ = v94
	var v96 int32
	_ = v96
	var v100 float64
	_ = v100
	var v101 float64
	_ = v101
	var v104 float64
	_ = v104
	var v105 int32
	_ = v105
	var v109 float64
	_ = v109
	var v110 float64
	_ = v110
	var v112 int32
	_ = v112
	var v117 float64
	_ = v117
	var v118 float64
	_ = v118
	var v121 float64
	_ = v121
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v143 float64
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 float64
	_ = v153
	var v155 float64
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v166 int32
	_ = v166
	var v176 float64
	_ = v176
	var v178 int32
	_ = v178
	var v182 float64
	_ = v182
	var v183 float64
	_ = v183
	var v186 float64
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 float64
	_ = v214
	var v216 float64
	_ = v216
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v227 int32
	_ = v227
	var v235 float64
	_ = v235
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v270 int32
	_ = v270
	var v276 float64
	_ = v276
	var v278 float64
	_ = v278
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v289 int32
	_ = v289
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v305 int32
	_ = v305
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v335 int32
	_ = v335
	var v343 int32
	_ = v343
	var v354 int32
	_ = v354
	var v360 int32
	_ = v360
	var v368 int32
	_ = v368
	var v372 float64
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v383 float64
	_ = v383
	var v385 float64
	_ = v385
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v399 float64
	_ = v399
	var v405 float64
	_ = v405
	var v407 float64
	_ = v407
	var v417 int64
	_ = v417
	var v423 int32
	_ = v423
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	v21 = *(*float64)(unsafe.Add(mBase, uint32(l3)+32))
	v22 = *(*float64)(unsafe.Add(mBase, uint32(l2)+32))
	v24 = *(*float64)(unsafe.Add(mBase, _c_F_initial_cost_hashjoin[0]))
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
	v32 = *(*float64)(unsafe.Add(mBase, _c_F_initial_cost_hashjoin[1]))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	v36 = *(*float64)(unsafe.Add(mBase, uint32(l2)+56))
	v37 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
	v39 = float64(0)
	v41 = base.F64_add(base.F64_mul(v29, v22), base.F64_add(base.F64_sub(v36, v37), v39))
	v44 = *(*float64)(unsafe.Add(mBase, uint32(l3)+56))
	v46 = base.F64_add(base.F64_mul(base.F64_add(v29, v32), v21), base.F64_add(base.F64_add(v37, v39), v44))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l3)+40))
	v49 = *(*int64)(unsafe.Add(mBase, uint32(l4)+40))
	if l5 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	v51 = base.F64_convert_i32_s(v50)
	v53 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_initial_cost_hashjoin[2])))
	if v53 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v68 = v21
	goto L6
L6:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+32))
	v78 = v19 + int32(4)
	v89 = (v71 + int32(7)) & int32(-8)
	v94 = *(*float64)(unsafe.Add(mBase, _c_F_initial_cost_hashjoin[3]))
	v96 = *(*int32)(unsafe.Add(mBase, _c_F_initial_cost_hashjoin[4]))
	v100 = base.F64_mul(base.F64_mul(v94, base.F64_convert_i32_s(v96)), float64(1024))
	v101 = float64(4.294967295e+09)
	if base.F64_lt(v100, v101) != 0 {
		goto L14
	} else {
		goto L15
	}
L7:
	;
	v59 = base.F64_add(base.F64_mul(v51, float64(-0.3)), float64(1))
	if base.F64_gt(v59, float64(0)) != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v66 = v51
	goto L9
L9:
	;
	v68 = base.F64_mul(v21, v66)
	goto L6
L10:
	;
	v63 = v59
	goto L12
L11:
	;
	v63 = math.Float64frombits(uint64(0x8000000000000000))
	goto L12
L12:
	;
	v66 = base.F64_add(v63, v51)
	goto L9
L13:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	if int32(2) <= v368 {
		goto L99
	} else {
		goto L100
	}
L14:
	;
	v104 = v100
	goto L16
L15:
	;
	v104 = v101
	goto L16
L16:
	;
	v105 = base.I32_trunc_sat_f64_u(v104)
	if base.F64_le(v68, float64(0)) != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v109 = float64(1000)
	goto L19
L18:
	;
	v109 = v68
	goto L19
L19:
	;
	v110 = base.F64_mul(v109, base.F64_convert_i32_s(v89+int32(24)))
	v112 = v89 + int32(68)
	if l5 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v117 = base.F64_mul(base.F64_convert_i32_s(v35+int32(1)), base.F64_convert_i32_u(v105))
	v118 = float64(4.294967295e+09)
	if base.F64_lt(v117, v118) != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v124 = v105
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v124
	goto L26
L23:
	;
	v121 = v117
	goto L25
L24:
	;
	v121 = v118
	goto L25
L25:
	;
	v124 = base.I32_trunc_sat_f64_u(v121)
	goto L22
L26:
	;
	v126 = base.I32_div_u_s(v124, v112)
	v127 = int32(50)
	v128 = base.I32_div_u_s(v126, v127)
	if base.Ui32(v127) <= base.Ui32(v126) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v78))) = v128
	v138 = int32(1)
	v143 = base.F64_ceil(v109)
	v145 = int32(268435455)
	v147 = int32(base.Ui32(v134) >> (uint(int32(2)) % 32))
	if base.Ui32(v145) <= base.Ui32(v147) {
		goto L33
	} else {
		goto L34
	}
L29:
	;
	v133 = v128 * v112
	goto L31
L30:
	;
	v133 = int32(0)
	goto L31
L31:
	;
	v134 = v124 - v133
	goto L28
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19+int32(12)))) = v360
	*(*int32)(unsafe.Add(mBase, uint32(v19+int32(8)))) = v354
	goto L13
L33:
	;
	v150 = v145
	goto L35
L34:
	;
	v150 = v147
	goto L35
L35:
	;
	v152 = int32(base.Ui32(int32(-2147483648)) >> (uint(base.I32_clz(v150)) % 32))
	v153 = base.F64_convert_i32_u(v152)
	if base.F64_gt(v153, v143) != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v155 = v143
	goto L38
L37:
	;
	v155 = v153
	goto L38
L38:
	;
	v156 = base.I32_trunc_sat_f64_s(v155)
	if v156 <= int32(1024) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v159 = int32(1024)
	goto L41
L40:
	;
	v159 = v156
	goto L41
L41:
	;
	if v159&(v159-int32(1)) != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v166 = v138 << (uint(int32(32)-base.I32_clz(v159)) % 32)
	goto L44
L43:
	;
	v166 = v159
	goto L44
L44:
	;
	if base.F64_lt(base.F64_convert_i32_u(v134), base.F64_add(v110, base.F64_convert_i32_u(v166<<(uint(int32(2))%32)))) == int32(0) {
		v354 = v138
		v360 = v166
		goto L32
	} else {
		goto L45
	}
L45:
	;
	if l5 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v176 = *(*float64)(unsafe.Add(mBase, _c_F_initial_cost_hashjoin[3]))
	v178 = *(*int32)(unsafe.Add(mBase, _c_F_initial_cost_hashjoin[4]))
	v182 = base.F64_mul(base.F64_mul(v176, base.F64_convert_i32_s(v178)), float64(1024))
	v183 = float64(4.294967295e+09)
	if base.F64_lt(v182, v183) != 0 {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	v235 = v153
	v236 = v134
	v240 = v152
	goto L48
L48:
	;
	v243 = v89 + int32(28)
	if base.Ui32(v243) < base.Ui32(v236) {
		goto L71
	} else {
		goto L72
	}
L49:
	;
	v186 = v182
	goto L51
L50:
	;
	v186 = v183
	goto L51
L51:
	;
	v187 = base.I32_trunc_sat_f64_u(v186)
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v187
	goto L52
L52:
	;
	v189 = base.I32_div_u_s(v187, v112)
	v190 = int32(50)
	v191 = base.I32_div_u_s(v189, v190)
	if base.Ui32(v190) <= base.Ui32(v189) {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v78))) = v191
	v206 = int32(268435455)
	v208 = int32(base.Ui32(v197) >> (uint(int32(2)) % 32))
	if base.Ui32(v206) <= base.Ui32(v208) {
		goto L58
	} else {
		goto L59
	}
L55:
	;
	v196 = v191 * v112
	goto L57
L56:
	;
	v196 = int32(0)
	goto L57
L57:
	;
	v197 = v187 - v196
	goto L54
L58:
	;
	v211 = v206
	goto L60
L59:
	;
	v211 = v208
	goto L60
L60:
	;
	v213 = int32(base.Ui32(int32(-2147483648)) >> (uint(base.I32_clz(v211)) % 32))
	v214 = base.F64_convert_i32_u(v213)
	if base.F64_gt(v214, v143) != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v216 = v143
	goto L63
L62:
	;
	v216 = v214
	goto L63
L63:
	;
	v217 = base.I32_trunc_sat_f64_s(v216)
	if v217 <= int32(1024) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v220 = int32(1024)
	goto L66
L65:
	;
	v220 = v217
	goto L66
L66:
	;
	if v220&(v220-int32(1)) != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v227 = int32(1) << (uint(int32(32)-base.I32_clz(v220)) % 32)
	goto L69
L68:
	;
	v227 = v220
	goto L69
L69:
	;
	if base.F64_lt(base.F64_convert_i32_u(v197), base.F64_add(v110, base.F64_convert_i32_u(v227<<(uint(int32(2))%32)))) == int32(0) {
		v354 = v138
		v360 = v227
		goto L32
	} else {
		goto L70
	}
L70:
	;
	v235 = v214
	v236 = v197
	v240 = v213
	goto L48
L71:
	;
	v245 = int32(1)
	v247 = base.I32_div_u_s(v236, v243)
	if v247&(v247-v245) != 0 {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	v258 = int32(1)
	goto L73
L73:
	;
	v259 = int32(1)
	v260 = int32(32)
	if v258&(v258-v259) != 0 {
		goto L81
	} else {
		goto L82
	}
L74:
	;
	v254 = v245 << (uint(int32(32)-base.I32_clz(v247)) % 32)
	goto L76
L75:
	;
	v254 = v247
	goto L76
L76:
	;
	if base.Ui32(v254) < base.Ui32(v240) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v256 = v254
	goto L79
L78:
	;
	v256 = v240
	goto L79
L79:
	;
	v258 = v256
	goto L73
L80:
	;
	v354 = v335
	v360 = v343
	goto L32
L81:
	;
	v270 = v259 << (uint(v260-base.I32_clz(v258)) % 32)
	goto L83
L82:
	;
	v270 = v258
	goto L83
L83:
	;
	v276 = base.F64_ceil(base.F64_div(v110, base.F64_convert_i32_u(v236-v270<<(uint(int32(2))%32))))
	if base.F64_gt(v235, v276) != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v278 = v276
	goto L86
L85:
	;
	v278 = v235
	goto L86
L86:
	;
	v279 = base.I32_trunc_sat_f64_s(v278)
	if v279 <= int32(2) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v282 = int32(2)
	goto L89
L88:
	;
	v282 = v279
	goto L89
L89:
	;
	if v282&(v282-int32(1)) != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v289 = v259 << (uint(v260-base.I32_clz(v282)) % 32)
	goto L92
L91:
	;
	v289 = v282
	goto L92
L92:
	;
	if base.B2i32(v289 < int32(2))|base.B2i32(base.Ui32(int32(134217727)) < base.Ui32(v270)) != 0 {
		v335 = v289
		v343 = v270
		goto L80
	} else {
		goto L93
	}
L93:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v297 = v289
	v299 = v295
	v305 = v270
	goto L94
L94:
	;
	if base.B2i32(v299 < int32(0))|base.B2i32(base.Ui32(v297) < base.Ui32(int32(base.Ui32(v299)>>(uint(int32(13))%32)))) != 0 {
		v335 = v297
		v343 = v305
		goto L80
	} else {
		goto L96
	}
L95:
	;
	v354 = v327
	v360 = v329
	goto L32
L96:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	v319 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v78))) = v318 << (uint(v319) % 32)
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v324 = v322 << (uint(v319) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v324
	v327 = int32(base.Ui32(v297) >> (uint(v319) % 32))
	v329 = v305 << (uint(v319) % 32)
	if base.Ui32(v297) < base.Ui32(int32(4)) {
		v354 = v327
		v360 = v329
		goto L32
	} else {
		goto L97
	}
L97:
	;
	if base.Ui32(v305) < base.Ui32(int32(67108864)) {
		v297 = v327
		v299 = v324
		v305 = v329
		goto L94
	} else {
		goto L98
	}
L98:
	;
	goto L95
L99:
	;
	v372 = *(*float64)(unsafe.Add(mBase, _c_F_initial_cost_hashjoin[5]))
	v373 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v373)+32))
	v375 = int32(7)
	v377 = int32(-8)
	v379 = int32(24)
	v383 = float64(0.0001220703125)
	v385 = base.F64_ceil(base.F64_mul(base.F64_mul(v22, base.F64_convert_i32_u((v374+v375)&v377+v379)), v383))
	v387 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v387)+32))
	v399 = base.F64_ceil(base.F64_mul(base.F64_mul(v21, base.F64_convert_i32_u((v388+v375)&v377+v379)), v383))
	v405 = base.F64_add(base.F64_mul(v372, v399), v46)
	v407 = base.F64_add(base.F64_mul(v372, base.F64_add(base.F64_add(v385, v385), v399)), v41)
	goto L101
L100:
	;
	v405 = v46
	v407 = v41
	goto L101
L101:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+24)) = v407
	*(*float64)(unsafe.Add(mBase, uint32(l0)+8)) = v405
	*(*float64)(unsafe.Add(mBase, uint32(l0)+16)) = base.F64_add(v407, v405)
	if v35 != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v417 = int64(-2049)
	goto L104
L103:
	;
	v417 = int64(-264193)
	goto L104
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v47 + v48 + base.B2i32(v417|v49 != int64(-1))
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+88)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v368
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v423
	m.G0 = v19 + int32(16)
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
	var v52 int32
	_ = v52
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int64
	_ = v129
	v4 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v7 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v13
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
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
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
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v36
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = int32(0)
	v47 = v40
	goto L7
L17:
	;
	goto L16
L18:
	;
	v84 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v84
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)) = uint8(v84)
	v88 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v88
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+76)) = uint16(v84)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v88
	*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = int64(4294967295)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+108)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+92)) = int64(-4294967295)
	if l1 == v84 {
		goto L33
	} else {
		goto L34
	}
L19:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+12)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v48&int32(-129) | v52<<(uint(int32(7))%32)&int32(128)
	if l2 != 0 {
		goto L18
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_initscan[0])))
	v63 = v47 & v62
	if l2 != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = int32(-1)
	goto L18
L23:
	;
	if v63 != 0 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	if v63 != 0 {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v48 | int32(128)
	goto L18
L27:
	;
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v48 & int32(-129)
	goto L18
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v48 | int32(128)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v75 = F_ss_get_location(m, v73, v74)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L5
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v48 & int32(-129)
	goto L18
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v75
	goto L18
L33:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v112&int32(1) == int32(0) {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v102 <= int32(0) {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v106 = v102 * int32(56)
	if v106 == int32(0) {
		goto L33
	} else {
		goto L36
	}
L36:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	base.MemoryCopy(m, v109, l1, v106)
	goto L33
L37:
	;
	return
L38:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+272))
	if v118 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+268)))
	if v121 != int32(1) {
		goto L37
	} else {
		goto L42
	}
L40:
	;
	v128 = v118
	goto L41
L41:
	;
	v129 = *(*int64)(unsafe.Add(mBase, uint32(v128)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v128)+16)) = v129 + int64(1)
	goto L37
L42:
	;
	F_pgstat_assoc_relation(m, v117)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L5
	} else {
		goto L43
	}
L43:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+272))
	v128 = v127
	goto L41
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
						F_errfinish(m, int32(_a_F_inittapes_1), int32(1780), int32(_a_F_inittapes_2))
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
func F_inline_function_in_from(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
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
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v88 int64
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v104 int64
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v161 int64
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
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
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
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
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	v3 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(48)
	m.G0 = v18
	F_check_stack_depth(m)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+72)))
	if v24 != 0 {
		v346 = v3
		goto L4
	} else {
		goto L5
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L1
	} else {
		goto L107
	}
L4:
	;
	m.G0 = v18 + int32(48)
	return v346
L5:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	if v25 == int32(0) {
		v346 = v3
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v28 != int32(1) {
		v346 = v3
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	if v34 != int32(15) {
		v346 = v3
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v33)+28))
	v40 = F_contain_volatile_functions_walker(m, v38, int32(0))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	if v40 != 0 {
		v346 = v3
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v33)+28))
	if v42 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	if base.Ui32(v43-int32(22)) < base.Ui32(int32(3)) {
		v346 = v3
		goto L4
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_inline_function_in_from[0]))
	v56 = F_object_aclcheck(m, int32(1255), v37, v54, int64(128))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L17
	}
L14:
	;
	v50 = F_expression_tree_walker_impl(m, v42, int32(905), int32(0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	if v50 != 0 {
		v346 = v3
		goto L4
	} else {
		goto L16
	}
L16:
	;
	goto L13
L17:
	;
	if v56 != 0 {
		v346 = v3
		goto L4
	} else {
		goto L18
	}
L18:
	;
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_inline_function_in_from[1]))
	if v59 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v60 = m.T0[v59].(func(*base.Module, int32) int32)(m, v37)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v64 = F_SearchSysCache1(m, int32(47), base.I64_extend_i32_u(v37))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	if v60 != 0 {
		v346 = v3
		goto L4
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	if v64 == int32(0) {
		goto L3
	} else {
		goto L25
	}
L25:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v64)+16))
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+22)))
	v72 = F_heap_attisnull(m, v64, int32(29), int32(0))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	if v72 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_inline_function_in_from[2]))
	v80 = F_AllocSetContextCreateInternal(m, v75, int32(_a_F_inline_function_in_from_3), int32(0), int32(_a_F_inline_function_in_from_4), int32(_a_F_inline_function_in_from_5))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	v334 = v3
	goto L29
L29:
	;
	F_ReleaseCatCache(m, v64)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L1
	} else {
		goto L106
	}
L30:
	;
	v82 = int32(_a_F_inline_function_in_from_6)
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_inline_function_in_from[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_inline_function_in_from[2])) = v80
	v88 = F_SysCacheGetAttrNotNull(m, int32(47), v64, int32(26))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v91 = F_text_to_cstring(m, base.I32_wrap_i64(v88))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v93 = v69 + v68
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+92))
	if v94 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = int32(465)
	v104 = F_OidFunctionCall1Coll(m, v94, int32(0), base.I64_extend_i32_u(v18+int32(12)))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L36
	}
L34:
	;
	v107 = v3
	goto L35
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v91
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = int32(921)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v93 + int32(4)
	v114 = int32(_a_F_inline_function_in_from_7)
	v115 = *(*int32)(unsafe.Add(mBase, _c_F_inline_function_in_from[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_inline_function_in_from[3])) = v18 + int32(28)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v18 + int32(40)
	if v107 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	v107 = base.I32_wrap_i64(v104)
	goto L35
L37:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_inline_function_in_from[2])) = v83
	F_MemoryContextDelete(m, v80)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L1
	} else {
		goto L105
	}
L38:
	;
	v126 = int32(0)
	v127 = m.G0
	v129 = v127 - int32(32)
	m.G0 = v129
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+12)))
	if v131 != int32(1) {
		v279 = v126
		goto L41
	} else {
		goto L42
	}
L39:
	;
	v289 = v107
	goto L40
L40:
	;
	v294 = int32(*(*int16)(unsafe.Add(mBase, uint32(v93)+104)))
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v33)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v295
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v294
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = int32(1)
	v304 = F_query_tree_mutator_impl(m, v289, int32(923), v18+int32(12), int32(0))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L1
	} else {
		goto L99
	}
L41:
	;
	m.G0 = v129 + int32(32)
	if v279 == int32(0) {
		goto L37
	} else {
		goto L98
	}
L42:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v93)+76))
	if v134 != int32(14) {
		v279 = v126
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93)+96)))
	if v137 != int32(102) {
		v279 = v126
		goto L41
	} else {
		goto L44
	}
L44:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93)+99)))
	if v140 != 0 {
		v279 = v126
		goto L41
	} else {
		goto L45
	}
L45:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93)+101)))
	if v141 == int32(118) {
		v279 = v126
		goto L41
	} else {
		goto L46
	}
L46:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v93)+108))
	if v144 == int32(2278) {
		v279 = v126
		goto L41
	} else {
		goto L47
	}
L47:
	;
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93)+97)))
	if v147 != 0 {
		v279 = v126
		goto L41
	} else {
		goto L48
	}
L48:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93)+100)))
	if v148 != int32(1) {
		v279 = v126
		goto L41
	} else {
		goto L49
	}
L49:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v33)+28))
	if v151 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)+4))
	v154 = v152
	goto L52
L51:
	;
	v154 = int32(0)
	goto L52
L52:
	;
	v155 = int32(*(*int16)(unsafe.Add(mBase, uint32(v93)+104)))
	if v154 != v155 {
		v279 = v126
		goto L41
	} else {
		goto L53
	}
L53:
	;
	v161 = F_SysCacheGetAttr(m, int32(47), v64, int32(28), v129+int32(31))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+31)))
	if v163 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v228)+12))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	if v231 != 0 {
		goto L84
	} else {
		goto L85
	}
L56:
	;
	v167 = F_text_to_cstring(m, base.I32_wrap_i64(v161))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L60
	}
L57:
	;
	goto L58
L58:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v33)+24))
	v206 = F_prepare_sql_fn_parse_info(m, v64, v33, v205)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L76
	}
L59:
	;
	if v183 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L60:
	;
	v169 = F_stringToNode(m, v167)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v169)))
	if v171 == int32(1) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v169)+12))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
	v183 = v175
	goto L59
L63:
	;
	goto L64
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v129)+12)) = v169
	*(*int32)(unsafe.Add(mBase, uint32(v129)+20)) = v169
	v181 = F_list_make1_impl(m, int32(1), v129+int32(12))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v183 = v181
	goto L59
L66:
	;
	v279 = int32(0)
	goto L41
L67:
	;
	goto L68
L68:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v183)+4))
	if v187 != int32(1) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v279 = int32(0)
	goto L41
L70:
	;
	goto L71
L71:
	;
	v191 = int32(0)
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v183)+12))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)))
	F_AcquireRewriteLocks(m, v193, int32(1), v191)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v198 = F_pg_rewrite_query(m, v193)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	if v198 == int32(0) {
		v279 = v191
		goto L41
	} else {
		goto L74
	}
L74:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v198)+4))
	if v202 == int32(1) {
		v228 = v198
		goto L55
	} else {
		goto L75
	}
L75:
	;
	v279 = v191
	goto L41
L76:
	;
	v208 = F_pg_parse_query(m, v91)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	if v208 == int32(0) {
		v279 = v126
		goto L41
	} else {
		goto L78
	}
L78:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v208)+4))
	if v212 != int32(1) {
		v279 = v126
		goto L41
	} else {
		goto L79
	}
L79:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v208)+12))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v215)))
	v219 = F_pg_analyze_and_rewrite_withcb(m, v216, v91, int32(509), v206, int32(0))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	if v219 == int32(0) {
		v279 = v126
		goto L41
	} else {
		goto L81
	}
L81:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v219)+4))
	if v223 != int32(1) {
		v279 = v126
		goto L41
	} else {
		goto L82
	}
L82:
	;
	v228 = v219
	goto L55
L83:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v230)))
	if v249 != int32(67) {
		goto L89
	} else {
		goto L90
	}
L84:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
	v235 = F_BuildDescFromLists(m, v231, v232, v233, v234)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	v242 = F_get_expr_result_type(m, v33, int32(0), v129+int32(24))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L1
	} else {
		goto L88
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v129)+24)) = v235
	v248 = int32(0)
	goto L83
L88:
	;
	v248 = base.B2i32(base.Ui32(v242-int32(4)) < base.Ui32(int32(-3)))
	goto L83
L89:
	;
	v279 = int32(0)
	goto L41
L90:
	;
	goto L91
L91:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v230)+4))
	if v253 != int32(1) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v279 = int32(0)
	goto L41
L93:
	;
	goto L94
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v129)+8)) = v228
	*(*int32)(unsafe.Add(mBase, uint32(v129)+16)) = v228
	v263 = F_list_make1_impl(m, int32(1), v129+int32(8))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v129)+24))
	v267 = int32(*(*int8)(unsafe.Add(mBase, uint32(v93)+96)))
	v269 = F_check_sql_fn_retval(m, v263, v265, v266, v267, int32(1))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	if (v269|v248)&int32(1) == int32(0) {
		v279 = int32(0)
		goto L41
	} else {
		goto L97
	}
L97:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v228)+12))
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v276)))
	v279 = v277
	goto L41
L98:
	;
	v289 = v279
	goto L40
L99:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_inline_function_in_from[2])) = v83
	v308 = F_copyObjectImpl(m, v304)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	F_MemoryContextDelete(m, v80)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v18)+28))
	*(*int32)(unsafe.Add(mBase, _c_F_inline_function_in_from[3])) = v313
	F_ReleaseCatCache(m, v64)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	F_record_plan_function_dependency(m, l0, v37)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v308)+44)))
	if v319 != int32(1) {
		v346 = v308
		goto L4
	} else {
		goto L104
	}
L104:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v323 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v322)+93)) = uint8(v323)
	v346 = v308
	goto L4
L105:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v18)+28))
	*(*int32)(unsafe.Add(mBase, _c_F_inline_function_in_from[3])) = v331
	v334 = int32(0)
	goto L29
L106:
	;
	v346 = v334
	goto L4
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v37
	F_errmsg_internal(m, int32(_a_F_inline_function_in_from_0), v18)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	F_errfinish(m, int32(_a_F_inline_function_in_from_1), int32(_a_F_inline_function_in_from_2), int32(_a_F_inline_function_in_from_3))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	base.Wasm_trap_unreachable()
	for {
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
func F_int24eq(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v3 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+24)))
	return base.I64_extend_i32_u(base.B2i32(v2 == v3))
}
func F_int28eq(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v3 int64
	_ = v3
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v3 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+24)))
	return base.I64_extend_i32_u(base.B2i32(v2 == v3))
}
func F_int2eqfast(m *base.Module, l0 int64, l1 int64) int32 {
	var v4 int32
	_ = v4
	v4 = int32(_a_F_int2eqfast_0)
	return base.B2i32(base.I32_wrap_i64(l0)&v4 == base.I32_wrap_i64(l1)&v4)
}
func F_int2ge(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+24)))
	v3 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+40)))
	return base.I64_extend_i32_u(base.B2i32(v3 <= v2))
}
func F_int2out(m *base.Module, l0 int32) int64 {
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
	v2 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+24)))
	v4 = F_palloc(m, int32(7))
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
func F_int2pl(m *base.Module, l0 int32) int64 {
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
	v4 = v2 + v3
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
				F_errmsg(m, int32(_a_F_int2pl_0), int32(0))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_int2pl_1), int32(944), int32(_a_F_int2pl_2))
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
func F_int42eq(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+40)))
	return base.I64_extend_i32_u(base.B2i32(v2 == v3))
}
func F_int48eq(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v3 int64
	_ = v3
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v3 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+24)))
	return base.I64_extend_i32_u(base.B2i32(v2 == v3))
}
func F_int48ge(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v3 int64
	_ = v3
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v3 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+24)))
	return base.I64_extend_i32_u(base.B2i32(v2 <= v3))
}
func F_int4mi(m *base.Module, l0 int32) int64 {
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
				F_errmsg(m, int32(_a_F_int4mi_0), int32(0))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_int4mi_1), int32(843), int32(_a_F_int4mi_2))
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
func F_int4mul(m *base.Module, l0 int32) int64 {
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
				F_errmsg(m, int32(_a_F_int4mul_0), int32(0))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_int4mul_1), int32(857), int32(_a_F_int4mul_2))
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
func F_int4um(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v26 int64
	_ = v26
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if base.I32_wrap_i64(v3) == int32(-2147483648) {
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
				F_errmsg(m, int32(_a_F_int4um_0), int32(0))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_int4um_1), int32(807), int32(_a_F_int4um_2))
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
		v26 = int64(32)
		return (int64(0) - v3<<(uint(v26)%64)) >> (uint(v26) % 64)
	}
}
func F_int82(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v8 int64
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v29 int64
	_ = v29
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if base.Ui64(int64(-65537)) < base.Ui64(v3-int64(32768)) {
		v29 = v3
		return v29
	} else {
		v8 = int64(0)
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v10 = F_errsave_start(m, v9)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int64(0)
		} else {
			if v10 == int32(0) {
				v29 = v8
				return v29
			} else {
				F_errcode(m, int32(50331778))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, int32(_a_F_int82_0), int32(0))
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int64(0)
					} else {
						F_errsave_finish(m, v9, int32(_a_F_int82_1), int32(1317), int32(_a_F_int82_2))
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return int64(0)
						} else {
							v29 = v8
							return v29
						}
					}
				}
			}
		}
	}
}
func F_int82mi(m *base.Module, l0 int32) int64 {
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
	v4 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+40)))
	v7 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = v7 - v4
	if base.B2i32(int64(0) < v4) != base.B2i32(v8 < v7) {
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
				F_errmsg(m, int32(_a_F_int82mi_0), int32(0))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_int82mi_1), int32(1094), int32(_a_F_int82mi_2))
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
func F_int84ge(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v3 int64
	_ = v3
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+40)))
	return base.I64_extend_i32_u(base.B2i32(v3 <= v2))
}
func F_int84mul(m *base.Module, l0 int32) int64 {
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v6 = Fn14245(m, l0, int32(_a_F_int84mul_0), int32(966), int32(_a_F_int84mul_1), int32(_a_F_int84mul_2))
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return v6
	}
}
func F_int84pl(m *base.Module, l0 int32) int64 {
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
	v4 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+40)))
	v7 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
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
				F_errmsg(m, int32(_a_F_int84pl_0), int32(0))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_int84pl_1), int32(938), int32(_a_F_int84pl_2))
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
func F_int8gcd(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v6 int64
	_ = v6
	var v7 int64
	_ = v7
	var v8 int64
	_ = v8
	var v12 int64
	_ = v12
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v17 int64
	_ = v17
	var v31 int64
	_ = v31
	var __phi31 int64
	_ = __phi31
	var v32 int64
	_ = v32
	var __phi32 int64
	_ = __phi32
	var v34 int64
	_ = v34
	var v39 int64
	_ = v39
	var v42 int64
	_ = v42
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v7 = int64(63)
	v8 = v5 >> (uint(v7) % 64)
	v12 = v6 >> (uint(v7) % 64)
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
	if v17 == int64(-9223372036854775807-1) {
		if v16&int64(9223372036854775807) == int64(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(50331778))
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, int32(_a_F_int8gcd_0), int32(0))
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_int8gcd_1), int32(645), int32(_a_F_int8gcd_2))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
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
			if v16 != int64(-1) {
				__phi31 = v16
				__phi32 = v17
				v31 = __phi31
				v32 = __phi32
				for {
					v34 = base.I64_rem_s(v32, v31)
					if v34 != int64(0) {
						__phi31 = v34
						__phi32 = v31
						v31 = __phi31
						v32 = __phi32
						continue
					} else {
						break
					}
					break
				}
				v39 = v31
				v42 = v39 >> (uint(int64(63)) % 64)
				return v39 ^ v42 - v42
			} else {
				return int64(1)
			}
		}
	} else {
		if v16 != int64(0) {
			__phi31 = v16
			__phi32 = v17
			v31 = __phi31
			v32 = __phi32
			for {
				v34 = base.I64_rem_s(v32, v31)
				if v34 != int64(0) {
					__phi31 = v34
					__phi32 = v31
					v31 = __phi31
					v32 = __phi32
					continue
				} else {
					break
				}
				break
			}
			v39 = v31
		} else {
			v39 = v17
		}
		v42 = v39 >> (uint(int64(63)) % 64)
		return v39 ^ v42 - v42
	}
}
func F_int8inc_support(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v7 int64
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v39 int32
	_ = v39
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
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
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v99 int64
	_ = v99
	v5 = int64(0)
	v7 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = base.I32_wrap_i64(v7)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	switch v9 - int32(464) {
	case 0:
		v54 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
		v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
		if v55 != int32(2147) {
			v99 = v5
			return v99
		} else {
			v58 = *(*int32)(unsafe.Add(mBase, uint32(v54)+32))
			if v58 == int32(0) {
				v99 = v5
				return v99
			} else {
				v61 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
				if v61 != int32(1) {
					v99 = v5
					return v99
				} else {
					v64 = *(*int32)(unsafe.Add(mBase, uint32(v54)+40))
					if v64 != 0 {
						v99 = v5
						return v99
					} else {
						v65 = *(*int32)(unsafe.Add(mBase, uint32(v54)+36))
						if v65 != 0 {
							v99 = v5
							return v99
						} else {
							v66 = *(*int32)(unsafe.Add(mBase, uint32(v54)+52))
							if v66 != 0 {
								v99 = v5
								return v99
							} else {
								v67 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
								v68 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
								v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
								v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
								v72 = F_expr_is_nonnullable(m, v67, v70, int32(1))
								mBase = m.M
								v75 = m.ExcPending
								if v75 != 0 {
									return int64(0)
								} else {
									if v72 == int32(0) {
										v99 = v5
										return v99
									} else {
										v79 = F_palloc0(m, int32(72))
										mBase = m.M
										v80 = m.ExcPending
										if v80 != 0 {
											return int64(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v79))) = int32(9)
											base.MemoryCopy(m, v79, v54, int32(68))
											*(*int32)(unsafe.Add(mBase, uint32(v79)+68)) = int32(-1)
											v87 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v79)+48)) = uint8(v87)
											v89 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v79)+32)) = v89
											*(*int32)(unsafe.Add(mBase, uint32(v79)+24)) = v89
											*(*int32)(unsafe.Add(mBase, uint32(v79)+4)) = int32(2803)
											v99 = base.I64_extend_i32_u(v79)
											return v99
										}
									}
								}
							}
						}
					}
				}
			}
		}
	default:
		v99 = v5
		return v99
	case 6:
		v13 = v7 & int64(4294967295)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
		if v15&int32(_a_F_int8inc_support_0) == int32(0) {
			if v15&int32(2) != 0 {
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
				if v39 == int32(0) {
					v51 = int32(3)
				} else {
					v51 = int32(base.Ui32(v15)>>(uint(int32(5))%32))&int32(1) | int32(base.Ui32(v15)>>(uint(int32(7))%32))&int32(2)
				}
			} else {
				v51 = int32(base.Ui32(v15)>>(uint(int32(5))%32))&int32(1) | int32(base.Ui32(v15)>>(uint(int32(7))%32))&int32(2)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v51
			return v13
		} else {
			if v15&int32(_a_F_int8inc_support_1) == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = int32(0)
				return v13
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
				if v25 != int32(2803) {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = int32(0)
					return v13
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
					if v28 == int32(0) {
						if v15&int32(2) != 0 {
							v39 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
							if v39 == int32(0) {
								v51 = int32(3)
							} else {
								v51 = int32(base.Ui32(v15)>>(uint(int32(5))%32))&int32(1) | int32(base.Ui32(v15)>>(uint(int32(7))%32))&int32(2)
							}
						} else {
							v51 = int32(base.Ui32(v15)>>(uint(int32(5))%32))&int32(1) | int32(base.Ui32(v15)>>(uint(int32(7))%32))&int32(2)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v51
						return v13
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = int32(0)
						return v13
					}
				}
			}
		}
	}
}
func F_int8mul(m *base.Module, l0 int32) int64 {
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v6 = Fn14246(m, l0, int32(_a_F_int8mul_0), int32(506), int32(_a_F_int8mul_1), int32(_a_F_int8mul_2))
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return v6
	}
}
func F_int8out(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
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
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v8 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if int64(0) <= v8 {
		v18 = v8
		v19 = int32(0)
	} else {
		v13 = int32(45)
		*(*uint8)(unsafe.Add(mBase, uint32(v6))) = uint8(v13)
		v18 = int64(0) - v8
		v19 = int32(1)
	}
	v21 = F_pg_ulltoa_n(m, v18, v6+v19)
	mBase = m.M
	v22 = v21 + v19
	v24 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v6+v22))) = uint8(v24)
	v27 = v22 + int32(1)
	v28 = F_palloc(m, v27)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		return int64(0)
	} else {
		if v27 != 0 {
			base.MemoryCopy(m, v28, v6, v27)
		} else {
		}
		m.G0 = v6 + int32(32)
		return base.I64_extend_i32_u(v28)
	}
}
func F_intervaltypmodin(m *base.Module, l0 int32) int64 {
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
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	return int64(0)
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
	F_errfinish(m, int32(_a_F_intervaltypmodin_3), int32(1086), int32(_a_F_intervaltypmodin_4))
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
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L47
	}
L28:
	;
	m.G0 = v7 + int32(32)
	return base.I64_extend_i32_s(v132)
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
	F_errfinish(m, int32(_a_F_intervaltypmodin_3), int32(1109), int32(_a_F_intervaltypmodin_4))
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
	F_errfinish(m, int32(_a_F_intervaltypmodin_3), int32(1119), int32(_a_F_intervaltypmodin_4))
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
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v146
	F_errmsg(m, int32(_a_F_intervaltypmodin_2), v7)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(_a_F_intervaltypmodin_3), int32(1103), int32(_a_F_intervaltypmodin_4))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
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
	v40 = l2 + v25<<(uint(int32(3))%32) + v30*int32(100)
	v42 = v40 + int32(28)
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
	v60 = v40 + int32(32)
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
func F_intset(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_int_to_intset(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v3)
	}
}
func F_intset_union_elem(m *base.Module, l0 int32) int64 {
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
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int64(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v15 = F_intarray_add_elem(m, v10, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int64(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			if v17 != v10 {
				F_pfree(m, v10)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int64(0)
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
					v24 = F_ArrayGetNItemsSafe(m, v21, v15+int32(16))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int64(0)
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
							return int64(0)
						} else {
							m.G0 = v7 + int32(16)
							return base.I64_extend_i32_u(v41)
						}
					}
				}
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
				v24 = F_ArrayGetNItemsSafe(m, v21, v15+int32(16))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int64(0)
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
						return int64(0)
					} else {
						m.G0 = v7 + int32(16)
						return base.I64_extend_i32_u(v41)
					}
				}
			}
		}
	}
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
func F_ismn_cast_from_ean13(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14315(m, l0, int32(4))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_isort(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v385 int32
	_ = v385
	var v397 int32
	_ = v397
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
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v442 int32
	_ = v442
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v593 int32
	_ = v593
	var v611 int32
	_ = v611
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v631 int32
	_ = v631
	var v634 int32
	_ = v634
	var v658 int32
	_ = v658
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v692 int32
	_ = v692
	var v699 int32
	_ = v699
	var v725 int32
	_ = v725
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v746 int32
	_ = v746
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v785 int32
	_ = v785
	if base.Ui32(l1) < base.Ui32(int32(7)) {
		v674 = l0
		v675 = l1
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return
L2:
	;
	if base.Ui32(v699) < base.Ui32(int32(2)) {
		goto L1
	} else {
		goto L194
	}
L3:
	;
	v692 = v674
	v699 = v675
	goto L2
L4:
	;
	v21 = l0
	v22 = l1
	goto L5
L5:
	;
	v40 = v21 + int32(4)
	v42 = v22
	goto L7
L7:
	;
	if base.Ui32(v42) < base.Ui32(int32(2)) {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v63 = v21 + v42<<(uint(int32(2))%32)
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	v66 = v64 & int32(1)
	v72 = v40
	goto L10
L10:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v72-int32(4))))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	if v66 != 0 {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	v98 = v21 + v42<<(uint(int32(1))%32)&int32(-4)
	if v42 != int32(7) {
		goto L20
	} else {
		goto L21
	}
L12:
	;
	goto L11
L13:
	;
	v92 = v72 + int32(4)
	if base.Ui32(v92) < base.Ui32(v63) {
		v72 = v92
		goto L10
	} else {
		goto L19
	}
L14:
	;
	if v87 <= v88 {
		goto L13
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	if v87 < v88 {
		goto L12
	} else {
		goto L18
	}
L17:
	;
	goto L12
L18:
	;
	goto L13
L19:
	;
	goto L1
L20:
	;
	v102 = v63 - int32(4)
	if base.Ui32(v42) < base.Ui32(int32(41)) {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	v228 = v98
	goto L22
L22:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v228)))
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v228))) = v231
	v236 = v63 - int32(4)
	v242 = v40
	v243 = v236
	v244 = v40
	v245 = v236
	goto L127
L23:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v197)))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v196)))
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v206 == int32(1) {
		goto L106
	} else {
		goto L107
	}
L24:
	;
	v196 = v21
	v197 = v98
	v198 = v102
	goto L23
L25:
	;
	goto L26
L26:
	;
	v105 = int32(1)
	v108 = int32(base.Ui32(v42)>>(uint(v105)%32)) & int32(2147483644)
	v109 = v21 + v108
	v111 = v42 & int32(-8)
	v112 = v21 + v111
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v119 == v105 {
		goto L31
	} else {
		goto L32
	}
L27:
	;
	v139 = v98 - v108
	v140 = v98 + v108
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v140)))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v139)))
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v147 == int32(1) {
		goto L56
	} else {
		goto L57
	}
L28:
	;
	v138 = v136
	goto L27
L29:
	;
	if v116 < v117 {
		v136 = v109
		goto L28
	} else {
		goto L48
	}
L30:
	;
	if v117 < v116 {
		v136 = v109
		goto L28
	} else {
		goto L44
	}
L31:
	;
	if v118 < v117 {
		goto L30
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	if v117 < v118 {
		goto L29
	} else {
		goto L39
	}
L34:
	;
	if v116 < v117 {
		v136 = v109
		goto L28
	} else {
		goto L35
	}
L35:
	;
	if v118 < v116 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v125 = v21
	goto L38
L37:
	;
	v125 = v112
	goto L38
L38:
	;
	v138 = v125
	goto L27
L39:
	;
	if v117 < v116 {
		v136 = v109
		goto L28
	} else {
		goto L40
	}
L40:
	;
	if v116 < v118 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v129 = v21
	goto L43
L42:
	;
	v129 = v112
	goto L43
L43:
	;
	v138 = v129
	goto L27
L44:
	;
	if v118 < v116 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v132 = v112
	goto L47
L46:
	;
	v132 = v21
	goto L47
L47:
	;
	v138 = v132
	goto L27
L48:
	;
	if v116 < v118 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v135 = v112
	goto L51
L50:
	;
	v135 = v21
	goto L51
L51:
	;
	v136 = v135
	goto L28
L52:
	;
	v167 = v102 - v111
	v168 = v102 - v108
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v168)))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v175 == int32(1) {
		goto L81
	} else {
		goto L82
	}
L53:
	;
	v166 = v164
	goto L52
L54:
	;
	if v144 < v145 {
		v164 = v98
		goto L53
	} else {
		goto L73
	}
L55:
	;
	if v145 < v144 {
		v164 = v98
		goto L53
	} else {
		goto L69
	}
L56:
	;
	if v146 < v145 {
		goto L55
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	if v145 < v146 {
		goto L54
	} else {
		goto L64
	}
L59:
	;
	if v144 < v145 {
		v164 = v98
		goto L53
	} else {
		goto L60
	}
L60:
	;
	if v146 < v144 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v153 = v139
	goto L63
L62:
	;
	v153 = v140
	goto L63
L63:
	;
	v166 = v153
	goto L52
L64:
	;
	if v145 < v144 {
		v164 = v98
		goto L53
	} else {
		goto L65
	}
L65:
	;
	if v144 < v146 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v157 = v139
	goto L68
L67:
	;
	v157 = v140
	goto L68
L68:
	;
	v166 = v157
	goto L52
L69:
	;
	if v146 < v144 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v160 = v140
	goto L72
L71:
	;
	v160 = v139
	goto L72
L72:
	;
	v166 = v160
	goto L52
L73:
	;
	if v144 < v146 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v163 = v140
	goto L76
L75:
	;
	v163 = v139
	goto L76
L76:
	;
	v164 = v163
	goto L53
L77:
	;
	v196 = v138
	v197 = v166
	v198 = v194
	goto L23
L78:
	;
	v194 = v192
	goto L77
L79:
	;
	if v172 < v173 {
		v192 = v168
		goto L78
	} else {
		goto L98
	}
L80:
	;
	if v173 < v172 {
		v192 = v168
		goto L78
	} else {
		goto L94
	}
L81:
	;
	if v174 < v173 {
		goto L80
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	if v173 < v174 {
		goto L79
	} else {
		goto L89
	}
L84:
	;
	if v172 < v173 {
		v192 = v168
		goto L78
	} else {
		goto L85
	}
L85:
	;
	if v174 < v172 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v181 = v167
	goto L88
L87:
	;
	v181 = v102
	goto L88
L88:
	;
	v194 = v181
	goto L77
L89:
	;
	if v173 < v172 {
		v192 = v168
		goto L78
	} else {
		goto L90
	}
L90:
	;
	if v172 < v174 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v185 = v167
	goto L93
L92:
	;
	v185 = v102
	goto L93
L93:
	;
	v194 = v185
	goto L77
L94:
	;
	if v174 < v172 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v188 = v102
	goto L97
L96:
	;
	v188 = v167
	goto L97
L97:
	;
	v194 = v188
	goto L77
L98:
	;
	if v172 < v174 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v191 = v102
	goto L101
L100:
	;
	v191 = v167
	goto L101
L101:
	;
	v192 = v191
	goto L78
L102:
	;
	v228 = v225
	goto L22
L103:
	;
	v225 = v223
	goto L102
L104:
	;
	if v203 < v204 {
		v223 = v197
		goto L103
	} else {
		goto L123
	}
L105:
	;
	if v204 < v203 {
		v223 = v197
		goto L103
	} else {
		goto L119
	}
L106:
	;
	if v205 < v204 {
		goto L105
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	if v204 < v205 {
		goto L104
	} else {
		goto L114
	}
L109:
	;
	if v203 < v204 {
		v223 = v197
		goto L103
	} else {
		goto L110
	}
L110:
	;
	if v205 < v203 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v212 = v196
	goto L113
L112:
	;
	v212 = v198
	goto L113
L113:
	;
	v225 = v212
	goto L102
L114:
	;
	if v204 < v203 {
		v223 = v197
		goto L103
	} else {
		goto L115
	}
L115:
	;
	if v203 < v205 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v216 = v196
	goto L118
L117:
	;
	v216 = v198
	goto L118
L118:
	;
	v225 = v216
	goto L102
L119:
	;
	if v205 < v203 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v219 = v198
	goto L122
L121:
	;
	v219 = v196
	goto L122
L122:
	;
	v225 = v219
	goto L102
L123:
	;
	if v203 < v205 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v222 = v198
	goto L126
L125:
	;
	v222 = v196
	goto L126
L126:
	;
	v223 = v222
	goto L103
L127:
	;
	if base.Ui32(v243) < base.Ui32(v242) {
		v295 = v242
		v297 = v244
		goto L129
	} else {
		goto L130
	}
L129:
	;
	if base.Ui32(v295) <= base.Ui32(v243) {
		goto L144
	} else {
		goto L145
	}
L130:
	;
	v261 = v242
	v263 = v244
	goto L131
L131:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v261)))
	if v66 != 0 {
		goto L135
	} else {
		goto L136
	}
L132:
	;
	v295 = v288
	v297 = v286
	goto L129
L133:
	;
	v288 = v261 + int32(4)
	if base.Ui32(v288) <= base.Ui32(v243) {
		v261 = v288
		v263 = v286
		goto L131
	} else {
		goto L142
	}
L134:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v263)))
	*(*int32)(unsafe.Add(mBase, uint32(v263))) = v275
	*(*int32)(unsafe.Add(mBase, uint32(v261))) = v280
	v286 = v263 + int32(4)
	goto L133
L135:
	;
	if v275 < v274 {
		v286 = v263
		goto L133
	} else {
		goto L138
	}
L136:
	;
	goto L137
L137:
	;
	if v274 < v275 {
		v286 = v263
		goto L133
	} else {
		goto L140
	}
L138:
	;
	if v275 <= v274 {
		goto L134
	} else {
		goto L139
	}
L139:
	;
	v295 = v261
	v297 = v263
	goto L129
L140:
	;
	if v275 < v274 {
		v295 = v261
		v297 = v263
		goto L129
	} else {
		goto L141
	}
L141:
	;
	goto L134
L142:
	;
	goto L132
L143:
	;
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v295)))
	*(*int32)(unsafe.Add(mBase, uint32(v295))) = v328
	*(*int32)(unsafe.Add(mBase, uint32(v315))) = v667
	v670 = int32(4)
	v242 = v295 + v670
	v243 = v315 - v670
	v244 = v297
	v245 = v317
	goto L127
L144:
	;
	v315 = v243
	v317 = v245
	goto L147
L145:
	;
	v348 = v243
	v350 = v245
	goto L146
L146:
	;
	v361 = int32(2)
	v362 = (v297 - v21) >> (uint(v361) % 32)
	v365 = (v295 - v297) >> (uint(v361) % 32)
	if v362 < v365 {
		goto L160
	} else {
		goto L161
	}
L147:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v315)))
	if v66 != 0 {
		goto L151
	} else {
		goto L152
	}
L148:
	;
	v348 = v340
	v350 = v338
	goto L146
L149:
	;
	v340 = v315 - int32(4)
	if base.Ui32(v295) <= base.Ui32(v340) {
		v315 = v340
		v317 = v338
		goto L147
	} else {
		goto L158
	}
L150:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v317)))
	*(*int32)(unsafe.Add(mBase, uint32(v315))) = v333
	*(*int32)(unsafe.Add(mBase, uint32(v317))) = v328
	v338 = v317 - int32(4)
	goto L149
L151:
	;
	if v328 < v327 {
		goto L143
	} else {
		goto L154
	}
L152:
	;
	goto L153
L153:
	;
	if v327 < v328 {
		goto L143
	} else {
		goto L156
	}
L154:
	;
	if v328 <= v327 {
		goto L150
	} else {
		goto L155
	}
L155:
	;
	v338 = v317
	goto L149
L156:
	;
	if v328 < v327 {
		v338 = v317
		goto L149
	} else {
		goto L157
	}
L157:
	;
	goto L150
L158:
	;
	goto L148
L159:
	;
	v507 = int32(2)
	v508 = (v350 - v348) >> (uint(v507) % 32)
	v513 = (v63-v350)>>(uint(v507)%32) - int32(1)
	if v508 < v513 {
		goto L175
	} else {
		goto L176
	}
L160:
	;
	v367 = v362
	goto L162
L161:
	;
	v367 = v365
	goto L162
L162:
	;
	if v367 == int32(0) {
		goto L159
	} else {
		goto L163
	}
L163:
	;
	v372 = v295 - v367<<(uint(int32(2))%32)
	v374 = v367 & int32(3)
	v375 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v367) {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v385 = v375
	v397 = int32(0)
	goto L167
L165:
	;
	v442 = v375
	goto L166
L166:
	;
	v460 = v442
	v464 = v375
	goto L171
L167:
	;
	v401 = v385 << (uint(int32(2)) % 32)
	v402 = v21 + v401
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v402)))
	v404 = v401 + v372
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v404)))
	*(*int32)(unsafe.Add(mBase, uint32(v402))) = v405
	*(*int32)(unsafe.Add(mBase, uint32(v404))) = v403
	v408 = int32(4)
	v409 = v401 | v408
	v410 = v21 + v409
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v410)))
	v412 = v409 + v372
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v412)))
	*(*int32)(unsafe.Add(mBase, uint32(v410))) = v413
	*(*int32)(unsafe.Add(mBase, uint32(v412))) = v411
	v417 = v401 | int32(8)
	v418 = v21 + v417
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v418)))
	v420 = v417 + v372
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v420)))
	*(*int32)(unsafe.Add(mBase, uint32(v418))) = v421
	*(*int32)(unsafe.Add(mBase, uint32(v420))) = v419
	v425 = v401 | int32(12)
	v426 = v21 + v425
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v426)))
	v428 = v425 + v372
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v428)))
	*(*int32)(unsafe.Add(mBase, uint32(v426))) = v429
	*(*int32)(unsafe.Add(mBase, uint32(v428))) = v427
	v433 = v385 + v408
	v435 = v397 + v408
	if v435 != v367&int32(-4) {
		v385 = v433
		v397 = v435
		goto L167
	} else {
		goto L169
	}
L168:
	;
	if v374 == int32(0) {
		goto L159
	} else {
		goto L170
	}
L169:
	;
	goto L168
L170:
	;
	v442 = v433
	goto L166
L171:
	;
	v476 = v460 << (uint(int32(2)) % 32)
	v477 = v21 + v476
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v477)))
	v479 = v476 + v372
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v479)))
	*(*int32)(unsafe.Add(mBase, uint32(v477))) = v480
	*(*int32)(unsafe.Add(mBase, uint32(v479))) = v478
	v483 = int32(1)
	v486 = v464 + v483
	if v486 != v374 {
		v460 = v460 + v483
		v464 = v486
		goto L171
	} else {
		goto L173
	}
L172:
	;
	goto L159
L173:
	;
	goto L172
L174:
	;
	if base.Ui32(v365) <= base.Ui32(v508) {
		goto L189
	} else {
		goto L190
	}
L175:
	;
	v515 = v508
	goto L177
L176:
	;
	v515 = v513
	goto L177
L177:
	;
	if v515 == int32(0) {
		goto L174
	} else {
		goto L178
	}
L178:
	;
	v520 = v63 - v515<<(uint(int32(2))%32)
	v522 = v515 & int32(3)
	v523 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v515) {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	v536 = v523
	v538 = int32(0)
	goto L182
L180:
	;
	v593 = v523
	goto L181
L181:
	;
	v611 = v593
	v622 = v523
	goto L186
L182:
	;
	v549 = v536 << (uint(int32(2)) % 32)
	v550 = v295 + v549
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v550)))
	v552 = v520 + v549
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v552)))
	*(*int32)(unsafe.Add(mBase, uint32(v550))) = v553
	*(*int32)(unsafe.Add(mBase, uint32(v552))) = v551
	v556 = int32(4)
	v557 = v549 | v556
	v558 = v295 + v557
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v558)))
	v560 = v557 + v520
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v560)))
	*(*int32)(unsafe.Add(mBase, uint32(v558))) = v561
	*(*int32)(unsafe.Add(mBase, uint32(v560))) = v559
	v565 = v549 | int32(8)
	v566 = v295 + v565
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v566)))
	v568 = v565 + v520
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v568)))
	*(*int32)(unsafe.Add(mBase, uint32(v566))) = v569
	*(*int32)(unsafe.Add(mBase, uint32(v568))) = v567
	v573 = v549 | int32(12)
	v574 = v295 + v573
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v574)))
	v576 = v573 + v520
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v576)))
	*(*int32)(unsafe.Add(mBase, uint32(v574))) = v577
	*(*int32)(unsafe.Add(mBase, uint32(v576))) = v575
	v581 = v536 + v556
	v583 = v538 + v556
	if v583 != v515&int32(-4) {
		v536 = v581
		v538 = v583
		goto L182
	} else {
		goto L184
	}
L183:
	;
	if v522 == int32(0) {
		goto L174
	} else {
		goto L185
	}
L184:
	;
	goto L183
L185:
	;
	v593 = v581
	goto L181
L186:
	;
	v624 = v611 << (uint(int32(2)) % 32)
	v625 = v295 + v624
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v625)))
	v627 = v624 + v520
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v627)))
	*(*int32)(unsafe.Add(mBase, uint32(v625))) = v628
	*(*int32)(unsafe.Add(mBase, uint32(v627))) = v626
	v631 = int32(1)
	v634 = v622 + v631
	if v634 != v522 {
		v611 = v611 + v631
		v622 = v634
		goto L186
	} else {
		goto L188
	}
L187:
	;
	goto L174
L188:
	;
	goto L187
L189:
	;
	F_isort(m, v21, v365, l2)
	mBase = m.M
	v658 = v63 - v508<<(uint(int32(2))%32)
	if base.Ui32(v508) < base.Ui32(int32(7)) {
		v692 = v658
		v699 = v508
		goto L2
	} else {
		goto L192
	}
L190:
	;
	goto L191
L191:
	;
	F_isort(m, v63-v508<<(uint(int32(2))%32), v508, l2)
	mBase = m.M
	if base.Ui32(v365) < base.Ui32(int32(7)) {
		v674 = v21
		v675 = v365
		goto L3
	} else {
		goto L193
	}
L192:
	;
	v21 = v658
	v22 = v508
	goto L5
L193:
	;
	v42 = v365
	goto L7
L194:
	;
	v725 = v692 + int32(4)
	goto L195
L195:
	;
	if base.Ui32(v725) <= base.Ui32(v692) {
		goto L197
	} else {
		goto L198
	}
L196:
	;
	goto L1
L197:
	;
	v785 = v725 + int32(4)
	if base.Ui32(v785) < base.Ui32(v692+v699<<(uint(int32(2))%32)) {
		v725 = v785
		goto L195
	} else {
		goto L208
	}
L198:
	;
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v725)))
	v737 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	v746 = v725
	goto L199
L199:
	;
	v759 = v746 - int32(4)
	v760 = *(*int32)(unsafe.Add(mBase, uint32(v759)))
	if v737&int32(1) != 0 {
		goto L202
	} else {
		goto L203
	}
L200:
	;
	goto L197
L201:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v746))) = v760
	*(*int32)(unsafe.Add(mBase, uint32(v759))) = v736
	if base.Ui32(v692) < base.Ui32(v759) {
		v746 = v759
		goto L199
	} else {
		goto L207
	}
L202:
	;
	if v736 < v760 {
		goto L201
	} else {
		goto L205
	}
L203:
	;
	goto L204
L204:
	;
	if v736 <= v760 {
		goto L197
	} else {
		goto L206
	}
L205:
	;
	goto L197
L206:
	;
	goto L201
L207:
	;
	goto L200
L208:
	;
	goto L196
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
func F_issn_in(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14257(m, l0, int32(5))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
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
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v247 int32
	_ = v247
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v313 int32
	_ = v313
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v422 int32
	_ = v422
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v471 int32
	_ = v471
	var v484 int32
	_ = v484
	var v489 int32
	_ = v489
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v507 int32
	_ = v507
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v524 int32
	_ = v524
	var v533 int32
	_ = v533
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v599 int32
	_ = v599
	var v612 int32
	_ = v612
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v649 int32
	_ = v649
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v683 int32
	_ = v683
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v698 int32
	_ = v698
	var v706 int32
	_ = v706
	var v714 int32
	_ = v714
	var v717 int32
	_ = v717
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v747 int32
	_ = v747
	var v749 int32
	_ = v749
	var v755 int32
	_ = v755
	var v768 int32
	_ = v768
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v781 int32
	_ = v781
	var v794 int32
	_ = v794
	var v796 int32
	_ = v796
	var v803 int32
	_ = v803
	var v807 int32
	_ = v807
	var v809 int32
	_ = v809
	var v811 int32
	_ = v811
	var v814 int32
	_ = v814
	var v818 int32
	_ = v818
	var v826 int32
	_ = v826
	var v834 int32
	_ = v834
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v849 int32
	_ = v849
	var v851 int32
	_ = v851
	var v857 int32
	_ = v857
	var v864 int32
	_ = v864
	var v866 int32
	_ = v866
	var v868 int32
	_ = v868
	var v874 int32
	_ = v874
	var v883 int32
	_ = v883
	var v892 int32
	_ = v892
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v907 int32
	_ = v907
	var v909 int32
	_ = v909
	var v916 int32
	_ = v916
	var v920 int32
	_ = v920
	var v922 int32
	_ = v922
	var v924 int32
	_ = v924
	var v927 int32
	_ = v927
	var v931 int32
	_ = v931
	var v939 int32
	_ = v939
	var v947 int32
	_ = v947
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v962 int32
	_ = v962
	var v964 int32
	_ = v964
	var v970 int32
	_ = v970
	var v977 int32
	_ = v977
	var v979 int32
	_ = v979
	var v981 int32
	_ = v981
	var v987 int32
	_ = v987
	var v996 int32
	_ = v996
	var v1005 int32
	_ = v1005
	var v1008 int32
	_ = v1008
	var v1013 int32
	_ = v1013
	var v1017 int32
	_ = v1017
	var v1019 int32
	_ = v1019
	var v1021 int32
	_ = v1021
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1040 int32
	_ = v1040
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1046 int32
	_ = v1046
	var v1048 int32
	_ = v1048
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1063 int32
	_ = v1063
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1073 int32
	_ = v1073
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1083 int32
	_ = v1083
	var v1087 int32
	_ = v1087
	var v1089 int32
	_ = v1089
	var v1092 int32
	_ = v1092
	var v1094 int32
	_ = v1094
	var v1097 int32
	_ = v1097
	var v1099 int32
	_ = v1099
	var v1101 int32
	_ = v1101
	var v1104 int32
	_ = v1104
	var v1107 int32
	_ = v1107
	var v1110 int32
	_ = v1110
	var v1114 int32
	_ = v1114
	var v1117 int32
	_ = v1117
	var v1119 int32
	_ = v1119
	var v1121 int32
	_ = v1121
	var v1124 int32
	_ = v1124
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1132 int32
	_ = v1132
	var v1136 int32
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1140 int32
	_ = v1140
	var v1144 int32
	_ = v1144
	var v1145 int32
	_ = v1145
	var v1148 int32
	_ = v1148
	var v1150 int32
	_ = v1150
	var v1153 int32
	_ = v1153
	var v1155 int32
	_ = v1155
	var v1158 int32
	_ = v1158
	var v1161 int32
	_ = v1161
	var v1162 int32
	_ = v1162
	var v1164 int32
	_ = v1164
	var v1166 int32
	_ = v1166
	var v1181 int32
	_ = v1181
	var v1182 int32
	_ = v1182
	var v1185 int32
	_ = v1185
	var v1187 int32
	_ = v1187
	var v1189 int32
	_ = v1189
	var v1194 int32
	_ = v1194
	var v1196 int32
	_ = v1196
	var v1198 int32
	_ = v1198
	var v1201 int32
	_ = v1201
	var v1204 int32
	_ = v1204
	var v1207 int32
	_ = v1207
	var v1211 int32
	_ = v1211
	var v1214 int32
	_ = v1214
	var v1216 int32
	_ = v1216
	var v1218 int32
	_ = v1218
	var v1221 int32
	_ = v1221
	var v1223 int32
	_ = v1223
	var v1226 int32
	_ = v1226
	var v1229 int32
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1232 int32
	_ = v1232
	var v1234 int32
	_ = v1234
	var v1249 int32
	_ = v1249
	var v1250 int32
	_ = v1250
	var v1253 int32
	_ = v1253
	var v1255 int32
	_ = v1255
	var v1257 int32
	_ = v1257
	var v1260 int32
	_ = v1260
	var v1262 int32
	_ = v1262
	var v1265 int32
	_ = v1265
	var v1267 int32
	_ = v1267
	var v1269 int32
	_ = v1269
	var v1272 int32
	_ = v1272
	var v1275 int32
	_ = v1275
	var v1278 int32
	_ = v1278
	var v1282 int32
	_ = v1282
	var v1285 int32
	_ = v1285
	var v1287 int32
	_ = v1287
	var v1289 int32
	_ = v1289
	var v1292 int32
	_ = v1292
	var v1294 int32
	_ = v1294
	var v1296 int32
	_ = v1296
	var v1299 int32
	_ = v1299
	var v1302 int32
	_ = v1302
	var v1305 int32
	_ = v1305
	var v1309 int32
	_ = v1309
	var v1312 int32
	_ = v1312
	var v1314 int32
	_ = v1314
	var v1316 int32
	_ = v1316
	var v1320 int32
	_ = v1320
	var v1322 int32
	_ = v1322
	var v1325 int32
	_ = v1325
	var v1330 int32
	_ = v1330
	var v1331 int32
	_ = v1331
	var v1332 int32
	_ = v1332
	var v1334 int32
	_ = v1334
	var v1342 int32
	_ = v1342
	var v1354 int32
	_ = v1354
	var v1366 int32
	_ = v1366
	var v1367 int32
	_ = v1367
	var v1371 int32
	_ = v1371
	var v1373 int32
	_ = v1373
	var v1379 int32
	_ = v1379
	var v1393 int32
	_ = v1393
	var v1397 int32
	_ = v1397
	var v1398 int32
	_ = v1398
	var v1400 int32
	_ = v1400
	var v1402 int32
	_ = v1402
	var v1405 int32
	_ = v1405
	var v1407 int32
	_ = v1407
	var v1409 int32
	_ = v1409
	var v1413 int32
	_ = v1413
	var v1417 int32
	_ = v1417
	var v1420 int32
	_ = v1420
	var v1422 int32
	_ = v1422
	var v1425 int32
	_ = v1425
	var v1428 int32
	_ = v1428
	var v1429 int32
	_ = v1429
	var v1434 int32
	_ = v1434
	var v1436 int32
	_ = v1436
	var v1439 int32
	_ = v1439
	var v1441 int32
	_ = v1441
	var v1445 int32
	_ = v1445
	var v1449 int32
	_ = v1449
	var v1461 int32
	_ = v1461
	var v1473 int32
	_ = v1473
	var v1474 int32
	_ = v1474
	var v1478 int32
	_ = v1478
	var v1480 int32
	_ = v1480
	var v1486 int32
	_ = v1486
	var v1500 int32
	_ = v1500
	var v1504 int32
	_ = v1504
	var v1505 int32
	_ = v1505
	var v1506 int32
	_ = v1506
	var v1508 int32
	_ = v1508
	var v1512 int32
	_ = v1512
	var v1519 int32
	_ = v1519
	var v1521 int32
	_ = v1521
	var v1523 int32
	_ = v1523
	var v1525 int32
	_ = v1525
	var v1527 int32
	_ = v1527
	var v1528 int32
	_ = v1528
	var v1538 int32
	_ = v1538
	var v1539 int32
	_ = v1539
	var v1540 int32
	_ = v1540
	var v1544 int32
	_ = v1544
	var v1547 int32
	_ = v1547
	var v1548 int32
	_ = v1548
	var v1553 int32
	_ = v1553
	var v1554 int32
	_ = v1554
	var v1559 int32
	_ = v1559
	var v1560 int32
	_ = v1560
	var v1567 int32
	_ = v1567
	var v1575 int32
	_ = v1575
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v8 = v6
	goto L2
L1:
	;
	return v1575
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v8
	v16 = F_find_among(m, l0, int32(_a_F_italian_ISO_8859_1_stem_0), int32(7), int32(0))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v69 = v6
	goto L31
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
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v8 = v66
	goto L2
L8:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v60 <= v20 {
		goto L27
	} else {
		goto L28
	}
L9:
	;
	v56 = F_slice_from_s(m, l0, int32(2), int32(_a_F_italian_ISO_8859_1_stem_1))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L5
	} else {
		goto L25
	}
L10:
	;
	v50 = F_slice_from_s(m, l0, int32(1), int32(_a_F_italian_ISO_8859_1_stem_2))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L5
	} else {
		goto L23
	}
L11:
	;
	v44 = F_slice_from_s(m, l0, int32(1), int32(_a_F_italian_ISO_8859_1_stem_3))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L5
	} else {
		goto L21
	}
L12:
	;
	v38 = F_slice_from_s(m, l0, int32(1), int32(_a_F_italian_ISO_8859_1_stem_4))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L5
	} else {
		goto L19
	}
L13:
	;
	v32 = F_slice_from_s(m, l0, int32(1), int32(_a_F_italian_ISO_8859_1_stem_5))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L5
	} else {
		goto L17
	}
L14:
	;
	v26 = F_slice_from_s(m, l0, int32(1), int32(_a_F_italian_ISO_8859_1_stem_6))
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
	v1575 = v26
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
	v1575 = v32
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
	v1575 = v38
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
	v1575 = v44
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
	v1575 = v50
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
	v1575 = v56
	goto L1
L27:
	;
	goto L4
L28:
	;
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v20 + int32(1)
	goto L7
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1436
	v1439 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1436 <= v1439 {
		goto L427
	} else {
		goto L428
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v69
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v82 < v69 {
		goto L35
	} else {
		goto L36
	}
L32:
	;
	v1434 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1434
	v1436 = v1434
	goto L30
L33:
	;
	goto L32
L34:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v125 != 0 {
		v269 = v126
		goto L49
	} else {
		goto L50
	}
L35:
	;
	v84 = v69
	goto L37
L36:
	;
	v84 = v82
	goto L37
L37:
	;
	goto L39
L38:
	;
	v125 = v120
	goto L34
L39:
	;
	if v69 == v84 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v120 = int32(0)
	goto L38
L41:
	;
	v125 = int32(-1)
	goto L34
L42:
	;
	goto L43
L43:
	;
	v96 = int32(1)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97+v69))))
	if int32(249) < v99 {
		v120 = v96
		goto L38
	} else {
		goto L44
	}
L44:
	;
	v101 = v99 - int32(97)
	if v101 < int32(0) {
		v120 = v96
		goto L38
	} else {
		goto L45
	}
L45:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v101)>>(uint(int32(3))%32)))+uint32(_c_F_italian_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v107)>>(uint(v101&int32(7))%32))&int32(1) == int32(0) {
		v120 = v96
		goto L38
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v69 + int32(1)
	goto L47
L47:
	;
	goto L40
L48:
	;
	v1428 = F_slice_from_s(m, l0, int32(1), int32(_a_F_italian_ISO_8859_1_stem_7))
	mBase = m.M
	v1429 = m.ExcPending
	if v1429 != 0 {
		goto L5
	} else {
		goto L425
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v69
	if v69 < v269 {
		goto L92
	} else {
		goto L93
	}
L50:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v127
	if v126 == v127 {
		v201 = v126
		goto L51
	} else {
		goto L52
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v127
	if v201 == v127 {
		goto L73
	} else {
		goto L74
	}
L52:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130+v127))))
	if v132 != int32(117) {
		v201 = v126
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v136 = v127 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v136
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v136
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v148 < v136 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	if v191 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L55:
	;
	v150 = v136
	goto L57
L56:
	;
	v150 = v148
	goto L57
L57:
	;
	goto L59
L58:
	;
	v191 = v186
	goto L54
L59:
	;
	if v136 == v150 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v186 = int32(0)
	goto L58
L61:
	;
	v191 = int32(-1)
	goto L54
L62:
	;
	goto L63
L63:
	;
	v162 = int32(1)
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163+v136))))
	if int32(249) < v165 {
		v186 = v162
		goto L58
	} else {
		goto L64
	}
L64:
	;
	v167 = v165 - int32(97)
	if v167 < int32(0) {
		v186 = v162
		goto L58
	} else {
		goto L65
	}
L65:
	;
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v167)>>(uint(int32(3))%32)))+uint32(_c_F_italian_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v173)>>(uint(v167&int32(7))%32))&int32(1) == int32(0) {
		v186 = v162
		goto L58
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v127 + int32(2)
	goto L67
L67:
	;
	goto L60
L68:
	;
	v196 = F_slice_from_s(m, l0, int32(1), int32(_a_F_italian_ISO_8859_1_stem_8))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L5
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v201 = v200
	goto L51
L71:
	;
	if v196 < int32(0) {
		v1575 = v196
		goto L1
	} else {
		goto L72
	}
L72:
	;
	goto L31
L73:
	;
	v269 = v127
	goto L49
L74:
	;
	goto L75
L75:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204+v127))))
	if v206 != int32(105) {
		v269 = v201
		goto L49
	} else {
		goto L76
	}
L76:
	;
	v210 = v127 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v210
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v210
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v222 < v210 {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	if v265 == int32(0) {
		goto L48
	} else {
		goto L91
	}
L78:
	;
	v224 = v210
	goto L80
L79:
	;
	v224 = v222
	goto L80
L80:
	;
	goto L82
L81:
	;
	v265 = v260
	goto L77
L82:
	;
	if v210 == v224 {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	v260 = int32(0)
	goto L81
L84:
	;
	v265 = int32(-1)
	goto L77
L85:
	;
	goto L86
L86:
	;
	v236 = int32(1)
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237+v210))))
	if int32(249) < v239 {
		v260 = v236
		goto L81
	} else {
		goto L87
	}
L87:
	;
	v241 = v239 - int32(97)
	if v241 < int32(0) {
		v260 = v236
		goto L81
	} else {
		goto L88
	}
L88:
	;
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v241)>>(uint(int32(3))%32)))+uint32(_c_F_italian_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v247)>>(uint(v241&int32(7))%32))&int32(1) == int32(0) {
		v260 = v236
		goto L81
	} else {
		goto L89
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v127 + int32(2)
	goto L90
L90:
	;
	goto L83
L91:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v269 = v268
	goto L49
L92:
	;
	v69 = v69 + int32(1)
	goto L31
L93:
	;
	goto L94
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v269
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v269
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v269
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v288 < v6 {
		goto L99
	} else {
		goto L100
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v6
	v794 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v794 < v6 {
		goto L250
	} else {
		goto L251
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v781
	goto L95
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v6
	v550 = int32(5)
	v552 = int32(0)
	v554 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v554-v6 < v550 {
		v564 = v552
		goto L177
	} else {
		goto L178
	}
L98:
	;
	if v331 != 0 {
		goto L97
	} else {
		goto L112
	}
L99:
	;
	v290 = v6
	goto L101
L100:
	;
	v290 = v288
	goto L101
L101:
	;
	goto L103
L102:
	;
	v331 = v326
	goto L98
L103:
	;
	if v6 == v290 {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	v326 = int32(0)
	goto L102
L105:
	;
	v331 = int32(-1)
	goto L98
L106:
	;
	goto L107
L107:
	;
	v302 = int32(1)
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v303+v6))))
	if int32(249) < v305 {
		v326 = v302
		goto L102
	} else {
		goto L108
	}
L108:
	;
	v307 = v305 - int32(97)
	if v307 < int32(0) {
		v326 = v302
		goto L102
	} else {
		goto L109
	}
L109:
	;
	v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v307)>>(uint(int32(3))%32)))+uint32(_c_F_italian_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v313)>>(uint(v307&int32(7))%32))&int32(1) == int32(0) {
		v326 = v302
		goto L102
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v6 + int32(1)
	goto L111
L111:
	;
	goto L104
L112:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v341 < v332 {
		goto L115
	} else {
		goto L116
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v332
	v446 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v446 < v332 {
		goto L147
	} else {
		goto L148
	}
L114:
	;
	if v381 != 0 {
		goto L113
	} else {
		goto L129
	}
L115:
	;
	v343 = v332
	goto L117
L116:
	;
	v343 = v341
	goto L117
L117:
	;
	goto L119
L118:
	;
	v381 = v378
	goto L114
L119:
	;
	if v332 == v343 {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	v378 = int32(0)
	goto L118
L121:
	;
	v381 = int32(-1)
	goto L114
L122:
	;
	goto L123
L123:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v354+v332))))
	if int32(249) < v356 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v332 + int32(1)
	goto L128
L125:
	;
	v358 = v356 - int32(97)
	if v358 < int32(0) {
		goto L124
	} else {
		goto L126
	}
L126:
	;
	v361 = int32(1)
	v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v358)>>(uint(int32(3))%32)))+uint32(_c_F_italian_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v365)>>(uint(v358&int32(7))%32))&v361 != 0 {
		v378 = v361
		goto L118
	} else {
		goto L127
	}
L127:
	;
	goto L124
L128:
	;
	goto L120
L129:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v390 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v390 < v389 {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	if v430 < int32(0) {
		goto L113
	} else {
		goto L145
	}
L131:
	;
	v392 = v389
	goto L133
L132:
	;
	v392 = v390
	goto L133
L133:
	;
	v399 = v389
	goto L135
L134:
	;
	v430 = v410
	goto L130
L135:
	;
	if v399 == v392 {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v430 = int32(-1)
	goto L130
L138:
	;
	goto L139
L139:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v405 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v403+v399))))
	if int32(249) < v405 {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v422 = v399 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v422
	v399 = v422
	goto L135
L141:
	;
	v407 = v405 - int32(97)
	if v407 < int32(0) {
		goto L140
	} else {
		goto L142
	}
L142:
	;
	v410 = int32(1)
	v414 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v407)>>(uint(int32(3))%32)))+uint32(_c_F_italian_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v414)>>(uint(v407&int32(7))%32))&v410 != 0 {
		goto L134
	} else {
		goto L143
	}
L143:
	;
	goto L140
L145:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v781 = v433 + v430
	goto L96
L146:
	;
	if v489 != 0 {
		goto L97
	} else {
		goto L160
	}
L147:
	;
	v448 = v332
	goto L149
L148:
	;
	v448 = v446
	goto L149
L149:
	;
	goto L151
L150:
	;
	v489 = v484
	goto L146
L151:
	;
	if v332 == v448 {
		goto L153
	} else {
		goto L154
	}
L152:
	;
	v484 = int32(0)
	goto L150
L153:
	;
	v489 = int32(-1)
	goto L146
L154:
	;
	goto L155
L155:
	;
	v460 = int32(1)
	v461 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v461+v332))))
	if int32(249) < v463 {
		v484 = v460
		goto L150
	} else {
		goto L156
	}
L156:
	;
	v465 = v463 - int32(97)
	if v465 < int32(0) {
		v484 = v460
		goto L150
	} else {
		goto L157
	}
L157:
	;
	v471 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v465)>>(uint(int32(3))%32)))+uint32(_c_F_italian_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v471)>>(uint(v465&int32(7))%32))&int32(1) == int32(0) {
		v484 = v460
		goto L150
	} else {
		goto L158
	}
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v332 + int32(1)
	goto L159
L159:
	;
	goto L152
L160:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v499 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v499 < v498 {
		goto L162
	} else {
		goto L163
	}
L161:
	;
	if v542 < int32(0) {
		goto L97
	} else {
		goto L175
	}
L162:
	;
	v501 = v498
	goto L164
L163:
	;
	v501 = v499
	goto L164
L164:
	;
	v507 = v498
	goto L166
L165:
	;
	v542 = int32(1)
	goto L161
L166:
	;
	if v507 == v501 {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v542 = int32(-1)
	goto L161
L169:
	;
	goto L170
L170:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v516 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514+v507))))
	if int32(249) < v516 {
		goto L165
	} else {
		goto L171
	}
L171:
	;
	v518 = v516 - int32(97)
	if v518 < int32(0) {
		goto L165
	} else {
		goto L172
	}
L172:
	;
	v524 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v518)>>(uint(int32(3))%32)))+uint32(_c_F_italian_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v524)>>(uint(v518&int32(7))%32))&int32(1) == int32(0) {
		goto L165
	} else {
		goto L173
	}
L173:
	;
	v533 = v507 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v533
	v507 = v533
	goto L166
L175:
	;
	v545 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v781 = v545 + v542
	goto L96
L176:
	;
	if v564 != 0 {
		goto L180
	} else {
		goto L181
	}
L177:
	;
	goto L176
L178:
	;
	v558 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v560 = F_memcmp(m, v558+v6, int32(_a_F_italian_ISO_8859_1_stem_9), v550)
	mBase = m.M
	if v560 != 0 {
		v564 = v552
		goto L177
	} else {
		goto L179
	}
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v550 + v6
	v564 = int32(1)
	goto L177
L180:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v781 = v565
	goto L96
L181:
	;
	goto L182
L182:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v6
	v575 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v575 < v6 {
		goto L184
	} else {
		goto L185
	}
L183:
	;
	if v615 != 0 {
		goto L95
	} else {
		goto L198
	}
L184:
	;
	v577 = v6
	goto L186
L185:
	;
	v577 = v575
	goto L186
L186:
	;
	goto L188
L187:
	;
	v615 = v612
	goto L183
L188:
	;
	if v6 == v577 {
		goto L190
	} else {
		goto L191
	}
L189:
	;
	v612 = int32(0)
	goto L187
L190:
	;
	v615 = int32(-1)
	goto L183
L191:
	;
	goto L192
L192:
	;
	v588 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v590 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v588+v6))))
	if int32(249) < v590 {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v6 + int32(1)
	goto L197
L194:
	;
	v592 = v590 - int32(97)
	if v592 < int32(0) {
		goto L193
	} else {
		goto L195
	}
L195:
	;
	v595 = int32(1)
	v599 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v592)>>(uint(int32(3))%32)))+uint32(_c_F_italian_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v599)>>(uint(v592&int32(7))%32))&v595 != 0 {
		v612 = v595
		goto L187
	} else {
		goto L196
	}
L196:
	;
	goto L193
L197:
	;
	goto L189
L198:
	;
	v616 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v625 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v625 < v616 {
		goto L201
	} else {
		goto L202
	}
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v616
	v730 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v730 < v616 {
		goto L233
	} else {
		goto L234
	}
L200:
	;
	if v665 != 0 {
		goto L199
	} else {
		goto L215
	}
L201:
	;
	v627 = v616
	goto L203
L202:
	;
	v627 = v625
	goto L203
L203:
	;
	goto L205
L204:
	;
	v665 = v662
	goto L200
L205:
	;
	if v616 == v627 {
		goto L207
	} else {
		goto L208
	}
L206:
	;
	v662 = int32(0)
	goto L204
L207:
	;
	v665 = int32(-1)
	goto L200
L208:
	;
	goto L209
L209:
	;
	v638 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v640 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v638+v616))))
	if int32(249) < v640 {
		goto L210
	} else {
		goto L211
	}
L210:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v616 + int32(1)
	goto L214
L211:
	;
	v642 = v640 - int32(97)
	if v642 < int32(0) {
		goto L210
	} else {
		goto L212
	}
L212:
	;
	v645 = int32(1)
	v649 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v642)>>(uint(int32(3))%32)))+uint32(_c_F_italian_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v649)>>(uint(v642&int32(7))%32))&v645 != 0 {
		v662 = v645
		goto L204
	} else {
		goto L213
	}
L213:
	;
	goto L210
L214:
	;
	goto L206
L215:
	;
	v673 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v674 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v674 < v673 {
		goto L217
	} else {
		goto L218
	}
L216:
	;
	if v714 < int32(0) {
		goto L199
	} else {
		goto L231
	}
L217:
	;
	v676 = v673
	goto L219
L218:
	;
	v676 = v674
	goto L219
L219:
	;
	v683 = v673
	goto L221
L220:
	;
	v714 = v694
	goto L216
L221:
	;
	if v683 == v676 {
		goto L223
	} else {
		goto L224
	}
L223:
	;
	v714 = int32(-1)
	goto L216
L224:
	;
	goto L225
L225:
	;
	v687 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v687+v683))))
	if int32(249) < v689 {
		goto L226
	} else {
		goto L227
	}
L226:
	;
	v706 = v683 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v706
	v683 = v706
	goto L221
L227:
	;
	v691 = v689 - int32(97)
	if v691 < int32(0) {
		goto L226
	} else {
		goto L228
	}
L228:
	;
	v694 = int32(1)
	v698 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v691)>>(uint(int32(3))%32)))+uint32(_c_F_italian_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v698)>>(uint(v691&int32(7))%32))&v694 != 0 {
		goto L220
	} else {
		goto L229
	}
L229:
	;
	goto L226
L231:
	;
	v717 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v781 = v717 + v714
	goto L96
L232:
	;
	if v773 != 0 {
		goto L95
	} else {
		goto L246
	}
L233:
	;
	v732 = v616
	goto L235
L234:
	;
	v732 = v730
	goto L235
L235:
	;
	goto L237
L236:
	;
	v773 = v768
	goto L232
L237:
	;
	if v616 == v732 {
		goto L239
	} else {
		goto L240
	}
L238:
	;
	v768 = int32(0)
	goto L236
L239:
	;
	v773 = int32(-1)
	goto L232
L240:
	;
	goto L241
L241:
	;
	v744 = int32(1)
	v745 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v747 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v745+v616))))
	if int32(249) < v747 {
		v768 = v744
		goto L236
	} else {
		goto L242
	}
L242:
	;
	v749 = v747 - int32(97)
	if v749 < int32(0) {
		v768 = v744
		goto L236
	} else {
		goto L243
	}
L243:
	;
	v755 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v749)>>(uint(int32(3))%32)))+uint32(_c_F_italian_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v755)>>(uint(v749&int32(7))%32))&int32(1) == int32(0) {
		v768 = v744
		goto L236
	} else {
		goto L244
	}
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v616 + int32(1)
	goto L245
L245:
	;
	goto L238
L246:
	;
	v774 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v775 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v775 <= v774 {
		goto L95
	} else {
		goto L247
	}
L247:
	;
	v781 = v774 + int32(1)
	goto L96
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v6
	v1013 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1013
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1013
	v1017 = v1013 - int32(1)
	if v1017 <= v6 {
		goto L311
	} else {
		goto L312
	}
L249:
	;
	if v834 < int32(0) {
		goto L248
	} else {
		goto L264
	}
L250:
	;
	v796 = v6
	goto L252
L251:
	;
	v796 = v794
	goto L252
L252:
	;
	v803 = v6
	goto L254
L253:
	;
	v834 = v814
	goto L249
L254:
	;
	if v803 == v796 {
		goto L256
	} else {
		goto L257
	}
L256:
	;
	v834 = int32(-1)
	goto L249
L257:
	;
	goto L258
L258:
	;
	v807 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v809 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v807+v803))))
	if int32(249) < v809 {
		goto L259
	} else {
		goto L260
	}
L259:
	;
	v826 = v803 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v826
	v803 = v826
	goto L254
L260:
	;
	v811 = v809 - int32(97)
	if v811 < int32(0) {
		goto L259
	} else {
		goto L261
	}
L261:
	;
	v814 = int32(1)
	v818 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v811)>>(uint(int32(3))%32)))+uint32(_c_F_italian_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v818)>>(uint(v811&int32(7))%32))&v814 != 0 {
		goto L253
	} else {
		goto L262
	}
L262:
	;
	goto L259
L264:
	;
	v837 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v838 = v837 + v834
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v838
	v849 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v849 < v838 {
		goto L266
	} else {
		goto L267
	}
L265:
	;
	if v892 < int32(0) {
		goto L248
	} else {
		goto L279
	}
L266:
	;
	v851 = v838
	goto L268
L267:
	;
	v851 = v849
	goto L268
L268:
	;
	v857 = v838
	goto L270
L269:
	;
	v892 = int32(1)
	goto L265
L270:
	;
	if v857 == v851 {
		goto L272
	} else {
		goto L273
	}
L272:
	;
	v892 = int32(-1)
	goto L265
L273:
	;
	goto L274
L274:
	;
	v864 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v866 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v864+v857))))
	if int32(249) < v866 {
		goto L269
	} else {
		goto L275
	}
L275:
	;
	v868 = v866 - int32(97)
	if v868 < int32(0) {
		goto L269
	} else {
		goto L276
	}
L276:
	;
	v874 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v868)>>(uint(int32(3))%32)))+uint32(_c_F_italian_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v874)>>(uint(v868&int32(7))%32))&int32(1) == int32(0) {
		goto L269
	} else {
		goto L277
	}
L277:
	;
	v883 = v857 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v883
	v857 = v883
	goto L270
L279:
	;
	v895 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v896 = v895 + v892
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v896
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v896
	v907 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v907 < v896 {
		goto L281
	} else {
		goto L282
	}
L280:
	;
	if v947 < int32(0) {
		goto L248
	} else {
		goto L295
	}
L281:
	;
	v909 = v896
	goto L283
L282:
	;
	v909 = v907
	goto L283
L283:
	;
	v916 = v896
	goto L285
L284:
	;
	v947 = v927
	goto L280
L285:
	;
	if v916 == v909 {
		goto L287
	} else {
		goto L288
	}
L287:
	;
	v947 = int32(-1)
	goto L280
L288:
	;
	goto L289
L289:
	;
	v920 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v922 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v920+v916))))
	if int32(249) < v922 {
		goto L290
	} else {
		goto L291
	}
L290:
	;
	v939 = v916 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v939
	v916 = v939
	goto L285
L291:
	;
	v924 = v922 - int32(97)
	if v924 < int32(0) {
		goto L290
	} else {
		goto L292
	}
L292:
	;
	v927 = int32(1)
	v931 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v924)>>(uint(int32(3))%32)))+uint32(_c_F_italian_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v931)>>(uint(v924&int32(7))%32))&v927 != 0 {
		goto L284
	} else {
		goto L293
	}
L293:
	;
	goto L290
L295:
	;
	v950 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v951 = v950 + v947
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v951
	v962 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v962 < v951 {
		goto L297
	} else {
		goto L298
	}
L296:
	;
	if v1005 < int32(0) {
		goto L248
	} else {
		goto L310
	}
L297:
	;
	v964 = v951
	goto L299
L298:
	;
	v964 = v962
	goto L299
L299:
	;
	v970 = v951
	goto L301
L300:
	;
	v1005 = int32(1)
	goto L296
L301:
	;
	if v970 == v964 {
		goto L303
	} else {
		goto L304
	}
L303:
	;
	v1005 = int32(-1)
	goto L296
L304:
	;
	goto L305
L305:
	;
	v977 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v979 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v977+v970))))
	if int32(249) < v979 {
		goto L300
	} else {
		goto L306
	}
L306:
	;
	v981 = v979 - int32(97)
	if v981 < int32(0) {
		goto L300
	} else {
		goto L307
	}
L307:
	;
	v987 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v981)>>(uint(int32(3))%32)))+uint32(_c_F_italian_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v987)>>(uint(v981&int32(7))%32))&int32(1) == int32(0) {
		goto L300
	} else {
		goto L308
	}
L308:
	;
	v996 = v970 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v996
	v970 = v996
	goto L301
L310:
	;
	v1008 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v1008 + v1005
	goto L248
L311:
	;
	v1073 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1073
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1073
	v1079 = F_find_among_b(m, l0, int32(_a_F_italian_ISO_8859_1_stem_10), int32(51), int32(0))
	mBase = m.M
	v1080 = m.ExcPending
	if v1080 != 0 {
		goto L5
	} else {
		goto L328
	}
L312:
	;
	v1019 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1021 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1019+v1017))))
	if base.B2i32(v1021&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v1021)%32)&int32(_a_F_italian_ISO_8859_1_stem_11) == int32(0)) != 0 {
		goto L311
	} else {
		goto L313
	}
L313:
	;
	v1036 = F_find_among_b(m, l0, int32(_a_F_italian_ISO_8859_1_stem_12), int32(37), int32(0))
	mBase = m.M
	v1037 = m.ExcPending
	if v1037 != 0 {
		goto L5
	} else {
		goto L314
	}
L314:
	;
	if v1036 == int32(0) {
		goto L311
	} else {
		goto L315
	}
L315:
	;
	v1040 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1040
	v1043 = v1040 - int32(1)
	v1044 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1043 <= v1044 {
		goto L311
	} else {
		goto L316
	}
L316:
	;
	v1046 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1048 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1046+v1043))))
	switch v1048 - int32(111) {
	case 0, 3:
		goto L317
	default:
		goto L311
	}
L317:
	;
	v1054 = F_find_among_b(m, l0, int32(_a_F_italian_ISO_8859_1_stem_13), int32(5), int32(0))
	mBase = m.M
	v1055 = m.ExcPending
	if v1055 != 0 {
		goto L5
	} else {
		goto L318
	}
L318:
	;
	if v1054 == int32(0) {
		goto L311
	} else {
		goto L319
	}
L319:
	;
	v1058 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v1059 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v1059 < v1058 {
		goto L311
	} else {
		goto L320
	}
L320:
	;
	switch v1054 - int32(1) {
	case 0:
		goto L322
	case 1:
		goto L321
	default:
		goto L311
	}
L321:
	;
	v1068 = F_slice_from_s(m, l0, int32(1), int32(_a_F_italian_ISO_8859_1_stem_14))
	mBase = m.M
	v1069 = m.ExcPending
	if v1069 != 0 {
		goto L5
	} else {
		goto L324
	}
L322:
	;
	v1063 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1063 {
		goto L311
	} else {
		goto L323
	}
L323:
	;
	v1575 = v1063
	goto L1
L324:
	;
	if v1068 < int32(0) {
		v1575 = v1068
		goto L1
	} else {
		goto L325
	}
L325:
	;
	goto L311
L326:
	;
	v1342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1342
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1342
	v1354 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L409
L327:
	;
	v1320 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1320
	v1322 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v1320 < v1322 {
		goto L326
	} else {
		goto L401
	}
L328:
	;
	if v1079 == int32(0) {
		goto L327
	} else {
		goto L329
	}
L329:
	;
	v1083 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1083
	switch v1079 - int32(1) {
	case 0:
		goto L338
	case 1:
		goto L337
	case 2:
		goto L336
	case 3:
		goto L335
	case 4:
		goto L334
	case 5:
		goto L333
	case 6:
		goto L332
	case 7:
		goto L331
	case 8:
		goto L330
	default:
		goto L326
	}
L330:
	;
	v1260 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1083 < v1260 {
		goto L327
	} else {
		goto L385
	}
L331:
	;
	v1221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1083 < v1221 {
		goto L327
	} else {
		goto L377
	}
L332:
	;
	v1153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1083 < v1153 {
		goto L327
	} else {
		goto L361
	}
L333:
	;
	v1148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v1083 < v1148 {
		goto L327
	} else {
		goto L359
	}
L334:
	;
	v1140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1083 < v1140 {
		goto L327
	} else {
		goto L356
	}
L335:
	;
	v1132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1083 < v1132 {
		goto L327
	} else {
		goto L353
	}
L336:
	;
	v1124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1083 < v1124 {
		goto L327
	} else {
		goto L350
	}
L337:
	;
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1083 < v1092 {
		goto L327
	} else {
		goto L341
	}
L338:
	;
	v1087 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1083 < v1087 {
		goto L327
	} else {
		goto L339
	}
L339:
	;
	v1089 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1089 {
		goto L326
	} else {
		goto L340
	}
L340:
	;
	v1575 = v1089
	goto L1
L341:
	;
	v1094 = F_slice_del(m, l0)
	mBase = m.M
	if v1094 < int32(0) {
		v1575 = v1094
		goto L1
	} else {
		goto L342
	}
L342:
	;
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1097
	v1099 = int32(2)
	v1101 = int32(0)
	v1104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1097-v1104 < v1099 {
		v1114 = v1101
		goto L344
	} else {
		goto L345
	}
L343:
	;
	if v1114 == int32(0) {
		goto L326
	} else {
		goto L347
	}
L344:
	;
	goto L343
L345:
	;
	v1107 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1110 = F_memcmp(m, v1107+v1097-v1099, int32(_a_F_italian_ISO_8859_1_stem_15), v1099)
	mBase = m.M
	if v1110 != 0 {
		v1114 = v1101
		goto L344
	} else {
		goto L346
	}
L346:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1097 - v1099
	v1114 = int32(1)
	goto L344
L347:
	;
	v1117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1117
	v1119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1117 < v1119 {
		goto L326
	} else {
		goto L348
	}
L348:
	;
	v1121 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1121 {
		goto L326
	} else {
		goto L349
	}
L349:
	;
	v1575 = v1121
	goto L1
L350:
	;
	v1128 = F_slice_from_s(m, l0, int32(3), int32(_a_F_italian_ISO_8859_1_stem_16))
	mBase = m.M
	v1129 = m.ExcPending
	if v1129 != 0 {
		goto L5
	} else {
		goto L351
	}
L351:
	;
	if int32(0) <= v1128 {
		goto L326
	} else {
		goto L352
	}
L352:
	;
	v1575 = v1128
	goto L1
L353:
	;
	v1136 = F_slice_from_s(m, l0, int32(1), int32(_a_F_italian_ISO_8859_1_stem_17))
	mBase = m.M
	v1137 = m.ExcPending
	if v1137 != 0 {
		goto L5
	} else {
		goto L354
	}
L354:
	;
	if int32(0) <= v1136 {
		goto L326
	} else {
		goto L355
	}
L355:
	;
	v1575 = v1136
	goto L1
L356:
	;
	v1144 = F_slice_from_s(m, l0, int32(4), int32(_a_F_italian_ISO_8859_1_stem_18))
	mBase = m.M
	v1145 = m.ExcPending
	if v1145 != 0 {
		goto L5
	} else {
		goto L357
	}
L357:
	;
	if int32(0) <= v1144 {
		goto L326
	} else {
		goto L358
	}
L358:
	;
	v1575 = v1144
	goto L1
L359:
	;
	v1150 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1150 {
		goto L326
	} else {
		goto L360
	}
L360:
	;
	v1575 = v1150
	goto L1
L361:
	;
	v1155 = F_slice_del(m, l0)
	mBase = m.M
	if v1155 < int32(0) {
		v1575 = v1155
		goto L1
	} else {
		goto L362
	}
L362:
	;
	v1158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1158
	v1161 = v1158 - int32(1)
	v1162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1161 <= v1162 {
		goto L326
	} else {
		goto L363
	}
L363:
	;
	v1164 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1164+v1161))))
	if base.B2i32(v1166&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v1166)%32)&int32(_a_F_italian_ISO_8859_1_stem_19) == int32(0)) != 0 {
		goto L326
	} else {
		goto L364
	}
L364:
	;
	v1181 = F_find_among_b(m, l0, int32(_a_F_italian_ISO_8859_1_stem_20), int32(4), int32(0))
	mBase = m.M
	v1182 = m.ExcPending
	if v1182 != 0 {
		goto L5
	} else {
		goto L365
	}
L365:
	;
	if v1181 == int32(0) {
		goto L326
	} else {
		goto L366
	}
L366:
	;
	v1185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1185
	v1187 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1185 < v1187 {
		goto L326
	} else {
		goto L367
	}
L367:
	;
	v1189 = F_slice_del(m, l0)
	mBase = m.M
	if v1189 < int32(0) {
		v1575 = v1189
		goto L1
	} else {
		goto L368
	}
L368:
	;
	if v1181 != int32(1) {
		goto L326
	} else {
		goto L369
	}
L369:
	;
	v1194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1194
	v1196 = int32(2)
	v1198 = int32(0)
	v1201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1194-v1201 < v1196 {
		v1211 = v1198
		goto L371
	} else {
		goto L372
	}
L370:
	;
	if v1211 == int32(0) {
		goto L326
	} else {
		goto L374
	}
L371:
	;
	goto L370
L372:
	;
	v1204 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1207 = F_memcmp(m, v1204+v1194-v1196, int32(_a_F_italian_ISO_8859_1_stem_21), v1196)
	mBase = m.M
	if v1207 != 0 {
		v1211 = v1198
		goto L371
	} else {
		goto L373
	}
L373:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1194 - v1196
	v1211 = int32(1)
	goto L371
L374:
	;
	v1214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1214
	v1216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1214 < v1216 {
		goto L326
	} else {
		goto L375
	}
L375:
	;
	v1218 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1218 {
		goto L326
	} else {
		goto L376
	}
L376:
	;
	v1575 = v1218
	goto L1
L377:
	;
	v1223 = F_slice_del(m, l0)
	mBase = m.M
	if v1223 < int32(0) {
		v1575 = v1223
		goto L1
	} else {
		goto L378
	}
L378:
	;
	v1226 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1226
	v1229 = v1226 - int32(1)
	v1230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1229 <= v1230 {
		goto L326
	} else {
		goto L379
	}
L379:
	;
	v1232 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1232+v1229))))
	if base.B2i32(v1234&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v1234)%32)&int32(_a_F_italian_ISO_8859_1_stem_22) == int32(0)) != 0 {
		goto L326
	} else {
		goto L380
	}
L380:
	;
	v1249 = F_find_among_b(m, l0, int32(_a_F_italian_ISO_8859_1_stem_23), int32(3), int32(0))
	mBase = m.M
	v1250 = m.ExcPending
	if v1250 != 0 {
		goto L5
	} else {
		goto L381
	}
L381:
	;
	if v1249 == int32(0) {
		goto L326
	} else {
		goto L382
	}
L382:
	;
	v1253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1253
	v1255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1253 < v1255 {
		goto L326
	} else {
		goto L383
	}
L383:
	;
	v1257 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1257 {
		goto L326
	} else {
		goto L384
	}
L384:
	;
	v1575 = v1257
	goto L1
L385:
	;
	v1262 = F_slice_del(m, l0)
	mBase = m.M
	if v1262 < int32(0) {
		v1575 = v1262
		goto L1
	} else {
		goto L386
	}
L386:
	;
	v1265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1265
	v1267 = int32(2)
	v1269 = int32(0)
	v1272 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1265-v1272 < v1267 {
		v1282 = v1269
		goto L388
	} else {
		goto L389
	}
L387:
	;
	if v1282 == int32(0) {
		goto L326
	} else {
		goto L391
	}
L388:
	;
	goto L387
L389:
	;
	v1275 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1278 = F_memcmp(m, v1275+v1265-v1267, int32(_a_F_italian_ISO_8859_1_stem_24), v1267)
	mBase = m.M
	if v1278 != 0 {
		v1282 = v1269
		goto L388
	} else {
		goto L390
	}
L390:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1265 - v1267
	v1282 = int32(1)
	goto L388
L391:
	;
	v1285 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1285
	v1287 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1285 < v1287 {
		goto L326
	} else {
		goto L392
	}
L392:
	;
	v1289 = F_slice_del(m, l0)
	mBase = m.M
	if v1289 < int32(0) {
		v1575 = v1289
		goto L1
	} else {
		goto L393
	}
L393:
	;
	v1292 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1292
	v1294 = int32(2)
	v1296 = int32(0)
	v1299 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1292-v1299 < v1294 {
		v1309 = v1296
		goto L395
	} else {
		goto L396
	}
L394:
	;
	if v1309 == int32(0) {
		goto L326
	} else {
		goto L398
	}
L395:
	;
	goto L394
L396:
	;
	v1302 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1305 = F_memcmp(m, v1302+v1292-v1294, int32(_a_F_italian_ISO_8859_1_stem_25), v1294)
	mBase = m.M
	if v1305 != 0 {
		v1309 = v1296
		goto L395
	} else {
		goto L397
	}
L397:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1292 - v1294
	v1309 = int32(1)
	goto L395
L398:
	;
	v1312 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1312
	v1314 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1312 < v1314 {
		goto L326
	} else {
		goto L399
	}
L399:
	;
	v1316 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1316 {
		goto L326
	} else {
		goto L400
	}
L400:
	;
	v1575 = v1316
	goto L1
L401:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1320
	v1325 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1322
	v1330 = F_find_among_b(m, l0, int32(_a_F_italian_ISO_8859_1_stem_26), int32(87), int32(0))
	mBase = m.M
	v1331 = m.ExcPending
	if v1331 != 0 {
		goto L5
	} else {
		goto L402
	}
L402:
	;
	if v1330 != 0 {
		goto L403
	} else {
		goto L404
	}
L403:
	;
	v1332 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1332
	v1334 = F_slice_del(m, l0)
	mBase = m.M
	if v1334 < int32(0) {
		v1575 = v1334
		goto L1
	} else {
		goto L406
	}
L404:
	;
	goto L405
L405:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1325
	goto L326
L406:
	;
	goto L405
L407:
	;
	if v1397 != 0 {
		goto L33
	} else {
		goto L418
	}
L408:
	;
	v1397 = v1393
	goto L407
L409:
	;
	if v1342 <= v1354 {
		goto L411
	} else {
		goto L412
	}
L410:
	;
	v1393 = int32(0)
	goto L408
L411:
	;
	v1397 = int32(-1)
	goto L407
L412:
	;
	goto L413
L413:
	;
	v1366 = int32(1)
	v1367 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1367+v1342-v1366))))
	if int32(242) < v1371 {
		v1393 = v1366
		goto L408
	} else {
		goto L414
	}
L414:
	;
	v1373 = v1371 - int32(97)
	if v1373 < int32(0) {
		v1393 = v1366
		goto L408
	} else {
		goto L415
	}
L415:
	;
	v1379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1373)>>(uint(int32(3))%32)))+uint32(_c_F_italian_ISO_8859_1_stem[1]))))
	if int32(base.Ui32(v1379)>>(uint(v1373&int32(7))%32))&int32(1) == int32(0) {
		v1393 = v1366
		goto L408
	} else {
		goto L416
	}
L416:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1342 - int32(1)
	goto L417
L417:
	;
	goto L410
L418:
	;
	v1398 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1398
	v1400 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v1398 < v1400 {
		goto L33
	} else {
		goto L419
	}
L419:
	;
	v1402 = F_slice_del(m, l0)
	mBase = m.M
	if v1402 < int32(0) {
		v1575 = v1402
		goto L1
	} else {
		goto L420
	}
L420:
	;
	v1405 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1405
	v1407 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1405 <= v1407 {
		goto L33
	} else {
		goto L421
	}
L421:
	;
	v1409 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1413 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1409+v1405-int32(1)))))
	if v1413 != int32(105) {
		goto L33
	} else {
		goto L422
	}
L422:
	;
	v1417 = v1405 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1417
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1417
	v1420 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v1405 <= v1420 {
		goto L33
	} else {
		goto L423
	}
L423:
	;
	v1422 = F_slice_del(m, l0)
	mBase = m.M
	if v1422 < int32(0) {
		v1575 = v1422
		goto L1
	} else {
		goto L424
	}
L424:
	;
	v1425 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1436 = v1425
	goto L30
L425:
	;
	if int32(0) <= v1428 {
		goto L31
	} else {
		goto L426
	}
L426:
	;
	v1575 = v1428
	goto L1
L427:
	;
	v1512 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1512
	goto L445
L428:
	;
	v1441 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1445 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1441+v1436-int32(1)))))
	if v1445 != int32(104) {
		goto L427
	} else {
		goto L429
	}
L429:
	;
	v1449 = v1436 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1449
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1449
	v1461 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L432
L430:
	;
	if v1504 != 0 {
		goto L427
	} else {
		goto L441
	}
L431:
	;
	v1504 = v1500
	goto L430
L432:
	;
	if v1449 <= v1461 {
		goto L434
	} else {
		goto L435
	}
L433:
	;
	v1500 = int32(0)
	goto L431
L434:
	;
	v1504 = int32(-1)
	goto L430
L435:
	;
	goto L436
L436:
	;
	v1473 = int32(1)
	v1474 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1478 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1474+v1449-v1473))))
	if int32(103) < v1478 {
		v1500 = v1473
		goto L431
	} else {
		goto L437
	}
L437:
	;
	v1480 = v1478 - int32(99)
	if v1480 < int32(0) {
		v1500 = v1473
		goto L431
	} else {
		goto L438
	}
L438:
	;
	v1486 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1480)>>(uint(int32(3))%32)))+uint32(_c_F_italian_ISO_8859_1_stem[2]))))
	if int32(base.Ui32(v1486)>>(uint(v1480&int32(7))%32))&int32(1) == int32(0) {
		v1500 = v1473
		goto L431
	} else {
		goto L439
	}
L439:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1449 - int32(1)
	goto L440
L440:
	;
	goto L433
L441:
	;
	v1505 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v1506 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v1506 < v1505 {
		goto L427
	} else {
		goto L442
	}
L442:
	;
	v1508 = F_slice_del(m, l0)
	mBase = m.M
	if v1508 < int32(0) {
		v1575 = v1508
		goto L1
	} else {
		goto L443
	}
L443:
	;
	goto L427
L444:
	;
	if v1567 < int32(0) {
		v1575 = v1567
		goto L1
	} else {
		goto L462
	}
L445:
	;
	v1519 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1519
	v1521 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v1521 <= v1519 {
		goto L448
	} else {
		goto L449
	}
L446:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1519
	v1567 = int32(1)
	goto L444
L447:
	;
	if v1560 < v1559 {
		goto L459
	} else {
		goto L460
	}
L448:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1519
	v1559 = v1521
	v1560 = v1519
	goto L447
L449:
	;
	v1523 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1525 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1523+v1519))))
	v1527 = v1525 - int32(73)
	v1528 = int32(0)
	if base.B2i32(v1527 == v1528)|base.B2i32(v1527 == int32(12)) == v1528 {
		goto L448
	} else {
		goto L450
	}
L450:
	;
	v1538 = F_find_among(m, l0, int32(_a_F_italian_ISO_8859_1_stem_27), int32(3), int32(0))
	mBase = m.M
	v1539 = m.ExcPending
	if v1539 != 0 {
		goto L5
	} else {
		goto L451
	}
L451:
	;
	v1540 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1540
	switch v1538 - int32(1) {
	case 0:
		goto L453
	case 1:
		goto L452
	case 2:
		goto L454
	default:
		goto L445
	}
L452:
	;
	v1553 = F_slice_from_s(m, l0, int32(1), int32(_a_F_italian_ISO_8859_1_stem_28))
	mBase = m.M
	v1554 = m.ExcPending
	if v1554 != 0 {
		goto L5
	} else {
		goto L457
	}
L453:
	;
	v1547 = F_slice_from_s(m, l0, int32(1), int32(_a_F_italian_ISO_8859_1_stem_29))
	mBase = m.M
	v1548 = m.ExcPending
	if v1548 != 0 {
		goto L5
	} else {
		goto L455
	}
L454:
	;
	v1544 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1559 = v1544
	v1560 = v1540
	goto L447
L455:
	;
	if int32(0) <= v1547 {
		goto L445
	} else {
		goto L456
	}
L456:
	;
	v1567 = v1547
	goto L444
L457:
	;
	if int32(0) <= v1553 {
		goto L445
	} else {
		goto L458
	}
L458:
	;
	v1567 = v1553
	goto L444
L459:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1560 + int32(1)
	goto L445
L460:
	;
	goto L461
L461:
	;
	goto L446
L462:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1512
	v1575 = int32(1)
	goto L1
}
