package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_InitializeTimeouts(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v24 int64
	_ = v24
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v121 int64
	_ = v121
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v159 int32
	_ = v159
	var v171 int32
	_ = v171
	var v182 int32
	_ = v182
	v1 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_InitializeTimeouts[0])) = v1
	*(*int32)(unsafe.Add(mBase, _c_F_InitializeTimeouts[1])) = v1
	v11 = v1
	for {
		v13 = int32(40)
		v14 = v11 * v13
		v17 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v14)+uint32(_c_F_InitializeTimeouts[2]))) = uint8(v17)
		*(*int32)(unsafe.Add(mBase, uint32(v14)+uint32(_c_F_InitializeTimeouts[3]))) = v11
		v24 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v14)+uint32(_c_F_InitializeTimeouts[4]))) = v24
		*(*int32)(unsafe.Add(mBase, uint32(v14)+uint32(_c_F_InitializeTimeouts[5]))) = v17
		*(*int64)(unsafe.Add(mBase, uint32(v14)+uint32(_c_F_InitializeTimeouts[6]))) = v24
		*(*int32)(unsafe.Add(mBase, uint32(v14)+uint32(_c_F_InitializeTimeouts[7]))) = v17
		*(*uint8)(unsafe.Add(mBase, uint32(v14)+uint32(_c_F_InitializeTimeouts[8]))) = uint8(v17)
		v43 = v11 | int32(1)
		v45 = v43 * v13
		*(*uint8)(unsafe.Add(mBase, uint32(v45)+uint32(_c_F_InitializeTimeouts[2]))) = uint8(v17)
		*(*int32)(unsafe.Add(mBase, uint32(v45)+uint32(_c_F_InitializeTimeouts[3]))) = v43
		*(*int64)(unsafe.Add(mBase, uint32(v45)+uint32(_c_F_InitializeTimeouts[4]))) = v24
		*(*int32)(unsafe.Add(mBase, uint32(v45)+uint32(_c_F_InitializeTimeouts[5]))) = v17
		*(*int64)(unsafe.Add(mBase, uint32(v45)+uint32(_c_F_InitializeTimeouts[6]))) = v24
		*(*int32)(unsafe.Add(mBase, uint32(v45)+uint32(_c_F_InitializeTimeouts[7]))) = v17
		*(*uint8)(unsafe.Add(mBase, uint32(v45)+uint32(_c_F_InitializeTimeouts[8]))) = uint8(v17)
		v74 = v11 | int32(2)
		v76 = v74 * v13
		*(*uint8)(unsafe.Add(mBase, uint32(v76)+uint32(_c_F_InitializeTimeouts[2]))) = uint8(v17)
		*(*int32)(unsafe.Add(mBase, uint32(v76)+uint32(_c_F_InitializeTimeouts[3]))) = v74
		*(*int64)(unsafe.Add(mBase, uint32(v76)+uint32(_c_F_InitializeTimeouts[4]))) = v24
		*(*int32)(unsafe.Add(mBase, uint32(v76)+uint32(_c_F_InitializeTimeouts[5]))) = v17
		*(*int64)(unsafe.Add(mBase, uint32(v76)+uint32(_c_F_InitializeTimeouts[6]))) = v24
		*(*int32)(unsafe.Add(mBase, uint32(v76)+uint32(_c_F_InitializeTimeouts[7]))) = v17
		*(*uint8)(unsafe.Add(mBase, uint32(v76)+uint32(_c_F_InitializeTimeouts[8]))) = uint8(v17)
		if base.B2i32(v11 == int32(20)) == v17 {
			v109 = v11 | int32(3)
			v111 = v109 * int32(40)
			v114 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v111)+uint32(_c_F_InitializeTimeouts[2]))) = uint8(v114)
			*(*int32)(unsafe.Add(mBase, uint32(v111)+uint32(_c_F_InitializeTimeouts[3]))) = v109
			v121 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v111)+uint32(_c_F_InitializeTimeouts[4]))) = v121
			*(*int32)(unsafe.Add(mBase, uint32(v111)+uint32(_c_F_InitializeTimeouts[5]))) = v114
			*(*int64)(unsafe.Add(mBase, uint32(v111)+uint32(_c_F_InitializeTimeouts[6]))) = v121
			*(*int32)(unsafe.Add(mBase, uint32(v111)+uint32(_c_F_InitializeTimeouts[7]))) = v114
			*(*uint8)(unsafe.Add(mBase, uint32(v111)+uint32(_c_F_InitializeTimeouts[8]))) = uint8(v114)
			v11 = v11 + int32(4)
			continue
		} else {
			break
		}
		break
	}
	v142 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_InitializeTimeouts[9])) = uint8(v142)
	v145 = int32(1784)
	v147 = m.G0
	v149 = v147 - int32(144)
	m.G0 = v149
	switch int32(1786) {
	case 0, 2:
		v159 = v145
	default:
		*(*int32)(unsafe.Add(mBase, _c_F_InitializeTimeouts[10])) = v145
		v159 = int32(_a_F_InitializeTimeouts_0)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v149)+4)) = v159
	F_sigemptyset(m, v149+int32(8))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v149)+136)) = int32(268435456)
	v171 = v149 + int32(4)
	if v171 != 0 {
		v182 = F___memcpy(m, int32(_a_F_InitializeTimeouts_1), v171, int32(140))
		mBase = m.M
	} else {
	}
	m.G0 = v149 + int32(144)
	return
}
func F_InputFunctionCall(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v18 int32
	_ = v18
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	v5 = int32(0)
	v7 = m.G0
	v9 = v7 + int32(-64)
	m.G0 = v9
	if l1 == v5 {
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+10)))
		if v13 != 0 {
			v77 = v5
			m.G0 = v9 - int32(-64)
			return v77
		} else {
			v14 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v9)+29)) = v14
			*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = v14
			v18 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v9)+60)) = uint8(v18)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l3
			*(*uint8)(unsafe.Add(mBase, uint32(v9)+52)) = uint8(v18)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = l2
			*(*uint8)(unsafe.Add(mBase, uint32(v9)+44)) = uint8(v18)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = l1
			v27 = int32(3)
			*(*uint16)(unsafe.Add(mBase, uint32(v9)+38)) = uint16(v27)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = l0
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v33 = m.T0[v32].(func(*base.Module, int32) int32)(m, v7+int32(-44))
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return int32(0)
			} else {
				v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+36)))
				if l1 == int32(0) {
					if v37&int32(1) != 0 {
						v77 = v33
						m.G0 = v9 - int32(-64)
						return v77
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int32(0)
						} else {
							v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v46
							F_errmsg_internal(m, int32(_a_F_InputFunctionCall_0), v9)
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_InputFunctionCall_1), int32(1554), int32(_a_F_InputFunctionCall_2))
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
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
					if v37&int32(1) == int32(0) {
						v77 = v33
						m.G0 = v9 - int32(-64)
						return v77
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v64
							F_errmsg_internal(m, int32(_a_F_InputFunctionCall_3), v7+int32(-48))
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_InputFunctionCall_1), int32(1560), int32(_a_F_InputFunctionCall_2))
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
				}
			}
		}
	} else {
		v14 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v9)+29)) = v14
		*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = v14
		v18 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v9)+60)) = uint8(v18)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l3
		*(*uint8)(unsafe.Add(mBase, uint32(v9)+52)) = uint8(v18)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = l2
		*(*uint8)(unsafe.Add(mBase, uint32(v9)+44)) = uint8(v18)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = l1
		v27 = int32(3)
		*(*uint16)(unsafe.Add(mBase, uint32(v9)+38)) = uint16(v27)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = l0
		v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v33 = m.T0[v32].(func(*base.Module, int32) int32)(m, v7+int32(-44))
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return int32(0)
		} else {
			v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+36)))
			if l1 == int32(0) {
				if v37&int32(1) != 0 {
					v77 = v33
					m.G0 = v9 - int32(-64)
					return v77
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int32(0)
					} else {
						v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v46
						F_errmsg_internal(m, int32(_a_F_InputFunctionCall_0), v9)
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_InputFunctionCall_1), int32(1554), int32(_a_F_InputFunctionCall_2))
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
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
				if v37&int32(1) == int32(0) {
					v77 = v33
					m.G0 = v9 - int32(-64)
					return v77
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return int32(0)
					} else {
						v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v64
						F_errmsg_internal(m, int32(_a_F_InputFunctionCall_3), v7+int32(-48))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_InputFunctionCall_1), int32(1560), int32(_a_F_InputFunctionCall_2))
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
			}
		}
	}
}
func F___isspace(m *base.Module, l0 int32) int32 {
	return base.B2i32(l0 == int32(32)) | base.B2i32(base.Ui32(l0-int32(9)) < base.Ui32(int32(5)))
}
func F_i2tod(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+20)))
	v4 = F_Float8GetDatum(m, base.F64_convert_i32_s(v2))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_i4tod(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_Float8GetDatum(m, base.F64_convert_i32_s(v2))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_i8tof(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	return base.I32_reinterpret_f32(base.F32_convert_i64_s(v3))
}
func F_inclusion_get_procinfo(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
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
	var v46 int32
	_ = v46
	var v47 int64
	_ = v47
	var v49 int64
	_ = v49
	var v51 int32
	_ = v51
	var v53 int64
	_ = v53
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1<<(uint(int32(2))%32)+l0)+16))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14+int32(112)))))
	if v17 != 0 {
		v59 = int32(0)
		m.G0 = v8 + int32(16)
		return v59
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
		if v18 != 0 {
			v59 = v14
			m.G0 = v8 + int32(16)
			return v59
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v20 = base.I32_extend16_s(l1)
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)+216))
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v19)+204))
			v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+6)))
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v22+v24*(v20-int32(1))<<(uint(int32(2))%32)+int32(44)-int32(4))))
			if v36 != 0 {
				v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v39 = F_index_getprocinfo(m, v37, v20, int32(11))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v46 = v14 + int32(16)
					v47 = *(*int64)(unsafe.Add(mBase, uint32(v39)+16))
					*(*int64)(unsafe.Add(mBase, uint32(v46))) = v47
					v49 = *(*int64)(unsafe.Add(mBase, uint32(v39)))
					*(*int64)(unsafe.Add(mBase, uint32(v14))) = v49
					v51 = *(*int32)(unsafe.Add(mBase, uint32(v39)+24))
					*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v51
					v53 = *(*int64)(unsafe.Add(mBase, uint32(v39)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v53
					*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v43
					*(*int32)(unsafe.Add(mBase, uint32(v46))) = int32(0)
					v59 = v14
					m.G0 = v8 + int32(16)
					return v59
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(117833860))
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(_a_F_inclusion_get_procinfo_0), int32(0))
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l1
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(11)
							F_errdetail_internal(m, int32(_a_F_inclusion_get_procinfo_1), v8)
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_inclusion_get_procinfo_2), int32(577), int32(_a_F_inclusion_get_procinfo_3))
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
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
func F_infobits_desc(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = l2
	F_appendStringInfo(m, l0, int32(_a_F_infobits_desc_0), v7)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if l1&int32(1) != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_infobits_desc_1))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	if l1&int32(2) != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L5
L7:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_infobits_desc_2))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	if l1&int32(4) != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L9
L11:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_infobits_desc_3))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	if l1&int32(8) != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	goto L13
L15:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_infobits_desc_4))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	if l1&int32(16) != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	goto L17
L19:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_infobits_desc_5))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38+v39-int32(1)))))
	if v43 == int32(32) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	goto L21
L23:
	;
	v47 = v39 - int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v47
	v50 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v38+v47))) = uint8(v50)
	goto L25
L24:
	;
	goto L25
