package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecMergeNotMatched(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
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
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 float64
	_ = v122
	var v127 int32
	_ = v127
	v4 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+172))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v4
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v4
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v22
	if v16 == v4 {
		v127 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v14 + int32(16)
	return v127
L2:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v28 <= int32(0) {
		v127 = v4
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v32 = v4
	goto L5
L4:
	;
	switch v48 - int32(3) {
	case 0:
		goto L12
	default:
		goto L13
	case 4:
		v127 = int32(0)
		goto L1
	}
L5:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42+v32<<(uint(int32(2))%32))))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
	if v49 == int32(0) {
		goto L4
	} else {
		goto L7
	}
L6:
	;
	v127 = int32(0)
	goto L1
L7:
	;
	v52 = int32(4554240)
	v53 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v55
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v49)+20))
	v60 = m.T0[v59].(func(*base.Module, int32, int32, int32) int32)(m, v49, v19, v14+int32(15))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return int32(0)
L9:
	;
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v53
	if v60 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v67 = v32 + int32(1)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v67 < v68 {
		v32 = v67
		goto L5
	} else {
		goto L11
	}
L11:
	;
	goto L6
L12:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+72))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+12))
	m.T0[v93].(func(*base.Module, int32))(m, v91)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L8
	} else {
		goto L17
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L8
	} else {
		goto L14
	}
L14:
	;
	F_errmsg_internal(m, int32(376854), int32(0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L8
	} else {
		goto L15
	}
L15:
	;
	F_errfinish(m, int32(522671), int32(3663), int32(479711))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L8
	} else {
		goto L16
	}
L16:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L17:
	;
	v96 = int32(4554240)
	v97 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v90)+20))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v99
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v89)+24))
	v105 = m.T0[v104].(func(*base.Module, int32, int32, int32) int32)(m, v89+int32(4), v90, int32(0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L8
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v97
	v109 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+4)))
	v111 = v109 & int32(65533)
	*(*uint16)(unsafe.Add(mBase, uint32(v91)+4)) = uint16(v111)
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	*(*uint16)(unsafe.Add(mBase, uint32(v91)+6)) = uint16(v114)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+216)) = v46
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v18)+120))
	v118 = int32(0)
	v120 = F_ExecInsert(m, l0, v117, v91, l2, v118, v118)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L8
	} else {
		goto L19
	}
L19:
	;
	v122 = *(*float64)(unsafe.Add(mBase, uint32(v18)+224))
	*(*float64)(unsafe.Add(mBase, uint32(v18)+224)) = base.F64_add(v122, float64(1))
	v127 = v120
	goto L1
}
func F_findNotNullConstraint(m *base.Module, l0 int32, l1 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v3 = F_get_attnum(m, l0, l1)
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		if v3 <= int32(0) {
			return int32(0)
		} else {
			v11 = F_findNotNullConstraintAttnum(m, l0, v3)
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				return v11
			}
		}
	}
}
func F_verifyNotNullPKCompatible(m *base.Module, l0 int32, l1 int32) {
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
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	v5 = m.G0
	v7 = v5 - int32(112)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
	v11 = v9 + v10
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+72)))
	if v12 == int32(110) {
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+106)))
		if v15 == int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return
			} else {
				F_errcode(m, int32(325))
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = l1
					F_errmsg(m, int32(743056), v7+int32(32))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return
					} else {
						v53 = *(*int32)(unsafe.Add(mBase, uint32(v11)+80))
						v54 = F_get_rel_name(m, v53)
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = int32(545466)
							*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v54
							*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = l1
							*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v11 + int32(4)
							F_errdetail(m, int32(602084), v7+int32(16))
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(545477)
								F_errhint(m, int32(631954), v7)
								mBase = m.M
								v72 = m.ExcPending
								if v72 != 0 {
									return
								} else {
									F_errfinish(m, int32(517363), int32(9595), int32(409043))
									mBase = m.M
									v77 = m.ExcPending
									if v77 != 0 {
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
				}
			}
		} else {
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+76)))
			if v18 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v81 = m.ExcPending
				if v81 != 0 {
					return
				} else {
					F_errcode(m, int32(325))
					mBase = m.M
					v84 = m.ExcPending
					if v84 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+80)) = l1
						F_errmsg(m, int32(743056), v7+int32(80))
						mBase = m.M
						v90 = m.ExcPending
						if v90 != 0 {
							return
						} else {
							v91 = *(*int32)(unsafe.Add(mBase, uint32(v11)+80))
							v92 = F_get_rel_name(m, v91)
							mBase = m.M
							v93 = m.ExcPending
							if v93 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7)+76)) = int32(568311)
								*(*int32)(unsafe.Add(mBase, uint32(v7)+72)) = v92
								*(*int32)(unsafe.Add(mBase, uint32(v7)+68)) = l1
								*(*int32)(unsafe.Add(mBase, uint32(v7)+64)) = v11 + int32(4)
								F_errdetail(m, int32(602084), v7-int32(-64))
								mBase = m.M
								v105 = m.ExcPending
								if v105 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = int32(542852)
									F_errhint(m, int32(631914), v7+int32(48))
									mBase = m.M
									v112 = m.ExcPending
									if v112 != 0 {
										return
									} else {
										F_errfinish(m, int32(517363), int32(9607), int32(409043))
										mBase = m.M
										v117 = m.ExcPending
										if v117 != 0 {
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
					}
				}
			} else {
				m.G0 = v7 + int32(112)
				return
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
			*(*int32)(unsafe.Add(mBase, uint32(v7)+96)) = v28
			F_errmsg_internal(m, int32(96256), v7+int32(96))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return
			} else {
				F_errfinish(m, int32(517363), int32(9583), int32(409043))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
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
