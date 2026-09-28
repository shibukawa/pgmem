package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecAssignScanProjectionInfo(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
	F_ExecConditionalAssignProjectionInfo(m, l0, v5)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		return
	}
}
func F_apply_projection_to_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 float64
	_ = v34
	var v35 float64
	_ = v35
	var v37 float64
	_ = v37
	var v38 float64
	_ = v38
	var v42 float64
	_ = v42
	var v43 float64
	_ = v43
	var v45 float64
	_ = v45
	var v47 float64
	_ = v47
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	v7 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	switch v8 - int32(336) {
	case 0, 1, 3, 4, 28, 29, 30, 31, 35, 38, 39, 40, 41:
		v23 = v7
		v25 = v23
	case 2:
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		if v16 != int32(293) {
			v23 = v7
			v25 = v23
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l2)+72))
			v25 = base.B2i32(v19 == int32(0))
		}
	default:
		v23 = int32(1)
		v25 = v23
	case 23:
		v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+72)))
		v25 = int32(base.Ui32(v11&int32(4)) >> (uint(int32(2)) % 32))
	}
	if v25 == int32(0) {
		v28 = F_create_projection_path(m, l0, l1, l2, l3)
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			return v28
		}
	} else {
		v33 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
		v34 = *(*float64)(unsafe.Add(mBase, uint32(v33)+24))
		v35 = *(*float64)(unsafe.Add(mBase, uint32(v33)+16))
		*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = l3
		v37 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
		v38 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
		*(*float64)(unsafe.Add(mBase, uint32(l2)+48)) = base.F64_add(v37, base.F64_sub(v38, v35))
		v42 = *(*float64)(unsafe.Add(mBase, uint32(l2)+56))
		v43 = *(*float64)(unsafe.Add(mBase, uint32(l3)+24))
		v45 = *(*float64)(unsafe.Add(mBase, uint32(l2)+32))
		v47 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
		*(*float64)(unsafe.Add(mBase, uint32(l2)+56)) = base.F64_add(v42, base.F64_add(base.F64_mul(base.F64_sub(v43, v34), v45), base.F64_sub(v47, v35)))
		v52 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		if v52&int32(-2) != int32(298) {
			v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+21)))
			if v68 != int32(1) {
				return l2
			} else {
				v71 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
				v72 = F_is_parallel_safe(m, l0, v71)
				mBase = m.M
				v73 = m.ExcPending
				if v73 != 0 {
					return int32(0)
				} else {
					if v72 != 0 {
					} else {
						v74 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l2)+21)) = uint8(v74)
					}
					return l2
				}
			}
		} else {
			v57 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
			v58 = F_is_parallel_safe(m, l0, v57)
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return int32(0)
			} else {
				if v58 == int32(0) {
					v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+21)))
					if v68 != int32(1) {
						return l2
					} else {
						v71 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
						v72 = F_is_parallel_safe(m, l0, v71)
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return int32(0)
						} else {
							if v72 != 0 {
							} else {
								v74 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l2)+21)) = uint8(v74)
							}
							return l2
						}
					}
				} else {
					v62 = *(*int32)(unsafe.Add(mBase, uint32(l2)+72))
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
					v64 = F_create_projection_path(m, l0, v63, v62, l3)
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l2)+72)) = v64
						return l2
					}
				}
			}
		}
	}
}
func F_create_projection_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 float64
	_ = v74
	var v76 int32
	_ = v76
	var v78 float64
	_ = v78
	var v79 float64
	_ = v79
	var v80 float64
	_ = v80
	var v84 float64
	_ = v84
	var v85 float64
	_ = v85
	var v87 float64
	_ = v87
	var v89 float64
	_ = v89
	var v90 float64
	_ = v90
	var v91 float64
	_ = v91
	var v97 int32
	_ = v97
	var v99 float64
	_ = v99
	var v101 int32
	_ = v101
	var v103 float64
	_ = v103
	var v104 float64
	_ = v104
	var v108 float64
	_ = v108
	var v109 float64
	_ = v109
	var v111 float64
	_ = v111
	var v113 float64
	_ = v113
	var v114 float64
	_ = v114
	v5 = int32(0)
	v9 = F_palloc0(m, int32(80))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = int32(303)
		*(*int32)(unsafe.Add(mBase, uint32(v9))) = v13
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		if v15 == v13 {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l2)+72))
			v19 = v18
		} else {
			v19 = l2
		}
		*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l3
		*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = int32(335)
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
		v25 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)) = uint8(v25)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v24
		v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
		if v28 != int32(1) {
			v37 = v5
			*(*uint8)(unsafe.Add(mBase, uint32(v9)+21)) = uint8(v37)
			v39 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
			*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v39
			v41 = *(*int32)(unsafe.Add(mBase, uint32(v19)+64))
			*(*int32)(unsafe.Add(mBase, uint32(v9)+72)) = v19
			*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = v41
			v44 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
			v45 = int32(0)
			v46 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
			switch v46 - int32(336) {
			case 0, 1, 3, 4, 28, 29, 30, 31, 35, 38, 39, 40, 41:
				v61 = v45
				v63 = v61
			case 2:
				v54 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
				if v54 != int32(293) {
					v61 = v45
					v63 = v61
				} else {
					v57 = *(*int32)(unsafe.Add(mBase, uint32(v19)+72))
					v63 = base.B2i32(v57 == int32(0))
				}
			default:
				v61 = int32(1)
				v63 = v61
			case 23:
				v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+72)))
				v63 = int32(base.Ui32(v49&int32(4)) >> (uint(int32(2)) % 32))
			}
			if v63 == int32(0) {
				v66 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
				v67 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
				v68 = F_equal(m, v66, v67)
				mBase = m.M
				v69 = m.ExcPending
				if v69 != 0 {
					return int32(0)
				} else {
					if v68 == int32(0) {
						v97 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v9)+76)) = uint8(v97)
						v99 = *(*float64)(unsafe.Add(mBase, uint32(v19)+32))
						*(*float64)(unsafe.Add(mBase, uint32(v9)+32)) = v99
						v101 = *(*int32)(unsafe.Add(mBase, uint32(v19)+40))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = v101
						v103 = *(*float64)(unsafe.Add(mBase, uint32(v19)+48))
						v104 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
						*(*float64)(unsafe.Add(mBase, uint32(v9)+48)) = base.F64_add(v103, v104)
						v108 = *(*float64)(unsafe.Add(mBase, _c_F_create_projection_path[0]))
						v109 = *(*float64)(unsafe.Add(mBase, uint32(l3)+24))
						v111 = *(*float64)(unsafe.Add(mBase, uint32(v19)+32))
						v113 = *(*float64)(unsafe.Add(mBase, uint32(v19)+56))
						v114 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
						*(*float64)(unsafe.Add(mBase, uint32(v9)+56)) = base.F64_add(base.F64_mul(base.F64_add(v108, v109), v111), base.F64_add(v113, v114))
						return v9
					} else {
						v72 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v9)+76)) = uint8(v72)
						v74 = *(*float64)(unsafe.Add(mBase, uint32(v19)+32))
						*(*float64)(unsafe.Add(mBase, uint32(v9)+32)) = v74
						v76 = *(*int32)(unsafe.Add(mBase, uint32(v19)+40))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = v76
						v78 = *(*float64)(unsafe.Add(mBase, uint32(v19)+48))
						v79 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
						v80 = *(*float64)(unsafe.Add(mBase, uint32(v44)+16))
						*(*float64)(unsafe.Add(mBase, uint32(v9)+48)) = base.F64_add(v78, base.F64_sub(v79, v80))
						v84 = *(*float64)(unsafe.Add(mBase, uint32(l3)+24))
						v85 = *(*float64)(unsafe.Add(mBase, uint32(v44)+24))
						v87 = *(*float64)(unsafe.Add(mBase, uint32(v19)+32))
						v89 = *(*float64)(unsafe.Add(mBase, uint32(v19)+56))
						v90 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
						v91 = *(*float64)(unsafe.Add(mBase, uint32(v44)+16))
						*(*float64)(unsafe.Add(mBase, uint32(v9)+56)) = base.F64_add(base.F64_mul(base.F64_sub(v84, v85), v87), base.F64_add(v89, base.F64_sub(v90, v91)))
						return v9
					}
				}
			} else {
				v72 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v9)+76)) = uint8(v72)
				v74 = *(*float64)(unsafe.Add(mBase, uint32(v19)+32))
				*(*float64)(unsafe.Add(mBase, uint32(v9)+32)) = v74
				v76 = *(*int32)(unsafe.Add(mBase, uint32(v19)+40))
				*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = v76
				v78 = *(*float64)(unsafe.Add(mBase, uint32(v19)+48))
				v79 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
				v80 = *(*float64)(unsafe.Add(mBase, uint32(v44)+16))
				*(*float64)(unsafe.Add(mBase, uint32(v9)+48)) = base.F64_add(v78, base.F64_sub(v79, v80))
				v84 = *(*float64)(unsafe.Add(mBase, uint32(l3)+24))
				v85 = *(*float64)(unsafe.Add(mBase, uint32(v44)+24))
				v87 = *(*float64)(unsafe.Add(mBase, uint32(v19)+32))
				v89 = *(*float64)(unsafe.Add(mBase, uint32(v19)+56))
				v90 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
				v91 = *(*float64)(unsafe.Add(mBase, uint32(v44)+16))
				*(*float64)(unsafe.Add(mBase, uint32(v9)+56)) = base.F64_add(base.F64_mul(base.F64_sub(v84, v85), v87), base.F64_add(v89, base.F64_sub(v90, v91)))
				return v9
			}
		} else {
			v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+21)))
			if v31 != int32(1) {
				v37 = v5
				*(*uint8)(unsafe.Add(mBase, uint32(v9)+21)) = uint8(v37)
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
				*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v39
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v19)+64))
				*(*int32)(unsafe.Add(mBase, uint32(v9)+72)) = v19
				*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = v41
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
				v45 = int32(0)
				v46 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
				switch v46 - int32(336) {
				case 0, 1, 3, 4, 28, 29, 30, 31, 35, 38, 39, 40, 41:
					v61 = v45
					v63 = v61
				case 2:
					v54 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
					if v54 != int32(293) {
						v61 = v45
						v63 = v61
					} else {
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v19)+72))
						v63 = base.B2i32(v57 == int32(0))
					}
				default:
					v61 = int32(1)
					v63 = v61
				case 23:
					v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+72)))
					v63 = int32(base.Ui32(v49&int32(4)) >> (uint(int32(2)) % 32))
				}
				if v63 == int32(0) {
					v66 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
					v67 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
					v68 = F_equal(m, v66, v67)
					mBase = m.M
					v69 = m.ExcPending
					if v69 != 0 {
						return int32(0)
					} else {
						if v68 == int32(0) {
							v97 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v9)+76)) = uint8(v97)
							v99 = *(*float64)(unsafe.Add(mBase, uint32(v19)+32))
							*(*float64)(unsafe.Add(mBase, uint32(v9)+32)) = v99
							v101 = *(*int32)(unsafe.Add(mBase, uint32(v19)+40))
							*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = v101
							v103 = *(*float64)(unsafe.Add(mBase, uint32(v19)+48))
							v104 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
							*(*float64)(unsafe.Add(mBase, uint32(v9)+48)) = base.F64_add(v103, v104)
							v108 = *(*float64)(unsafe.Add(mBase, _c_F_create_projection_path[0]))
							v109 = *(*float64)(unsafe.Add(mBase, uint32(l3)+24))
							v111 = *(*float64)(unsafe.Add(mBase, uint32(v19)+32))
							v113 = *(*float64)(unsafe.Add(mBase, uint32(v19)+56))
							v114 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
							*(*float64)(unsafe.Add(mBase, uint32(v9)+56)) = base.F64_add(base.F64_mul(base.F64_add(v108, v109), v111), base.F64_add(v113, v114))
							return v9
						} else {
							v72 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v9)+76)) = uint8(v72)
							v74 = *(*float64)(unsafe.Add(mBase, uint32(v19)+32))
							*(*float64)(unsafe.Add(mBase, uint32(v9)+32)) = v74
							v76 = *(*int32)(unsafe.Add(mBase, uint32(v19)+40))
							*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = v76
							v78 = *(*float64)(unsafe.Add(mBase, uint32(v19)+48))
							v79 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
							v80 = *(*float64)(unsafe.Add(mBase, uint32(v44)+16))
							*(*float64)(unsafe.Add(mBase, uint32(v9)+48)) = base.F64_add(v78, base.F64_sub(v79, v80))
							v84 = *(*float64)(unsafe.Add(mBase, uint32(l3)+24))
							v85 = *(*float64)(unsafe.Add(mBase, uint32(v44)+24))
							v87 = *(*float64)(unsafe.Add(mBase, uint32(v19)+32))
							v89 = *(*float64)(unsafe.Add(mBase, uint32(v19)+56))
							v90 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
							v91 = *(*float64)(unsafe.Add(mBase, uint32(v44)+16))
							*(*float64)(unsafe.Add(mBase, uint32(v9)+56)) = base.F64_add(base.F64_mul(base.F64_sub(v84, v85), v87), base.F64_add(v89, base.F64_sub(v90, v91)))
							return v9
						}
					}
				} else {
					v72 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v9)+76)) = uint8(v72)
					v74 = *(*float64)(unsafe.Add(mBase, uint32(v19)+32))
					*(*float64)(unsafe.Add(mBase, uint32(v9)+32)) = v74
					v76 = *(*int32)(unsafe.Add(mBase, uint32(v19)+40))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = v76
					v78 = *(*float64)(unsafe.Add(mBase, uint32(v19)+48))
					v79 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
					v80 = *(*float64)(unsafe.Add(mBase, uint32(v44)+16))
					*(*float64)(unsafe.Add(mBase, uint32(v9)+48)) = base.F64_add(v78, base.F64_sub(v79, v80))
					v84 = *(*float64)(unsafe.Add(mBase, uint32(l3)+24))
					v85 = *(*float64)(unsafe.Add(mBase, uint32(v44)+24))
					v87 = *(*float64)(unsafe.Add(mBase, uint32(v19)+32))
					v89 = *(*float64)(unsafe.Add(mBase, uint32(v19)+56))
					v90 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
					v91 = *(*float64)(unsafe.Add(mBase, uint32(v44)+16))
					*(*float64)(unsafe.Add(mBase, uint32(v9)+56)) = base.F64_add(base.F64_mul(base.F64_sub(v84, v85), v87), base.F64_add(v89, base.F64_sub(v90, v91)))
					return v9
				}
			} else {
				v34 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
				v35 = F_is_parallel_safe(m, l0, v34)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					v37 = v35
					*(*uint8)(unsafe.Add(mBase, uint32(v9)+21)) = uint8(v37)
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v39
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v19)+64))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+72)) = v19
					*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = v41
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
					v45 = int32(0)
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
					switch v46 - int32(336) {
					case 0, 1, 3, 4, 28, 29, 30, 31, 35, 38, 39, 40, 41:
						v61 = v45
						v63 = v61
					case 2:
						v54 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
						if v54 != int32(293) {
							v61 = v45
							v63 = v61
						} else {
							v57 = *(*int32)(unsafe.Add(mBase, uint32(v19)+72))
							v63 = base.B2i32(v57 == int32(0))
						}
					default:
						v61 = int32(1)
						v63 = v61
					case 23:
						v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+72)))
						v63 = int32(base.Ui32(v49&int32(4)) >> (uint(int32(2)) % 32))
					}
					if v63 == int32(0) {
						v66 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
						v67 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
						v68 = F_equal(m, v66, v67)
						mBase = m.M
						v69 = m.ExcPending
						if v69 != 0 {
							return int32(0)
						} else {
							if v68 == int32(0) {
								v97 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v9)+76)) = uint8(v97)
								v99 = *(*float64)(unsafe.Add(mBase, uint32(v19)+32))
								*(*float64)(unsafe.Add(mBase, uint32(v9)+32)) = v99
								v101 = *(*int32)(unsafe.Add(mBase, uint32(v19)+40))
								*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = v101
								v103 = *(*float64)(unsafe.Add(mBase, uint32(v19)+48))
								v104 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
								*(*float64)(unsafe.Add(mBase, uint32(v9)+48)) = base.F64_add(v103, v104)
								v108 = *(*float64)(unsafe.Add(mBase, _c_F_create_projection_path[0]))
								v109 = *(*float64)(unsafe.Add(mBase, uint32(l3)+24))
								v111 = *(*float64)(unsafe.Add(mBase, uint32(v19)+32))
								v113 = *(*float64)(unsafe.Add(mBase, uint32(v19)+56))
								v114 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
								*(*float64)(unsafe.Add(mBase, uint32(v9)+56)) = base.F64_add(base.F64_mul(base.F64_add(v108, v109), v111), base.F64_add(v113, v114))
								return v9
							} else {
								v72 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v9)+76)) = uint8(v72)
								v74 = *(*float64)(unsafe.Add(mBase, uint32(v19)+32))
								*(*float64)(unsafe.Add(mBase, uint32(v9)+32)) = v74
								v76 = *(*int32)(unsafe.Add(mBase, uint32(v19)+40))
								*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = v76
								v78 = *(*float64)(unsafe.Add(mBase, uint32(v19)+48))
								v79 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
								v80 = *(*float64)(unsafe.Add(mBase, uint32(v44)+16))
								*(*float64)(unsafe.Add(mBase, uint32(v9)+48)) = base.F64_add(v78, base.F64_sub(v79, v80))
								v84 = *(*float64)(unsafe.Add(mBase, uint32(l3)+24))
								v85 = *(*float64)(unsafe.Add(mBase, uint32(v44)+24))
								v87 = *(*float64)(unsafe.Add(mBase, uint32(v19)+32))
								v89 = *(*float64)(unsafe.Add(mBase, uint32(v19)+56))
								v90 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
								v91 = *(*float64)(unsafe.Add(mBase, uint32(v44)+16))
								*(*float64)(unsafe.Add(mBase, uint32(v9)+56)) = base.F64_add(base.F64_mul(base.F64_sub(v84, v85), v87), base.F64_add(v89, base.F64_sub(v90, v91)))
								return v9
							}
						}
					} else {
						v72 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v9)+76)) = uint8(v72)
						v74 = *(*float64)(unsafe.Add(mBase, uint32(v19)+32))
						*(*float64)(unsafe.Add(mBase, uint32(v9)+32)) = v74
						v76 = *(*int32)(unsafe.Add(mBase, uint32(v19)+40))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = v76
						v78 = *(*float64)(unsafe.Add(mBase, uint32(v19)+48))
						v79 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
						v80 = *(*float64)(unsafe.Add(mBase, uint32(v44)+16))
						*(*float64)(unsafe.Add(mBase, uint32(v9)+48)) = base.F64_add(v78, base.F64_sub(v79, v80))
						v84 = *(*float64)(unsafe.Add(mBase, uint32(l3)+24))
						v85 = *(*float64)(unsafe.Add(mBase, uint32(v44)+24))
						v87 = *(*float64)(unsafe.Add(mBase, uint32(v19)+32))
						v89 = *(*float64)(unsafe.Add(mBase, uint32(v19)+56))
						v90 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
						v91 = *(*float64)(unsafe.Add(mBase, uint32(v44)+16))
						*(*float64)(unsafe.Add(mBase, uint32(v9)+56)) = base.F64_add(base.F64_mul(base.F64_sub(v84, v85), v87), base.F64_add(v89, base.F64_sub(v90, v91)))
						return v9
					}
				}
			}
		}
	}
}