L25:
	;
	F_appendStringInfoChar(m, l0, int32(93))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	m.G0 = v7 + int32(16)
	return
}
func F_init_MultiFuncCall(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int64
	_ = v56
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v5 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(1088))
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_init_MultiFuncCall_0), int32(0))
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_init_MultiFuncCall_1), int32(143), int32(_a_F_init_MultiFuncCall_2))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
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
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
		if v8 != int32(383) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(1088))
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_init_MultiFuncCall_0), int32(0))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_init_MultiFuncCall_1), int32(143), int32(_a_F_init_MultiFuncCall_2))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
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
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
			if v12 == int32(0) {
				v46 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
				v51 = F_AllocSetContextCreateInternal(m, v46, int32(_a_F_init_MultiFuncCall_3), int32(0), int32(1024), int32(_a_F_init_MultiFuncCall_4))
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return int32(0)
				} else {
					v54 = F_MemoryContextAllocZero(m, v51, int32(32))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						v56 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v54))) = v56
						*(*int32)(unsafe.Add(mBase, uint32(v54)+28)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v54)+24)) = v51
						*(*int64)(unsafe.Add(mBase, uint32(v54)+8)) = v56
						*(*int64)(unsafe.Add(mBase, uint32(v54)+16)) = v56
						v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v65)+16)) = v54
						v67 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
						v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						F_RegisterExprContextCallback(m, v67, int32(1628), v69)
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return int32(0)
						} else {
							return v54
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(_a_F_init_MultiFuncCall_5), int32(0))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_init_MultiFuncCall_1), int32(193), int32(_a_F_init_MultiFuncCall_2))
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
		}
	}
}
func F_initcap(m *base.Module, l0 int32) int32 {
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
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = int32(1)
		v12 = v7 + v11
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
		v17 = v15 & v11
		if v17 != 0 {
			v18 = v12
		} else {
			v18 = v7 + int32(4)
		}
		if v15 == int32(1) {
			v21 = int32(4)
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
			if v23&int32(254) == int32(2) {
				v32 = v21
			} else {
				v32 = base.B2i32(v23 == int32(18)) << (uint(v21) % 32)
			}
			if v23 == int32(1) {
				v35 = v21
			} else {
				v35 = v32
			}
			v46 = v35
		} else {
			v36 = int32(1)
			if v17 != 0 {
				v46 = int32(base.Ui32(v15)>>(uint(v36)%32)) - v36
			} else {
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
				v46 = int32(base.Ui32(v40)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v48 = F_str_initcap(m, v18, v46, v47)
		mBase = m.M
		v49 = m.ExcPending
		if v49 != 0 {
			return int32(0)
		} else {
			v50 = F_cstring_to_text(m, v48)
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return int32(0)
			} else {
				F_pfree(m, v48)
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int32(0)
				} else {
					return v50
				}
			}
		}
	}
}
func F_initialize_reloptions(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v441 int32
	_ = v441
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v544 int32
	_ = v544
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v564 int32
	_ = v564
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v579 int32
	_ = v579
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v634 int32
	_ = v634
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v647 int32
	_ = v647
	var v652 int32
	_ = v652
	var v656 int32
	_ = v656
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v667 int32
	_ = v667
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v682 int32
	_ = v682
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v703 int32
	_ = v703
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v720 int32
	_ = v720
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v746 int32
	_ = v746
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v760 int32
	_ = v760
	var v763 int32
	_ = v763
	var v766 int32
	_ = v766
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v793 int32
	_ = v793
	var v802 int32
	_ = v802
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
	var v813 int32
	_ = v813
	var v815 int32
	_ = v815
	var v818 int32
	_ = v818
	var v835 int32
	_ = v835
	v1 = int32(0)
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_initialize_reloptions[0]))
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v17 = v1
	goto L4
L2:
	;
	v36 = v1
	goto L3
L3:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_initialize_reloptions[1]))
	if v47 != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v28 = v17 + int32(1)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v28*int32(28))+uint32(_c_F_initialize_reloptions[0])))
	if v33 != 0 {
		v17 = v28
		goto L4
	} else {
		goto L6
	}
L5:
	;
	v36 = v28
	goto L3
L6:
	;
	goto L5
L7:
	;
	v49 = v1
	v50 = v36
	goto L10
L8:
	;
	v71 = v36
	goto L9
L9:
	;
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_initialize_reloptions[2]))
	if v83 != 0 {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v60 = int32(1)
	v61 = v50 + v60
	v63 = v49 + v60
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v63*int32(36))+uint32(_c_F_initialize_reloptions[1])))
	if v68 != 0 {
		v49 = v63
		v50 = v61
		goto L10
	} else {
		goto L12
	}
L11:
	;
	v71 = v61
	goto L9
L12:
	;
	goto L11
L13:
	;
	v85 = int32(0)
	v86 = v71
	goto L16
L14:
	;
	v107 = v71
	goto L15
L15:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_initialize_reloptions[3]))
	if v119 != 0 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v96 = int32(1)
	v97 = v86 + v96
	v99 = v85 + v96
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v99*int32(48))+uint32(_c_F_initialize_reloptions[2])))
	if v104 != 0 {
		v85 = v99
		v86 = v97
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v107 = v97
	goto L15
L18:
	;
	goto L17
L19:
	;
	v121 = int32(0)
	v122 = v107
	goto L22
L20:
	;
	v143 = v107
	goto L21
L21:
	;
	v155 = *(*int32)(unsafe.Add(mBase, _c_F_initialize_reloptions[4]))
	if v155 != 0 {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v132 = int32(1)
	v133 = v122 + v132
	v135 = v121 + v132
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v135*int32(36))+uint32(_c_F_initialize_reloptions[3])))
	if v140 != 0 {
		v121 = v135
		v122 = v133
		goto L22
	} else {
		goto L24
	}
L23:
	;
	v143 = v133
	goto L21
L24:
	;
	goto L23
L25:
	;
	v157 = int32(0)
	v158 = v143
	goto L28
L26:
	;
	v179 = v143
	goto L27
L27:
	;
	v189 = int32(0)
	v191 = *(*int32)(unsafe.Add(mBase, _c_F_initialize_reloptions[5]))
	v194 = *(*int32)(unsafe.Add(mBase, _c_F_initialize_reloptions[6]))
	if v194 != 0 {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	v168 = int32(1)
	v169 = v158 + v168
	v171 = v157 + v168
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v171*int32(44))+uint32(_c_F_initialize_reloptions[4])))
	if v176 != 0 {
		v157 = v171
		v158 = v169
		goto L28
	} else {
		goto L30
	}
L29:
	;
	v179 = v169
	goto L27
L30:
	;
	goto L29
L31:
	;
	F_pfree(m, v194)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	goto L33
L33:
	;
	v199 = *(*int32)(unsafe.Add(mBase, _c_F_initialize_reloptions[7]))
	v204 = F_MemoryContextAlloc(m, v199, (v191+v179)<<(uint(int32(2))%32)+int32(4))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L34
	} else {
		goto L36
	}
L34:
	;
	return
L35:
	;
	goto L33
L36:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_initialize_reloptions[6])) = v204
	v208 = *(*int32)(unsafe.Add(mBase, _c_F_initialize_reloptions[0]))
	if v208 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v210 = int32(_a_F_initialize_reloptions_0)
	v211 = v189
	goto L40
L38:
	;
	v296 = v189
	goto L39
L39:
	;
	v309 = *(*int32)(unsafe.Add(mBase, _c_F_initialize_reloptions[1]))
	if v309 != 0 {
		goto L60
	} else {
		goto L61
	}
L40:
	;
	v224 = v204 + v211<<(uint(int32(2))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v224))) = v210
	v226 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v210)+20)) = v226
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v228)))
	if v229&int32(3) == v226 {
		v253 = v229
		goto L44
	} else {
		goto L45
	}
L41:
	;
	v296 = v289
	goto L39
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v228)+16)) = v286
	v289 = v211 + int32(1)
	v291 = v289 * int32(28)
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v291)+uint32(_c_F_initialize_reloptions[0])))
	if v294 != 0 {
		v210 = v291 + int32(_a_F_initialize_reloptions_0)
		v211 = v289
		goto L40
	} else {
		goto L59
	}
L43:
	;
	v286 = v278 - v229
	goto L42
L44:
	;
	v257 = v253
	goto L53
L45:
	;
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229))))
	if v237 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v286 = int32(0)
	goto L42
L47:
	;
	goto L48
L48:
	;
	v242 = v229
	goto L49
L49:
	;
	v246 = v242 + int32(1)
	if v246&int32(3) == int32(0) {
		v253 = v246
		goto L44
	} else {
		goto L51
	}
L50:
	;
	v278 = v246
	goto L43
L51:
	;
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246))))
	if v251 != 0 {
		v242 = v246
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v257)))
	v266 = int32(-2139062144)
	if (int32(16843008)-v263|v263)&v266 == v266 {
		v257 = v257 + int32(4)
		goto L53
	} else {
		goto L55
	}
L54:
	;
	v272 = v257
	goto L56
L55:
	;
	goto L54
L56:
	;
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272))))
	if v276 != 0 {
		v272 = v272 + int32(1)
		goto L56
	} else {
		goto L58
	}
L57:
	;
	v278 = v272
	goto L43
L58:
	;
	goto L57
L59:
	;
	goto L41
L60:
	;
	v311 = int32(_a_F_initialize_reloptions_1)
	v312 = v296
	v314 = int32(0)
	goto L63
L61:
	;
	v399 = v296
	goto L62
L62:
	;
	v412 = *(*int32)(unsafe.Add(mBase, _c_F_initialize_reloptions[2]))
	if v412 != 0 {
		goto L83
	} else {
		goto L84
	}
L63:
	;
	v325 = v204 + v312<<(uint(int32(2))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v325))) = v311
	*(*int32)(unsafe.Add(mBase, uint32(v311)+20)) = int32(1)
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v325)))
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v329)))
	if v330&int32(3) == int32(0) {
		v354 = v330
		goto L67
	} else {
		goto L68
	}
L64:
	;
	v399 = v390
	goto L62
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v329)+16)) = v387
	v389 = int32(1)
	v390 = v312 + v389
	v392 = v314 + v389
	v394 = v392 * int32(36)
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v394)+uint32(_c_F_initialize_reloptions[1])))
	if v397 != 0 {
		v311 = v394 + int32(_a_F_initialize_reloptions_1)
		v312 = v390
		v314 = v392
		goto L63
	} else {
		goto L82
	}
L66:
	;
	v387 = v379 - v330
	goto L65
L67:
	;
	v358 = v354
	goto L76
L68:
	;
	v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v330))))
	if v338 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v387 = int32(0)
	goto L65
L70:
	;
	goto L71
L71:
	;
	v343 = v330
	goto L72
L72:
	;
	v347 = v343 + int32(1)
	if v347&int32(3) == int32(0) {
		v354 = v347
		goto L67
	} else {
		goto L74
	}
L73:
	;
	v379 = v347
	goto L66
L74:
	;
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v347))))
	if v352 != 0 {
		v343 = v347
		goto L72
	} else {
		goto L75
	}
L75:
	;
	goto L73
L76:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v358)))
	v367 = int32(-2139062144)
	if (int32(16843008)-v364|v364)&v367 == v367 {
		v358 = v358 + int32(4)
		goto L76
	} else {
		goto L78
	}
L77:
	;
	v373 = v358
	goto L79
L78:
	;
	goto L77
L79:
	;
	v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373))))
	if v377 != 0 {
		v373 = v373 + int32(1)
		goto L79
	} else {
		goto L81
	}
L80:
	;
	v379 = v373
	goto L66
L81:
	;
	goto L80
L82:
	;
	goto L64
L83:
	;
	v414 = int32(_a_F_initialize_reloptions_2)
	v415 = v399
	v417 = int32(0)
	goto L86
L84:
	;
	v502 = v399
	goto L85
L85:
	;
	v515 = *(*int32)(unsafe.Add(mBase, _c_F_initialize_reloptions[3]))
	if v515 != 0 {
		goto L106
	} else {
		goto L107
	}
L86:
	;
	v426 = int32(2)
	v428 = v204 + v415<<(uint(v426)%32)
	*(*int32)(unsafe.Add(mBase, uint32(v428))) = v414
	*(*int32)(unsafe.Add(mBase, uint32(v414)+20)) = v426
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v428)))
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v432)))
	if v433&int32(3) == int32(0) {
		v457 = v433
		goto L90
	} else {
		goto L91
	}
L87:
	;
	v502 = v493
	goto L85
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v432)+16)) = v490
	v492 = int32(1)
	v493 = v415 + v492
	v495 = v417 + v492
	v497 = v495 * int32(48)
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v497)+uint32(_c_F_initialize_reloptions[2])))
	if v500 != 0 {
		v414 = v497 + int32(_a_F_initialize_reloptions_2)
		v415 = v493
		v417 = v495
		goto L86
	} else {
		goto L105
	}
L89:
	;
	v490 = v482 - v433
	goto L88
L90:
	;
	v461 = v457
	goto L99
L91:
	;
	v441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v433))))
	if v441 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v490 = int32(0)
	goto L88
L93:
	;
	goto L94
L94:
	;
	v446 = v433
	goto L95
L95:
	;
	v450 = v446 + int32(1)
	if v450&int32(3) == int32(0) {
		v457 = v450
		goto L90
	} else {
		goto L97
	}
L96:
	;
	v482 = v450
	goto L89
L97:
	;
	v455 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v450))))
	if v455 != 0 {
		v446 = v450
		goto L95
	} else {
		goto L98
	}
L98:
	;
	goto L96
L99:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v461)))
	v470 = int32(-2139062144)
	if (int32(16843008)-v467|v467)&v470 == v470 {
		v461 = v461 + int32(4)
		goto L99
	} else {
		goto L101
	}
L100:
	;
	v476 = v461
	goto L102
L101:
	;
	goto L100
L102:
	;
	v480 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v476))))
	if v480 != 0 {
		v476 = v476 + int32(1)
		goto L102
	} else {
		goto L104
	}
L103:
	;
	v482 = v476
	goto L89
L104:
	;
	goto L103
L105:
	;
	goto L87
L106:
	;
	v517 = int32(_a_F_initialize_reloptions_3)
	v518 = v502
	v520 = int32(0)
	goto L109
L107:
	;
	v605 = v502
	goto L108
L108:
	;
	v618 = *(*int32)(unsafe.Add(mBase, _c_F_initialize_reloptions[4]))
	if v618 != 0 {
		goto L129
	} else {
		goto L130
	}
L109:
	;
	v531 = v204 + v518<<(uint(int32(2))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v531))) = v517
	v533 = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v517)+20)) = v533
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v531)))
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v535)))
	if v536&v533 == int32(0) {
		v560 = v536
		goto L113
	} else {
		goto L114
	}
L110:
	;
	v605 = v596
	goto L108
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v535)+16)) = v593
	v595 = int32(1)
	v596 = v518 + v595
	v598 = v520 + v595
	v600 = v598 * int32(36)
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v600)+uint32(_c_F_initialize_reloptions[3])))
	if v603 != 0 {
		v517 = v600 + int32(_a_F_initialize_reloptions_3)
		v518 = v596
		v520 = v598
		goto L109
	} else {
		goto L128
	}
