package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetConfigOptionFlags(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	v5 = F_find_option(m, l0, int32(0), int32(1), int32(21))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 != 0 {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(v5)+20))
			v11 = v9
		} else {
			v11 = int32(0)
		}
		return v11
	}
}
func F_config_enum_get_options(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	F_initStringInfo(m, v9)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	F_appendStringInfoString(m, v9, l1)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if l3&int32(3) == int32(0) {
		v40 = l3
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if v74 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L5:
	;
	v73 = v65 - l3
	goto L4
L6:
	;
	v44 = v40
	goto L15
L7:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v24 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v73 = int32(0)
	goto L4
L9:
	;
	goto L10
L10:
	;
	v29 = l3
	goto L11
L11:
	;
	v33 = v29 + int32(1)
	if v33&int32(3) == int32(0) {
		v40 = v33
		goto L6
	} else {
		goto L13
	}
L12:
	;
	v65 = v33
	goto L5
L13:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	if v38 != 0 {
		v29 = v33
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	v53 = int32(-2139062144)
	if (int32(16843008)-v50|v50)&v53 == v53 {
		v44 = v44 + int32(4)
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v59 = v44
	goto L18
L17:
	;
	goto L16
L18:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	if v63 != 0 {
		v59 = v59 + int32(1)
		goto L18
	} else {
		goto L20
	}
L19:
	;
	v65 = v59
	goto L5
L20:
	;
	goto L19
L21:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v73 <= v101 {
		goto L32
	} else {
		goto L33
	}
L22:
	;
	v78 = v74
	goto L23
L23:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	if v83 == int32(0) {
		goto L21
	} else {
		goto L25
	}
L24:
	;
	goto L21
L25:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+8)))
	if v86 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	F_appendStringInfoString(m, v9, v83)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v94 = v78 + int32(12)
	if v94 != 0 {
		v78 = v94
		goto L23
	} else {
		goto L31
	}
L29:
	;
	F_appendBinaryStringInfo(m, v9, l3, v73)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	goto L24
L32:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v106 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v103+(v101-v73)))) = uint8(v106)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v108 - v73
	goto L34
L33:
	;
	goto L34
L34:
	;
	F_appendStringInfoString(m, v9, l2)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	m.G0 = v9 + int32(16)
	return v113
}
func F_get_config_unit_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v19 int32
	_ = v19
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	v3 = m.G0
	v5 = v3 - int32(48)
	m.G0 = v5
	v8 = l0 & int32(2130706432)
	if base.Ui32(v8) <= base.Ui32(int32(67108863)) {
		if v8 <= int32(33554431) {
			if v8 != 0 {
				if v8 != int32(16777216) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v76 = m.ExcPending
					if v76 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v5))) = v8
						F_errmsg_internal(m, int32(476286), v5)
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(492061), int32(2854), int32(372750))
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
				} else {
					v68 = int32(536571)
					m.G0 = v5 + int32(48)
					return v68
				}
			} else {
				v68 = int32(0)
				m.G0 = v5 + int32(48)
				return v68
			}
		} else {
			if v8 == int32(33554432) {
				v39 = int32(*(*uint8)(unsafe.Add(mBase, _consts[961])))
				if v39 != 0 {
					v68 = int32(4468800)
					m.G0 = v5 + int32(48)
					return v68
				} else {
					v41 = int32(8)
					*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = v41
					v43 = int32(4468800)
					v49 = F_pg_snprintf(m, v43, v41, int32(536506), v5+int32(16))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						v68 = v43
						m.G0 = v5 + int32(48)
						return v68
					}
				}
			} else {
				if v8 != int32(50331648) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v76 = m.ExcPending
					if v76 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v5))) = v8
						F_errmsg_internal(m, int32(476286), v5)
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(492061), int32(2854), int32(372750))
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
				} else {
					v19 = int32(*(*uint8)(unsafe.Add(mBase, _consts[962])))
					if v19 == int32(0) {
						v53 = int32(8)
						*(*int32)(unsafe.Add(mBase, uint32(v5)+32)) = v53
						v55 = int32(4468808)
						v61 = F_pg_snprintf(m, v55, v53, int32(536506), v5+int32(32))
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int32(0)
						} else {
							v68 = v55
							m.G0 = v5 + int32(48)
							return v68
						}
					} else {
						v68 = int32(4468808)
						m.G0 = v5 + int32(48)
						return v68
					}
				}
			}
		}
	} else {
		if v8 <= int32(268435455) {
			if v8 == int32(67108864) {
				v68 = int32(536599)
				m.G0 = v5 + int32(48)
				return v68
			} else {
				if v8 != int32(83886080) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v76 = m.ExcPending
					if v76 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v5))) = v8
						F_errmsg_internal(m, int32(476286), v5)
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(492061), int32(2854), int32(372750))
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
				} else {
					v68 = int32(536621)
					m.G0 = v5 + int32(48)
					return v68
				}
			}
		} else {
			if v8 == int32(268435456) {
				v68 = int32(149397)
				m.G0 = v5 + int32(48)
				return v68
			} else {
				if v8 == int32(536870912) {
					v68 = int32(203559)
					m.G0 = v5 + int32(48)
					return v68
				} else {
					if v8 != int32(805306368) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v76 = m.ExcPending
						if v76 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v5))) = v8
							F_errmsg_internal(m, int32(476286), v5)
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(492061), int32(2854), int32(372750))
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
					} else {
						v68 = int32(272374)
						m.G0 = v5 + int32(48)
						return v68
					}
				}
			}
		}
	}
}
