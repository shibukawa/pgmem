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
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	v6 = int32(0)
	v8 = m.G0
	v10 = v8 + int32(-64)
	m.G0 = v10
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v6
	*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = int64(0)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if l2 != 0 {
		v21 = int32(1337)
	} else {
		v21 = int32(1338)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v10)+56)) = v21
	if l2 != 0 {
		v25 = int32(1339)
	} else {
		v25 = int32(1340)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = v25
	v28 = int32(base.Ui32(v18) >> (uint(int32(31)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+60)) = uint8(v28)
	F_jspInit(m, v8+int32(-40), l0)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		return int32(0)
	} else {
		if l1 == int32(15) {
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
			*(*int32)(unsafe.Add(mBase, uint32(v10))) = v38
			v45 = F_extract_jsp_path_expr(m, v8+int32(-12), v10, v8+int32(-40), int32(0))
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return int32(0)
			} else {
				v58 = v45
				if v58 == int32(0) {
					v61 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v61
					v81 = v61
					m.G0 = v10 - int32(-64)
					return v81
				} else {
					F_emit_jsp_gin_entries(m, v58, v8+int32(-56))
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return int32(0)
					} else {
						v68 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v68
						v70 = int32(0)
						if v68 == v70 {
							v81 = v70
							m.G0 = v10 - int32(-64)
							return v81
						} else {
							v75 = F_palloc0(m, v68<<(uint(int32(2))%32))
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l4))) = v75
								*(*int32)(unsafe.Add(mBase, uint32(v75))) = v58
								v79 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
								v81 = v79
								m.G0 = v10 - int32(-64)
								return v81
							}
						}
					}
				}
			}
		} else {
			v47 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
			*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v47
			v56 = F_extract_jsp_bool_expr(m, v8+int32(-12), v8+int32(-60), v8+int32(-40), int32(0))
			mBase = m.M
			v57 = m.ExcPending
			if v57 != 0 {
				return int32(0)
			} else {
				v58 = v56
				if v58 == int32(0) {
					v61 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v61
					v81 = v61
					m.G0 = v10 - int32(-64)
					return v81
				} else {
					F_emit_jsp_gin_entries(m, v58, v8+int32(-56))
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return int32(0)
					} else {
						v68 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v68
						v70 = int32(0)
						if v68 == v70 {
							v81 = v70
							m.G0 = v10 - int32(-64)
							return v81
						} else {
							v75 = F_palloc0(m, v68<<(uint(int32(2))%32))
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l4))) = v75
								*(*int32)(unsafe.Add(mBase, uint32(v75))) = v58
								v79 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
								v81 = v79
								m.G0 = v10 - int32(-64)
								return v81
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
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if l1 == int32(0) {
		return base.B2i32(int32(0) < v4)
	} else {
		if v4 <= int32(0) {
			return base.B2i32(int32(0) < v4)
		} else {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			F_jspInitByBuffer(m, l1, v9, v4)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				return base.B2i32(int32(0) < v4)
			}
		}
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
