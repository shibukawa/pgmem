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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	v5 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(112)
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
			v20 = v10 + int32(24)
			F_jspGetArg(m, l2, v20)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v23
				v27 = F_extract_jsp_bool_expr(m, l0, v10+int32(8), v20, l3)
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
						*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v31
						v35 = F_extract_jsp_bool_expr(m, l0, v10+int32(4), v20, l3)
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int32(0)
						} else {
							v37 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
							if v35 != 0 {
								v39 = v27
							} else {
								v39 = int32(0)
							}
							if v39 == int32(0) {
								if v27 != 0 {
									v42 = v27
								} else {
									v42 = v35
								}
								if v37 != int32(5) {
									v46 = v42
								} else {
									v46 = int32(0)
								}
								v137 = v46
								m.G0 = v10 + int32(112)
								return v137
							} else {
								v48 = F_palloc(m, int32(24))
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v48)+20)) = v35
									*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v27
									*(*int32)(unsafe.Add(mBase, uint32(v48)+8)) = int32(2)
									*(*int32)(unsafe.Add(mBase, uint32(v48))) = l3 ^ base.B2i32(v37 == int32(4))
									v137 = v48
									m.G0 = v10 + int32(112)
									return v137
								}
							}
						}
					}
				}
			}
		case 2:
			v59 = v10 + int32(24)
			F_jspGetArg(m, l2, v59)
			mBase = m.M
			v61 = m.ExcPending
			if v61 != 0 {
				return int32(0)
			} else {
				v62 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v62
				v68 = F_extract_jsp_bool_expr(m, l0, v10+int32(12), v59, l3^int32(1))
				mBase = m.M
				v69 = m.ExcPending
				if v69 != 0 {
					return int32(0)
				} else {
					v137 = v68
					m.G0 = v10 + int32(112)
					return v137
				}
			}
		default:
			v137 = v5
			m.G0 = v10 + int32(112)
			return v137
		case 4:
			if l3 != 0 {
				v137 = v5
				m.G0 = v10 + int32(112)
				return v137
			} else {
				F_jspGetArg(m, l2, v10+int32(84))
				mBase = m.M
				v84 = m.ExcPending
				if v84 != 0 {
					return int32(0)
				} else {
					v86 = v10 + int32(56)
					F_jspGetRightArg(m, l2, v86)
					mBase = m.M
					v88 = m.ExcPending
					if v88 != 0 {
						return int32(0)
					} else {
						v89 = *(*int32)(unsafe.Add(mBase, uint32(v10)+84))
						if base.Ui32(v89) < base.Ui32(int32(4)) {
							v101 = v89
							v102 = v10 + int32(96)
							v103 = v86
							switch v101 - int32(1) {
							case 0:
								*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = int32(1)
								v121 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
								*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v121
								v123 = *(*int32)(unsafe.Add(mBase, uint32(v102)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v123
							case 1:
								*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = int32(2)
								v117 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
								*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v117
							case 2:
								*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = int32(3)
								v110 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
								v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110))))
								*(*uint8)(unsafe.Add(mBase, uint32(v10)+32)) = uint8(base.B2i32(v111 != int32(0)))
							default:
								*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = int32(0)
							}
							v125 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v125
							v131 = F_extract_jsp_path_expr(m, l0, v10+int32(20), v103, v10+int32(24))
							mBase = m.M
							v132 = m.ExcPending
							if v132 != 0 {
								return int32(0)
							} else {
								v137 = v131
								m.G0 = v10 + int32(112)
								return v137
							}
						} else {
							v94 = *(*int32)(unsafe.Add(mBase, uint32(v10)+56))
							if base.Ui32(int32(3)) < base.Ui32(v94) {
								v137 = v5
								m.G0 = v10 + int32(112)
								return v137
							} else {
								v101 = v94
								v102 = v10 + int32(68)
								v103 = v10 + int32(84)
								switch v101 - int32(1) {
								case 0:
									*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = int32(1)
									v121 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
									*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v121
									v123 = *(*int32)(unsafe.Add(mBase, uint32(v102)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v123
								case 1:
									*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = int32(2)
									v117 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
									*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v117
								case 2:
									*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = int32(3)
									v110 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
									v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110))))
									*(*uint8)(unsafe.Add(mBase, uint32(v10)+32)) = uint8(base.B2i32(v111 != int32(0)))
								default:
									*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = int32(0)
								}
								v125 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v125
								v131 = F_extract_jsp_path_expr(m, l0, v10+int32(20), v103, v10+int32(24))
								mBase = m.M
								v132 = m.ExcPending
								if v132 != 0 {
									return int32(0)
								} else {
									v137 = v131
									m.G0 = v10 + int32(112)
									return v137
								}
							}
						}
					}
				}
			}
		case 26:
			if l3 != 0 {
				v137 = v5
				m.G0 = v10 + int32(112)
				return v137
			} else {
				v71 = v10 + int32(24)
				F_jspGetArg(m, l2, v71)
				mBase = m.M
				v73 = m.ExcPending
				if v73 != 0 {
					return int32(0)
				} else {
					v74 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v74
					v79 = F_extract_jsp_path_expr(m, l0, v10+int32(16), v71, int32(0))
					mBase = m.M
					v80 = m.ExcPending
					if v80 != 0 {
						return int32(0)
					} else {
						v137 = v79
						m.G0 = v10 + int32(112)
						return v137
					}
				}
			}
		}
	}
}
