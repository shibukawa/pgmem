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
	*(*int32)(unsafe.Add(mBase, _consts[636])) = v1
	*(*int32)(unsafe.Add(mBase, _consts[637])) = v1
	v11 = v1
	for {
		v13 = int32(40)
		v14 = v11 * v13
		v17 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v14)+uint32(_consts[638]))) = uint8(v17)
		*(*int32)(unsafe.Add(mBase, uint32(v14)+uint32(_consts[639]))) = v11
		v24 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v14)+uint32(_consts[640]))) = v24
		*(*int32)(unsafe.Add(mBase, uint32(v14)+uint32(_consts[641]))) = v17
		*(*int64)(unsafe.Add(mBase, uint32(v14)+uint32(_consts[642]))) = v24
		*(*int32)(unsafe.Add(mBase, uint32(v14)+uint32(_consts[643]))) = v17
		*(*uint8)(unsafe.Add(mBase, uint32(v14)+uint32(_consts[644]))) = uint8(v17)
		v43 = v11 | int32(1)
		v45 = v43 * v13
		*(*uint8)(unsafe.Add(mBase, uint32(v45)+uint32(_consts[638]))) = uint8(v17)
		*(*int32)(unsafe.Add(mBase, uint32(v45)+uint32(_consts[639]))) = v43
		*(*int64)(unsafe.Add(mBase, uint32(v45)+uint32(_consts[640]))) = v24
		*(*int32)(unsafe.Add(mBase, uint32(v45)+uint32(_consts[641]))) = v17
		*(*int64)(unsafe.Add(mBase, uint32(v45)+uint32(_consts[642]))) = v24
		*(*int32)(unsafe.Add(mBase, uint32(v45)+uint32(_consts[643]))) = v17
		*(*uint8)(unsafe.Add(mBase, uint32(v45)+uint32(_consts[644]))) = uint8(v17)
		v74 = v11 | int32(2)
		v76 = v74 * v13
		*(*uint8)(unsafe.Add(mBase, uint32(v76)+uint32(_consts[638]))) = uint8(v17)
		*(*int32)(unsafe.Add(mBase, uint32(v76)+uint32(_consts[639]))) = v74
		*(*int64)(unsafe.Add(mBase, uint32(v76)+uint32(_consts[640]))) = v24
		*(*int32)(unsafe.Add(mBase, uint32(v76)+uint32(_consts[641]))) = v17
		*(*int64)(unsafe.Add(mBase, uint32(v76)+uint32(_consts[642]))) = v24
		*(*int32)(unsafe.Add(mBase, uint32(v76)+uint32(_consts[643]))) = v17
		*(*uint8)(unsafe.Add(mBase, uint32(v76)+uint32(_consts[644]))) = uint8(v17)
		if base.B2i32(v11 == int32(20)) == v17 {
			v109 = v11 | int32(3)
			v111 = v109 * int32(40)
			v114 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v111)+uint32(_consts[638]))) = uint8(v114)
			*(*int32)(unsafe.Add(mBase, uint32(v111)+uint32(_consts[639]))) = v109
			v121 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v111)+uint32(_consts[640]))) = v121
			*(*int32)(unsafe.Add(mBase, uint32(v111)+uint32(_consts[641]))) = v114
			*(*int64)(unsafe.Add(mBase, uint32(v111)+uint32(_consts[642]))) = v121
			*(*int32)(unsafe.Add(mBase, uint32(v111)+uint32(_consts[643]))) = v114
			*(*uint8)(unsafe.Add(mBase, uint32(v111)+uint32(_consts[644]))) = uint8(v114)
			v11 = v11 + int32(4)
			continue
		} else {
			break
		}
		break
	}
	v142 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[645])) = uint8(v142)
	v145 = int32(1784)
	v147 = m.G0
	v149 = v147 - int32(144)
	m.G0 = v149
	switch int32(1786) {
	case 0, 2:
		v159 = v145
	default:
		*(*int32)(unsafe.Add(mBase, _consts[724])) = v145
		v159 = int32(4729)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v149)+4)) = v159
	F_sigemptyset(m, v149+int32(8))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v149)+136)) = int32(268435456)
	v171 = v149 + int32(4)
	if v171 != 0 {
		v182 = F___memcpy(m, int32(4609864), v171, int32(140))
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
							F_errmsg_internal(m, int32(508571), v9)
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(472279), int32(1554), int32(290259))
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
							F_errmsg_internal(m, int32(508692), v7+int32(-48))
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(472279), int32(1560), int32(290259))
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
						F_errmsg_internal(m, int32(508571), v9)
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(472279), int32(1554), int32(290259))
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
						F_errmsg_internal(m, int32(508692), v7+int32(-48))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(472279), int32(1560), int32(290259))
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
						F_errmsg_internal(m, int32(237838), int32(0))
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l1
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(11)
							F_errdetail_internal(m, int32(612902), v8)
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(473405), int32(577), int32(229272))
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
	F_appendStringInfo(m, l0, int32(484327), v7)
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
	F_appendStringInfoString(m, l0, int32(704177))
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
	F_appendStringInfoString(m, l0, int32(704124))
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
	F_appendStringInfoString(m, l0, int32(704165))
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
	F_appendStringInfoString(m, l0, int32(704151))
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
	F_appendStringInfoString(m, l0, int32(704198))
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
				F_errmsg(m, int32(99381), int32(0))
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(474666), int32(143), int32(290344))
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
					F_errmsg(m, int32(99381), int32(0))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(474666), int32(143), int32(290344))
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
				v51 = F_AllocSetContextCreateInternal(m, v46, int32(58233), int32(0), int32(1024), int32(8192))
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
					F_errmsg_internal(m, int32(395759), int32(0))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(474666), int32(193), int32(290344))
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
func F_init_work(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int64
	_ = v18
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
	var v40 int32
	_ = v40
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v105 int32
	_ = v105
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v155 int32
	_ = v155
	var v164 int32
	_ = v164
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v204 int32
	_ = v204
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v240 int32
	_ = v240
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v270 int32
	_ = v270
	var v281 int32
	_ = v281
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v361 int32
	_ = v361
	var v368 int32
	_ = v368
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v445 int32
	_ = v445
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v525 int32
	_ = v525
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v574 int32
	_ = v574
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v605 int32
	_ = v605
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
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
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v640 int32
	_ = v640
	var v643 int32
	_ = v643
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v675 int32
	_ = v675
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v690 int32
	_ = v690
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v711 int32
	_ = v711
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v725 int32
	_ = v725
	var v728 int32
	_ = v728
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v741 int32
	_ = v741
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v763 int32
	_ = v763
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v775 int32
	_ = v775
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v797 int32
	_ = v797
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v809 int32
	_ = v809
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v831 int32
	_ = v831
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v846 int32
	_ = v846
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v867 int32
	_ = v867
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v874 int32
	_ = v874
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v881 int32
	_ = v881
	var v884 int32
	_ = v884
	var v890 int32
	_ = v890
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v914 int32
	_ = v914
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v929 int32
	_ = v929
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v950 int32
	_ = v950
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v957 int32
	_ = v957
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v964 int32
	_ = v964
	var v967 int32
	_ = v967
	var v973 int32
	_ = v973
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v997 int32
	_ = v997
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1012 int32
	_ = v1012
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1033 int32
	_ = v1033
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1040 int32
	_ = v1040
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1047 int32
	_ = v1047
	var v1050 int32
	_ = v1050
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1061 int32
	_ = v1061
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1077 int32
	_ = v1077
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1090 int32
	_ = v1090
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1111 int32
	_ = v1111
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1118 int32
	_ = v1118
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1125 int32
	_ = v1125
	var v1128 int32
	_ = v1128
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1139 int32
	_ = v1139
	var v1142 int32
	_ = v1142
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1154 int32
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1158 int32
	_ = v1158
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1173 int32
	_ = v1173
	var v1178 int32
	_ = v1178
	var v1179 int32
	_ = v1179
	var v1180 int32
	_ = v1180
	var v1181 int32
	_ = v1181
	var v1187 int32
	_ = v1187
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1192 int32
	_ = v1192
	var v1194 int32
	_ = v1194
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1199 int32
	_ = v1199
	var v1201 int32
	_ = v1201
	var v1203 int32
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1208 int32
	_ = v1208
	var v1211 int32
	_ = v1211
	var v1217 int32
	_ = v1217
	var v1220 int32
	_ = v1220
	var v1223 int32
	_ = v1223
	var v1224 int32
	_ = v1224
	var v1228 int32
	_ = v1228
	var v1229 int32
	_ = v1229
	var v1232 int32
	_ = v1232
	var v1233 int32
	_ = v1233
	var v1236 int32
	_ = v1236
	var v1243 int32
	_ = v1243
	var v1244 int32
	_ = v1244
	var v1255 int32
	_ = v1255
	var v1263 int32
	_ = v1263
	var v1271 int32
	_ = v1271
	var v1279 int32
	_ = v1279
	var v1287 int32
	_ = v1287
	var v1295 int32
	_ = v1295
	var v1303 int32
	_ = v1303
	var v1311 int32
	_ = v1311
	var v1320 int32
	_ = v1320
	var v1324 int32
	_ = v1324
	var v1325 int32
	_ = v1325
	var v1326 int32
	_ = v1326
	var v1329 int32
	_ = v1329
	var v1332 int32
	_ = v1332
	var v1333 int32
	_ = v1333
	var v1337 int32
	_ = v1337
	var v1338 int32
	_ = v1338
	var v1341 int32
	_ = v1341
	var v1342 int32
	_ = v1342
	var v1345 int32
	_ = v1345
	var v1352 int32
	_ = v1352
	var v1353 int32
	_ = v1353
	var v1362 int32
	_ = v1362
	var v1367 int32
	_ = v1367
	var v1368 int32
	_ = v1368
	var v1369 int32
	_ = v1369
	var v1370 int32
	_ = v1370
	var v1376 int32
	_ = v1376
	var v1377 int32
	_ = v1377
	var v1378 int32
	_ = v1378
	var v1379 int32
	_ = v1379
	var v1380 int32
	_ = v1380
	var v1381 int32
	_ = v1381
	var v1383 int32
	_ = v1383
	var v1386 int32
	_ = v1386
	var v1387 int32
	_ = v1387
	var v1388 int32
	_ = v1388
	var v1390 int32
	_ = v1390
	var v1392 int32
	_ = v1392
	var v1393 int32
	_ = v1393
	var v1397 int32
	_ = v1397
	var v1400 int32
	_ = v1400
	var v1406 int32
	_ = v1406
	var v1409 int32
	_ = v1409
	var v1412 int32
	_ = v1412
	var v1413 int32
	_ = v1413
	var v1417 int32
	_ = v1417
	var v1418 int32
	_ = v1418
	var v1421 int32
	_ = v1421
	var v1422 int32
	_ = v1422
	var v1425 int32
	_ = v1425
	var v1432 int32
	_ = v1432
	var v1433 int32
	_ = v1433
	var v1442 int32
	_ = v1442
	var v1447 int32
	_ = v1447
	var v1448 int32
	_ = v1448
	var v1449 int32
	_ = v1449
	var v1450 int32
	_ = v1450
	var v1456 int32
	_ = v1456
	var v1457 int32
	_ = v1457
	var v1458 int32
	_ = v1458
	var v1459 int32
	_ = v1459
	var v1460 int32
	_ = v1460
	var v1461 int32
	_ = v1461
	var v1463 int32
	_ = v1463
	var v1466 int32
	_ = v1466
	var v1467 int32
	_ = v1467
	var v1468 int32
	_ = v1468
	var v1470 int32
	_ = v1470
	var v1472 int32
	_ = v1472
	var v1473 int32
	_ = v1473
	var v1477 int32
	_ = v1477
	var v1480 int32
	_ = v1480
	var v1486 int32
	_ = v1486
	var v1489 int32
	_ = v1489
	var v1492 int32
	_ = v1492
	var v1493 int32
	_ = v1493
	var v1497 int32
	_ = v1497
	var v1498 int32
	_ = v1498
	var v1501 int32
	_ = v1501
	var v1502 int32
	_ = v1502
	var v1505 int32
	_ = v1505
	var v1512 int32
	_ = v1512
	var v1513 int32
	_ = v1513
	var v1522 int32
	_ = v1522
	var v1527 int32
	_ = v1527
	var v1528 int32
	_ = v1528
	var v1529 int32
	_ = v1529
	var v1530 int32
	_ = v1530
	var v1536 int32
	_ = v1536
	var v1537 int32
	_ = v1537
	var v1538 int32
	_ = v1538
	var v1539 int32
	_ = v1539
	var v1540 int32
	_ = v1540
	var v1541 int32
	_ = v1541
	var v1543 int32
	_ = v1543
	var v1546 int32
	_ = v1546
	var v1547 int32
	_ = v1547
	var v1548 int32
	_ = v1548
	var v1550 int32
	_ = v1550
	var v1552 int32
	_ = v1552
	var v1553 int32
	_ = v1553
	var v1557 int32
	_ = v1557
	var v1560 int32
	_ = v1560
	var v1566 int32
	_ = v1566
	var v1569 int32
	_ = v1569
	var v1572 int32
	_ = v1572
	var v1573 int32
	_ = v1573
	var v1577 int32
	_ = v1577
	var v1578 int32
	_ = v1578
	var v1581 int32
	_ = v1581
	var v1582 int32
	_ = v1582
	var v1585 int32
	_ = v1585
	var v1592 int32
	_ = v1592
	var v1593 int32
	_ = v1593
	var v1602 int32
	_ = v1602
	var v1607 int32
	_ = v1607
	var v1608 int32
	_ = v1608
	var v1609 int32
	_ = v1609
	var v1610 int32
	_ = v1610
	var v1616 int32
	_ = v1616
	var v1617 int32
	_ = v1617
	var v1618 int32
	_ = v1618
	var v1619 int32
	_ = v1619
	var v1620 int32
	_ = v1620
	var v1621 int32
	_ = v1621
	var v1623 int32
	_ = v1623
	var v1626 int32
	_ = v1626
	var v1627 int32
	_ = v1627
	var v1628 int32
	_ = v1628
	var v1630 int32
	_ = v1630
	var v1632 int32
	_ = v1632
	var v1633 int32
	_ = v1633
	var v1637 int32
	_ = v1637
	var v1640 int32
	_ = v1640
	var v1646 int32
	_ = v1646
	var v1649 int32
	_ = v1649
	var v1652 int32
	_ = v1652
	var v1653 int32
	_ = v1653
	var v1657 int32
	_ = v1657
	var v1658 int32
	_ = v1658
	var v1661 int32
	_ = v1661
	var v1662 int32
	_ = v1662
	var v1665 int32
	_ = v1665
	var v1672 int32
	_ = v1672
	var v1673 int32
	_ = v1673
	var v1684 int32
	_ = v1684
	var v1692 int32
	_ = v1692
	var v1700 int32
	_ = v1700
	var v1708 int32
	_ = v1708
	var v1716 int32
	_ = v1716
	var v1724 int32
	_ = v1724
	var v1733 int32
	_ = v1733
	var v1737 int32
	_ = v1737
	var v1738 int32
	_ = v1738
	var v1739 int32
	_ = v1739
	var v1742 int32
	_ = v1742
	var v1745 int32
	_ = v1745
	var v1746 int32
	_ = v1746
	var v1750 int32
	_ = v1750
	var v1751 int32
	_ = v1751
	var v1754 int32
	_ = v1754
	var v1755 int32
	_ = v1755
	var v1758 int32
	_ = v1758
	var v1765 int32
	_ = v1765
	var v1766 int32
	_ = v1766
	var v1777 int32
	_ = v1777
	var v1785 int32
	_ = v1785
	var v1793 int32
	_ = v1793
	var v1801 int32
	_ = v1801
	var v1809 int32
	_ = v1809
	var v1817 int32
	_ = v1817
	var v1825 int32
	_ = v1825
	var v1833 int32
	_ = v1833
	var v1842 int32
	_ = v1842
	var v1846 int32
	_ = v1846
	var v1847 int32
	_ = v1847
	var v1848 int32
	_ = v1848
	var v1851 int32
	_ = v1851
	var v1854 int32
	_ = v1854
	var v1855 int32
	_ = v1855
	var v1859 int32
	_ = v1859
	var v1860 int32
	_ = v1860
	var v1863 int32
	_ = v1863
	var v1864 int32
	_ = v1864
	var v1867 int32
	_ = v1867
	var v1874 int32
	_ = v1874
	var v1875 int32
	_ = v1875
	var v1884 int32
	_ = v1884
	var v1889 int32
	_ = v1889
	var v1890 int32
	_ = v1890
	var v1891 int32
	_ = v1891
	var v1892 int32
	_ = v1892
	var v1898 int32
	_ = v1898
	var v1899 int32
	_ = v1899
	var v1900 int32
	_ = v1900
	var v1901 int32
	_ = v1901
	var v1902 int32
	_ = v1902
	var v1903 int32
	_ = v1903
	var v1905 int32
	_ = v1905
	var v1908 int32
	_ = v1908
	var v1909 int32
	_ = v1909
	var v1910 int32
	_ = v1910
	var v1912 int32
	_ = v1912
	var v1914 int32
	_ = v1914
	var v1915 int32
	_ = v1915
	var v1919 int32
	_ = v1919
	var v1922 int32
	_ = v1922
	var v1928 int32
	_ = v1928
	var v1931 int32
	_ = v1931
	var v1934 int32
	_ = v1934
	var v1935 int32
	_ = v1935
	var v1939 int32
	_ = v1939
	var v1940 int32
	_ = v1940
	var v1943 int32
	_ = v1943
	var v1944 int32
	_ = v1944
	var v1947 int32
	_ = v1947
	var v1954 int32
	_ = v1954
	var v1955 int32
	_ = v1955
	var v1962 int32
	_ = v1962
	var v1967 int32
	_ = v1967
	var v1968 int32
	_ = v1968
	var v1969 int32
	_ = v1969
	var v1970 int32
	_ = v1970
	var v1976 int32
	_ = v1976
	var v1977 int32
	_ = v1977
	var v1978 int32
	_ = v1978
	var v1979 int32
	_ = v1979
	var v1980 int32
	_ = v1980
	var v1981 int32
	_ = v1981
	var v1983 int32
	_ = v1983
	var v1986 int32
	_ = v1986
	var v1987 int32
	_ = v1987
	var v1988 int32
	_ = v1988
	var v1990 int32
	_ = v1990
	var v1992 int32
	_ = v1992
	var v1993 int32
	_ = v1993
	var v1997 int32
	_ = v1997
	var v2000 int32
	_ = v2000
	var v2006 int32
	_ = v2006
	var v2009 int32
	_ = v2009
	var v2010 int32
	_ = v2010
	var v2017 int32
	_ = v2017
	var v2027 int32
	_ = v2027
	var v2034 int32
	_ = v2034
	var v2045 int32
	_ = v2045
	var v2049 int32
	_ = v2049
	var v2053 int32
	_ = v2053
	v16 = F_pgp_init(m, l0)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v18 = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(l3)+8)) = v18
	*(*int64)(unsafe.Add(mBase, uint32(l3))) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l3)+16)) = v18
	*(*int64)(unsafe.Add(mBase, uint32(l3)+24)) = v18
	*(*int64)(unsafe.Add(mBase, uint32(l3)+32)) = v18
	*(*int32)(unsafe.Add(mBase, uint32(l3)+40)) = int32(-1)
	if l2 == int32(0) {
		v2034 = v16
		goto L3
	} else {
		goto L4
	}
