package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__jumbleDropTableSpaceStmt(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v4 != 0 {
		v5 = F_strlen(m, v4)
		mBase = m.M
		F_AppendJumble(m, l0, v4, v5+int32(1))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			F_AppendJumble8(m, l0, l1+int32(8))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				return
			}
		}
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v10 + int32(1)
		F_AppendJumble8(m, l0, l1+int32(8))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			return
		}
	}
}
func F__jumbleRangeTableSample(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	F__jumbleNode(m, l0, v3)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		F__jumbleNode(m, l0, v6)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			F__jumbleNode(m, l0, v9)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return
			} else {
				v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
				F__jumbleNode(m, l0, v12)
				mBase = m.M
				v14 = m.ExcPending
				if v14 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F__jumbleRelabelType(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	F__jumbleNode(m, l0, v3)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		F_AppendJumble32(m, l0, l1+int32(8))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			return
		}
	}
}
func F__jumbleResTarget(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v4 != 0 {
		v5 = F_strlen(m, v4)
		mBase = m.M
		F_AppendJumble(m, l0, v4, v5+int32(1))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			F__jumbleNode(m, l0, v14)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				F__jumbleNode(m, l0, v17)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					return
				}
			}
		}
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v10 + int32(1)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		F__jumbleNode(m, l0, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			F__jumbleNode(m, l0, v17)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F__jumbleRoleSpec(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	F_AppendJumble32(m, l0, l1+int32(4))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		if v7 != 0 {
			v8 = F_strlen(m, v7)
			mBase = m.M
			F_AppendJumble(m, l0, v7, v8+int32(1))
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				return
			}
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v13 + int32(1)
			return
		}
	}
}
func F__jumbleUpdateStmt(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	F__jumbleNode(m, l0, v3)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		F__jumbleNode(m, l0, v6)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			F__jumbleNode(m, l0, v9)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return
			} else {
				v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
				F__jumbleNode(m, l0, v12)
				mBase = m.M
				v14 = m.ExcPending
				if v14 != 0 {
					return
				} else {
					v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
					F__jumbleNode(m, l0, v15)
					mBase = m.M
					v17 = m.ExcPending
					if v17 != 0 {
						return
					} else {
						v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
						F__jumbleNode(m, l0, v18)
						mBase = m.M
						v20 = m.ExcPending
						if v20 != 0 {
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
