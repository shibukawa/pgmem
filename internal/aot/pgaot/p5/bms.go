package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_bms_equal(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	v3 = int32(0)
	if base.B2i32(l0 == v3)|base.B2i32(l1 == v3) != 0 {
		v49 = base.B2i32(l0|l1 == v3)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v49
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v17 != v18 {
		v49 = int32(0)
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v20 = int32(1)
	if v17 <= v20 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v23 = v20
	goto L6
L5:
	;
	v23 = v17
	goto L6
L6:
	;
	v24 = int32(8)
	v29 = int32(0)
	goto L7
L7:
	;
	v37 = v29 << (uint(int32(2)) % 32)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0+v24+v37)))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1+v24+v37)))
	v42 = base.B2i32(v39 == v41)
	if v39 != v41 {
		v49 = v42
		goto L1
	} else {
		goto L9
	}
L8:
	;
	v49 = v42
	goto L1
L9:
	;
	v45 = v29 + int32(1)
	if v45 != v23 {
		v29 = v45
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L8
}
func F_bms_membership(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	v2 = int32(0)
	if l0 == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v11 = int32(1)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v12 <= v11 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v15 = v11
	goto L6
L5:
	;
	v15 = v12
	goto L6
L6:
	;
	v19 = int32(0)
	v21 = v2
	goto L7
L7:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(8)+v19<<(uint(int32(2))%32))))
	if v28 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	return v40
L9:
	;
	goto L8
L10:
	;
	v29 = int32(2)
	if v21 != 0 {
		v40 = v29
		goto L9
	} else {
		goto L13
	}
L11:
	;
	v35 = v21
	goto L12
L12:
	;
	v37 = v19 + int32(1)
	if v37 != v15 {
		v19 = v37
		v21 = v35
		goto L7
	} else {
		goto L15
	}
L13:
	;
	v30 = int32(1)
	if base.Ui32(v30) < base.Ui32(base.I32_popcnt(v28)) {
		v40 = v29
		goto L9
	} else {
		goto L14
	}
L14:
	;
	v35 = v30
	goto L12
L15:
	;
	v40 = v35
	goto L9
}
func F_bms_num_members(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	v2 = int32(0)
	if l0 == v2 {
		return int32(0)
	} else {
		v10 = int32(1)
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v11 <= v10 {
			v14 = v10
		} else {
			v14 = v11
		}
		v18 = int32(0)
		v20 = v2
		for {
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(8)+v18<<(uint(int32(2))%32))))
			if v26 != 0 {
				v29 = v20 + base.I32_popcnt(v26)
			} else {
				v29 = v20
			}
			v31 = v18 + int32(1)
			if v31 != v14 {
				v18 = v31
				v20 = v29
				continue
			} else {
				break
			}
			break
		}
		return v29
	}
}
func F_bms_subset_compare(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	if l0 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return base.B2i32(l1 != int32(0))
L2:
	;
	goto L3
L3:
	;
	if l1 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(2)
L5:
	;
	goto L6
L6:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v21 < v22 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v24 = v21
	goto L9
L8:
	;
	v24 = v22
	goto L9
L9:
	;
	if v24 <= int32(1) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v27 = int32(1)
	goto L12
L11:
	;
	v27 = v24
	goto L12
L12:
	;
	v28 = int32(8)
	v32 = int32(0)
	v34 = v32
	v35 = v32
	goto L15
L13:
	;
	return int32(3)
L14:
	;
	return v86
L15:
	;
	v45 = v35 << (uint(int32(2)) % 32)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0+v28+v45)))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v45+(l1+v28))))
	if v47&(v49^int32(-1)) != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	if v22 < v21 {
		goto L25
	} else {
		goto L26
	}
L17:
	;
	v71 = v35 + int32(1)
	if v71 != v27 {
		v34 = v69
		v35 = v71
		goto L15
	} else {
		goto L24
	}
L18:
	;
	if base.B2i32(v34 == int32(1))|v49&(v47^int32(-1)) != 0 {
		v86 = int32(3)
		goto L14
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	if v49&(v47^int32(-1)) == int32(0) {
		v69 = v34
		goto L17
	} else {
		goto L22
	}
L21:
	;
	v69 = int32(2)
	goto L17
L22:
	;
	if v34 == int32(2) {
		goto L13
	} else {
		goto L23
	}
L23:
	;
	v69 = int32(1)
	goto L17
L24:
	;
	goto L16
L25:
	;
	if v69 == int32(1) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	if v22 <= v21 {
		v86 = v69
		goto L14
	} else {
		goto L31
	}
L28:
	;
	v78 = int32(3)
	goto L30
L29:
	;
	v78 = int32(2)
	goto L30
L30:
	;
	return v78
L31:
	;
	if v69 == int32(2) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v85 = int32(3)
	goto L34
L33:
	;
	v85 = int32(1)
	goto L34
L34:
	;
	v86 = v85
	goto L14
}