L3:
	;
	if v2034 == int32(0) {
		goto L661
	} else {
		goto L662
	}
L4:
	;
	if v16 != 0 {
		v2034 = v16
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v32 = int32(1)
	v33 = l2 + v32
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	v36 = v34 & v32
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v34 == v32 {
		goto L11
	} else {
		goto L12
	}
L6:
	;
	v181 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v174+v177))) = uint8(v181)
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177))))
	if v185 == v181 {
		v2017 = v181
		goto L38
	} else {
		goto L39
	}
L7:
	;
	if v36 != 0 {
		goto L19
	} else {
		goto L20
	}
L8:
	;
	v70 = F_palloc(m, v67+int32(1))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L17
	}
L9:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v67 = int32(base.Ui32(v61)>>(uint(int32(2))%32)) - int32(4)
	goto L8
L10:
	;
	v59 = F_palloc(m, int32(5))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L16
	}
L11:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	if base.Ui32((v40-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	if v36 == int32(0) {
		goto L9
	} else {
		goto L15
	}
L14:
	;
	v67 = base.B2i32(v40 == int32(18)) << (uint(int32(4)) % 32)
	goto L8
L15:
	;
	v53 = int32(1)
	v67 = int32(base.Ui32(v34)>>(uint(v53)%32)) - v53
	goto L8
L16:
	;
	v75 = int32(4)
	v76 = v59
	goto L7
L17:
	;
	if v67 <= int32(0) {
		v174 = v67
		v177 = v70
		goto L6
	} else {
		goto L18
	}
L18:
	;
	v75 = v67
	v76 = v70
	goto L7
L19:
	;
	v79 = v33
	goto L21
L20:
	;
	v79 = l2 + int32(4)
	goto L21
L21:
	;
	v80 = int32(1)
	v82 = int32(0)
	if v75 != v80 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v90 = v82
	v97 = int32(0)
	goto L25
L23:
	;
	v138 = v82
	goto L24
L24:
	;
	if v75&v80 == int32(0) {
		v174 = v75
		v177 = v76
		goto L6
	} else {
		goto L34
	}
L25:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90+v79))))
	if base.Ui32((v105-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v138 = v132
	goto L24
L27:
	;
	v114 = v105 | int32(32)
	goto L29
L28:
	;
	v114 = v105
	goto L29
L29:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v90+v76))) = uint8(v114)
	v117 = v90 | int32(1)
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117+v79))))
	if base.Ui32((v120-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v129 = v120 | int32(32)
	goto L32
L31:
	;
	v129 = v120
	goto L32
L32:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v76+v117))) = uint8(v129)
	v131 = int32(2)
	v132 = v90 + v131
	v134 = v97 + v131
	if v134 != v75&int32(2147483646) {
		v90 = v132
		v97 = v134
		goto L25
	} else {
		goto L33
	}
L33:
	;
	goto L26
L34:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138+v79))))
	if base.Ui32((v155-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v164 = v155 | int32(32)
	goto L37
L36:
	;
	v164 = v155
	goto L37
L37:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v138+v76))) = uint8(v164)
	v174 = v75
	v177 = v76
	goto L6
L38:
	;
	F_pfree(m, v177)
	mBase = m.M
	v2027 = m.ExcPending
	if v2027 != 0 {
		goto L1
	} else {
		goto L660
	}
L39:
	;
	v192 = v177
	v195 = v185
	goto L40
L40:
	;
	v204 = v195 & int32(255)
	switch v204 - int32(32) {
	case 0:
		goto L45
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28:
		goto L44
	case 12, 29:
		goto L43
	default:
		goto L46
	}
L41:
	;
	v2017 = v2009
	goto L38
L42:
	;
	v254 = int32(-13)
	v257 = v240
	goto L53
L43:
	;
	v240 = v192 + int32(1)
	goto L42
L44:
	;
	v216 = v192
	v221 = v195
	goto L48
L45:
	;
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+1)))
	v192 = v192 + int32(1)
	v195 = v213
	goto L40
L46:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v204-int32(9)) {
		goto L44
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	v230 = v221 & int32(255)
	switch v230 {
	case 0, 9, 10, 32, 44:
		v240 = v216
		goto L42
	case 1, 2, 3, 4, 5, 6, 7, 8, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43:
		goto L50
	default:
		goto L51
	}
L50:
	;
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216)+1)))
	v216 = v216 + int32(1)
	v221 = v233
	goto L48
L51:
	;
	if v230 == int32(61) {
		v240 = v216
		goto L42
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v257))))
	if base.Ui32(v270-int32(9)) < base.Ui32(int32(2)) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	v345 = v281 + v342
	v350 = v345
	goto L70
L55:
	;
	goto L54
L56:
	;
	v257 = v257 + int32(1)
	goto L53
L57:
	;
	if v270 == int32(32) {
		goto L56
	} else {
		goto L58
	}
L58:
	;
	if v270 != int32(61) {
		v2017 = v254
		goto L38
	} else {
		goto L59
	}
L59:
	;
	v281 = v257
	goto L60
L60:
	;
	v294 = int32(1)
	v296 = v281 + v294
	v297 = int32(2)
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v281)+1)))
	if base.Ui32(v298-int32(9)) < base.Ui32(v297) {
		v281 = v296
		goto L60
	} else {
		goto L62
	}
L61:
	;
	v313 = v298
	v314 = v294
	goto L64
L62:
	;
	switch v298 - int32(32) {
	case 0:
		v281 = v296
		goto L60
	default:
		goto L63
	case 12, 29:
		v342 = v297
		goto L55
	}
L63:
	;
	goto L61
L64:
	;
	v321 = v313 & int32(255)
	switch v321 {
	case 0, 9, 10, 32, 44:
		v342 = v314
		goto L55
	case 1, 2, 3, 4, 5, 6, 7, 8, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43:
		goto L66
	default:
		goto L67
	}
L66:
	;
	v325 = v314 + int32(1)
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v281+v325))))
	v313 = v327
	v314 = v325
	goto L64
L67:
	;
	if v321 != int32(61) {
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v342 = v314
	goto L55
L69:
	;
	if v204 == int32(0) {
		v2017 = v254
		goto L38
	} else {
		goto L75
	}
L70:
	;
	v361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v350))))
	switch v361 {
	case 0:
		v368 = v350
		goto L69
	default:
		goto L72
	case 9, 10, 32:
		goto L73
	}
L71:
	;
	if v361 != int32(44) {
		v2017 = v254
		goto L38
	} else {
		goto L74
	}
L72:
	;
	goto L71
L73:
	;
	v350 = v350 + int32(1)
	goto L70
L74:
	;
	v368 = v350 + int32(1)
	goto L69
L75:
	;
	if v298 == int32(0) {
		v2017 = v254
		goto L38
	} else {
		goto L76
	}