L112:
	;
	v593 = v585 - v536
	goto L111
L113:
	;
	v564 = v560
	goto L122
L114:
	;
	v544 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v536))))
	if v544 == int32(0) {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v593 = int32(0)
	goto L111
L116:
	;
	goto L117
L117:
	;
	v549 = v536
	goto L118
L118:
	;
	v553 = v549 + int32(1)
	if v553&int32(3) == int32(0) {
		v560 = v553
		goto L113
	} else {
		goto L120
	}
L119:
	;
	v585 = v553
	goto L112
L120:
	;
	v558 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v553))))
	if v558 != 0 {
		v549 = v553
		goto L118
	} else {
		goto L121
	}
L121:
	;
	goto L119
L122:
	;
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v564)))
	v573 = int32(-2139062144)
	if (int32(16843008)-v570|v570)&v573 == v573 {
		v564 = v564 + int32(4)
		goto L122
	} else {
		goto L124
	}
L123:
	;
	v579 = v564
	goto L125
L124:
	;
	goto L123
L125:
	;
	v583 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v579))))
	if v583 != 0 {
		v579 = v579 + int32(1)
		goto L125
	} else {
		goto L127
	}
L126:
	;
	v585 = v579
	goto L112
L127:
	;
	goto L126
L128:
	;
	goto L110
L129:
	;
	v620 = int32(_a_F_initialize_reloptions_4)
	v621 = v605
	v623 = int32(0)
	goto L132
L130:
	;
	v708 = v605
	goto L131
L131:
	;
	v720 = *(*int32)(unsafe.Add(mBase, _c_F_initialize_reloptions[5]))
	if v720 <= int32(0) {
		v818 = v708
		goto L152
	} else {
		goto L153
	}
L132:
	;
	v634 = v204 + v621<<(uint(int32(2))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v634))) = v620
	*(*int32)(unsafe.Add(mBase, uint32(v620)+20)) = int32(4)
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v634)))
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v638)))
	if v639&int32(3) == int32(0) {
		v663 = v639
		goto L136
	} else {
		goto L137
	}
L133:
	;
	v708 = v699
	goto L131
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v638)+16)) = v696
	v698 = int32(1)
	v699 = v621 + v698
	v701 = v623 + v698
	v703 = v701 * int32(44)
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v703)+uint32(_c_F_initialize_reloptions[4])))
	if v706 != 0 {
		v620 = v703 + int32(_a_F_initialize_reloptions_4)
		v621 = v699
		v623 = v701
		goto L132
	} else {
		goto L151
	}
L135:
	;
	v696 = v688 - v639
	goto L134
L136:
	;
	v667 = v663
	goto L145
L137:
	;
	v647 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v639))))
	if v647 == int32(0) {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v696 = int32(0)
	goto L134
L139:
	;
	goto L140
L140:
	;
	v652 = v639
	goto L141
L141:
	;
	v656 = v652 + int32(1)
	if v656&int32(3) == int32(0) {
		v663 = v656
		goto L136
	} else {
		goto L143
	}
L142:
	;
	v688 = v656
	goto L135
L143:
	;
	v661 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v656))))
	if v661 != 0 {
		v652 = v656
		goto L141
	} else {
		goto L144
	}
L144:
	;
	goto L142
L145:
	;
	v673 = *(*int32)(unsafe.Add(mBase, uint32(v667)))
	v676 = int32(-2139062144)
	if (int32(16843008)-v673|v673)&v676 == v676 {
		v667 = v667 + int32(4)
		goto L145
	} else {
		goto L147
	}
L146:
	;
	v682 = v667
	goto L148
L147:
	;
	goto L146
L148:
	;
	v686 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v682))))
	if v686 != 0 {
		v682 = v682 + int32(1)
		goto L148
	} else {
		goto L150
	}
L149:
	;
	v688 = v682
	goto L135
L150:
	;
	goto L149
L151:
	;
	goto L133
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v204+v818<<(uint(int32(2))%32)))) = int32(0)
	v835 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_initialize_reloptions[8])) = uint8(v835)
	return
L153:
	;
	v724 = v720 & int32(3)
	v726 = *(*int32)(unsafe.Add(mBase, _c_F_initialize_reloptions[9]))
	if base.Ui32(v720) < base.Ui32(int32(4)) {
		goto L155
	} else {
		goto L156
	}
L154:
	;
	if v724 == int32(0) {
		v818 = v776
		goto L152
	} else {
		goto L161
	}
L155:
	;
	v775 = int32(0)
	v776 = v708
	goto L154
L156:
	;
	goto L157
L157:
	;
	v739 = int32(0)
	v740 = v708
	v746 = v1
	goto L158
L158:
	;
	v751 = int32(2)
	v752 = v740 << (uint(v751) % 32)
	v756 = v726 + v739<<(uint(v751)%32)
	v757 = *(*int32)(unsafe.Add(mBase, uint32(v756)))
	*(*int32)(unsafe.Add(mBase, uint32(v204+v752))) = v757
	v760 = *(*int32)(unsafe.Add(mBase, uint32(v756)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v752+(v204+int32(4))))) = v760
	v763 = *(*int32)(unsafe.Add(mBase, uint32(v756)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v752+(v204+int32(8))))) = v763
	v766 = *(*int32)(unsafe.Add(mBase, uint32(v756)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v752+(v204+int32(12))))) = v766
	v768 = int32(4)
	v769 = v739 + v768
	v771 = v740 + v768
	v773 = v746 + v768
	if v773 != v720&int32(2147483644) {
		v739 = v769
		v740 = v771
		v746 = v773
		goto L158
	} else {
		goto L160
	}
L159:
	;
	v775 = v769
	v776 = v771
	goto L154
L160:
	;
	goto L159
L161:
	;
	v790 = v775
	v791 = v776
	v793 = int32(0)
	goto L162
L162:
	;
	v802 = int32(2)
	v808 = *(*int32)(unsafe.Add(mBase, uint32(v726+v790<<(uint(v802)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v204+v791<<(uint(v802)%32)))) = v808
	v810 = int32(1)
	v813 = v791 + v810
	v815 = v793 + v810
	if v815 != v724 {
		v790 = v790 + v810
		v791 = v813
		v793 = v815
		goto L162
	} else {
		goto L164
	}
L163:
	;
	v818 = v813
	goto L152
L164:
	;
	goto L163
}
func F_innerrel_is_unique_ext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v272 int32
	_ = v272
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v327 int32
	_ = v327
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v358 int32
	_ = v358
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v408 int32
	_ = v408
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v431 int32
	_ = v431
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v472 int32
	_ = v472
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v495 int32
	_ = v495
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v529 int32
	_ = v529
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v552 int32
	_ = v552
	var v561 int32
	_ = v561
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v587 int32
	_ = v587
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v610 int32
	_ = v610
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v644 int32
	_ = v644
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v667 int32
	_ = v667
	var v676 int32
	_ = v676
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v705 int32
	_ = v705
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v736 int32
	_ = v736
	var v739 int32
	_ = v739
	var v741 int32
	_ = v741
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v747 int32
	_ = v747
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v765 int32
	_ = v765
	v9 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v9
	if l5 == v9 {
		v765 = v9
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v17 + int32(16)
	return v765
L2:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v25 != 0 {
		v87 = int32(0)
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if v93 == int32(0) {
		v765 = v9
		goto L1
	} else {
		goto L31
	}
L4:
	;
	v93 = v87
	goto L3
L5:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l3)+76))
	switch v26 {
	case 0:
		goto L8
	case 1:
		goto L7
	default:
		goto L6
	}
L6:
	;
	v87 = int32(0)
	goto L4
L7:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l3)+68))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v58+v59<<(uint(int32(2))%32))))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+36))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+120))
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+38)))
	if v66 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L8:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l3)+108))
	if v27 == int32(0) {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if v30 <= int32(0) {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v33 = int32(0)
	if v33 < v30 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v37 = v30
	goto L13
L12:
	;
	v37 = v33
	goto L13
L13:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v40 = v33
	goto L14
L14:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v38+v40<<(uint(int32(2))%32))))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+101)))
	if v47 != int32(1) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L6
L16:
	;
	v56 = v40 + int32(1)
	if v56 != v37 {
		v40 = v56
		goto L14
	} else {
		goto L20
	}
L17:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+103)))
	if v50 != int32(1) {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v46)+88))
	if v53 != 0 {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v93 = int32(1)
	goto L3
L20:
	;
	goto L15
L21:
	;
	v69 = int32(1)
	if v65 != 0 {
		v87 = v69
		goto L4
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	if v65 == int32(0) {
		goto L6
	} else {
		goto L30
	}
L24:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v64)+100))
	if v70 != 0 {
		v87 = v69
		goto L4
	} else {
		goto L25
	}
L25:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v64)+108))
	if v71 != 0 {
		v87 = v69
		goto L4
	} else {
		goto L26
	}
L26:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+36)))
	if v72 != 0 {
		v87 = v69
		goto L4
	} else {
		goto L27
	}
L27:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v64)+112))
	if v73 != 0 {
		v87 = v69
		goto L4
	} else {
		goto L28
	}
L28:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v64)+144))
	if v74 == int32(0) {
		goto L6
	} else {
		goto L29
	}
L29:
	;
	v87 = v69
	goto L4
L30:
	;
	v93 = int32(1)
	goto L3
L31:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l3)+176))
	if v96 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(l3)+180))
	if v257 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L33:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	if v99 <= int32(0) {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v110 = v9
	goto L35
L35:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v96)+12))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v116+v110<<(uint(int32(2))%32))))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
	if l7 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	goto L32
L37:
	;
	v240 = v110 + int32(1)
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	if v240 < v241 {
		v110 = v240
		goto L35
	} else {
		goto L70
	}
L38:
	;
	v124 = int32(0)
	if v121 == v124 {
		goto L42
	} else {
		goto L43
	}
L39:
	;
	goto L40
L40:
	;
	v181 = int32(0)
	v188 = base.B2i32(v121|l2 == v181)
	if v121 == v181 {
		v227 = v188
		goto L57
	} else {
		goto L58
	}
L41:
	;
	if v177 == int32(0) {
		goto L37
	} else {
		goto L55
	}
L42:
	;
	v177 = int32(1)
	goto L41
L43:
	;
	goto L44
L44:
	;
	if l2 == int32(0) {
		v168 = v124
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v177 = v168
	goto L41
L46:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v121)+4))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v134 < v133 {
		v168 = v124
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v136 = int32(1)
	if v133 <= v136 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v139 = v136
	goto L50
L49:
	;
	v139 = v133
	goto L50
L50:
	;
	v140 = int32(8)
	v145 = int32(0)
	goto L51
L51:
	;
	v152 = v145 << (uint(int32(2)) % 32)
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v121+v140+v152)))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v152+(l2+v140))))
	v159 = v154 & (v156 ^ int32(-1))
	v161 = base.B2i32(v159 == int32(0))
	if v159 != 0 {
		v168 = v161
		goto L45
	} else {
		goto L53
	}
L52:
	;
	v168 = v161
	goto L45
L53:
	;
	v163 = v145 + int32(1)
	if v163 != v139 {
		v145 = v163
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	v765 = int32(1)
	goto L1
L56:
	;
	if v227 == int32(0) {
		goto L37
	} else {
		goto L68
	}
L57:
	;
	goto L56
L58:
	;
	if l2 == int32(0) {
		v227 = v188
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v121)+4))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v194 != v195 {
		v227 = int32(0)
		goto L57
	} else {
		goto L60
	}
L60:
	;
	v197 = int32(1)
	if v194 <= v197 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v200 = v197
	goto L63
L62:
	;
	v200 = v194
	goto L63
L63:
	;
	v201 = int32(8)
	v206 = int32(0)
	goto L64
L64:
	;
	v214 = v206 << (uint(int32(2)) % 32)
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v121+v201+v214)))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v214+(l2+v201))))
	v219 = base.B2i32(v216 == v218)
	if v218 != v216 {
		v227 = v219
		goto L57
	} else {
		goto L66
	}
L65:
	;
	v227 = v219
	goto L57
L66:
	;
	v222 = v206 + int32(1)
	if v222 != v200 {
		v206 = v222
		goto L64
	} else {
		goto L67
	}
L67:
	;
	goto L65
L68:
	;
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+8)))
	if v233 != int32(1) {
		goto L37
	} else {
		goto L69
	}
L69:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v120)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v236
	v765 = int32(1)
	goto L1
L70:
	;
	goto L36
L71:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if int32(0) < v358 {
		goto L94
	} else {
		goto L95
	}
L72:
	;
	v260 = int32(0)
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v257)+4))
	if v261 <= v260 {
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v272 = v260
	goto L74
L74:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v257)+12))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v278+v272<<(uint(int32(2))%32))))
	v283 = int32(0)
	if l2 == v283 {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	v765 = int32(0)
	goto L1
L76:
	;
	if v336 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L77:
	;
	v336 = int32(1)
	goto L76
L78:
	;
	goto L79
L79:
	;
	if v282 == int32(0) {
		v327 = v283
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v336 = v327
	goto L76
L81:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v282)+4))
	if v293 < v292 {
		v327 = v283
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v295 = int32(1)
	if v292 <= v295 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v298 = v295
	goto L85
L84:
	;
	v298 = v292
	goto L85
L85:
	;
	v299 = int32(8)
	v304 = int32(0)
	goto L86
L86:
	;
	v311 = v304 << (uint(int32(2)) % 32)
	v313 = *(*int32)(unsafe.Add(mBase, uint32(l2+v299+v311)))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v311+(v282+v299))))
	v318 = v313 & (v315 ^ int32(-1))
	v320 = base.B2i32(v318 == int32(0))
	if v318 != 0 {
		v327 = v320
		goto L80
	} else {
		goto L88
	}
