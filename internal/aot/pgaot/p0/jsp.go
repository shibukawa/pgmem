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
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
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
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
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
			v20 = v10 + int32(68)
			F_jspGetArg(m, l2, v20)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v23
				v27 = F_extract_jsp_bool_expr(m, l0, v10+int32(4), v20, l3)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					F_jspGetRightArg(m, l2, v20)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = v31
						v33 = F_extract_jsp_bool_expr(m, l0, v10, v20, l3)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							v35 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
							if v33 != 0 {
								v37 = v27
							} else {
								v37 = int32(0)
							}
							if v37 == int32(0) {
								if v27 != 0 {
									v40 = v27
								} else {
									v40 = v33
								}
								if v35 != int32(5) {
									v44 = v40
								} else {
									v44 = int32(0)
								}
								v135 = v44
								m.G0 = v10 + int32(96)
								return v135
							} else {
								v46 = F_palloc(m, int32(16))
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v46)+12)) = v33
									*(*int32)(unsafe.Add(mBase, uint32(v46)+8)) = v27
									*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = int32(2)
									*(*int32)(unsafe.Add(mBase, uint32(v46))) = l3 ^ base.B2i32(v35 == int32(4))
									v135 = v46
									m.G0 = v10 + int32(96)
									return v135
								}
							}
						}
					}
				}
			}
		case 2:
			v57 = v10 + int32(68)
			F_jspGetArg(m, l2, v57)
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return int32(0)
			} else {
				v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v60
				v66 = F_extract_jsp_bool_expr(m, l0, v10+int32(8), v57, l3^int32(1))
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
					return int32(0)
				} else {
					v135 = v66
					m.G0 = v10 + int32(96)
					return v135
				}
			}
		default:
			v135 = v5
			m.G0 = v10 + int32(96)
			return v135
		case 4:
			if l3 != 0 {
				v135 = v5
				m.G0 = v10 + int32(96)
				return v135
			} else {
				F_jspGetArg(m, l2, v10+int32(68))
				mBase = m.M
				v82 = m.ExcPending
				if v82 != 0 {
					return int32(0)
				} else {
					v84 = v10 + int32(40)
					F_jspGetRightArg(m, l2, v84)
					mBase = m.M
					v86 = m.ExcPending
					if v86 != 0 {
						return int32(0)
					} else {
						v87 = *(*int32)(unsafe.Add(mBase, uint32(v10)+68))
						if base.Ui32(v87) < base.Ui32(int32(4)) {
							v99 = v87
							v100 = v10 + int32(80)
							v101 = v84
							switch v99 - int32(1) {
							case 0:
								*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = int32(1)
								v119 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
								*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v119
								v121 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v121
							case 1:
								*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = int32(2)
								v115 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
								*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v115
							case 2:
								*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = int32(3)
								v108 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
								v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108))))
								*(*uint8)(unsafe.Add(mBase, uint32(v10)+24)) = uint8(base.B2i32(v109 != int32(0)))
							default:
								*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = int32(0)
							}
							v123 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v123
							v129 = F_extract_jsp_path_expr(m, l0, v10+int32(16), v101, v10+int32(20))
							mBase = m.M
							v130 = m.ExcPending
							if v130 != 0 {
								return int32(0)
							} else {
								v135 = v129
								m.G0 = v10 + int32(96)
								return v135
							}
						} else {
							v92 = *(*int32)(unsafe.Add(mBase, uint32(v10)+40))
							if base.Ui32(int32(3)) < base.Ui32(v92) {
								v135 = v5
								m.G0 = v10 + int32(96)
								return v135
							} else {
								v99 = v92
								v100 = v10 + int32(52)
								v101 = v10 + int32(68)
								switch v99 - int32(1) {
								case 0:
									*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = int32(1)
									v119 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
									*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v119
									v121 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v121
								case 1:
									*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = int32(2)
									v115 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
									*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v115
								case 2:
									*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = int32(3)
									v108 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
									v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108))))
									*(*uint8)(unsafe.Add(mBase, uint32(v10)+24)) = uint8(base.B2i32(v109 != int32(0)))
								default:
									*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = int32(0)
								}
								v123 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v123
								v129 = F_extract_jsp_path_expr(m, l0, v10+int32(16), v101, v10+int32(20))
								mBase = m.M
								v130 = m.ExcPending
								if v130 != 0 {
									return int32(0)
								} else {
									v135 = v129
									m.G0 = v10 + int32(96)
									return v135
								}
							}
						}
					}
				}
			}
		case 26:
			if l3 != 0 {
				v135 = v5
				m.G0 = v10 + int32(96)
				return v135
			} else {
				v69 = v10 + int32(68)
				F_jspGetArg(m, l2, v69)
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return int32(0)
				} else {
					v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v72
					v77 = F_extract_jsp_path_expr(m, l0, v10+int32(12), v69, int32(0))
					mBase = m.M
					v78 = m.ExcPending
					if v78 != 0 {
						return int32(0)
					} else {
						v135 = v77
						m.G0 = v10 + int32(96)
						return v135
					}
				}
			}
		}
	}
}