L76:
	;
	if v342 == int32(1) {
		v2017 = v254
		goto L38
	} else {
		goto L77
	}
L77:
	;
	v376 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v192+(v240-v192)))) = uint8(v376)
	*(*uint8)(unsafe.Add(mBase, uint32(v345))) = uint8(v376)
	v380 = int32(228960)
	v383 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1531])))
	v384 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
	if v384 == v376 {
		v403 = v383
		v404 = v384
		goto L82
	} else {
		goto L83
	}
L78:
	;
	v2010 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v368))))
	if v2010 != 0 {
		v192 = v368
		v195 = v2010
		goto L40
	} else {
		goto L659
	}
L79:
	;
	v1142 = int32(310594)
	v1145 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1532])))
	v1146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
	if v1146 == int32(0) {
		v1165 = v1145
		v1166 = v1146
		goto L362
	} else {
		goto L363
	}
L80:
	;
	if int32(0) <= v1139 {
		v2009 = v1139
		goto L78
	} else {
		goto L360
	}
L81:
	;
	if v404-v403 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L82:
	;
	goto L81
L83:
	;
	if v383 != v384 {
		v403 = v383
		v404 = v384
		goto L82
	} else {
		goto L84
	}
L84:
	;
	v388 = v192
	v389 = v380
	goto L85
L85:
	;
	v392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v389)+1)))
	v393 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v388)+1)))
	if v393 == int32(0) {
		v403 = v392
		v404 = v393
		goto L82
	} else {
		goto L87
	}
L86:
	;
	v403 = v392
	v404 = v393
	goto L82
L87:
	;
	v396 = int32(1)
	if v392 == v393 {
		v388 = v388 + v396
		v389 = v389 + v396
		goto L85
	} else {
		goto L88
	}
L88:
	;
	goto L86
L89:
	;
	v408 = F_pgp_get_cipher_code(m, v296)
	mBase = m.M
	if v408 < int32(0) {
		goto L93
	} else {
		goto L94
	}
L90:
	;
	goto L91
L91:
	;
	v414 = int32(469202)
	v417 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1533])))
	v418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
	if v418 == int32(0) {
		v437 = v417
		v438 = v418
		goto L97
	} else {
		goto L98
	}
L92:
	;
	v1139 = v413
	goto L80
L93:
	;
	v413 = v408
	goto L92
L94:
	;
	goto L95
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+60)) = v408
	v413 = int32(0)
	goto L92
L96:
	;
	if v438-v437 == int32(0) {
		goto L104
	} else {
		goto L105
	}
L97:
	;
	goto L96
L98:
	;
	if v417 != v418 {
		v437 = v417
		v438 = v418
		goto L97
	} else {
		goto L99
	}
L99:
	;
	v422 = v192
	v423 = v414
	goto L100
L100:
	;
	v426 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v423)+1)))
	v427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v422)+1)))
	if v427 == int32(0) {
		v437 = v426
		v438 = v427
		goto L97
	} else {
		goto L102
	}
L101:
	;
	v437 = v426
	v438 = v427
	goto L97
L102:
	;
	v430 = int32(1)
	if v426 == v427 {
		v422 = v422 + v430
		v423 = v423 + v430
		goto L100
	} else {
		goto L103
	}
L103:
	;
	goto L101
L104:
	;
	v445 = v296
	goto L108
L105:
	;
	goto L106
L106:
	;
	v494 = int32(19836)
	v497 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1534])))
	v498 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
	if v498 == int32(0) {
		v517 = v497
		v518 = v498
		goto L125
	} else {
		goto L126
	}
L107:
	;
	v490 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v37)+72)) = base.B2i32(v489 != v490)
	goto L123
L108:
	;
	v450 = v445 + int32(1)
	v451 = int32(*(*int8)(unsafe.Add(mBase, uint32(v445))))
	v452 = F___isspace(m, v451)
	mBase = m.M
	if v452 != 0 {
		v445 = v450
		goto L108
	} else {
		goto L110
	}
L109:
	;
	v453 = int32(1)
	switch v451&int32(255) - int32(43) {
	case 0:
		v459 = v453
		goto L112
	default:
		v461 = v451
		v462 = v445
		v463 = v453
		goto L111
	case 2:
		goto L113
	}
L110:
	;
	goto L109
L111:
	;
	v464 = int32(0)
	v466 = v461 - int32(48)
	if base.Ui32(v466) <= base.Ui32(int32(9)) {
		goto L114
	} else {
		goto L115
	}
L112:
	;
	v460 = int32(*(*int8)(unsafe.Add(mBase, uint32(v450))))
	v461 = v460
	v462 = v450
	v463 = v459
	goto L111
L113:
	;
	v459 = int32(0)
	goto L112
L114:
	;
	v469 = v464
	v470 = v466
	v471 = v462
	goto L117
L115:
	;
	v483 = v464
	goto L116
L116:
	;
	if v463 != 0 {
		goto L120
	} else {
		goto L121
	}
L117:
	;
	v473 = int32(10)
	v475 = v469*v473 - v470
	v476 = int32(*(*int8)(unsafe.Add(mBase, uint32(v471)+1)))
	v480 = v476 - int32(48)
	if base.Ui32(v480) < base.Ui32(v473) {
		v469 = v475
		v470 = v480
		v471 = v471 + int32(1)
		goto L117
	} else {
		goto L119
	}
L118:
	;
	v483 = v475
	goto L116
L119:
	;
	goto L118
L120:
	;
	v489 = int32(0) - v483
	goto L122
L121:
	;
	v489 = v483
	goto L122
L122:
	;
	goto L107
L123:
	;
	v1139 = v490
	goto L80
L124:
	;
	if v518-v517 == int32(0) {
		goto L132
	} else {
		goto L133
	}
L125:
	;
	goto L124
L126:
	;
	if v497 != v498 {
		v517 = v497
		v518 = v498
		goto L125
	} else {
		goto L127
	}
L127:
	;
	v502 = v192
	v503 = v494
	goto L128
L128:
	;
	v506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v503)+1)))
	v507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v502)+1)))
	if v507 == int32(0) {
		v517 = v506
		v518 = v507
		goto L125
	} else {
		goto L130
	}
L129:
	;
	v517 = v506
	v518 = v507
	goto L125
L130:
	;
	v510 = int32(1)
	if v506 == v507 {
		v502 = v502 + v510
		v503 = v503 + v510
		goto L128
	} else {
		goto L131
	}
L131:
	;
	goto L129
L132:
	;
	v525 = v296
	goto L136
L133:
	;
	goto L134
L134:
	;
	v574 = int32(392180)
	v577 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1535])))
	v578 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
	if v578 == int32(0) {
		v597 = v577
		v598 = v578
		goto L153
	} else {
		goto L154
	}
L135:
	;
	v570 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v37)+76)) = base.B2i32(v569 != v570)
	goto L151
L136:
	;
	v530 = v525 + int32(1)
	v531 = int32(*(*int8)(unsafe.Add(mBase, uint32(v525))))
	v532 = F___isspace(m, v531)
	mBase = m.M
	if v532 != 0 {
		v525 = v530
		goto L136
	} else {
		goto L138
	}
L137:
	;
	v533 = int32(1)
	switch v531&int32(255) - int32(43) {
	case 0:
		v539 = v533
		goto L140
	default:
		v541 = v531
		v542 = v525
		v543 = v533
		goto L139
	case 2:
		goto L141
	}
L138:
	;
	goto L137
L139:
	;
	v544 = int32(0)
	v546 = v541 - int32(48)
	if base.Ui32(v546) <= base.Ui32(int32(9)) {
		goto L142
	} else {
		goto L143
	}
L140:
	;
	v540 = int32(*(*int8)(unsafe.Add(mBase, uint32(v530))))
	v541 = v540
	v542 = v530
	v543 = v539
	goto L139
L141:
	;
	v539 = int32(0)
	goto L140
L142:
	;
	v549 = v544
	v550 = v546
	v551 = v542
	goto L145
L143:
	;
	v563 = v544
	goto L144
L144:
	;
	if v543 != 0 {
		goto L148
	} else {
		goto L149
	}
L145:
	;
	v553 = int32(10)
	v555 = v549*v553 - v550
	v556 = int32(*(*int8)(unsafe.Add(mBase, uint32(v551)+1)))
	v560 = v556 - int32(48)
	if base.Ui32(v560) < base.Ui32(v553) {
		v549 = v555
		v550 = v560
		v551 = v551 + int32(1)
		goto L145
	} else {
		goto L147
	}
L146:
	;
	v563 = v555
	goto L144
L147:
	;
	goto L146
L148:
	;
	v569 = int32(0) - v563
	goto L150
L149:
	;
	v569 = v563
	goto L150
L150:
	;
	goto L135
L151:
	;
	v1139 = v570
	goto L80
L152:
	;
	if v598-v597 == int32(0) {
		goto L160
	} else {
		goto L161
	}
L153:
	;
	goto L152
L154:
	;
	if v577 != v578 {
		v597 = v577
		v598 = v578
		goto L153
	} else {
		goto L155
	}
L155:
	;
	v582 = v192
	v583 = v574
	goto L156
L156:
	;
	v586 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v583)+1)))
	v587 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v582)+1)))
	if v587 == int32(0) {
		v597 = v586
		v598 = v587
		goto L153
	} else {
		goto L158
	}
L157:
	;
	v597 = v586
	v598 = v587
	goto L153
L158:
	;
	v590 = int32(1)
	if v586 == v587 {
		v582 = v582 + v590
		v583 = v583 + v590
		goto L156
	} else {
		goto L159
	}
L159:
	;
	goto L157
L160:
	;
	v605 = v296
	goto L164
L161:
	;
	goto L162
L162:
	;
	v659 = int32(83362)
	v662 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1536])))
	v663 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
	if v663 == int32(0) {
		v682 = v662
		v683 = v663
		goto L184
	} else {
		goto L185
	}
L163:
	;
	v651 = int32(-13)
	if base.Ui32(int32(3)) < base.Ui32(v649) {
		v658 = v651
		goto L180
	} else {
		goto L181
	}
L164:
	;
	v610 = v605 + int32(1)
	v611 = int32(*(*int8)(unsafe.Add(mBase, uint32(v605))))
	v612 = F___isspace(m, v611)
	mBase = m.M
	if v612 != 0 {
		v605 = v610
		goto L164
	} else {
		goto L166
	}
L165:
	;
	v613 = int32(1)
	switch v611&int32(255) - int32(43) {
	case 0:
		v619 = v613
		goto L168
	default:
		v621 = v611
		v622 = v605
		v623 = v613
		goto L167
	case 2:
		goto L169
	}
L166:
	;
	goto L165
L167:
	;
	v624 = int32(0)
	v626 = v621 - int32(48)
	if base.Ui32(v626) <= base.Ui32(int32(9)) {
		goto L170
	} else {
		goto L171
	}
L168:
	;
	v620 = int32(*(*int8)(unsafe.Add(mBase, uint32(v610))))
	v621 = v620
	v622 = v610
	v623 = v619
	goto L167
L169:
	;
	v619 = int32(0)
	goto L168
L170:
	;
	v629 = v624
	v630 = v626
	v631 = v622
	goto L173
L171:
	;
	v643 = v624
	goto L172
L172:
	;
	if v623 != 0 {
		goto L176
	} else {
		goto L177
	}
L173:
	;
	v633 = int32(10)
	v635 = v629*v633 - v630
	v636 = int32(*(*int8)(unsafe.Add(mBase, uint32(v631)+1)))
	v640 = v636 - int32(48)
	if base.Ui32(v640) < base.Ui32(v633) {
		v629 = v635
		v630 = v640
		v631 = v631 + int32(1)
		goto L173
	} else {
		goto L175
	}
L174:
	;
	v643 = v635
	goto L172
L175:
	;
	goto L174
L176:
	;
	v649 = int32(0) - v643
	goto L178
L177:
	;
	v649 = v643
	goto L178
L178:
	;
	goto L163
L179:
	;
	v1139 = v658
	goto L80
L180:
	;
	goto L179
L181:
	;
	if v649 == int32(2) {
		v658 = v651
		goto L180
	} else {
		goto L182
	}
L182:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+44)) = v649
	v658 = int32(0)
	goto L180
L183:
	;
	if v683-v682 == int32(0) {
		goto L191
	} else {
		goto L192
	}
L184:
	;
	goto L183
L185:
	;
	if v662 != v663 {
		v682 = v662
		v683 = v663
		goto L184
	} else {
		goto L186
	}
L186:
	;
	v667 = v192
	v668 = v659
	goto L187
L187:
	;
	v671 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v668)+1)))
	v672 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v667)+1)))
	if v672 == int32(0) {
		v682 = v671
		v683 = v672
		goto L184
	} else {
		goto L189
	}
L188:
	;
	v682 = v671
	v683 = v672
	goto L184
L189:
	;
	v675 = int32(1)
	if v671 == v672 {
		v667 = v667 + v675
		v668 = v668 + v675
		goto L187
	} else {
		goto L190
	}
L190:
	;
	goto L188
L191:
	;
	v690 = v296
	goto L195
L192:
	;
	goto L193