L87:
	;
	v327 = v320
	goto L80
L88:
	;
	v322 = v304 + int32(1)
	if v322 != v298 {
		v304 = v322
		goto L86
	} else {
		goto L89
	}
L89:
	;
	goto L87
L90:
	;
	v340 = v272 + int32(1)
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v257)+4))
	if v340 < v341 {
		v272 = v340
		goto L74
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	goto L75
L93:
	;
	goto L71
L94:
	;
	v375 = int32(0)
	v379 = v9
	goto L97
L95:
	;
	v705 = v9
	goto L96
L96:
	;
	if l7 != 0 {
		goto L186
	} else {
		goto L187
	}
L97:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(l5)+12))
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v380+v375<<(uint(int32(2))%32))))
	if int32(1)<<(uint(l4)%32)&int32(174) != 0 {
		goto L100
	} else {
		goto L101
	}
L98:
	;
	v705 = v687
	goto L96
L99:
	;
	v689 = v375 + int32(1)
	v690 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if v689 < v690 {
		v375 = v689
		v379 = v687
		goto L97
	} else {
		goto L185
	}
L100:
	;
	v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384)+8)))
	if v385 != 0 {
		v687 = v379
		goto L99
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	v443 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384)+9)))
	if v443 != int32(1) {
		v687 = v379
		goto L99
	} else {
		goto L119
	}
L103:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v384)+32))
	v387 = int32(0)
	if v386 == v387 {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	if v440 == int32(0) {
		v687 = v379
		goto L99
	} else {
		goto L118
	}
L105:
	;
	v440 = int32(1)
	goto L104
L106:
	;
	goto L107
L107:
	;
	if l1 == int32(0) {
		v431 = v387
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v440 = v431
	goto L104
L109:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v386)+4))
	v397 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v397 < v396 {
		v431 = v387
		goto L108
	} else {
		goto L110
	}
L110:
	;
	v399 = int32(1)
	if v396 <= v399 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v402 = v399
	goto L113
L112:
	;
	v402 = v396
	goto L113
L113:
	;
	v403 = int32(8)
	v408 = int32(0)
	goto L114
L114:
	;
	v415 = v408 << (uint(int32(2)) % 32)
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v386+v403+v415)))
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v415+(l1+v403))))
	v422 = v417 & (v419 ^ int32(-1))
	v424 = base.B2i32(v422 == int32(0))
	if v422 != 0 {
		v431 = v424
		goto L108
	} else {
		goto L116
	}
L115:
	;
	v431 = v424
	goto L108
L116:
	;
	v426 = v408 + int32(1)
	if v426 != v402 {
		v408 = v426
		goto L114
	} else {
		goto L117
	}
L117:
	;
	goto L115
L118:
	;
	goto L102
L119:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v384)+96))
	if v446 == int32(0) {
		v687 = v379
		goto L99
	} else {
		goto L120
	}
L120:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v384)+44))
	v451 = int32(0)
	if v450 == v451 {
		goto L124
	} else {
		goto L125
	}
L121:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v384)+120)) = uint8(v680)
	v682 = F_lappend(m, v379, v384)
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L183
	} else {
		goto L184
	}
L122:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v384)+44))
	v566 = int32(0)
	if v565 == v566 {
		goto L154
	} else {
		goto L155
	}
L123:
	;
	if v504 == int32(0) {
		goto L122
	} else {
		goto L137
	}
L124:
	;
	v504 = int32(1)
	goto L123
L125:
	;
	goto L126
L126:
	;
	if l2 == int32(0) {
		v495 = v451
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v504 = v495
	goto L123
L128:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v450)+4))
	v461 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v461 < v460 {
		v495 = v451
		goto L127
	} else {
		goto L129
	}
L129:
	;
	v463 = int32(1)
	if v460 <= v463 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v466 = v463
	goto L132
L131:
	;
	v466 = v460
	goto L132
L132:
	;
	v467 = int32(8)
	v472 = int32(0)
	goto L133
L133:
	;
	v479 = v472 << (uint(int32(2)) % 32)
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v450+v467+v479)))
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v479+(l2+v467))))
	v486 = v481 & (v483 ^ int32(-1))
	v488 = base.B2i32(v486 == int32(0))
	if v486 != 0 {
		v495 = v488
		goto L127
	} else {
		goto L135
	}
L134:
	;
	v495 = v488
	goto L127
L135:
	;
	v490 = v472 + int32(1)
	if v490 != v466 {
		v472 = v490
		goto L133
	} else {
		goto L136
	}
L136:
	;
	goto L134
L137:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v384)+48))
	v508 = int32(0)
	if v507 == v508 {
		goto L139
	} else {
		goto L140
	}
L138:
	;
	if v561 == int32(0) {
		goto L122
	} else {
		goto L152
	}
L139:
	;
	v561 = int32(1)
	goto L138
L140:
	;
	goto L141
L141:
	;
	if v449 == int32(0) {
		v552 = v508
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v561 = v552
	goto L138
L143:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v507)+4))
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v449)+4))
	if v518 < v517 {
		v552 = v508
		goto L142
	} else {
		goto L144
	}
L144:
	;
	v520 = int32(1)
	if v517 <= v520 {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v523 = v520
	goto L147
L146:
	;
	v523 = v517
	goto L147
L147:
	;
	v524 = int32(8)
	v529 = int32(0)
	goto L148
L148:
	;
	v536 = v529 << (uint(int32(2)) % 32)
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v507+v524+v536)))
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v536+(v449+v524))))
	v543 = v538 & (v540 ^ int32(-1))
	v545 = base.B2i32(v543 == int32(0))
	if v543 != 0 {
		v552 = v545
		goto L142
	} else {
		goto L150
	}
L149:
	;
	v552 = v545
	goto L142
L150:
	;
	v547 = v529 + int32(1)
	if v547 != v523 {
		v529 = v547
		goto L148
	} else {
		goto L151
	}
L151:
	;
	goto L149
L152:
	;
	v680 = int32(1)
	goto L121
L153:
	;
	if v619 == int32(0) {
		v687 = v379
		goto L99
	} else {
		goto L167
	}
L154:
	;
	v619 = int32(1)
	goto L153
L155:
	;
	goto L156
L156:
	;
	if v449 == int32(0) {
		v610 = v566
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v619 = v610
	goto L153
L158:
	;
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v565)+4))
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v449)+4))
	if v576 < v575 {
		v610 = v566
		goto L157
	} else {
		goto L159
	}
L159:
	;
	v578 = int32(1)
	if v575 <= v578 {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v581 = v578
	goto L162
L161:
	;
	v581 = v575
	goto L162
L162:
	;
	v582 = int32(8)
	v587 = int32(0)
	goto L163
L163:
	;
	v594 = v587 << (uint(int32(2)) % 32)
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v565+v582+v594)))
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v594+(v449+v582))))
	v601 = v596 & (v598 ^ int32(-1))
	v603 = base.B2i32(v601 == int32(0))
	if v601 != 0 {
		v610 = v603
		goto L157
	} else {
		goto L165
	}
L164:
	;
	v610 = v603
	goto L157
L165:
	;
	v605 = v587 + int32(1)
	if v605 != v581 {
		v587 = v605
		goto L163
	} else {
		goto L166
	}
L166:
	;
	goto L164
L167:
	;
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v384)+48))
	v623 = int32(0)
	if v622 == v623 {
		goto L169
	} else {
		goto L170
	}
L168:
	;
	if v676 == int32(0) {
		v687 = v379
		goto L99
	} else {
		goto L182
	}
L169:
	;
	v676 = int32(1)
	goto L168
L170:
	;
	goto L171
L171:
	;
	if l2 == int32(0) {
		v667 = v623
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v676 = v667
	goto L168
L173:
	;
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v622)+4))
	v633 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v633 < v632 {
		v667 = v623
		goto L172
	} else {
		goto L174
	}
L174:
	;
	v635 = int32(1)
	if v632 <= v635 {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v638 = v635
	goto L177
L176:
	;
	v638 = v632
	goto L177
L177:
	;
	v639 = int32(8)
	v644 = int32(0)
	goto L178
L178:
	;
	v651 = v644 << (uint(int32(2)) % 32)
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v622+v639+v651)))
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v651+(l2+v639))))
	v658 = v653 & (v655 ^ int32(-1))
	v660 = base.B2i32(v658 == int32(0))
	if v658 != 0 {
		v667 = v660
		goto L172
	} else {
		goto L180
	}
L179:
	;
	v667 = v660
	goto L172
L180:
	;
	v662 = v644 + int32(1)
	if v662 != v638 {
		v644 = v662
		goto L178
	} else {
		goto L181
	}
L181:
	;
	goto L179
L182:
	;
	v680 = int32(0)
	goto L121
L183:
	;
	return int32(0)
L184:
	;
	v687 = v682
	goto L99
L185:
	;
	goto L98
L186:
	;
	v709 = v17 + int32(12)
	goto L188
L187:
	;
	v709 = int32(0)
	goto L188
L188:
	;
	v710 = F_rel_is_distinct_for(m, l0, l3, v705, v709)
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L183
	} else {
		goto L189
	}
L189:
	;
	if v710 != 0 {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v712 = int32(_a_F_innerrel_is_unique_ext_0)
	v713 = *(*int32)(unsafe.Add(mBase, _c_F_innerrel_is_unique_ext[0]))
	v715 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
	*(*int32)(unsafe.Add(mBase, _c_F_innerrel_is_unique_ext[0])) = v715
	v718 = F_palloc0(m, int32(16))
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L183
	} else {
		goto L193
	}
L191:
	;
	goto L192
L192:
	;
	if l6 != 0 {
		goto L197
	} else {
		goto L198
	}
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v718))) = int32(329)
	v722 = F_bms_copy(m, l2)
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		goto L183
	} else {
		goto L194
	}
L194:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v718)+8)) = uint8(base.B2i32(l7 != int32(0)))
	*(*int32)(unsafe.Add(mBase, uint32(v718)+4)) = v722
	v728 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v718)+12)) = v728
	v730 = *(*int32)(unsafe.Add(mBase, uint32(l3)+176))
	v731 = F_lappend(m, v730, v718)
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L183
	} else {
		goto L195
	}
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+176)) = v731
	*(*int32)(unsafe.Add(mBase, _c_F_innerrel_is_unique_ext[0])) = v713
	v736 = int32(1)
	if l7 == int32(0) {
		v765 = v736
		goto L1
	} else {
		goto L196
	}
L196:
	;
	v739 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v739
	v765 = v736
	goto L1
L197:
	;
	v744 = int32(_a_F_innerrel_is_unique_ext_0)
	v745 = *(*int32)(unsafe.Add(mBase, _c_F_innerrel_is_unique_ext[0]))
	v747 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
	*(*int32)(unsafe.Add(mBase, _c_F_innerrel_is_unique_ext[0])) = v747
	v749 = *(*int32)(unsafe.Add(mBase, uint32(l3)+180))
	v750 = F_bms_copy(m, l2)
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L183
	} else {
		goto L200
	}
L198:
	;
	v741 = *(*int32)(unsafe.Add(mBase, uint32(l0)+368))
	if v741 != 0 {
		goto L197
	} else {
		goto L199
	}
L199:
	;
	v765 = int32(0)
	goto L1
L200:
	;
	v752 = F_lappend(m, v749, v750)
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L183
	} else {
		goto L201
	}
