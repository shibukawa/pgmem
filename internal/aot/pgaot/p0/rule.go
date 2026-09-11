package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_get_rule_list_toplevel(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	if l0 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v8 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_appendStringInfoString(m, v13, int32(_a_F_get_rule_list_toplevel_0))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	if v12 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v27 = int32(1)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v28 <= v27 {
		goto L1
	} else {
		goto L12
	}
L7:
	;
	F_get_rule_expr(m, v12, l1, l2)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L11
	}
L8:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	if v19 != int32(6) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v23 = F_get_variable(m, v12, int32(1), l1)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	goto L6
L11:
	;
	goto L6
L12:
	;
	v35 = v27
	goto L13
L13:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v36+v35<<(uint(int32(2))%32))))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_appendStringInfoString(m, v41, int32(_a_F_get_rule_list_toplevel_1))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L4
	} else {
		goto L15
	}
L14:
	;
	goto L1
L15:
	;
	if v40 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v56 = v35 + int32(1)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v56 < v57 {
		v35 = v56
		goto L13
	} else {
		goto L22
	}
L17:
	;
	F_get_rule_expr(m, v40, l1, l2)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L4
	} else {
		goto L21
	}
L18:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	if v47 != int32(6) {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v51 = F_get_variable(m, v40, int32(1), l1)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	goto L16
L21:
	;
	goto L16
L22:
	;
	goto L14
}
func F_get_rule_sortgroupclause(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
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
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v12 = F_get_sortgroupref_tle(m, l0, l1)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
		if l2 != 0 {
			v17 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+8)))
			*(*int32)(unsafe.Add(mBase, uint32(v9))) = v17
			F_appendStringInfo(m, v11, int32(_a_F_get_rule_sortgroupclause_0), v9)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				m.G0 = v9 + int32(16)
				return v16
			}
		} else {
			if v16 == int32(0) {
				m.G0 = v9 + int32(16)
				return v16
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
				switch v24 - int32(6) {
				case 0:
					v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+35)))
					v31 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+35)) = uint8(v31)
					v34 = F_get_variable(m, v16, int32(0), l3)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						*(*uint8)(unsafe.Add(mBase, uint32(l3)+35)) = uint8(v30)
						m.G0 = v9 + int32(16)
						return v16
					}
				case 1:
					F_get_const_expr(m, v16, l3, int32(1))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						m.G0 = v9 + int32(16)
						return v16
					}
				default:
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
					if v37&int32(1) != 0 {
						v45 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
						F_appendStringInfoChar(m, v45, int32(40))
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return int32(0)
						} else {
							F_get_rule_expr(m, v16, l3, int32(1))
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return int32(0)
							} else {
								v52 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
								F_appendStringInfoChar(m, v52, int32(41))
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
									return int32(0)
								} else {
									m.G0 = v9 + int32(16)
									return v16
								}
							}
						}
					} else {
						switch v24 - int32(9) {
						case 0, 2, 6, 36:
							v45 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
							F_appendStringInfoChar(m, v45, int32(40))
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int32(0)
							} else {
								F_get_rule_expr(m, v16, l3, int32(1))
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return int32(0)
								} else {
									v52 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
									F_appendStringInfoChar(m, v52, int32(41))
									mBase = m.M
									v55 = m.ExcPending
									if v55 != 0 {
										return int32(0)
									} else {
										m.G0 = v9 + int32(16)
										return v16
									}
								}
							}
						default:
							F_get_rule_expr(m, v16, l3, int32(1))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return int32(0)
							} else {
								m.G0 = v9 + int32(16)
								return v16
							}
						}
					}
				}
			}
		}
	}
}