L193:
	;
	v747 = int32(228893)
	v750 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1537])))
	v751 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
	if v751 == int32(0) {
		v770 = v750
		v771 = v751
		goto L215
	} else {
		goto L216
	}
L194:
	;
	v736 = int32(-13)
	if base.Ui32(int32(65010688)) < base.Ui32(v734-int32(1024)) {
		v746 = v736
		goto L211
	} else {
		goto L212
	}
L195:
	;
	v695 = v690 + int32(1)
	v696 = int32(*(*int8)(unsafe.Add(mBase, uint32(v690))))
	v697 = F___isspace(m, v696)
	mBase = m.M
	if v697 != 0 {
		v690 = v695
		goto L195
	} else {
		goto L197
	}
L196:
	;
	v698 = int32(1)
	switch v696&int32(255) - int32(43) {
	case 0:
		v704 = v698
		goto L199
	default:
		v706 = v696
		v707 = v690
		v708 = v698
		goto L198
	case 2:
		goto L200
	}
L197:
	;
	goto L196
L198:
	;
	v709 = int32(0)
	v711 = v706 - int32(48)
	if base.Ui32(v711) <= base.Ui32(int32(9)) {
		goto L201
	} else {
		goto L202
	}
L199:
	;
	v705 = int32(*(*int8)(unsafe.Add(mBase, uint32(v695))))
	v706 = v705
	v707 = v695
	v708 = v704
	goto L198
L200:
	;
	v704 = int32(0)
	goto L199
L201:
	;
	v714 = v709
	v715 = v711
	v716 = v707
	goto L204
L202:
	;
	v728 = v709
	goto L203
L203:
	;
	if v708 != 0 {
		goto L207
	} else {
		goto L208
	}
L204:
	;
	v718 = int32(10)
	v720 = v714*v718 - v715
	v721 = int32(*(*int8)(unsafe.Add(mBase, uint32(v716)+1)))
	v725 = v721 - int32(48)
	if base.Ui32(v725) < base.Ui32(v718) {
		v714 = v720
		v715 = v725
		v716 = v716 + int32(1)
		goto L204
	} else {
		goto L206
	}
L205:
	;
	v728 = v720
	goto L203
L206:
	;
	goto L205
L207:
	;
	v734 = int32(0) - v728
	goto L209
L208:
	;
	v734 = v728
	goto L209
L209:
	;
	goto L194
L210:
	;
	v1139 = v746
	goto L80
L211:
	;
	goto L210
L212:
	;
	v741 = *(*int32)(unsafe.Add(mBase, uint32(v37)+44))
	if v741 != int32(3) {
		v746 = v736
		goto L211
	} else {
		goto L213
	}
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+48)) = v734
	v746 = int32(0)
	goto L211
L214:
	;
	if v771-v770 == int32(0) {
		goto L222
	} else {
		goto L223
	}
L215:
	;
	goto L214
L216:
	;
	if v750 != v751 {
		v770 = v750
		v771 = v751
		goto L215
	} else {
		goto L217
	}
L217:
	;
	v755 = v192
	v756 = v747
	goto L218
L218:
	;
	v759 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v756)+1)))
	v760 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v755)+1)))
	if v760 == int32(0) {
		v770 = v759
		v771 = v760
		goto L215
	} else {
		goto L220
	}
L219:
	;
	v770 = v759
	v771 = v760
	goto L215
L220:
	;
	v763 = int32(1)
	if v759 == v760 {
		v755 = v755 + v763
		v756 = v756 + v763
		goto L218
	} else {
		goto L221
	}
L221:
	;
	goto L219
L222:
	;
	v775 = F_pgp_get_digest_code(m, v296)
	mBase = m.M
	if v775 < int32(0) {
		goto L226
	} else {
		goto L227
	}
L223:
	;
	goto L224
L224:
	;
	v781 = int32(228956)
	v784 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1538])))
	v785 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
	if v785 == int32(0) {
		v804 = v784
		v805 = v785
		goto L230
	} else {
		goto L231
	}
L225:
	;
	v1139 = v780
	goto L80
L226:
	;
	v780 = v775
	goto L225
L227:
	;
	goto L228
L228:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+52)) = v775
	v780 = int32(0)
	goto L225
L229:
	;
	if v805-v804 == int32(0) {
		goto L237
	} else {
		goto L238
	}
L230:
	;
	goto L229
L231:
	;
	if v784 != v785 {
		v804 = v784
		v805 = v785
		goto L230
	} else {
		goto L232
	}
L232:
	;
	v789 = v192
	v790 = v781
	goto L233
L233:
	;
	v793 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v790)+1)))
	v794 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v789)+1)))
	if v794 == int32(0) {
		v804 = v793
		v805 = v794
		goto L230
	} else {
		goto L235
	}
L234:
	;
	v804 = v793
	v805 = v794
	goto L230
L235:
	;
	v797 = int32(1)
	if v793 == v794 {
		v789 = v789 + v797
		v790 = v790 + v797
		goto L233
	} else {
		goto L236
	}
L236:
	;
	goto L234
L237:
	;
	v809 = F_pgp_get_cipher_code(m, v296)
	mBase = m.M
	if v809 < int32(0) {
		goto L241
	} else {
		goto L242
	}
L238:
	;
	goto L239
L239:
	;
	v815 = int32(228916)
	v818 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1539])))
	v819 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
	if v819 == int32(0) {
		v838 = v818
		v839 = v819
		goto L245
	} else {
		goto L246
	}
L240:
	;
	v1139 = v814
	goto L80
L241:
	;
	v814 = v809
	goto L240
L242:
	;
	goto L243
L243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+56)) = v809
	v814 = int32(0)
	goto L240
L244:
	;
	if v839-v838 == int32(0) {
		goto L252
	} else {
		goto L253
	}
L245:
	;
	goto L244
L246:
	;
	if v818 != v819 {
		v838 = v818
		v839 = v819
		goto L245
	} else {
		goto L247
	}
L247:
	;
	v823 = v192
	v824 = v815
	goto L248
L248:
	;
	v827 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v824)+1)))
	v828 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v823)+1)))
	if v828 == int32(0) {
		v838 = v827
		v839 = v828
		goto L245
	} else {
		goto L250
	}
L249:
	;
	v838 = v827
	v839 = v828
	goto L245
L250:
	;
	v831 = int32(1)
	if v827 == v828 {
		v823 = v823 + v831
		v824 = v824 + v831
		goto L248
	} else {
		goto L251
	}
L251:
	;
	goto L249
L252:
	;
	v846 = v296
	goto L256
L253:
	;
	goto L254
L254:
	;
	v898 = int32(291177)
	v901 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1540])))
	v902 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
	if v902 == int32(0) {
		v921 = v901
		v922 = v902
		goto L276
	} else {
		goto L277
	}
L255:
	;
	if base.Ui32(v890) <= base.Ui32(int32(3)) {
		goto L272
	} else {
		goto L273
	}
L256:
	;
	v851 = v846 + int32(1)
	v852 = int32(*(*int8)(unsafe.Add(mBase, uint32(v846))))
	v853 = F___isspace(m, v852)
	mBase = m.M
	if v853 != 0 {
		v846 = v851
		goto L256
	} else {
		goto L258
	}
L257:
	;
	v854 = int32(1)
	switch v852&int32(255) - int32(43) {
	case 0:
		v860 = v854
		goto L260
	default:
		v862 = v852
		v863 = v846
		v864 = v854
		goto L259
	case 2:
		goto L261
	}
L258:
	;
	goto L257
L259:
	;
	v865 = int32(0)
	v867 = v862 - int32(48)
	if base.Ui32(v867) <= base.Ui32(int32(9)) {
		goto L262
	} else {
		goto L263
	}
L260:
	;
	v861 = int32(*(*int8)(unsafe.Add(mBase, uint32(v851))))
	v862 = v861
	v863 = v851
	v864 = v860
	goto L259
L261:
	;
	v860 = int32(0)
	goto L260
L262:
	;
	v870 = v865
	v871 = v867
	v872 = v863
	goto L265
L263:
	;
	v884 = v865
	goto L264
L264:
	;
	if v864 != 0 {
		goto L268
	} else {
		goto L269
	}
L265:
	;
	v874 = int32(10)
	v876 = v870*v874 - v871
	v877 = int32(*(*int8)(unsafe.Add(mBase, uint32(v872)+1)))
	v881 = v877 - int32(48)
	if base.Ui32(v881) < base.Ui32(v874) {
		v870 = v876
		v871 = v881
		v872 = v872 + int32(1)
		goto L265
	} else {
		goto L267
	}
L266:
	;
	v884 = v876
	goto L264
L267:
	;
	goto L266
L268:
	;
	v890 = int32(0) - v884
	goto L270
L269:
	;
	v890 = v884
	goto L270
L270:
	;
	goto L255
L271:
	;
	v1139 = v897
	goto L80
L272:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+64)) = v890
	v897 = int32(0)
	goto L274
L273:
	;
	v897 = int32(-13)
	goto L274
L274:
	;
	goto L271
L275:
	;
	if v922-v921 == int32(0) {
		goto L283
	} else {
		goto L284
	}
L276:
	;
	goto L275
L277:
	;
	if v901 != v902 {
		v921 = v901
		v922 = v902
		goto L276
	} else {
		goto L278
	}
L278:
	;
	v906 = v192
	v907 = v898
	goto L279
L279:
	;
	v910 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v907)+1)))
	v911 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v906)+1)))
	if v911 == int32(0) {
		v921 = v910
		v922 = v911
		goto L276
	} else {
		goto L281
	}
L280:
	;
	v921 = v910
	v922 = v911
	goto L276
L281:
	;
	v914 = int32(1)
	if v910 == v911 {
		v906 = v906 + v914
		v907 = v907 + v914
		goto L279
	} else {
		goto L282
	}
L282:
	;
	goto L280
L283:
	;
	v929 = v296
	goto L287
L284:
	;
	goto L285
L285:
	;
	v981 = int32(322174)
	v984 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1541])))
	v985 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
	if v985 == int32(0) {
		v1004 = v984
		v1005 = v985
		goto L307
	} else {
		goto L308
	}
L286:
	;
	if base.Ui32(v973) <= base.Ui32(int32(9)) {
		goto L303
	} else {
		goto L304
	}
L287:
	;
	v934 = v929 + int32(1)
	v935 = int32(*(*int8)(unsafe.Add(mBase, uint32(v929))))
	v936 = F___isspace(m, v935)
	mBase = m.M
	if v936 != 0 {
		v929 = v934
		goto L287
	} else {
		goto L289
	}
L288:
	;
	v937 = int32(1)
	switch v935&int32(255) - int32(43) {
	case 0:
		v943 = v937
		goto L291
	default:
		v945 = v935
		v946 = v929
		v947 = v937
		goto L290
	case 2:
		goto L292
	}
L289:
	;
	goto L288
L290:
	;
	v948 = int32(0)
	v950 = v945 - int32(48)
	if base.Ui32(v950) <= base.Ui32(int32(9)) {
		goto L293
	} else {
		goto L294
	}
L291:
	;
	v944 = int32(*(*int8)(unsafe.Add(mBase, uint32(v934))))
	v945 = v944
	v946 = v934
	v947 = v943
	goto L290
L292:
	;
	v943 = int32(0)
	goto L291
L293:
	;
	v953 = v948
	v954 = v950
	v955 = v946
	goto L296
L294:
	;
	v967 = v948
	goto L295
L295:
	;
	if v947 != 0 {
		goto L299
	} else {
		goto L300
	}
L296:
	;
	v957 = int32(10)
	v959 = v953*v957 - v954
	v960 = int32(*(*int8)(unsafe.Add(mBase, uint32(v955)+1)))
	v964 = v960 - int32(48)
	if base.Ui32(v964) < base.Ui32(v957) {
		v953 = v959
		v954 = v964
		v955 = v955 + int32(1)
		goto L296
	} else {
		goto L298
	}
L297:
	;
	v967 = v959
	goto L295
L298:
	;
	goto L297
L299:
	;
	v973 = int32(0) - v967
	goto L301
L300:
	;
	v973 = v967
	goto L301
L301:
	;
	goto L286
L302:
	;
	v1139 = v980
	goto L80
L303:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+68)) = v973
	v980 = int32(0)
	goto L305
L304:
	;
	v980 = int32(-13)
	goto L305
L305:
	;
	goto L302
L306:
	;
	if v1005-v1004 == int32(0) {
		goto L314
	} else {
		goto L315
	}
L307:
	;
	goto L306
L308:
	;
	if v984 != v985 {
		v1004 = v984
		v1005 = v985
		goto L307
	} else {
		goto L309
	}
L309:
	;
	v989 = v192
	v990 = v981
	goto L310
L310:
	;
	v993 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v990)+1)))
	v994 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v989)+1)))
	if v994 == int32(0) {
		v1004 = v993
		v1005 = v994
		goto L307
	} else {
		goto L312
	}
L311:
	;
	v1004 = v993
	v1005 = v994
	goto L307
L312:
	;
	v997 = int32(1)
	if v993 == v994 {
		v989 = v989 + v997
		v990 = v990 + v997
		goto L310
	} else {
		goto L313
	}
L313:
	;
	goto L311