L201:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+180)) = v752
	*(*int32)(unsafe.Add(mBase, _c_F_innerrel_is_unique_ext[0])) = v745
	v765 = int32(0)
	goto L1
}
func F_int24div(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v3 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(33816706))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_int24div_0), int32(0))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_int24div_1), int32(1068), int32(_a_F_int24div_2))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
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
		v24 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+20)))
		v25 = base.I32_div_s(v24, v3)
		return v25
	}
}
func F_int24ne(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+20)))
	return base.B2i32(v2 != v3)
}
func F_int28div(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int64
	_ = v25
	var v26 int64
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	if v4 == int64(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(33816706))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_int28div_0), int32(0))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_int28div_1), int32(1164), int32(_a_F_int28div_2))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
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
		v25 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+20)))
		v26 = base.I64_div_s(v25, v4)
		v27 = F_Int64GetDatum(m, v26)
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			return v27
		}
	}
}
func F_int2div(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
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
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = int32(_a_F_int2div_0)
	v7 = v5 & v6
	if v7 != v6 {
		if v7 != 0 {
			v38 = base.I32_div_s(base.I32_extend16_s(v4), base.I32_extend16_s(v5))
			v41 = v38 << (uint(int32(16)) % 32)
			return v41 >> (uint(int32(16)) % 32)
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(33816706))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_int2div_1), int32(0))
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_int2div_2), int32(988), int32(_a_F_int2div_3))
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
		}
	} else {
		if v4&int32(_a_F_int2div_0) == int32(_a_F_int2div_4) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50331778))
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_int2div_5), int32(0))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_int2div_2), int32(1004), int32(_a_F_int2div_3))
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
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
			v41 = int32(0) - v4<<(uint(int32(16))%32)
			return v41 >> (uint(int32(16)) % 32)
		}
	}
}
func F_int2gt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+20)))
	v3 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+28)))
	return base.B2i32(v3 < v2)
}
func F_int2or(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	return base.I32_extend16_s(v2 | v3)
}
func F_int2shr(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+20)))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	return v2 >> (uint(v3) % 32)
}
func F_int2vectorsend(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_array_send(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_int42ne(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+28)))
	return base.B2i32(v2 != v3)
}
func F_int48div(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int64
	_ = v25
	var v26 int64
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	if v4 == int64(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(33816706))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_int48div_0), int32(0))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_int48div_1), int32(1022), int32(_a_F_int48div_2))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
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
		v25 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+20)))
		v26 = base.I64_div_s(v25, v4)
		v27 = F_Int64GetDatum(m, v26)
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			return v27
		}
	}
}
func F_int48gt(m *base.Module, l0 int32) int32 {
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
	return base.B2i32(v3 < v4)
}
func F_int4and(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	return v2 & v3
}
func F_int4div(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	switch v4 + int32(1) {
	case 0:
		if v3 == int32(-2147483648) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50331778))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_int4div_0), int32(0))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_int4div_1), int32(888), int32(_a_F_int4div_2))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
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
			return int32(0) - v3
		}
	case 1:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(33816706))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_int4div_3), int32(0))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_int4div_1), int32(872), int32(_a_F_int4div_2))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	default:
		v30 = base.I32_div_s(v3, v4)
		return v30
	}
}
func F_int4or(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	return v2 | v3
}
func F_int4range_subdiff(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = F_Float8GetDatum(m, base.F64_sub(base.F64_convert_i32_s(v2), base.F64_convert_i32_s(v4)))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_int4shr(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	return v2 >> (uint(v3) % 32)
}
func F_int64_div_fast_to_numeric(m *base.Module, l0 int64, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v14 int64
	_ = v14
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int64
	_ = v34
	var v35 int64
	_ = v35
	var v42 int64
	_ = v42
	var v44 int64
	_ = v44
	var v49 int64
	_ = v49
	var v50 int64
	_ = v50
	var v52 int64
	_ = v52
	var v55 int64
	_ = v55
	var v56 int64
	_ = v56
	var v58 int64
	_ = v58
	var v59 int64
	_ = v59
	var v63 int64
	_ = v63
	var v70 int64
	_ = v70
	var v81 int64
	_ = v81
	var v82 int64
	_ = v82
	var v87 int32
	_ = v87
	var v92 int64
	_ = v92
	var v93 int64
	_ = v93
	var v95 int64
	_ = v95
	var v98 int64
	_ = v98
	var v99 int64
	_ = v99
	var v101 int64
	_ = v101
	var v102 int64
	_ = v102
	var v106 int64
	_ = v106
	var v113 int64
	_ = v113
	var v124 int64
	_ = v124
	var v125 int64
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v152 int64
	_ = v152
	var v160 int64
	_ = v160
	var v165 int64
	_ = v165
	var v166 int64
	_ = v166
	var v171 int64
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v184 int64
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v198 int64
	_ = v198
	var v199 int64
	_ = v199
	var v205 int64
	_ = v205
	var v206 int64
	_ = v206
	var v207 int64
	_ = v207
	var v208 int64
	_ = v208
	var v213 int64
	_ = v213
	var v216 int64
	_ = v216
	var v219 int64
	_ = v219
	var v222 int64
	_ = v222
	var v223 int64
	_ = v223
	var v227 int64
	_ = v227
	var v234 int64
	_ = v234
	var v246 int32
	_ = v246
	var v247 int64
	_ = v247
	var v248 int64
	_ = v248
	var v252 int64
	_ = v252
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v300 int64
	_ = v300
	var v304 int64
	_ = v304
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v320 int64
	_ = v320
	var v325 int32
	_ = v325
	var v327 int64
	_ = v327
	var v330 int64
	_ = v330
	var v333 int32
	_ = v333
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v394 int64
	_ = v394
	var v397 int64
	_ = v397
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v415 int32
	_ = v415
	var v417 int64
	_ = v417
	var v420 int64
	_ = v420
	var v423 int32
	_ = v423
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	v3 = int32(0)
	v14 = int64(0)
	v18 = m.G0
	v20 = v18 + int32(-64)
	m.G0 = v20
	*(*int64)(unsafe.Add(mBase, uint32(v20)+56)) = v14
	*(*int64)(unsafe.Add(mBase, uint32(v20)+48)) = v14
	*(*int64)(unsafe.Add(mBase, uint32(v20)+40)) = v14
	v29 = l1 >> (uint(int32(2)) % 32)
	v31 = l1 & int32(3)
	if v31 != 0 {
		v33 = v18 + int32(-40)
		v34 = int64(63)
		v35 = l0 >> (uint(v34) % 64)
		v42 = int64(*(*int32)(unsafe.Add(mBase, uint32((int32(0)-v31)<<(uint(int32(2))%32))+uint32(_c_F_int64_div_fast_to_numeric[0]))))
		v44 = v42 >> (uint(v34) % 64)
		v49 = int64(32)
		v50 = int64(base.Ui64(v42) >> (uint(v49) % 64))
		v52 = int64(base.Ui64(l0) >> (uint(v49) % 64))
		v55 = int64(4294967295)
		v56 = v42 & v55
		v58 = l0 & v55
		v59 = v56 * v58
		v63 = int64(base.Ui64(v59)>>(uint(v49)%64)) + v56*v52
		v70 = v58*v50 + v63&v55
		*(*int64)(unsafe.Add(mBase, uint32(v33)+8)) = l0*v44 + v35*v42 + v50*v52 + int64(base.Ui64(v63)>>(uint(v49)%64)) + int64(base.Ui64(v70)>>(uint(v49)%64))
		*(*int64)(unsafe.Add(mBase, uint32(v33))) = v59&v55 | v70<<(uint(v49)%64)
		v81 = *(*int64)(unsafe.Add(mBase, uint32(v20)+32))
		v82 = *(*int64)(unsafe.Add(mBase, uint32(v20)+24))
		if v81 != v82>>(uint(int64(63))%64) {
			v87 = v18 + int32(-56)
			v92 = int64(32)
			v93 = int64(base.Ui64(l0) >> (uint(v92) % 64))
			v95 = int64(base.Ui64(v42) >> (uint(v92) % 64))
			v98 = int64(4294967295)
			v99 = l0 & v98
			v101 = v42 & v98
			v102 = v99 * v101
			v106 = int64(base.Ui64(v102)>>(uint(v92)%64)) + v99*v95
			v113 = v101*v93 + v106&v98
			*(*int64)(unsafe.Add(mBase, uint32(v87)+8)) = v42*v35 + v44*l0 + v93*v95 + int64(base.Ui64(v106)>>(uint(v92)%64)) + int64(base.Ui64(v113)>>(uint(v92)%64))
			*(*int64)(unsafe.Add(mBase, uint32(v87))) = v102&v98 | v113<<(uint(v92)%64)
			v124 = *(*int64)(unsafe.Add(mBase, uint32(v20)+8))
			v125 = *(*int64)(unsafe.Add(mBase, uint32(v20)+16))
			v126 = m.G0
			v128 = v126 - int32(32)
			m.G0 = v128
			v131 = v18 + int32(-24)
			v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+16))
			if v132 != 0 {
				F_pfree(m, v132)
				mBase = m.M
				v136 = m.ExcPending
				if v136 != 0 {
					return int32(0)
				} else {
					v138 = F_palloc(m, int32(22))
					mBase = m.M
					v139 = m.ExcPending
					if v139 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v131)+16)) = v138
						v141 = int32(0)
						*(*uint16)(unsafe.Add(mBase, uint32(v138))) = uint16(v141)
						v144 = *(*int32)(unsafe.Add(mBase, uint32(v131)+16))
						*(*int32)(unsafe.Add(mBase, uint32(v131)+20)) = v144 + int32(2)
						if v125 < int64(0) {
							*(*int64)(unsafe.Add(mBase, uint32(v131)+8)) = int64(16384)
							v152 = int64(0)
							v165 = v152 - v124
							v166 = v152 - (v125 + base.I64_extend_i32_u(base.B2i32(v124 != v152)))
							v171 = v165
							v174 = v144 + int32(22)
							v175 = v141
							v184 = v166
							for {
								v188 = int32(16)
								v189 = v128 + v188
								v192 = m.G0
								v194 = v192 - v188
								m.G0 = v194
								F___udivmodti4(m, v194, v171, v184, int64(10000), int64(0))
								mBase = m.M
								v198 = *(*int64)(unsafe.Add(mBase, uint32(v194)))
								v199 = *(*int64)(unsafe.Add(mBase, uint32(v194)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v189)+8)) = v199
								*(*int64)(unsafe.Add(mBase, uint32(v189))) = v198
								m.G0 = v194 + v188
								v205 = *(*int64)(unsafe.Add(mBase, uint32(v128)+16))
								v206 = *(*int64)(unsafe.Add(mBase, uint32(v128+int32(24))))
								v207 = int64(55536)
								v208 = int64(0)
								v213 = int64(32)
								v216 = int64(base.Ui64(v205) >> (uint(v213) % 64))
								v219 = int64(4294967295)
								v222 = v205 & v219
								v223 = v207 * v222
								v227 = int64(base.Ui64(v223)>>(uint(v213)%64)) + v207*v216
								v234 = v222*v208 + v227&v219
								*(*int64)(unsafe.Add(mBase, uint32(v128)+8)) = v205*v208 + v206*v207 + v208*v216 + int64(base.Ui64(v227)>>(uint(v213)%64)) + int64(base.Ui64(v234)>>(uint(v213)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v128))) = v223&v219 | v234<<(uint(v213)%64)
								v246 = v174 - int32(2)
								v247 = *(*int64)(unsafe.Add(mBase, uint32(v128)))
								v248 = v247 + v171
								*(*uint16)(unsafe.Add(mBase, uint32(v246))) = uint16(v248)
								v252 = int64(0)
								v257 = v175 + int32(1)
								if v184 == v252 {
									v258 = base.B2i32(base.Ui64(int64(9999)) < base.Ui64(v171))
								} else {
									v258 = base.B2i32(v184 != v252)
								}
								if v258 != 0 {
									v171 = v205
									v174 = v246
									v175 = v257
									v184 = v206
									continue
								} else {
									break
								}
								break
							}
							*(*int32)(unsafe.Add(mBase, uint32(v131)+20)) = v246
							v264 = v257
							v267 = v175
						} else {
							v160 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v131)+8)) = v160
							if v124|v125 == v160 {
								v264 = v141
								v267 = v3
							} else {
								v165 = v124
								v166 = v125
								v171 = v165
								v174 = v144 + int32(22)
								v175 = v141
								v184 = v166
								for {
									v188 = int32(16)
									v189 = v128 + v188
									v192 = m.G0
									v194 = v192 - v188
									m.G0 = v194
									F___udivmodti4(m, v194, v171, v184, int64(10000), int64(0))
									mBase = m.M
									v198 = *(*int64)(unsafe.Add(mBase, uint32(v194)))
									v199 = *(*int64)(unsafe.Add(mBase, uint32(v194)+8))
									*(*int64)(unsafe.Add(mBase, uint32(v189)+8)) = v199
									*(*int64)(unsafe.Add(mBase, uint32(v189))) = v198
									m.G0 = v194 + v188
									v205 = *(*int64)(unsafe.Add(mBase, uint32(v128)+16))
									v206 = *(*int64)(unsafe.Add(mBase, uint32(v128+int32(24))))
									v207 = int64(55536)
									v208 = int64(0)
									v213 = int64(32)
									v216 = int64(base.Ui64(v205) >> (uint(v213) % 64))
									v219 = int64(4294967295)
									v222 = v205 & v219
									v223 = v207 * v222
									v227 = int64(base.Ui64(v223)>>(uint(v213)%64)) + v207*v216
									v234 = v222*v208 + v227&v219
									*(*int64)(unsafe.Add(mBase, uint32(v128)+8)) = v205*v208 + v206*v207 + v208*v216 + int64(base.Ui64(v227)>>(uint(v213)%64)) + int64(base.Ui64(v234)>>(uint(v213)%64))
									*(*int64)(unsafe.Add(mBase, uint32(v128))) = v223&v219 | v234<<(uint(v213)%64)
									v246 = v174 - int32(2)
									v247 = *(*int64)(unsafe.Add(mBase, uint32(v128)))
									v248 = v247 + v171
									*(*uint16)(unsafe.Add(mBase, uint32(v246))) = uint16(v248)
									v252 = int64(0)
									v257 = v175 + int32(1)
									if v184 == v252 {
										v258 = base.B2i32(base.Ui64(int64(9999)) < base.Ui64(v171))
									} else {
										v258 = base.B2i32(v184 != v252)
									}
									if v258 != 0 {
										v171 = v205
										v174 = v246
										v175 = v257
										v184 = v206
										continue
									} else {
										break
									}
									break
								}
								*(*int32)(unsafe.Add(mBase, uint32(v131)+20)) = v246
								v264 = v257
								v267 = v175
							}
						}
						*(*int32)(unsafe.Add(mBase, uint32(v131)+4)) = v267
						*(*int32)(unsafe.Add(mBase, uint32(v131))) = v264
						m.G0 = v128 + int32(32)
						v282 = *(*int32)(unsafe.Add(mBase, uint32(v20)+56))
						v283 = *(*int32)(unsafe.Add(mBase, uint32(v20)+44))
						v358 = v282
						v362 = v283
						v448 = v358
						v452 = v362
						v453 = v29 + int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v20)+44)) = v452 - v453
						v464 = int32(0)
						if v464 < l1 {
							v467 = l1
						} else {
							v467 = v464
						}
						*(*int32)(unsafe.Add(mBase, uint32(v20)+52)) = v467
						v472 = F_make_result_opt_error(m, v18+int32(-24), int32(0))
						mBase = m.M
						v473 = m.ExcPending
						if v473 != 0 {
							return int32(0)
						} else {
							if v448 != 0 {
								F_pfree(m, v448)
								mBase = m.M
								v475 = m.ExcPending
								if v475 != 0 {
									return int32(0)
								} else {
									m.G0 = v20 - int32(-64)
									return v472
								}
							} else {
								m.G0 = v20 - int32(-64)
								return v472
							}
						}
					}
				}
			} else {
				v138 = F_palloc(m, int32(22))
				mBase = m.M
				v139 = m.ExcPending
				if v139 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v131)+16)) = v138
					v141 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(v138))) = uint16(v141)
					v144 = *(*int32)(unsafe.Add(mBase, uint32(v131)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v131)+20)) = v144 + int32(2)
					if v125 < int64(0) {
						*(*int64)(unsafe.Add(mBase, uint32(v131)+8)) = int64(16384)
						v152 = int64(0)
						v165 = v152 - v124
						v166 = v152 - (v125 + base.I64_extend_i32_u(base.B2i32(v124 != v152)))
						v171 = v165
						v174 = v144 + int32(22)
						v175 = v141
						v184 = v166
						for {
							v188 = int32(16)
							v189 = v128 + v188
							v192 = m.G0
							v194 = v192 - v188
							m.G0 = v194
							F___udivmodti4(m, v194, v171, v184, int64(10000), int64(0))
							mBase = m.M
							v198 = *(*int64)(unsafe.Add(mBase, uint32(v194)))
							v199 = *(*int64)(unsafe.Add(mBase, uint32(v194)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v189)+8)) = v199
							*(*int64)(unsafe.Add(mBase, uint32(v189))) = v198
							m.G0 = v194 + v188
							v205 = *(*int64)(unsafe.Add(mBase, uint32(v128)+16))
							v206 = *(*int64)(unsafe.Add(mBase, uint32(v128+int32(24))))
							v207 = int64(55536)
							v208 = int64(0)
							v213 = int64(32)
							v216 = int64(base.Ui64(v205) >> (uint(v213) % 64))
							v219 = int64(4294967295)
							v222 = v205 & v219
							v223 = v207 * v222
							v227 = int64(base.Ui64(v223)>>(uint(v213)%64)) + v207*v216
							v234 = v222*v208 + v227&v219
							*(*int64)(unsafe.Add(mBase, uint32(v128)+8)) = v205*v208 + v206*v207 + v208*v216 + int64(base.Ui64(v227)>>(uint(v213)%64)) + int64(base.Ui64(v234)>>(uint(v213)%64))
							*(*int64)(unsafe.Add(mBase, uint32(v128))) = v223&v219 | v234<<(uint(v213)%64)
							v246 = v174 - int32(2)
							v247 = *(*int64)(unsafe.Add(mBase, uint32(v128)))
							v248 = v247 + v171
							*(*uint16)(unsafe.Add(mBase, uint32(v246))) = uint16(v248)
							v252 = int64(0)
							v257 = v175 + int32(1)
							if v184 == v252 {
								v258 = base.B2i32(base.Ui64(int64(9999)) < base.Ui64(v171))
							} else {
								v258 = base.B2i32(v184 != v252)
							}
							if v258 != 0 {
								v171 = v205
								v174 = v246
								v175 = v257
								v184 = v206
								continue
							} else {
								break
							}
							break
						}
						*(*int32)(unsafe.Add(mBase, uint32(v131)+20)) = v246
						v264 = v257
						v267 = v175
					} else {
						v160 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v131)+8)) = v160
						if v124|v125 == v160 {
							v264 = v141
							v267 = v3
						} else {
							v165 = v124
							v166 = v125
							v171 = v165
							v174 = v144 + int32(22)
							v175 = v141
							v184 = v166
							for {
								v188 = int32(16)
								v189 = v128 + v188
								v192 = m.G0
								v194 = v192 - v188
								m.G0 = v194
								F___udivmodti4(m, v194, v171, v184, int64(10000), int64(0))
								mBase = m.M
								v198 = *(*int64)(unsafe.Add(mBase, uint32(v194)))
								v199 = *(*int64)(unsafe.Add(mBase, uint32(v194)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v189)+8)) = v199
								*(*int64)(unsafe.Add(mBase, uint32(v189))) = v198
								m.G0 = v194 + v188
								v205 = *(*int64)(unsafe.Add(mBase, uint32(v128)+16))
								v206 = *(*int64)(unsafe.Add(mBase, uint32(v128+int32(24))))
								v207 = int64(55536)
								v208 = int64(0)
								v213 = int64(32)
								v216 = int64(base.Ui64(v205) >> (uint(v213) % 64))
								v219 = int64(4294967295)
								v222 = v205 & v219
								v223 = v207 * v222
								v227 = int64(base.Ui64(v223)>>(uint(v213)%64)) + v207*v216
								v234 = v222*v208 + v227&v219
								*(*int64)(unsafe.Add(mBase, uint32(v128)+8)) = v205*v208 + v206*v207 + v208*v216 + int64(base.Ui64(v227)>>(uint(v213)%64)) + int64(base.Ui64(v234)>>(uint(v213)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v128))) = v223&v219 | v234<<(uint(v213)%64)
								v246 = v174 - int32(2)
								v247 = *(*int64)(unsafe.Add(mBase, uint32(v128)))
								v248 = v247 + v171
								*(*uint16)(unsafe.Add(mBase, uint32(v246))) = uint16(v248)
								v252 = int64(0)
								v257 = v175 + int32(1)
								if v184 == v252 {
									v258 = base.B2i32(base.Ui64(int64(9999)) < base.Ui64(v171))
								} else {
									v258 = base.B2i32(v184 != v252)
								}
								if v258 != 0 {
									v171 = v205
									v174 = v246
									v175 = v257
									v184 = v206
									continue
								} else {
									break
								}
								break
							}
							*(*int32)(unsafe.Add(mBase, uint32(v131)+20)) = v246
							v264 = v257
							v267 = v175
						}
					}
					*(*int32)(unsafe.Add(mBase, uint32(v131)+4)) = v267
					*(*int32)(unsafe.Add(mBase, uint32(v131))) = v264
					m.G0 = v128 + int32(32)
					v282 = *(*int32)(unsafe.Add(mBase, uint32(v20)+56))
					v283 = *(*int32)(unsafe.Add(mBase, uint32(v20)+44))
					v358 = v282
					v362 = v283
					v448 = v358
					v452 = v362
					v453 = v29 + int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v20)+44)) = v452 - v453
					v464 = int32(0)
					if v464 < l1 {
						v467 = l1
					} else {
						v467 = v464
					}
					*(*int32)(unsafe.Add(mBase, uint32(v20)+52)) = v467
					v472 = F_make_result_opt_error(m, v18+int32(-24), int32(0))
					mBase = m.M
					v473 = m.ExcPending
					if v473 != 0 {
						return int32(0)
					} else {
						if v448 != 0 {
							F_pfree(m, v448)
							mBase = m.M
							v475 = m.ExcPending
							if v475 != 0 {
								return int32(0)
							} else {
								m.G0 = v20 - int32(-64)
								return v472
							}
						} else {
							m.G0 = v20 - int32(-64)
							return v472
						}
					}
				}
			}
		} else {
			v285 = F_palloc(m, int32(12))
			mBase = m.M
			v286 = m.ExcPending
			if v286 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v20)+56)) = v285
				v288 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(v285))) = uint16(v288)
				*(*int32)(unsafe.Add(mBase, uint32(v20)+60)) = v285 + int32(2)
				if v82 < int64(0) {
					*(*int64)(unsafe.Add(mBase, uint32(v20)+48)) = int64(16384)
					v304 = int64(0) - v82
					v311 = v288
					v312 = v285 + int32(12)
					v320 = v304
					for {
						v325 = v312 - int32(2)
						v327 = base.I64_div_u_s(v320, int64(10000))
						v330 = v327*int64(55536) + v320
						*(*uint16)(unsafe.Add(mBase, uint32(v325))) = uint16(v330)
						v333 = v311 + int32(1)
						if base.Ui64(int64(9999)) < base.Ui64(v320) {
							v311 = v333
							v312 = v325
							v320 = v327
							continue
						} else {
							break
						}
						break
					}
					*(*int32)(unsafe.Add(mBase, uint32(v20)+60)) = v325
					v341 = v333
					v344 = v311
				} else {
					v300 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v20)+48)) = v300
					if v82 == v300 {
						v341 = v288
						v344 = v3
					} else {
						v304 = v82
						v311 = v288
						v312 = v285 + int32(12)
						v320 = v304
						for {
							v325 = v312 - int32(2)
							v327 = base.I64_div_u_s(v320, int64(10000))
							v330 = v327*int64(55536) + v320
							*(*uint16)(unsafe.Add(mBase, uint32(v325))) = uint16(v330)
							v333 = v311 + int32(1)
							if base.Ui64(int64(9999)) < base.Ui64(v320) {
								v311 = v333
								v312 = v325
								v320 = v327
								continue
							} else {
								break
							}
							break
						}
						*(*int32)(unsafe.Add(mBase, uint32(v20)+60)) = v325
						v341 = v333
						v344 = v311
					}
				}
				*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = v341
				v358 = v285
				v362 = v344
				v448 = v358
				v452 = v362
				v453 = v29 + int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v20)+44)) = v452 - v453
				v464 = int32(0)
				if v464 < l1 {
					v467 = l1
				} else {
					v467 = v464
				}
				*(*int32)(unsafe.Add(mBase, uint32(v20)+52)) = v467
				v472 = F_make_result_opt_error(m, v18+int32(-24), int32(0))
				mBase = m.M
				v473 = m.ExcPending
				if v473 != 0 {
					return int32(0)
				} else {
					if v448 != 0 {
						F_pfree(m, v448)
						mBase = m.M
						v475 = m.ExcPending
						if v475 != 0 {
							return int32(0)
						} else {
							m.G0 = v20 - int32(-64)
							return v472
						}
					} else {
						m.G0 = v20 - int32(-64)
						return v472
					}
				}
			}
		}
	} else {
		v375 = F_palloc(m, int32(12))
		mBase = m.M
		v376 = m.ExcPending
		if v376 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v20)+56)) = v375
			v378 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(v375))) = uint16(v378)
			*(*int32)(unsafe.Add(mBase, uint32(v20)+60)) = v375 + int32(2)
			if l0 < int64(0) {
				*(*int64)(unsafe.Add(mBase, uint32(v20)+48)) = int64(16384)
				v394 = int64(0) - l0
				v397 = v394
				v401 = v378
				v402 = v375 + int32(12)
				for {
					v415 = v402 - int32(2)
					v417 = base.I64_div_u_s(v397, int64(10000))
					v420 = v417*int64(55536) + v397
					*(*uint16)(unsafe.Add(mBase, uint32(v415))) = uint16(v420)
					v423 = v401 + int32(1)
					if base.Ui64(int64(9999)) < base.Ui64(v397) {
						v397 = v417
						v401 = v423
						v402 = v415
						continue
					} else {
						break
					}
					break
				}
				*(*int32)(unsafe.Add(mBase, uint32(v20)+60)) = v415
				v431 = v423
				v434 = v401
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = int32(0)
				if l0 == int64(0) {
					v431 = v378
					v434 = v3
				} else {
					v394 = l0
					v397 = v394
					v401 = v378
					v402 = v375 + int32(12)
					for {
						v415 = v402 - int32(2)
						v417 = base.I64_div_u_s(v397, int64(10000))
						v420 = v417*int64(55536) + v397
						*(*uint16)(unsafe.Add(mBase, uint32(v415))) = uint16(v420)
						v423 = v401 + int32(1)
						if base.Ui64(int64(9999)) < base.Ui64(v397) {
							v397 = v417
							v401 = v423
							v402 = v415
							continue
						} else {
							break
						}
						break
					}
					*(*int32)(unsafe.Add(mBase, uint32(v20)+60)) = v415
					v431 = v423
					v434 = v401
				}
			}
			*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = v431
			v448 = v375
			v452 = v434
			v453 = v29
			*(*int32)(unsafe.Add(mBase, uint32(v20)+44)) = v452 - v453
			v464 = int32(0)
			if v464 < l1 {
				v467 = l1
			} else {
				v467 = v464
			}
			*(*int32)(unsafe.Add(mBase, uint32(v20)+52)) = v467
			v472 = F_make_result_opt_error(m, v18+int32(-24), int32(0))
			mBase = m.M
			v473 = m.ExcPending
			if v473 != 0 {
				return int32(0)
			} else {
				if v448 != 0 {
					F_pfree(m, v448)
					mBase = m.M
					v475 = m.ExcPending
					if v475 != 0 {
						return int32(0)
					} else {
						m.G0 = v20 - int32(-64)
						return v472
					}
				} else {
					m.G0 = v20 - int32(-64)
					return v472
				}
			}
		}
	}
}
func F_int82le(m *base.Module, l0 int32) int32 {
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
	v4 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+28)))
	return base.B2i32(v3 <= v4)
}
func F_int8in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v4 = F_pg_strtoint64_safe(m, v2, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = F_Int64GetDatum(m, v4)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			return v8
		}
	}
}
func F_int8inc_any(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_int8inc(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_int8or(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
	v7 = F_Int64GetDatum(m, v3|v5)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_inter_sb(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_box_interpt_lseg(m, int32(0), v3, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_interpret_func_support(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v8 = F_defGetQualifiedName(m, l0)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v6)+28)) = int32(2281)
		v14 = int32(1)
		v18 = F_LookupFuncName(m, v8, v14, v6+int32(28), v14)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			if v18 != 0 {
				v20 = F_get_func_rettype(m, v18)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					if v20 != int32(2281) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(117833860))
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return int32(0)
							} else {
								v61 = F_NameListToString(m, v8)
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = int32(_a_F_interpret_func_support_0)
									*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v61
									F_errmsg(m, int32(_a_F_interpret_func_support_1), v6+int32(16))
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_interpret_func_support_2), int32(708), int32(_a_F_interpret_func_support_3))
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
						}
					} else {
						v24 = F_superuser(m)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return int32(0)
						} else {
							if v24 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v79 = m.ExcPending
								if v79 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(16797828))
									mBase = m.M
									v82 = m.ExcPending
									if v82 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(_a_F_interpret_func_support_4), int32(0))
										mBase = m.M
										v86 = m.ExcPending
										if v86 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_interpret_func_support_2), int32(718), int32(_a_F_interpret_func_support_3))
											mBase = m.M
											v91 = m.ExcPending
											if v91 != 0 {
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
								m.G0 = v6 + int32(32)
								return v18
							}
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(52461700))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						v43 = F_func_signature_string(m, v8, int32(1), int32(0), v6+int32(28))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v6))) = v43
							F_errmsg(m, int32(_a_F_interpret_func_support_5), v6)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_interpret_func_support_2), int32(702), int32(_a_F_interpret_func_support_3))
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
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
func F_intervaltypmodout(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_palloc(m, int32(64))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if v10 < int32(0) {
			v18 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v12))) = uint8(v18)
			m.G0 = v8 + int32(48)
			return v12
		} else {
			v21 = int32(base.Ui32(v10) >> (uint(int32(16)) % 32))
			if base.Ui32(v21) <= base.Ui32(int32(3071)) {
				switch v21 - int32(2) {
				case 0:
					v70 = int32(_a_F_intervaltypmodout_0)
					v71 = int32(_a_F_intervaltypmodout_1)
					v72 = v10 & v71
					if v72 != v71 {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v72
						*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v70
						v81 = F_pg_snprintf(m, v12, int32(64), int32(_a_F_intervaltypmodout_2), v8+int32(32))
						mBase = m.M
						v82 = m.ExcPending
						if v82 != 0 {
							return int32(0)
						} else {
							m.G0 = v8 + int32(48)
							return v12
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v70
						v88 = F_pg_snprintf(m, v12, int32(64), int32(_a_F_intervaltypmodout_3), v8+int32(16))
						mBase = m.M
						v89 = m.ExcPending
						if v89 != 0 {
							return int32(0)
						} else {
							m.G0 = v8 + int32(48)
							return v12
						}
					}
				case 1, 3, 5:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v10
						F_errmsg_internal(m, int32(_a_F_intervaltypmodout_4), v8)
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_intervaltypmodout_5), int32(1188), int32(_a_F_intervaltypmodout_6))
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				case 2:
					v70 = int32(_a_F_intervaltypmodout_7)
					v71 = int32(_a_F_intervaltypmodout_1)
					v72 = v10 & v71
					if v72 != v71 {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v72
						*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v70
						v81 = F_pg_snprintf(m, v12, int32(64), int32(_a_F_intervaltypmodout_2), v8+int32(32))
						mBase = m.M
						v82 = m.ExcPending
						if v82 != 0 {
							return int32(0)
						} else {
							m.G0 = v8 + int32(48)
							return v12
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v70
						v88 = F_pg_snprintf(m, v12, int32(64), int32(_a_F_intervaltypmodout_3), v8+int32(16))
						mBase = m.M
						v89 = m.ExcPending
						if v89 != 0 {
							return int32(0)
						} else {
							m.G0 = v8 + int32(48)
							return v12
						}
					}
				case 4:
					v70 = int32(_a_F_intervaltypmodout_8)
					v71 = int32(_a_F_intervaltypmodout_1)
					v72 = v10 & v71
					if v72 != v71 {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v72
						*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v70
						v81 = F_pg_snprintf(m, v12, int32(64), int32(_a_F_intervaltypmodout_2), v8+int32(32))
						mBase = m.M
						v82 = m.ExcPending
						if v82 != 0 {
							return int32(0)
						} else {
							m.G0 = v8 + int32(48)
							return v12
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v70
						v88 = F_pg_snprintf(m, v12, int32(64), int32(_a_F_intervaltypmodout_3), v8+int32(16))
						mBase = m.M
						v89 = m.ExcPending
						if v89 != 0 {
							return int32(0)
						} else {
							m.G0 = v8 + int32(48)
							return v12
						}
					}
				case 6:
					v70 = int32(_a_F_intervaltypmodout_9)
					v71 = int32(_a_F_intervaltypmodout_1)
					v72 = v10 & v71
					if v72 != v71 {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v72
						*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v70
						v81 = F_pg_snprintf(m, v12, int32(64), int32(_a_F_intervaltypmodout_2), v8+int32(32))
						mBase = m.M
						v82 = m.ExcPending
						if v82 != 0 {
							return int32(0)
						} else {
							m.G0 = v8 + int32(48)
							return v12
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v70
						v88 = F_pg_snprintf(m, v12, int32(64), int32(_a_F_intervaltypmodout_3), v8+int32(16))
						mBase = m.M
						v89 = m.ExcPending
						if v89 != 0 {
							return int32(0)
						} else {
							m.G0 = v8 + int32(48)
							return v12
						}
					}
				default:
					switch v21 - int32(1024) {
					case 0:
						v70 = int32(_a_F_intervaltypmodout_10)
						v71 = int32(_a_F_intervaltypmodout_1)
						v72 = v10 & v71
						if v72 != v71 {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v72
							*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v70
							v81 = F_pg_snprintf(m, v12, int32(64), int32(_a_F_intervaltypmodout_2), v8+int32(32))
							mBase = m.M
							v82 = m.ExcPending
							if v82 != 0 {
								return int32(0)
							} else {
								m.G0 = v8 + int32(48)
								return v12
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v70
							v88 = F_pg_snprintf(m, v12, int32(64), int32(_a_F_intervaltypmodout_3), v8+int32(16))
							mBase = m.M
							v89 = m.ExcPending
							if v89 != 0 {
								return int32(0)
							} else {
								m.G0 = v8 + int32(48)
								return v12
							}
						}
					case 1, 2, 3, 4, 5, 6, 7:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v10
							F_errmsg_internal(m, int32(_a_F_intervaltypmodout_4), v8)
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_intervaltypmodout_5), int32(1188), int32(_a_F_intervaltypmodout_6))
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					case 8:
						v70 = int32(_a_F_intervaltypmodout_11)
						v71 = int32(_a_F_intervaltypmodout_1)
						v72 = v10 & v71
						if v72 != v71 {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v72
							*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v70
							v81 = F_pg_snprintf(m, v12, int32(64), int32(_a_F_intervaltypmodout_2), v8+int32(32))
							mBase = m.M
							v82 = m.ExcPending
							if v82 != 0 {
								return int32(0)
							} else {
								m.G0 = v8 + int32(48)
								return v12
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v70
							v88 = F_pg_snprintf(m, v12, int32(64), int32(_a_F_intervaltypmodout_3), v8+int32(16))
							mBase = m.M
							v89 = m.ExcPending
							if v89 != 0 {
								return int32(0)
							} else {
								m.G0 = v8 + int32(48)
								return v12
							}
						}
					default:
						if v21 == int32(2048) {
							v70 = int32(_a_F_intervaltypmodout_12)
							v71 = int32(_a_F_intervaltypmodout_1)
							v72 = v10 & v71
							if v72 != v71 {
								*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v72
								*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v70
								v81 = F_pg_snprintf(m, v12, int32(64), int32(_a_F_intervaltypmodout_2), v8+int32(32))
								mBase = m.M
								v82 = m.ExcPending
								if v82 != 0 {
									return int32(0)
								} else {
									m.G0 = v8 + int32(48)
									return v12
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v70
								v88 = F_pg_snprintf(m, v12, int32(64), int32(_a_F_intervaltypmodout_3), v8+int32(16))
								mBase = m.M
								v89 = m.ExcPending
								if v89 != 0 {
									return int32(0)
								} else {
									m.G0 = v8 + int32(48)
									return v12
								}
							}
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = v10
								F_errmsg_internal(m, int32(_a_F_intervaltypmodout_4), v8)
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_intervaltypmodout_5), int32(1188), int32(_a_F_intervaltypmodout_6))
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
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
			} else {
				if base.Ui32(v21) <= base.Ui32(int32(_a_F_intervaltypmodout_13)) {
					switch v21 - int32(3072) {
					case 0:
						v70 = int32(_a_F_intervaltypmodout_14)
						v71 = int32(_a_F_intervaltypmodout_1)
						v72 = v10 & v71
						if v72 != v71 {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v72
							*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v70
							v81 = F_pg_snprintf(m, v12, int32(64), int32(_a_F_intervaltypmodout_2), v8+int32(32))
							mBase = m.M
							v82 = m.ExcPending
							if v82 != 0 {
								return int32(0)
							} else {
								m.G0 = v8 + int32(48)
								return v12
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v70
							v88 = F_pg_snprintf(m, v12, int32(64), int32(_a_F_intervaltypmodout_3), v8+int32(16))
							mBase = m.M
							v89 = m.ExcPending
							if v89 != 0 {
								return int32(0)
							} else {
								m.G0 = v8 + int32(48)
								return v12
							}
						}
					case 1, 2, 3, 4, 5, 6, 7:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v10
							F_errmsg_internal(m, int32(_a_F_intervaltypmodout_4), v8)
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_intervaltypmodout_5), int32(1188), int32(_a_F_intervaltypmodout_6))
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					case 8:
						v70 = int32(_a_F_intervaltypmodout_15)
						v71 = int32(_a_F_intervaltypmodout_1)
						v72 = v10 & v71
						if v72 != v71 {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v72
							*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v70
							v81 = F_pg_snprintf(m, v12, int32(64), int32(_a_F_intervaltypmodout_2), v8+int32(32))
							mBase = m.M
							v82 = m.ExcPending
							if v82 != 0 {
								return int32(0)
							} else {
								m.G0 = v8 + int32(48)
								return v12
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v70
							v88 = F_pg_snprintf(m, v12, int32(64), int32(_a_F_intervaltypmodout_3), v8+int32(16))
							mBase = m.M
							v89 = m.ExcPending
							if v89 != 0 {
								return int32(0)
							} else {
								m.G0 = v8 + int32(48)
								return v12
							}
						}
					default:
						if v21 != int32(_a_F_intervaltypmodout_16) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = v10
								F_errmsg_internal(m, int32(_a_F_intervaltypmodout_4), v8)
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_intervaltypmodout_5), int32(1188), int32(_a_F_intervaltypmodout_6))
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v70 = int32(_a_F_intervaltypmodout_17)
							v71 = int32(_a_F_intervaltypmodout_1)
							v72 = v10 & v71
							if v72 != v71 {
								*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v72
								*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v70
								v81 = F_pg_snprintf(m, v12, int32(64), int32(_a_F_intervaltypmodout_2), v8+int32(32))
								mBase = m.M
								v82 = m.ExcPending
								if v82 != 0 {
									return int32(0)
								} else {
									m.G0 = v8 + int32(48)
									return v12
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v70
								v88 = F_pg_snprintf(m, v12, int32(64), int32(_a_F_intervaltypmodout_3), v8+int32(16))
								mBase = m.M
								v89 = m.ExcPending
								if v89 != 0 {
									return int32(0)
								} else {
									m.G0 = v8 + int32(48)
									return v12
								}
							}
						}
					}
				} else {
					switch v21 - int32(_a_F_intervaltypmodout_18) {
					case 0:
						v70 = int32(_a_F_intervaltypmodout_19)
						v71 = int32(_a_F_intervaltypmodout_1)
						v72 = v10 & v71
						if v72 != v71 {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v72
							*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v70
							v81 = F_pg_snprintf(m, v12, int32(64), int32(_a_F_intervaltypmodout_2), v8+int32(32))
							mBase = m.M
							v82 = m.ExcPending
							if v82 != 0 {
								return int32(0)
							} else {
								m.G0 = v8 + int32(48)
								return v12
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v70
							v88 = F_pg_snprintf(m, v12, int32(64), int32(_a_F_intervaltypmodout_3), v8+int32(16))
							mBase = m.M
							v89 = m.ExcPending
							if v89 != 0 {
								return int32(0)
							} else {
								m.G0 = v8 + int32(48)
								return v12
							}
						}
					case 1, 2, 3, 4, 5, 6, 7:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v10
							F_errmsg_internal(m, int32(_a_F_intervaltypmodout_4), v8)
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_intervaltypmodout_5), int32(1188), int32(_a_F_intervaltypmodout_6))
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					case 8:
						v70 = int32(_a_F_intervaltypmodout_20)
						v71 = int32(_a_F_intervaltypmodout_1)
						v72 = v10 & v71
						if v72 != v71 {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v72
							*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v70
							v81 = F_pg_snprintf(m, v12, int32(64), int32(_a_F_intervaltypmodout_2), v8+int32(32))
							mBase = m.M
							v82 = m.ExcPending
							if v82 != 0 {
								return int32(0)
							} else {
								m.G0 = v8 + int32(48)
								return v12
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v70
							v88 = F_pg_snprintf(m, v12, int32(64), int32(_a_F_intervaltypmodout_3), v8+int32(16))
							mBase = m.M
							v89 = m.ExcPending
							if v89 != 0 {
								return int32(0)
							} else {
								m.G0 = v8 + int32(48)
								return v12
							}
						}
					default:
						if v21 == int32(_a_F_intervaltypmodout_21) {
							v70 = int32(_a_F_intervaltypmodout_22)
							v71 = int32(_a_F_intervaltypmodout_1)
							v72 = v10 & v71
							if v72 != v71 {
								*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v72
								*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v70
								v81 = F_pg_snprintf(m, v12, int32(64), int32(_a_F_intervaltypmodout_2), v8+int32(32))
								mBase = m.M
								v82 = m.ExcPending
								if v82 != 0 {
									return int32(0)
								} else {
									m.G0 = v8 + int32(48)
									return v12
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v70
								v88 = F_pg_snprintf(m, v12, int32(64), int32(_a_F_intervaltypmodout_3), v8+int32(16))
								mBase = m.M
								v89 = m.ExcPending
								if v89 != 0 {
									return int32(0)
								} else {
									m.G0 = v8 + int32(48)
									return v12
								}
							}
						} else {
							if v21 != int32(_a_F_intervaltypmodout_23) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v8))) = v10
									F_errmsg_internal(m, int32(_a_F_intervaltypmodout_4), v8)
									mBase = m.M
									v62 = m.ExcPending
									if v62 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_intervaltypmodout_5), int32(1188), int32(_a_F_intervaltypmodout_6))
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								v70 = int32(_a_F_intervaltypmodout_24)
								v71 = int32(_a_F_intervaltypmodout_1)
								v72 = v10 & v71
								if v72 != v71 {
									*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v72
									*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v70
									v81 = F_pg_snprintf(m, v12, int32(64), int32(_a_F_intervaltypmodout_2), v8+int32(32))
									mBase = m.M
									v82 = m.ExcPending
									if v82 != 0 {
										return int32(0)
									} else {
										m.G0 = v8 + int32(48)
										return v12
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v70
									v88 = F_pg_snprintf(m, v12, int32(64), int32(_a_F_intervaltypmodout_3), v8+int32(16))
									mBase = m.M
									v89 = m.ExcPending
									if v89 != 0 {
										return int32(0)
									} else {
										m.G0 = v8 + int32(48)
										return v12
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
func F_inv_seek(m *base.Module, l0 int32, l1 int64, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v98 int64
	_ = v98
	var v101 int64
	_ = v101
	var v105 int32
	_ = v105
	var v111 int64
	_ = v111
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v131 int64
	_ = v131
	var v133 int64
	_ = v133
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int64
	_ = v168
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	v11 = m.G0
	v13 = v11 - int32(96)
	m.G0 = v13
	switch l2 {
	case 0:
		v133 = l1
		goto L4
	case 1:
		goto L5
	case 2:
		goto L7
	default:
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L17
	} else {
		goto L51
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L17
	} else {
		goto L47
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L17
	} else {
		goto L44
	}
L4:
	;
	if base.Ui64(int64(4398046509057)) <= base.Ui64(v133) {
		goto L1
	} else {
		goto L43
	}
L5:
	;
	v131 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v133 = v131 + l1
	goto L4
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L17
	} else {
		goto L39
	}
L7:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_inv_seek[0]))
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_inv_seek[1]))
	if v19 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v20 = v16
	goto L10
L9:
	;
	v20 = int32(0)
	goto L10
L10:
	;
	if v20 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v23 = int32(_a_F_inv_seek_3)
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_inv_seek[2]))
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_inv_seek[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_inv_seek[2])) = v27
	if v16 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_ScanKeyInit(m, v13+int32(48), int32(1), int32(3), int32(184), v56)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L17
	} else {
		goto L23
	}
