package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_extract_jsp_query(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	v6 = int32(0)
	v8 = m.G0
	v10 = v8 + int32(-64)
	m.G0 = v10
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v6
	*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = int64(0)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if l2 != 0 {
		v21 = int32(1321)
	} else {
		v21 = int32(1322)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v10)+56)) = v21
	if l2 != 0 {
		v25 = int32(1323)
	} else {
		v25 = int32(1324)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = v25
	v28 = int32(base.Ui32(v18) >> (uint(int32(31)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+60)) = uint8(v28)
	v31 = v8 + int32(-40)
	F_jspInit(m, v31, l0)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		return int32(0)
	} else {
		if l1 == int32(15) {
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
			*(*int32)(unsafe.Add(mBase, uint32(v10))) = v38
			v43 = F_extract_jsp_path_expr(m, v8+int32(-12), v10, v31, int32(0))
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return int32(0)
			} else {
				v56 = v43
				if v56 == int32(0) {
					v59 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v59
					v79 = v59
					m.G0 = v10 - int32(-64)
					return v79
				} else {
					F_emit_jsp_gin_entries(m, v56, v8+int32(-56))
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return int32(0)
					} else {
						v66 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v66
						v68 = int32(0)
						if v66 == v68 {
							v79 = v68
							m.G0 = v10 - int32(-64)
							return v79
						} else {
							v73 = F_palloc0(m, v66<<(uint(int32(2))%32))
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l4))) = v73
								*(*int32)(unsafe.Add(mBase, uint32(v73))) = v56
								v77 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
								v79 = v77
								m.G0 = v10 - int32(-64)
								return v79
							}
						}
					}
				}
			}
		} else {
			v45 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
			*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v45
			v54 = F_extract_jsp_bool_expr(m, v8+int32(-12), v8+int32(-60), v8+int32(-40), int32(0))
			mBase = m.M
			v55 = m.ExcPending
			if v55 != 0 {
				return int32(0)
			} else {
				v56 = v54
				if v56 == int32(0) {
					v59 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v59
					v79 = v59
					m.G0 = v10 - int32(-64)
					return v79
				} else {
					F_emit_jsp_gin_entries(m, v56, v8+int32(-56))
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return int32(0)
					} else {
						v66 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v66
						v68 = int32(0)
						if v66 == v68 {
							v79 = v68
							m.G0 = v10 - int32(-64)
							return v79
						} else {
							v73 = F_palloc0(m, v66<<(uint(int32(2))%32))
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l4))) = v73
								*(*int32)(unsafe.Add(mBase, uint32(v73))) = v56
								v77 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
								v79 = v77
								m.G0 = v10 - int32(-64)
								return v79
							}
						}
					}
				}
			}
		}
	}
}
func F_jspGetNext(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	v3 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.B2i32(l1 == v3)|base.B2i32(v6 <= v3) == v3 {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		F_jspInitByBuffer(m, l1, v12, v6)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			return base.B2i32(int32(0) < v6)
		}
	} else {
		return base.B2i32(int32(0) < v6)
	}
}
func F_jspGetString(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	if l1 != 0 {
		v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v3
	} else {
	}
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	return v5
}