L314:
	;
	v1012 = v296
	goto L318
L315:
	;
	goto L316
L316:
	;
	v1061 = int32(392196)
	v1064 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1542])))
	v1065 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
	if v1065 == int32(0) {
		v1084 = v1064
		v1085 = v1065
		goto L335
	} else {
		goto L336
	}
L317:
	;
	v1057 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v37)+84)) = base.B2i32(v1056 != v1057)
	goto L333
L318:
	;
	v1017 = v1012 + int32(1)
	v1018 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1012))))
	v1019 = F___isspace(m, v1018)
	mBase = m.M
	if v1019 != 0 {
		v1012 = v1017
		goto L318
	} else {
		goto L320
	}
L319:
	;
	v1020 = int32(1)
	switch v1018&int32(255) - int32(43) {
	case 0:
		v1026 = v1020
		goto L322
	default:
		v1028 = v1018
		v1029 = v1012
		v1030 = v1020
		goto L321
	case 2:
		goto L323
	}
L320:
	;
	goto L319
L321:
	;
	v1031 = int32(0)
	v1033 = v1028 - int32(48)
	if base.Ui32(v1033) <= base.Ui32(int32(9)) {
		goto L324
	} else {
		goto L325
	}
L322:
	;
	v1027 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1017))))
	v1028 = v1027
	v1029 = v1017
	v1030 = v1026
	goto L321
L323:
	;
	v1026 = int32(0)
	goto L322
L324:
	;
	v1036 = v1031
	v1037 = v1033
	v1038 = v1029
	goto L327
L325:
	;
	v1050 = v1031
	goto L326
L326:
	;
	if v1030 != 0 {
		goto L330
	} else {
		goto L331
	}
L327:
	;
	v1040 = int32(10)
	v1042 = v1036*v1040 - v1037
	v1043 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1038)+1)))
	v1047 = v1043 - int32(48)
	if base.Ui32(v1047) < base.Ui32(v1040) {
		v1036 = v1042
		v1037 = v1047
		v1038 = v1038 + int32(1)
		goto L327
	} else {
		goto L329
	}
L328:
	;
	v1050 = v1042
	goto L326
L329:
	;
	goto L328
L330:
	;
	v1056 = int32(0) - v1050
	goto L332
L331:
	;
	v1056 = v1050
	goto L332
L332:
	;
	goto L317
L333:
	;
	v1139 = v1057
	goto L80
L334:
	;
	if v1085-v1084 != 0 {
		goto L79
	} else {
		goto L342
	}
L335:
	;
	goto L334
L336:
	;
	if v1064 != v1065 {
		v1084 = v1064
		v1085 = v1065
		goto L335
	} else {
		goto L337
	}
L337:
	;
	v1069 = v192
	v1070 = v1061
	goto L338
L338:
	;
	v1073 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1070)+1)))
	v1074 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1069)+1)))
	if v1074 == int32(0) {
		v1084 = v1073
		v1085 = v1074
		goto L335
	} else {
		goto L340
	}
L339:
	;
	v1084 = v1073
	v1085 = v1074
	goto L335
L340:
	;
	v1077 = int32(1)
	if v1073 == v1074 {
		v1069 = v1069 + v1077
		v1070 = v1070 + v1077
		goto L338
	} else {
		goto L341
	}
L341:
	;
	goto L339
L342:
	;
	v1090 = v296
	goto L344
L343:
	;
	v1135 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v37)+88)) = base.B2i32(v1134 != v1135)
	goto L359
L344:
	;
	v1095 = v1090 + int32(1)
	v1096 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1090))))
	v1097 = F___isspace(m, v1096)
	mBase = m.M
	if v1097 != 0 {
		v1090 = v1095
		goto L344
	} else {
		goto L346
	}
L345:
	;
	v1098 = int32(1)
	switch v1096&int32(255) - int32(43) {
	case 0:
		v1104 = v1098
		goto L348
	default:
		v1106 = v1096
		v1107 = v1090
		v1108 = v1098
		goto L347
	case 2:
		goto L349
	}
L346:
	;
	goto L345
L347:
	;
	v1109 = int32(0)
	v1111 = v1106 - int32(48)
	if base.Ui32(v1111) <= base.Ui32(int32(9)) {
		goto L350
	} else {
		goto L351
	}
L348:
	;
	v1105 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1095))))
	v1106 = v1105
	v1107 = v1095
	v1108 = v1104
	goto L347
L349:
	;
	v1104 = int32(0)
	goto L348
L350:
	;
	v1114 = v1109
	v1115 = v1111
	v1116 = v1107
	goto L353
L351:
	;
	v1128 = v1109
	goto L352
L352:
	;
	if v1108 != 0 {
		goto L356
	} else {
		goto L357
	}
L353:
	;
	v1118 = int32(10)
	v1120 = v1114*v1118 - v1115
	v1121 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1116)+1)))
	v1125 = v1121 - int32(48)
	if base.Ui32(v1125) < base.Ui32(v1118) {
		v1114 = v1120
		v1115 = v1125
		v1116 = v1116 + int32(1)
		goto L353
	} else {
		goto L355
	}
L354:
	;
	v1128 = v1120
	goto L352
L355:
	;
	goto L354
L356:
	;
	v1134 = int32(0) - v1128
	goto L358
L357:
	;
	v1134 = v1128
	goto L358
L358:
	;
	goto L343
L359:
	;
	v1139 = v1135
	goto L80
L360:
	;
	v2017 = v1139
	goto L38
L361:
	;
	if v1166-v1165 == int32(0) {
		goto L369
	} else {
		goto L370
	}
L362:
	;
	goto L361
L363:
	;
	if v1145 != v1146 {
		v1165 = v1145
		v1166 = v1146
		goto L362
	} else {
		goto L364
	}
L364:
	;
	v1150 = v192
	v1151 = v1142
	goto L365
L365:
	;
	v1154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1151)+1)))
	v1155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1150)+1)))
	if v1155 == int32(0) {
		v1165 = v1154
		v1166 = v1155
		goto L362
	} else {
		goto L367
	}
L366:
	;
	v1165 = v1154
	v1166 = v1155
	goto L362
L367:
	;
	v1158 = int32(1)
	if v1154 == v1155 {
		v1150 = v1150 + v1158
		v1151 = v1151 + v1158
		goto L365
	} else {
		goto L368
	}
L368:
	;
	goto L366
L369:
	;
	v1173 = v296
	goto L373
L370:
	;
	goto L371
L371:
	;
	v1220 = int32(228930)
	v1223 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1543])))
	v1224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
	if v1224 == int32(0) {
		v1243 = v1223
		v1244 = v1224
		goto L389
	} else {
		goto L390
	}
L372:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v1217
	v2009 = int32(0)
	goto L78
L373:
	;
	v1178 = v1173 + int32(1)
	v1179 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1173))))
	v1180 = F___isspace(m, v1179)
	mBase = m.M
	if v1180 != 0 {
		v1173 = v1178
		goto L373
	} else {
		goto L375
	}
L374:
	;
	v1181 = int32(1)
	switch v1179&int32(255) - int32(43) {
	case 0:
		v1187 = v1181
		goto L377
	default:
		v1189 = v1179
		v1190 = v1173
		v1191 = v1181
		goto L376
	case 2:
		goto L378
	}
L375:
	;
	goto L374
L376:
	;
	v1192 = int32(0)
	v1194 = v1189 - int32(48)
	if base.Ui32(v1194) <= base.Ui32(int32(9)) {
		goto L379
	} else {
		goto L380
	}
L377:
	;
	v1188 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1178))))
	v1189 = v1188
	v1190 = v1178
	v1191 = v1187
	goto L376
L378:
	;
	v1187 = int32(0)
	goto L377
L379:
	;
	v1197 = v1192
	v1198 = v1194
	v1199 = v1190
	goto L382
L380:
	;
	v1211 = v1192
	goto L381
L381:
	;
	if v1191 != 0 {
		goto L385
	} else {
		goto L386
	}
L382:
	;
	v1201 = int32(10)
	v1203 = v1197*v1201 - v1198
	v1204 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1199)+1)))
	v1208 = v1204 - int32(48)
	if base.Ui32(v1208) < base.Ui32(v1201) {
		v1197 = v1203
		v1198 = v1208
		v1199 = v1199 + int32(1)
		goto L382
	} else {
		goto L384
	}
L383:
	;
	v1211 = v1203
	goto L381
L384:
	;
	goto L383
L385:
	;
	v1217 = int32(0) - v1211
	goto L387
L386:
	;
	v1217 = v1211
	goto L387
L387:
	;
	goto L372
L388:
	;
	if v1244-v1243 == int32(0) {
		goto L396
	} else {
		goto L397
	}
L389:
	;
	goto L388
L390:
	;
	if v1223 != v1224 {
		v1243 = v1223
		v1244 = v1224
		goto L389
	} else {
		goto L391
	}
L391:
	;
	v1228 = v192
	v1229 = v1220
	goto L392
L392:
	;
	v1232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1229)+1)))
	v1233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1228)+1)))
	if v1233 == int32(0) {
		v1243 = v1232
		v1244 = v1233
		goto L389
	} else {
		goto L394
	}
L393:
	;
	v1243 = v1232
	v1244 = v1233
	goto L389
L394:
	;
	v1236 = int32(1)
	if v1232 == v1233 {
		v1228 = v1228 + v1236
		v1229 = v1229 + v1236
		goto L392
	} else {
		goto L395
	}
L395:
	;
	goto L393
L396:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = int32(1)
	v1255 = F_pg_strcasecmp(m, int32(161458), v296)
	mBase = m.M
	if v1255 == int32(0) {
		goto L402
	} else {
		goto L403
	}
L397:
	;
	goto L398
L398:
	;
	v1329 = int32(469195)
	v1332 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1544])))
	v1333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
	if v1333 == int32(0) {
		v1352 = v1332
		v1353 = v1333
		goto L428
	} else {
		goto L429
	}
L399:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v1326
	v2009 = int32(0)
	goto L78
L400:
	;
	goto L399
L401:
	;
	v1325 = *(*int32)(unsafe.Add(mBase, uint32(v1324)+4))
	v1326 = v1325
	goto L400
L402:
	;
	v1324 = int32(4335776)
	goto L401
L403:
	;
	goto L404
L404:
	;
	v1263 = F_pg_strcasecmp(m, int32(526064), v296)
	mBase = m.M
	if v1263 == int32(0) {
		goto L405
	} else {
		goto L406
	}
L405:
	;
	v1324 = int32(4335796)
	goto L401
L406:
	;
	goto L407
L407:
	;
	v1271 = F_pg_strcasecmp(m, int32(323202), v296)
	mBase = m.M
	if v1271 == int32(0) {
		goto L408
	} else {
		goto L409
	}
L408:
	;
	v1324 = int32(4335816)
	goto L401
L409:
	;
	goto L410
L410:
	;
	v1279 = F_pg_strcasecmp(m, int32(307068), v296)
	mBase = m.M
	if v1279 == int32(0) {
		goto L411
	} else {
		goto L412
	}
L411:
	;
	v1324 = int32(4335836)
	goto L401
L412:
	;
	goto L413
L413:
	;
	v1287 = F_pg_strcasecmp(m, int32(161958), v296)
	mBase = m.M
	if v1287 == int32(0) {
		goto L414
	} else {
		goto L415
	}
L414:
	;
	v1324 = int32(4335856)
	goto L401
L415:
	;
	goto L416
L416:
	;
	v1295 = F_pg_strcasecmp(m, int32(525448), v296)
	mBase = m.M
	if v1295 == int32(0) {
		goto L417
	} else {
		goto L418
	}
L417:
	;
	v1324 = int32(4335876)
	goto L401
L418:
	;
	goto L419
L419:
	;
	v1303 = F_pg_strcasecmp(m, int32(528684), v296)
	mBase = m.M
	if v1303 == int32(0) {
		goto L420
	} else {
		goto L421
	}
L420:
	;
	v1324 = int32(4335896)
	goto L401
L421:
	;
	goto L422
L422:
	;
	v1311 = F_pg_strcasecmp(m, int32(525828), v296)
	mBase = m.M
	if v1311 == int32(0) {
		goto L423
	} else {
		goto L424
	}
L423:
	;
	v1324 = int32(4335916)
	goto L401
L424:
	;
	goto L425
L425:
	;
	v1320 = F_pg_strcasecmp(m, int32(307077), v296)
	mBase = m.M
	if v1320 != 0 {
		v1326 = int32(-103)
		goto L400
	} else {
		goto L426
	}
L426:
	;
	v1324 = int32(4335936)
	goto L401
L427:
	;
	if v1353-v1352 == int32(0) {
		goto L435
	} else {
		goto L436
	}
L428:
	;
	goto L427
L429:
	;
	if v1332 != v1333 {
		v1352 = v1332
		v1353 = v1333
		goto L428
	} else {
		goto L430
	}
L430:
	;
	v1337 = v192
	v1338 = v1329
	goto L431
L431:
	;
	v1341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1338)+1)))
	v1342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1337)+1)))
	if v1342 == int32(0) {
		v1352 = v1341
		v1353 = v1342
		goto L428
	} else {
		goto L433
	}