L14:
	;
	v39 = v19
	goto L16
L15:
	;
	v32 = F_table_open(m, int32(2613), int32(3))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	if v39 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	return int64(0)
L18:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_inv_seek[0])) = v32
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_inv_seek[1]))
	v39 = v38
	goto L16
L19:
	;
	v45 = F_index_open(m, int32(2683), int32(3))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L17
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_inv_seek[2])) = v24
	goto L13
L22:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_inv_seek[1])) = v45
	goto L21
L23:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_inv_seek[0]))
	v62 = *(*int32)(unsafe.Add(mBase, _c_F_inv_seek[1]))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v67 = F_systable_beginscan_ordered(m, v60, v62, v63, int32(1), v13+int32(48))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L17
	} else {
		goto L25
	}
L24:
	;
	F_systable_endscan_ordered(m, v67)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L17
	} else {
		goto L38
	}
L25:
	;
	v70 = F_systable_getnext_ordered(m, v67, int32(-1))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L17
	} else {
		goto L26
	}
L26:
	;
	if v70 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v111 = int64(0)
	goto L24
L28:
	;
	goto L29
L29:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v70)+16))
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+20)))
	if v75&int32(1) != 0 {
		goto L3
	} else {
		goto L30
	}
L30:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+22)))
	v79 = v74 + v78
	v81 = v79 + int32(8)
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+8)))
	v84 = v82 & int32(3)
	if v84 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v85 = F_detoast_attr(m, v81)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L17
	} else {
		goto L34
	}
