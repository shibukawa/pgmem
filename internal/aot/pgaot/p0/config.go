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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
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
	v17 = F_strlen(m, l3)
	mBase = m.M
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if v18 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v17 <= v45 {
		goto L15
	} else {
		goto L16
	}
L5:
	;
	v22 = v18
	goto L6
L6:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	if v27 == int32(0) {
		goto L4
	} else {
		goto L8
	}
L7:
	;
	goto L4
L8:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+8)))
	if v30 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	F_appendStringInfoString(m, v9, v27)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v38 = v22 + int32(12)
	if v38 != 0 {
		v22 = v38
		goto L6
	} else {
		goto L14
	}
L12:
	;
	F_appendBinaryStringInfo(m, v9, l3, v17)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	goto L7
L15:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v50 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v47+(v45-v17)))) = uint8(v50)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v52 - v17
	goto L17
L16:
	;
	goto L17
L17:
	;
	F_appendStringInfoString(m, v9, l2)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	m.G0 = v9 + int32(16)
	return v57
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
						F_errmsg_internal(m, int32(480654), v5)
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(496606), int32(2854), int32(376237))
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
					v68 = int32(541526)
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
					v68 = int32(4485088)
					m.G0 = v5 + int32(48)
					return v68
				} else {
					v41 = int32(8)
					*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = v41
					v43 = int32(4485088)
					v49 = F_pg_snprintf(m, v43, v41, int32(541461), v5+int32(16))
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
						F_errmsg_internal(m, int32(480654), v5)
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(496606), int32(2854), int32(376237))
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
						v55 = int32(4485096)
						v61 = F_pg_snprintf(m, v55, v53, int32(541461), v5+int32(32))
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
						v68 = int32(4485096)
						m.G0 = v5 + int32(48)
						return v68
					}
				}
			}
		}
	} else {
		if v8 <= int32(268435455) {
			if v8 == int32(67108864) {
				v68 = int32(541555)
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
						F_errmsg_internal(m, int32(480654), v5)
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(496606), int32(2854), int32(376237))
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
					v68 = int32(541583)
					m.G0 = v5 + int32(48)
					return v68
				}
			}
		} else {
			if v8 == int32(268435456) {
				v68 = int32(150882)
				m.G0 = v5 + int32(48)
				return v68
			} else {
				if v8 == int32(536870912) {
					v68 = int32(205285)
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
							F_errmsg_internal(m, int32(480654), v5)
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(496606), int32(2854), int32(376237))
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
						v68 = int32(274782)
						m.G0 = v5 + int32(48)
						return v68
					}
				}
			}
		}
	}
}