L432:
	;
	v1352 = v1341
	v1353 = v1342
	goto L428
L433:
	;
	v1345 = int32(1)
	if v1341 == v1342 {
		v1337 = v1337 + v1345
		v1338 = v1338 + v1345
		goto L431
	} else {
		goto L434
	}
L434:
	;
	goto L432
L435:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = int32(1)
	v1362 = v296
	goto L439
L436:
	;
	goto L437
L437:
	;
	v1409 = int32(19829)
	v1412 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1545])))
	v1413 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
	if v1413 == int32(0) {
		v1432 = v1412
		v1433 = v1413
		goto L455
	} else {
		goto L456
	}
L438:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+36)) = v1406
	v2009 = int32(0)
	goto L78
L439:
	;
	v1367 = v1362 + int32(1)
	v1368 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1362))))
	v1369 = F___isspace(m, v1368)
	mBase = m.M
	if v1369 != 0 {
		v1362 = v1367
		goto L439
	} else {
		goto L441
	}
L440:
	;
	v1370 = int32(1)
	switch v1368&int32(255) - int32(43) {
	case 0:
		v1376 = v1370
		goto L443
	default:
		v1378 = v1368
		v1379 = v1362
		v1380 = v1370
		goto L442
	case 2:
		goto L444
	}
L441:
	;
	goto L440
L442:
	;
	v1381 = int32(0)
	v1383 = v1378 - int32(48)
	if base.Ui32(v1383) <= base.Ui32(int32(9)) {
		goto L445
	} else {
		goto L446
	}
L443:
	;
	v1377 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1367))))
	v1378 = v1377
	v1379 = v1367
	v1380 = v1376
	goto L442
L444:
	;
	v1376 = int32(0)
	goto L443
L445:
	;
	v1386 = v1381
	v1387 = v1383
	v1388 = v1379
	goto L448
L446:
	;
	v1400 = v1381
	goto L447
L447:
	;
	if v1380 != 0 {
		goto L451
	} else {
		goto L452
	}
L448:
	;
	v1390 = int32(10)
	v1392 = v1386*v1390 - v1387
	v1393 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1388)+1)))
	v1397 = v1393 - int32(48)
	if base.Ui32(v1397) < base.Ui32(v1390) {
		v1386 = v1392
		v1387 = v1397
		v1388 = v1388 + int32(1)
		goto L448
	} else {
		goto L450
	}
L449:
	;
	v1400 = v1392
	goto L447
L450:
	;
	goto L449
L451:
	;
	v1406 = int32(0) - v1400
	goto L453
L452:
	;
	v1406 = v1400
	goto L453
L453:
	;
	goto L438
L454:
	;
	if v1433-v1432 == int32(0) {
		goto L462
	} else {
		goto L463
	}
L455:
	;
	goto L454
L456:
	;
	if v1412 != v1413 {
		v1432 = v1412
		v1433 = v1413
		goto L455
	} else {
		goto L457
	}
L457:
	;
	v1417 = v192
	v1418 = v1409
	goto L458
L458:
	;
	v1421 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1418)+1)))
	v1422 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1417)+1)))
	if v1422 == int32(0) {
		v1432 = v1421
		v1433 = v1422
		goto L455
	} else {
		goto L460
	}
L459:
	;
	v1432 = v1421
	v1433 = v1422
	goto L455
L460:
	;
	v1425 = int32(1)
	if v1421 == v1422 {
		v1417 = v1417 + v1425
		v1418 = v1418 + v1425
		goto L458
	} else {
		goto L461
	}
L461:
	;
	goto L459
L462:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = int32(1)
	v1442 = v296
	goto L466
L463:
	;
	goto L464
L464:
	;
	v1489 = int32(392173)
	v1492 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1546])))
	v1493 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
	if v1493 == int32(0) {
		v1512 = v1492
		v1513 = v1493
		goto L482
	} else {
		goto L483
	}
L465:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+32)) = v1486
	v2009 = int32(0)
	goto L78
L466:
	;
	v1447 = v1442 + int32(1)
	v1448 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1442))))
	v1449 = F___isspace(m, v1448)
	mBase = m.M
	if v1449 != 0 {
		v1442 = v1447
		goto L466
	} else {
		goto L468
	}
L467:
	;
	v1450 = int32(1)
	switch v1448&int32(255) - int32(43) {
	case 0:
		v1456 = v1450
		goto L470
	default:
		v1458 = v1448
		v1459 = v1442
		v1460 = v1450
		goto L469
	case 2:
		goto L471
	}
L468:
	;
	goto L467
L469:
	;
	v1461 = int32(0)
	v1463 = v1458 - int32(48)
	if base.Ui32(v1463) <= base.Ui32(int32(9)) {
		goto L472
	} else {
		goto L473
	}
L470:
	;
	v1457 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1447))))
	v1458 = v1457
	v1459 = v1447
	v1460 = v1456
	goto L469
L471:
	;
	v1456 = int32(0)
	goto L470
L472:
	;
	v1466 = v1461
	v1467 = v1463
	v1468 = v1459
	goto L475
L473:
	;
	v1480 = v1461
	goto L474
L474:
	;
	if v1460 != 0 {
		goto L478
	} else {
		goto L479
	}
L475:
	;
	v1470 = int32(10)
	v1472 = v1466*v1470 - v1467
	v1473 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1468)+1)))
	v1477 = v1473 - int32(48)
	if base.Ui32(v1477) < base.Ui32(v1470) {
		v1466 = v1472
		v1467 = v1477
		v1468 = v1468 + int32(1)
		goto L475
	} else {
		goto L477
	}
L476:
	;
	v1480 = v1472
	goto L474
L477:
	;
	goto L476
L478:
	;
	v1486 = int32(0) - v1480
	goto L480
L479:
	;
	v1486 = v1480
	goto L480
L480:
	;
	goto L465
L481:
	;
	if v1513-v1512 == int32(0) {
		goto L489
	} else {
		goto L490
	}
L482:
	;
	goto L481
L483:
	;
	if v1492 != v1493 {
		v1512 = v1492
		v1513 = v1493
		goto L482
	} else {
		goto L484
	}
L484:
	;
	v1497 = v192
	v1498 = v1489
	goto L485
L485:
	;
	v1501 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1498)+1)))
	v1502 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1497)+1)))
	if v1502 == int32(0) {
		v1512 = v1501
		v1513 = v1502
		goto L482
	} else {
		goto L487
	}
L486:
	;
	v1512 = v1501
	v1513 = v1502
	goto L482
L487:
	;
	v1505 = int32(1)
	if v1501 == v1502 {
		v1497 = v1497 + v1505
		v1498 = v1498 + v1505
		goto L485
	} else {
		goto L488
	}
L488:
	;
	goto L486
L489:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = int32(1)
	v1522 = v296
	goto L493
L490:
	;
	goto L491
L491:
	;
	v1569 = int32(83355)
	v1572 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1547])))
	v1573 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
	if v1573 == int32(0) {
		v1592 = v1572
		v1593 = v1573
		goto L509
	} else {
		goto L510
	}
L492:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v1566
	v2009 = int32(0)
	goto L78
L493:
	;
	v1527 = v1522 + int32(1)
	v1528 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1522))))
	v1529 = F___isspace(m, v1528)
	mBase = m.M
	if v1529 != 0 {
		v1522 = v1527
		goto L493
	} else {
		goto L495
	}
L494:
	;
	v1530 = int32(1)
	switch v1528&int32(255) - int32(43) {
	case 0:
		v1536 = v1530
		goto L497
	default:
		v1538 = v1528
		v1539 = v1522
		v1540 = v1530
		goto L496
	case 2:
		goto L498
	}
L495:
	;
	goto L494
L496:
	;
	v1541 = int32(0)
	v1543 = v1538 - int32(48)
	if base.Ui32(v1543) <= base.Ui32(int32(9)) {
		goto L499
	} else {
		goto L500
	}
L497:
	;
	v1537 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1527))))
	v1538 = v1537
	v1539 = v1527
	v1540 = v1536
	goto L496
L498:
	;
	v1536 = int32(0)
	goto L497
L499:
	;
	v1546 = v1541
	v1547 = v1543
	v1548 = v1539
	goto L502
L500:
	;
	v1560 = v1541
	goto L501
L501:
	;
	if v1540 != 0 {
		goto L505
	} else {
		goto L506
	}
L502:
	;
	v1550 = int32(10)
	v1552 = v1546*v1550 - v1547
	v1553 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1548)+1)))
	v1557 = v1553 - int32(48)
	if base.Ui32(v1557) < base.Ui32(v1550) {
		v1546 = v1552
		v1547 = v1557
		v1548 = v1548 + int32(1)
		goto L502
	} else {
		goto L504
	}
L503:
	;
	v1560 = v1552
	goto L501
L504:
	;
	goto L503
L505:
	;
	v1566 = int32(0) - v1560
	goto L507
L506:
	;
	v1566 = v1560
	goto L507
L507:
	;
	goto L492
L508:
	;
	if v1593-v1592 == int32(0) {
		goto L516
	} else {
		goto L517
	}
L509:
	;
	goto L508
L510:
	;
	if v1572 != v1573 {
		v1592 = v1572
		v1593 = v1573
		goto L509
	} else {
		goto L511
	}
L511:
	;
	v1577 = v192
	v1578 = v1569
	goto L512
L512:
	;
	v1581 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1578)+1)))
	v1582 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1577)+1)))
	if v1582 == int32(0) {
		v1592 = v1581
		v1593 = v1582
		goto L509
	} else {
		goto L514
	}
L513:
	;
	v1592 = v1581
	v1593 = v1582
	goto L509
L514:
	;
	v1585 = int32(1)
	if v1581 == v1582 {
		v1577 = v1577 + v1585
		v1578 = v1578 + v1585
		goto L512
	} else {
		goto L515
	}
L515:
	;
	goto L513
L516:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = int32(1)
	v1602 = v296
	goto L520
L517:
	;
	goto L518
L518:
	;
	v1649 = int32(228886)
	v1652 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1548])))
	v1653 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
	if v1653 == int32(0) {
		v1672 = v1652
		v1673 = v1653
		goto L536
	} else {
		goto L537
	}
L519:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+16)) = v1646
	v2009 = int32(0)
	goto L78
L520:
	;
	v1607 = v1602 + int32(1)
	v1608 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1602))))
	v1609 = F___isspace(m, v1608)
	mBase = m.M
	if v1609 != 0 {
		v1602 = v1607
		goto L520
	} else {
		goto L522
	}
L521:
	;
	v1610 = int32(1)
	switch v1608&int32(255) - int32(43) {
	case 0:
		v1616 = v1610
		goto L524
	default:
		v1618 = v1608
		v1619 = v1602
		v1620 = v1610
		goto L523
	case 2:
		goto L525
	}
L522:
	;
	goto L521
L523:
	;
	v1621 = int32(0)
	v1623 = v1618 - int32(48)
	if base.Ui32(v1623) <= base.Ui32(int32(9)) {
		goto L526
	} else {
		goto L527
	}
L524:
	;
	v1617 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1607))))
	v1618 = v1617
	v1619 = v1607
	v1620 = v1616
	goto L523
L525:
	;
	v1616 = int32(0)
	goto L524
L526:
	;
	v1626 = v1621
	v1627 = v1623
	v1628 = v1619
	goto L529
L527:
	;
	v1640 = v1621
	goto L528
L528:
	;
	if v1620 != 0 {
		goto L532
	} else {
		goto L533
	}
L529:
	;
	v1630 = int32(10)
	v1632 = v1626*v1630 - v1627
	v1633 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1628)+1)))
	v1637 = v1633 - int32(48)
	if base.Ui32(v1637) < base.Ui32(v1630) {
		v1626 = v1632
		v1627 = v1637
		v1628 = v1628 + int32(1)
		goto L529
	} else {
		goto L531
	}
L530:
	;
	v1640 = v1632
	goto L528
L531:
	;
	goto L530
L532:
	;
	v1646 = int32(0) - v1640
	goto L534
L533:
	;
	v1646 = v1640
	goto L534
L534:
	;
	goto L519
L535:
	;
	if v1673-v1672 == int32(0) {
		goto L543
	} else {
		goto L544
	}
L536:
	;
	goto L535
L537:
	;
	if v1652 != v1653 {
		v1672 = v1652
		v1673 = v1653
		goto L536
	} else {
		goto L538
	}
L538:
	;
	v1657 = v192
	v1658 = v1649
	goto L539
L539:
	;
	v1661 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1658)+1)))
	v1662 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1657)+1)))
	if v1662 == int32(0) {
		v1672 = v1661
		v1673 = v1662
		goto L536
	} else {
		goto L541
	}
L540:
	;
	v1672 = v1661
	v1673 = v1662
	goto L536
L541:
	;
	v1665 = int32(1)
	if v1661 == v1662 {
		v1657 = v1657 + v1665
		v1658 = v1658 + v1665
		goto L539
	} else {
		goto L542
	}
L542:
	;
	goto L540
L543:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = int32(1)
	v1684 = F_pg_strcasecmp(m, int32(526229), v296)
	mBase = m.M
	if v1684 == int32(0) {
		goto L549
	} else {
		goto L550
	}