L32:
	;
	v87 = v81
	goto L33
L33:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	v92 = int32(base.Ui32(v88)>>(uint(int32(2))%32)) - int32(4)
	if base.Ui32(v88-int32(_a_F_inv_seek_6)) <= base.Ui32(int32(-8197)) {
		goto L2
	} else {
		goto L35
	}
L34:
	;
	v87 = v85
	goto L33
L35:
	;
	v98 = int64(*(*int32)(unsafe.Add(mBase, uint32(v79)+4)))
	v101 = base.I64_extend_i32_s(v92) + v98<<(uint(int64(11))%64)
	if v84 == int32(0) {
		v111 = v101
		goto L24
	} else {
		goto L36
	}
L36:
	;
	F_pfree(m, v87)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L17
	} else {
		goto L37
	}
L37:
	;
	v111 = v101
	goto L24
L38:
	;
	v133 = l1 + v111
	goto L4
L39:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L17
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = l2
	F_errmsg(m, int32(_a_F_inv_seek_9), v13)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L17
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_inv_seek_1), int32(417), int32(_a_F_inv_seek_2))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L17
	} else {
		goto L42
	}
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L43:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v133
	m.G0 = v13 + int32(96)
	return v133
L44:
	;
	F_errmsg_internal(m, int32(_a_F_inv_seek_4), int32(0))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L17
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_inv_seek_1), int32(374), int32(_a_F_inv_seek_5))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L17
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
	F_errcode(m, int32(16779816))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L17
	} else {
		goto L48
	}
