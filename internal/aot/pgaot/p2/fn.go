package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_get_fn_expr_arg_stable(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	v3 = int32(0)
	if l0 == v3 {
		v48 = v3
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		if v7 == int32(0) {
			v48 = v3
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
			v12 = v10 - int32(11)
			if base.Ui32(int32(9)) < base.Ui32(v12) {
				v48 = v3
			} else {
				if int32(base.Ui32(int32(977))>>(uint(v12)%32))&int32(1) == int32(0) {
					v48 = v3
				} else {
					if l1 < int32(0) {
						v48 = v3
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v12<<(uint(int32(2))%32))+uint32(_consts[1104])))
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v7+v27)))
						if v29 == int32(0) {
							v48 = v3
						} else {
							v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
							if v32 <= l1 {
								v48 = v3
							} else {
								v34 = int32(1)
								v35 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
								v39 = *(*int32)(unsafe.Add(mBase, uint32(v35+l1<<(uint(int32(2))%32))))
								v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
								switch v40 - int32(7) {
								case 0:
									v48 = v34
								case 1:
									v43 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
									if v43 == int32(0) {
										v48 = v34
									} else {
										v48 = int32(0)
									}
								default:
									v48 = int32(0)
								}
							}
						}
					}
				}
			}
		}
	}
	return v48
}
func F_has_fn_opclass_options(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	v2 = int32(0)
	if l0 == v2 {
		v18 = v2
	} else {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		if v5 == int32(0) {
			v18 = v2
		} else {
			v8 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
			if v8 != int32(7) {
				v18 = v2
			} else {
				v11 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
				if v11 != int32(17) {
					v18 = v2
				} else {
					v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+24)))
					v18 = v14 ^ int32(1)
				}
			}
		}
	}
	return v18 & int32(1)
}
func F_make_fn_arguments(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	v5 = int32(0)
	if l1 == v5 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v12 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v20 = v5
	goto L4
L4:
	;
	v25 = v20 << (uint(int32(2)) % 32)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l2+v25)))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l3+v25)))
	if v27 == v29 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L1
L6:
	;
	v55 = v20 + int32(1)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v55 < v56 {
		v20 = v55
		goto L4
	} else {
		goto L14
	}
L7:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v32 = v31 + v25
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	if v34 == int32(16) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	v38 = int32(-1)
	v42 = F_coerce_type(m, l0, v37, v27, v29, v38, int32(0), int32(2), v38)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v45 = int32(-1)
	v49 = F_coerce_type(m, l0, v33, v27, v29, v45, int32(0), int32(2), v45)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L11
	} else {
		goto L13
	}
L11:
	;
	return
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v42
	goto L6
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = v49
	goto L6
L14:
	;
	goto L5
}
