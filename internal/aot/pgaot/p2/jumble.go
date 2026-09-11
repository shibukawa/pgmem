package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__jumbleAlias(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v4 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F__jumbleNode(m, l0, v70)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L22
	} else {
		goto L24
	}
L2:
	;
	if v4&int32(3) == int32(0) {
		v28 = v4
		goto L7
	} else {
		goto L8
	}
L3:
	;
	goto L4
L4:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v66 + int32(1)
	goto L1
L5:
	;
	F_AppendJumble(m, l0, v4, v61+int32(1))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L22
	} else {
		goto L23
	}
L6:
	;
	v61 = v53 - v4
	goto L5
L7:
	;
	v32 = v28
	goto L16
L8:
	;
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4))))
	if v12 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v61 = int32(0)
	goto L5
L10:
	;
	goto L11
L11:
	;
	v17 = v4
	goto L12
L12:
	;
	v21 = v17 + int32(1)
	if v21&int32(3) == int32(0) {
		v28 = v21
		goto L7
	} else {
		goto L14
	}
L13:
	;
	v53 = v21
	goto L6
L14:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v26 != 0 {
		v17 = v21
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v41 = int32(-2139062144)
	if (int32(16843008)-v38|v38)&v41 == v41 {
		v32 = v32 + int32(4)
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v47 = v32
	goto L19
L18:
	;
	goto L17
L19:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
	if v51 != 0 {
		v47 = v47 + int32(1)
		goto L19
	} else {
		goto L21
	}
L20:
	;
	v53 = v47
	goto L6
L21:
	;
	goto L20
L22:
	;
	return
L23:
	;
	goto L1
L24:
	;
	return
}
func F__jumbleAlterTableSpaceOptionsStmt(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v4 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F__jumbleNode(m, l0, v70)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L22
	} else {
		goto L24
	}
L2:
	;
	if v4&int32(3) == int32(0) {
		v28 = v4
		goto L7
	} else {
		goto L8
	}
L3:
	;
	goto L4
L4:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v66 + int32(1)
	goto L1
L5:
	;
	F_AppendJumble(m, l0, v4, v61+int32(1))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L22
	} else {
		goto L23
	}
L6:
	;
	v61 = v53 - v4
	goto L5
L7:
	;
	v32 = v28
	goto L16
L8:
	;
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4))))
	if v12 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v61 = int32(0)
	goto L5
L10:
	;
	goto L11
L11:
	;
	v17 = v4
	goto L12
L12:
	;
	v21 = v17 + int32(1)
	if v21&int32(3) == int32(0) {
		v28 = v21
		goto L7
	} else {
		goto L14
	}
L13:
	;
	v53 = v21
	goto L6
L14:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v26 != 0 {
		v17 = v21
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v41 = int32(-2139062144)
	if (int32(16843008)-v38|v38)&v41 == v41 {
		v32 = v32 + int32(4)
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v47 = v32
	goto L19
L18:
	;
	goto L17
L19:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
	if v51 != 0 {
		v47 = v47 + int32(1)
		goto L19
	} else {
		goto L21
	}
L20:
	;
	v53 = v47
	goto L6
L21:
	;
	goto L20
L22:
	;
	return
L23:
	;
	goto L1
L24:
	;
	F_AppendJumble8(m, l0, l1+int32(12))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	return
}
func F__jumbleJsonReturning(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
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
			F_AppendJumble32(m, l0, l1+int32(12))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F__jumbleJsonValueExpr(m *base.Module, l0 int32, l1 int32) {
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
				return
			}
		}
	}
}
func F__jumblePartitionCmd(m *base.Module, l0 int32, l1 int32) {
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
	var v12 int32
	_ = v12
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
			F_AppendJumble8(m, l0, l1+int32(12))
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				return
			}
		}
	}
}
