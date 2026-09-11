package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__jumbleCreateSeqStmt(m *base.Module, l0 int32, l1 int32) {
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
	var v16 int32
	_ = v16
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
			F_AppendJumble32(m, l0, l1+int32(12))
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				F_AppendJumble8(m, l0, l1+int32(16))
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return
				} else {
					F_AppendJumble8(m, l0, l1+int32(17))
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
func F__jumbleCreateUserMappingStmt(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	F__jumbleNode(m, l0, v4)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v7 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	F_AppendJumble8(m, l0, l1+int32(12))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L25
	}
L4:
	;
	if v7&int32(3) == int32(0) {
		v31 = v7
		goto L9
	} else {
		goto L10
	}
L5:
	;
	goto L6
L6:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v69 + int32(1)
	goto L3
L7:
	;
	F_AppendJumble(m, l0, v7, v64+int32(1))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L24
	}
L8:
	;
	v64 = v56 - v7
	goto L7
L9:
	;
	v35 = v31
	goto L18
L10:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v15 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v64 = int32(0)
	goto L7
L12:
	;
	goto L13
L13:
	;
	v20 = v7
	goto L14
L14:
	;
	v24 = v20 + int32(1)
	if v24&int32(3) == int32(0) {
		v31 = v24
		goto L9
	} else {
		goto L16
	}
L15:
	;
	v56 = v24
	goto L8
L16:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v29 != 0 {
		v20 = v24
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v44 = int32(-2139062144)
	if (int32(16843008)-v41|v41)&v44 == v44 {
		v35 = v35 + int32(4)
		goto L18
	} else {
		goto L20
	}
L19:
	;
	v50 = v35
	goto L21
L20:
	;
	goto L19
L21:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	if v54 != 0 {
		v50 = v50 + int32(1)
		goto L21
	} else {
		goto L23
	}
L22:
	;
	v56 = v50
	goto L8
L23:
	;
	goto L22
L24:
	;
	goto L3
L25:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	F__jumbleNode(m, l0, v77)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	return
}
func F__jumbleJsonArrayQueryConstructor(m *base.Module, l0 int32, l1 int32) {
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
	var v15 int32
	_ = v15
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
				F_AppendJumble8(m, l0, l1+int32(16))
				mBase = m.M
				v15 = m.ExcPending
				if v15 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F__jumbleJsonIsPredicate(m *base.Module, l0 int32, l1 int32) {
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
	var v16 int32
	_ = v16
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
			F_AppendJumble32(m, l0, l1+int32(12))
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				F_AppendJumble8(m, l0, l1+int32(16))
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F__jumbleJsonTablePath(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	F__jumbleNode(m, l0, v3)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v6 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	if v6&int32(3) == int32(0) {
		v30 = v6
		goto L8
	} else {
		goto L9
	}
L4:
	;
	goto L5
L5:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v68 + int32(1)
	return
L6:
	;
	F_AppendJumble(m, l0, v6, v63+int32(1))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L23
	}
L7:
	;
	v63 = v55 - v6
	goto L6
L8:
	;
	v34 = v30
	goto L17
L9:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	if v14 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v63 = int32(0)
	goto L6
L11:
	;
	goto L12
L12:
	;
	v19 = v6
	goto L13
L13:
	;
	v23 = v19 + int32(1)
	if v23&int32(3) == int32(0) {
		v30 = v23
		goto L8
	} else {
		goto L15
	}
L14:
	;
	v55 = v23
	goto L7
L15:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	if v28 != 0 {
		v19 = v23
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v43 = int32(-2139062144)
	if (int32(16843008)-v40|v40)&v43 == v43 {
		v34 = v34 + int32(4)
		goto L17
	} else {
		goto L19
	}
L18:
	;
	v49 = v34
	goto L20
L19:
	;
	goto L18
L20:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	if v53 != 0 {
		v49 = v49 + int32(1)
		goto L20
	} else {
		goto L22
	}
L21:
	;
	v55 = v49
	goto L7
L22:
	;
	goto L21
L23:
	;
	return
}
func F__jumblePLAssignStmt(m *base.Module, l0 int32, l1 int32) {
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
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
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
	F_AppendJumble32(m, l0, l1+int32(12))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	F__jumbleNode(m, l0, v77)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L22
	} else {
		goto L26
	}
L26:
	;
	return
}
func F__jumbleTableSampleClause(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	F_AppendJumble32(m, l0, l1+int32(4))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		F__jumbleNode(m, l0, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			F__jumbleNode(m, l0, v10)
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
func F__jumbleVariableShowStmt(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v3 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if v3&int32(3) == int32(0) {
		v27 = v3
		goto L6
	} else {
		goto L7
	}
L2:
	;
	goto L3
L3:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v65 + int32(1)
	return
L4:
	;
	F_AppendJumble(m, l0, v3, v60+int32(1))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L21
	} else {
		goto L22
	}
L5:
	;
	v60 = v52 - v3
	goto L4
L6:
	;
	v31 = v27
	goto L15
L7:
	;
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3))))
	if v11 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v60 = int32(0)
	goto L4
L9:
	;
	goto L10
L10:
	;
	v16 = v3
	goto L11
L11:
	;
	v20 = v16 + int32(1)
	if v20&int32(3) == int32(0) {
		v27 = v20
		goto L6
	} else {
		goto L13
	}
L12:
	;
	v52 = v20
	goto L5
L13:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v25 != 0 {
		v16 = v20
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	v40 = int32(-2139062144)
	if (int32(16843008)-v37|v37)&v40 == v40 {
		v31 = v31 + int32(4)
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v46 = v31
	goto L18
L17:
	;
	goto L16
L18:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if v50 != 0 {
		v46 = v46 + int32(1)
		goto L18
	} else {
		goto L20
	}
L19:
	;
	v52 = v46
	goto L5
L20:
	;
	goto L19
L21:
	;
	return
L22:
	;
	return
}