L544:
	;
	goto L545
L545:
	;
	v1742 = int32(228949)
	v1745 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1549])))
	v1746 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
	if v1746 == int32(0) {
		v1765 = v1745
		v1766 = v1746
		goto L569
	} else {
		goto L570
	}
L546:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+24)) = v1739
	v2009 = int32(0)
	goto L78
L547:
	;
	goto L546
L548:
	;
	v1738 = *(*int32)(unsafe.Add(mBase, uint32(v1737)+4))
	v1739 = v1738
	goto L547
L549:
	;
	v1737 = int32(4335712)
	goto L548
L550:
	;
	goto L551
L551:
	;
	v1692 = F_pg_strcasecmp(m, int32(529161), v296)
	mBase = m.M
	if v1692 == int32(0) {
		goto L552
	} else {
		goto L553
	}
L552:
	;
	v1737 = int32(4335720)
	goto L548
L553:
	;
	goto L554
L554:
	;
	v1700 = F_pg_strcasecmp(m, int32(529484), v296)
	mBase = m.M
	if v1700 == int32(0) {
		goto L555
	} else {
		goto L556
	}
L555:
	;
	v1737 = int32(4335728)
	goto L548
L556:
	;
	goto L557
L557:
	;
	v1708 = F_pg_strcasecmp(m, int32(529925), v296)
	mBase = m.M
	if v1708 == int32(0) {
		goto L558
	} else {
		goto L559
	}
L558:
	;
	v1737 = int32(4335736)
	goto L548
L559:
	;
	goto L560
L560:
	;
	v1716 = F_pg_strcasecmp(m, int32(525863), v296)
	mBase = m.M
	if v1716 == int32(0) {
		goto L561
	} else {
		goto L562
	}
L561:
	;
	v1737 = int32(4335744)
	goto L548
L562:
	;
	goto L563
L563:
	;
	v1724 = F_pg_strcasecmp(m, int32(527231), v296)
	mBase = m.M
	if v1724 == int32(0) {
		goto L564
	} else {
		goto L565
	}
L564:
	;
	v1737 = int32(4335752)
	goto L548
L565:
	;
	goto L566
L566:
	;
	v1733 = F_pg_strcasecmp(m, int32(528839), v296)
	mBase = m.M
	if v1733 != 0 {
		v1739 = int32(-104)
		goto L547
	} else {
		goto L567
	}
L567:
	;
	v1737 = int32(4335760)
	goto L548
L568:
	;
	if v1766-v1765 == int32(0) {
		goto L576
	} else {
		goto L577
	}
L569:
	;
	goto L568
L570:
	;
	if v1745 != v1746 {
		v1765 = v1745
		v1766 = v1746
		goto L569
	} else {
		goto L571
	}
L571:
	;
	v1750 = v192
	v1751 = v1742
	goto L572
L572:
	;
	v1754 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1751)+1)))
	v1755 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1750)+1)))
	if v1755 == int32(0) {
		v1765 = v1754
		v1766 = v1755
		goto L569
	} else {
		goto L574
	}
L573:
	;
	v1765 = v1754
	v1766 = v1755
	goto L569
L574:
	;
	v1758 = int32(1)
	if v1754 == v1755 {
		v1750 = v1750 + v1758
		v1751 = v1751 + v1758
		goto L572
	} else {
		goto L575
	}
L575:
	;
	goto L573
L576:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = int32(1)
	v1777 = F_pg_strcasecmp(m, int32(161458), v296)
	mBase = m.M
	if v1777 == int32(0) {
		goto L582
	} else {
		goto L583
	}
L577:
	;
	goto L578
L578:
	;
	v1851 = int32(228909)
	v1854 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1550])))
	v1855 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
	if v1855 == int32(0) {
		v1874 = v1854
		v1875 = v1855
		goto L608
	} else {
		goto L609
	}
L579:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+20)) = v1848
	v2009 = int32(0)
	goto L78
L580:
	;
	goto L579
L581:
	;
	v1847 = *(*int32)(unsafe.Add(mBase, uint32(v1846)+4))
	v1848 = v1847
	goto L580
L582:
	;
	v1846 = int32(4335776)
	goto L581
L583:
	;
	goto L584
L584:
	;
	v1785 = F_pg_strcasecmp(m, int32(526064), v296)
	mBase = m.M
	if v1785 == int32(0) {
		goto L585
	} else {
		goto L586
	}
L585:
	;
	v1846 = int32(4335796)
	goto L581
L586:
	;
	goto L587
L587:
	;
	v1793 = F_pg_strcasecmp(m, int32(323202), v296)
	mBase = m.M
	if v1793 == int32(0) {
		goto L588
	} else {
		goto L589
	}
L588:
	;
	v1846 = int32(4335816)
	goto L581
L589:
	;
	goto L590
L590:
	;
	v1801 = F_pg_strcasecmp(m, int32(307068), v296)
	mBase = m.M
	if v1801 == int32(0) {
		goto L591
	} else {
		goto L592
	}
L591:
	;
	v1846 = int32(4335836)
	goto L581
L592:
	;
	goto L593
L593:
	;
	v1809 = F_pg_strcasecmp(m, int32(161958), v296)
	mBase = m.M
	if v1809 == int32(0) {
		goto L594
	} else {
		goto L595
	}
L594:
	;
	v1846 = int32(4335856)
	goto L581
L595:
	;
	goto L596
L596:
	;
	v1817 = F_pg_strcasecmp(m, int32(525448), v296)
	mBase = m.M
	if v1817 == int32(0) {
		goto L597
	} else {
		goto L598
	}
L597:
	;
	v1846 = int32(4335876)
	goto L581
L598:
	;
	goto L599
L599:
	;
	v1825 = F_pg_strcasecmp(m, int32(528684), v296)
	mBase = m.M
	if v1825 == int32(0) {
		goto L600
	} else {
		goto L601
	}
L600:
	;
	v1846 = int32(4335896)
	goto L581
L601:
	;
	goto L602
L602:
	;
	v1833 = F_pg_strcasecmp(m, int32(525828), v296)
	mBase = m.M
	if v1833 == int32(0) {
		goto L603
	} else {
		goto L604
	}
L603:
	;
	v1846 = int32(4335916)
	goto L581
L604:
	;
	goto L605
L605:
	;
	v1842 = F_pg_strcasecmp(m, int32(307077), v296)
	mBase = m.M
	if v1842 != 0 {
		v1848 = int32(-103)
		goto L580
	} else {
		goto L606
	}
L606:
	;
	v1846 = int32(4335936)
	goto L581
L607:
	;
	if v1875-v1874 == int32(0) {
		goto L615
	} else {
		goto L616
	}
L608:
	;
	goto L607
L609:
	;
	if v1854 != v1855 {
		v1874 = v1854
		v1875 = v1855
		goto L608
	} else {
		goto L610
	}
L610:
	;
	v1859 = v192
	v1860 = v1851
	goto L611
L611:
	;
	v1863 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1860)+1)))
	v1864 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1859)+1)))
	if v1864 == int32(0) {
		v1874 = v1863
		v1875 = v1864
		goto L608
	} else {
		goto L613
	}
L612:
	;
	v1874 = v1863
	v1875 = v1864
	goto L608
L613:
	;
	v1867 = int32(1)
	if v1863 == v1864 {
		v1859 = v1859 + v1867
		v1860 = v1860 + v1867
		goto L611
	} else {
		goto L614
	}
L614:
	;
	goto L612
L615:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = int32(1)
	v1884 = v296
	goto L619
L616:
	;
	goto L617
L617:
	;
	v1931 = int32(392189)
	v1934 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1551])))
	v1935 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
	if v1935 == int32(0) {
		v1954 = v1934
		v1955 = v1935
		goto L635
	} else {
		goto L636
	}
L618:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+28)) = v1928
	v2009 = int32(0)
	goto L78
L619:
	;
	v1889 = v1884 + int32(1)
	v1890 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1884))))
	v1891 = F___isspace(m, v1890)
	mBase = m.M
	if v1891 != 0 {
		v1884 = v1889
		goto L619
	} else {
		goto L621
	}
L620:
	;
	v1892 = int32(1)
	switch v1890&int32(255) - int32(43) {
	case 0:
		v1898 = v1892
		goto L623
	default:
		v1900 = v1890
		v1901 = v1884
		v1902 = v1892
		goto L622
	case 2:
		goto L624
	}
L621:
	;
	goto L620
L622:
	;
	v1903 = int32(0)
	v1905 = v1900 - int32(48)
	if base.Ui32(v1905) <= base.Ui32(int32(9)) {
		goto L625
	} else {
		goto L626
	}
L623:
	;
	v1899 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1889))))
	v1900 = v1899
	v1901 = v1889
	v1902 = v1898
	goto L622
L624:
	;
	v1898 = int32(0)
	goto L623
L625:
	;
	v1908 = v1903
	v1909 = v1905
	v1910 = v1901
	goto L628
L626:
	;
	v1922 = v1903
	goto L627
L627:
	;
	if v1902 != 0 {
		goto L631
	} else {
		goto L632
	}
L628:
	;
	v1912 = int32(10)
	v1914 = v1908*v1912 - v1909
	v1915 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1910)+1)))
	v1919 = v1915 - int32(48)
	if base.Ui32(v1919) < base.Ui32(v1912) {
		v1908 = v1914
		v1909 = v1919
		v1910 = v1910 + int32(1)
		goto L628
	} else {
		goto L630
	}
L629:
	;
	v1922 = v1914
	goto L627
L630:
	;
	goto L629
L631:
	;
	v1928 = int32(0) - v1922
	goto L633
L632:
	;
	v1928 = v1922
	goto L633
L633:
	;
	goto L618
L634:
	;
	if v1955-v1954 != 0 {
		v2017 = v254
		goto L38
	} else {
		goto L642
	}
L635:
	;
	goto L634
L636:
	;
	if v1934 != v1935 {
		v1954 = v1934
		v1955 = v1935
		goto L635
	} else {
		goto L637
	}
L637:
	;
	v1939 = v192
	v1940 = v1931
	goto L638
L638:
	;
	v1943 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1940)+1)))
	v1944 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1939)+1)))
	if v1944 == int32(0) {
		v1954 = v1943
		v1955 = v1944
		goto L635
	} else {
		goto L640
	}
L639:
	;
	v1954 = v1943
	v1955 = v1944
	goto L635
L640:
	;
	v1947 = int32(1)
	if v1943 == v1944 {
		v1939 = v1939 + v1947
		v1940 = v1940 + v1947
		goto L638
	} else {
		goto L641
	}
L641:
	;
	goto L639
L642:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = int32(1)
	v1962 = v296
	goto L644
L643:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+40)) = v2006
	v2009 = int32(0)
	goto L78
L644:
	;
	v1967 = v1962 + int32(1)
	v1968 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1962))))
	v1969 = F___isspace(m, v1968)
	mBase = m.M
	if v1969 != 0 {
		v1962 = v1967
		goto L644
	} else {
		goto L646
	}
L645:
	;
	v1970 = int32(1)
	switch v1968&int32(255) - int32(43) {
	case 0:
		v1976 = v1970
		goto L648
	default:
		v1978 = v1968
		v1979 = v1962
		v1980 = v1970
		goto L647
	case 2:
		goto L649
	}
L646:
	;
	goto L645
L647:
	;
	v1981 = int32(0)
	v1983 = v1978 - int32(48)
	if base.Ui32(v1983) <= base.Ui32(int32(9)) {
		goto L650
	} else {
		goto L651
	}
L648:
	;
	v1977 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1967))))
	v1978 = v1977
	v1979 = v1967
	v1980 = v1976
	goto L647
L649:
	;
	v1976 = int32(0)
	goto L648
L650:
	;
	v1986 = v1981
	v1987 = v1983
	v1988 = v1979
	goto L653
L651:
	;
	v2000 = v1981
	goto L652
L652:
	;
	if v1980 != 0 {
		goto L656
	} else {
		goto L657
	}
L653:
	;
	v1990 = int32(10)
	v1992 = v1986*v1990 - v1987
	v1993 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1988)+1)))
	v1997 = v1993 - int32(48)
	if base.Ui32(v1997) < base.Ui32(v1990) {
		v1986 = v1992
		v1987 = v1997
		v1988 = v1988 + int32(1)
		goto L653
	} else {
		goto L655
	}
L654:
	;
	v2000 = v1992
	goto L652
L655:
	;
	goto L654
L656:
	;
	v2006 = int32(0) - v2000
	goto L658
L657:
	;
	v2006 = v2000
	goto L658
L658:
	;
	goto L643
L659:
	;
	goto L41
L660:
	;
	v2034 = v2017
	goto L3
L661:
	;
	v2045 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v2045 != 0 {
		goto L664
	} else {
		goto L665
	}
L662:
	;
	goto L663
L663:
	;
	F_px_THROW_ERROR(m, v2034)
	mBase = m.M
	v2053 = m.ExcPending
	if v2053 != 0 {
		goto L1
	} else {
		goto L669
	}
L664:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1552])) = int32(5456)
	goto L667
L665:
	;
	goto L666
L666:
	;
	v2049 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v2049)+80)) = l1
	goto L668
L667:
	;
	goto L666
