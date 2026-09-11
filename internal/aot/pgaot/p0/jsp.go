package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_extract_jsp_bool_expr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	v5 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(96)
	m.G0 = v10
	F_check_stack_depth(m)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		switch v16 - int32(4) {
		case 0, 1:
			F_jspGetArg(m, l2, v10+int32(68))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v23
				v29 = F_extract_jsp_bool_expr(m, l0, v10+int32(4), v10+int32(68), l3)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					F_jspGetRightArg(m, l2, v10+int32(68))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = v35
						v39 = F_extract_jsp_bool_expr(m, l0, v10, v10+int32(68), l3)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							v41 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
							if v39 != 0 {
								v43 = v29
							} else {
								v43 = int32(0)
							}
							if v43 == int32(0) {
								if v29 != 0 {
									v46 = v29
								} else {
									v46 = v39
								}
								if v41 != int32(5) {
									v50 = v46
								} else {
									v50 = int32(0)
								}
								v146 = v50
								m.G0 = v10 + int32(96)
								return v146
							} else {
								v52 = F_palloc(m, int32(16))
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v52)+8)) = v29
									*(*int32)(unsafe.Add(mBase, uint32(v52)+4)) = int32(2)
									*(*int32)(unsafe.Add(mBase, uint32(v52)+12)) = v39
									*(*int32)(unsafe.Add(mBase, uint32(v52))) = l3 ^ base.B2i32(v41 == int32(4))
									v146 = v52
									m.G0 = v10 + int32(96)
									return v146
								}
							}
						}
					}
				}
			}
		case 2:
			F_jspGetArg(m, l2, v10+int32(68))
			mBase = m.M
			v65 = m.ExcPending
			if v65 != 0 {
				return int32(0)
			} else {
				v66 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v66
				v74 = F_extract_jsp_bool_expr(m, l0, v10+int32(8), v10+int32(68), l3^int32(1))
				mBase = m.M
				v75 = m.ExcPending
				if v75 != 0 {
					return int32(0)
				} else {
					v146 = v74
					m.G0 = v10 + int32(96)
					return v146
				}
			}
		default:
			v146 = v5
			m.G0 = v10 + int32(96)
			return v146
		case 4:
			if l3 != 0 {
				v146 = v5
				m.G0 = v10 + int32(96)
				return v146
			} else {
				F_jspGetArg(m, l2, v10+int32(68))
				mBase = m.M
				v92 = m.ExcPending
				if v92 != 0 {
					return int32(0)
				} else {
					F_jspGetRightArg(m, l2, v10+int32(40))
					mBase = m.M
					v96 = m.ExcPending
					if v96 != 0 {
						return int32(0)
					} else {
						v97 = *(*int32)(unsafe.Add(mBase, uint32(v10)+68))
						if base.Ui32(v97) < base.Ui32(int32(4)) {
							v111 = v97
							v112 = v10 + int32(80)
							v113 = v10 + int32(40)
							switch v111 - int32(1) {
							case 0:
								*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = int32(1)
								v131 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
								*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v131
								v133 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v133
							case 1:
								*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = int32(2)
								v127 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
								*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v127
							case 2:
								*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = int32(3)
								v120 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
								v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120))))
								*(*uint8)(unsafe.Add(mBase, uint32(v10)+24)) = uint8(base.B2i32(v121 != int32(0)))
							default:
								*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = int32(0)
							}
							v135 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v135
							v141 = F_extract_jsp_path_expr(m, l0, v10+int32(16), v113, v10+int32(20))
							mBase = m.M
							v142 = m.ExcPending
							if v142 != 0 {
								return int32(0)
							} else {
								v146 = v141
								m.G0 = v10 + int32(96)
								return v146
							}
						} else {
							v104 = *(*int32)(unsafe.Add(mBase, uint32(v10)+40))
							if base.Ui32(int32(3)) < base.Ui32(v104) {
								v146 = v5
								m.G0 = v10 + int32(96)
								return v146
							} else {
								v111 = v104
								v112 = v10 + int32(52)
								v113 = v10 + int32(68)
								switch v111 - int32(1) {
								case 0:
									*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = int32(1)
									v131 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
									*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v131
									v133 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v133
								case 1:
									*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = int32(2)
									v127 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
									*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v127
								case 2:
									*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = int32(3)
									v120 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
									v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120))))
									*(*uint8)(unsafe.Add(mBase, uint32(v10)+24)) = uint8(base.B2i32(v121 != int32(0)))
								default:
									*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = int32(0)
								}
								v135 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v135
								v141 = F_extract_jsp_path_expr(m, l0, v10+int32(16), v113, v10+int32(20))
								mBase = m.M
								v142 = m.ExcPending
								if v142 != 0 {
									return int32(0)
								} else {
									v146 = v141
									m.G0 = v10 + int32(96)
									return v146
								}
							}
						}
					}
				}
			}
		case 26:
			if l3 != 0 {
				v146 = v5
				m.G0 = v10 + int32(96)
				return v146
			} else {
				F_jspGetArg(m, l2, v10+int32(68))
				mBase = m.M
				v79 = m.ExcPending
				if v79 != 0 {
					return int32(0)
				} else {
					v80 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v80
					v87 = F_extract_jsp_path_expr(m, l0, v10+int32(12), v10+int32(68), int32(0))
					mBase = m.M
					v88 = m.ExcPending
					if v88 != 0 {
						return int32(0)
					} else {
						v146 = v87
						m.G0 = v10 + int32(96)
						return v146
					}
				}
			}
		}
	}
}