L48:
	;
	v168 = *(*int64)(unsafe.Add(mBase, uint32(v79)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+40)) = v92
	*(*int64)(unsafe.Add(mBase, uint32(v13)+32)) = v168
	F_errmsg(m, int32(_a_F_inv_seek_7), v13+int32(32))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L17
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(_a_F_inv_seek_1), int32(153), int32(_a_F_inv_seek_8))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L17
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L51:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L17
	} else {
		goto L52
	}
L52:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v133
	F_errmsg_internal(m, int32(_a_F_inv_seek_0), v13+int32(16))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L17
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(_a_F_inv_seek_1), int32(430), int32(_a_F_inv_seek_2))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L17
	} else {
		goto L54
	}
L54:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_is_pseudo_constant_clause_relids(m *base.Module, l0 int32, l1 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	if l1 == int32(0) {
		v7 = F_contain_volatile_functions_walker(m, l0, int32(0))
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			if v7 == int32(0) {
				v15 = int32(1)
			} else {
				v15 = int32(0)
			}
			return v15
		}
	} else {
		v15 = int32(0)
		return v15
	}
}
func F_isalpha(m *base.Module, l0 int32) int32 {
	return base.B2i32(base.Ui32(l0|int32(32)-int32(97)) < base.Ui32(int32(26)))
}
func F_isatty(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	v3 = m.G0
	v5 = v3 - int32(32)
	m.G0 = v5
	v9 = m.Wasi_snapshot_preview1.Fd_fdstat_get(m, l0, v5+int32(8))
	mBase = m.M
	if v9 == int32(0) {
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+8)))
		if v14 == int32(2) {
			v22 = int32(1)
		} else {
			v17 = int32(59)
			*(*int32)(unsafe.Add(mBase, _c_F_isatty[0])) = v17
			v22 = int32(0)
		}
	} else {
		v17 = v9
		*(*int32)(unsafe.Add(mBase, _c_F_isatty[0])) = v17
		v22 = int32(0)
	}
	m.G0 = v5 + int32(32)
	return v22
}
func F_iso_to_win866(m *base.Module, l0 int32) int32 {
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
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v8, v9, v10, int32(25), int32(20))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v22 = F_local2local(m, v6, v5, v10, int32(25), int32(20), int32(_a_F_iso_to_win866_0), base.B2i32(v7 != int32(0)))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			return v22
		}
	}
}
func F_iswalnum(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	if base.Ui32(int32(10)) <= base.Ui32(l0-int32(48)) {
		if base.Ui32(l0) <= base.Ui32(int32(_a_F_iswalnum_0)) {
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(l0)>>(uint(int32(8))%32)))+uint32(_c_F_iswalnum[0]))))
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(l0)>>(uint(int32(3))%32))&int32(31)|v18<<(uint(int32(5))%32))+uint32(_c_F_iswalnum[0]))))
			v32 = int32(base.Ui32(v24)>>(uint(l0&int32(7))%32)) & int32(1)
		} else {
			v32 = base.B2i32(base.Ui32(l0) < base.Ui32(int32(_a_F_iswalnum_1)))
		}
		v35 = base.B2i32(v32 != int32(0))
	} else {
		v35 = int32(1)
	}
	return v35
}
func F_iterate_json_values(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	v10 = m.G0
	v12 = v10 - int32(80)
	m.G0 = v12
	v15 = F_palloc0(m, int32(40))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		v18 = F_palloc0(m, int32(16))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			v22 = F_pg_detoast_datum_packed(m, l0)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				v24 = int32(1)
				v25 = v22 + v24
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
				v30 = v28 & v24
				if v30 != 0 {
					v31 = v25
				} else {
					v31 = v22 + int32(4)
				}
				if v28 == int32(1) {
					v34 = int32(4)
					v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
					if v36&int32(254) == int32(2) {
						v45 = v34
					} else {
						v45 = base.B2i32(v36 == int32(18)) << (uint(v34) % 32)
					}
					if v36 == int32(1) {
						v48 = v34
					} else {
						v48 = v45
					}
					v59 = v48
				} else {
					v49 = int32(1)
					if v30 != 0 {
						v59 = int32(base.Ui32(v28)>>(uint(v49)%32)) - v49
					} else {
						v53 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
						v59 = int32(base.Ui32(v53)>>(uint(int32(2))%32)) - int32(4)
					}
				}
				v61 = *(*int32)(unsafe.Add(mBase, _c_F_iterate_json_values[0]))
				v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
				v64 = F_makeJsonLexContextCstringLen(m, v12+int32(12), v31, v59, v62, int32(1))
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = int32(1172)
					*(*int32)(unsafe.Add(mBase, uint32(v18))) = v64
					*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = int32(1393)
					*(*int32)(unsafe.Add(mBase, uint32(v15))) = v18
					*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = int32(1394)
					v78 = F_pg_parse_json(m, v12+int32(12), v15)
					mBase = m.M
					v79 = m.ExcPending
					if v79 != 0 {
						return
					} else {
						if v78 != 0 {
							F_json_errsave_error(m, v78, v12+int32(12), int32(0))
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
								return
							} else {
								F_freeJsonLexContext(m, v12+int32(12))
								mBase = m.M
								v88 = m.ExcPending
								if v88 != 0 {
									return
								} else {
									m.G0 = v12 + int32(80)
									return
								}
							}
						} else {
							F_freeJsonLexContext(m, v12+int32(12))
							mBase = m.M
							v88 = m.ExcPending
							if v88 != 0 {
								return
							} else {
								m.G0 = v12 + int32(80)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_itmin2interval(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	var v8 int64
	_ = v8
	var v14 int32
	_ = v14
	var v16 int64
	_ = v16
	v4 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+12)))
	v5 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+16)))
	v8 = v4 + v5*int64(12)
	if base.Ui64(int64(-4294967296)) <= base.Ui64(v8-int64(2147483648)) {
		*(*uint32)(unsafe.Add(mBase, uint32(l1)+12)) = uint32(v8)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v14
		v16 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		*(*int64)(unsafe.Add(mBase, uint32(l1))) = v16
	} else {
	}
	return
}