L668:
	;
	return
L669:
	;
	base.Wasm_trap_unreachable()
	for {
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
	v14 = *(*int32)(unsafe.Add(mBase, _consts[29]))
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
	v47 = *(*int32)(unsafe.Add(mBase, _consts[30]))
	if v47 != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v28 = v17 + int32(1)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v28*int32(28))+uint32(_consts[29])))
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
	v83 = *(*int32)(unsafe.Add(mBase, _consts[31]))
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
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v63*int32(36))+uint32(_consts[30])))
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
	v119 = *(*int32)(unsafe.Add(mBase, _consts[32]))
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
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v99*int32(48))+uint32(_consts[31])))
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
	v155 = *(*int32)(unsafe.Add(mBase, _consts[33]))
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
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v135*int32(36))+uint32(_consts[32])))
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
	v191 = *(*int32)(unsafe.Add(mBase, _consts[34]))
	v194 = *(*int32)(unsafe.Add(mBase, _consts[35]))
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
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v171*int32(44))+uint32(_consts[33])))
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
	v199 = *(*int32)(unsafe.Add(mBase, _consts[36]))
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
	*(*int32)(unsafe.Add(mBase, _consts[35])) = v204
	v208 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	if v208 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v210 = int32(4053328)
	v211 = v189
	goto L40
L38:
	;
	v296 = v189
	goto L39
L39:
	;
	v309 = *(*int32)(unsafe.Add(mBase, _consts[30]))
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
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v291)+uint32(_consts[29])))
	if v294 != 0 {
		v210 = v291 + int32(4053328)
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
	v311 = int32(4051792)
	v312 = v296
	v314 = int32(0)
	goto L63
L61:
	;
	v399 = v296
	goto L62
L62:
	;
	v412 = *(*int32)(unsafe.Add(mBase, _consts[31]))
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
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v394)+uint32(_consts[30])))
	if v397 != 0 {
		v311 = v394 + int32(4051792)
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
	v414 = int32(4052656)
	v415 = v399
	v417 = int32(0)
	goto L86
L84:
	;
	v502 = v399
	goto L85
L85:
	;
	v515 = *(*int32)(unsafe.Add(mBase, _consts[32]))
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
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v497)+uint32(_consts[31])))
	if v500 != 0 {
		v414 = v497 + int32(4052656)
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
	v517 = int32(4053184)
	v518 = v502
	v520 = int32(0)
	goto L109
L107:
	;
	v605 = v502
	goto L108
L108:
	;
	v618 = *(*int32)(unsafe.Add(mBase, _consts[33]))
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
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v600)+uint32(_consts[32])))
	if v603 != 0 {
		v517 = v600 + int32(4053184)
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
	v620 = int32(4337408)
	v621 = v605
	v623 = int32(0)
	goto L132
L130:
	;
	v708 = v605
	goto L131
L131:
	;
	v720 = *(*int32)(unsafe.Add(mBase, _consts[34]))
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
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v703)+uint32(_consts[33])))
	if v706 != 0 {
		v620 = v703 + int32(4337408)
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
	*(*uint8)(unsafe.Add(mBase, _consts[37])) = uint8(v835)
	return
L153:
	;
	v724 = v720 & int32(3)
	v726 = *(*int32)(unsafe.Add(mBase, _consts[38]))
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
	v712 = int32(4442992)
	v713 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v715 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v715
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
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v713
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
	v744 = int32(4442992)
	v745 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v747 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v747
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
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v745
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
				F_errmsg(m, int32(227355), int32(0))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(470274), int32(1068), int32(33253))
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
				F_errmsg(m, int32(227355), int32(0))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(477043), int32(1164), int32(33206))
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
	v6 = int32(65535)
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
					F_errmsg(m, int32(227355), int32(0))
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(470274), int32(988), int32(33262))
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
		if v4&int32(65535) == int32(32768) {
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
					F_errmsg(m, int32(383009), int32(0))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(470274), int32(1004), int32(33262))
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
				F_errmsg(m, int32(227355), int32(0))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(477043), int32(1022), int32(33186))
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
					F_errmsg(m, int32(383321), int32(0))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(470274), int32(888), int32(33215))
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
				F_errmsg(m, int32(227355), int32(0))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(470274), int32(872), int32(33215))
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
		v42 = int64(*(*int32)(unsafe.Add(mBase, uint32((int32(0)-v31)<<(uint(int32(2))%32))+uint32(_consts[1300]))))
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
									*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = int32(297118)
									*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v61
									F_errmsg(m, int32(180805), v6+int32(16))
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(471492), int32(708), int32(76373))
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
										F_errmsg(m, int32(239370), int32(0))
										mBase = m.M
										v86 = m.ExcPending
										if v86 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(471492), int32(718), int32(76373))
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
							F_errmsg(m, int32(66479), v6)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(471492), int32(702), int32(76373))
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
					v70 = int32(305004)
					v71 = int32(65535)
					v72 = v10 & v71
					if v72 != v71 {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v72
						*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v70
						v81 = F_pg_snprintf(m, v12, int32(64), int32(637338), v8+int32(32))
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
						v88 = F_pg_snprintf(m, v12, int32(64), int32(195849), v8+int32(16))
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
						F_errmsg_internal(m, int32(27608), v8)
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(472939), int32(1188), int32(63446))
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
					v70 = int32(218990)
					v71 = int32(65535)
					v72 = v10 & v71
					if v72 != v71 {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v72
						*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v70
						v81 = F_pg_snprintf(m, v12, int32(64), int32(637338), v8+int32(32))
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
						v88 = F_pg_snprintf(m, v12, int32(64), int32(195849), v8+int32(16))
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
					v70 = int32(304996)
					v71 = int32(65535)
					v72 = v10 & v71
					if v72 != v71 {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v72
						*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v70
						v81 = F_pg_snprintf(m, v12, int32(64), int32(637338), v8+int32(32))
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
						v88 = F_pg_snprintf(m, v12, int32(64), int32(195849), v8+int32(16))
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
					v70 = int32(25066)
					v71 = int32(65535)
					v72 = v10 & v71
					if v72 != v71 {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v72
						*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v70
						v81 = F_pg_snprintf(m, v12, int32(64), int32(637338), v8+int32(32))
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
						v88 = F_pg_snprintf(m, v12, int32(64), int32(195849), v8+int32(16))
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
						v70 = int32(195927)
						v71 = int32(65535)
						v72 = v10 & v71
						if v72 != v71 {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v72
							*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v70
							v81 = F_pg_snprintf(m, v12, int32(64), int32(637338), v8+int32(32))
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
							v88 = F_pg_snprintf(m, v12, int32(64), int32(195849), v8+int32(16))
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
							F_errmsg_internal(m, int32(27608), v8)
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(472939), int32(1188), int32(63446))
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
						v70 = int32(195920)
						v71 = int32(65535)
						v72 = v10 & v71
						if v72 != v71 {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v72
							*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v70
							v81 = F_pg_snprintf(m, v12, int32(64), int32(637338), v8+int32(32))
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
							v88 = F_pg_snprintf(m, v12, int32(64), int32(195849), v8+int32(16))
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
							v70 = int32(331769)
							v71 = int32(65535)
							v72 = v10 & v71
							if v72 != v71 {
								*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v72
								*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v70
								v81 = F_pg_snprintf(m, v12, int32(64), int32(637338), v8+int32(32))
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
								v88 = F_pg_snprintf(m, v12, int32(64), int32(195849), v8+int32(16))
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
								F_errmsg_internal(m, int32(27608), v8)
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(472939), int32(1188), int32(63446))
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
				if base.Ui32(v21) <= base.Ui32(int32(6143)) {
					switch v21 - int32(3072) {
					case 0:
						v70 = int32(331761)
						v71 = int32(65535)
						v72 = v10 & v71
						if v72 != v71 {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v72
							*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v70
							v81 = F_pg_snprintf(m, v12, int32(64), int32(637338), v8+int32(32))
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
							v88 = F_pg_snprintf(m, v12, int32(64), int32(195849), v8+int32(16))
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
							F_errmsg_internal(m, int32(27608), v8)
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(472939), int32(1188), int32(63446))
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
						v70 = int32(331746)
						v71 = int32(65535)
						v72 = v10 & v71
						if v72 != v71 {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v72
							*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v70
							v81 = F_pg_snprintf(m, v12, int32(64), int32(637338), v8+int32(32))
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
							v88 = F_pg_snprintf(m, v12, int32(64), int32(195849), v8+int32(16))
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
						if v21 != int32(4096) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = v10
								F_errmsg_internal(m, int32(27608), v8)
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(472939), int32(1188), int32(63446))
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
							v70 = int32(405116)
							v71 = int32(65535)
							v72 = v10 & v71
							if v72 != v71 {
								*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v72
								*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v70
								v81 = F_pg_snprintf(m, v12, int32(64), int32(637338), v8+int32(32))
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
								v88 = F_pg_snprintf(m, v12, int32(64), int32(195849), v8+int32(16))
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
					switch v21 - int32(7168) {
					case 0:
						v70 = int32(404983)
						v71 = int32(65535)
						v72 = v10 & v71
						if v72 != v71 {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v72
							*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v70
							v81 = F_pg_snprintf(m, v12, int32(64), int32(637338), v8+int32(32))
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
							v88 = F_pg_snprintf(m, v12, int32(64), int32(195849), v8+int32(16))
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
							F_errmsg_internal(m, int32(27608), v8)
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(472939), int32(1188), int32(63446))
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
						v70 = int32(404968)
						v71 = int32(65535)
						v72 = v10 & v71
						if v72 != v71 {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v72
							*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v70
							v81 = F_pg_snprintf(m, v12, int32(64), int32(637338), v8+int32(32))
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
							v88 = F_pg_snprintf(m, v12, int32(64), int32(195849), v8+int32(16))
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
						if v21 == int32(6144) {
							v70 = int32(404999)
							v71 = int32(65535)
							v72 = v10 & v71
							if v72 != v71 {
								*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v72
								*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v70
								v81 = F_pg_snprintf(m, v12, int32(64), int32(637338), v8+int32(32))
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
								v88 = F_pg_snprintf(m, v12, int32(64), int32(195849), v8+int32(16))
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
							if v21 != int32(32767) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v8))) = v10
									F_errmsg_internal(m, int32(27608), v8)
									mBase = m.M
									v62 = m.ExcPending
									if v62 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(472939), int32(1188), int32(63446))
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
								v70 = int32(715480)
								v71 = int32(65535)
								v72 = v10 & v71
								if v72 != v71 {
									*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v72
									*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v70
									v81 = F_pg_snprintf(m, v12, int32(64), int32(637338), v8+int32(32))
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
									v88 = F_pg_snprintf(m, v12, int32(64), int32(195849), v8+int32(16))
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
	v16 = *(*int32)(unsafe.Add(mBase, _consts[511]))
	v19 = *(*int32)(unsafe.Add(mBase, _consts[512]))
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
	v23 = int32(4443044)
	v24 = *(*int32)(unsafe.Add(mBase, _consts[173]))
	v27 = *(*int32)(unsafe.Add(mBase, _consts[205]))
	*(*int32)(unsafe.Add(mBase, _consts[173])) = v27
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
	*(*int32)(unsafe.Add(mBase, _consts[511])) = v32
	v38 = *(*int32)(unsafe.Add(mBase, _consts[512]))
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
	*(*int32)(unsafe.Add(mBase, _consts[173])) = v24
	goto L13
L22:
	;
	*(*int32)(unsafe.Add(mBase, _consts[512])) = v45
	goto L21
L23:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _consts[511]))
	v62 = *(*int32)(unsafe.Add(mBase, _consts[512]))
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
	if base.Ui32(v88-int32(8212)) <= base.Ui32(int32(-8197)) {
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
	F_errmsg(m, int32(461041), v13)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L17
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(474676), int32(417), int32(300552))
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
	F_errmsg_internal(m, int32(102760), int32(0))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L17
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(474676), int32(374), int32(323809))
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
	F_errmsg(m, int32(453009), v13+int32(32))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L17
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(474676), int32(153), int32(411105))
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
	F_errmsg_internal(m, int32(409982), v13+int32(16))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L17
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(474676), int32(430), int32(300552))
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
			*(*int32)(unsafe.Add(mBase, _consts[140])) = v17
			v22 = int32(0)
		}
	} else {
		v17 = v9
		*(*int32)(unsafe.Add(mBase, _consts[140])) = v17
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
		v22 = F_local2local(m, v6, v5, v10, int32(25), int32(20), int32(2177824), base.B2i32(v7 != int32(0)))
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
		if base.Ui32(l0) <= base.Ui32(int32(131071)) {
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(l0)>>(uint(int32(8))%32)))+uint32(_consts[1561]))))
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(l0)>>(uint(int32(3))%32))&int32(31)|v18<<(uint(int32(5))%32))+uint32(_consts[1561]))))
			v32 = int32(base.Ui32(v24)>>(uint(l0&int32(7))%32)) & int32(1)
		} else {
			v32 = base.B2i32(base.Ui32(l0) < base.Ui32(int32(196606)))
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
				v61 = *(*int32)(unsafe.Add(mBase, _consts[485]))
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
