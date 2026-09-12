package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_get_rule_expr_funccall(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if l0 == int32(0) {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		F_appendStringInfoString(m, v18, int32(648131))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			F_get_rule_expr(m, l0, l1, int32(0))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return
			} else {
				v25 = F_exprType(m, l0)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					v27 = F_exprTypmod(m, l0)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						v29 = F_format_type_with_typemod(m, v25, v27)
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = v29
							F_appendStringInfo(m, v18, int32(638425), v7)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return
							} else {
								m.G0 = v7 + int32(16)
								return
							}
						}
					}
				}
			}
		}
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		switch v11 - int32(15) {
		case 0:
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			switch v14 {
			case 0, 3:
				F_get_rule_expr(m, l0, l1, int32(1))
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return
				} else {
					m.G0 = v7 + int32(16)
					return
				}
			default:
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				F_appendStringInfoString(m, v18, int32(648131))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return
				} else {
					F_get_rule_expr(m, l0, l1, int32(0))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return
					} else {
						v25 = F_exprType(m, l0)
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return
						} else {
							v27 = F_exprTypmod(m, l0)
							mBase = m.M
							v28 = m.ExcPending
							if v28 != 0 {
								return
							} else {
								v29 = F_format_type_with_typemod(m, v25, v27)
								mBase = m.M
								v30 = m.ExcPending
								if v30 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7))) = v29
									F_appendStringInfo(m, v18, int32(638425), v7)
									mBase = m.M
									v34 = m.ExcPending
									if v34 != 0 {
										return
									} else {
										m.G0 = v7 + int32(16)
										return
									}
								}
							}
						}
					}
				}
			}
		default:
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			F_appendStringInfoString(m, v18, int32(648131))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				F_get_rule_expr(m, l0, l1, int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					v25 = F_exprType(m, l0)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return
					} else {
						v27 = F_exprTypmod(m, l0)
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return
						} else {
							v29 = F_format_type_with_typemod(m, v25, v27)
							mBase = m.M
							v30 = m.ExcPending
							if v30 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7))) = v29
								F_appendStringInfo(m, v18, int32(638425), v7)
								mBase = m.M
								v34 = m.ExcPending
								if v34 != 0 {
									return
								} else {
									m.G0 = v7 + int32(16)
									return
								}
							}
						}
					}
				}
			}
		case 4, 23, 24, 25, 26, 33:
			F_get_rule_expr(m, l0, l1, int32(1))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				m.G0 = v7 + int32(16)
				return
			}
		}
	}
}
func F_get_rule_expr_paren(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v6&int32(1) == int32(0) {
		F_get_rule_expr(m, l0, l1, l2)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return
		} else {
			return
		}
	} else {
		v11 = F_isSimpleNode(m, l0, l3, v6)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			if v11 != 0 {
				F_get_rule_expr(m, l0, l1, l2)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					return
				}
			} else {
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				F_appendStringInfoChar(m, v13, int32(40))
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return
				} else {
					F_get_rule_expr(m, l0, l1, l2)
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return
					} else {
						v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						F_appendStringInfoChar(m, v19, int32(41))
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
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
